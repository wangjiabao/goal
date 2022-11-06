package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Role struct {
	ID   int64
	Type string
}

type UserRole struct {
	ID     int64
	UserId int64
	RoleId int64
}

type RoleRepo interface {
	GetRoleById(ctx context.Context, Id int64) (*Role, error)
	GetRoleByType(ctx context.Context, Type string) (*Role, error)
}

type UserRoleRepo interface {
	CreateUserRole(ctx context.Context, u *User, role *Role) (*UserRole, error)
}

type RoleUseCase struct {
	repo   RoleRepo
	urRepo UserRoleRepo
	log    *log.Helper
}

func NewRoleUseCase(repo RoleRepo, logger log.Logger) *RoleUseCase {
	return &RoleUseCase{repo: repo, log: log.NewHelper(logger)}
}
