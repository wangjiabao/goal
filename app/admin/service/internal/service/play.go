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

func (p *PlayService) SortPlayGrant(ctx context.Context, req *v1.SortPlayGrantRequest) (*v1.SortPlayGrantReply, error) {
	return p.uc.SortPlayGrant(ctx, req)
}

// CreatePlayGame 创建房间和比赛玩法
func (p *PlayService) CreatePlayGame(ctx context.Context, req *v1.CreatePlayGameRequest) (*v1.CreatePlayGameReply, error) {
	return p.uc.CreatePlayGame(ctx, req)
}

// CreatePlaySort 创建房间和比赛排名玩法
func (p *PlayService) CreatePlaySort(ctx context.Context, req *v1.CreatePlaySortRequest) (*v1.CreatePlaySortReply, error) {
	return p.uc.CreatePlaySort(ctx, req)
}

// GetPlayList .
func (p *PlayService) GetPlayList(ctx context.Context, req *v1.GetPlayListRequest) (*v1.GetPlayListReply, error) {
	return p.uc.GetPlayList(ctx, req)
}

// GetPlayRelList .
func (p *PlayService) GetPlayRelList(ctx context.Context, req *v1.GetPlayRelListRequest) (*v1.GetPlayRelListReply, error) {
	return p.uc.GetPlayUserRelList(ctx, req)
}

// GetRoomList .
func (p *PlayService) GetRoomList(ctx context.Context, req *v1.GetRoomListRequest) (*v1.GetRoomListReply, error) {
	return p.uc.GetRooms(ctx)
}

// GetRoomPlayList .
func (p *PlayService) GetRoomPlayList(ctx context.Context, req *v1.GetRoomPlayListRequest) (*v1.GetRoomPlayListReply, error) {
	return p.uc.GetRoomPlayList(ctx, req)
}
