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

type UserWithdraw struct {
	ID        int64     `gorm:"primarykey;type:int"`
	Amount    int64     `gorm:"type:int;not null"`
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

func (ub *UserBalanceRepo) GetUserBalanceRecord(ctx context.Context) ([]*biz.UserBalanceRecord, error) {
	var userBalanceRecord []*UserBalanceRecord
	if err := ub.data.DB(ctx).Table("user_balance_record").Find(&userBalanceRecord).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("USER_BALANCE_NOT_FOUND", "用户余额记录不存在")
		}

		return nil, errors.New(500, "USER_BALANCE_NOT_FOUND", err.Error())
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

	return res, nil
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
	if err := ub.data.DB(ctx).Where("id=?", id).First(&userWithdraw).Error; err != nil {
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

func (ub *UserBalanceRepo) WithdrawList(ctx context.Context, status string, b *biz.Pagination) ([]*biz.UserWithdraw, error) {
	var userWithdraw []*UserWithdraw
	instance := ub.data.DB(ctx)

	if "" != status {
		instance = instance.Where("status=?", status)
	}

	if err := instance.Scopes(Paginate(b.PageNum, b.PageSize)).Table("user_withdraw").
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
			Tx:        item.Tx,
			CreatedAt: item.CreatedAt,
		})
	}

	return res, nil
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

func (ui *UserInfoRepo) GetUserInfoListByRecommendCode(ctx context.Context, recommendCode string) ([]*biz.UserInfo, error) {
	var userInfo []*UserInfo
	if err := ui.data.DB(ctx).
		Table("user_info").
		Where("recommend_code Like ?", recommendCode+"%").
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
