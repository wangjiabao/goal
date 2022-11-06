package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "goal/api/game/service/v1"
	"time"
)

type GameSort struct {
	ID       int64
	SortName string
	SortType string
	EndTime  time.Time
}

type GameSortRepo interface {
	GetGameSortList(ctx context.Context) ([]*GameSort, error)
}

type GameSortUseCase struct {
	gameSortRepo GameSortRepo
	log          *log.Helper
}

func NewGameSortUseCase(gameSortRepo GameSortRepo, logger log.Logger) *GameSortUseCase {
	return &GameSortUseCase{gameSortRepo: gameSortRepo, log: log.NewHelper(logger)}
}

func (g *GameSortUseCase) GetGameSortList(ctx context.Context) (*v1.GetGameSortListReply, error) {
	var (
		gameSort []*GameSort
		err      error
	)
	gameSort, err = g.gameSortRepo.GetGameSortList(ctx)
	if err != nil {
		return nil, err
	}

	res := &v1.GetGameSortListReply{
		Sorts: make([]*v1.GetGameSortListReply_Sort, 0),
	}

	for _, item := range gameSort {
		res.Sorts = append(res.Sorts, &v1.GetGameSortListReply_Sort{
			SortId:   item.ID,
			SortType: item.SortType,
			EndTime:  item.EndTime.Format("2006-01-02 15:04:05"),
			SortName: item.SortName,
		})
	}

	return res, nil
}
