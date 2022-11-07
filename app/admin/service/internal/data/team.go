package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"goal/app/admin/service/internal/biz"
	"time"
)

type Team struct {
	ID        int64     `gorm:"primarykey;type:int"`
	Name      string    `gorm:"type:varchar(45);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type TeamRepo struct {
	data *Data
	log  *log.Helper
}

func NewTeamRepo(data *Data, logger log.Logger) biz.TeamRepo {
	return &TeamRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (t *TeamRepo) GetTeamByIds(ctx context.Context, ids ...int64) (map[int64]*biz.Team, error) {

	var ts []*Team
	if err := t.data.DB(ctx).Table("soccer_team").Where("id IN (?)", ids).Find(&ts).Error; err != nil {
		return nil, errors.NotFound("TEAMS_NOT_FOUND", "队伍不存在")
	}

	res := make(map[int64]*biz.Team)
	for _, item := range ts {
		res[item.ID] = &biz.Team{ID: item.ID, Name: item.Name}
	}

	return res, nil
}

func (t *TeamRepo) GetTeamList(ctx context.Context) ([]*biz.Team, error) {
	var team []*Team
	if err := t.data.DB(ctx).Table("soccer_team").Find(&team).Error; err != nil {
		return nil, errors.NotFound("TEAMS_NOT_FOUND", "队伍不存在")
	}

	res := make([]*biz.Team, 0)
	for _, item := range team {
		res = append(res, &biz.Team{
			ID:   item.ID,
			Name: item.Name,
		})
	}

	return res, nil
}

// CreateTeam .
func (t *TeamRepo) CreateTeam(ctx context.Context, ct *biz.Team) (*biz.Team, error) {
	var team Team
	team.Name = ct.Name
	res := t.data.DB(ctx).Table("soccer_team").Create(&team)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_TEAM_ERROR", "创建队伍比赛失败")
	}

	return &biz.Team{
		ID:   team.ID,
		Name: team.Name,
	}, nil
}
