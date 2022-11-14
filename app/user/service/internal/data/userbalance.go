package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"goal/user/internal/biz"
	"gorm.io/gorm"
	"time"
)

type UserBalance struct {
	ID        int64     `gorm:"primarykey;type:int"`
	UserId    int64     `gorm:"type:int"`
	Balance   int64     `gorm:"type:bigint"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type UserBalanceRecord struct {
	ID        int64     `gorm:"primarykey;type:int"`
	UserId    int64     `gorm:"type:int;not null"`
	Balance   int64     `gorm:"type:bigint;not null"`
	Type      string    `gorm:"type:varchar(45);not null"`
	Amount    int64     `gorm:"type:bigint;not null"`
	Reason    string    `gorm:"type:varchar(45);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type AddressEthBalance struct {
	ID        int64     `gorm:"primarykey;type:int"`
	Balance   string    `gorm:"type:varchar(45);not null"`
	Address   string    `gorm:"type:varchar(45);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type UserWithdraw struct {
	ID        int64     `gorm:"primarykey;type:int"`
	Amount    int64     `gorm:"type:bigint;not null"`
	UserId    int64     `gorm:"type:int;not null"`
	Status    string    `gorm:"type:varchar(45);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type UserBalanceRepo struct {
	data *Data
	log  *log.Helper
}

type AddressEthBalanceRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserBalanceRepo(data *Data, logger log.Logger) biz.UserBalanceRepo {
	return &UserBalanceRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func NewAddressEthBalanceRepo(data *Data, logger log.Logger) biz.AddressEthBalanceRepo {
	return &AddressEthBalanceRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// CreateUserBalance .
func (ub UserBalanceRepo) CreateUserBalance(ctx context.Context, u *biz.User) (*biz.UserBalance, error) {
	var userBalance UserBalance
	userBalance.UserId = u.ID
	res := ub.data.DB(ctx).Table("user_balance").Create(&userBalance)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_USER_BALANCE_ERROR", "用户余额信息创建失败")
	}

	return &biz.UserBalance{
		ID:      userBalance.ID,
		UserId:  userBalance.UserId,
		Balance: userBalance.Balance,
	}, nil
}

// CreateAddressEthBalance .
func (ab AddressEthBalanceRepo) CreateAddressEthBalance(ctx context.Context, address string) (*biz.AddressEthBalance, error) {
	var addressEthBalance AddressEthBalance
	addressEthBalance.Balance = ""
	addressEthBalance.Address = address
	res := ab.data.DB(ctx).Table("address_eth_balance").Create(&addressEthBalance)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_ADDRESS_USER_BALANCE_ERROR", "地址余额信息创建失败")
	}

	return &biz.AddressEthBalance{
		ID: addressEthBalance.ID,
	}, nil
}

// GetUserBalanceByUserId .
func (ub *UserBalanceRepo) GetUserBalanceByUserId(ctx context.Context, userId int64) (*biz.UserBalance, error) {
	var userBalance UserBalance
	if err := ub.data.db.Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("USER_BALANCE_NOT_FOUND", "user balance not found")
		}

		return nil, errors.New(500, "USER_BALANCE_NOT_FOUND", err.Error())
	}

	return &biz.UserBalance{
		ID:      userBalance.ID,
		UserId:  userBalance.UserId,
		Balance: userBalance.Balance,
	}, nil
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

// TransferIntoProxy 在事务中使用
func (ub *UserBalanceRepo) TransferIntoProxy(ctx context.Context, userId int64, amount int64) (*biz.UserBalance, error) {
	var err error
	if res := ub.data.DB(ctx).Table("user_balance").
		Where("user_id=? and balance>=?", userId, amount).
		Updates(map[string]interface{}{"balance": gorm.Expr("balance - ?", amount)}); 0 == res.RowsAffected || nil != res.Error {
		return nil, errors.NotFound("user balance err", "user balance error")
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return nil, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.Balance
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "pay"
	userBalanceRecode.Reason = "user_proxy"
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

// TransferIntoProxyRecommendReward 在事务中使用
func (ub *UserBalanceRepo) TransferIntoProxyRecommendReward(ctx context.Context, userId int64, amount int64) (*biz.UserBalance, error) {
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
	userBalanceRecode.Type = "transfer_into"
	userBalanceRecode.Reason = "recommend_user_proxy_reward"
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

// Withdraw .
func (ub *UserBalanceRepo) Withdraw(ctx context.Context, userId int64, amount int64) (bool, error) {
	var userWithdraw UserWithdraw
	userWithdraw.UserId = userId
	userWithdraw.Amount = amount
	userWithdraw.Status = "wait"
	if err := ub.data.DB(ctx).Table("user_withdraw").Create(&userWithdraw).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (ub *UserBalanceRepo) WithdrawList(ctx context.Context, userId int64) ([]*biz.UserWithdraw, error) {
	var userWithdraw []*UserWithdraw
	if err := ub.data.DB(ctx).Table("user_withdraw").
		Where("user_id=?", userId).
		Order("created_at desc").
		Find(&userWithdraw).Error; err != nil {
		return nil, errors.NotFound("USER_WITHDRAW_NOT_FOUND", "未查到记录不存在")
	}

	res := make([]*biz.UserWithdraw, 0)
	for _, item := range userWithdraw {
		res = append(res, &biz.UserWithdraw{
			ID:        item.ID,
			Status:    item.Status,
			Amount:    item.Amount,
			CreatedAt: item.CreatedAt,
		})
	}

	return res, nil
}

func (ub *UserBalanceRepo) GetUserBalanceRecordByUserId(ctx context.Context, userId int64, recordType string, reason string) ([]*biz.UserBalanceRecord, error) {
	var userBalanceRecord []*UserBalanceRecord
	if err := ub.data.DB(ctx).Table("user_balance_record").
		Where("user_id=?", userId).
		Where("type=?", recordType).
		Where("reason=?", reason).
		Order("created_at desc").Find(&userBalanceRecord).Error; err != nil {
		return nil, errors.NotFound("USER_BALANCE_RECORD_NOT_FOUND", "未查到记录不存在")
	}

	res := make([]*biz.UserBalanceRecord, 0)
	for _, item := range userBalanceRecord {
		res = append(res, &biz.UserBalanceRecord{
			ID:        item.ID,
			UserId:    item.UserId,
			Type:      item.Type,
			Reason:    item.Reason,
			Amount:    item.Amount,
			CreatedAt: item.CreatedAt,
		})
	}

	return res, nil
}
