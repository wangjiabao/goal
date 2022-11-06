package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"goal/app/play/service/internal/biz"
	"gorm.io/gorm"
	"time"
)

type Sort struct {
	ID        int64     `gorm:"primarykey;type:int"`
	SortName  string    `gorm:"type:varchar(45);not null"`
	Type      string    `gorm:"type:varchar(45);not null"`
	EndTime   time.Time `gorm:"type:datetime;not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type SortRepo struct {
	data *Data
	log  *log.Helper
}

func NewSortGRepo(data *Data, logger log.Logger) biz.SortRepo {
	return &SortRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (s *SortRepo) GetGameSortById(ctx context.Context, gameId int64) (*biz.Sort, error) {
	var sort Sort
	if err := s.data.DB(ctx).Where(&Game{ID: gameId}).Table("soccer_game_team_sort").First(&sort).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("GAME_SORT_NOT_FOUND", "game sort not found")
		}

		return nil, errors.New(500, "GAME_SORT_NOT_FOUND", err.Error())
	}

	return &biz.Sort{
		ID:      sort.ID,
		EndTime: sort.EndTime,
		Type:    sort.Type,
	}, nil
}
