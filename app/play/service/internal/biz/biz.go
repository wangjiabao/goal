package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwt2 "github.com/golang-jwt/jwt/v4"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewPlayUseCase, NewRoomUseCase, NewGameUseCase, NewSortUseCase)

// Transaction 新增事务接口方法
type Transaction interface {
	ExecTx(context.Context, func(ctx context.Context) error) error
}

func getUserFromJwt(ctx context.Context) (int64, string, error) {
	// 在上下文 context 中取出 claims 对象
	var userId int64
	var userType string
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["UserId"] == nil {
			return 0, "", errors.New(500, "ERROR_TOKEN", "无效TOKEN")
		}
		if c["UserType"] == nil {
			return 0, "", errors.New(500, "ERROR_TOKEN", "无效TOKEN")
		}
		userId = int64(c["UserId"].(float64))
		userType = c["UserType"].(string)
	}

	return userId, userType, nil
}
