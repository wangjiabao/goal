package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "goal/api/game/service/v1"
)

type Team struct {
	ID   int64
	Name string
}

type TeamRepo interface {
	GetTeamByIds(ctx context.Context, ids ...int64) (map[int64]*Team, error)
	GetTeamList(ctx context.Context) ([]*Team, error)
}

type TeamUseCase struct {
	teamRepo TeamRepo
	log      *log.Helper
}

func NewTeamUseCase(repo TeamRepo, logger log.Logger) *TeamUseCase {
	return &TeamUseCase{teamRepo: repo, log: log.NewHelper(logger)}
}

func (t *TeamUseCase) GetTeamList(ctx context.Context) (*v1.GetTeamListReply, error) {
	var (
		team []*Team
		err  error
	)
	team, err = t.teamRepo.GetTeamList(ctx)
	if err != nil {
		return nil, err
	}

	res := &v1.GetTeamListReply{
		Teams: make([]*v1.GetTeamListReply_Team, 0),
	}

	for _, item := range team {
		res.Teams = append(res.Teams, &v1.GetTeamListReply_Team{
			TeamId:   item.ID,
			TeamName: item.Name,
		})
	}

	return res, nil
}
