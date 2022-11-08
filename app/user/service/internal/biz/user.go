package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "goal/api/user/service/v1"
)

type User struct {
	ID      int64
	Address string
}

type UserInfo struct {
	ID              int64
	UserId          int64
	Name            string
	Avatar          string
	RecommendCode   string
	MyRecommendCode string
	Code            string
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUserByAddress(ctx context.Context, address string) (*User, error)
	GetUserById(ctx context.Context, Id int64) (*User, error)
}

type UserInfoRepo interface {
	GetUserInfoByUserId(ctx context.Context, userId int64) (*UserInfo, error)
	CreateUserInfo(ctx context.Context, u *User, recommendCode string) (*UserInfo, error)
	GetUserInfoByMyRecommendCode(ctx context.Context, myRecommendCode string) (*UserInfo, error)
	GetUserInfoListByRecommendCode(ctx context.Context, recommendCode string) ([]*UserInfo, error)
}

type UserUseCase struct {
	repo   UserRepo
	rRepo  RoleRepo
	uiRepo UserInfoRepo
	ubRepo UserBalanceRepo
	tx     Transaction
	log    *log.Helper
}

func NewUserUseCase(repo UserRepo, tx Transaction, rRepo RoleRepo, uiRepo UserInfoRepo, ubRepo UserBalanceRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo:   repo,
		tx:     tx,
		rRepo:  rRepo,
		uiRepo: uiRepo,
		ubRepo: ubRepo,
		log:    log.NewHelper(logger),
	}
}

func (uc *UserUseCase) EthAuthorize(ctx context.Context, u *User, req *v1.EthAuthorizeRequest) (*User, error) {
	var (
		user        *User
		userInfo    *UserInfo
		userBalance *UserBalance
		err         error
	)

	recommendCode := req.SendBody.Code // 查询推荐码
	if 0 != len(recommendCode) {
		userInfo, err = uc.uiRepo.GetUserInfoByMyRecommendCode(ctx, recommendCode)
		if err != nil {
			return nil, errors.New(500, "USER_ERROR", "无效的推荐码")
		}
	}

	user, err = uc.repo.GetUserByAddress(ctx, u.Address) // 查询用户
	if err != nil {
		err = uc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			user, err = uc.repo.CreateUser(ctx, u) // 用户创建
			if err != nil {
				return err
			}

			userInfo, err = uc.uiRepo.CreateUserInfo(ctx, user, recommendCode) // 创建用户信息
			if err != nil {
				return err
			}

			userBalance, err = uc.ubRepo.CreateUserBalance(ctx, user) // 创建余额信息
			if err != nil {
				return err
			}

			return nil
		})
	}

	return user, nil
}

func (uc *UserUseCase) GetUserWithInfoAndBalance(ctx context.Context, u *User) (*v1.GetUserReply, error) {
	var (
		user        *User
		userInfo    *UserInfo
		userBalance *UserBalance
		base        int64 = 100000 // 基础精度0.00001 todo 加配置文件
		err         error
	)

	user, err = uc.repo.GetUserById(ctx, u.ID)
	if err != nil {
		return nil, err
	}

	userInfo, err = uc.uiRepo.GetUserInfoByUserId(ctx, u.ID)
	if err != nil {
		return nil, err
	}

	userBalance, err = uc.ubRepo.GetUserBalanceByUserId(ctx, u.ID)
	if err != nil {
		return nil, err
	}

	return &v1.GetUserReply{
		Address:         user.Address,
		Balance:         userBalance.Balance / base,
		Avatar:          userInfo.Avatar,
		MyRecommendCode: userInfo.MyRecommendCode,
	}, nil
}

func (uc *UserUseCase) Deposit(ctx context.Context, u *User, req *v1.DepositRequest) (*v1.DepositReply, error) {
	var (
		userBalance *UserBalance
		base        int64 = 100000 // 基础精度0.00001 todo 加配置文件
		err         error
	)

	// todo 以太坊
	amount := req.SendBody.Amount * base                         // 系统的数值
	if err = uc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务

		userBalance, err = uc.ubRepo.Deposit(ctx, u.ID, amount)
		if err != nil {
			return err
		}

		return nil
	}); nil != err {
		return nil, err
	}

	return &v1.DepositReply{
		Balance: userBalance.Balance / base,
	}, nil
}

func (uc *UserUseCase) GetUserRecommendList(ctx context.Context, u *User, req *v1.GetUserRecommendListRequest) (*v1.GetUserRecommendListReply, error) {
	var (
		userInfo          *UserInfo
		recommendUserInfo []*UserInfo
		userBalanceRecord []*UserBalanceRecord
		base              int64 = 100000 // 基础精度0.00001 todo 加配置文件
		err               error
	)

	userInfo, err = uc.uiRepo.GetUserInfoByUserId(ctx, u.ID)
	if err != nil {
		return nil, err
	}

	recommendUserInfo, _ = uc.uiRepo.GetUserInfoListByRecommendCode(ctx, userInfo.MyRecommendCode)
	userBalanceRecord, _ = uc.ubRepo.GetUserBalanceRecordByUserId(ctx, u.ID, "transfer_into", "recommend_user_goal_reward")

	res := &v1.GetUserRecommendListReply{
		Records: make([]*v1.GetUserRecommendListReply_Record, 0),
	}

	for _, v := range userBalanceRecord {
		tmpAmount := v.Amount / base
		res.RewardCount += tmpAmount
		res.Records = append(res.Records, &v1.GetUserRecommendListReply_Record{
			Amount:    tmpAmount,
			CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	for _, v := range recommendUserInfo {
		res.UserCount += 1
		res.UserInfos = append(res.UserInfos, &v1.GetUserRecommendListReply_UserInfo{
			Name: v.Name,
		})
	}

	return res, nil
}
