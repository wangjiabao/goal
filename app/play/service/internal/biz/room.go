package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "goal/api/play/service/v1"
	"strconv"
	"time"
)

type User struct {
	ID                  int64
	Address             string
	ToAddress           string
	ToAddressPrivateKey string
}

type Room struct {
	ID           int64
	CreateUserId int64
	Account      string
	CreatedAt    time.Time
	Type         string
}

type RoomUserRel struct {
	ID     int64
	RoomId int64
	UserId int64
}

type RoomGameRel struct {
	ID     int64
	RoomId int64
	GameId int64
}

type RoomRepo interface {
	GetRoomByAccount(ctx context.Context, account string) (*Room, error)
	GetRoomByID(ctx context.Context, roomId int64) (*Room, error)
	CreateRoom(ctx context.Context, rc *Room) (*Room, error)
	GetUserByUseIds(ctx context.Context, userIds ...int64) (map[int64]*User, error)
	GetRoomByIds(ctx context.Context, roomIds ...int64) ([]*Room, error)
}

type RoomUserRelRepo interface {
	GetRoomUserRel(ctx context.Context, userId int64, roomId int64) (*RoomUserRel, error)
	GetRoomUserRelByRoomId(ctx context.Context, roomId int64) (map[int64]*RoomUserRel, error)
	CreateRoomUserRel(ctx context.Context, userId int64, roomId int64) (*RoomUserRel, error)
	GetRoomUsers(ctx context.Context, roomId int64) ([]*RoomUserRel, error)
	GetRoomByUserId(ctx context.Context, userId int64) ([]*RoomUserRel, error)
}

type RoomGameRelRepo interface {
	GetRoomGame(ctx context.Context, roomId int64) (*RoomGameRel, error)
	CreateRoomGameRel(ctx context.Context, gameId int64, roomId int64) (*RoomGameRel, error)
}

type RoomUseCase struct {
	roomRepo         RoomRepo
	playRepo         PlayRepo
	gameRepo         GameRepo
	sortRepo         SortRepo
	systemConfigRepo SystemConfigRepo
	playRoomRelRepo  PlayRoomRelRepo
	roomUserRelRepo  RoomUserRelRepo
	playGameRelRepo  PlayGameRelRepo
	roomGameRelRepo  RoomGameRelRepo
	playSortRelRepo  PlaySortRelRepo
	userBalanceRepo  UserBalanceRepo
	tx               Transaction
	log              *log.Helper
}

func NewRoomUseCase(repo RoomRepo, roomUserRelRepo RoomUserRelRepo, userBalanceRepo UserBalanceRepo, systemConfigRepo SystemConfigRepo, roomGameRelRepo RoomGameRelRepo, playRepo PlayRepo, gameRepo GameRepo, playSortRelRepo PlaySortRelRepo, playRoomRelRepo PlayRoomRelRepo, playGameRelRepo PlayGameRelRepo, sortRepo SortRepo, tx Transaction, logger log.Logger) *RoomUseCase {
	return &RoomUseCase{roomRepo: repo, playRepo: playRepo, systemConfigRepo: systemConfigRepo, userBalanceRepo: userBalanceRepo, roomGameRelRepo: roomGameRelRepo, roomUserRelRepo: roomUserRelRepo, gameRepo: gameRepo, playSortRelRepo: playSortRelRepo, playRoomRelRepo: playRoomRelRepo, playGameRelRepo: playGameRelRepo, sortRepo: sortRepo, tx: tx, log: log.NewHelper(logger)}
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
		userId       int64
		room         *Room
		roomUserRel  *RoomUserRel
		roomGameRel  *RoomGameRel
		systemConfig *SystemConfig
		base         int64 = 100000
		err          error
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

	systemConfig, err = r.systemConfigRepo.GetSystemConfigByName(ctx, "room_rate")
	if nil != err {
		return nil, err
	}
	fee := systemConfig.Value * base

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

		roomGameRel, err = r.roomGameRelRepo.CreateRoomGameRel(ctx, req.SendBody.GameId, room.ID)
		if err != nil {
			return err
		}

		var recordId int64
		recordId, err = r.userBalanceRepo.RoomFee(ctx, userId, fee) // 余额扣除，先扣后加以防，给用户是代理余额增长了
		if err != nil {
			return err
		}

		err = r.userBalanceRepo.CreateBalanceRecordIdRel(ctx, recordId, "room_fee", room.ID)
		if err != nil {
			return err
		}

		return nil
	})

	return &v1.CreateRoomReply{
		Account:  room.Account,
		RoomId:   room.ID,
		GameId:   req.SendBody.GameId,
		RoomType: room.Type,
	}, err
}

// RoomInfo 房间内信息
func (r *RoomUseCase) RoomInfo(ctx context.Context, req *v1.RoomInfoRequest) (*v1.RoomInfoReply, error) {
	var (
		room           *Room
		roomUserRel    []*RoomUserRel
		roomGameRel    *RoomGameRel
		createRoomUser bool = false
		userIds        []int64
		users          map[int64]*User
	)

	userId, _, err := getUserFromJwt(ctx) // 获取用户id
	if nil != err {
		return nil, err
	}

	room, err = r.roomRepo.GetRoomByID(ctx, req.RoomId)
	if nil != err {
		return nil, errors.New(500, "ROOM_USER_ERROR", "房间信息错误")
	}

	roomUserRel, err = r.roomUserRelRepo.GetRoomUsers(ctx, room.ID)
	if nil != err {
		return nil, errors.New(500, "ROOM_USER_ERROR", "房间信息错误")
	}

	roomGameRel, err = r.roomGameRelRepo.GetRoomGame(ctx, room.ID)
	if nil != err {
		return nil, errors.New(500, "ROOM_USER_ERROR", "房间信息错误")
	}

	for _, v := range roomUserRel {
		userIds = append(userIds, v.UserId)
	}

	users, _ = r.roomRepo.GetUserByUseIds(ctx, userIds...)

	if room.CreateUserId == userId {
		createRoomUser = true
	}

	res := &v1.RoomInfoReply{
		CreatedRoomUser: createRoomUser,
		GameId:          roomGameRel.GameId,
		Users:           make([]*v1.RoomInfoReply_User, 0),
	}

	for _, v := range roomUserRel {
		tmpAddress := ""
		if tmpUser, ok := users[v.UserId]; ok {
			tmpAddress = tmpUser.Address
		}
		res.Users = append(res.Users, &v1.RoomInfoReply_User{
			Address: tmpAddress,
			ID:      v.UserId,
		})
	}

	return res, nil
}
