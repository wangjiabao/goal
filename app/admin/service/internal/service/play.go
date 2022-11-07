package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "goal/api/admin/service/v1"
	"goal/app/admin/service/internal/biz"
)

type PlayService struct {
	v1.UnimplementedAdminServer

	uc  *biz.PlayUseCase
	log *log.Helper
}

// NewPlayService new a game service.
func NewPlayService(uc *biz.PlayUseCase, logger log.Logger) *PlayService {
	return &PlayService{uc: uc, log: log.NewHelper(logger)}
}

func (p *PlayService) GamePlayGrant(ctx context.Context, req *v1.GamePlayGrantRequest) (*v1.GamePlayGrantReply, error) {
	return p.uc.GamePlayGrant(ctx, req)
}

// CreatePlayGame 创建房间和比赛玩法
func (p *PlayService) CreatePlayGame(ctx context.Context, req *v1.CreatePlayGameRequest) (*v1.CreatePlayGameReply, error) {
	return p.uc.CreatePlayGame(ctx, req)
}

// CreatePlaySort 创建房间和比赛排名玩法
func (p *PlayService) CreatePlaySort(ctx context.Context, req *v1.CreatePlaySortRequest) (*v1.CreatePlaySortReply, error) {
	return p.uc.CreatePlaySort(ctx, req)
}
