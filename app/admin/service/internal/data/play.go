package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"goal/app/admin/service/internal/biz"
	"time"
)

type Play struct {
	ID             int64     `gorm:"primarykey;type:int"`
	CreateUserId   int64     `gorm:"type:int;not null"`
	CreateUserType string    `gorm:"type:varchar(45);not null"`
	Type           string    `gorm:"type:varchar(45);not null"`
	StartTime      time.Time `gorm:"type:datetime;not null"`
	EndTime        time.Time `gorm:"type:datetime;not null"`
	CreatedAt      time.Time `gorm:"type:datetime;not null"`
	UpdatedAt      time.Time `gorm:"type:datetime;not null"`
}

type PlayGameRel struct {
	ID        int64     `gorm:"primarykey;type:int"`
	PlayId    int64     `gorm:"type:int;not null"`
	GameId    int64     `gorm:"type:int;not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type PlayRoomRel struct {
	ID        int64     `gorm:"primarykey;type:int"`
	PlayId    int64     `gorm:"type:int;not null"`
	RoomId    int64     `gorm:"type:int;not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type PlaySortRel struct {
	ID        int64     `gorm:"primarykey;type:int"`
	PlayId    int64     `gorm:"type:int;not null"`
	SortId    int64     `gorm:"type:int;not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type PlayGameScoreUserRel struct {
	ID        int64     `gorm:"primarykey;type:int"`
	PlayId    int64     `gorm:"type:int;not null"`
	UserId    int64     `gorm:"type:int;not null"`
	Pay       int64     `gorm:"type:int;not null"`
	Content   string    `gorm:"type:varchar(45);not null"`
	Status    string    `gorm:"type:varchar(45);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type PlayGameTeamResultUserRel struct {
	ID        int64     `gorm:"primarykey;type:int"`
	PlayId    int64     `gorm:"type:int;not null"`
	UserId    int64     `gorm:"type:int;not null"`
	Pay       int64     `gorm:"type:int;not null"`
	Content   string    `gorm:"type:varchar(45);not null"`
	Status    string    `gorm:"type:varchar(45);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type PlayGameTeamSortUserRel struct {
	ID        int64     `gorm:"primarykey;type:int"`
	PlayId    int64     `gorm:"type:int;not null"`
	UserId    int64     `gorm:"type:int;not null"`
	Pay       int64     `gorm:"type:int;not null"`
	Content   string    `gorm:"type:varchar(500);not null"`
	Status    string    `gorm:"type:varchar(45);not null"`
	SortId    int64     `gorm:"type:int;not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type PlayGameTeamGoalUserRel struct {
	ID        int64     `gorm:"primarykey;type:int"`
	PlayId    int64     `gorm:"type:int;not null"`
	UserId    int64     `gorm:"type:int;not null"`
	TeamId    int64     `gorm:"type:int;not null"`
	Goal      int64     `gorm:"type:int;not null"`
	Type      string    `gorm:"type:varchar(45);not null"`
	Pay       int64     `gorm:"type:int;not null"`
	Status    string    `gorm:"type:varchar(45);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type PlayRepo struct {
	data *Data
	log  *log.Helper
}

type PlayGameRelRepo struct {
	data *Data
	log  *log.Helper
}

type PlayRoomRelRepo struct {
	data *Data
	log  *log.Helper
}

type PlaySortRelRepo struct {
	data *Data
	log  *log.Helper
}

type PlayGameScoreUserRelRepo struct {
	data *Data
	log  *log.Helper
}

type PlayGameTeamResultUserRelRepo struct {
	data *Data
	log  *log.Helper
}

type PlayGameTeamGoalUserRelRepo struct {
	data *Data
	log  *log.Helper
}

type PlayGameTeamSortUserRelRepo struct {
	data *Data
	log  *log.Helper
}

func NewPlayRepo(data *Data, logger log.Logger) biz.PlayRepo {
	return &PlayRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func NewPlayGameScoreUserRelRepo(data *Data, logger log.Logger) biz.PlayGameScoreUserRelRepo {
	return &PlayGameScoreUserRelRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func NewPlayGameTeamResultUserRelRepo(data *Data, logger log.Logger) biz.PlayGameTeamResultUserRelRepo {
	return &PlayGameTeamResultUserRelRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func NewPlayGameTeamGoalUserRelRepo(data *Data, logger log.Logger) biz.PlayGameTeamGoalUserRelRepo {
	return &PlayGameTeamGoalUserRelRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func NewPlayGameTeamSortUserRelRepo(data *Data, logger log.Logger) biz.PlayGameTeamSortUserRelRepo {
	return &PlayGameTeamSortUserRelRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func NewPlayGameRelRepo(data *Data, logger log.Logger) biz.PlayGameRelRepo {
	return &PlayGameRelRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func NewPlayRoomRelRepo(data *Data, logger log.Logger) biz.PlayRoomRelRepo {
	return &PlayRoomRelRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func NewPlaySortRelRepo(data *Data, logger log.Logger) biz.PlaySortRelRepo {
	return &PlaySortRelRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (pgr *PlayGameRelRepo) GetPlayGameRelByGameId(ctx context.Context, gameId int64) ([]*biz.PlayGameRel, error) {
	var l []*PlayGameRel
	if result := pgr.data.DB(ctx).Table("goal_play_game_rel").Where(&PlayGameRel{GameId: gameId}).Find(&l); result.Error != nil {
		return nil, errors.InternalServer("SELECT_PLAY_GAME_REL_ERROR", "查询比赛玩法关系失败")
	}

	pl := make([]*biz.PlayGameRel, 0)
	for _, v := range l {
		pl = append(pl, &biz.PlayGameRel{
			ID:     v.ID,
			PlayId: v.PlayId,
			GameId: v.GameId,
		})
	}
	return pl, nil
}

func (prr *PlayRoomRelRepo) GetPlayRoomRelByRoomId(ctx context.Context, roomId int64) ([]*biz.PlayRoomRel, error) {
	var l []*PlayRoomRel
	if result := prr.data.DB(ctx).Table("goal_play_room_rel").Where(&PlayRoomRel{RoomId: roomId}).Find(&l); result.Error != nil {
		return nil, errors.InternalServer("SELECT_PLAY_GAME_REL_ERROR", "查询房间内比赛玩法关系失败")
	}

	pl := make([]*biz.PlayRoomRel, 0)
	for _, v := range l {
		pl = append(pl, &biz.PlayRoomRel{
			ID:     v.ID,
			PlayId: v.PlayId,
			RoomId: v.RoomId,
		})
	}
	return pl, nil
}

func (psr *PlaySortRelRepo) GetPlaySortRelBySortId(ctx context.Context, sortId int64) ([]*biz.PlaySortRel, error) {
	var l []*PlaySortRel
	if result := psr.data.DB(ctx).Table("goal_play_game_sort_rel").Where("sort_id=?", sortId).Find(&l); result.Error != nil {
		return nil, errors.InternalServer("SELECT_PLAY_GAME_SORT_REL_ERROR", "查询比赛排名玩法关系失败")
	}

	pl := make([]*biz.PlaySortRel, 0)
	for _, v := range l {
		pl = append(pl, &biz.PlaySortRel{
			ID:     v.ID,
			PlayId: v.PlayId,
			SortId: v.SortId,
		})
	}
	return pl, nil
}

func (p *PlayRepo) GetPlayListByIds(ctx context.Context, ids ...int64) ([]*biz.Play, error) {
	var l []*Play
	if err := p.data.DB(ctx).Table("goal_play").Where("id IN (?)", ids).Find(&l).Error; err != nil {
		return nil, errors.NotFound("TEAMS_NOT_FOUND", "查询玩法列表失败")
	}

	pl := make([]*biz.Play, 0)
	for _, v := range l {
		pl = append(pl, &biz.Play{
			ID:        v.ID,
			Type:      v.Type,
			StartTime: v.StartTime,
			EndTime:   v.EndTime,
		})
	}

	return pl, nil
}

func (p *PlayRepo) GetPlayById(ctx context.Context, id int64) (*biz.Play, error) {
	var play *Play
	if err := p.data.DB(ctx).Table("goal_play").Where("id=?", id).First(&play).Error; err != nil {
		return nil, errors.NotFound("TEAMS_NOT_FOUND", "查询玩法失败")
	}

	return &biz.Play{
		ID:             play.ID,
		CreateUserId:   play.CreateUserId,
		CreateUserType: play.CreateUserType,
		Type:           play.Type,
		StartTime:      play.StartTime,
		EndTime:        play.EndTime,
	}, nil
}

func (psr *PlayGameScoreUserRelRepo) GetPlayGameScoreUserRelByPlayIds(ctx context.Context, playIds ...int64) (map[int64][]*biz.PlayGameScoreUserRel, error) {
	var l []*PlayGameScoreUserRel
	if result := psr.data.DB(ctx).Table("play_game_score_user_rel").Where("play_id IN (?)", playIds).Find(&l); result.Error != nil {
		return nil, errors.InternalServer("SELECT_PLAY_GAME_SCORE_USER_ERROR", "查询比分玩法用户关系失败")
	}

	pl := make(map[int64][]*biz.PlayGameScoreUserRel, 0)
	for _, v := range l {
		pl[v.PlayId] = append(pl[v.PlayId], &biz.PlayGameScoreUserRel{
			ID:      v.ID,
			PlayId:  v.PlayId,
			UserId:  v.UserId,
			Content: v.Content,
			Pay:     v.Pay,
			Status:  v.Status,
		})
	}
	return pl, nil
}

func (psr *PlayGameScoreUserRelRepo) GetPlayGameScoreUserRelByPlayId(ctx context.Context, playId int64) ([]*biz.PlayGameScoreUserRel, error) {
	var l []*PlayGameScoreUserRel
	if result := psr.data.DB(ctx).Table("play_game_score_user_rel").Where("play_id=?", playId).Find(&l); result.Error != nil {
		return nil, errors.InternalServer("SELECT_PLAY_SCORE_RESULT_USER_ERROR", "查询比赛得分玩法用户关系失败")
	}

	pl := make([]*biz.PlayGameScoreUserRel, 0)
	for _, v := range l {
		pl = append(pl, &biz.PlayGameScoreUserRel{
			ID:      v.ID,
			PlayId:  v.PlayId,
			UserId:  v.UserId,
			Content: v.Content,
			Pay:     v.Pay,
			Status:  v.Status,
		})
	}
	return pl, nil
}

// SetRewarded 在事务中使用
func (psr *PlayGameScoreUserRelRepo) SetRewarded(ctx context.Context, id int64) error {
	var err error
	if err = psr.data.DB(ctx).Table("play_game_score_user_rel").
		Where("id=?", id).
		Update("status", "rewarded").Error; nil != err {
		return errors.NotFound("play game score rel error", "play game score rel found")
	}

	return nil
}

func (pgtR *PlayGameTeamResultUserRelRepo) GetPlayGameTeamResultUserRelByPlayIds(ctx context.Context, playIds ...int64) (map[int64][]*biz.PlayGameTeamResultUserRel, error) {
	var l []*PlayGameTeamResultUserRel
	if result := pgtR.data.DB(ctx).Table("play_game_team_result_user_rel").Where("play_id IN (?)", playIds).Find(&l); result.Error != nil {
		return nil, errors.InternalServer("SELECT_PLAY_GAME_TEAM_RESULT_USER_ERROR", "查询比赛结果玩法用户关系失败")
	}

	pl := make(map[int64][]*biz.PlayGameTeamResultUserRel, 0)
	for _, v := range l {
		pl[v.PlayId] = append(pl[v.PlayId], &biz.PlayGameTeamResultUserRel{
			ID:      v.ID,
			PlayId:  v.PlayId,
			UserId:  v.UserId,
			Content: v.Content,
			Pay:     v.Pay,
			Status:  v.Status,
		})
	}
	return pl, nil
}

func (pgtR *PlayGameTeamResultUserRelRepo) GetPlayGameTeamResultUserRelByPlayId(ctx context.Context, playId int64) ([]*biz.PlayGameTeamResultUserRel, error) {
	var l []*PlayGameTeamResultUserRel
	if result := pgtR.data.DB(ctx).Table("play_game_team_result_user_rel").Where("play_id=?", playId).Find(&l); result.Error != nil {
		return nil, errors.InternalServer("SELECT_PLAY_GAME_TEAM_RESULT_USER_ERROR", "查询比赛结果玩法用户关系失败")
	}

	pl := make([]*biz.PlayGameTeamResultUserRel, 0)
	for _, v := range l {
		pl = append(pl, &biz.PlayGameTeamResultUserRel{
			ID:      v.ID,
			PlayId:  v.PlayId,
			UserId:  v.UserId,
			Content: v.Content,
			Pay:     v.Pay,
			Status:  v.Status,
		})
	}
	return pl, nil
}

// SetRewarded 在事务中使用
func (pgtR *PlayGameTeamResultUserRelRepo) SetRewarded(ctx context.Context, id int64) error {
	var err error
	if err = pgtR.data.DB(ctx).Table("play_game_team_result_user_rel").
		Where("id=?", id).
		Update("status", "rewarded").Error; nil != err {
		return errors.NotFound("play game team result rel error", "play game team result rel not found")
	}

	return nil
}

func (pgtG *PlayGameTeamGoalUserRelRepo) GetPlayGameTeamGoalUserRelByPlayIds(ctx context.Context, playIds ...int64) (map[int64][]*biz.PlayGameTeamGoalUserRel, error) {
	var l []*PlayGameTeamGoalUserRel
	if result := pgtG.data.DB(ctx).Table("play_game_team_goal_user_rel").Where("play_id IN (?)", playIds).Find(&l); result.Error != nil {
		return nil, errors.InternalServer("SELECT_PLAY_GAME_TEAM_GOAL_USER_ERROR", "查询比赛进球数玩法用户关系失败")
	}

	pl := make(map[int64][]*biz.PlayGameTeamGoalUserRel, 0)
	for _, v := range l {
		pl[v.PlayId] = append(pl[v.PlayId], &biz.PlayGameTeamGoalUserRel{
			ID:     v.ID,
			PlayId: v.PlayId,
			UserId: v.UserId,
			TeamId: v.TeamId,
			Goal:   v.Goal,
			Type:   v.Type,
			Pay:    v.Pay,
			Status: v.Status,
		})
	}
	return pl, nil
}

func (pgtG *PlayGameTeamGoalUserRelRepo) GetPlayGameTeamGoalUserRelByPlayId(ctx context.Context, playId int64) ([]*biz.PlayGameTeamGoalUserRel, error) {
	var l []*PlayGameTeamGoalUserRel
	if result := pgtG.data.DB(ctx).Table("play_game_team_goal_user_rel").Where("play_id=?", playId).Find(&l); result.Error != nil {
		return nil, errors.InternalServer("SELECT_PLAY_GAME_TEAM_GOAL_USER_ERROR", "查询比赛进球数玩法用户关系失败")
	}

	pl := make([]*biz.PlayGameTeamGoalUserRel, 0)
	for _, v := range l {
		pl = append(pl, &biz.PlayGameTeamGoalUserRel{
			ID:     v.ID,
			PlayId: v.PlayId,
			UserId: v.UserId,
			TeamId: v.TeamId,
			Goal:   v.Goal,
			Type:   v.Type,
			Pay:    v.Pay,
			Status: v.Status,
		})
	}
	return pl, nil
}

// SetRewarded 在事务中使用
func (pgtG *PlayGameTeamGoalUserRelRepo) SetRewarded(ctx context.Context, id int64) error {
	var err error
	if err = pgtG.data.DB(ctx).Table("play_game_team_goal_user_rel").
		Where("id=?", id).
		Update("status", "rewarded").Error; nil != err {
		return errors.NotFound("play game team result rel error", "play game team goal rel not found")
	}

	return nil
}
func (pgtS *PlayGameTeamSortUserRelRepo) GetPlayGameTeamSortUserRelByPlayIds(ctx context.Context, playIds ...int64) (map[int64][]*biz.PlayGameTeamSortUserRel, error) {
	var l []*PlayGameTeamSortUserRel
	if result := pgtS.data.DB(ctx).Table("play_game_team_sort_user_rel").Where("play_id IN (?)", playIds).Find(&l); result.Error != nil {
		return nil, errors.InternalServer("SELECT_PLAY_GAME_SORT_RESULT_USER_ERROR", "查询比赛排名玩法用户关系失败")
	}

	pl := make(map[int64][]*biz.PlayGameTeamSortUserRel, 0)
	for _, v := range l {
		pl[v.PlayId] = append(pl[v.PlayId], &biz.PlayGameTeamSortUserRel{
			ID:      v.ID,
			PlayId:  v.PlayId,
			UserId:  v.UserId,
			SortId:  v.SortId,
			Content: v.Content,
			Pay:     v.Pay,
			Status:  v.Status,
		})
	}
	return pl, nil
}

// SetRewarded 在事务中使用
func (pgtS *PlayGameTeamSortUserRelRepo) SetRewarded(ctx context.Context, id int64) error {
	var err error
	if err = pgtS.data.DB(ctx).Table("play_game_team_sort_user_rel").
		Where("id=?", id).
		Update("status", "rewarded").Error; nil != err {
		return errors.NotFound("play game team sort rel error", "play game team sort rel not found")
	}

	return nil
}

// CreatePlay .
func (p *PlayRepo) CreatePlay(ctx context.Context, pc *biz.Play) (*biz.Play, error) {
	var play Play
	play.CreateUserId = pc.CreateUserId
	play.CreateUserType = pc.CreateUserType
	play.Type = pc.Type
	play.StartTime = pc.StartTime
	play.EndTime = pc.EndTime
	res := p.data.DB(ctx).Table("goal_play").Create(&play)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_PLAY_ERROR", "玩法创建失败")
	}

	return &biz.Play{
		ID:             play.ID,
		CreateUserId:   play.CreateUserId,
		CreateUserType: play.CreateUserType,
		Type:           play.Type,
		StartTime:      play.StartTime,
		EndTime:        play.EndTime,
	}, nil
}

// CreatePlayGameRel .
func (pgr *PlayGameRelRepo) CreatePlayGameRel(ctx context.Context, rel *biz.PlayGameRel) (*biz.PlayGameRel, error) {
	var playGameRel PlayGameRel
	playGameRel.GameId = rel.GameId
	playGameRel.PlayId = rel.PlayId
	res := pgr.data.DB(ctx).Table("goal_play_game_rel").Create(&playGameRel)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_PLAY_GAME_REL_ERROR", "玩法和比赛关系创建失败")
	}

	return &biz.PlayGameRel{
		ID:     playGameRel.ID,
		PlayId: playGameRel.PlayId,
		GameId: playGameRel.GameId,
	}, nil
}

// CreatePlaySortRel .
func (psr *PlaySortRelRepo) CreatePlaySortRel(ctx context.Context, rel *biz.PlaySortRel) (*biz.PlaySortRel, error) {
	var playSortRel PlaySortRel
	playSortRel.SortId = rel.SortId
	playSortRel.PlayId = rel.PlayId
	res := psr.data.DB(ctx).Table("goal_play_game_sort_rel").Create(&playSortRel)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_PLAY_SORT_REL_ERROR", "玩法和比赛排名关系创建失败")
	}

	return &biz.PlaySortRel{
		ID:     playSortRel.ID,
		PlayId: playSortRel.PlayId,
		SortId: playSortRel.SortId,
	}, nil
}
