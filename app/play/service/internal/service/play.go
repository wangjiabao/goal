package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "goal/api/play/service/v1"
	"goal/app/play/service/internal/biz"
)

type PlayService struct {
	v1.UnimplementedPlayServer

	uc  *biz.PlayUseCase
	ruc *biz.RoomUseCase
	log *log.Helper
}

// NewPlayService new a game service.
func NewPlayService(uc *biz.PlayUseCase, ruc *biz.RoomUseCase, logger log.Logger) *PlayService {
	return &PlayService{uc: uc, ruc: ruc, log: log.NewHelper(logger)}
}

func (p *PlayService) AllowedPlayList(ctx context.Context, req *v1.AllowedPlayListRequest) (*v1.AllowedPlayListReply, error) {
	sortIds := []int64{1, 2, 3} // 简单点，提前入库的三个排名规则
	return p.uc.GetAdminCreateGameAndSortPlayList(ctx, req.GameId, sortIds...)
}

func (p *PlayService) RoomPlayList(ctx context.Context, req *v1.RoomPlayListRequest) (*v1.RoomPlayListReply, error) {
	_, err := p.ruc.GetRoomUserRel(ctx, req.RoomId) // 检查用户是否在房间
	if err != nil {
		return nil, err
	}

	return p.uc.GetRoomGameAndSortPlayList(ctx, req.RoomId)
}

func (p *PlayService) RoomInfo(ctx context.Context, req *v1.RoomInfoRequest) (*v1.RoomInfoReply, error) {
	_, err := p.ruc.GetRoomUserRel(ctx, req.RoomId) // 检查用户是否在房间
	if err != nil {
		return nil, err
	}

	return p.ruc.RoomInfo(ctx, req)
}

// CreatePlayGame 创建房间和比赛玩法
func (p *PlayService) CreatePlayGame(ctx context.Context, req *v1.CreatePlayGameRequest) (*v1.CreatePlayGameReply, error) {
	return p.ruc.CreatePlayGame(ctx, req)
}

// CreatePlaySort 创建房间和比赛排名玩法
func (p *PlayService) CreatePlaySort(ctx context.Context, req *v1.CreatePlaySortRequest) (*v1.CreatePlaySortReply, error) {
	return p.ruc.CreatePlaySort(ctx, req)
}

// RoomAccount 进入房间不存在用户关系记录则根据情况创建
func (p *PlayService) RoomAccount(ctx context.Context, req *v1.RoomAccountRequest) (*v1.RoomAccountReply, error) {
	return p.ruc.InRoomByAccount(ctx, req.SendBody.Account)
}

// CreateRoom 创建房间和比赛玩法
func (p *PlayService) CreateRoom(ctx context.Context, req *v1.CreateRoomRequest) (*v1.CreateRoomReply, error) {
	return p.ruc.CreateRoom(ctx, req)
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

// CreatePlayGameSort 创建比赛排名玩法竞猜
func (p *PlayService) CreatePlayGameSort(ctx context.Context, req *v1.CreatePlayGameSortRequest) (*v1.CreatePlayGameSortReply, error) {
	return p.uc.CreatePlayGameSort(ctx, req)
}

// GetUserPlayList 获取用户参与玩法记录
func (p *PlayService) GetUserPlayList(ctx context.Context, req *v1.GetUserPlayListRequest) (*v1.GetUserPlayListReply, error) {
	return p.uc.GetUserPlayList(ctx)
}

func (p *PlayService) GameUserList(ctx context.Context, req *v1.GameUserListRequest) (*v1.GameUserListReply, error) {
	sortIds := []int64{1, 2, 3} // 简单点，提前入库的三个排名规则
	return p.uc.GetAdminCreateGameAndSortPlayUserList(ctx, req.GameId, sortIds...)
}
