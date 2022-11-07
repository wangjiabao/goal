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
	log *log.Helper
}

// NewGameService new a game service.
func NewGameService(uc *biz.GameUseCase, suc *biz.SortUseCase, logger log.Logger) *GameService {
	return &GameService{uc: uc, suc: suc, log: log.NewHelper(logger)}
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

func (g *GameService) GetGameSortList(ctx context.Context, req *v1.GetGameSortListRequest) (*v1.GetGameSortListReply, error) {
	return g.suc.GetGameSortList(ctx)
}
