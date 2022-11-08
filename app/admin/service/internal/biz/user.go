package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"strconv"
)

type User struct {
	ID                  int64
	Address             string
	ToAddress           string
	ToAddressPrivateKey string
}

type AddressEthBalance struct {
	ID      int64
	Address string
	Balance string
}

type UserRepo interface {
	GetUserList(ctx context.Context) ([]*User, error)
}

type UserUseCase struct {
	repo   UserRepo
	uiRepo UserInfoRepo
	ubRepo UserBalanceRepo
	tx     Transaction
	log    *log.Helper
}

func NewUserUseCase(repo UserRepo, tx Transaction, uiRepo UserInfoRepo, ubRepo UserBalanceRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		repo:   repo,
		tx:     tx,
		uiRepo: uiRepo,
		ubRepo: ubRepo,
		log:    log.NewHelper(logger),
	}
}

func (u *UserUseCase) GetUserList(ctx context.Context) ([]*User, error) {
	return u.repo.GetUserList(ctx)
}

//func (u *UserUseCase) GetAddressEthBalance(ctx context.Context, address string) (*AddressEthBalance, error) {
//	return u.ubRepo.GetAddressEthBalanceByAddress(ctx, address)
//}

func (u *UserUseCase) Deposit(ctx context.Context, balance string, address string, userId int64) (bool, error) {
	var (
		currentBalance, lastBalance int64
		base                        int64 = 100000 // 基础精度0.00001 todo 加配置文件
	)

	addressEthBalance, err := u.ubRepo.GetAddressEthBalanceByAddress(ctx, address)
	if err != nil {
		return false, err
	}
	lenBalance := len(balance)
	if lenBalance > 18 {
		if currentBalance, err = strconv.ParseInt(balance[:lenBalance-18], 10, 64); err != nil {
			return false, err
		}
	} else {
		currentBalance = 0
	}
	lenLastBalance := len(addressEthBalance.Balance)
	if lenLastBalance > 18 {
		if lastBalance, err = strconv.ParseInt(addressEthBalance.Balance[:lenLastBalance-18], 10, 64); err != nil {
			return false, err
		}
	} else {
		lastBalance = 0
	}

	fmt.Println(currentBalance, lastBalance)
	if currentBalance <= lastBalance {
		return false, err
	}

	depositBalanceNow := (currentBalance - lastBalance) * base

	fmt.Println(depositBalanceNow)
	//if err = u.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
	_, err = u.ubRepo.Deposit(ctx, userId, depositBalanceNow) // todo 事务
	//if nil != err {
	//	return err
	//}
	_, err = u.ubRepo.UpdateEthBalanceByAddress(ctx, addressEthBalance.Address, strconv.FormatInt(currentBalance, 10))
	//if err != nil {
	//	return err
	//}
	//return nil
	//}); nil != err {
	//	fmt.Println(4444)
	//	return false, nil
	//}

	return true, nil
}
