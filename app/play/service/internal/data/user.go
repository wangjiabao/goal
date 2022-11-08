package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"goal/app/play/service/internal/biz"
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

type UserBalanceRecord struct {
	ID        int64     `gorm:"primarykey;type:int"`
	UserId    int64     `gorm:"type:int;not null"`
	Balance   int64     `gorm:"type:int;not null"`
	Type      string    `gorm:"type:varchar(45);not null"`
	Amount    int64     `gorm:"type:int;not null"`
	Reason    string    `gorm:"type:varchar(45);not null"`
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

type UserBalanceRepo struct {
	data *Data
	log  *log.Helper
}

type UserProxyRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserBalanceRepo(data *Data, logger log.Logger) biz.UserBalanceRepo {
	return &UserBalanceRepo{
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

// Pay 在事务中使用
func (ub *UserBalanceRepo) Pay(ctx context.Context, userId int64, pay int64) error {
	var err error
	if err = ub.data.DB(ctx).Table("user_balance").
		Where("user_id=? and balance>=?", userId, pay).
		Updates(map[string]interface{}{"balance": gorm.Expr("balance - ?", pay)}).Error; nil != err {
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
	userBalanceRecode.Type = "pay"
	userBalanceRecode.Reason = "user_play_pay"
	userBalanceRecode.Amount = pay
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return err
	}

	return nil
}

// TransferIntoUserPlayProxyReward 在事务中使用
func (ub *UserBalanceRepo) TransferIntoUserPlayProxyReward(ctx context.Context, userId int64, amount int64) error {
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
	userBalanceRecode.Amount = amount
	userBalanceRecode.Reason = "proxy_use_play_reward"
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return err
	}

	return nil
}
