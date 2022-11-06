package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"goal/app/game/service/internal/biz"
	"time"
)

type GameSort struct {
	ID        int64     `gorm:"primarykey;type:int"`
	SortName  string    `gorm:"type:varchar(45);not null"`
	Type      string    `gorm:"type:varchar(45);not null"`
	EndTime   time.Time `gorm:"type:datetime;not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type GameSortRepo struct {
	data *Data
	log  *log.Helper
}

func NewGameSortRepo(data *Data, logger log.Logger) biz.GameSortRepo {
	return &GameSortRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (g *GameSortRepo) GetGameSortList(ctx context.Context) ([]*biz.GameSort, error) {
	var gameSort []*GameSort
	if err := g.data.DB(ctx).Table("soccer_game_team_sort").Find(&gameSort).Error; err != nil {
		return nil, errors.NotFound("TEAMS_NOT_FOUND", "比赛排名不存在")
	}

	res := make([]*biz.GameSort, 0)
	for _, item := range gameSort {
		res = append(res, &biz.GameSort{
			ID:       item.ID,
			SortName: item.SortName,
			SortType: item.Type,
			EndTime:  item.EndTime,
		})
	}

	return res, nil
}
