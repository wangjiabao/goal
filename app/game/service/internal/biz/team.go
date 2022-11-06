package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Team struct {
	ID   int64
	Name string
}

type TeamRepo interface {
	GetTeamByIds(ctx context.Context, ids ...int64) (map[int64]*Team, error)
}

type TeamUseCase struct {
	teamRepo TeamRepo
	log      *log.Helper
}

func NewTeamUseCase(repo TeamRepo, logger log.Logger) *TeamUseCase {
	return &TeamUseCase{teamRepo: repo, log: log.NewHelper(logger)}
}
