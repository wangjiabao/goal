package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
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

type UserBalanceRecordTotal struct {
	Total int64
}

type UserBalanceTotal struct {
	Total int64
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
	Status  int64
}

type Pagination struct {
	PageNum  int
	PageSize int
}

type UserRepo interface {
	CreateUserProxy(ctx context.Context, userId int64, rate int64) (*UserProxy, error)
	CreateDownUserProxy(ctx context.Context, userId int64, upUserId int64, rate int64) (*UserProxy, error)
	UpdateUserProxy(ctx context.Context, userId int64, rate int64) (*UserProxy, error)
	GetUserProxyByUserId(ctx context.Context, userId int64) (*UserProxy, error)
	GetUserList(ctx context.Context, address string, b *Pagination) ([]*User, error, int64)
	GetUserById(ctx context.Context, userId int64) (*User, error)
	GetUserByAddress(ctx context.Context, address string) (*User, error)
	GetUserListByUserIds(ctx context.Context, userIds ...int64) ([]*User, error)
	GetUserMap(ctx context.Context, userIds ...int64) (map[int64]*User, error)
	GetUserProxyList(ctx context.Context, userId ...int64) ([]*UserProxy, error)
}

type UserUseCase struct {
	repo             UserRepo
	uiRepo           UserInfoRepo
	ubRepo           UserBalanceRepo
	systemConfigRepo SystemConfigRepo
	tx               Transaction
	log              *log.Helper
}

func NewUserUseCase(repo UserRepo, tx Transaction, uiRepo UserInfoRepo, systemConfigRepo SystemConfigRepo, ubRepo UserBalanceRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo:             repo,
		systemConfigRepo: systemConfigRepo,
		tx:               tx,
		uiRepo:           uiRepo,
		ubRepo:           ubRepo,
		log:              log.NewHelper(logger),
	}
}

func (u *UserUseCase) GetUserById(ctx context.Context, userId int64) (*User, error) {
	return u.repo.GetUserById(ctx, userId)
}

func (u *UserUseCase) GetUsers(ctx context.Context, req *v1.GetUserListRequest) (*v1.GetUserListReply, error) {
	var (
		user  []*User
		count int64
		err   error
	)

	user, err, count = u.repo.GetUserList(ctx, req.Address, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 10,
	})
	if err != nil {
		return nil, err
	}

	res := &v1.GetUserListReply{
		Count: count,
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
	amount := int64(req.SendBody.Amount * base)
	_, err := u.ubRepo.UpdateUserBalance(ctx, req.SendBody.UserId, amount)
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
		userSearch        []*User
		searchUserId      []int64
		userBalanceRecord []*UserBalanceRecord
		base              int64 = 100000 // 基础精度0.00001 todo 加配置文件
		userId            []int64
		err               error
		count             int64
	)

	if "" != req.Address {
		userSearch, err, _ = u.repo.GetUserList(ctx, req.Address, &Pagination{
			PageNum:  int(req.Page),
			PageSize: 10,
		})
		if err != nil {
			return nil, err
		}
		for _, v := range userSearch {
			searchUserId = append(searchUserId, v.ID)
		}
	}

	userBalanceRecord, err, count = u.ubRepo.GetUserBalanceRecord(ctx, "user_deposit", &Pagination{
		PageNum:  int(req.Page),
		PageSize: 10,
	}, searchUserId...)
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
			Balance:   fmt.Sprintf("%.2f", float64(item.Balance)/float64(base)),
			Type:      item.Type,
			Amount:    fmt.Sprintf("%.2f", float64(item.Amount)/float64(base)),
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
			Balance:   fmt.Sprintf("%.2f", float64(item.Balance)/float64(base)),
			Type:      item.Type,
			Amount:    fmt.Sprintf("%.2f", float64(item.Amount)/float64(base)),
			Reason:    item.Reason,
			CreatedAt: item.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return res, nil
}

func (u *UserUseCase) UserBalanceRecordTotal(ctx context.Context, req *v1.UserBalanceRecordTotalRequest) (*v1.UserBalanceRecordTotalReply, error) {
	var (
		todayDeposit  *UserBalanceRecordTotal
		totalDeposit  *UserBalanceRecordTotal
		todayWithdraw *UserBalanceRecordTotal
		totalWithdraw *UserBalanceRecordTotal
		totalBalance  *UserBalanceTotal
		base          int64 = 100000 // 基础精度0.00001 todo 加配置文件
		err           error
	)
	todayDeposit, err = u.ubRepo.GetUserBalanceRecordTotal(ctx, "deposit", true)
	if nil != err {
		return nil, err
	}

	totalDeposit, err = u.ubRepo.GetUserBalanceRecordTotal(ctx, "deposit", false)
	if nil != err {
		return nil, err
	}

	todayWithdraw, err = u.ubRepo.GetUserBalanceRecordTotal(ctx, "withdraw", true)
	if nil != err {
		return nil, err
	}

	totalWithdraw, err = u.ubRepo.GetUserBalanceRecordTotal(ctx, "withdraw", false)
	if nil != err {
		return nil, err
	}

	totalBalance, err = u.ubRepo.GetUserBalanceTotal(ctx)
	if nil != err {
		return nil, err
	}

	res := &v1.UserBalanceRecordTotalReply{
		TodayDeposit:  todayDeposit.Total / base,
		TotalDeposit:  totalDeposit.Total / base,
		TodayWithdraw: todayWithdraw.Total / base,
		TotalWithdraw: totalWithdraw.Total / base,
		TotalBalance:  totalBalance.Total / base,
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
			UserId:    item.UserId,
			Rate:      item.Rate,
			CreatedAt: item.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return res, nil
}

func (u *UserUseCase) GetUserWithById(ctx context.Context, id int64) (*UserWithdraw, error) {
	return u.ubRepo.WithdrawById(ctx, id)
}

func (u *UserUseCase) GetUserByToAddress(ctx context.Context, address string) (*User, error) {
	return u.ubRepo.GetUserByToAddress(ctx, address)
}

func (u *UserUseCase) GetUserWithdrawList(ctx context.Context, req *v1.GetUserWithdrawListRequest) (*v1.GetUserWithdrawListReply, error) {
	var (
		userMap      map[int64]*User
		userWithdraw []*UserWithdraw
		user         []*User
		userId       []int64
		searchUserId []int64
		err          error
		base         int64 = 100000
		count        int64
	)

	if "" != req.Address {
		user, err, _ = u.repo.GetUserList(ctx, req.Address, &Pagination{
			PageNum:  int(req.Page),
			PageSize: 10,
		})
		if err != nil {
			return nil, err
		}
		for _, v := range user {
			searchUserId = append(searchUserId, v.ID)
		}
	}
	userWithdraw, err, count = u.ubRepo.WithdrawList(ctx, req.Status, &Pagination{
		PageNum:  int(req.Page),
		PageSize: 10,
	}, searchUserId...)

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
			Amount:    fmt.Sprintf("%.2f", float64(item.Amount)/float64(base)),
			Tx:        item.Tx,
			CreatedAt: item.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return res, nil
}

func (u *UserUseCase) LockAddressEthBalance(ctx context.Context, address string) (bool, error) {
	return u.ubRepo.LockEthBalanceByAddress(ctx, address)
}

func (u *UserUseCase) UnLockAddressEthBalance(ctx context.Context, address string) (bool, error) {
	return u.ubRepo.UnLockEthBalanceByAddress(ctx, address)
}

func (u *UserUseCase) GetAddressEthBalance(ctx context.Context) ([]*AddressEthBalance, error) {
	return u.ubRepo.GetAddressEthBalance(ctx)
}

func (u *UserUseCase) UpdateAddressEthBalance(ctx context.Context, address string, balance string) (bool, error) {
	return u.ubRepo.UpdateEthBalanceByAddress(ctx, address, balance)
}

func (u *UserUseCase) UserWithdraw(ctx context.Context, withdraw *UserWithdraw, user *User) (bool, error) {
	var (
		err error
		//base              int64 = 100000 // 基础精度0.00001 todo 加配置文件
		//addressEthBalance *AddressEthBalance
		//lastBalance       int64
		//nowAmount         int64
	)
	//addressEthBalance, err = u.ubRepo.GetAddressEthBalanceByAddress(ctx, user.ToAddress)
	//if err != nil {
	//	return false, err
	//}
	//nowAmount = withdraw.Amount / base
	//lastBalance, _ = strconv.ParseInt(addressEthBalance.Balance, 10, 64)
	//if lastBalance < nowAmount {
	//	return false, errors.New(500, "BALANCE_ETH_ERROR", "余额不足eth")
	//}
	//lastBalance -= nowAmount

	if err = u.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = u.ubRepo.Withdraw(ctx, user.ID, withdraw.Amount)
		if nil != err {
			return err
		}
		err = u.ubRepo.UpdateWithdraw(ctx, withdraw.ID, "pass", "")
		if nil != err {
			return err
		}

		//_, err = u.ubRepo.UpdateEthBalanceByAddress(ctx, user.ToAddress, strconv.FormatInt(lastBalance, 10))
		//if nil != err {
		//	return err
		//}

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
	return &v1.GetUserReply{UserBalance: fmt.Sprintf("%.2f", float64(userBalance.Balance)/float64(base))}, nil
}

//func (u *UserUseCase) GetAddressEthBalance(ctx context.Context, address string) (*AddressEthBalance, error) {
//	return u.ubRepo.GetAddressEthBalanceByAddress(ctx, address)
//}

func (u *UserUseCase) CreateProxy(ctx context.Context, user *User, req *v1.CreateProxyRequest) (*v1.CreateProxyReply, error) {
	var (
		rate      int64 = 5
		err       error
		userProxy *UserProxy
	)

	rate = req.SendBody.Rate

	userProxy, err = u.repo.GetUserProxyByUserId(ctx, user.ID)
	if nil != userProxy && err == nil {
		_, err = u.repo.UpdateUserProxy(ctx, user.ID, rate)
	} else {
		_, err = u.repo.CreateUserProxy(ctx, user.ID, rate)
	}

	if err != nil {
		return nil, err
	}
	return &v1.CreateProxyReply{
		Result: "提交成功",
	}, nil
}

func (u *UserUseCase) CreateDownProxy(ctx context.Context, user *User, req *v1.CreateDownProxyRequest) (*v1.CreateDownProxyReply, error) {
	var (
		rate         int64 = 5
		err          error
		proxyUser    *User
		userProxy    *UserProxy
		ok           bool
		systemConfig map[string]*SystemConfig
	)
	proxyUser, err = u.repo.GetUserByAddress(ctx, req.SendBody.Address)
	if err != nil {
		return nil, errors.New(500, "USER_NO_FOUND", "用户地址有误")
	}

	systemConfig, err = u.systemConfigRepo.GetSystemConfigByNames(ctx, "down_proxy_rate")
	if _, ok = systemConfig["down_proxy_rate"]; !ok {
		return nil, errors.New(500, "USER_NO_FOUND", "配置有误")
	}
	rate = systemConfig["down_proxy_rate"].Value
	userProxy, err = u.repo.GetUserProxyByUserId(ctx, proxyUser.ID)
	if nil != userProxy && userProxy.UpUserId == user.ID {
		return nil, errors.New(500, "USER_NO_FOUND", "用户已是代理")
	}
	_, err = u.repo.CreateDownUserProxy(ctx, proxyUser.ID, user.ID, rate)
	if err != nil {
		return nil, err
	}
	return &v1.CreateDownProxyReply{
		Result: "提交成功",
	}, nil
}
