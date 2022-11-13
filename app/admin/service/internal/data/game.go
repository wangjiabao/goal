package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"goal/app/admin/service/internal/biz"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Game struct {
	ID               int64     `gorm:"primarykey;type:int"`
	Name             string    `gorm:"type:varchar(45);not null"`
	RedTeamId        int64     `gorm:"type:int;not null"`
	BlueTeamId       int64     `gorm:"type:int;not null"`
	WinTeamId        int64     `gorm:"type:int;not null"`
	RedTeamUpGoal    int64     `gorm:"type:int;not null"`
	BlueTeamUpGoal   int64     `gorm:"type:int;not null"`
	RedTeamDownGoal  int64     `gorm:"type:int;not null"`
	BlueTeamDownGoal int64     `gorm:"type:int;not null"`
	UpEndTime        time.Time `gorm:"type:datetime;not null"`
	DownStartTime    time.Time `gorm:"type:datetime;not null"`
	Status           string    `gorm:"type:varchar(45);not null"`
	Result           string    `gorm:"type:varchar(45);not null"`
	StartTime        time.Time `gorm:"type:datetime;not null"`
	EndTime          time.Time `gorm:"type:datetime;not null"`
	CreatedAt        time.Time `gorm:"type:datetime;not null"`
	UpdatedAt        time.Time `gorm:"type:datetime;not null"`
}

type DisplayGame struct {
	ID        int64     `gorm:"primarykey;type:int"`
	GameId    int64     `gorm:"type:int;not null"`
	Type      string    `gorm:"type:varchar(45);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
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

type DisplayGameRepo struct {
	data *Data
	log  *log.Helper
}

func (g *GameRepo) GetGameById(ctx context.Context, gameId int64) (*biz.Game, error) {
	var game Game
	if err := g.data.db.Where(&Game{ID: gameId}).Table("soccer_game").First(&game).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("GAME_NOT_FOUND", "game not found")
		}

		return nil, errors.New(500, "GAME_NOT_FOUND", err.Error())
	}

	return &biz.Game{
		ID:               game.ID,
		RedTeamId:        game.RedTeamId,
		BlueTeamId:       game.BlueTeamId,
		Name:             game.Name,
		StartTime:        game.StartTime,
		EndTime:          game.EndTime,
		UpEndTime:        game.UpEndTime,
		DownStartTime:    game.DownStartTime,
		RedTeamDownGoal:  game.RedTeamDownGoal,
		RedTeamUpGoal:    game.RedTeamUpGoal,
		BlueTeamDownGoal: game.BlueTeamDownGoal,
		BlueTeamUpGoal:   game.BlueTeamUpGoal,
		WinTeamId:        game.WinTeamId,
		Status:           game.Status,
		Result:           game.Result,
	}, nil
}

// CreateGame .
func (g *GameRepo) CreateGame(ctx context.Context, gc *biz.Game) (*biz.Game, error) {
	var game Game
	game.StartTime = gc.StartTime
	game.EndTime = gc.EndTime
	game.Name = gc.Name
	game.RedTeamId = gc.RedTeamId
	game.BlueTeamId = gc.BlueTeamId
	game.DownStartTime = gc.DownStartTime
	game.UpEndTime = gc.UpEndTime
	res := g.data.DB(ctx).Table("soccer_game").Create(&game)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_PLAY_ERROR", "创建比赛失败")
	}

	return &biz.Game{
		ID: game.ID,
	}, nil
}

// UpdateGame .
func (g *GameRepo) UpdateGame(ctx context.Context, gc *biz.Game) (*biz.Game, error) {
	var game Game
	game.ID = gc.ID
	game.StartTime = gc.StartTime
	game.EndTime = gc.EndTime
	game.Name = gc.Name
	game.DownStartTime = gc.DownStartTime
	game.UpEndTime = gc.UpEndTime
	game.BlueTeamUpGoal = gc.BlueTeamUpGoal
	game.BlueTeamDownGoal = gc.BlueTeamDownGoal
	game.RedTeamUpGoal = gc.RedTeamUpGoal
	game.RedTeamDownGoal = gc.RedTeamDownGoal
	game.Status = gc.Status
	game.WinTeamId = gc.WinTeamId
	game.Result = strconv.FormatInt(gc.RedTeamUpGoal+gc.RedTeamDownGoal, 10) + ":" + strconv.FormatInt(gc.BlueTeamUpGoal+gc.BlueTeamDownGoal, 10)
	res := g.data.DB(ctx).Table("soccer_game").Updates(&game)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_PLAY_ERROR", "修改比赛失败")
	}

	return &biz.Game{
		ID: game.ID,
	}, nil
}

func (g *GameRepo) GetGameList(ctx context.Context) ([]*biz.Game, error) {
	var game []*Game
	if err := g.data.DB(ctx).Table("soccer_game").Order("created_at desc").Find(&game).Error; err != nil {
		return nil, errors.NotFound("TEAMS_NOT_FOUND", "比赛不存在")
	}

	res := make([]*biz.Game, 0)
	for _, item := range game {
		res = append(res, &biz.Game{
			ID:            item.ID,
			Name:          item.Name,
			RedTeamId:     item.RedTeamId,
			BlueTeamId:    item.BlueTeamId,
			StartTime:     item.StartTime,
			EndTime:       item.EndTime,
			UpEndTime:     item.UpEndTime,
			DownStartTime: item.DownStartTime,
		})
	}

	return res, nil
}

func (g *GameRepo) GetDisplayGame() (*biz.DisplayGame, error) {
	var displayGame DisplayGame
	if err := g.data.db.Table("display_game").First(&displayGame).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("DISPLAY_GAME_NOT_FOUND", "display game not found")
		}

		return nil, errors.New(500, "DISPLAY_GAME_NOT_FOUND", err.Error())
	}

	return &biz.DisplayGame{
		ID:     displayGame.ID,
		GameId: displayGame.GameId,
	}, nil
}

func (g *GameRepo) GetDisplayGameList() ([]*biz.DisplayGame, error) {
	var displayGame []*DisplayGame
	if err := g.data.db.Table("display_game").Find(&displayGame).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("DISPLAY_GAME_NOT_FOUND", "display game not found")
		}

		return nil, errors.New(500, "DISPLAY_GAME_NOT_FOUND", err.Error())
	}

	res := make([]*biz.DisplayGame, 0)
	for _, v := range displayGame {
		res = append(res, &biz.DisplayGame{
			ID:     v.ID,
			GameId: v.GameId,
		})
	}

	return res, nil
}

// CreateDisplayGame .
func (g *GameRepo) CreateDisplayGame(ctx context.Context, gameId int64) (*biz.DisplayGame, error) {
	var displayGame DisplayGame
	displayGame.GameId = gameId
	displayGame.Type = "index"
	res := g.data.DB(ctx).Table("display_game").Create(&displayGame)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_PLAY_ERROR", "创建首页展示比赛失败")
	}

	return &biz.DisplayGame{
		GameId: displayGame.GameId,
	}, nil
}

// DeleteDisplayGame .
func (g *GameRepo) DeleteDisplayGame(ctx context.Context, gameId int64) (bool, error) {
	var displayGame DisplayGame
	res := g.data.DB(ctx).Table("display_game").Where("game_id=?", gameId).Delete(&displayGame)
	if res.Error != nil {
		return false, errors.New(500, "CREATE_PLAY_ERROR", "删除首页展示比赛失败")
	}

	return true, nil
}

// UpdateDisplayGame .
func (g *GameRepo) UpdateDisplayGame(ctx context.Context, displayGame *biz.DisplayGame, gameId int64) (*biz.DisplayGame, error) {
	displayGame.GameId = gameId
	displayGame.Type = "index"
	res := g.data.DB(ctx).Table("display_game").Updates(&displayGame)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_PLAY_ERROR", "更新首页展示比赛失败")
	}

	return &biz.DisplayGame{
		GameId: displayGame.GameId,
	}, nil
}
