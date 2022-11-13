package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "goal/api/game/service/v1"
	"time"
)

type Game struct {
	ID            int64
	Name          string
	RedTeamId     int64
	BlueTeamId    int64
	StartTime     time.Time
	UpEndTime     time.Time
	DownStartTime time.Time
	EndTime       time.Time
}

type GameRepo interface {
	GetGameById(ctx context.Context, gameId int64) (*Game, error)
	GetGameByIds(ctx context.Context, ids ...int64) (map[int64]*Game, error)
	GetGameList(ctx context.Context) ([]*Game, error)
}

type GameUseCase struct {
	gameRepo GameRepo
	teamRepo TeamRepo
	log      *log.Helper
}

func NewGameUseCase(repo GameRepo, teamRepo TeamRepo, logger log.Logger) *GameUseCase {
	return &GameUseCase{gameRepo: repo, teamRepo: teamRepo, log: log.NewHelper(logger)}
}

func (g *GameUseCase) GetGameList(ctx context.Context) (*v1.GetGameListReply, error) {
	var (
		game   []*Game
		teamId []int64
		teams  map[int64]*Team
		err    error
	)

	game, err = g.gameRepo.GetGameList(ctx)
	if err != nil {
		return nil, err
	}
	for _, item := range game {
		teamId = append(teamId, item.RedTeamId)
		teamId = append(teamId, item.BlueTeamId)
	}

	teams, err = g.teamRepo.GetTeamByIds(ctx, teamId...)
	if err != nil {
		return nil, err
	}

	res := &v1.GetGameListReply{
		Items: make([]*v1.GetGameListReply_Item, 0),
	}

	for _, item := range game {
		teamTmp := make([]*v1.GetGameListReply_Item_Team, 0) // 队伍信息
		if value, ok := teams[item.RedTeamId]; ok {
			teamTmp = append(teamTmp, &v1.GetGameListReply_Item_Team{
				TeamId:   value.ID,
				TeamName: value.Name,
				TeamType: "red",
			})
		}
		if value, ok := teams[item.BlueTeamId]; ok {
			teamTmp = append(teamTmp, &v1.GetGameListReply_Item_Team{
				TeamId:   value.ID,
				TeamName: value.Name,
				TeamType: "blue",
			})
		}

		res.Items = append(res.Items, &v1.GetGameListReply_Item{
			GameId:        item.ID,
			Name:          item.Name,
			StartTime:     item.StartTime.Format("2006-01-02 15:04:05"),
			EndTime:       item.EndTime.Format("2006-01-02 15:04:05"),
			UpEndTime:     item.UpEndTime.Format("2006-01-02 15:04:05"),
			DownStartTime: item.DownStartTime.Format("2006-01-02 15:04:05"),
			Teams:         teamTmp,
		})
	}

	return res, nil
}
