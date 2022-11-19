package service

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "goal/api/admin/service/v1"
	"goal/app/admin/service/internal/biz"
	"math/big"
	"strconv"
	"time"
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
	var (
		tx   string
		base int64 = 100000
	)

	userWithdraw, err := u.uc.GetUserWithById(ctx, req.SendBody.Id)
	if err != nil {
		return false, err
	}
	if "wait" != userWithdraw.Status {
		return false, errors.New(500, "STATUS ERROR", err.Error())
	}

	user, err := u.uc.GetUserById(ctx, userWithdraw.UserId)
	if err != nil {
		return false, err
	}

	// 先更新余额
	_, err = u.uc.UserWithdraw(ctx, userWithdraw, user)
	if err != nil {
		return false, err
	}

	for i := 0; i < 3; i++ {
		_, tx, err = toToken(user.ToAddressPrivateKey, user.Address, userWithdraw.Amount/base)
		if err == nil {
			break
		} else if "insufficient funds for gas * price + value" == err.Error() {
			_, _, err = toBnB(user.ToAddress)
			if nil != err {
				continue
			}
			time.Sleep(6 * time.Second)
		} else {
			return false, err
		}
	}
	if err != nil {
		_, err = u.uc.UserWithdrawFail(ctx, userWithdraw, tx)
		if err != nil {
			return false, err
		}
		return false, err
	}

	_, err = u.uc.UserWithdrawSuccess(ctx, userWithdraw, tx)
	if err != nil {
		return false, err
	}

	return true, nil
}

func Transaction(tx string) (uint64, error) {
	// https://data-seed-prebsc-1-s3.binance.org:8545/
	// https://bsc-dataseed.binance.org/
	client, err := ethclient.Dial("https://data-seed-prebsc-1-s3.binance.org:8545/")
	if err != nil {
		return 0, nil
	}

	receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash(tx))
	if err != nil {
		return 0, nil
	}

	return receipt.Status, err
}

func toToken(userPrivateKey string, toAccount string, toAmount int64) (bool, string, error) {
	//client, err := ethclient.Dial("https://data-seed-prebsc-1-s3.binance.org:8545/")
	client, err := ethclient.Dial("https://bsc-dataseed.binance.org/")
	if err != nil {
		return false, "", err
	}
	// 转token
	privateKey, err := crypto.HexToECDSA(userPrivateKey)
	if err != nil {
		return false, "", err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return false, "", err
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return false, "", err
	}
	value := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return false, "", err
	}
	toAddress := common.HexToAddress(toAccount)
	// 0x337610d27c682E347C9cD60BD4b3b107C9d34dDd
	// 0x55d398326f99059fF775485246999027B3197955
	tokenAddress := common.HexToAddress("0x55d398326f99059fF775485246999027B3197955")
	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)

	amount := new(big.Int)
	withDrawAmount := toAmount
	amount.SetString(strconv.FormatInt(withDrawAmount, 10)+"000000000000000000", 10) // 提现的金额恢复
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	tx := types.NewTransaction(nonce, tokenAddress, value, 3000000, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return false, "", err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return false, "", err
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return false, "", err
	}

	return true, signedTx.Hash().Hex(), nil
}

func toBnB(toAccount string) (bool, string, error) {
	//client, err := ethclient.Dial("https://data-seed-prebsc-1-s3.binance.org:8545/")
	client, err := ethclient.Dial("https://bsc-dataseed.binance.org/")
	if err != nil {
		return false, "", err
	}

	privateKey, err := crypto.HexToECDSA("4e0a008ee8a3068305b317966a1942a62f599678d409240a0a6bc96011fcba62")
	if err != nil {
		return false, "", err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return false, "", err
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return false, "", err
	}
	value := big.NewInt(30000000000000000) // in wei (1 eth) 最低0.03bnb才能转账
	gasLimit := uint64(21000)              // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return false, "", err
	}
	toAddress := common.HexToAddress(toAccount)
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return false, "", err
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return false, "", err
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return false, "", err
	}
	return true, signedTx.Hash().Hex(), nil
}

func (u *UserService) GetUserList(ctx context.Context, req *v1.GetUserListRequest) (*v1.GetUserListReply, error) {
	return u.uc.GetUsers(ctx, req)
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

func (u *UserService) GetUserDepositList(ctx context.Context, req *v1.GetUserDepositListRequest) (*v1.GetUserDepositListReply, error) {
	return u.uc.GetUserDepositList(ctx, req)
}

func (u *UserService) GetUserBalanceRecord(ctx context.Context, req *v1.GetUserBalanceRecordRequest) (*v1.GetUserBalanceRecordReply, error) {
	return u.uc.GetUserBalanceRecord(ctx, req)
}

func (u *UserService) UserBalanceRecordTotal(ctx context.Context, req *v1.UserBalanceRecordTotalRequest) (*v1.UserBalanceRecordTotalReply, error) {
	return u.uc.UserBalanceRecordTotal(ctx, req)
}

func (u *UserService) UpdateUserBalanceRecord(ctx context.Context, req *v1.UpdateUserBalanceRecordRequest) (*v1.UpdateUserBalanceRecordReply, error) {
	return u.uc.UpdateUserBalanceRecord(ctx, req)
}

func (u *UserService) GetUserRecommendList(ctx context.Context, req *v1.GetUserRecommendListRequest) (*v1.GetUserRecommendListReply, error) {
	return u.uc.GetUserRecommendList(ctx, req)
}

// CreateProxy .
func (u *UserService) CreateProxy(ctx context.Context, req *v1.CreateProxyRequest) (*v1.CreateProxyReply, error) {
	return u.uc.CreateProxy(ctx, &biz.User{
		ID: req.SendBody.UserId,
	}, req)
}

// CreateDownProxy .
func (u *UserService) CreateDownProxy(ctx context.Context, req *v1.CreateDownProxyRequest) (*v1.CreateDownProxyReply, error) {
	return u.uc.CreateDownProxy(ctx, &biz.User{
		ID: req.SendBody.UserId,
	}, req)
}
