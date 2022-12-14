package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"goal/app/play/service/internal/biz"
	"gorm.io/gorm"
	"time"
)

type Room struct {
	ID           int64     `gorm:"primarykey;type:int"`
	CreateUserId int64     `gorm:"type:int;not null"`
	Account      string    `gorm:"type:varchar(45);not null"`
	Type         string    `gorm:"type:varchar(45);not null"`
	CreatedAt    time.Time `gorm:"type:datetime;not null"`
	UpdatedAt    time.Time `gorm:"type:datetime;not null"`
}

type RoomUserRel struct {
	ID        int64     `gorm:"primarykey;type:int"`
	UserId    int64     `gorm:"type:int;not null"`
	RoomId    int64     `gorm:"type:int;not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type RoomGameRel struct {
	ID        int64     `gorm:"primarykey;type:int"`
	GameId    int64     `gorm:"type:int;not null"`
	RoomId    int64     `gorm:"type:int;not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type RoomRepo struct {
	data *Data
	log  *log.Helper
}

type RoomUserRelRepo struct {
	data *Data
	log  *log.Helper
}

type RoomGameRelRepo struct {
	data *Data
	log  *log.Helper
}

func NewRoomRepo(data *Data, logger log.Logger) biz.RoomRepo {
	return &RoomRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func NewRoomUserRelRepo(data *Data, logger log.Logger) biz.RoomUserRelRepo {
	return &RoomUserRelRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func NewRoomGameRelRepo(data *Data, logger log.Logger) biz.RoomGameRelRepo {
	return &RoomGameRelRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *RoomUserRelRepo) GetRoomUserRel(ctx context.Context, userId int64, roomId int64) (*biz.RoomUserRel, error) {
	var roomUserRel RoomUserRel
	if err := r.data.DB(ctx).Where(&RoomUserRel{RoomId: roomId, UserId: userId}).Table("room_user_rel").First(&roomUserRel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("ROOM_USER_REL_NOT_FOUND", "?????????????????????")
		}

		return nil, errors.New(500, "ROOM_USER_REL_NOT_FOUND", err.Error())
	}

	return &biz.RoomUserRel{
		ID:     roomUserRel.ID,
		RoomId: roomUserRel.RoomId,
		UserId: roomUserRel.UserId,
	}, nil
}

func (r *RoomUserRelRepo) GetRoomUserRelByRoomId(ctx context.Context, roomId int64) (map[int64]*biz.RoomUserRel, error) {
	var l []*RoomUserRel
	if result := r.data.DB(ctx).Table("room_user_rel").Where(&RoomUserRel{RoomId: roomId}).Find(&l); result.Error != nil {
		return nil, errors.InternalServer("ROOM_USER_REL_NOT_FOUND", "?????????????????????????????????")
	}

	pl := make(map[int64]*biz.RoomUserRel)
	for _, v := range l {
		pl[v.UserId] = &biz.RoomUserRel{
			ID:     v.ID,
			UserId: v.UserId,
			RoomId: v.RoomId,
		}
	}
	return pl, nil
}

func (r *RoomUserRelRepo) GetRoomUsers(ctx context.Context, roomId int64) ([]*biz.RoomUserRel, error) {
	var l []*RoomUserRel
	if result := r.data.DB(ctx).Table("room_user_rel").Where(&RoomUserRel{RoomId: roomId}).Find(&l); result.Error != nil {
		return nil, errors.InternalServer("ROOM_USER_REL_NOT_FOUND", "?????????????????????????????????")
	}

	pl := make([]*biz.RoomUserRel, 0)
	for _, v := range l {
		pl = append(pl, &biz.RoomUserRel{
			ID:     v.ID,
			UserId: v.UserId,
			RoomId: v.RoomId,
		})
	}
	return pl, nil
}

func (r *RoomUserRelRepo) GetRoomByUserId(ctx context.Context, userId int64) ([]*biz.RoomUserRel, error) {
	var l []*RoomUserRel
	if result := r.data.DB(ctx).Table("room_user_rel").
		Where(&RoomUserRel{UserId: userId}).
		Order("created_at desc").
		Find(&l); result.Error != nil {
		return nil, errors.InternalServer("ROOM_USER_REL_NOT_FOUND", "?????????????????????????????????")
	}

	pl := make([]*biz.RoomUserRel, 0)
	for _, v := range l {
		pl = append(pl, &biz.RoomUserRel{
			ID:     v.ID,
			UserId: v.UserId,
			RoomId: v.RoomId,
		})
	}
	return pl, nil
}

func (r *RoomRepo) GetUserByUseIds(ctx context.Context, userIds ...int64) (map[int64]*biz.User, error) {
	var l []*User
	if result := r.data.DB(ctx).Table("user").Where("id IN(?)", userIds).Find(&l); result.Error != nil {
		return nil, errors.InternalServer("ROOM_USER_REL_NOT_FOUND", "???????????????????????????")
	}

	pl := make(map[int64]*biz.User)
	for _, v := range l {
		pl[v.ID] = &biz.User{
			ID:      v.ID,
			Address: v.Address,
		}
	}
	return pl, nil
}

func (r *RoomRepo) GetRoomByIds(ctx context.Context, roomIds ...int64) ([]*biz.Room, error) {
	var l []*Room
	if result := r.data.DB(ctx).Table("room").Where("id IN(?)", roomIds).Order("created_at desc").Find(&l); result.Error != nil {
		return nil, errors.InternalServer("ROOM_USER_REL_NOT_FOUND", "??????????????????")
	}

	pl := make([]*biz.Room, 0)
	for _, v := range l {
		pl = append(pl, &biz.Room{
			ID:        v.ID,
			Account:   v.Account,
			CreatedAt: v.CreatedAt,
		})
	}
	return pl, nil
}

func (r *RoomRepo) GetRoomByAccount(ctx context.Context, account string) (*biz.Room, error) {
	var room Room
	if err := r.data.DB(ctx).Where(&Room{Account: account}).Table("room").First(&room).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("ROOM_NOT_FOUND", "???????????????")
		}

		return nil, errors.New(500, "ROOM_NOT_FOUND", err.Error())
	}

	return &biz.Room{
		ID:   room.ID,
		Type: room.Type,
	}, nil
}

func (r *RoomRepo) GetRoomByID(ctx context.Context, roomId int64) (*biz.Room, error) {
	var room Room
	if err := r.data.DB(ctx).Where(&Room{ID: roomId}).Table("room").First(&room).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("ROOM_USER_REL_NOT_FOUND", "???????????????")
		}

		return nil, errors.New(500, "ROOM_USER_REL_NOT_FOUND", err.Error())
	}

	return &biz.Room{
		ID:           room.ID,
		Type:         room.Type,
		CreateUserId: room.CreateUserId,
	}, nil
}

// CreateRoomUserRel .
func (r *RoomUserRelRepo) CreateRoomUserRel(ctx context.Context, userId int64, roomId int64) (*biz.RoomUserRel, error) {
	var roomUserRel RoomUserRel
	roomUserRel.UserId = userId
	roomUserRel.RoomId = roomId
	res := r.data.DB(ctx).Table("room_user_rel").Create(&roomUserRel)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_ROOM_USER_REL_ERROR", "?????????????????????????????????")
	}

	return &biz.RoomUserRel{
		ID:     roomUserRel.ID,
		UserId: roomUserRel.UserId,
		RoomId: roomUserRel.RoomId,
	}, nil
}

// CreateRoomGameRel .
func (r *RoomGameRelRepo) CreateRoomGameRel(ctx context.Context, gameId int64, roomId int64) (*biz.RoomGameRel, error) {
	var roomGameRel RoomGameRel
	roomGameRel.GameId = gameId
	roomGameRel.RoomId = roomId
	res := r.data.DB(ctx).Table("room_game_rel").Create(&roomGameRel)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_ROOM_USER_REL_ERROR", "?????????????????????????????????")
	}

	return &biz.RoomGameRel{
		ID:     roomGameRel.ID,
		GameId: roomGameRel.GameId,
		RoomId: roomGameRel.RoomId,
	}, nil
}

func (r *RoomGameRelRepo) GetRoomGame(ctx context.Context, roomId int64) (*biz.RoomGameRel, error) {
	var roomGameRel RoomGameRel
	if result := r.data.DB(ctx).Table("room_game_rel").Where(&RoomUserRel{RoomId: roomId}).First(&roomGameRel); result.Error != nil {
		return nil, errors.InternalServer("ROOM_USER_REL_NOT_FOUND", "??????????????????????????????")
	}

	return &biz.RoomGameRel{
		RoomId: roomGameRel.RoomId,
		GameId: roomGameRel.GameId,
	}, nil
}

// CreateRoom .
func (r *RoomRepo) CreateRoom(ctx context.Context, rc *biz.Room) (*biz.Room, error) {
	var room Room
	room.CreateUserId = rc.CreateUserId
	room.Account = rc.Account
	room.Type = rc.Type
	res := r.data.DB(ctx).Table("room").Create(&room)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_ROOM_ERROR", "??????????????????")
	}

	return &biz.Room{
		ID:           room.ID,
		CreateUserId: room.CreateUserId,
		Account:      room.Account,
		Type:         room.Type,
	}, nil
}
