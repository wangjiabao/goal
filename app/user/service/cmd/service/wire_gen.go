// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"goal/app/user/service/internal/biz"
	"goal/app/user/service/internal/conf"
	"goal/app/user/service/internal/data"
	"goal/app/user/service/internal/server"
	"goal/app/user/service/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, auth *conf.Auth, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewDB(confData)
	client := data.NewRedis(confData)
	dataData, cleanup, err := data.NewData(confData, logger, db, client)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	transaction := data.NewTransaction(dataData)
	systemConfigRepo := data.NewSystemConfigRepo(dataData, logger)
	addressEthBalanceRepo := data.NewAddressEthBalanceRepo(dataData, logger)
	roleRepo := data.NewRoleRepo(dataData, logger)
	userInfoRepo := data.NewUserInfoRepo(dataData, logger)
	userBalanceRepo := data.NewUserBalanceRepo(dataData, logger)
	userUseCase := biz.NewUserUseCase(userRepo, transaction, systemConfigRepo, addressEthBalanceRepo, roleRepo, userInfoRepo, userBalanceRepo, logger)
	userService := service.NewUserService(userUseCase, logger, auth)
	httpServer := server.NewHTTPServer(confServer, auth, userService, logger)
	app := newApp(logger, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
