package service

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-kratos/kratos/v2/log"
	v1 "goal/api/admin/service/v1"
	"goal/app/admin/service/internal/biz"
)

type UserService struct {
	v1.UnimplementedUserServer

	uc  *biz.UserUseCase
	log *log.Helper
}

// NewUserService new a greeter service.
func NewUserService(uc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{uc: uc, log: log.NewHelper(logger)}
}

func (u *UserService) UserDeposit(ctx context.Context, req *v1.UserDepositRequest) (*v1.UserDepositReply, error) {
	_, err := u.Deposit(ctx)
	if nil != err {
		return nil, err
	}

	return &v1.UserDepositReply{}, nil
}

func (u *UserService) Deposit(ctx context.Context) (bool, error) {
	var (
		user []*biz.User
	)
	client, err := ethclient.Dial("https://data-seed-prebsc-1-s3.binance.org:8545/")
	if err != nil {
		return false, err
	}

	user, _ = u.uc.GetUserList(ctx)

	for _, v := range user {
		fmt.Println(v)
		tokenAddress := common.HexToAddress("0x337610d27c682E347C9cD60BD4b3b107C9d34dDd")
		instance, err := NewToken(tokenAddress, client)
		if err != nil {
			fmt.Println(1222)
			continue
		}
		address := common.HexToAddress(v.ToAddress)
		bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
		if err != nil {
			continue
		}

		_, err = u.uc.Deposit(ctx, bal.String(), v.ToAddress, v.ID)
		if err != nil {
			continue
		}
	}

	return false, nil
}
