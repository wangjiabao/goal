package biz

import (
	"context"
)

type UserBalance struct {
	ID      int64
	UserId  int64
	Balance int64
}

type UserBalanceRepo interface {
	CreateUserBalance(ctx context.Context, u *User) (*UserBalance, error)
	GetUserBalanceByUserId(ctx context.Context, userId int64) (*UserBalance, error)
	Deposit(ctx context.Context, userId int64, amount int64) (*UserBalance, error)
}
