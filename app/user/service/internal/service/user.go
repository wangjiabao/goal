package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwt2 "github.com/golang-jwt/jwt/v4"
	v1 "goal/api/user/service/v1"
	"goal/app/user/service/internal/biz"
	"goal/app/user/service/internal/conf"
	"goal/app/user/service/internal/pkg/middleware/auth"
	"time"
)

type UserService struct {
	v1.UnimplementedUserServer

	uc  *biz.UserUseCase
	log *log.Helper
	ca  *conf.Auth
}

// NewUserService new a greeter service.
func NewUserService(uc *biz.UserUseCase, logger log.Logger, ca *conf.Auth) *UserService {
	return &UserService{uc: uc, log: log.NewHelper(logger), ca: ca}
}

// EthAuthorize exist user or create and ethAuthorize
func (u *UserService) EthAuthorize(ctx context.Context, req *v1.EthAuthorizeRequest) (*v1.EthAuthorizeReply, error) {
	// TODO 以太坊验证用户真实性
	userAddress := req.SendBody.Address // 以太坊账户
	if "" == userAddress {
		return nil, errors.New(500, "CREATE_TOKEN_ERROR", "账户地址参数错误")
	}

	user, err := u.uc.EthAuthorize(ctx, &biz.User{
		Address: userAddress,
	}, req)
	if err != nil {
		return nil, err
	}

	claims := auth.CustomClaims{
		UserId:   user.ID,
		UserType: "user",
		StandardClaims: jwt2.StandardClaims{
			NotBefore: time.Now().Unix(),              // 签名的生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*7, // 7天过期
			Issuer:    "Goal",
		},
	}
	token, err := auth.CreateToken(claims, u.ca.JwtKey)
	if err != nil {
		return nil, errors.New(500, "CREATE_TOKEN_ERROR", "生成token失败")
	}

	userInfoRsp := v1.EthAuthorizeReply{
		Token: token,
	}
	return &userInfoRsp, nil
}

// GetUser .
func (u *UserService) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.GetUserReply, error) {
	// 在上下文 context 中取出 claims 对象
	var userId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
		}
		userId = int64(c["UserId"].(float64))
	}

	return u.uc.GetUserWithInfoAndBalance(ctx, &biz.User{
		ID: userId,
	})
}

// Deposit .
func (u *UserService) Deposit(ctx context.Context, req *v1.DepositRequest) (*v1.DepositReply, error) {
	// 在上下文 context 中取出 claims 对象
	var userId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return nil, errors.New(500, "ERROR_TOKEN", "无效TOKEN")
		}
		userId = int64(c["UserId"].(float64))
	}

	return u.uc.Deposit(ctx, &biz.User{
		ID: userId,
	}, req)
}
