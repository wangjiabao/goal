package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "goal/api/admin/service/v1"
	"goal/app/admin/service/internal/biz"
)

type GameService struct {
	v1.UnimplementedGameServer

	uc  *biz.GameUseCase
	suc *biz.SortUseCase
	tc  *biz.TeamUseCase
	log *log.Helper
}

// NewGameService new a game service.
func NewGameService(uc *biz.GameUseCase, tc *biz.TeamUseCase, suc *biz.SortUseCase, logger log.Logger) *GameService {
	return &GameService{uc: uc, suc: suc, tc: tc, log: log.NewHelper(logger)}
}

func (g *GameService) CreateGame(ctx context.Context, req *v1.CreateGameRequest) (*v1.CreateGameReply, error) {
	return g.uc.CreateGame(ctx, req)
}

func (g *GameService) UpdateGame(ctx context.Context, req *v1.UpdateGameRequest) (*v1.UpdateGameReply, error) {
	return g.uc.UpdateGame(ctx, req)
}

func (g *GameService) GetGameList(ctx context.Context, req *v1.GetGameListRequest) (*v1.GetGameListReply, error) {
	return g.uc.GetGameList(ctx)
}

func (g *GameService) GetGame(ctx context.Context, req *v1.GetGameRequest) (*v1.GetGameReply, error) {
	return g.uc.GetGameByGameId(ctx, req)
}

func (g *GameService) DisplayGameIndex(ctx context.Context, req *v1.DisplayGameIndexRequest) (*v1.DisplayGameIndexReply, error) {
	return g.uc.GetDisplayGameIndex(ctx)
}

func (g *GameService) SaveDisplayGameIndex(ctx context.Context, req *v1.SaveDisplayGameIndexRequest) (*v1.SaveDisplayGameIndexReply, error) {
	return g.uc.SaveDisplayGameIndex(ctx, req)
}

func (g *GameService) DeleteDisplayGameIndex(ctx context.Context, req *v1.DeleteDisplayGameIndexRequest) (*v1.DeleteDisplayGameIndexReply, error) {
	return g.uc.DeleteDisplayGameIndex(ctx, req)
}

func (g *GameService) GetGameSortList(ctx context.Context, req *v1.GetGameSortListRequest) (*v1.GetGameSortListReply, error) {
	return g.suc.GetGameSortList(ctx)
}

func (g *GameService) CreateSort(ctx context.Context, req *v1.CreateSortRequest) (*v1.CreateSortReply, error) {
	return g.suc.CreateSort(ctx, req)
}

func (g *GameService) UpdateSort(ctx context.Context, req *v1.UpdateSortRequest) (*v1.UpdateSortReply, error) {
	return g.suc.UpdateSort(ctx, req)
}

func (g *GameService) GetTeamList(ctx context.Context, req *v1.GetTeamListRequest) (*v1.GetTeamListReply, error) {
	return g.tc.GetTeamList(ctx)
}

func (g *GameService) CreateTeam(ctx context.Context, req *v1.CreateTeamRequest) (*v1.CreateTeamReply, error) {
	return g.tc.CreateTeam(ctx, req)
}

func (g *GameService) GameIndexStatistics(ctx context.Context, req *v1.GameIndexStatisticsRequest) (*v1.GameIndexStatisticsReply, error) {
	return g.uc.GameIndexStatistics(ctx, req)
}
