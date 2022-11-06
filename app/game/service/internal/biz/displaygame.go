package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "goal/api/game/service/v1"
)

type DisplayGame struct {
	ID     int64
	GameId int64
	Type   string
}

type DisplayGameRepo interface {
	GetDisplayGameByType(ctx context.Context, displayGameType string) (*DisplayGame, error)
}

type DisplayGameUseCase struct {
	displayGameRepo DisplayGameRepo
	gameRepo        GameRepo
	teamRepo        TeamRepo
	log             *log.Helper
}

func NewDisplayGameUseCase(repo DisplayGameRepo, gameRepo GameRepo, teamRepo TeamRepo, logger log.Logger) *DisplayGameUseCase {
	return &DisplayGameUseCase{
		displayGameRepo: repo,
		gameRepo:        gameRepo,
		teamRepo:        teamRepo,
		log:             log.NewHelper(logger)}
}

func (dg *DisplayGameUseCase) GetDisplayGame(ctx context.Context, disPlayGameType string) (*v1.DisplayGameReply, error) {
	var (
		game        *Game
		disPlayGame *DisplayGame
		teams       map[int64]*Team
		err         error
	)

	disPlayGame, err = dg.displayGameRepo.GetDisplayGameByType(ctx, disPlayGameType)
	if err != nil {
		return nil, err
	}

	game, err = dg.gameRepo.GetGameById(ctx, disPlayGame.GameId)
	if err != nil {
		return nil, err
	}

	teams, err = dg.teamRepo.GetTeamByIds(ctx, game.RedTeamId, game.BlueTeamId)
	if err != nil {
		return nil, err
	}

	res := &v1.DisplayGameReply{
		GameId: disPlayGame.GameId,
		Name:   game.Name,
		Teams:  make([]*v1.DisplayGameReply_Team, 0),
	}

	if v, ok := teams[game.RedTeamId]; ok {
		res.Teams = append(res.Teams, &v1.DisplayGameReply_Team{
			TeamId:   v.ID,
			TeamName: v.Name,
			TeamType: "red",
		})
	}

	if v, ok := teams[game.BlueTeamId]; ok {
		res.Teams = append(res.Teams, &v1.DisplayGameReply_Team{
			TeamId:   v.ID,
			TeamName: v.Name,
			TeamType: "blue",
		})
	}

	return res, nil
}
