package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "goal/api/admin/service/v1"
	"time"
)

type Sort struct {
	ID       int64
	SortName string
	Type     string
	EndTime  time.Time
}

type SortRepo interface {
	GetGameSortById(ctx context.Context, gameId int64) (*Sort, error)
	GetGameSortList(ctx context.Context) ([]*Sort, error)
}

type SortUseCase struct {
	sortRepo SortRepo
	log      *log.Helper
}

func NewSortUseCase(repo SortRepo, logger log.Logger) *SortUseCase {
	return &SortUseCase{sortRepo: repo, log: log.NewHelper(logger)}
}

func (s *SortUseCase) GetGameSortList(ctx context.Context) (*v1.GetGameSortListReply, error) {
	var (
		sort []*Sort
		err  error
	)

	sort, err = s.sortRepo.GetGameSortList(ctx)
	if err != nil {
		return nil, err
	}

	res := &v1.GetGameSortListReply{
		Sorts: make([]*v1.GetGameSortListReply_Sort, 0),
	}

	for _, item := range sort {
		res.Sorts = append(res.Sorts, &v1.GetGameSortListReply_Sort{
			SortId:   item.ID,
			SortType: item.Type,
			EndTime:  item.EndTime.Format("2006-01-02 15:04:05"),
			SortName: item.SortName,
		})
	}

	return res, nil
}
