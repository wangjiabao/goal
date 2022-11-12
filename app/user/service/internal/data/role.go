package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"goal/user/internal/biz"
	"gorm.io/gorm"
	"time"
)

type Role struct {
	ID        int64     `gorm:"primarykey;type:int"`
	Type      string    `gorm:"type:varchar(45);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type UserRole struct {
	ID        int64     `gorm:"primarykey;type:int"`
	UserId    int64     `gorm:"type:int;not null"`
	RoleId    int64     `gorm:"type:int;not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type RoleRepo struct {
	data *Data
	log  *log.Helper
}

type UserRoleRepo struct {
	data *Data
	log  *log.Helper
}

func NewRoleRepo(data *Data, logger log.Logger) biz.RoleRepo {
	return &RoleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func NewUserRoleRepo(data *Data, logger log.Logger) biz.UserRoleRepo {
	return &UserRoleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// GetRoleById .
func (r *RoleRepo) GetRoleById(ctx context.Context, Id int64) (*biz.Role, error) {
	var role Role
	if err := r.data.db.Where(&Role{ID: Id}).Table("role").First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("ROLE_NOT_FOUND", "role not found")
		}

		return nil, errors.New(500, "ROLE_NOT_FOUND", err.Error())
	}

	return &biz.Role{
		ID:   role.ID,
		Type: role.Type,
	}, nil
}

// GetRoleByType .
func (r *RoleRepo) GetRoleByType(ctx context.Context, Type string) (*biz.Role, error) {
	var role Role
	if err := r.data.db.Where(&Role{Type: Type}).Table("role").First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("ROLE_NOT_FOUND", "role not found")
		}

		return nil, errors.New(500, "ROLE_NOT_FOUND", err.Error())
	}

	return &biz.Role{
		ID:   role.ID,
		Type: role.Type,
	}, nil
}

// CreateUserRole creat user and role relation
func (ur *UserRoleRepo) CreateUserRole(ctx context.Context, u *biz.User, role *biz.Role) (*biz.UserRole, error) {
	var userRole UserRole
	userRole.RoleId = role.ID
	userRole.UserId = u.ID

	res := ur.data.DB(ctx).Table("user_role").Create(&userRole)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_USER_ROLE_ERROR", "用户角色关系创建失败")
	}

	return &biz.UserRole{
		ID:     userRole.ID,
		UserId: userRole.UserId,
		RoleId: userRole.RoleId,
	}, nil
}
