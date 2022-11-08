package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"goal/app/user/service/internal/biz"
	"gorm.io/gorm"
	"strconv"
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

type UserProxy struct {
	ID        int64     `gorm:"primarykey;type:int"`
	UserId    int64     `gorm:"type:int;not null"`
	UpUserId  int64     `gorm:"type:int;not null"`
	Rate      int64     `gorm:"type:int;not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type UserRepo struct {
	data *Data
	log  *log.Helper
}

type UserInfoRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &UserRepo{
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

// CreateUser .
func (u *UserRepo) CreateUser(ctx context.Context, uc *biz.User) (*biz.User, error) {
	var user User
	user.Address = uc.Address
	user.ToAddress = uc.ToAddress
	user.ToAddressPrivateKey = uc.ToAddressPrivateKey
	res := u.data.DB(ctx).Table("user").Create(&user)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_USER_ERROR", "用户创建失败")
	}

	return &biz.User{
		ID:        user.ID,
		Address:   user.Address,
		ToAddress: user.ToAddress,
	}, nil
}

// GetUserByAddress .
func (u *UserRepo) GetUserByAddress(ctx context.Context, address string) (*biz.User, error) {
	var user User
	if err := u.data.db.Where(&User{Address: address}).Table("user").First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("USER_NOT_FOUND", "user not found")
		}

		return nil, errors.New(500, "USER_NOT_FOUND", err.Error())
	}

	return &biz.User{
		ID:      user.ID,
		Address: user.Address,
	}, nil
}

// GetUserById .
func (u *UserRepo) GetUserById(ctx context.Context, Id int64) (*biz.User, error) {
	var user User
	if err := u.data.db.Where(&User{ID: Id}).Table("user").First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("USER_NOT_FOUND", "user not found")
		}

		return nil, errors.New(500, "USER_NOT_FOUND", err.Error())
	}

	return &biz.User{
		ID:      user.ID,
		Address: user.Address,
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

// CreateUserInfo .
func (ui *UserInfoRepo) CreateUserInfo(ctx context.Context, u *biz.User, recommendCode string) (*biz.UserInfo, error) {
	var userInfo UserInfo
	userInfo.UserId = u.ID
	userInfo.RecommendCode = recommendCode
	userInfo.Code = "GA" + strconv.FormatInt(u.ID, 10)
	userInfo.MyRecommendCode = userInfo.RecommendCode + userInfo.Code

	res := ui.data.DB(ctx).Table("user_info").Create(&userInfo)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_USER_INFO_ERROR", "用户信息创建失败")
	}

	return &biz.UserInfo{
		ID:     userInfo.ID,
		Name:   userInfo.Name,
		Avatar: userInfo.Avatar,
		UserId: userInfo.UserId,
	}, nil
}

func (u *UserRepo) GetUserListByIds(ctx context.Context, userIds ...int64) (map[int64]*biz.User, error) {
	var user []*User
	if err := u.data.DB(ctx).
		Table("user").
		Where("id IN (?)", userIds).
		Find(&user).Error; err != nil {
		return nil, errors.NotFound("USER_NOT_FOUND", "用户信息不存在")
	}

	res := make(map[int64]*biz.User, 0)
	for _, item := range user {
		res[item.ID] = &biz.User{
			ID:      item.ID,
			Address: item.Address,
		}
	}

	return res, nil
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
		})
	}

	return res, nil
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

// CreateUserUpProxy .
func (u *UserRepo) CreateUserUpProxy(ctx context.Context, userId int64, upUserId int64, rate int64) (*biz.UserProxy, error) {
	var userProxy UserProxy
	userProxy.UserId = userId
	userProxy.UpUserId = upUserId
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
