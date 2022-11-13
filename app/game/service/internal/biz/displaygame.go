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
	GetDisplayGameByType(ctx context.Context, displayGameType string) ([]*DisplayGame, error)
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
		games       map[int64]*Game
		disPlayGame []*DisplayGame
		gameIds     []int64
		teamIds     []int64
		teams       map[int64]*Team
		err         error
	)

	disPlayGame, err = dg.displayGameRepo.GetDisplayGameByType(ctx, "index")
	if err != nil {
		return nil, err
	}
	for _, v := range disPlayGame {
		gameIds = append(gameIds, v.GameId)
	}

	games, err = dg.gameRepo.GetGameByIds(ctx, gameIds...)
	if err != nil {
		return nil, err
	}
	for _, v := range games {
		teamIds = append(teamIds, v.RedTeamId, v.BlueTeamId)
	}

	teams, err = dg.teamRepo.GetTeamByIds(ctx, teamIds...)
	if err != nil {
		return nil, err
	}

	res := &v1.DisplayGameReply{
		Games: make([]*v1.DisplayGameReply_Game, 0),
	}

	for _, game := range games {

		tmpTeam := make([]*v1.DisplayGameReply_Game_Team, 0)
		if v, ok := teams[game.RedTeamId]; ok {
			tmpTeam = append(tmpTeam, &v1.DisplayGameReply_Game_Team{
				TeamId:   v.ID,
				TeamName: v.Name,
				TeamType: "red",
			})
		}

		if v, ok := teams[game.BlueTeamId]; ok {
			tmpTeam = append(tmpTeam, &v1.DisplayGameReply_Game_Team{
				TeamId:   v.ID,
				TeamName: v.Name,
				TeamType: "blue",
			})
		}

		res.Games = append(res.Games, &v1.DisplayGameReply_Game{
			GameId: game.ID,
			Name:   game.Name,
			Teams:  tmpTeam,
		})
	}

	return res, nil
}
