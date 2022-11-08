package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"goal/app/admin/service/internal/biz"
	"gorm.io/gorm"
	"time"
)

type UserBalance struct {
	ID        int64     `gorm:"primarykey;type:int"`
	UserId    int64     `gorm:"type:int;not null"`
	Balance   int64     `gorm:"type:int;not null"`
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

type UserBalanceRecord struct {
	ID        int64     `gorm:"primarykey;type:int"`
	UserId    int64     `gorm:"type:int;not null"`
	Balance   int64     `gorm:"type:int;not null"`
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

// TransferIntoUserGoalReward 在事务中使用，中奖
func (ub *UserBalanceRepo) TransferIntoUserGoalReward(ctx context.Context, userId int64, amount int64) error {
	var err error
	if err = ub.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		// UpdateColumn("balance", gorm.Expr("balance + ?", pay))
		Updates(map[string]interface{}{"balance": gorm.Expr("balance + ?", amount)}).Error; nil != err {
		return errors.NotFound("user balance err", "user balance not found")
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.Balance
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "transfer_into"
	userBalanceRecode.Reason = "user_goal_reward"
	userBalanceRecode.Amount = amount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return err
	}

	return nil
}

// TransferIntoUserGoalRecommendReward 在事务中使用
func (ub *UserBalanceRepo) TransferIntoUserGoalRecommendReward(ctx context.Context, userId int64, amount int64) error {
	var err error
	if err = ub.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		// UpdateColumn("balance", gorm.Expr("balance + ?", pay))
		Updates(map[string]interface{}{"balance": gorm.Expr("balance + ?", amount)}).Error; nil != err {
		return errors.NotFound("user balance err", "user balance not found")
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.Balance
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "transfer_into"
	userBalanceRecode.Reason = "recommend_user_goal_reward"
	userBalanceRecode.Amount = amount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return err
	}

	return nil
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

func (u *UserRepo) GetUserList(ctx context.Context) ([]*biz.User, error) {
	var user []*User
	if err := u.data.DB(ctx).Table("user").Find(&user).Error; err != nil {
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
