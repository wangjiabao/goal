package biz

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "goal/api/user/service/v1"
)

type User struct {
	ID                  int64
	Address             string
	ToAddress           string
	ToAddressPrivateKey string
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

type UserProxy struct {
	ID       int64
	UserId   int64
	UpUserId int64
	Rate     int64
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUserByAddress(ctx context.Context, address string) (*User, error)
	GetUserById(ctx context.Context, Id int64) (*User, error)
	GetUserListByIds(ctx context.Context, userIds ...int64) (map[int64]*User, error)
	CreateUserProxy(ctx context.Context, userId int64, rate int64) (*UserProxy, error)
	CreateUserUpProxy(ctx context.Context, userId int64, upUserId int64, rate int64) (*UserProxy, error)
	GetUserProxyByUserId(ctx context.Context, userId int64) (*UserProxy, error)
}

type UserInfoRepo interface {
	GetUserInfoByUserId(ctx context.Context, userId int64) (*UserInfo, error)
	CreateUserInfo(ctx context.Context, u *User, recommendCode string) (*UserInfo, error)
	GetUserInfoByMyRecommendCode(ctx context.Context, myRecommendCode string) (*UserInfo, error)
	GetUserInfoListByRecommendCode(ctx context.Context, recommendCode string) ([]*UserInfo, error)
}

type AddressEthBalanceRepo interface {
	CreateAddressEthBalance(ctx context.Context, address string) (*AddressEthBalance, error)
}

type UserUseCase struct {
	repo   UserRepo
	abRepo AddressEthBalanceRepo
	rRepo  RoleRepo
	uiRepo UserInfoRepo
	ubRepo UserBalanceRepo
	tx     Transaction
	log    *log.Helper
}

func NewUserUseCase(repo UserRepo, tx Transaction, abRepo AddressEthBalanceRepo, rRepo RoleRepo, uiRepo UserInfoRepo, ubRepo UserBalanceRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo:   repo,
		tx:     tx,
		rRepo:  rRepo,
		abRepo: abRepo,
		uiRepo: uiRepo,
		ubRepo: ubRepo,
		log:    log.NewHelper(logger),
	}
}

func (uc *UserUseCase) EthAuthorize(ctx context.Context, u *User, req *v1.EthAuthorizeRequest) (*User, error) {
	var (
		user              *User
		userInfo          *UserInfo
		userBalance       *UserBalance
		addressEthBalance *AddressEthBalance
		privateKey        string
		publicAddress     string
		err               error
	)

	recommendCode := req.SendBody.Code // 查询推荐码
	if 0 != len(recommendCode) {
		userInfo, err = uc.uiRepo.GetUserInfoByMyRecommendCode(ctx, recommendCode)
		if err != nil {
			return nil, errors.New(500, "USER_ERROR", "无效的推荐码")
		}
	}
	if privateKey, publicAddress = ethAccount(); 0 == len(privateKey) || 0 == len(publicAddress) {
		return nil, errors.New(500, "USER_ERROR", "生成账户失败，请重试")
	}

	user, err = uc.repo.GetUserByAddress(ctx, u.Address) // 查询用户
	if err != nil {
		err = uc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
			u.ToAddressPrivateKey = privateKey
			u.ToAddress = publicAddress
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

			addressEthBalance, err = uc.abRepo.CreateAddressEthBalance(ctx, user.ToAddress)
			if err != nil {
				return err
			}
			return nil
		})
	}

	return user, nil
}

func ethAccount() (string, string) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", ""
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return hexutil.Encode(privateKeyBytes)[2:], address
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
		ToAddress:       user.ToAddress,
	}, nil
}

func (uc *UserUseCase) Deposit(ctx context.Context, u *User, req *v1.DepositRequest) (*v1.DepositReply, error) {
	var (
		userBalance *UserBalance
		base        int64 = 100000 // 基础精度0.00001 todo 加配置文件
		err         error
	)

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

func (uc *UserUseCase) Withdraw(ctx context.Context, u *User, req *v1.WithdrawRequest) (*v1.WithdrawReply, error) {
	_, err := uc.ubRepo.Withdraw(ctx, u.ID, req.SendBody.Amount)
	if err != nil {
		return &v1.WithdrawReply{Result: "提交审核失败"}, err
	}

	return &v1.WithdrawReply{Result: "提交审核成功"}, nil
}

func (uc *UserUseCase) WithdrawList(ctx context.Context, u *User, req *v1.GetUserWithdrawListRequest) (*v1.GetUserWithdrawListReply, error) {
	var (
		userWithDraw []*UserWithdraw
		err          error
	)
	userWithDraw, err = uc.ubRepo.WithdrawList(ctx, u.ID)
	res := &v1.GetUserWithdrawListReply{
		Records: make([]*v1.GetUserWithdrawListReply_Record, 0),
	}
	if err != nil {
		return res, err
	}

	for _, v := range userWithDraw {
		res.Records = append(res.Records, &v1.GetUserWithdrawListReply_Record{
			Status:    v.Status,
			Amount:    v.Amount,
			CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return res, err
}

func (uc *UserUseCase) DepositList(ctx context.Context, u *User, req *v1.GetUserDepositListRequest) (*v1.GetUserDepositListReply, error) {
	var (
		userBalanceRecord []*UserBalanceRecord
		err               error
	)

	userBalanceRecord, _ = uc.ubRepo.GetUserBalanceRecordByUserId(ctx, u.ID, "deposit", "user_deposit")
	res := &v1.GetUserDepositListReply{
		Records: make([]*v1.GetUserDepositListReply_Record, 0),
	}
	if err != nil {
		return res, err
	}

	for _, v := range userBalanceRecord {
		res.Records = append(res.Records, &v1.GetUserDepositListReply_Record{
			Amount:    v.Amount,
			CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return res, err
}

func (uc *UserUseCase) GetUserRecommendList(ctx context.Context, u *User, req *v1.GetUserRecommendListRequest) (*v1.GetUserRecommendListReply, error) {
	var (
		userInfo          *UserInfo
		recommendUserInfo []*UserInfo
		userBalanceRecord []*UserBalanceRecord
		userId            []int64
		user              map[int64]*User
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
		userId = append(userId, v.UserId)
	}

	user, _ = uc.repo.GetUserListByIds(ctx, userId...)

	for _, v := range recommendUserInfo {
		res.UserCount += 1
		res.UserInfos = append(res.UserInfos, &v1.GetUserRecommendListReply_UserInfo{
			Name:    v.Name,
			Address: user[v.UserId].Address,
		})
	}

	return res, nil
}

func (uc *UserUseCase) CreateProxy(ctx context.Context, u *User, req *v1.CreateProxyRequest) (*v1.CreateProxyReply, error) {
	var (
		userProxy            *UserProxy
		recommendUserProxy   *UserProxy
		userBalance          *UserBalance
		userInfo             *UserInfo
		recommendUserInfo    *UserInfo
		rate                 int64 = 5
		amount               int64 = 500
		recommendProxyReward int64 = 20
		base                 int64 = 100000 // 基础精度0.00001 todo 加配置文件
		err                  error
	)

	// todo 后台比例
	if req.SendBody.Amount == amount {
		rate = 20
	}

	userProxy, err = uc.repo.GetUserProxyByUserId(ctx, u.ID)
	if err == nil {
		return nil, errors.New(500, "USER_PROXY_ALREADY", "已经是代理了")
	}

	// 查找上级是否大代理
	userInfo, err = uc.uiRepo.GetUserInfoByUserId(ctx, u.ID)
	if err != nil {
		return nil, err
	}
	recommendUserInfo, _ = uc.uiRepo.GetUserInfoByMyRecommendCode(ctx, userInfo.RecommendCode)
	if nil != recommendUserInfo {
		recommendUserProxy, err = uc.repo.GetUserProxyByUserId(ctx, recommendUserInfo.UserId)
	}

	if err = uc.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		userProxy, err = uc.repo.CreateUserProxy(ctx, u.ID, rate)
		if err != nil {
			return err
		}

		userBalance, err = uc.ubRepo.TransferIntoProxy(ctx, u.ID, amount*base)
		if err != nil {
			return err
		}

		if nil != recommendUserProxy && 0 != recommendUserProxy.UserId {
			userBalance, err = uc.ubRepo.TransferIntoProxyRecommendReward(ctx, recommendUserProxy.UserId, amount*base*recommendProxyReward/100)
		}

		return nil
	}); nil != err {
		return nil, err
	}

	return &v1.CreateProxyReply{
		Result: "提交成功",
	}, nil
}

func (uc *UserUseCase) CreateDownProxy(ctx context.Context, u *User, req *v1.CreateDownProxyRequest) (*v1.CreateDownProxyReply, error) {
	var (
		user *User
		err  error
	)

	user, err = uc.repo.GetUserByAddress(ctx, req.SendBody.Address)
	if err != nil {
		return nil, errors.New(500, "USER_NO_FOUND", "用户地址有误")
	}

	_, err = uc.repo.GetUserProxyByUserId(ctx, u.ID)
	if err != nil {
		return nil, errors.New(500, "USER_PROXY_NO_FOUND", "你不是代理")
	}

	_, err = uc.repo.GetUserProxyByUserId(ctx, user.ID)
	if err == nil {
		return nil, errors.New(500, "USER_PROXY_NO_FOUND", "用户已经是代理")
	}

	_, err = uc.repo.CreateUserUpProxy(ctx, user.ID, u.ID, req.SendBody.Rate)
	if err != nil {
		return nil, err
	}

	return &v1.CreateDownProxyReply{
		Result: "提交成功",
	}, nil
}

func (uc *UserUseCase) GetUserProxyList(ctx context.Context, u *User, req *v1.GetUserProxyListRequest) (*v1.GetUserProxyListReply, error) {
	var (
		userProxy         *UserProxy
		userBalanceRecord []*UserBalanceRecord
		base              int64 = 100000 // 基础精度0.00001 todo 加配置文件
		err               error
	)

	userProxy, err = uc.repo.GetUserProxyByUserId(ctx, u.ID)
	if err != nil {
		return nil, errors.New(500, "USER_PROXY_NO_FOUND", "你不是代理")
	}

	userBalanceRecord, _ = uc.ubRepo.GetUserBalanceRecordByUserId(ctx, u.ID, "transfer_into", "proxy_user_play_reward")

	res := &v1.GetUserProxyListReply{
		Rate:    userProxy.Rate,
		Records: make([]*v1.GetUserProxyListReply_Record, 0),
	}

	for _, v := range userBalanceRecord {
		tmpAmount := v.Amount / base
		res.RewardCount += tmpAmount
		res.Records = append(res.Records, &v1.GetUserProxyListReply_Record{
			Amount:    tmpAmount,
			CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return res, nil
}
