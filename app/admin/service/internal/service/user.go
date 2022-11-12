package service

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-kratos/kratos/v2/log"
	v1 "goal/api/admin/service/v1"
	"goal/app/admin/service/internal/biz"
	"math/big"
	"strconv"
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

func (u *UserService) UserWithdraw(ctx context.Context, req *v1.UserWithdrawRequest) (*v1.UserWithdrawReply, error) {
	_, err := u.Withdraw(ctx, req)
	if nil != err {
		return &v1.UserWithdrawReply{Result: "失败"}, nil
	}

	return &v1.UserWithdrawReply{Result: "成功"}, nil
}

func (u *UserService) Withdraw(ctx context.Context, req *v1.UserWithdrawRequest) (bool, error) {
	var base = int64(100000)
	userWithdraw, err := u.uc.GetUserWithById(ctx, req.SendBody.Id)
	if err != nil {
		return false, err
	}
	user, err := u.uc.GetUserById(ctx, userWithdraw.UserId)
	if err != nil {
		return false, err
	}

	client, err := ethclient.Dial("https://data-seed-prebsc-1-s3.binance.org:8545/")
	//client, err := ethclient.Dial("https://bsc-dataseed.binance.org/")
	if err != nil {
		return false, err
	}
	privateKey, err := crypto.HexToECDSA(user.ToAddressPrivateKey)
	if err != nil {
		return false, err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return false, err
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return false, err
	}
	value := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return false, err
	}
	toAddress := common.HexToAddress(user.Address)
	// 0x337610d27c682E347C9cD60BD4b3b107C9d34dDd
	// 0x55d398326f99059fF775485246999027B3197955
	tokenAddress := common.HexToAddress("0x337610d27c682E347C9cD60BD4b3b107C9d34dDd")
	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)

	amount := new(big.Int)
	withDrawAmount := userWithdraw.Amount / base
	amount.SetString(strconv.FormatInt(withDrawAmount, 10)+"000000000000000000", 10) // 提现的金额恢复
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	tx := types.NewTransaction(nonce, tokenAddress, value, 3000000, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return false, err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return false, err
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return false, err
	}

	fmt.Println(signedTx.Hash().Hex())

	return false, nil
}

func (u *UserService) GetUserList(ctx context.Context, req *v1.GetUserListRequest) (*v1.GetUserListReply, error) {
	return u.uc.GetUsers(ctx)
}

func (u *UserService) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.GetUserReply, error) {
	return u.uc.GetUserInfo(ctx, req)
}

func (u *UserService) GetUserProxyList(ctx context.Context, req *v1.GetUserProxyListRequest) (*v1.GetUserProxyListReply, error) {
	return u.uc.GetUserProxyList(ctx, req)
}

func (u *UserService) GetUserWithdrawList(ctx context.Context, req *v1.GetUserWithdrawListRequest) (*v1.GetUserWithdrawListReply, error) {
	return u.uc.GetUserWithdrawList(ctx, req)
}

func (u *UserService) GetUserBalanceRecord(ctx context.Context, req *v1.GetUserBalanceRecordRequest) (*v1.GetUserBalanceRecordReply, error) {
	return u.uc.GetUserBalanceRecord(ctx)
}

func (u *UserService) GetUserRecommendList(ctx context.Context, req *v1.GetUserRecommendListRequest) (*v1.GetUserRecommendListReply, error) {
	return u.uc.GetUserRecommendList(ctx, req)
}
