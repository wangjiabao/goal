package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"goal/app/admin/service/internal/biz"
	"time"
)

type Room struct {
	ID        int64     `gorm:"primarykey;type:int"`
	Account   string    `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type RoomRepo struct {
	data *Data
	log  *log.Helper
}

func NewRoomRepo(data *Data, logger log.Logger) biz.RoomRepo {
	return &RoomRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *RoomRepo) GetRoomList(ctx context.Context) ([]*biz.Room, error) {
	var room []*Room
	if err := r.data.DB(ctx).Table("room").Find(&room).Error; err != nil {
		return nil, errors.NotFound("ROOM_NOT_FOUND", "房间不存在")
	}

	res := make([]*biz.Room, 0)
	for _, item := range room {
		res = append(res, &biz.Room{
			ID:        item.ID,
			Account:   item.Account,
			CreatedAt: item.CreatedAt,
		})
	}

	return res, nil
}
