package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"goal/app/play/service/internal/biz"
	"gorm.io/gorm"
	"time"
)

type Game struct {
	ID         int64     `gorm:"primarykey;type:int"`
	Name       string    `gorm:"type:varchar(45);not null"`
	RedTeamId  int64     `gorm:"type:int;not null"`
	BlueTeamId int64     `gorm:"type:int;not null"`
	StartTime  time.Time `gorm:"type:datetime;not null"`
	EndTime    time.Time `gorm:"type:datetime;not null"`
	CreatedAt  time.Time `gorm:"type:datetime;not null"`
	UpdatedAt  time.Time `gorm:"type:datetime;not null"`
}

type GameRepo struct {
	data *Data
	log  *log.Helper
}

func NewGameRepo(data *Data, logger log.Logger) biz.GameRepo {
	return &GameRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (d *GameRepo) GetGameById(ctx context.Context, gameId int64) (*biz.Game, error) {

	var game Game
	if err := d.data.db.Where(&Game{ID: gameId}).Table("soccer_game").First(&game).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("GAME_NOT_FOUND", "game not found")
		}

		return nil, errors.New(500, "GAME_NOT_FOUND", err.Error())
	}

	return &biz.Game{
		ID:         game.ID,
		RedTeamId:  game.RedTeamId,
		BlueTeamId: game.BlueTeamId,
		Name:       game.Name,
		StartTime:  game.StartTime,
		EndTime:    game.EndTime,
	}, nil
}
