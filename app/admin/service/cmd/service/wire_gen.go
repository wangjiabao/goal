// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"goal/app/admin/service/internal/biz"
	"goal/app/admin/service/internal/conf"
	"goal/app/admin/service/internal/data"
	"goal/app/admin/service/internal/server"
	"goal/app/admin/service/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewDB(confData)
	client := data.NewRedis(confData)
	dataData, cleanup, err := data.NewData(confData, logger, db, client)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	playRepo := data.NewPlayRepo(dataData, logger)
	roomRepo := data.NewRoomRepo(dataData, logger)
	playGameRelRepo := data.NewPlayGameRelRepo(dataData, logger)
	playSortRelRepo := data.NewPlaySortRelRepo(dataData, logger)
	playRoomRelRepo := data.NewPlayRoomRelRepo(dataData, logger)
	playGameScoreUserRelRepo := data.NewPlayGameScoreUserRelRepo(dataData, logger)
	playGameTeamGoalUserRelRepo := data.NewPlayGameTeamGoalUserRelRepo(dataData, logger)
	playGameTeamResultUserRelRepo := data.NewPlayGameTeamResultUserRelRepo(dataData, logger)
	playGameTeamSortUserRelRepo := data.NewPlayGameTeamSortUserRelRepo(dataData, logger)
	gameRepo := data.NewGameRepo(dataData, logger)
	sortRepo := data.NewSortRepo(dataData, logger)
	userBalanceRepo := data.NewUserBalanceRepo(dataData, logger)
	userProxyRepo := data.NewUserProxyRepo(dataData, logger)
	userInfoRepo := data.NewUserInfoRepo(dataData, logger)
	transaction := data.NewTransaction(dataData)
	playUseCase := biz.NewPlayUseCase(userRepo, playRepo, roomRepo, playGameRelRepo, playSortRelRepo, playRoomRelRepo, playGameScoreUserRelRepo, playGameTeamGoalUserRelRepo, playGameTeamResultUserRelRepo, playGameTeamSortUserRelRepo, gameRepo, sortRepo, userBalanceRepo, userProxyRepo, userInfoRepo, transaction, logger)
	playService := service.NewPlayService(playUseCase, logger)
	userUseCase := biz.NewUserUseCase(userRepo, transaction, userInfoRepo, userBalanceRepo, logger)
	userService := service.NewUserService(userUseCase, logger)
	gameUseCase := biz.NewGameUseCase(gameRepo, playGameRelRepo, playRepo, playGameScoreUserRelRepo, playGameTeamGoalUserRelRepo, playGameTeamResultUserRelRepo, logger)
	teamRepo := data.NewTeamRepo(dataData, logger)
	teamUseCase := biz.NewTeamUseCase(teamRepo, logger)
	sortUseCase := biz.NewSortUseCase(sortRepo, logger)
	gameService := service.NewGameService(gameUseCase, teamUseCase, sortUseCase, logger)
	httpServer := server.NewHTTPServer(confServer, playService, userService, gameService, logger)
	app := newApp(logger, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
