package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type Sort struct {
	ID      int64
	Name    string
	Type    string
	EndTime time.Time
}

type SortRepo interface {
	GetGameSortById(ctx context.Context, gameId int64) (*Sort, error)
}

type SortUseCase struct {
	sortRepo SortRepo
	log      *log.Helper
}

func NewSortUseCase(repo SortRepo, logger log.Logger) *SortUseCase {
	return &SortUseCase{sortRepo: repo, log: log.NewHelper(logger)}
}
