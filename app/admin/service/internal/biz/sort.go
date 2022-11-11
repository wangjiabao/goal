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
	Content  string
	Status   string
	EndTime  time.Time
}

type SortRepo interface {
	GetGameSortById(ctx context.Context, gameId int64) (*Sort, error)
	GetGameSortList(ctx context.Context) ([]*Sort, error)
	CreateSort(ctx context.Context, sc *Sort) (*Sort, error)
	UpdateSort(ctx context.Context, sc *Sort) (*Sort, error)
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

func (s *SortUseCase) CreateSort(ctx context.Context, req *v1.CreateSortRequest) (*v1.CreateSortReply, error) {
	var (
		err     error
		sort    *Sort
		endTime time.Time
	)

	endTime, err = time.Parse("2006-01-02 15:04:05", req.SendBody.EndTime) // 时间进行格式校验
	if nil != err {
		return nil, err
	}

	if sort, err = s.sortRepo.CreateSort(ctx, &Sort{
		SortName: req.SendBody.SortName, // todo 参数校验
		Type:     req.SendBody.Type,
		EndTime:  endTime,
	}); nil != err {
		return nil, err
	}

	return &v1.CreateSortReply{SortId: sort.ID}, nil
}

func (s *SortUseCase) UpdateSort(ctx context.Context, req *v1.UpdateSortRequest) (*v1.UpdateSortReply, error) {
	var (
		err     error
		endTime time.Time
		sort    *Sort
	)

	endTime, err = time.Parse("2006-01-02 15:04:05", req.SendBody.EndTime) // 时间进行格式校验
	if nil != err {
		return nil, err
	}

	if sort, err = s.sortRepo.UpdateSort(ctx, &Sort{
		Content: req.SendBody.Content, // todo 参数校验
		ID:      req.SendBody.SortId,
		Status:  req.SendBody.Status,
		EndTime: endTime,
	}); nil != err {
		return nil, err
	}

	return &v1.UpdateSortReply{SortId: sort.ID}, nil
}
