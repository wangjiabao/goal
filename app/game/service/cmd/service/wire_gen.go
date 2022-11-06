// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"goal/app/game/service/internal/biz"
	"goal/app/game/service/internal/conf"
	"goal/app/game/service/internal/data"
	"goal/app/game/service/internal/server"
	"goal/app/game/service/internal/service"
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
	gameRepo := data.NewGameRepo(dataData, logger)
	teamRepo := data.NewTeamRepo(dataData, logger)
	gameUseCase := biz.NewGameUseCase(gameRepo, teamRepo, logger)
	teamUseCase := biz.NewTeamUseCase(teamRepo, logger)
	displayGameRepo := data.NewDisplayGameRepo(dataData, logger)
	displayGameUseCase := biz.NewDisplayGameUseCase(displayGameRepo, gameRepo, teamRepo, logger)
	gameSortRepo := data.NewGameSortRepo(dataData, logger)
	gameSortUseCase := biz.NewGameSortUseCase(gameSortRepo, logger)
	gameService := service.NewGameService(gameUseCase, teamUseCase, displayGameUseCase, gameSortUseCase, logger)
	httpServer := server.NewHTTPServer(confServer, auth, gameService, logger)
	app := newApp(logger, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
