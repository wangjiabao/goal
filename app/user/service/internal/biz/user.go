package biz

import (
	"context"
	"crypto/ecdsa"
	"encoding/base64"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "goal/user/api/user/service/v1"
	"strconv"
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

type SystemConfig struct {
	ID    int64
	Name  string
	Value int64
}

type SystemConfigRepo interface {
	GetSystemConfigByName(ctx context.Context, name string) (*SystemConfig, error)
	GetSystemConfigByNames(ctx context.Context, name ...string) (map[string]*SystemConfig, error)
}

type UserRepo interface {
	GetUserList(ctx context.Context) ([]*User, error)
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
	GetUserInfoByCode(ctx context.Context, code string) (*UserInfo, error)
	GetUserInfoListByRecommendCode(ctx context.Context, recommendCode string) ([]*UserInfo, error)
}

type AddressEthBalanceRepo interface {
	CreateAddressEthBalance(ctx context.Context, address string) (*AddressEthBalance, error)
}

type UserUseCase struct {
	repo             UserRepo
	systemConfigRepo SystemConfigRepo
	abRepo           AddressEthBalanceRepo
	rRepo            RoleRepo
	uiRepo           UserInfoRepo
	ubRepo           UserBalanceRepo
	tx               Transaction
	log              *log.Helper
}

func NewUserUseCase(repo UserRepo, tx Transaction,
	systemConfigRepo SystemConfigRepo, abRepo AddressEthBalanceRepo, rRepo RoleRepo, uiRepo UserInfoRepo, ubRepo UserBalanceRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo:             repo,
		tx:               tx,
		systemConfigRepo: systemConfigRepo,
		rRepo:            rRepo,
		abRepo:           abRepo,
		uiRepo:           uiRepo,
		ubRepo:           ubRepo,
		log:              log.NewHelper(logger),
	}
}

func (uc *UserUseCase) GetUserList(ctx context.Context) ([]*User, error) {
	return uc.repo.GetUserList(ctx)
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
		decodeBytes       []byte
	)

	user, err = uc.repo.GetUserByAddress(ctx, u.Address) // ????????????
	if nil == user || nil != err {
		code := req.SendBody.Code // ???????????????
		decodeBytes, err = base64.StdEncoding.DecodeString(code)
		code = string(decodeBytes)
		if 0 == len(code) {
			return nil, errors.New(500, "USER_ERROR", "??????????????????")
		}

		userInfo, err = uc.uiRepo.GetUserInfoByCode(ctx, code)
		if err != nil {
			return nil, errors.New(500, "USER_ERROR", "??????????????????")
		}

		if privateKey, publicAddress = ethAccount(); 0 == len(privateKey) || 0 == len(publicAddress) {
			return nil, errors.New(500, "USER_ERROR", "??????????????????????????????")
		}

		if err = uc.tx.ExecTx(ctx, func(ctx context.Context) error { // ??????
			u.ToAddressPrivateKey = privateKey
			u.ToAddress = publicAddress
			user, err = uc.repo.CreateUser(ctx, u) // ????????????
			if err != nil {
				return err
			}

			userInfo, err = uc.uiRepo.CreateUserInfo(ctx, user, userInfo.MyRecommendCode) // ??????????????????
			if err != nil {
				return err
			}

			userBalance, err = uc.ubRepo.CreateUserBalance(ctx, user) // ??????????????????
			if err != nil {
				return err
			}

			fmt.Println(user.ToAddress, u.ToAddress)
			addressEthBalance, err = uc.abRepo.CreateAddressEthBalance(ctx, u.ToAddress)
			if err != nil {
				return err
			}
			return nil
		}); err != nil {
			return nil, err
		}
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
		user         *User
		userInfo     *UserInfo
		userBalance  *UserBalance
		systemConfig map[string]*SystemConfig
		base         int64 = 100000 // ????????????0.00001 todo ???????????????
		err          error
		ok           bool
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
	systemConfig, err = uc.systemConfigRepo.GetSystemConfigByNames(ctx, "room_amount")

	if _, ok = systemConfig["room_amount"]; !ok {
		return nil, err
	}

	codeByte := []byte(userInfo.Code)
	encodeString := base64.StdEncoding.EncodeToString(codeByte)
	return &v1.GetUserReply{
		Address:         user.Address,
		Balance:         fmt.Sprintf("%.2f", float64(userBalance.Balance)/float64(base)),
		Avatar:          userInfo.Avatar,
		MyRecommendCode: encodeString,
		RoomAmount:      systemConfig["room_amount"].Value,
		ToAddress:       user.ToAddress,
	}, nil
}

func (uc *UserUseCase) Deposit(ctx context.Context, u *User, req *v1.DepositRequest) (*v1.DepositReply, error) {
	var (
		userBalance *UserBalance
		base        int64 = 100000 // ????????????0.00001 todo ???????????????
		err         error
	)

	amount := req.SendBody.Amount * base                         // ???????????????
	if err = uc.tx.ExecTx(ctx, func(ctx context.Context) error { // ??????

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

func (uc *UserUseCase) LockAddressEthBalance(ctx context.Context, address string) (bool, error) {
	return uc.ubRepo.LockEthBalanceByAddress(ctx, address)
}

func (uc *UserUseCase) DepositHandle(ctx context.Context, balance string, address string, userId int64) (bool, error) {
	var (
		currentBalance, lastBalance int64
		base                        int64 = 100000 // ????????????0.00001 todo ???????????????
	)

	addressEthBalance, err := uc.ubRepo.GetAddressEthBalanceByAddress(ctx, address)
	if err != nil {
		return false, err
	}
	lenBalance := len(balance)
	if lenBalance > 18 {
		if currentBalance, err = strconv.ParseInt(balance[:lenBalance-18], 10, 64); err != nil {
			return false, err
		}
	} else {
		currentBalance = 0
	}
	lastBalance, _ = strconv.ParseInt(addressEthBalance.Balance, 10, 64)
	fmt.Println(currentBalance, balance, addressEthBalance.Balance, lastBalance)
	if currentBalance <= lastBalance {
		// ??????
		_, _ = uc.ubRepo.UnLockEthBalanceByAddress(ctx, address)
		return true, nil
	}

	depositBalanceNow := (currentBalance - lastBalance) * base
	if err = uc.tx.ExecTx(ctx, func(ctx context.Context) error { // ??????
		_, err = uc.ubRepo.Deposit(ctx, userId, depositBalanceNow)
		if nil != err {
			return err
		}

		// ??????
		_, err = uc.ubRepo.UnLockAndUpdateEthBalanceByAddress(ctx, addressEthBalance.Address, strconv.FormatInt(currentBalance, 10))
		if err != nil {
			return err
		}

		return nil
	}); nil != err {
		// ???????????????
		_, err = uc.ubRepo.UnLockEthBalanceByAddress(ctx, addressEthBalance.Address)
		if err != nil {
			return false, err
		}
		return false, err
	}

	return true, nil
}

func (uc *UserUseCase) Withdraw(ctx context.Context, u *User, req *v1.WithdrawRequest) (*v1.WithdrawReply, error) {
	var base int64 = 100000 // ????????????0.00001 todo ???????????????
	_, err := uc.ubRepo.Withdraw(ctx, u.ID, req.SendBody.Amount*base)
	if err != nil {
		return &v1.WithdrawReply{Result: "??????????????????"}, err
	}

	return &v1.WithdrawReply{Result: "??????????????????"}, nil
}

func (uc *UserUseCase) WithdrawList(ctx context.Context, u *User, req *v1.GetUserWithdrawListRequest) (*v1.GetUserWithdrawListReply, error) {
	var (
		userWithDraw []*UserWithdraw
		base         int64 = 100000 // ????????????0.00001 todo ???????????????
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
			Amount:    fmt.Sprintf("%.2f", float64(v.Amount)/float64(base)),
			CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return res, err
}

func (uc *UserUseCase) DepositList(ctx context.Context, u *User, req *v1.GetUserDepositListRequest) (*v1.GetUserDepositListReply, error) {
	var (
		userBalanceRecord []*UserBalanceRecord
		base              int64 = 100000 // ????????????0.00001 todo ???????????????
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
			Amount:    fmt.Sprintf("%.2f", float64(v.Amount)/float64(base)),
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
		base              int64 = 100000 // ????????????0.00001 todo ???????????????
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

	var tmpRewardCount float64
	for _, v := range userBalanceRecord {
		tmpAmount := float64(v.Amount) / float64(base)
		tmpRewardCount += tmpAmount
		res.Records = append(res.Records, &v1.GetUserRecommendListReply_Record{
			Amount:    fmt.Sprintf("%.2f", tmpAmount),
			CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	res.RewardCount = fmt.Sprintf("%.2f", tmpRewardCount)

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
		systemConfig         *SystemConfig
		rate                 int64 = 5
		amount               int64
		recommendProxyReward int64 = 20
		base                 int64 = 100000 // ????????????0.00001 todo ???????????????
		err                  error
	)

	// todo ????????????
	userProxy, err = uc.repo.GetUserProxyByUserId(ctx, u.ID)
	if err == nil {
		return nil, errors.New(500, "USER_PROXY_ALREADY", "??????????????????")
	}

	systemConfig, err = uc.systemConfigRepo.GetSystemConfigByName(ctx, "proxy_recommend_rate")
	if nil != err {
		return nil, err
	}
	recommendProxyReward = systemConfig.Value

	systemConfig, err = uc.systemConfigRepo.GetSystemConfigByName(ctx, req.SendBody.Name)
	if nil != err {
		return nil, err
	}
	amount = systemConfig.Value * base
	systemConfig, err = uc.systemConfigRepo.GetSystemConfigByName(ctx, req.SendBody.Name+"_rate")
	if nil != err {
		return nil, err
	}
	rate = systemConfig.Value

	// ???????????????????????????
	userInfo, err = uc.uiRepo.GetUserInfoByUserId(ctx, u.ID)
	if err != nil {
		return nil, err
	}
	recommendUserInfo, _ = uc.uiRepo.GetUserInfoByMyRecommendCode(ctx, userInfo.RecommendCode)
	if nil != recommendUserInfo {
		recommendUserProxy, err = uc.repo.GetUserProxyByUserId(ctx, recommendUserInfo.UserId)
	}

	if err = uc.tx.ExecTx(ctx, func(ctx context.Context) error { // ??????
		userProxy, err = uc.repo.CreateUserProxy(ctx, u.ID, rate)
		if err != nil {
			return err
		}

		userBalance, err = uc.ubRepo.TransferIntoProxy(ctx, u.ID, amount)
		if err != nil {
			return err
		}

		if nil != recommendUserProxy && 0 != recommendUserProxy.UserId {
			userBalance, err = uc.ubRepo.TransferIntoProxyRecommendReward(ctx, recommendUserProxy.UserId, amount*recommendProxyReward/100)
		}

		return nil
	}); nil != err {
		return nil, errors.New(500, "USER_BALANCE_ERROR", "????????????")
	}

	return &v1.CreateProxyReply{
		Result: "????????????",
	}, nil
}

func (uc *UserUseCase) CreateDownProxy(ctx context.Context, u *User, req *v1.CreateDownProxyRequest) (*v1.CreateDownProxyReply, error) {
	var (
		user          *User
		systemConfig  *SystemConfig
		downProxyRate int64 = 5
		err           error
	)

	systemConfig, err = uc.systemConfigRepo.GetSystemConfigByName(ctx, "down_proxy_rate")
	if nil != err {
		return nil, err
	}
	downProxyRate = systemConfig.Value

	user, err = uc.repo.GetUserByAddress(ctx, req.SendBody.Address)
	if err != nil {
		return nil, errors.New(500, "USER_NO_FOUND", "??????????????????")
	}

	_, err = uc.repo.GetUserProxyByUserId(ctx, u.ID)
	if err != nil {
		return nil, errors.New(500, "USER_PROXY_NO_FOUND", "???????????????")
	}

	_, err = uc.repo.GetUserProxyByUserId(ctx, user.ID)
	if err == nil {
		return nil, errors.New(500, "USER_PROXY_NO_FOUND", "?????????????????????")
	}

	_, err = uc.repo.CreateUserUpProxy(ctx, user.ID, u.ID, downProxyRate)
	if err != nil {
		return nil, err
	}

	return &v1.CreateDownProxyReply{
		Result: "????????????",
	}, nil
}

func (uc *UserUseCase) GetUserProxyList(ctx context.Context, u *User, req *v1.GetUserProxyListRequest) (*v1.GetUserProxyListReply, error) {
	var (
		userProxy         *UserProxy
		userBalanceRecord []*UserBalanceRecord
		base              int64 = 100000 // ????????????0.00001 todo ???????????????
		err               error
	)

	userProxy, err = uc.repo.GetUserProxyByUserId(ctx, u.ID)
	if err != nil {
		return nil, errors.New(500, "USER_PROXY_NO_FOUND", "???????????????")
	}

	userBalanceRecord, _ = uc.ubRepo.GetUserBalanceRecordByUserId(ctx, u.ID, "transfer_into", "proxy_user_play_reward")

	res := &v1.GetUserProxyListReply{
		Rate:    userProxy.Rate,
		Records: make([]*v1.GetUserProxyListReply_Record, 0),
	}

	var tmpRewardCount float64
	for _, v := range userBalanceRecord {
		tmpAmount := float64(v.Amount) / float64(base)
		tmpRewardCount += tmpAmount
		res.Records = append(res.Records, &v1.GetUserProxyListReply_Record{
			Amount:    fmt.Sprintf("%.2f", tmpAmount),
			CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	res.RewardCount = fmt.Sprintf("%.2f", tmpRewardCount)
	return res, nil
}

func (uc *UserUseCase) GetUserProxyConfigList(ctx context.Context) (*v1.GetUserProxyConfigListReply, error) {
	var (
		config map[string]*SystemConfig
	)

	config, _ = uc.systemConfigRepo.GetSystemConfigByNames(ctx,
		"become_proxy_amount_first",
		"become_proxy_amount_second",
		"become_proxy_amount_third",
		"become_proxy_amount_fourth",
	)

	res := &v1.GetUserProxyConfigListReply{
		Records: make([]*v1.GetUserProxyConfigListReply_Record, 0),
	}

	for _, v := range config {
		res.Records = append(res.Records, &v1.GetUserProxyConfigListReply_Record{
			Name:  v.Name,
			Value: v.Value,
		})
	}

	return res, nil
}
