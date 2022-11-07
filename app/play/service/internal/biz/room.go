package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "goal/api/play/service/v1"
	"strconv"
	"time"
)

type Room struct {
	ID           int64
	CreateUserId int64
	Account      string
	Type         string
}

type RoomUserRel struct {
	ID     int64
	RoomId int64
	UserId int64
}
type RoomRepo interface {
	GetRoomByAccount(ctx context.Context, account string) (*Room, error)
	GetRoomByID(ctx context.Context, roomId int64) (*Room, error)
	CreateRoom(ctx context.Context, rc *Room) (*Room, error)
}

type RoomUserRelRepo interface {
	GetRoomUserRel(ctx context.Context, userId int64, roomId int64) (*RoomUserRel, error)
	GetRoomUserRelByRoomId(ctx context.Context, roomId int64) (map[int64]*RoomUserRel, error)
	CreateRoomUserRel(ctx context.Context, userId int64, roomId int64) (*RoomUserRel, error)
}

type RoomUseCase struct {
	roomRepo        RoomRepo
	playRepo        PlayRepo
	gameRepo        GameRepo
	sortRepo        SortRepo
	playRoomRelRepo PlayRoomRelRepo
	roomUserRelRepo RoomUserRelRepo
	playGameRelRepo PlayGameRelRepo
	playSortRelRepo PlaySortRelRepo
	tx              Transaction
	log             *log.Helper
}

func NewRoomUseCase(repo RoomRepo, roomUserRelRepo RoomUserRelRepo, playRepo PlayRepo, gameRepo GameRepo, playSortRelRepo PlaySortRelRepo, playRoomRelRepo PlayRoomRelRepo, playGameRelRepo PlayGameRelRepo, sortRepo SortRepo, tx Transaction, logger log.Logger) *RoomUseCase {
	return &RoomUseCase{roomRepo: repo, playRepo: playRepo, roomUserRelRepo: roomUserRelRepo, gameRepo: gameRepo, playSortRelRepo: playSortRelRepo, playRoomRelRepo: playRoomRelRepo, playGameRelRepo: playGameRelRepo, sortRepo: sortRepo, tx: tx, log: log.NewHelper(logger)}
}

// GetRoomUserRel 获取房间内用户
func (r *RoomUseCase) GetRoomUserRel(ctx context.Context, roomId int64) (*RoomUserRel, error) {
	userId, _, err := getUserFromJwt(ctx) // 获取用户id
	if nil != err {
		return nil, err
	}

	return r.roomUserRelRepo.GetRoomUserRel(ctx, userId, roomId)
}

// InRoomByAccount 进入房间根据房间号
func (r *RoomUseCase) InRoomByAccount(ctx context.Context, account string) (*v1.RoomAccountReply, error) {
	var (
		userId         int64
		room           *Room
		roomUserRel    *RoomUserRel
		roomUserRelMap map[int64]*RoomUserRel
		err            error
	)
	userId, _, err = getUserFromJwt(ctx) // 获取用户id
	if nil != err {
		return nil, err
	}

	room, err = r.roomRepo.GetRoomByAccount(ctx, account) // 获取房间
	if nil != err {
		return nil, err
	}

	roomUserRelMap, err = r.roomUserRelRepo.GetRoomUserRelByRoomId(ctx, room.ID) // 获取房间内用户
	if nil != err {
		return nil, err
	}

	if v, ok := roomUserRelMap[userId]; ok { // 在房间里
		return &v1.RoomAccountReply{
			RoomId:   v.RoomId,
			RoomType: room.Type,
		}, nil
	}

	if roomUserCount := len(roomUserRelMap); roomUserCount > 2 {
		return nil, errors.New(500, "ROOM_USER_FULL", "房间已满3人")
	}

	roomUserRel, err = r.roomUserRelRepo.CreateRoomUserRel(ctx, userId, room.ID) // 创建房间用户关系
	return &v1.RoomAccountReply{
		RoomId:   roomUserRel.RoomId,
		RoomType: room.Type,
	}, nil
}

// CreateRoom 创建房间
func (r *RoomUseCase) CreateRoom(ctx context.Context, req *v1.CreateRoomRequest) (*v1.CreateRoomReply, error) {
	var (
		userId      int64
		room        *Room
		roomUserRel *RoomUserRel
		err         error
	)

	if "game_team_goal" != req.SendBody.RoomType && // 验证type类型
		"game_score" != req.SendBody.RoomType &&
		"game_team_sort" != req.SendBody.RoomType &&
		"game_team_result" != req.SendBody.RoomType {
		return nil, errors.New(500, "TIME_ERROR", "房间类型输入错误")
	}

	userId, _, err = getUserFromJwt(ctx) // 获取用户id
	if nil != err {
		return nil, err
	}

	err = r.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		room, err = r.roomRepo.CreateRoom(ctx, &Room{ // 新增房间
			CreateUserId: userId,
			Account:      strconv.FormatInt(time.Now().UnixNano(), 10) + strconv.FormatInt(userId, 10),
			Type:         req.SendBody.RoomType,
		})
		if err != nil {
			return err
		}

		roomUserRel, err = r.roomUserRelRepo.CreateRoomUserRel(ctx, userId, room.ID) // 新增用户和房间的关系
		if err != nil {
			return err
		}

		return nil
	})

	return &v1.CreateRoomReply{
		Account:  room.Account,
		RoomId:   room.ID,
		RoomType: room.Type,
	}, err
}
