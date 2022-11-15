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

// CreatePlaySort 创建比赛排名玩法
func (p *PlayService) CreatePlaySort(ctx context.Context, req *v1.CreatePlaySortRequest) (*v1.CreatePlaySortReply, error) {
	return p.uc.CreatePlaySort(ctx, req)
}

// DeletePlayGame 删除比赛玩法
func (p *PlayService) DeletePlayGame(ctx context.Context, req *v1.DeletePlayGameRequest) (*v1.DeletePlayGameReply, error) {
	return p.uc.DeletePlayGame(ctx, req)
}

// DeletePlaySort 删除比赛排名
func (p *PlayService) DeletePlaySort(ctx context.Context, req *v1.DeletePlaySortRequest) (*v1.DeletePlaySortReply, error) {
	return p.uc.DeletePlaySort(ctx, req)
}

// GetPlayList .
func (p *PlayService) GetPlayList(ctx context.Context, req *v1.GetPlayListRequest) (*v1.GetPlayListReply, error) {
	return p.uc.GetPlayList(ctx, req)
}

// GetPlaySortList .
func (p *PlayService) GetPlaySortList(ctx context.Context, req *v1.GetPlaySortListRequest) (*v1.GetPlaySortListReply, error) {
	return p.uc.GetPlaySortList(ctx, req)
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

// CreatePlayGameScore 创建比分玩法竞猜
func (p *PlayService) CreatePlayGameScore(ctx context.Context, req *v1.CreatePlayGameScoreRequest) (*v1.CreatePlayGameScoreReply, error) {
	return p.uc.CreatePlayGameScore(ctx, req)
}

// CreatePlayGameResult 创建输赢平玩法竞猜
func (p *PlayService) CreatePlayGameResult(ctx context.Context, req *v1.CreatePlayGameResultRequest) (*v1.CreatePlayGameResultReply, error) {
	return p.uc.CreatePlayGameResult(ctx, req)
}

// CreatePlayGameGoal 创建进球数玩法竞猜
func (p *PlayService) CreatePlayGameGoal(ctx context.Context, req *v1.CreatePlayGameGoalRequest) (*v1.CreatePlayGameGoalReply, error) {
	return p.uc.CreatePlayGameGoal(ctx, req)
}

// GetConfigList .
func (p *PlayService) GetConfigList(ctx context.Context, req *v1.GetConfigListRequest) (*v1.GetConfigListReply, error) {
	return p.uc.GetConfigList(ctx)
}

// UpdateConfig .
func (p *PlayService) UpdateConfig(ctx context.Context, req *v1.UpdateConfigRequest) (*v1.UpdateConfigReply, error) {
	return p.uc.UpdateConfig(ctx, req)
}
