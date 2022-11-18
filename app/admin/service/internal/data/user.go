package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"goal/app/admin/service/internal/biz"
	"gorm.io/gorm"
	"time"
)

type UserBalance struct {
	ID        int64     `gorm:"primarykey;type:int"`
	UserId    int64     `gorm:"type:int;not null"`
	Balance   int64     `gorm:"type:bigint;not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type AddressEthBalance struct {
	ID        int64     `gorm:"primarykey;type:int"`
	Address   string    `gorm:"type:varchar(100);not null"`
	Balance   string    `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type User struct {
	ID                  int64     `gorm:"primarykey;type:int"`
	Address             string    `gorm:"type:varchar(100);not null"`
	ToAddress           string    `gorm:"type:varchar(100);not null"`
	ToAddressPrivateKey string    `gorm:"type:varchar(100);not null"`
	CreatedAt           time.Time `gorm:"type:datetime;not null"`
	UpdatedAt           time.Time `gorm:"type:datetime;not null"`
}

type Admin struct {
	ID        int64     `gorm:"primarykey;type:int"`
	Account   string    `gorm:"type:varchar(100);not null"`
	Password  string    `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type BalanceRecordIdRel struct {
	ID        int64     `gorm:"primarykey;type:int"`
	RecordId  int64     `gorm:"type:int;not null"`
	RelType   string    `gorm:"type:varchar(100)"`
	RelId     int64     `gorm:"type:int;not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type UserBalanceRecordTotal struct {
	Total int64
}

type UserBalanceTotal struct {
	Total int64
}

type UserBalanceRecord struct {
	ID        int64     `gorm:"primarykey;type:int"`
	UserId    int64     `gorm:"type:int;not null"`
	Balance   int64     `gorm:"type:bigint;not null"`
	Type      string    `gorm:"type:varchar(45);not null"`
	Reason    string    `gorm:"type:varchar(45);not null"`
	Amount    int64     `gorm:"type:int;not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type UserProxy struct {
	ID        int64     `gorm:"primarykey;type:int"`
	UserId    int64     `gorm:"type:int;not null"`
	UpUserId  int64     `gorm:"type:int;not null"`
	Rate      int64     `gorm:"type:int;not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type UserInfo struct {
	ID              int64     `gorm:"primarykey;type:int"`
	UserId          int64     `gorm:"type:int;not null"`
	Name            string    `gorm:"type:varchar(45)"`
	Avatar          string    `gorm:"type:varchar(45)"`
	RecommendCode   string    `gorm:"type:varchar(2000)"`
	MyRecommendCode string    `gorm:"type:varchar(2000)"`
	Code            string    `gorm:"type:varchar(45)"`
	CreatedAt       time.Time `gorm:"type:datetime;not null"`
	UpdatedAt       time.Time `gorm:"type:datetime;not null"`
}

type UserWithdraw struct {
	ID        int64     `gorm:"primarykey;type:int"`
	Amount    int64     `gorm:"type:bigint;not null"`
	UserId    int64     `gorm:"type:int;not null"`
	Tx        string    `gorm:"type:varchar(100);not null"`
	Status    string    `gorm:"type:varchar(45);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type UserBalanceRepo struct {
	data *Data
	log  *log.Helper
}

type UserProxyRepo struct {
	data *Data
	log  *log.Helper
}

type UserRepo struct {
	data *Data
	log  *log.Helper
}

type UserInfoRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserBalanceRepo(data *Data, logger log.Logger) biz.UserBalanceRepo {
	return &UserBalanceRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &UserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func NewUserProxyRepo(data *Data, logger log.Logger) biz.UserProxyRepo {
	return &UserProxyRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func NewUserInfoRepo(data *Data, logger log.Logger) biz.UserInfoRepo {
	return &UserInfoRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (up *UserProxyRepo) GetUserProxyAndDown(ctx context.Context) ([]*biz.UserProxy, map[int64][]*biz.UserProxy, error) {
	var l []*UserProxy
	if err := up.data.DB(ctx).Table("user_proxy").Find(&l).Error; err != nil {
		return nil, nil, errors.InternalServer("SELECT_PLAY_ERROR", "查询代理失败")
	}

	ul := make([]*biz.UserProxy, 0)
	dl := make(map[int64][]*biz.UserProxy, 0)
	for _, v := range l {
		if 0 != v.UpUserId {
			dl[v.UpUserId] = append(dl[v.UpUserId], &biz.UserProxy{
				ID:       v.ID,
				UserId:   v.UserId,
				UpUserId: v.UpUserId,
				Rate:     v.Rate,
			})
			continue
		}

		ul = append(ul, &biz.UserProxy{
			ID:       v.ID,
			UserId:   v.UserId,
			UpUserId: v.UpUserId,
			Rate:     v.Rate,
		})
	}
	return ul, dl, nil
}

func (ub *UserBalanceRepo) GetUserBalance(ctx context.Context, userId int64) (*biz.UserBalance, error) {
	var userBalance UserBalance
	if err := ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("USER_BALANCE_NOT_FOUND", "用户余额记录不存在")
		}

		return nil, errors.New(500, "USER_BALANCE_NOT_FOUND", err.Error())
	}

	return &biz.UserBalance{
		ID:      userBalance.ID,
		Balance: userBalance.Balance,
		UserId:  userBalance.UserId,
	}, nil
}

func (ub *UserBalanceRepo) GetUserBalanceRecord(ctx context.Context, reason string, b *biz.Pagination, userIds ...int64) ([]*biz.UserBalanceRecord, error, int64) {
	var (
		userBalanceRecord []*UserBalanceRecord
		count             int64
	)

	instance := ub.data.DB(ctx).Table("user_balance_record")
	if "" != reason {
		instance = instance.Where("reason=?", reason)
	}

	if 0 < len(userIds) {
		instance = instance.Where("user_id IN(?)", userIds)
	}

	instance = instance.Count(&count)
	if err := instance.Scopes(Paginate(b.PageNum, b.PageSize)).Find(&userBalanceRecord).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("USER_BALANCE_NOT_FOUND", "用户余额记录不存在"), 0
		}

		return nil, errors.New(500, "USER_BALANCE_NOT_FOUND", err.Error()), 0
	}

	res := make([]*biz.UserBalanceRecord, 0)
	for _, item := range userBalanceRecord {
		res = append(res, &biz.UserBalanceRecord{
			UserId:    item.UserId,
			Balance:   item.Balance,
			Type:      item.Type,
			Amount:    item.Amount,
			Reason:    item.Reason,
			CreatedAt: item.CreatedAt,
		})
	}

	return res, nil, count
}

func (ub *UserBalanceRepo) GetUserBalanceRecordTotal(ctx context.Context, recordType string, today bool) (*biz.UserBalanceRecordTotal, error) {
	var userBalanceRecordTotal UserBalanceRecordTotal
	instance := ub.data.DB(ctx).Table("user_balance_record")
	if "" != recordType {
		instance = instance.Where("type=?", recordType)
	}

	if today {
		t := time.Now().UTC().Add(8 * time.Hour)
		createdAt := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
		fmt.Println(t.Day(), createdAt)
		instance = instance.Where("created_at>=?", createdAt)
	}

	if err := instance.Select("sum(amount) as total").Take(&userBalanceRecordTotal).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &biz.UserBalanceRecordTotal{
				Total: 0,
			}, nil
		}

		return nil, errors.New(500, "USER_BALANCE_RECORD_NOT_FOUND", err.Error())
	}

	return &biz.UserBalanceRecordTotal{
		Total: userBalanceRecordTotal.Total,
	}, nil
}

func (ub *UserBalanceRepo) GetUserBalanceTotal(ctx context.Context) (*biz.UserBalanceTotal, error) {
	var userBalanceTotal UserBalanceTotal
	instance := ub.data.DB(ctx).Table("user_balance")
	if err := instance.Select("sum(balance) as total").Take(&userBalanceTotal).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &biz.UserBalanceTotal{
				Total: 0,
			}, nil
		}

		return nil, errors.New(500, "USER_BALANCE_NOT_FOUND", err.Error())
	}

	return &biz.UserBalanceTotal{
		Total: userBalanceTotal.Total,
	}, nil
}

func (ub *UserBalanceRepo) GetAddressEthBalanceByAddress(ctx context.Context, address string) (*biz.AddressEthBalance, error) {
	var addressEthBalance AddressEthBalance
	if err := ub.data.DB(ctx).Where("address=?", address).Table("address_eth_balance").First(&addressEthBalance).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("ADDRESS_ETH_BALANCE_NOT_FOUND", "地址余额记录不存在")
		}

		return nil, errors.New(500, "ADDRESS_ETH_BALANCE_NOT_FOUND", err.Error())
	}

	return &biz.AddressEthBalance{
		ID:      addressEthBalance.ID,
		Balance: addressEthBalance.Balance,
		Address: addressEthBalance.Address,
	}, nil
}

func (ub *UserBalanceRepo) WithdrawById(ctx context.Context, id int64) (*biz.UserWithdraw, error) {
	var userWithdraw UserWithdraw
	if err := ub.data.DB(ctx).Table("user_withdraw").Where("id=?", id).First(&userWithdraw).Error; err != nil {
		return nil, errors.NotFound("USER_WITHDRAW_NOT_FOUND", "未查到记录不存在")
	}

	return &biz.UserWithdraw{
		ID:        userWithdraw.ID,
		UserId:    userWithdraw.UserId,
		Amount:    userWithdraw.Amount,
		Status:    userWithdraw.Status,
		Tx:        userWithdraw.Tx,
		CreatedAt: userWithdraw.CreatedAt,
	}, nil
}

func (ub *UserBalanceRepo) WithdrawList(ctx context.Context, status string, b *biz.Pagination, userIds ...int64) ([]*biz.UserWithdraw, error, int64) {
	var (
		count        int64
		userWithdraw []*UserWithdraw
	)
	instance := ub.data.DB(ctx).Table("user_withdraw")

	if "" != status {
		instance = instance.Where("status=?", status)
	}

	if 0 < len(userIds) {
		instance = instance.Where("user_id IN(?)", userIds)
	}

	instance = instance.Count(&count)
	if err := instance.Scopes(Paginate(b.PageNum, b.PageSize)).
		Order("created_at desc").
		Find(&userWithdraw).Error; err != nil {
		return nil, errors.NotFound("USER_WITHDRAW_NOT_FOUND", "未查到记录不存在"), 0
	}

	res := make([]*biz.UserWithdraw, 0)
	for _, item := range userWithdraw {
		res = append(res, &biz.UserWithdraw{
			ID:        item.ID,
			UserId:    item.UserId,
			Status:    item.Status,
			Amount:    item.Amount,
			Tx:        item.Tx,
			CreatedAt: item.CreatedAt,
		})
	}

	return res, nil, count
}

func (ub *UserBalanceRepo) UpdateEthBalanceByAddress(ctx context.Context, address string, balance string) (bool, error) {
	if err := ub.data.DB(ctx).Where("address=?", address).
		Table("address_eth_balance").
		Update("balance", balance).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, errors.NotFound("ADDRESS_ETH_BALANCE_NOT_FOUND", "地址余额不存在")
		}

		return false, errors.New(500, "ADDRESS_ETH_BALANCE_ERROR", err.Error())
	}

	return true, nil
}

// TransferIntoUserBack 在事务中使用，退款
func (ub *UserBalanceRepo) TransferIntoUserBack(ctx context.Context, userId int64, amount int64) (int64, error) {
	var err error
	if err = ub.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		// UpdateColumn("balance", gorm.Expr("balance + ?", pay))
		Updates(map[string]interface{}{"balance": gorm.Expr("balance + ?", amount)}).Error; nil != err {
		return 0, errors.NotFound("user balance err", "user balance not found")
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.Balance
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "transfer_into"
	userBalanceRecode.Reason = "user_goal_back"
	userBalanceRecode.Amount = amount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// TransferIntoUserGoalReward 在事务中使用，中奖
func (ub *UserBalanceRepo) TransferIntoUserGoalReward(ctx context.Context, userId int64, amount int64) (int64, error) {
	var err error
	if err = ub.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		// UpdateColumn("balance", gorm.Expr("balance + ?", pay))
		Updates(map[string]interface{}{"balance": gorm.Expr("balance + ?", amount)}).Error; nil != err {
		return 0, errors.NotFound("user balance err", "user balance not found")
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.Balance
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "transfer_into"
	userBalanceRecode.Reason = "user_goal_reward"
	userBalanceRecode.Amount = amount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

func (ub *UserBalanceRepo) CreateBalanceRecordIdRel(ctx context.Context, recordId int64, relType string, id int64) error {
	var balanceRecordIdRel BalanceRecordIdRel
	balanceRecordIdRel.RecordId = recordId
	balanceRecordIdRel.RelType = relType
	balanceRecordIdRel.RelId = id
	err := ub.data.DB(ctx).Table("balance_record_id_rel").Create(&balanceRecordIdRel).Error
	if err != nil {
		return err
	}

	return nil
}

// TransferIntoUserGoalRecommendReward 在事务中使用
func (ub *UserBalanceRepo) TransferIntoUserGoalRecommendReward(ctx context.Context, userId int64, amount int64) (int64, error) {
	var err error
	if err = ub.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		// UpdateColumn("balance", gorm.Expr("balance + ?", pay))
		Updates(map[string]interface{}{"balance": gorm.Expr("balance + ?", amount)}).Error; nil != err {
		return 0, errors.NotFound("user balance err", "user balance not found")
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.Balance
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "transfer_into"
	userBalanceRecode.Reason = "recommend_user_goal_reward"
	userBalanceRecode.Amount = amount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

func (ub *UserBalanceRepo) UpdateUserBalance(ctx context.Context, userId int64, amount int64) (bool, error) {
	var err error
	if err = ub.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Update("balance", amount).Error; nil != err {
		return false, errors.NotFound("user balance err", "user balance not found")
	}

	return true, nil
}

// Deposit 在事务中使用
func (ub *UserBalanceRepo) Deposit(ctx context.Context, userId int64, amount int64) (*biz.UserBalance, error) {
	var err error
	if err = ub.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance": gorm.Expr("balance + ?", amount)}).Error; nil != err {
		return nil, errors.NotFound("user balance err", "user balance not found")
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return nil, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.Balance
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "deposit"
	userBalanceRecode.Reason = "user_deposit"
	userBalanceRecode.Amount = amount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return nil, err
	}

	return &biz.UserBalance{
		UserId:  userBalance.UserId,
		Balance: userBalance.Balance,
	}, nil
}

// Withdraw 在事务中使用
func (ub *UserBalanceRepo) Withdraw(ctx context.Context, userId int64, amount int64) error {
	var err error
	if res := ub.data.DB(ctx).Table("user_balance").
		Where("user_id=? and balance>=?", userId, amount).
		Updates(map[string]interface{}{"balance": gorm.Expr("balance - ?", amount)}); 0 == res.RowsAffected || nil != res.Error {
		return errors.NotFound("user balance err", "user balance error")
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.Balance
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "withdraw"
	userBalanceRecode.Reason = "user_withdraw"
	userBalanceRecode.Amount = amount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return err
	}

	return nil
}

// UpdateWithdraw 在事务中使用
func (ub *UserBalanceRepo) UpdateWithdraw(ctx context.Context, Id int64, status string, tx string) error {
	if res := ub.data.DB(ctx).Table("user_withdraw").
		Where("id=?", Id).
		Updates(&UserWithdraw{Status: status, Tx: tx}); nil != res.Error {
		return errors.NotFound("user balance err", "user withdraw error")
	}

	return nil
}

// GetUserInfoByMyRecommendCode .
func (ui *UserInfoRepo) GetUserInfoByMyRecommendCode(ctx context.Context, myRecommendCode string) (*biz.UserInfo, error) {
	var userInfo UserInfo
	if err := ui.data.db.Where(&UserInfo{MyRecommendCode: myRecommendCode}).Table("user_info").First(&userInfo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("USER_NOT_FOUND", "user not found")
		}

		return nil, errors.New(500, "USER_NOT_FOUND", err.Error())
	}

	return &biz.UserInfo{
		ID:              userInfo.ID,
		Name:            userInfo.Name,
		Avatar:          userInfo.Avatar,
		UserId:          userInfo.UserId,
		MyRecommendCode: userInfo.MyRecommendCode,
	}, nil
}

// GetUserInfoByUserId .
func (ui *UserInfoRepo) GetUserInfoByUserId(ctx context.Context, userId int64) (*biz.UserInfo, error) {
	var userInfo UserInfo
	if err := ui.data.db.Where(&UserInfo{UserId: userId}).Table("user_info").First(&userInfo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("USER_NOT_FOUND", "user not found")
		}

		return nil, errors.New(500, "USER_NOT_FOUND", err.Error())
	}

	return &biz.UserInfo{
		ID:              userInfo.ID,
		Name:            userInfo.Name,
		Avatar:          userInfo.Avatar,
		UserId:          userInfo.UserId,
		MyRecommendCode: userInfo.MyRecommendCode,
		RecommendCode:   userInfo.RecommendCode,
	}, nil
}

func (ui *UserInfoRepo) GetUserInfoListByRecommendCode(ctx context.Context, recommendCode string) ([]*biz.UserInfo, error) {
	var userInfo []*UserInfo
	if err := ui.data.DB(ctx).
		Table("user_info").
		Where("recommend_code=?", recommendCode).
		Find(&userInfo).Error; err != nil {
		return nil, errors.NotFound("USER_INFO_NOT_FOUND", "用户信息不存在")
	}

	res := make([]*biz.UserInfo, 0)
	for _, item := range userInfo {
		res = append(res, &biz.UserInfo{
			ID:              item.ID,
			Name:            item.Name,
			Avatar:          item.Avatar,
			UserId:          item.UserId,
			MyRecommendCode: item.MyRecommendCode,
			CreatedAt:       item.CreatedAt,
		})
	}

	return res, nil
}

func (u *UserRepo) GetUserList(ctx context.Context, address string, b *biz.Pagination) ([]*biz.User, error, int64) {
	var (
		user  []*User
		count int64
	)

	instance := u.data.DB(ctx).Table("user")
	if "" != address {
		instance = instance.Where("address=?", address)
	}

	instance = instance.Count(&count)
	if err := instance.Scopes(Paginate(b.PageNum, b.PageSize)).Find(&user).Error; err != nil {
		return nil, errors.NotFound("USER_NOT_FOUND", "用户不存在"), 0
	}

	res := make([]*biz.User, 0)
	for _, item := range user {
		res = append(res, &biz.User{
			ID:                  item.ID,
			Address:             item.Address,
			ToAddress:           item.ToAddress,
			ToAddressPrivateKey: item.ToAddressPrivateKey,
		})
	}

	return res, nil, count
}

// CreateUserProxy .
func (u *UserRepo) CreateUserProxy(ctx context.Context, userId int64, rate int64) (*biz.UserProxy, error) {
	var userProxy UserProxy
	userProxy.UserId = userId
	userProxy.Rate = rate
	res := u.data.DB(ctx).Table("user_proxy").Create(&userProxy)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_USER_PROXY_ERROR", "用户代理创建失败")
	}

	return &biz.UserProxy{
		ID:     userProxy.ID,
		Rate:   userProxy.Rate,
		UserId: userProxy.UserId,
	}, nil
}

// UpdateUserProxy .
func (u *UserRepo) UpdateUserProxy(ctx context.Context, userId int64, rate int64) (*biz.UserProxy, error) {
	var userProxy UserProxy
	userProxy.Rate = rate
	res := u.data.DB(ctx).Table("user_proxy").Where("user_id=? and up_user_id=?", userId, 0).Updates(&userProxy)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_USER_PROXY_ERROR", "用户代理修改失败")
	}

	return &biz.UserProxy{
		Rate:   userProxy.Rate,
		UserId: userProxy.UserId,
	}, nil
}

// GetUserProxyByUserId .
func (u *UserRepo) GetUserProxyByUserId(ctx context.Context, userId int64) (*biz.UserProxy, error) {
	var userProxy UserProxy
	if err := u.data.db.Where("user_id=?", userId).
		Table("user_proxy").First(&userProxy).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("USER_PROXY_NOT_FOUND", "user proxy not found")
		}

		return nil, errors.New(500, "USER_PROXY_NOT_FOUND", err.Error())
	}

	return &biz.UserProxy{
		UserId: userProxy.UserId,
		Rate:   userProxy.Rate,
	}, nil
}

func (u *UserRepo) GetUserById(ctx context.Context, userId int64) (*biz.User, error) {
	var user *User
	if err := u.data.DB(ctx).Where("ID=?", userId).Table("user").First(&user).Error; err != nil {
		return nil, errors.NotFound("USER_NOT_FOUND", "用户不存在")
	}

	return &biz.User{
		ID:                  user.ID,
		Address:             user.Address,
		ToAddress:           user.ToAddress,
		ToAddressPrivateKey: user.ToAddressPrivateKey,
	}, nil
}

func (u *UserRepo) GetUserListByUserIds(ctx context.Context, userIds ...int64) ([]*biz.User, error) {
	var user []*User
	if err := u.data.DB(ctx).Table("user").
		Where("ID IN (?)", userIds).
		Find(&user).Error; err != nil {
		return nil, errors.NotFound("USER_NOT_FOUND", "用户不存在")
	}

	res := make([]*biz.User, 0)
	for _, item := range user {
		res = append(res, &biz.User{
			ID:                  item.ID,
			Address:             item.Address,
			ToAddress:           item.ToAddress,
			ToAddressPrivateKey: item.ToAddressPrivateKey,
		})
	}

	return res, nil
}

func (u *UserRepo) GetUserProxyList(ctx context.Context, userId ...int64) ([]*biz.UserProxy, error) {
	var userProxy []*UserProxy
	if err := u.data.DB(ctx).
		Where("up_user_id", userId).
		Table("user_proxy").Find(&userProxy).Error; err != nil {
		return nil, errors.NotFound("USER_NOT_FOUND", "用户代理不存在")
	}

	res := make([]*biz.UserProxy, 0)
	for _, item := range userProxy {
		res = append(res, &biz.UserProxy{
			ID:        item.ID,
			UpUserId:  item.UpUserId,
			UserId:    item.UserId,
			CreatedAt: item.CreatedAt,
			Rate:      item.Rate,
		})
	}

	return res, nil
}

func (u *UserRepo) GetUserMap(ctx context.Context, userIds ...int64) (map[int64]*biz.User, error) {
	var user []*User
	if err := u.data.DB(ctx).Table("user").
		Where("ID IN (?)", userIds).
		Find(&user).Error; err != nil {
		return nil, errors.NotFound("USER_NOT_FOUND", "用户不存在")
	}

	res := make(map[int64]*biz.User, 0)
	for _, item := range user {
		res[item.ID] = &biz.User{
			ID:                  item.ID,
			Address:             item.Address,
			ToAddress:           item.ToAddress,
			ToAddressPrivateKey: item.ToAddressPrivateKey,
		}
	}

	return res, nil
}
