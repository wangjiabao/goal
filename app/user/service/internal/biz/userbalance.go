package biz

import (
	"context"
	"time"
)

type UserBalance struct {
	ID      int64
	UserId  int64
	Balance int64
}

type AddressEthBalance struct {
	ID      int64
	Balance string
	Address string
	Version int64
	Status  int64
}

type UserBalanceRecord struct {
	ID        int64
	UserId    int64
	Amount    int64
	Type      string
	Reason    string
	CreatedAt time.Time
}

type UserWithdraw struct {
	ID        int64
	UserId    int64
	Amount    int64
	Status    string
	CreatedAt time.Time
}

type UserBalanceRepo interface {
	UnLockAndUpdateEthBalanceByAddress(ctx context.Context, address string, balance string) (bool, error)
	LockEthBalanceByAddress(ctx context.Context, address string) (bool, error)
	UnLockEthBalanceByAddress(ctx context.Context, address string) (bool, error)
	GetAddressEthBalanceByAddress(ctx context.Context, address string) (*AddressEthBalance, error)
	CreateUserBalance(ctx context.Context, u *User) (*UserBalance, error)
	GetUserBalanceByUserId(ctx context.Context, userId int64) (*UserBalance, error)
	Deposit(ctx context.Context, userId int64, amount int64) (*UserBalance, error)
	Withdraw(ctx context.Context, userId int64, amount int64) (bool, error)
	WithdrawList(ctx context.Context, userId int64) ([]*UserWithdraw, error)
	TransferIntoProxy(ctx context.Context, userId int64, amount int64) (*UserBalance, error)
	TransferIntoProxyRecommendReward(ctx context.Context, userId int64, amount int64) (*UserBalance, error)
	GetUserBalanceRecordByUserId(ctx context.Context, userId int64, recordType string, reason string) ([]*UserBalanceRecord, error)
}
