package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "goal/api/admin/service/v1"
	"time"
)

type User struct {
	ID                  int64
	Address             string
	ToAddress           string
	ToAddressPrivateKey string
	CreatedAt           time.Time
}

type UserWithdraw struct {
	ID        int64
	UserId    int64
	Amount    int64
	Status    string
	Tx        string
	CreatedAt time.Time
}

type AddressEthBalance struct {
	ID      int64
	Address string
	Balance string
}

type Pagination struct {
	PageNum  int
	PageSize int
}

type UserRepo interface {
	GetUserList(ctx context.Context) ([]*User, error)
	GetUserById(ctx context.Context, userId int64) (*User, error)
	GetUserListByUserIds(ctx context.Context, userIds ...int64) ([]*User, error)
	GetUserMap(ctx context.Context, userIds ...int64) (map[int64]*User, error)
	GetUserProxyList(ctx context.Context, userId ...int64) ([]*UserProxy, error)
}

type UserUseCase struct {
	repo   UserRepo
	uiRepo UserInfoRepo
	ubRepo UserBalanceRepo
	tx     Transaction
	log    *log.Helper
}

func NewUserUseCase(repo UserRepo, tx Transaction, uiRepo UserInfoRepo, ubRepo UserBalanceRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo:   repo,
		tx:     tx,
		uiRepo: uiRepo,
		ubRepo: ubRepo,
		log:    log.NewHelper(logger),
	}
}

func (u *UserUseCase) GetUserList(ctx context.Context) ([]*User, error) {
	return u.repo.GetUserList(ctx)
}

func (u *UserUseCase) GetUserById(ctx context.Context, userId int64) (*User, error) {
	return u.repo.GetUserById(ctx, userId)
}

func (u *UserUseCase) GetUsers(ctx context.Context) (*v1.GetUserListReply, error) {
	var (
		user []*User
		err  error
	)

	user, err = u.repo.GetUserList(ctx)
	if err != nil {
		return nil, err
	}

	res := &v1.GetUserListReply{
		Items: make([]*v1.GetUserListReply_Item, 0),
	}

	for _, item := range user {
		res.Items = append(res.Items, &v1.GetUserListReply_Item{
			UserId:    item.ID,
			Address:   item.Address,
			ToAddress: item.ToAddress,
			CreatedAt: item.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return res, nil
}

func (u *UserUseCase) UpdateUserBalanceRecord(ctx context.Context, req *v1.UpdateUserBalanceRecordRequest) (*v1.UpdateUserBalanceRecordReply, error) {
	var base int64 = 100000 // 基础精度0.00001 todo 加配置文件
	_, err := u.ubRepo.UpdateUserBalance(ctx, req.SendBody.UserId, req.SendBody.Amount*base)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateUserBalanceRecordReply{
		Result: "成功",
	}, nil
}
func (u *UserUseCase) GetUserDepositList(ctx context.Context, req *v1.GetUserDepositListRequest) (*v1.GetUserDepositListReply, error) {
	var (
		user              map[int64]*User
		userBalanceRecord []*UserBalanceRecord
		base              int64 = 100000 // 基础精度0.00001 todo 加配置文件
		userId            []int64
		err               error
		count             int64
	)

	userBalanceRecord, err, count = u.ubRepo.GetUserBalanceRecord(ctx, "user_deposit", &Pagination{
		PageNum:  int(req.Page),
		PageSize: 10,
	})
	if err != nil {
		return nil, err
	}

	for _, v := range userBalanceRecord {
		userId = append(userId, v.UserId)
	}

	user, err = u.repo.GetUserMap(ctx, userId...)
	if err != nil {
		return nil, err
	}

	res := &v1.GetUserDepositListReply{
		Count: count,
		Items: make([]*v1.GetUserDepositListReply_Item, 0),
	}

	for _, item := range userBalanceRecord {
		tempAddress := ""
		if v, ok := user[item.UserId]; ok {
			tempAddress = v.Address
		}
		res.Items = append(res.Items, &v1.GetUserDepositListReply_Item{
			Address:   tempAddress,
			Balance:   item.Balance / base,
			Type:      item.Type,
			Amount:    item.Amount / base,
			Reason:    item.Reason,
			CreatedAt: item.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return res, nil
}

func (u *UserUseCase) GetUserBalanceRecord(ctx context.Context, req *v1.GetUserBalanceRecordRequest) (*v1.GetUserBalanceRecordReply, error) {
	var (
		user              map[int64]*User
		userBalanceRecord []*UserBalanceRecord
		base              int64 = 100000 // 基础精度0.00001 todo 加配置文件
		userId            []int64
		err               error
		count             int64
	)

	userBalanceRecord, err, count = u.ubRepo.GetUserBalanceRecord(ctx, req.Reason, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 10,
	})
	if err != nil {
		return nil, err
	}

	for _, v := range userBalanceRecord {
		userId = append(userId, v.UserId)
	}

	user, err = u.repo.GetUserMap(ctx, userId...)
	if err != nil {
		return nil, err
	}

	res := &v1.GetUserBalanceRecordReply{
		Count: count,
		Items: make([]*v1.GetUserBalanceRecordReply_Item, 0),
	}

	for _, item := range userBalanceRecord {
		tempAddress := ""
		if v, ok := user[item.UserId]; ok {
			tempAddress = v.Address
		}
		res.Items = append(res.Items, &v1.GetUserBalanceRecordReply_Item{
			Address:   tempAddress,
			Balance:   item.Balance / base,
			Type:      item.Type,
			Amount:    item.Amount / base,
			Reason:    item.Reason,
			CreatedAt: item.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return res, nil
}

func (u *UserUseCase) GetUserProxyList(ctx context.Context, req *v1.GetUserProxyListRequest) (*v1.GetUserProxyListReply, error) {
	var (
		user      map[int64]*User
		userProxy []*UserProxy
		userId    []int64
		err       error
	)

	userProxy, err = u.repo.GetUserProxyList(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	for _, v := range userProxy {
		userId = append(userId, v.UserId)
	}

	user, err = u.repo.GetUserMap(ctx, userId...)
	if err != nil {
		return nil, err
	}

	res := &v1.GetUserProxyListReply{
		Items: make([]*v1.GetUserProxyListReply_Item, 0),
	}

	for _, item := range userProxy {
		tempAddress := ""
		if v, ok := user[item.UserId]; ok {
			tempAddress = v.Address
		}
		res.Items = append(res.Items, &v1.GetUserProxyListReply_Item{
			Address:   tempAddress,
			Rate:      item.Rate,
			CreatedAt: item.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return res, nil
}

func (u *UserUseCase) GetUserWithById(ctx context.Context, id int64) (*UserWithdraw, error) {
	return u.ubRepo.WithdrawById(ctx, id)
}

func (u *UserUseCase) GetUserWithdrawList(ctx context.Context, req *v1.GetUserWithdrawListRequest) (*v1.GetUserWithdrawListReply, error) {
	var (
		userMap      map[int64]*User
		userWithdraw []*UserWithdraw
		userId       []int64
		err          error
		base         int64 = 100000
		count        int64
	)

	userWithdraw, err, count = u.ubRepo.WithdrawList(ctx, req.Status, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 10,
	})

	if err != nil {
		return nil, err
	}

	for _, v := range userWithdraw {
		userId = append(userId, v.UserId)
	}

	userMap, err = u.repo.GetUserMap(ctx, userId...)
	if err != nil {
		return nil, err
	}

	res := &v1.GetUserWithdrawListReply{
		Count: count,
		Items: make([]*v1.GetUserWithdrawListReply_Item, 0),
	}

	for _, item := range userWithdraw {
		tempAddress := ""
		if v, ok := userMap[item.UserId]; ok {
			tempAddress = v.Address
		}
		res.Items = append(res.Items, &v1.GetUserWithdrawListReply_Item{
			Address:   tempAddress,
			Status:    item.Status,
			ID:        item.ID,
			Amount:    item.Amount / base,
			Tx:        item.Tx,
			CreatedAt: item.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return res, nil
}

func (u *UserUseCase) UserWithdraw(ctx context.Context, withdraw *UserWithdraw, user *User) (bool, error) {
	var (
		err error
	)

	if err = u.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = u.ubRepo.Withdraw(ctx, user.ID, withdraw.Amount)
		if nil != err {
			return err
		}
		err = u.ubRepo.UpdateWithdraw(ctx, withdraw.ID, "pass", "")
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		return false, err
	}

	return true, nil
}

func (u *UserUseCase) UserWithdrawSuccess(ctx context.Context, withdraw *UserWithdraw, tx string) (bool, error) {
	err := u.ubRepo.UpdateWithdraw(ctx, withdraw.ID, "success", tx)
	if nil != err {
		return false, err
	}

	return true, nil
}

func (u *UserUseCase) UserWithdrawFail(ctx context.Context, withdraw *UserWithdraw, tx string) (bool, error) {
	err := u.ubRepo.UpdateWithdraw(ctx, withdraw.ID, "fail", tx)
	if nil != err {
		return false, err
	}

	return true, nil
}

func (u *UserUseCase) GetUserRecommendList(ctx context.Context, req *v1.GetUserRecommendListRequest) (*v1.GetUserRecommendListReply, error) {
	var (
		user              map[int64]*User
		userInfo          *UserInfo
		recommendUserInfo []*UserInfo
		recommendUserIds  []int64
		err               error
	)

	userInfo, err = u.uiRepo.GetUserInfoByUserId(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	recommendUserInfo, err = u.uiRepo.GetUserInfoListByRecommendCode(ctx, userInfo.MyRecommendCode)
	for _, v := range recommendUserInfo {
		recommendUserIds = append(recommendUserIds, v.UserId)
	}

	user, err = u.repo.GetUserMap(ctx, recommendUserIds...)
	if err != nil {
		return nil, err
	}

	res := &v1.GetUserRecommendListReply{
		Items: make([]*v1.GetUserRecommendListReply_Item, 0),
	}

	for _, item := range recommendUserInfo {
		tempAddress := ""
		if v, ok := user[item.UserId]; ok {
			tempAddress = v.Address
		}
		res.Items = append(res.Items, &v1.GetUserRecommendListReply_Item{
			Address:   tempAddress,
			UserId:    item.ID,
			CreatedAt: item.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return res, nil
}

func (u *UserUseCase) GetUserInfo(ctx context.Context, req *v1.GetUserRequest) (*v1.GetUserReply, error) {
	var base int64 = 100000 // 基础精度0.00001 todo 加配置文件
	userBalance, err := u.ubRepo.GetUserBalance(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return &v1.GetUserReply{UserBalance: userBalance.Balance / base}, nil
}

//func (u *UserUseCase) GetAddressEthBalance(ctx context.Context, address string) (*AddressEthBalance, error) {
//	return u.ubRepo.GetAddressEthBalanceByAddress(ctx, address)
//}
