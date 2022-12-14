package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"goal/app/admin/service/internal/biz"
	"gorm.io/gorm"
	"time"
)

type Sort struct {
	ID        int64     `gorm:"primarykey;type:int"`
	SortName  string    `gorm:"type:varchar(45);not null"`
	Status    string    `gorm:"type:varchar(45);not null"`
	Content   string    `gorm:"type:varchar(500);not null"`
	Type      string    `gorm:"type:varchar(45);not null"`
	EndTime   time.Time `gorm:"type:datetime;not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type SortRepo struct {
	data *Data
	log  *log.Helper
}

func NewSortRepo(data *Data, logger log.Logger) biz.SortRepo {
	return &SortRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (s *SortRepo) GetGameSortById(ctx context.Context, sortId int64) (*biz.Sort, error) {
	var sort Sort
	if err := s.data.DB(ctx).Where("id=?", sortId).Table("soccer_game_team_sort").First(&sort).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("GAME_SORT_NOT_FOUND", "game sort not found")
		}

		return nil, errors.New(500, "GAME_SORT_NOT_FOUND", err.Error())
	}

	return &biz.Sort{
		ID:       sort.ID,
		EndTime:  sort.EndTime,
		SortName: sort.SortName,
		Type:     sort.Type,
		Status:   sort.Status,
		Content:  sort.Content,
	}, nil
}

func (s *SortRepo) GetNexGameSort(ctx context.Context, endTime time.Time) (*biz.Sort, error) {
	var sort Sort
	if err := s.data.DB(ctx).Where("end_time>?", endTime).Table("soccer_game_team_sort").Order("end_time asc").First(&sort).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("GAME_SORT_NOT_FOUND", "game sort not found")
		}

		return nil, errors.New(500, "GAME_SORT_NOT_FOUND", err.Error())
	}

	return &biz.Sort{
		ID:       sort.ID,
		EndTime:  sort.EndTime,
		SortName: sort.SortName,
		Type:     sort.Type,
		Status:   sort.Status,
		Content:  sort.Content,
	}, nil
}

func (s *SortRepo) GetGameSortList(ctx context.Context) ([]*biz.Sort, error) {
	var gameSort []*Sort
	if err := s.data.DB(ctx).Table("soccer_game_team_sort").Find(&gameSort).Error; err != nil {
		return nil, errors.NotFound("TEAMS_NOT_FOUND", "?????????????????????")
	}

	res := make([]*biz.Sort, 0)
	for _, item := range gameSort {
		res = append(res, &biz.Sort{
			ID:       item.ID,
			SortName: item.SortName,
			Content:  item.Content,
			Status:   item.Status,
			Type:     item.Type,
			EndTime:  item.EndTime,
		})
	}

	return res, nil
}

// CreateSort .
func (s *SortRepo) CreateSort(ctx context.Context, sc *biz.Sort) (*biz.Sort, error) {
	var sort Sort
	sort.EndTime = sc.EndTime
	sort.SortName = sc.SortName
	sort.Type = sc.Type
	res := s.data.DB(ctx).Table("soccer_game_team_sort").Create(&sort)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_SORT_ERROR", "??????????????????")
	}

	return &biz.Sort{
		ID: sort.ID,
	}, nil
}

// UpdateSort .
func (s *SortRepo) UpdateSort(ctx context.Context, sc *biz.Sort) (*biz.Sort, error) {
	var sort Sort
	sort.ID = sc.ID
	sort.Content = sc.Content
	sort.Status = sc.Status
	sort.EndTime = sc.EndTime
	res := s.data.DB(ctx).Table("soccer_game_team_sort").Updates(&sort)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_SORT_ERROR", "??????????????????")
	}

	return &biz.Sort{
		ID: sort.ID,
	}, nil
}
