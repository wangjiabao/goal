package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type Game struct {
	ID         int64
	Name       string
	RedTeamId  int64
	BlueTeamId int64
	StartTime  time.Time
	EndTime    time.Time
}

type GameRepo interface {
	GetGameById(ctx context.Context, gameId int64) (*Game, error)
}

type GameUseCase struct {
	gameRepo GameRepo
	log      *log.Helper
}

func NewGameUseCase(repo GameRepo, logger log.Logger) *GameUseCase {
	return &GameUseCase{gameRepo: repo, log: log.NewHelper(logger)}
}
