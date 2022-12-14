package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"goal/app/play/service/internal/biz"
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

type LastTermPool struct {
	ID             int64     `gorm:"primarykey;type:int"`
	OriginGameId   int64     `gorm:"type:int;not null"`
	OriginPlayId   int64     `gorm:"type:int;not null"`
	GameId         int64     `gorm:"type:int;not null"`
	PlayId         int64     `gorm:"type:int;not null"`
	Total          int64     `gorm:"type:bigint;not null"`
	PlayType       string    `gorm:"type:varchar(45);not null"`
	OriginPlayType string    `gorm:"type:varchar(45);not null"`
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

type PlaySortRel struct {
	ID        int64     `gorm:"primarykey;type:int"`
	PlayId    int64     `gorm:"type:int;not null"`
	SortId    int64     `gorm:"type:int;not null"`
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

type PlayGameScoreUserRel struct {
	ID        int64     `gorm:"primarykey;type:int"`
	PlayId    int64     `gorm:"type:int;not null"`
	UserId    int64     `gorm:"type:int;not null"`
	Content   string    `gorm:"type:varchar(45);not null"`
	Pay       int64     `gorm:"type:bigint;not null"`
	OriginPay int64     `gorm:"type:bigint;not null"`
	Status    string    `gorm:"type:varchar(45);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type PlayGameTeamSortUserRel struct {
	ID        int64     `gorm:"primarykey;type:int"`
	PlayId    int64     `gorm:"type:int;not null"`
	UserId    int64     `gorm:"type:int;not null"`
	SortId    int64     `gorm:"type:int;not null"`
	OriginPay int64     `gorm:"type:bigint;not null"`
	Content   string    `gorm:"type:varchar(45);not null"`
	Pay       int64     `gorm:"type:bigint;not null"`
	Status    string    `gorm:"type:varchar(45);not null"`
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
	Pay       int64     `gorm:"type:bigint;not null"`
	OriginPay int64     `gorm:"type:bigint;not null"`
	Status    string    `gorm:"type:varchar(45);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type PlayGameTeamGoalUserRelTotal struct {
	Total int64
}

type PlayGameScoreUserRelTotal struct {
	Total int64
}

type PlayGameTeamResultUserRelTotal struct {
	Total int64
}

type PlayGameTeamSortUserRelTotal struct {
	Total int64
}

type PlayGameTeamResultUserRel struct {
	ID        int64     `gorm:"primarykey;type:int"`
	PlayId    int64     `gorm:"type:int;not null"`
	UserId    int64     `gorm:"type:int;not null"`
	Content   string    `gorm:"type:varchar(45);not null"`
	Pay       int64     `gorm:"type:bigint;not null"`
	OriginPay int64     `gorm:"type:bigint;not null"`
	Status    string    `gorm:"type:varchar(45);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type SystemConfig struct {
	ID        int64     `gorm:"primarykey;type:int"`
	Name      string    `gorm:"type:varchar(100);not null"`
	Value     int64     `gorm:"type:int;not null"`
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

type PlaySortRelRepo struct {
	data *Data
	log  *log.Helper
}

type PlayRoomRelRepo struct {
	data *Data
	log  *log.Helper
}

type PlayGameTeamGoalUserRelRepo struct {
	data *Data
	log  *log.Helper
}

type PlayGameScoreUserRelRepo struct {
	data *Data
	log  *log.Helper
}

type SystemConfigRepo struct {
	data *Data
	log  *log.Helper
}

type PlayGameTeamSortUserRelRepo struct {
	data *Data
	log  *log.Helper
}

type PlayGameTeamResultUserRelRepo struct {
	data *Data
	log  *log.Helper
}

func NewPlayRepo(data *Data, logger log.Logger) biz.PlayRepo {
	return &PlayRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func NewSystemConfigRepo(data *Data, logger log.Logger) biz.SystemConfigRepo {
	return &SystemConfigRepo{
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

func NewPlaySortRelRepo(data *Data, logger log.Logger) biz.PlaySortRelRepo {
	return &PlaySortRelRepo{
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

func NewPlayGameTeamResultUserRepo(data *Data, logger log.Logger) biz.PlayGameTeamResultUserRelRepo {
	return &PlayGameTeamResultUserRelRepo{
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

func NewPlayGameScoreUserRelRepo(data *Data, logger log.Logger) biz.PlayGameScoreUserRelRepo {
	return &PlayGameScoreUserRelRepo{
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

func (p *PlayRepo) GetAdminCreatePlayListByType(ctx context.Context, playType string) ([]*biz.Play, error) {
	var l []*Play
	if result := p.data.DB(ctx).Table("goal_play").Where(&Play{CreateUserId: 1, CreateUserType: "admin", Type: playType}).Find(&l); result.Error != nil {
		return nil, errors.InternalServer("SELECT_PLAY_ERROR", "????????????????????????")
	}

	pl := make([]*biz.Play, 0)
	for _, v := range l {
		pl = append(pl, &biz.Play{
			ID:        v.ID,
			StartTime: v.StartTime,
			EndTime:   v.EndTime,
		})
	}
	return pl, nil
}

func (p *PlayRepo) GetAdminCreatePlayByType(ctx context.Context, playType string) (*biz.Play, error) {
	var play *Play
	if result := p.data.DB(ctx).Table("goal_play").Where(&Play{CreateUserId: 1, CreateUserType: "admin", Type: playType}).First(&play); result.Error != nil {
		return nil, errors.InternalServer("SELECT_PLAY_ERROR", "??????????????????")
	}

	return &biz.Play{
		ID:        play.ID,
		StartTime: play.StartTime,
		EndTime:   play.EndTime,
	}, nil
}

func (p *PlayRepo) GetAdminCreatePlayList(ctx context.Context) ([]*biz.Play, error) {
	var l []*Play
	if result := p.data.DB(ctx).Table("goal_play").Where(&Play{CreateUserId: 1, CreateUserType: "admin"}).Find(&l); result.Error != nil {
		return nil, errors.InternalServer("SELECT_PLAY_ERROR", "????????????????????????")
	}

	pl := make([]*biz.Play, 0)
	for _, v := range l {
		pl = append(pl, &biz.Play{
			ID:   v.ID,
			Type: v.Type,
		})
	}
	return pl, nil
}

func (p *PlayRepo) GetAdminCreatePlayListByIds(ctx context.Context, ids ...int64) ([]*biz.Play, error) {
	var l []*Play
	if err := p.data.DB(ctx).Table("goal_play").
		Where(&Play{CreateUserId: 1, CreateUserType: "admin"}).
		Where("id IN (?)", ids).Find(&l).Error; err != nil {
		return nil, errors.NotFound("TEAMS_NOT_FOUND", "????????????????????????")
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

func (p *PlayRepo) GetPlayListByIds(ctx context.Context, ids ...int64) ([]*biz.Play, error) {
	var l []*Play
	if err := p.data.DB(ctx).Table("goal_play").Where("id IN (?)", ids).Find(&l).Error; err != nil {
		return nil, errors.NotFound("TEAMS_NOT_FOUND", "????????????????????????")
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

func (p *PlayRepo) GetPlayMapByIds(ctx context.Context, ids ...int64) (map[int64]*biz.Play, error) {
	var l []*Play
	if err := p.data.DB(ctx).Table("goal_play").Where("id IN (?)", ids).Find(&l).Error; err != nil {
		return nil, errors.NotFound("TEAMS_NOT_FOUND", "????????????????????????")
	}

	pl := make(map[int64]*biz.Play, 0)
	for _, v := range l {
		pl[v.ID] = &biz.Play{
			ID:        v.ID,
			Type:      v.Type,
			StartTime: v.StartTime,
			EndTime:   v.EndTime,
		}
	}

	return pl, nil
}

func (p *PlayRepo) GetGameMapByIds(ctx context.Context, ids ...int64) (map[int64]*biz.Game, error) {
	var l []*Game
	if err := p.data.DB(ctx).Table("soccer_game").Where("id IN (?)", ids).Find(&l).Error; err != nil {
		return nil, errors.NotFound("TEAMS_NOT_FOUND", "????????????????????????")
	}

	pl := make(map[int64]*biz.Game, 0)
	for _, v := range l {
		pl[v.ID] = &biz.Game{
			ID:         v.ID,
			Name:       v.Name,
			RedTeamId:  v.RedTeamId,
			BlueTeamId: v.BlueTeamId,
			StartTime:  v.StartTime,
			EndTime:    v.EndTime,
		}
	}

	return pl, nil
}

func (p *PlayRepo) GetPlayById(ctx context.Context, playId int64) (*biz.Play, error) {
	var play Play
	if result := p.data.DB(ctx).Table("goal_play").Where(&Play{ID: playId}).Find(&play); result.Error != nil {
		return nil, errors.InternalServer("SELECT_PLAY_ERROR", "??????????????????")
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

func (pgr *PlayGameRelRepo) GetPlayGameRelByGameId(ctx context.Context, gameId int64) ([]*biz.PlayGameRel, error) {
	var l []*PlayGameRel
	if result := pgr.data.DB(ctx).Table("goal_play_game_rel").Where(&PlayGameRel{GameId: gameId}).Find(&l); result.Error != nil {
		return nil, errors.InternalServer("SELECT_PLAY_GAME_REL_ERROR", "??????????????????????????????")
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

func (pgr *PlayGameRelRepo) GetPlayGameRelByGameIdAndPlayIds(ctx context.Context, gameId int64, playIds ...int64) ([]*biz.PlayGameRel, error) {
	var l []*PlayGameRel
	if result := pgr.data.DB(ctx).Table("goal_play_game_rel").Where("game_id = ? and play_id IN (?)", gameId, playIds).Find(&l); result.Error != nil {
		return nil, errors.InternalServer("SELECT_PLAY_GAME_REL_ERROR", "??????????????????????????????")
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

func (pgr *PlayGameRelRepo) GetPlayGameRelByPlayId(ctx context.Context, playId int64) (*biz.PlayGameRel, error) {
	var playGameRel PlayGameRel
	if result := pgr.data.DB(ctx).Table("goal_play_game_rel").Where(&PlayGameRel{PlayId: playId}).First(&playGameRel); result.Error != nil {
		return nil, errors.InternalServer("SELECT_PLAY_GAME_REL_ERROR", "??????????????????????????????")
	}

	return &biz.PlayGameRel{
		ID:     playGameRel.ID,
		PlayId: playGameRel.PlayId,
		GameId: playGameRel.GameId,
	}, nil
}

func (pgr *PlayGameRelRepo) GetPlayGameRelMapByPlayId(ctx context.Context, playIds ...int64) (map[int64]*biz.PlayGameRel, error) {
	var l []*PlayGameRel
	if result := pgr.data.DB(ctx).Table("goal_play_game_rel").Where("play_id IN (?)", playIds).Find(&l); result.Error != nil {
		return nil, errors.InternalServer("SELECT_PLAY_GAME_REL_ERROR", "??????????????????????????????")
	}

	pl := make(map[int64]*biz.PlayGameRel, 0)
	for _, v := range l {
		pl[v.PlayId] = &biz.PlayGameRel{
			ID:     v.ID,
			PlayId: v.PlayId,
			GameId: v.GameId,
		}
	}
	return pl, nil
}

func (psr *PlaySortRelRepo) GetPlaySortRelBySortIds(ctx context.Context, sortIds ...int64) ([]*biz.PlaySortRel, error) {
	var l []*PlaySortRel
	if result := psr.data.DB(ctx).Table("goal_play_game_sort_rel").Where("sort_id IN (?)", sortIds).Find(&l); result.Error != nil {
		return nil, errors.InternalServer("SELECT_PLAY_ERROR", "??????????????????????????????")
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

func (psr *PlaySortRelRepo) GetPlaySortRelByPlayIds(ctx context.Context, playIds ...int64) ([]*biz.PlaySortRel, error) {
	var l []*PlaySortRel
	if result := psr.data.DB(ctx).Table("goal_play_game_sort_rel").Where("play_id IN (?)", playIds).Find(&l); result.Error != nil {
		return nil, errors.InternalServer("SELECT_PLAY_ERROR", "??????????????????????????????")
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

func (prr *PlayRoomRelRepo) GetPlayRoomRelByRoomId(ctx context.Context, roomId int64) ([]*biz.PlayRoomRel, error) {
	var l []*PlayRoomRel
	if result := prr.data.DB(ctx).Table("goal_play_room_rel").Where(&PlayRoomRel{RoomId: roomId}).Find(&l); result.Error != nil {
		return nil, errors.InternalServer("SELECT_PLAY_ERROR", "???????????????????????????????????????")
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
		return nil, errors.New(500, "CREATE_PLAY_ERROR", "??????????????????")
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
		return nil, errors.New(500, "CREATE_PLAY_GAME_REL_ERROR", "?????????????????????????????????")
	}

	return &biz.PlayGameRel{
		ID:     playGameRel.ID,
		PlayId: playGameRel.PlayId,
		GameId: playGameRel.GameId,
	}, nil
}

// CreatePlayRoomRel .
func (prr *PlayRoomRelRepo) CreatePlayRoomRel(ctx context.Context, rel *biz.PlayRoomRel) (*biz.PlayRoomRel, error) {
	var playRoomRel PlayRoomRel
	playRoomRel.RoomId = rel.RoomId
	playRoomRel.PlayId = rel.PlayId
	res := prr.data.DB(ctx).Table("goal_play_room_rel").Create(&playRoomRel)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_PLAY_ROOM_ERROR", "????????????????????????????????????")
	}

	return &biz.PlayRoomRel{
		ID:     playRoomRel.ID,
		RoomId: playRoomRel.RoomId,
		PlayId: playRoomRel.PlayId,
	}, nil
}

// CreatePlaySortRel .
func (psr *PlaySortRelRepo) CreatePlaySortRel(ctx context.Context, rel *biz.PlaySortRel) (*biz.PlaySortRel, error) {
	var playSortRel PlaySortRel
	playSortRel.SortId = rel.SortId
	playSortRel.PlayId = rel.PlayId
	res := psr.data.DB(ctx).Table("goal_play_game_sort_rel").Create(&playSortRel)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_PLAY_SORT_REL_ERROR", "???????????????????????????????????????")
	}

	return &biz.PlaySortRel{
		ID:     playSortRel.ID,
		PlayId: playSortRel.PlayId,
		SortId: playSortRel.SortId,
	}, nil
}

// CreatePlayGameScoreUserRel .
func (p *PlayGameScoreUserRelRepo) CreatePlayGameScoreUserRel(ctx context.Context, pr *biz.PlayGameScoreUserRel) (*biz.PlayGameScoreUserRel, error) {
	var playGameScoreUserRel PlayGameScoreUserRel
	playGameScoreUserRel.UserId = pr.UserId
	playGameScoreUserRel.PlayId = pr.PlayId
	playGameScoreUserRel.Pay = pr.Pay
	playGameScoreUserRel.OriginPay = pr.OriginPay
	playGameScoreUserRel.Content = pr.Content
	playGameScoreUserRel.Status = pr.Status
	res := p.data.DB(ctx).Table("play_game_score_user_rel").Create(&playGameScoreUserRel)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_PLAY_GAME_SCORE_REL_ERROR", "????????????????????????????????????")
	}

	return &biz.PlayGameScoreUserRel{
		ID:      playGameScoreUserRel.ID,
		UserId:  playGameScoreUserRel.UserId,
		PlayId:  playGameScoreUserRel.PlayId,
		Pay:     playGameScoreUserRel.Pay,
		Content: playGameScoreUserRel.Content,
		Status:  playGameScoreUserRel.Status,
	}, nil
}

// UpdatePlayGameScoreUserRel .
func (p *PlayGameScoreUserRelRepo) UpdatePlayGameScoreUserRel(ctx context.Context, pr *biz.PlayGameScoreUserRel) (*biz.PlayGameScoreUserRel, error) {
	var playGameScoreUserRel PlayGameScoreUserRel
	playGameScoreUserRel.Pay = pr.Pay
	res := p.data.DB(ctx).Table("play_game_score_user_rel").Where("id=?", pr.ID).Updates(&playGameScoreUserRel)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_PLAY_GAME_SCORE_REL_ERROR", "????????????????????????????????????")
	}

	return &biz.PlayGameScoreUserRel{
		ID:      playGameScoreUserRel.ID,
		UserId:  playGameScoreUserRel.UserId,
		PlayId:  playGameScoreUserRel.PlayId,
		Pay:     playGameScoreUserRel.Pay,
		Content: playGameScoreUserRel.Content,
		Status:  playGameScoreUserRel.Status,
	}, nil
}

func (p *PlayGameScoreUserRelRepo) GetPlayGameScoreUserRelByUserId(ctx context.Context, userId int64) ([]*biz.PlayGameScoreUserRel, error) {
	var l []*PlayGameScoreUserRel
	if result := p.data.DB(ctx).Table("play_game_score_user_rel").Where(&PlayGameScoreUserRel{UserId: userId}).Find(&l); result.Error != nil {
		return nil, errors.InternalServer("SELECT_PLAY_GAME_SCORE_REL_ERROR", "??????????????????????????????")
	}

	pl := make([]*biz.PlayGameScoreUserRel, 0)
	for _, playGameScoreUserRel := range l {
		pl = append(pl, &biz.PlayGameScoreUserRel{
			ID:        playGameScoreUserRel.ID,
			UserId:    playGameScoreUserRel.UserId,
			PlayId:    playGameScoreUserRel.PlayId,
			Pay:       playGameScoreUserRel.Pay,
			OriginPay: playGameScoreUserRel.OriginPay,
			Content:   playGameScoreUserRel.Content,
			Status:    playGameScoreUserRel.Status,
			CreatedAt: playGameScoreUserRel.CreatedAt,
		})
	}
	return pl, nil
}

func (p *PlayGameScoreUserRelRepo) GetPlayGameScoreUserRelByPlayIds(ctx context.Context, playIds ...int64) ([]*biz.PlayGameScoreUserRel, error) {
	var l []*PlayGameScoreUserRel
	if result := p.data.DB(ctx).Table("play_game_score_user_rel").Where("play_id IN(?)", playIds).Find(&l); result.Error != nil {
		return nil, errors.InternalServer("SELECT_PLAY_GAME_SCORE_REL_ERROR", "??????????????????????????????")
	}

	pl := make([]*biz.PlayGameScoreUserRel, 0)
	for _, playGameScoreUserRel := range l {
		pl = append(pl, &biz.PlayGameScoreUserRel{
			ID:        playGameScoreUserRel.ID,
			UserId:    playGameScoreUserRel.UserId,
			CreatedAt: playGameScoreUserRel.CreatedAt,
		})
	}
	return pl, nil
}

func (p *PlayGameTeamResultUserRelRepo) GetPlayGameTeamResultUserRelByPlayIds(ctx context.Context, playIds ...int64) ([]*biz.PlayGameTeamResultUserRel, error) {
	var l []*PlayGameTeamResultUserRel
	if result := p.data.DB(ctx).Table("play_game_team_result_user_rel").Where("play_id IN(?)", playIds).Find(&l); result.Error != nil {
		return nil, errors.InternalServer("SELECT_PLAY_GAME_SCORE_REL_ERROR", "??????????????????????????????")
	}

	pl := make([]*biz.PlayGameTeamResultUserRel, 0)
	for _, playGameTeamResultUserRel := range l {
		pl = append(pl, &biz.PlayGameTeamResultUserRel{
			ID:        playGameTeamResultUserRel.ID,
			UserId:    playGameTeamResultUserRel.UserId,
			CreatedAt: playGameTeamResultUserRel.CreatedAt,
		})
	}
	return pl, nil
}

func (p *PlayGameTeamResultUserRelRepo) GetPlayGameTeamResultUserRelByPlayId(ctx context.Context, playId int64) ([]*biz.PlayGameTeamResultUserRel, error) {
	var l []*PlayGameTeamResultUserRel
	if result := p.data.DB(ctx).Table("play_game_team_result_user_rel").Where("play_id=?", playId).Find(&l); result.Error != nil {
		return nil, errors.InternalServer("SELECT_PLAY_GAME_RESULT_REL_ERROR", "??????????????????????????????")
	}

	pl := make([]*biz.PlayGameTeamResultUserRel, 0)
	for _, playGameTeamResultUserRel := range l {
		pl = append(pl, &biz.PlayGameTeamResultUserRel{
			ID:        playGameTeamResultUserRel.ID,
			PlayId:    playGameTeamResultUserRel.PlayId,
			Pay:       playGameTeamResultUserRel.Pay,
			Content:   playGameTeamResultUserRel.Content,
			UserId:    playGameTeamResultUserRel.UserId,
			CreatedAt: playGameTeamResultUserRel.CreatedAt,
		})
	}
	return pl, nil
}

func (p *PlayGameScoreUserRelRepo) GetPlayGameScoreUserRelByPlayId(ctx context.Context, playId int64) ([]*biz.PlayGameScoreUserRel, error) {
	var l []*PlayGameScoreUserRel
	if result := p.data.DB(ctx).Table("play_game_score_user_rel").Where("play_id=?", playId).Find(&l); result.Error != nil {
		return nil, errors.InternalServer("SELECT_PLAY_GAME_SCORE_REL_ERROR", "??????????????????????????????")
	}

	pl := make([]*biz.PlayGameScoreUserRel, 0)
	for _, playGameScoreUserRel := range l {
		pl = append(pl, &biz.PlayGameScoreUserRel{
			ID:        playGameScoreUserRel.ID,
			PlayId:    playGameScoreUserRel.PlayId,
			Content:   playGameScoreUserRel.Content,
			Pay:       playGameScoreUserRel.Pay,
			UserId:    playGameScoreUserRel.UserId,
			CreatedAt: playGameScoreUserRel.CreatedAt,
		})
	}
	return pl, nil
}

func (p *PlayGameTeamSortUserRelRepo) GetPlayTeamSortUserRelByPlayId(ctx context.Context, playId int64) ([]*biz.PlayGameTeamSortUserRel, error) {
	var l []*PlayGameTeamSortUserRel
	if result := p.data.DB(ctx).Table("play_game_team_sort_user_rel").Where("play_id=?", playId).Find(&l); result.Error != nil {
		return nil, errors.InternalServer("SELECT_PLAY_GAME_SCORE_REL_ERROR", "??????????????????????????????")
	}

	pl := make([]*biz.PlayGameTeamSortUserRel, 0)
	for _, playGameTeamSortUserRel := range l {
		pl = append(pl, &biz.PlayGameTeamSortUserRel{
			UserId:    playGameTeamSortUserRel.UserId,
			PlayId:    playGameTeamSortUserRel.PlayId,
			SortId:    playGameTeamSortUserRel.SortId,
			Content:   playGameTeamSortUserRel.Content,
			Pay:       playGameTeamSortUserRel.Pay,
			CreatedAt: playGameTeamSortUserRel.CreatedAt,
		})
	}
	return pl, nil
}

func (p *PlayGameTeamGoalUserRelRepo) GetPlayGameGoalUserRelByPlayId(ctx context.Context, playId int64) ([]*biz.PlayGameTeamGoalUserRel, error) {
	var l []*PlayGameTeamGoalUserRel
	if result := p.data.DB(ctx).Table("play_game_team_goal_user_rel").Where("play_id=?", playId).Find(&l); result.Error != nil {
		return nil, errors.InternalServer("SELECT_PLAY_GAME_SCORE_REL_ERROR", "?????????????????????????????????")
	}

	pl := make([]*biz.PlayGameTeamGoalUserRel, 0)
	for _, playGameTeamGoalUserRel := range l {
		pl = append(pl, &biz.PlayGameTeamGoalUserRel{
			ID:        playGameTeamGoalUserRel.ID,
			UserId:    playGameTeamGoalUserRel.UserId,
			PlayId:    playGameTeamGoalUserRel.PlayId,
			Pay:       playGameTeamGoalUserRel.Pay,
			OriginPay: playGameTeamGoalUserRel.OriginPay,
			Status:    playGameTeamGoalUserRel.Status,
			Goal:      playGameTeamGoalUserRel.Goal,
			TeamId:    playGameTeamGoalUserRel.TeamId,
			Type:      playGameTeamGoalUserRel.Type,
			CreatedAt: playGameTeamGoalUserRel.CreatedAt,
		})
	}
	return pl, nil
}

func (p *PlayGameTeamGoalUserRelRepo) GetPlayGameTeamGoalUserRelByPlayIdTotal(ctx context.Context, playId int64) (*biz.PlayGameTeamGoalUserRelTotal, error) {
	var playGameTeamGoalUserRelTotal *PlayGameTeamGoalUserRelTotal
	if result := p.data.DB(ctx).Table("play_game_team_goal_user_rel").Select("sum(pay) as total").Where("play_id", playId).Take(&playGameTeamGoalUserRelTotal); result.Error != nil {
		return nil, errors.InternalServer("SELECT_PLAY_GAME_SCORE_REL_ERROR", "???????????????????????????")
	}

	return &biz.PlayGameTeamGoalUserRelTotal{
		Total: playGameTeamGoalUserRelTotal.Total,
	}, nil
}

func (p *PlayGameTeamResultUserRelRepo) GetPlayGameTeamResultUserRelByPlayIdTotal(ctx context.Context, playId int64) (*biz.PlayGameTeamResultUserRelTotal, error) {
	var playGameTeamResultUserRelTotal *PlayGameTeamResultUserRelTotal
	if result := p.data.DB(ctx).Table("play_game_team_result_user_rel").Select("sum(pay) as total").Where("play_id", playId).Take(&playGameTeamResultUserRelTotal); result.Error != nil {
		return nil, errors.InternalServer("SELECT_PLAY_GAME_RESULT_REL_ERROR", "????????????????????????")
	}

	return &biz.PlayGameTeamResultUserRelTotal{
		Total: playGameTeamResultUserRelTotal.Total,
	}, nil
}

func (p *PlayGameTeamSortUserRelRepo) GetPlayGameTeamSortUserRelByPlayIdTotal(ctx context.Context, playId int64) (*biz.PlayGameTeamSortUserRelTotal, error) {
	var playGameTeamSortUserRelTotal *PlayGameTeamGoalUserRelTotal
	if result := p.data.DB(ctx).Table("play_game_team_sort_user_rel").Select("sum(pay) as total").Where("play_id", playId).Take(&playGameTeamSortUserRelTotal); result.Error != nil {
		return nil, errors.InternalServer("SELECT_PLAY_GAME_SORT_REL_ERROR", "????????????????????????")
	}

	return &biz.PlayGameTeamSortUserRelTotal{
		Total: playGameTeamSortUserRelTotal.Total,
	}, nil
}
func (p *PlayGameScoreUserRelRepo) GetPlayGameScoreUserRelByPlayIdTotal(ctx context.Context, playId int64) (*biz.PlayGameScoreUserRelTotal, error) {
	var playGameScoreUserRelTotal *PlayGameScoreUserRelTotal
	if result := p.data.DB(ctx).Table("play_game_score_user_rel").Select("sum(pay) as total").Where("play_id", playId).Take(&playGameScoreUserRelTotal); result.Error != nil {
		return nil, errors.InternalServer("SELECT_PLAY_GAME_SCORE_REL_ERROR", "????????????????????????")
	}

	return &biz.PlayGameScoreUserRelTotal{
		Total: playGameScoreUserRelTotal.Total,
	}, nil
}

func (p *PlayGameTeamGoalUserRelRepo) GetPlayGameTeamGoalUserRelByPlayIds(ctx context.Context, playIds ...int64) ([]*biz.PlayGameTeamGoalUserRel, error) {
	var l []*PlayGameTeamGoalUserRel
	if result := p.data.DB(ctx).Table("play_game_team_goal_user_rel").Where("play_id IN(?)", playIds).Find(&l); result.Error != nil {
		return nil, errors.InternalServer("SELECT_PLAY_GAME_SCORE_REL_ERROR", "?????????????????????????????????")
	}

	pl := make([]*biz.PlayGameTeamGoalUserRel, 0)
	for _, playGameScoreUserRel := range l {
		pl = append(pl, &biz.PlayGameTeamGoalUserRel{
			ID:        playGameScoreUserRel.ID,
			UserId:    playGameScoreUserRel.UserId,
			CreatedAt: playGameScoreUserRel.CreatedAt,
		})
	}
	return pl, nil
}

func (p *PlayGameTeamSortUserRelRepo) GetPlayGameTeamScoreUserRelByPlayIds(ctx context.Context, playIds ...int64) ([]*biz.PlayGameTeamSortUserRel, error) {
	var l []*PlayGameTeamSortUserRel
	if result := p.data.DB(ctx).Table("play_game_team_sort_user_rel").Where("play_id IN(?)", playIds).Find(&l); result.Error != nil {
		return nil, errors.InternalServer("SELECT_PLAY_GAME_SCORE_REL_ERROR", "??????????????????????????????")
	}

	pl := make([]*biz.PlayGameTeamSortUserRel, 0)
	for _, playGameScoreUserRel := range l {
		pl = append(pl, &biz.PlayGameTeamSortUserRel{
			ID:        playGameScoreUserRel.ID,
			UserId:    playGameScoreUserRel.UserId,
			CreatedAt: playGameScoreUserRel.CreatedAt,
		})
	}
	return pl, nil
}

// CreatePlayGameTeamResultUserRel .
func (p *PlayGameTeamResultUserRelRepo) CreatePlayGameTeamResultUserRel(ctx context.Context, pr *biz.PlayGameTeamResultUserRel) (*biz.PlayGameTeamResultUserRel, error) {
	var playGameTeamResultUserRel PlayGameTeamResultUserRel
	playGameTeamResultUserRel.UserId = pr.UserId
	playGameTeamResultUserRel.PlayId = pr.PlayId
	playGameTeamResultUserRel.Pay = pr.Pay
	playGameTeamResultUserRel.OriginPay = pr.OriginPay
	playGameTeamResultUserRel.Content = pr.Content
	playGameTeamResultUserRel.Status = pr.Status
	res := p.data.DB(ctx).Table("play_game_team_result_user_rel").Create(&playGameTeamResultUserRel)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_PLAY_GAME_RESULT_REL_ERROR", "????????????????????????????????????")
	}

	return &biz.PlayGameTeamResultUserRel{
		ID:      playGameTeamResultUserRel.ID,
		UserId:  playGameTeamResultUserRel.UserId,
		PlayId:  playGameTeamResultUserRel.PlayId,
		Pay:     playGameTeamResultUserRel.Pay,
		Content: playGameTeamResultUserRel.Content,
	}, nil
}

// UpdatePlayGameTeamResultUserRel .
func (p *PlayGameTeamResultUserRelRepo) UpdatePlayGameTeamResultUserRel(ctx context.Context, pr *biz.PlayGameTeamResultUserRel) (*biz.PlayGameTeamResultUserRel, error) {
	var playGameTeamResultUserRel PlayGameTeamResultUserRel
	playGameTeamResultUserRel.Pay = pr.Pay
	res := p.data.DB(ctx).Table("play_game_team_result_user_rel").Where("id=?", pr.ID).Updates(&playGameTeamResultUserRel)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_PLAY_GAME_RESULT_REL_ERROR", "????????????????????????????????????")
	}

	return &biz.PlayGameTeamResultUserRel{
		ID:      playGameTeamResultUserRel.ID,
		UserId:  playGameTeamResultUserRel.UserId,
		PlayId:  playGameTeamResultUserRel.PlayId,
		Pay:     playGameTeamResultUserRel.Pay,
		Content: playGameTeamResultUserRel.Content,
	}, nil
}

func (p *PlayGameTeamResultUserRelRepo) GetPlayGameTeamResultUserRelByUserId(ctx context.Context, userId int64) ([]*biz.PlayGameTeamResultUserRel, error) {
	var l []*PlayGameTeamResultUserRel
	if result := p.data.DB(ctx).Table("play_game_team_result_user_rel").Where(&PlayGameTeamResultUserRel{UserId: userId}).Find(&l); result.Error != nil {
		return nil, errors.InternalServer("SELECT_PLAY_GAME_RESULT_REL_ERROR", "??????????????????????????????")
	}

	pl := make([]*biz.PlayGameTeamResultUserRel, 0)
	for _, playGameTeamResultUserRel := range l {
		pl = append(pl, &biz.PlayGameTeamResultUserRel{
			ID:        playGameTeamResultUserRel.ID,
			UserId:    playGameTeamResultUserRel.UserId,
			PlayId:    playGameTeamResultUserRel.PlayId,
			Pay:       playGameTeamResultUserRel.Pay,
			OriginPay: playGameTeamResultUserRel.OriginPay,
			Content:   playGameTeamResultUserRel.Content,
			CreatedAt: playGameTeamResultUserRel.CreatedAt,
			Status:    playGameTeamResultUserRel.Status,
		})
	}
	return pl, nil
}

// CreatePlayGameTeamGoalUserRel .
func (p *PlayGameTeamGoalUserRelRepo) CreatePlayGameTeamGoalUserRel(ctx context.Context, pr *biz.PlayGameTeamGoalUserRel) (*biz.PlayGameTeamGoalUserRel, error) {
	var playGameTeamGoalUserRel PlayGameTeamGoalUserRel
	playGameTeamGoalUserRel.UserId = pr.UserId
	playGameTeamGoalUserRel.PlayId = pr.PlayId
	playGameTeamGoalUserRel.Pay = pr.Pay
	playGameTeamGoalUserRel.TeamId = pr.TeamId
	playGameTeamGoalUserRel.Status = pr.Status
	playGameTeamGoalUserRel.Goal = pr.Goal
	playGameTeamGoalUserRel.Type = pr.Type
	playGameTeamGoalUserRel.OriginPay = pr.OriginPay
	res := p.data.DB(ctx).Table("play_game_team_goal_user_rel").Create(&playGameTeamGoalUserRel)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_PLAY_GAME_GOAL_REL_ERROR", "???????????????????????????????????????")
	}

	return &biz.PlayGameTeamGoalUserRel{
		ID:     playGameTeamGoalUserRel.ID,
		UserId: playGameTeamGoalUserRel.UserId,
		PlayId: playGameTeamGoalUserRel.PlayId,
		Pay:    playGameTeamGoalUserRel.Pay,
		Status: playGameTeamGoalUserRel.Status,
		Goal:   playGameTeamGoalUserRel.Goal,
		TeamId: playGameTeamGoalUserRel.TeamId,
		Type:   playGameTeamGoalUserRel.Type,
	}, nil
}

// UpdatePlayGameTeamGoalUserRel .
func (p *PlayGameTeamGoalUserRelRepo) UpdatePlayGameTeamGoalUserRel(ctx context.Context, pr *biz.PlayGameTeamGoalUserRel) (*biz.PlayGameTeamGoalUserRel, error) {
	var playGameTeamGoalUserRel PlayGameTeamGoalUserRel
	playGameTeamGoalUserRel.Pay = pr.Pay
	res := p.data.DB(ctx).Table("play_game_team_goal_user_rel").Where("id=?", pr.ID).Updates(&playGameTeamGoalUserRel)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_PLAY_GAME_GOAL_REL_ERROR", "???????????????????????????????????????")
	}

	return &biz.PlayGameTeamGoalUserRel{
		ID:     playGameTeamGoalUserRel.ID,
		UserId: playGameTeamGoalUserRel.UserId,
		PlayId: playGameTeamGoalUserRel.PlayId,
		Pay:    playGameTeamGoalUserRel.Pay,
		Status: playGameTeamGoalUserRel.Status,
		Goal:   playGameTeamGoalUserRel.Goal,
		TeamId: playGameTeamGoalUserRel.TeamId,
		Type:   playGameTeamGoalUserRel.Type,
	}, nil
}

func (p *PlayGameTeamGoalUserRelRepo) GetPlayGameTeamGoalUserRelByUserId(ctx context.Context, userId int64) ([]*biz.PlayGameTeamGoalUserRel, error) {
	var l []*PlayGameTeamGoalUserRel
	if result := p.data.DB(ctx).Table("play_game_team_goal_user_rel").Where(&PlayGameTeamGoalUserRel{UserId: userId}).Find(&l); result.Error != nil {
		return nil, errors.InternalServer("SELECT_PLAY_GAME_GOAL_REL_ERROR", "?????????????????????????????????")
	}

	pl := make([]*biz.PlayGameTeamGoalUserRel, 0)
	for _, playGameTeamGoalUserRel := range l {
		pl = append(pl, &biz.PlayGameTeamGoalUserRel{
			ID:        playGameTeamGoalUserRel.ID,
			UserId:    playGameTeamGoalUserRel.UserId,
			PlayId:    playGameTeamGoalUserRel.PlayId,
			Pay:       playGameTeamGoalUserRel.Pay,
			OriginPay: playGameTeamGoalUserRel.OriginPay,
			Status:    playGameTeamGoalUserRel.Status,
			Goal:      playGameTeamGoalUserRel.Goal,
			TeamId:    playGameTeamGoalUserRel.TeamId,
			Type:      playGameTeamGoalUserRel.Type,
			CreatedAt: playGameTeamGoalUserRel.CreatedAt,
		})
	}
	return pl, nil
}

// CreatePlayGameTeamSortUserRel .
func (p *PlayGameTeamSortUserRelRepo) CreatePlayGameTeamSortUserRel(ctx context.Context, pr *biz.PlayGameTeamSortUserRel) (*biz.PlayGameTeamSortUserRel, error) {
	var playGameTeamSortUserRel PlayGameTeamSortUserRel
	playGameTeamSortUserRel.UserId = pr.UserId
	playGameTeamSortUserRel.PlayId = pr.PlayId
	playGameTeamSortUserRel.Pay = pr.Pay
	playGameTeamSortUserRel.OriginPay = pr.OriginPay
	playGameTeamSortUserRel.Status = pr.Status
	playGameTeamSortUserRel.Content = pr.Content
	playGameTeamSortUserRel.SortId = pr.SortId
	res := p.data.DB(ctx).Table("play_game_team_sort_user_rel").Create(&playGameTeamSortUserRel)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_PLAY_GAME_SORT_REL_ERROR", "??????????????????????????????????????????")
	}

	return &biz.PlayGameTeamSortUserRel{
		ID:      playGameTeamSortUserRel.ID,
		UserId:  playGameTeamSortUserRel.UserId,
		PlayId:  playGameTeamSortUserRel.PlayId,
		Pay:     playGameTeamSortUserRel.Pay,
		Status:  playGameTeamSortUserRel.Status,
		Content: playGameTeamSortUserRel.Content,
		SortId:  playGameTeamSortUserRel.SortId,
	}, nil

}

// UpdatePlayGameTeamSortUserRel .
func (p *PlayGameTeamSortUserRelRepo) UpdatePlayGameTeamSortUserRel(ctx context.Context, pr *biz.PlayGameTeamSortUserRel) (*biz.PlayGameTeamSortUserRel, error) {
	var playGameTeamSortUserRel PlayGameTeamSortUserRel
	playGameTeamSortUserRel.Pay = pr.Pay
	res := p.data.DB(ctx).Table("play_game_team_sort_user_rel").Where("id=?", pr.ID).Updates(&playGameTeamSortUserRel)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_PLAY_GAME_SORT_REL_ERROR", "??????????????????????????????????????????")
	}

	return &biz.PlayGameTeamSortUserRel{
		ID:      playGameTeamSortUserRel.ID,
		UserId:  playGameTeamSortUserRel.UserId,
		PlayId:  playGameTeamSortUserRel.PlayId,
		Pay:     playGameTeamSortUserRel.Pay,
		Status:  playGameTeamSortUserRel.Status,
		Content: playGameTeamSortUserRel.Content,
		SortId:  playGameTeamSortUserRel.SortId,
	}, nil

}

func (p *PlayGameTeamSortUserRelRepo) GetPlayGameTeamSortUserRelByUserId(ctx context.Context, userId int64) ([]*biz.PlayGameTeamSortUserRel, error) {
	var l []*PlayGameTeamSortUserRel
	if result := p.data.DB(ctx).Table("play_game_team_sort_user_rel").Where(&PlayGameTeamSortUserRel{UserId: userId}).Find(&l); result.Error != nil {
		return nil, errors.InternalServer("SELECT_PLAY_SORT_GOAL_REL_ERROR", "??????????????????????????????")
	}

	pl := make([]*biz.PlayGameTeamSortUserRel, 0)
	for _, playGameTeamSortUserRel := range l {
		pl = append(pl, &biz.PlayGameTeamSortUserRel{
			ID:        playGameTeamSortUserRel.ID,
			UserId:    playGameTeamSortUserRel.UserId,
			PlayId:    playGameTeamSortUserRel.PlayId,
			Pay:       playGameTeamSortUserRel.Pay,
			OriginPay: playGameTeamSortUserRel.OriginPay,
			Status:    playGameTeamSortUserRel.Status,
			Content:   playGameTeamSortUserRel.Content,
			SortId:    playGameTeamSortUserRel.SortId,
			CreatedAt: playGameTeamSortUserRel.CreatedAt,
		})
	}
	return pl, nil
}

func (p *PlayRepo) GetUserByUserIds(ctx context.Context, userIds ...int64) ([]*biz.User, error) {
	var l []*User
	if result := p.data.DB(ctx).Table("user").Where("ID IN(?)", userIds).Find(&l); result.Error != nil {
		return nil, errors.InternalServer("SELECT_PLAY_GAME_SCORE_REL_ERROR", "??????????????????????????????")
	}

	pl := make([]*biz.User, 0)
	for _, v := range l {
		pl = append(pl, &biz.User{
			ID:      v.ID,
			Address: v.Address,
		})
	}
	return pl, nil
}

func (s *SystemConfigRepo) GetSystemConfigByName(ctx context.Context, name string) (*biz.SystemConfig, error) {
	var config *SystemConfig
	if err := s.data.DB(ctx).Table("system_config").Where("name=?", name).First(&config).Error; err != nil {
		return nil, errors.NotFound("TEAMS_NOT_FOUND", "??????????????????")
	}

	return &biz.SystemConfig{
		ID:    config.ID,
		Name:  config.Name,
		Value: config.Value,
	}, nil
}

func (s *SystemConfigRepo) GetSystemConfigByNames(ctx context.Context, name ...string) (map[string]*biz.SystemConfig, error) {
	var l []*SystemConfig
	if err := s.data.DB(ctx).Table("system_config").Where("name IN (?)", name).Find(&l).Error; err != nil {
		return nil, errors.NotFound("TEAMS_NOT_FOUND", "????????????????????????")
	}

	pl := make(map[string]*biz.SystemConfig, 0)
	for _, v := range l {
		pl[v.Name] = &biz.SystemConfig{
			Value: v.Value,
		}
	}

	return pl, nil
}

// GetLastTermPoolByPlayIdAndType .
func (p *PlayRepo) GetLastTermPoolByPlayIdAndType(ctx context.Context, playId int64, playType string) (*biz.LastTermPool, error) {
	var pool LastTermPool
	res := p.data.DB(ctx).Table("last_term_pool").Where("play_id=? and play_type=?", playId, playType).First(&pool)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_PLAY_ERROR", "?????????????????????")
	}

	return &biz.LastTermPool{
		ID:             pool.ID,
		GameId:         pool.GameId,
		OriginGameId:   pool.OriginGameId,
		PlayId:         pool.PlayId,
		OriginPlayId:   pool.OriginPlayId,
		Total:          pool.Total,
		PlayType:       pool.PlayType,
		OriginPlayType: pool.OriginPlayType,
	}, nil
}
