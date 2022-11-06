package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "goal/api/game/service/v1"
	"goal/app/game/service/internal/biz"
)

type GameService struct {
	v1.UnimplementedGameServer

	dc  *biz.DisplayGameUseCase
	gc  *biz.GameUseCase
	tc  *biz.TeamUseCase
	gsc *biz.GameSortUseCase
	log *log.Helper
}

// NewGameService new a game service.
func NewGameService(gc *biz.GameUseCase, tc *biz.TeamUseCase, dc *biz.DisplayGameUseCase, gsc *biz.GameSortUseCase, logger log.Logger) *GameService {
	return &GameService{gc: gc, dc: dc, tc: tc, gsc: gsc, log: log.NewHelper(logger)}
}

func (g *GameService) DisplayGame(ctx context.Context, req *v1.DisplayGameRequest) (*v1.DisplayGameReply, error) {
	return g.dc.GetDisplayGame(ctx, req.Type)
}

func (g *GameService) GetGameList(ctx context.Context, req *v1.GetGameListRequest) (*v1.GetGameListReply, error) {
	return g.gc.GetGameList(ctx)
}

func (g *GameService) GetGameSortList(ctx context.Context, req *v1.GetGameSortListRequest) (*v1.GetGameSortListReply, error) {
	return g.gsc.GetGameSortList(ctx)
}

func (g *GameService) GetTeamList(ctx context.Context, req *v1.GetTeamListRequest) (*v1.GetTeamListReply, error) {
	return g.tc.GetTeamList(ctx)
}
