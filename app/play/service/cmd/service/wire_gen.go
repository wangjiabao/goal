// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"goal/app/play/service/internal/biz"
	"goal/app/play/service/internal/conf"
	"goal/app/play/service/internal/data"
	"goal/app/play/service/internal/server"
	"goal/app/play/service/internal/service"
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
	playRepo := data.NewPlayRepo(dataData, logger)
	playGameRelRepo := data.NewPlayGameRelRepo(dataData, logger)
	systemConfigRepo := data.NewSystemConfigRepo(dataData, logger)
	playSortRelRepo := data.NewPlaySortRelRepo(dataData, logger)
	playRoomRelRepo := data.NewPlayRoomRelRepo(dataData, logger)
	roomUserRelRepo := data.NewRoomUserRelRepo(dataData, logger)
	roomRepo := data.NewRoomRepo(dataData, logger)
	playGameScoreUserRelRepo := data.NewPlayGameScoreUserRelRepo(dataData, logger)
	playGameTeamSortUserRelRepo := data.NewPlayGameTeamSortUserRelRepo(dataData, logger)
	playGameTeamGoalUserRelRepo := data.NewPlayGameTeamGoalUserRelRepo(dataData, logger)
	playGameTeamResultUserRelRepo := data.NewPlayGameTeamResultUserRepo(dataData, logger)
	userBalanceRepo := data.NewUserBalanceRepo(dataData, logger)
	userProxyRepo := data.NewUserProxyRepo(dataData, logger)
	transaction := data.NewTransaction(dataData)
	playUseCase := biz.NewPlayUseCase(playRepo, playGameRelRepo, systemConfigRepo, playSortRelRepo, playRoomRelRepo, roomUserRelRepo, roomRepo, playGameScoreUserRelRepo, playGameTeamSortUserRelRepo, playGameTeamGoalUserRelRepo, playGameTeamResultUserRelRepo, userBalanceRepo, userProxyRepo, transaction, logger)
	roomGameRelRepo := data.NewRoomGameRelRepo(dataData, logger)
	gameRepo := data.NewGameRepo(dataData, logger)
	sortRepo := data.NewSortRepo(dataData, logger)
	roomUseCase := biz.NewRoomUseCase(roomRepo, roomUserRelRepo, roomGameRelRepo, playRepo, gameRepo, playSortRelRepo, playRoomRelRepo, playGameRelRepo, sortRepo, transaction, logger)
	playService := service.NewPlayService(playUseCase, roomUseCase, logger)
	httpServer := server.NewHTTPServer(confServer, auth, playService, logger)
	app := newApp(logger, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
