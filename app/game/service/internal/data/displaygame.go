package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"goal/app/game/service/internal/biz"
	"gorm.io/gorm"
	"time"
)

type DisplayGame struct {
	ID        int64     `gorm:"primarykey;type:int"`
	GameId    int64     `gorm:"type:int;not null"`
	Type      string    `gorm:"type:varchar(45);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type DisplayGameRepo struct {
	data *Data
	log  *log.Helper
}

func NewDisplayGameRepo(data *Data, logger log.Logger) biz.DisplayGameRepo {
	return &DisplayGameRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (d *DisplayGameRepo) GetDisplayGameByType(ctx context.Context, displayGameType string) (*biz.DisplayGame, error) {
	var displayGame DisplayGame
	if err := d.data.db.Where(&DisplayGame{Type: displayGameType}).Table("display_game").First(&displayGame).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("DISPLAY_GAME_NOT_FOUND", "display game not found")
		}

		return nil, errors.New(500, "DISPLAY_GAME_NOT_FOUND", err.Error())
	}

	return &biz.DisplayGame{
		ID:     displayGame.ID,
		GameId: displayGame.GameId,
		Type:   displayGame.Type,
	}, nil
}
