package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"goal/app/play/service/internal/biz"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID                  int64     `gorm:"primarykey;type:int"`
	Address             string    `gorm:"type:varchar(45)"`
	ToAddress           string    `gorm:"type:varchar(45)"`
	ToAddressPrivateKey string    `gorm:"type:varchar(100)"`
	CreatedAt           time.Time `gorm:"type:datetime;not null"`
	UpdatedAt           time.Time `gorm:"type:datetime;not null"`
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

type UserBalance struct {
	ID        int64     `gorm:"primarykey;type:int"`
	UserId    int64     `gorm:"type:int;not null"`
	Balance   int64     `gorm:"type:bigint;not null"`
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

func (up *UserProxyRepo) GetUserProxyAndDown(ctx context.Context) (map[int64]*biz.UserProxy, map[int64]*biz.UserProxy, error) {
	var l []*UserProxy
	if err := up.data.DB(ctx).Table("user_proxy").Find(&l).Error; err != nil {
		return nil, nil, errors.InternalServer("SELECT_PLAY_ERROR", "查询代理失败")
	}

	ul := make(map[int64]*biz.UserProxy, 0)
	dl := make(map[int64]*biz.UserProxy, 0)
	for _, v := range l {
		if 0 != v.UpUserId {
			dl[v.UserId] = &biz.UserProxy{
				ID:       v.ID,
				UserId:   v.UserId,
				UpUserId: v.UpUserId,
				Rate:     v.Rate,
			}
			continue
		}

		ul[v.UserId] = &biz.UserProxy{
			ID:       v.ID,
			UserId:   v.UserId,
			UpUserId: v.UpUserId,
			Rate:     v.Rate,
		}
	}
	return ul, dl, nil
}

func (ub *UserBalanceRepo) GetUserBalanceRecordGoalReward(ctx context.Context, ids ...int64) (map[int64]*biz.UserBalanceRecord, error) {
	var userBalanceRecord []*UserBalanceRecord
	res := make(map[int64]*biz.UserBalanceRecord, 0)
	if err := ub.data.DB(ctx).Table("user_balance_record").Where("id IN (?) and reason=?", ids, "user_goal_reward").Find(&userBalanceRecord).Error; err != nil {
		return res, errors.NotFound("ROOM_NOT_FOUND", "记录不存在")
	}

	for _, item := range userBalanceRecord {
		fmt.Println(item)
		res[item.ID] = &biz.UserBalanceRecord{
			ID:     item.ID,
			Amount: item.Amount,
		}
	}

	return res, nil
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
func (ub *UserBalanceRepo) Pay(ctx context.Context, userId int64, pay int64) (int64, error) {
	var err error
	if res := ub.data.DB(ctx).Table("user_balance").
		Where("user_id=? and balance>=?", userId, pay).
		Updates(map[string]interface{}{"balance": gorm.Expr("balance - ?", pay)}); 0 == res.RowsAffected || nil != res.Error {
		return 0, errors.NotFound("user balance err", "user balance error")
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.Balance
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "pay"
	userBalanceRecode.Reason = "user_play_pay"
	userBalanceRecode.Amount = pay
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
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

// RoomFee 在事务中使用
func (ub *UserBalanceRepo) RoomFee(ctx context.Context, userId int64, pay int64) (int64, error) {
	var err error
	if res := ub.data.DB(ctx).Table("user_balance").
		Where("user_id=? and balance>=?", userId, pay).
		Updates(map[string]interface{}{"balance": gorm.Expr("balance - ?", pay)}); 0 == res.RowsAffected || nil != res.Error {
		return 0, errors.NotFound("user balance err", "user balance error")
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.Balance
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "pay"
	userBalanceRecode.Reason = "room_fee"
	userBalanceRecode.Amount = pay
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

func (ub *UserBalanceRepo) GetBalanceRecordIdRelMap(ctx context.Context, relType string, id ...int64) (map[int64]*biz.BalanceRecordIdRel, error) {
	var l []*BalanceRecordIdRel
	if err := ub.data.DB(ctx).Table("balance_record_id_rel").Where("rel_id IN (?) and rel_type=?", id, relType).Find(&l).Error; err != nil {
		return nil, errors.InternalServer("SELECT_PLAY_ERROR", "查询代理失败")
	}

	res := make(map[int64]*biz.BalanceRecordIdRel, 0)
	for _, v := range l {
		res[v.RelId] = &biz.BalanceRecordIdRel{
			ID:       v.ID,
			RecordId: v.RecordId,
			RelId:    v.RelId,
		}
	}

	return res, nil
}

// TransferIntoUserPlayProxyReward 在事务中使用
func (ub *UserBalanceRepo) TransferIntoUserPlayProxyReward(ctx context.Context, userId int64, amount int64) (int64, error) {
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
	userBalanceRecode.Amount = amount
	userBalanceRecode.Reason = "proxy_user_play_reward"
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}
