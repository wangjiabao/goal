package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "goal/api/play/service/v1"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Play struct {
	ID             int64
	CreateUserId   int64
	CreateUserType string
	Type           string
	StartTime      time.Time
	EndTime        time.Time
}

type LastTermPool struct {
	ID             int64
	GameId         int64
	OriginGameId   int64
	PlayId         int64
	OriginPlayId   int64
	Total          int64
	PlayType       string
	OriginPlayType string
}

type BalanceRecordIdRel struct {
	ID       int64
	RecordId int64
	RelType  string
	RelId    int64
}

type PlayGameRel struct {
	ID     int64
	PlayId int64
	GameId int64
}

type PlaySortRel struct {
	ID     int64
	PlayId int64
	SortId int64
}

type PlayRoomRel struct {
	ID     int64
	RoomId int64
	PlayId int64
}

type UserBalanceRecord struct {
	ID     int64
	Amount int64
}

type PlayGameScoreUserRel struct {
	ID        int64
	UserId    int64
	PlayId    int64
	Content   string
	Pay       int64
	OriginPay int64
	Status    string
	CreatedAt time.Time
}

type PlayGameTeamSortUserRel struct {
	ID        int64
	UserId    int64
	PlayId    int64
	SortId    int64
	Status    string
	Content   string
	Pay       int64
	OriginPay int64
	CreatedAt time.Time
}

type PlayGameTeamGoalUserRel struct {
	ID        int64
	UserId    int64
	PlayId    int64
	TeamId    int64
	Type      string
	Pay       int64
	OriginPay int64
	Goal      int64
	Status    string
	CreatedAt time.Time
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
	ID        int64
	UserId    int64
	PlayId    int64
	Content   string
	Pay       int64
	OriginPay int64
	Status    string
	CreatedAt time.Time
}

type PlayAllTypeUserRel struct {
	ID         int64
	GameName   string
	RedTeamId  int64
	BlueTeamId int64
	PlayId     int64
	Pay        int64
	Status     string
	Content    string
	Type       string
	Goal       int64
	TeamId     int64
	SortId     int64
	CreatedAt  time.Time
}

type UserInfo struct {
	ID              int64
	UserId          int64
	Name            string
	Avatar          string
	RecommendCode   string
	MyRecommendCode string
	Code            string
	CreatedAt       time.Time
}

type PlayAllTypeUserRelSlice []*PlayAllTypeUserRel

func (p PlayAllTypeUserRelSlice) Len() int           { return len(p) }
func (p PlayAllTypeUserRelSlice) Less(i, j int) bool { return p[i].CreatedAt.After(p[j].CreatedAt) }
func (p PlayAllTypeUserRelSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type UserBalance struct {
	ID      int64
	UserId  int64
	Balance int64
}

type UserInfoRepo interface {
	GetUserInfoByUserId(ctx context.Context, userId int64) (*UserInfo, error)
}

type UserProxy struct {
	ID       int64
	UserId   int64
	UpUserId int64
	Rate     int64
}

type PlayRepo interface {
	GetAdminCreatePlayList(ctx context.Context) ([]*Play, error)
	GetAdminCreatePlayListByIds(ctx context.Context, ids ...int64) ([]*Play, error)
	GetPlayListByIds(ctx context.Context, ids ...int64) ([]*Play, error)
	GetGameMapByIds(ctx context.Context, ids ...int64) (map[int64]*Game, error)
	GetPlayMapByIds(ctx context.Context, ids ...int64) (map[int64]*Play, error)
	CreatePlay(ctx context.Context, pc *Play) (*Play, error)
	GetAdminCreatePlayListByType(ctx context.Context, playType string) ([]*Play, error)
	GetPlayById(ctx context.Context, playId int64) (*Play, error)
	GetAdminCreatePlayByType(ctx context.Context, playType string) (*Play, error)
	GetUserByUserIds(ctx context.Context, ids ...int64) ([]*User, error)
	GetLastTermPoolByPlayIdAndType(ctx context.Context, playId int64, playType string) (*LastTermPool, error)
}

type PlayGameRelRepo interface {
	GetPlayGameRelByGameId(ctx context.Context, gameId int64) ([]*PlayGameRel, error)
	GetPlayGameRelByPlayId(ctx context.Context, playId int64) (*PlayGameRel, error)
	GetPlayGameRelByGameIdAndPlayIds(ctx context.Context, gameId int64, playIds ...int64) ([]*PlayGameRel, error)
	GetPlayGameRelMapByPlayId(ctx context.Context, playIds ...int64) (map[int64]*PlayGameRel, error)
	CreatePlayGameRel(ctx context.Context, rel *PlayGameRel) (*PlayGameRel, error)
}

type PlaySortRelRepo interface {
	GetPlaySortRelBySortIds(ctx context.Context, sortIds ...int64) ([]*PlaySortRel, error)
	GetPlaySortRelByPlayIds(ctx context.Context, playIds ...int64) ([]*PlaySortRel, error)
	CreatePlaySortRel(ctx context.Context, rel *PlaySortRel) (*PlaySortRel, error)
}

type PlayRoomRelRepo interface {
	GetPlayRoomRelByRoomId(ctx context.Context, roomId int64) ([]*PlayRoomRel, error)
	CreatePlayRoomRel(ctx context.Context, pc *PlayRoomRel) (*PlayRoomRel, error)
}

type UserBalanceRepo interface {
	Pay(ctx context.Context, userId int64, pay int64) (int64, error)
	TransferIntoUserGoalRecommendReward(ctx context.Context, userId int64, amount int64) (int64, error)
	RoomFee(ctx context.Context, userId int64, pay int64) (int64, error)
	GetBalanceRecordIdRelMap(ctx context.Context, relType string, id ...int64) (map[int64]*BalanceRecordIdRel, error)
	CreateBalanceRecordIdRel(ctx context.Context, recordId int64, relType string, id int64) error
	GetUserBalance(ctx context.Context, userId int64) (*UserBalance, error)
	GetUserBalanceRecordGoalReward(ctx context.Context, ids ...int64) (map[int64]*UserBalanceRecord, error)
	TransferIntoUserPlayProxyReward(ctx context.Context, userId int64, amount int64) (int64, error)
}

type UserProxyRepo interface {
	GetUserProxyAndDown(ctx context.Context) (map[int64]*UserProxy, map[int64]*UserProxy, error)
}

type PlayGameTeamResultUserRelRepo interface {
	GetPlayGameTeamResultUserRelByPlayIdTotal(ctx context.Context, playId int64) (*PlayGameTeamResultUserRelTotal, error)
	GetPlayGameTeamResultUserRelByUserId(ctx context.Context, userId int64) ([]*PlayGameTeamResultUserRel, error)
	CreatePlayGameTeamResultUserRel(ctx context.Context, pr *PlayGameTeamResultUserRel) (*PlayGameTeamResultUserRel, error)
	UpdatePlayGameTeamResultUserRel(ctx context.Context, pr *PlayGameTeamResultUserRel) (*PlayGameTeamResultUserRel, error)
	GetPlayGameTeamResultUserRelByPlayIds(ctx context.Context, playIds ...int64) ([]*PlayGameTeamResultUserRel, error)
	GetPlayGameTeamResultUserRelByPlayId(ctx context.Context, playId int64) ([]*PlayGameTeamResultUserRel, error)
}

type PlayGameTeamGoalUserRelRepo interface {
	GetPlayGameGoalUserRelByPlayId(ctx context.Context, playId int64) ([]*PlayGameTeamGoalUserRel, error)
	GetPlayGameTeamGoalUserRelByPlayIdTotal(ctx context.Context, playId int64) (*PlayGameTeamGoalUserRelTotal, error)
	GetPlayGameTeamGoalUserRelByUserId(ctx context.Context, userId int64) ([]*PlayGameTeamGoalUserRel, error)
	CreatePlayGameTeamGoalUserRel(ctx context.Context, pr *PlayGameTeamGoalUserRel) (*PlayGameTeamGoalUserRel, error)
	UpdatePlayGameTeamGoalUserRel(ctx context.Context, pr *PlayGameTeamGoalUserRel) (*PlayGameTeamGoalUserRel, error)
	GetPlayGameTeamGoalUserRelByPlayIds(ctx context.Context, playIds ...int64) ([]*PlayGameTeamGoalUserRel, error)
}

type PlayGameTeamSortUserRelRepo interface {
	GetPlayTeamSortUserRelByPlayId(ctx context.Context, playId int64) ([]*PlayGameTeamSortUserRel, error)
	GetPlayGameTeamSortUserRelByPlayIdTotal(ctx context.Context, playId int64) (*PlayGameTeamSortUserRelTotal, error)
	GetPlayGameTeamSortUserRelByUserId(ctx context.Context, userId int64) ([]*PlayGameTeamSortUserRel, error)
	CreatePlayGameTeamSortUserRel(ctx context.Context, pr *PlayGameTeamSortUserRel) (*PlayGameTeamSortUserRel, error)
	UpdatePlayGameTeamSortUserRel(ctx context.Context, pr *PlayGameTeamSortUserRel) (*PlayGameTeamSortUserRel, error)
	GetPlayGameTeamScoreUserRelByPlayIds(ctx context.Context, playIds ...int64) ([]*PlayGameTeamSortUserRel, error)
}

type PlayGameScoreUserRelRepo interface {
	GetPlayGameScoreUserRelByPlayIdTotal(ctx context.Context, playId int64) (*PlayGameScoreUserRelTotal, error)
	CreatePlayGameScoreUserRel(ctx context.Context, pr *PlayGameScoreUserRel) (*PlayGameScoreUserRel, error)
	UpdatePlayGameScoreUserRel(ctx context.Context, pr *PlayGameScoreUserRel) (*PlayGameScoreUserRel, error)
	GetPlayGameScoreUserRelByUserId(ctx context.Context, userId int64) ([]*PlayGameScoreUserRel, error)
	GetPlayGameScoreUserRelByPlayIds(ctx context.Context, playIds ...int64) ([]*PlayGameScoreUserRel, error)
	GetPlayGameScoreUserRelByPlayId(ctx context.Context, playId int64) ([]*PlayGameScoreUserRel, error)
}

type SystemConfigRepo interface {
	GetSystemConfigByName(ctx context.Context, name string) (*SystemConfig, error)
	GetSystemConfigByNames(ctx context.Context, name ...string) (map[string]*SystemConfig, error)
}

type SystemConfig struct {
	ID    int64
	Name  string
	Value int64
}

type PlayUseCase struct {
	systemConfigRepo              SystemConfigRepo
	playRepo                      PlayRepo
	gameRepo                      GameRepo
	sortRepo                      SortRepo
	playGameRelRepo               PlayGameRelRepo
	playSortRelRepo               PlaySortRelRepo
	roomUserRelRepo               RoomUserRelRepo
	roomRepo                      RoomRepo
	userInfoRepo                  UserInfoRepo
	playRoomRelRepo               PlayRoomRelRepo
	playGameScoreUserRelRepo      PlayGameScoreUserRelRepo
	playGameTeamSortUserRelRepo   PlayGameTeamSortUserRelRepo
	playGameTeamGoalUserRelRepo   PlayGameTeamGoalUserRelRepo
	playGameTeamResultUserRelRepo PlayGameTeamResultUserRelRepo
	userBalanceRepo               UserBalanceRepo
	userProxyRepo                 UserProxyRepo
	tx                            Transaction
	log                           *log.Helper
}

func NewPlayUseCase(repo PlayRepo,
	playGameRelRepo PlayGameRelRepo,
	systemConfigRepo SystemConfigRepo,
	playSortRelRepo PlaySortRelRepo,
	playRoomRelRepo PlayRoomRelRepo,
	roomUserRelRepo RoomUserRelRepo,
	userInfoRepo UserInfoRepo,
	roomRepo RoomRepo,
	gameRepo GameRepo,
	sortRepo SortRepo,
	playGameScoreUserRelRepo PlayGameScoreUserRelRepo,
	playGameTeamSortUserRelRepo PlayGameTeamSortUserRelRepo,
	playGameTeamGoalUserRelRepo PlayGameTeamGoalUserRelRepo,
	playGameTeamResultUserRelRepo PlayGameTeamResultUserRelRepo,
	userBalanceRepo UserBalanceRepo,
	userProxyRepo UserProxyRepo,
	tx Transaction,
	logger log.Logger) *PlayUseCase {
	return &PlayUseCase{
		playRepo:                      repo,
		roomRepo:                      roomRepo,
		sortRepo:                      sortRepo,
		gameRepo:                      gameRepo,
		userInfoRepo:                  userInfoRepo,
		playGameRelRepo:               playGameRelRepo,
		systemConfigRepo:              systemConfigRepo,
		playSortRelRepo:               playSortRelRepo,
		playRoomRelRepo:               playRoomRelRepo,
		roomUserRelRepo:               roomUserRelRepo,
		playGameScoreUserRelRepo:      playGameScoreUserRelRepo,
		playGameTeamSortUserRelRepo:   playGameTeamSortUserRelRepo,
		playGameTeamGoalUserRelRepo:   playGameTeamGoalUserRelRepo,
		playGameTeamResultUserRelRepo: playGameTeamResultUserRelRepo,
		userBalanceRepo:               userBalanceRepo,
		userProxyRepo:                 userProxyRepo,
		tx:                            tx,
		log:                           log.NewHelper(logger)}
}

// GetAdminCreateGameAndSortPlayList ???????????????????????????????????????????????????????????????
func (p *PlayUseCase) GetAdminCreateGameAndSortPlayList(ctx context.Context, gameId int64, sortIds ...int64) (*v1.AllowedPlayListReply, error) {
	var (
		playIds            []int64 // todo ????????????????????????????????????????????????????????????????????????????????????????????????
		plays              []*Play
		adminCreatePlayIds []int64
		playGameRel        []*PlayGameRel
		playSortRel        []*PlaySortRel
		err                error
	)

	plays, err = p.playRepo.GetAdminCreatePlayList(ctx) // ??????admin???????????????
	for _, v := range plays {
		playIds = append(playIds, v.ID)
	}

	playGameRel, _ = p.playGameRelRepo.GetPlayGameRelByGameIdAndPlayIds(ctx, gameId, playIds...)
	if err != nil {
		return nil, err
	}
	for _, v := range playGameRel {
		adminCreatePlayIds = append(adminCreatePlayIds, v.PlayId)
	}

	playSortRel, err = p.playSortRelRepo.GetPlaySortRelByPlayIds(ctx, playIds...) // ???????????????????????????
	if err != nil {
		return nil, err
	}
	for _, v := range playSortRel {
		adminCreatePlayIds = append(adminCreatePlayIds, v.PlayId)
	}

	plays, err = p.playRepo.GetAdminCreatePlayListByIds(ctx, adminCreatePlayIds...) // ??????admin???????????????
	if err != nil {
		return nil, err
	}
	rep := &v1.AllowedPlayListReply{
		Items: make([]*v1.AllowedPlayListReply_Item, 0),
	}
	for _, v := range plays {
		rep.Items = append(rep.Items, &v1.AllowedPlayListReply_Item{
			ID:        v.ID,
			Type:      v.Type,
			StartTime: v.StartTime.Format("2006-01-02 15:04:05"),
			EndTime:   v.EndTime.Format("2006-01-02 15:04:05"),
		})
	}

	return rep, nil
}

// GetAdminCreateGameAndSortPlayUserList .
func (p *PlayUseCase) GetAdminCreateGameAndSortPlayUserList(ctx context.Context, gameId int64, sortIds ...int64) (*v1.GameUserListReply, error) {
	var (
		playIds                   []int64 // todo ????????????????????????????????????????????????????????????????????????????????????????????????
		adminCreatePlayIds        []int64
		plays                     []*Play
		playGameRel               []*PlayGameRel
		playSortRel               []*PlaySortRel
		playGameScoreUserRel      []*PlayGameScoreUserRel
		playGameTeamResultUserRel []*PlayGameTeamResultUserRel
		playGameTeamGoalUserRel   []*PlayGameTeamGoalUserRel
		playGameTeamSortUserRel   []*PlayGameTeamSortUserRel
		users                     []*User
		err                       error
	)

	plays, err = p.playRepo.GetAdminCreatePlayList(ctx) // ??????admin???????????????
	for _, v := range plays {
		playIds = append(playIds, v.ID)
	}

	playGameRel, _ = p.playGameRelRepo.GetPlayGameRelByGameIdAndPlayIds(ctx, gameId, playIds...)
	if err != nil {
		return nil, err
	}
	for _, v := range playGameRel {
		adminCreatePlayIds = append(adminCreatePlayIds, v.PlayId)
	}

	playSortRel, err = p.playSortRelRepo.GetPlaySortRelByPlayIds(ctx, playIds...) // ???????????????????????????
	if err != nil {
		return nil, err
	}
	for _, v := range playSortRel {
		adminCreatePlayIds = append(adminCreatePlayIds, v.PlayId)
	}

	rep := &v1.GameUserListReply{
		Items: []*v1.GameUserListReply_Item{},
	}

	playGameScoreUserRel, err = p.playGameScoreUserRelRepo.GetPlayGameScoreUserRelByPlayIds(ctx, adminCreatePlayIds...)
	if nil == err {
		var scoreUserIds []int64
		for _, v := range playGameScoreUserRel {
			scoreUserIds = append(scoreUserIds, v.UserId)
		}

		// ??????in??????
		users, err = p.playRepo.GetUserByUserIds(ctx, scoreUserIds...)
		for _, v := range users {
			rep.Items = append(rep.Items, &v1.GameUserListReply_Item{Address: v.Address})
		}
	}

	playGameTeamResultUserRel, err = p.playGameTeamResultUserRelRepo.GetPlayGameTeamResultUserRelByPlayIds(ctx, adminCreatePlayIds...)
	if nil == err {
		var resultUserIds []int64
		for _, v := range playGameTeamResultUserRel {
			resultUserIds = append(resultUserIds, v.UserId)
		}
		// ??????in??????
		users, err = p.playRepo.GetUserByUserIds(ctx, resultUserIds...)
		for _, v := range users {
			rep.Items = append(rep.Items, &v1.GameUserListReply_Item{Address: v.Address})
		}

	}

	playGameTeamGoalUserRel, err = p.playGameTeamGoalUserRelRepo.GetPlayGameTeamGoalUserRelByPlayIds(ctx, adminCreatePlayIds...)
	if nil == err {
		var goalUserIds []int64
		for _, v := range playGameTeamGoalUserRel {
			goalUserIds = append(goalUserIds, v.UserId)
		}
		// ??????in??????
		users, err = p.playRepo.GetUserByUserIds(ctx, goalUserIds...)
		for _, v := range users {
			rep.Items = append(rep.Items, &v1.GameUserListReply_Item{Address: v.Address})
		}
	}

	playGameTeamSortUserRel, err = p.playGameTeamSortUserRelRepo.GetPlayGameTeamScoreUserRelByPlayIds(ctx, adminCreatePlayIds...)
	if nil == err {
		var sortUserIds []int64
		for _, v := range playGameTeamSortUserRel {
			sortUserIds = append(sortUserIds, v.UserId)
		}
		// ??????in??????
		users, err = p.playRepo.GetUserByUserIds(ctx, sortUserIds...)
		for _, v := range users {
			rep.Items = append(rep.Items, &v1.GameUserListReply_Item{Address: v.Address})
		}
	}

	return rep, nil
}

// GetRoomUserList .
func (p *PlayUseCase) GetRoomUserList(ctx context.Context) (*v1.GetRoomUserListReply, error) {
	var (
		roomIds     []int64
		room        []*Room
		userId      int64
		roomUserRel []*RoomUserRel
		err         error
	)

	userId, _, err = getUserFromJwt(ctx) // ????????????id
	if nil != err {
		return nil, err
	}

	roomUserRel, _ = p.roomUserRelRepo.GetRoomByUserId(ctx, userId)
	for _, v := range roomUserRel {
		roomIds = append(roomIds, v.RoomId)
	}

	room, _ = p.roomRepo.GetRoomByIds(ctx, roomIds...)

	rep := &v1.GetRoomUserListReply{
		Items: make([]*v1.GetRoomUserListReply_Item, 0),
	}

	for _, v := range room {
		rep.Items = append(rep.Items, &v1.GetRoomUserListReply_Item{
			Account:   v.Account,
			CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return rep, nil
}

// GetRoomGameAndSortPlayList ????????????????????????????????????????????????????????????
func (p *PlayUseCase) GetRoomGameAndSortPlayList(ctx context.Context, roomId int64) (*v1.RoomPlayListReply, error) {
	var (
		playIds     []int64
		plays       []*Play
		playRoomRel []*PlayRoomRel
		err         error
	)
	playRoomRel, err = p.playRoomRelRepo.GetPlayRoomRelByRoomId(ctx, roomId)
	if err != nil {
		return nil, err
	}
	for _, v := range playRoomRel {
		playIds = append(playIds, v.PlayId)
	}

	plays, err = p.playRepo.GetPlayListByIds(ctx, playIds...) // ??????admin???????????????
	rep := &v1.RoomPlayListReply{
		Items: make([]*v1.RoomPlayListReply_Item, 0),
	}
	for _, v := range plays {
		rep.Items = append(rep.Items, &v1.RoomPlayListReply_Item{
			ID:        v.ID,
			Type:      v.Type,
			StartTime: v.StartTime.Format("2006-01-02 15:04:05"),
			EndTime:   v.EndTime.Format("2006-01-02 15:04:05"),
		})
	}

	return rep, nil
}

// GetUserPlayList ?????????????????????????????????
func (p *PlayUseCase) GetUserPlayList(ctx context.Context) (*v1.GetUserPlayListReply, error) {
	var (
		playGameScoreUserRel        []*PlayGameScoreUserRel
		playGameTeamGoalUserRel     []*PlayGameTeamGoalUserRel
		playGameTeamResultUserRel   []*PlayGameTeamResultUserRel
		playGameTeamSortUserRel     []*PlayGameTeamSortUserRel
		playAllTypeUserRel          PlayAllTypeUserRelSlice
		userBalanceRecordGoalReward map[int64]*UserBalanceRecord
		playGameRel                 map[int64]*PlayGameRel
		base                        int64 = 100000 // ????????????0.00001 todo ???????????????
		userId                      int64
		playIds                     []int64
		gameIds                     []int64
		play                        map[int64]*Play
		game                        map[int64]*Game
		err                         error
	)

	userId, _, err = getUserFromJwt(ctx) // ????????????id
	if nil != err {
		return nil, err
	}

	var (
		balanceRecordIdRelScore    map[int64]*BalanceRecordIdRel
		recordIds                  []int64
		tmpPlayGameScoreUserRelIds []int64
	)
	playGameScoreUserRel, err = p.playGameScoreUserRelRepo.GetPlayGameScoreUserRelByUserId(ctx, userId) // ??????admin???????????????
	for _, v := range playGameScoreUserRel {
		playAllTypeUserRel = append(playAllTypeUserRel, &PlayAllTypeUserRel{
			ID:        v.ID,
			PlayId:    v.PlayId,
			Pay:       v.OriginPay,
			Status:    v.Status,
			CreatedAt: v.CreatedAt,
			Content:   v.Content,
		})
		tmpPlayGameScoreUserRelIds = append(tmpPlayGameScoreUserRelIds, v.ID)
	}
	balanceRecordIdRelScore, err = p.userBalanceRepo.GetBalanceRecordIdRelMap(ctx, "game_score", tmpPlayGameScoreUserRelIds...)
	for _, v := range balanceRecordIdRelScore {
		recordIds = append(recordIds, v.RecordId)
	}

	var (
		balanceRecordIdRelResult    map[int64]*BalanceRecordIdRel
		tmpPlayGameResultUserRelIds []int64
	)
	playGameTeamResultUserRel, err = p.playGameTeamResultUserRelRepo.GetPlayGameTeamResultUserRelByUserId(ctx, userId) // ??????admin???????????????
	for _, v := range playGameTeamResultUserRel {
		playAllTypeUserRel = append(playAllTypeUserRel, &PlayAllTypeUserRel{
			ID:        v.ID,
			PlayId:    v.PlayId,
			Pay:       v.OriginPay,
			Status:    v.Status,
			CreatedAt: v.CreatedAt,
			Content:   v.Content,
		})
		tmpPlayGameResultUserRelIds = append(tmpPlayGameResultUserRelIds, v.ID)
	}
	balanceRecordIdRelResult, err = p.userBalanceRepo.GetBalanceRecordIdRelMap(ctx, "game_team_result", tmpPlayGameResultUserRelIds...)
	for _, v := range balanceRecordIdRelResult {
		recordIds = append(recordIds, v.RecordId)
	}

	var (
		balanceRecordIdRelGoal     map[int64]*BalanceRecordIdRel
		balanceRecordIdRelGoalUp   map[int64]*BalanceRecordIdRel
		balanceRecordIdRelGoalDown map[int64]*BalanceRecordIdRel
		tmpPlayGameGoalUserRelIds  []int64
	)
	playGameTeamGoalUserRel, err = p.playGameTeamGoalUserRelRepo.GetPlayGameTeamGoalUserRelByUserId(ctx, userId) // ??????admin???????????????
	for _, v := range playGameTeamGoalUserRel {
		playAllTypeUserRel = append(playAllTypeUserRel, &PlayAllTypeUserRel{
			ID:        v.ID,
			PlayId:    v.PlayId,
			Pay:       v.OriginPay,
			Status:    v.Status,
			CreatedAt: v.CreatedAt,
			Goal:      v.Goal,
			TeamId:    v.TeamId,
		})
		tmpPlayGameGoalUserRelIds = append(tmpPlayGameGoalUserRelIds, v.ID)
	}
	balanceRecordIdRelGoal, err = p.userBalanceRepo.GetBalanceRecordIdRelMap(ctx, "game_team_goal_all", tmpPlayGameGoalUserRelIds...)
	for _, v := range balanceRecordIdRelGoal {
		recordIds = append(recordIds, v.RecordId)
	}
	balanceRecordIdRelGoalUp, err = p.userBalanceRepo.GetBalanceRecordIdRelMap(ctx, "game_team_goal_up", tmpPlayGameGoalUserRelIds...)
	for _, v := range balanceRecordIdRelGoalUp {
		recordIds = append(recordIds, v.RecordId)
	}
	balanceRecordIdRelGoalDown, err = p.userBalanceRepo.GetBalanceRecordIdRelMap(ctx, "game_team_goal_down", tmpPlayGameGoalUserRelIds...)
	for _, v := range balanceRecordIdRelGoalDown {
		recordIds = append(recordIds, v.RecordId)
	}

	var (
		balanceRecordIdRelSort        map[int64]*BalanceRecordIdRel
		balanceRecordIdRelSortEight   map[int64]*BalanceRecordIdRel
		balanceRecordIdRelSortSixteen map[int64]*BalanceRecordIdRel
		tmpPlayGameSortUserRelIds     []int64
	)
	playGameTeamSortUserRel, err = p.playGameTeamSortUserRelRepo.GetPlayGameTeamSortUserRelByUserId(ctx, userId) // ??????admin???????????????
	for _, v := range playGameTeamSortUserRel {
		playAllTypeUserRel = append(playAllTypeUserRel, &PlayAllTypeUserRel{
			ID:        v.ID,
			PlayId:    v.PlayId,
			Pay:       v.OriginPay,
			Status:    v.Status,
			CreatedAt: v.CreatedAt,
			SortId:    v.SortId,
			Content:   v.Content,
		})

		tmpPlayGameSortUserRelIds = append(tmpPlayGameSortUserRelIds, v.ID)
	}
	balanceRecordIdRelSort, err = p.userBalanceRepo.GetBalanceRecordIdRelMap(ctx, "team_sort_three", tmpPlayGameSortUserRelIds...)
	for _, v := range balanceRecordIdRelSort {
		recordIds = append(recordIds, v.RecordId)
	}
	balanceRecordIdRelSortEight, err = p.userBalanceRepo.GetBalanceRecordIdRelMap(ctx, "team_sort_eight", tmpPlayGameSortUserRelIds...)
	for _, v := range balanceRecordIdRelSortEight {
		recordIds = append(recordIds, v.RecordId)
	}
	balanceRecordIdRelSortSixteen, err = p.userBalanceRepo.GetBalanceRecordIdRelMap(ctx, "team_sort_sixteen", tmpPlayGameSortUserRelIds...)
	for _, v := range balanceRecordIdRelSortSixteen {
		recordIds = append(recordIds, v.RecordId)
	}

	userBalanceRecordGoalReward, _ = p.userBalanceRepo.GetUserBalanceRecordGoalReward(ctx, recordIds...)
	sort.Sort(playAllTypeUserRel)

	for _, v := range playAllTypeUserRel {
		playIds = append(playIds, v.PlayId)
	}

	play, _ = p.playRepo.GetPlayMapByIds(ctx, playIds...)

	playGameRel, _ = p.playGameRelRepo.GetPlayGameRelMapByPlayId(ctx, playIds...)
	for _, v := range playGameRel {
		gameIds = append(gameIds, v.GameId)
	}
	game, _ = p.playRepo.GetGameMapByIds(ctx, gameIds...)

	rep := &v1.GetUserPlayListReply{
		Items: make([]*v1.GetUserPlayListReply_Item, 0),
	}

	var (
		gameName   string
		RedTeamId  int64
		BlueTeamId int64
		playType   string
		ok         bool
	)
	for _, v := range playAllTypeUserRel {
		var tmpPlayGameRel *PlayGameRel
		var tmpGame *Game
		var tmpPlay *Play

		if tmpPlay, ok = play[v.PlayId]; ok {
			playType = tmpPlay.Type
		}
		if tmpPlayGameRel, ok = playGameRel[v.PlayId]; ok {
			if tmpGame, ok = game[tmpPlayGameRel.GameId]; ok {
				gameName = tmpGame.Name
				RedTeamId = tmpGame.RedTeamId
				BlueTeamId = tmpGame.BlueTeamId
			}
		}

		tmpAmount := int64(0)
		if "game_score" == playType {
			if _, ok = balanceRecordIdRelScore[v.ID]; ok {
				if _, ok = userBalanceRecordGoalReward[balanceRecordIdRelScore[v.ID].RecordId]; ok {
					tmpAmount = userBalanceRecordGoalReward[balanceRecordIdRelScore[v.ID].RecordId].Amount
				}
			}
		} else if "game_team_result" == playType {
			if _, ok = balanceRecordIdRelResult[v.ID]; ok {
				if _, ok = userBalanceRecordGoalReward[balanceRecordIdRelResult[v.ID].RecordId]; ok {
					tmpAmount = userBalanceRecordGoalReward[balanceRecordIdRelResult[v.ID].RecordId].Amount
				}
			}
		} else if "game_team_goal_all" == playType || "game_team_goal_down" == playType || "game_team_goal_up" == playType {
			if _, ok = balanceRecordIdRelGoal[v.ID]; ok {
				if _, ok = userBalanceRecordGoalReward[balanceRecordIdRelGoal[v.ID].RecordId]; ok {
					tmpAmount = userBalanceRecordGoalReward[balanceRecordIdRelGoal[v.ID].RecordId].Amount
				}
			}
		} else if "game_team_goal_up" == playType {
			if _, ok = balanceRecordIdRelGoalUp[v.ID]; ok {
				if _, ok = userBalanceRecordGoalReward[balanceRecordIdRelGoalUp[v.ID].RecordId]; ok {
					tmpAmount = userBalanceRecordGoalReward[balanceRecordIdRelGoalUp[v.ID].RecordId].Amount
				}
			}
		} else if "game_team_goal_down" == playType {
			if _, ok = balanceRecordIdRelGoalDown[v.ID]; ok {
				if _, ok = userBalanceRecordGoalReward[balanceRecordIdRelGoalDown[v.ID].RecordId]; ok {
					tmpAmount = userBalanceRecordGoalReward[balanceRecordIdRelGoalDown[v.ID].RecordId].Amount
				}
			}
		} else if "team_sort_three" == playType {
			if _, ok = balanceRecordIdRelSort[v.ID]; ok {
				if _, ok = userBalanceRecordGoalReward[balanceRecordIdRelSort[v.ID].RecordId]; ok {
					tmpAmount = userBalanceRecordGoalReward[balanceRecordIdRelSort[v.ID].RecordId].Amount
				}
			}
		} else if "team_sort_eight" == playType {
			if _, ok = balanceRecordIdRelSortEight[v.ID]; ok {
				if _, ok = userBalanceRecordGoalReward[balanceRecordIdRelSortEight[v.ID].RecordId]; ok {
					tmpAmount = userBalanceRecordGoalReward[balanceRecordIdRelSortEight[v.ID].RecordId].Amount
				}
			}
		} else if "team_sort_sixteen" == playType {
			if _, ok = balanceRecordIdRelSortSixteen[v.ID]; ok {
				if _, ok = userBalanceRecordGoalReward[balanceRecordIdRelSortSixteen[v.ID].RecordId]; ok {
					tmpAmount = userBalanceRecordGoalReward[balanceRecordIdRelSortSixteen[v.ID].RecordId].Amount
				}
			}
		}

		rep.Items = append(rep.Items, &v1.GetUserPlayListReply_Item{
			Id:         v.ID,
			GameName:   gameName,
			Status:     v.Status,
			Pay:        fmt.Sprintf("%.2f", float64(v.Pay)/float64(base)), // ?????????????????????????????????????????????
			PlayId:     v.PlayId,
			CreatedAt:  v.CreatedAt.Format("2006-01-02 15:04:05"),
			RedTeamId:  RedTeamId,
			BlueTeamId: BlueTeamId,
			Type:       playType,
			Content:    v.Content,
			Goal:       v.Goal,
			TeamId:     v.TeamId,
			SortId:     v.SortId,
			Amount:     fmt.Sprintf("%.2f", float64(tmpAmount)/float64(base)),
		})
	}

	return rep, nil
}

// CreatePlayGame ?????????????????????????????????
func (r *RoomUseCase) CreatePlayGame(ctx context.Context, req *v1.CreatePlayGameRequest) (*v1.CreatePlayGameReply, error) {
	var (
		userId          int64
		userType        string
		room            *Room
		playRoomRel     *PlayRoomRel
		playRoomRelList []*PlayRoomRel
		playGameRel     *PlayGameRel
		play            *Play
		plays           []*Play
		playIds         []int64
		game            *Game
		err             error
		startTime       time.Time
		endTime         time.Time
	)

	game, err = r.gameRepo.GetGameById(ctx, req.SendBody.GameId) // ??????????????????????????????????????????
	if nil != err {
		return nil, err
	}

	startTime, err = time.Parse("2006-01-02 15:04:05", req.SendBody.StartTime) // ????????????????????????
	if nil != err {
		return nil, err
	}
	endTime, err = time.Parse("2006-01-02 15:04:05", req.SendBody.EndTime) // ????????????????????????
	if nil != err {
		return nil, err
	}
	if endTime.Before(startTime) || endTime.After(game.EndTime) {
		return nil, errors.New(500, "TIME_ERROR", "????????????????????????????????????????????????????????????")
	}

	room, err = r.roomRepo.GetRoomByID(ctx, req.SendBody.RoomId) // ??????????????? todo ??????
	if nil != err {
		return nil, err
	}

	if "game_team_goal_all" != req.SendBody.PlayType && // ??????type??????
		"game_score" != req.SendBody.PlayType &&
		"game_team_result" != req.SendBody.PlayType &&
		"game_team_goal_up" != req.SendBody.PlayType &&
		"game_team_goal_down" != req.SendBody.PlayType {
		return nil, errors.New(500, "PLAY_TYPE_ERROR", "????????????????????????")
	}

	playRoomRelList, err = r.playRoomRelRepo.GetPlayRoomRelByRoomId(ctx, req.SendBody.RoomId)
	for _, v := range playRoomRelList {
		playIds = append(playIds, v.PlayId)
	}
	plays, err = r.playRepo.GetPlayListByIds(ctx, playIds...)
	for _, v := range plays {
		if v.Type == req.SendBody.PlayType {
			return nil, errors.New(500, "ALREADY_PLAY_TYPE", "???????????????????????????")
		}
	}

	userId, userType, err = getUserFromJwt(ctx) // ????????????id
	if nil != err {
		return nil, err
	}
	if "user" != userType && "admin" != userType {
		return nil, errors.New(500, "USER_ERROR", "??????????????????")
	}

	err = r.tx.ExecTx(ctx, func(ctx context.Context) error { // ??????
		play, err = r.playRepo.CreatePlay(ctx, &Play{ // ????????????
			CreateUserId:   userId,
			CreateUserType: userType,
			Type:           req.SendBody.PlayType, // todo ???????????????????????????
			StartTime:      startTime,
			EndTime:        endTime,
		})
		if err != nil {
			return err
		}

		playRoomRel, err = r.playRoomRelRepo.CreatePlayRoomRel(ctx, &PlayRoomRel{ // ???????????????????????????
			PlayId: play.ID,
			RoomId: room.ID,
		})
		if err != nil {
			return err
		}

		playGameRel, err = r.playGameRelRepo.CreatePlayGameRel(ctx, &PlayGameRel{ // ???????????????????????????
			PlayId: play.ID,
			GameId: game.ID,
		})
		if err != nil {
			return err
		}

		return nil
	})

	return &v1.CreatePlayGameReply{
		PlayId: play.ID,
	}, err
}

// CreatePlaySort  ?????????????????????????????????
func (r *RoomUseCase) CreatePlaySort(ctx context.Context, req *v1.CreatePlaySortRequest) (*v1.CreatePlaySortReply, error) {
	var (
		userId          int64
		userType        string
		room            *Room
		playRoomRel     *PlayRoomRel
		playSortRel     *PlaySortRel
		playRoomRelList []*PlayRoomRel
		plays           []*Play
		playIds         []int64
		play            *Play
		playSort        *Sort
		err             error
		startTime       time.Time
		endTime         time.Time
	)

	playSort, err = r.sortRepo.GetGameSortById(ctx, req.SendBody.SortId) // ????????????????????????????????????????????????
	if nil != err {
		return nil, err
	}

	startTime, err = time.Parse("2006-01-02 15:04:05", req.SendBody.StartTime) // ????????????????????????
	if nil != err {
		return nil, err
	}
	endTime, err = time.Parse("2006-01-02 15:04:05", req.SendBody.EndTime)
	if nil != err {
		return nil, err
	}
	if endTime.Before(startTime) || endTime.After(playSort.EndTime) {
		return nil, errors.New(500, "TIME_ERROR", "??????????????????????????????????????????????????????????????????")
	}

	room, err = r.roomRepo.GetRoomByID(ctx, req.SendBody.RoomId) // ??????????????? todo ??????
	if nil != err {
		return nil, err
	}

	playRoomRelList, err = r.playRoomRelRepo.GetPlayRoomRelByRoomId(ctx, req.SendBody.RoomId)
	for _, v := range playRoomRelList {
		playIds = append(playIds, v.PlayId)
	}
	plays, err = r.playRepo.GetPlayListByIds(ctx, playIds...)
	for _, v := range plays {
		if v.Type == req.SendBody.PlayType {
			return nil, errors.New(500, "ALREADY_PLAY_TYPE", "???????????????????????????")
		}
	}

	userId, userType, err = getUserFromJwt(ctx) // ????????????id
	if nil != err {
		return nil, err
	}
	if "user" != userType && "admin" != userType {
		return nil, errors.New(500, "USER_ERROR", "??????????????????")
	}

	if err = r.tx.ExecTx(ctx, func(ctx context.Context) error { // ??????
		play, err = r.playRepo.CreatePlay(ctx, &Play{ // ????????????
			CreateUserId:   userId,
			CreateUserType: userType,
			Type:           playSort.Type,
			StartTime:      startTime,
			EndTime:        endTime,
		})
		if err != nil {
			return err
		}

		playRoomRel, err = r.playRoomRelRepo.CreatePlayRoomRel(ctx, &PlayRoomRel{ // ???????????????????????????
			PlayId: play.ID,
			RoomId: room.ID,
		})
		if err != nil {
			return err
		}

		playSortRel, err = r.playSortRelRepo.CreatePlaySortRel(ctx, &PlaySortRel{ // ???????????????????????????
			PlayId: play.ID,
			SortId: playSort.ID,
		})
		if err != nil {
			return err
		}

		return nil
	}); nil != err {
		return nil, err
	}

	return &v1.CreatePlaySortReply{
		PlayId: play.ID,
	}, err
}

func (p *PlayUseCase) CreatePlayGameScore(ctx context.Context, req *v1.CreatePlayGameScoreRequest) (*v1.CreatePlayGameScoreReply, error) {

	var (
		userId               int64
		playGameScoreUserRel *PlayGameScoreUserRel
		play                 *Play
		pay                  int64
		userBalance          *UserBalance
		err                  error
		recordId             int64
		originPay            int64
		base                 int64 = 100000 // ????????????0.00001
		payLimit             int64 = 100    // ??????
	)
	// todo ???????????????????????????????????????
	play, err = p.playRepo.GetPlayById(ctx, req.SendBody.PlayId) // ?????????
	if nil != err {
		return nil, err
	}
	if "game_score" != play.Type {
		return nil, errors.New(500, "PLAY_ERROR", "?????????????????????")
	}

	if play.EndTime.Before(time.Now().UTC().Add(8 * time.Hour)) {
		return nil, errors.New(500, "TIME_ERROR", "???????????????")
	}

	pay = req.SendBody.Pay * 100       // ???????????????????????????100???????????????*100
	if 0 != pay%payLimit || pay <= 0 { // ??????????????????
		return nil, errors.New(500, "PAY_ERROR", "??????????????????100")
	}

	userId, _, err = getUserFromJwt(ctx) // ????????????id
	if nil != err {
		return nil, errors.New(500, "USER_ERROR", "??????????????????")
	}
	userBalance, err = p.userBalanceRepo.GetUserBalance(ctx, userId) // ?????????
	if nil != err {
		return nil, err
	}
	pay = pay * base // todo ??????????????????????????????????????????????????????
	if pay > userBalance.Balance {
		return nil, errors.New(500, "USER_BALANCE_ERROR", "????????????")
	}
	originPay = pay

	/* todo ?????????
	 * ????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????,
	 * mysql??????innodb????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????0????????????
	 * ??????????????????update???????????????????????????????????????????????????????????????????????????
	 */
	if err = p.tx.ExecTx(ctx, func(ctx context.Context) error { // ??????
		recordId, err = p.userBalanceRepo.Pay(ctx, userId, pay) // ?????????????????????????????????????????????????????????????????????
		if err != nil {
			return err
		}

		fee := pay * 6 / 100 // ???????????????
		pay -= fee
		playGameScoreUserRel, err = p.playGameScoreUserRelRepo.CreatePlayGameScoreUserRel(ctx, &PlayGameScoreUserRel{
			ID:        0,
			UserId:    userId,
			PlayId:    play.ID,
			Content:   strconv.FormatInt(req.SendBody.RedScore, 10) + ":" + strconv.FormatInt(req.SendBody.BlueScore, 10),
			Pay:       pay,
			OriginPay: originPay,
			Status:    "default",
		})
		if err != nil {
			return err
		}
		err = p.userBalanceRepo.CreateBalanceRecordIdRel(ctx, recordId, play.Type, playGameScoreUserRel.ID)
		if err != nil {
			return err
		}

		return nil
	}); nil != err {
		return nil, err
	}

	return &v1.CreatePlayGameScoreReply{PlayId: playGameScoreUserRel.PlayId}, nil
}

func (p *PlayUseCase) CreatePlayGameResult(ctx context.Context, req *v1.CreatePlayGameResultRequest) (*v1.CreatePlayGameResultReply, error) {

	var (
		userId                    int64
		playGameTeamResultUserRel *PlayGameTeamResultUserRel
		play                      *Play
		pay                       int64
		gameResult                string
		userBalance               *UserBalance
		recordId                  int64
		originPay                 int64
		err                       error
		base                      int64 = 100000 // ????????????0.00001 todo ???????????????
		payLimit                  int64 = 100    // ?????? todo ??????????????????
	)

	if strings.EqualFold("red", req.SendBody.Result) {
		gameResult = "red"
	} else if strings.EqualFold("blue", req.SendBody.Result) {
		gameResult = "blue"
	} else if strings.EqualFold("draw", req.SendBody.Result) {
		gameResult = "draw"
	} else {
		return nil, errors.New(500, "RESULT_ERROR", "?????????????????????")
	}

	userId, _, err = getUserFromJwt(ctx) // ????????????id
	if nil != err {
		return nil, errors.New(500, "USER_ERROR", "??????????????????")
	}

	play, err = p.playRepo.GetPlayById(ctx, req.SendBody.PlayId) // ?????????
	if nil != err {
		return nil, err
	}
	if "game_team_result" != play.Type {
		return nil, errors.New(500, "PLAY_ERROR", "?????????????????????")
	}

	if play.EndTime.Before(time.Now().UTC().Add(8 * time.Hour)) {
		return nil, errors.New(500, "TIME_ERROR", "???????????????")
	}

	pay = req.SendBody.Pay * 100       // ???????????????????????????100???????????????*100
	if 0 != pay%payLimit || pay <= 0 { // ??????????????????
		return nil, errors.New(500, "PAY_ERROR", "??????????????????100")
	}

	userBalance, err = p.userBalanceRepo.GetUserBalance(ctx, userId) // ?????????
	if nil != err {
		return nil, err
	}
	pay = pay * base // todo ??????????????????????????????????????????????????????
	if pay > userBalance.Balance {
		return nil, errors.New(500, "USER_BALANCE_ERROR", "????????????")
	}
	originPay = pay

	/* todo ?????????
	 * ????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????,
	 * mysql??????innodb????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????0????????????
	 * ??????????????????update???????????????????????????????????????????????????????????????????????????
	 */
	if err = p.tx.ExecTx(ctx, func(ctx context.Context) error { // ??????
		recordId, err = p.userBalanceRepo.Pay(ctx, userId, pay) // ?????????????????????????????????????????????????????????????????????
		if err != nil {
			return err
		}

		fee := pay * 6 / 100 // ???????????????
		pay -= fee
		playGameTeamResultUserRel, err = p.playGameTeamResultUserRelRepo.CreatePlayGameTeamResultUserRel(ctx, &PlayGameTeamResultUserRel{
			ID:        0,
			UserId:    userId,
			PlayId:    play.ID,
			Content:   gameResult,
			OriginPay: originPay,
			Pay:       pay,
			Status:    "default",
		})
		if err != nil {
			return err
		}

		err = p.userBalanceRepo.CreateBalanceRecordIdRel(ctx, recordId, play.Type, playGameTeamResultUserRel.ID)
		if err != nil {
			return err
		}

		return nil
	}); nil != err {
		return nil, err
	}

	return &v1.CreatePlayGameResultReply{PlayId: playGameTeamResultUserRel.PlayId}, nil
}

func (p *PlayUseCase) CreatePlayGameSort(ctx context.Context, req *v1.CreatePlayGameSortRequest) (*v1.CreatePlayGameSortReply, error) {

	var (
		userId                  int64
		playGameTeamSortUserRel *PlayGameTeamSortUserRel
		play                    *Play
		pay                     int64
		userBalance             *UserBalance
		err                     error
		recordId                int64
		originPay               int64
		base                    int64 = 100000 // ????????????0.00001 todo ???????????????
		payLimit                int64 = 100    // ?????? todo ??????????????????
	)

	// todo ???????????????????????????????????????
	play, err = p.playRepo.GetPlayById(ctx, req.SendBody.PlayId) // ?????????
	if nil != err {
		return nil, err
	}

	if "team_sort_eight" != play.Type && "team_sort_three" != play.Type && "team_sort_sixteen" != play.Type {
		return nil, errors.New(500, "PLAY_ERROR", "?????????????????????")
	}

	if "team_sort_three" == play.Type {
		tmpPlay, tmpErr := p.playRepo.GetAdminCreatePlayByType(ctx, "team_sort_eight")
		if nil != tmpErr {
			return nil, tmpErr
		}

		if tmpPlay.EndTime.Before(time.Now().UTC().Add(8 * time.Hour)) {
			return nil, errors.New(500, "TIME_ERROR", "??????????????????????????????????????????")
		}
	}

	if play.EndTime.Before(time.Now().UTC().Add(8 * time.Hour)) {
		return nil, errors.New(500, "TIME_ERROR", "???????????????")
	}

	pay = req.SendBody.Pay * 100       // ???????????????????????????100???????????????*100
	if 0 != pay%payLimit || pay <= 0 { // ??????????????????
		return nil, errors.New(500, "PAY_ERROR", "??????????????????100")
	}

	userId, _, err = getUserFromJwt(ctx) // ????????????id
	if nil != err {
		return nil, errors.New(500, "USER_ERROR", "??????????????????")
	}
	userBalance, err = p.userBalanceRepo.GetUserBalance(ctx, userId) // ?????????
	if nil != err {
		return nil, err
	}
	pay = pay * base // todo ??????????????????????????????????????????????????????
	if pay > userBalance.Balance {
		return nil, errors.New(500, "USER_BALANCE_ERROR", "????????????")
	}
	originPay = pay

	/* todo ?????????
	 * ????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????,
	 * mysql??????innodb????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????0????????????
	 * ??????????????????update???????????????????????????????????????????????????????????????????????????
	 */
	if err = p.tx.ExecTx(ctx, func(ctx context.Context) error { // ??????
		recordId, err = p.userBalanceRepo.Pay(ctx, userId, pay) // ?????????????????????????????????????????????????????????????????????
		if err != nil {
			return err
		}

		fee := pay * 6 / 100 // ???????????????
		pay -= fee
		playGameTeamSortUserRel, err = p.playGameTeamSortUserRelRepo.CreatePlayGameTeamSortUserRel(ctx, &PlayGameTeamSortUserRel{
			ID:        0,
			UserId:    userId,
			PlayId:    play.ID,
			SortId:    req.SendBody.SortId,
			Content:   req.SendBody.Content,
			Pay:       pay,
			OriginPay: originPay,
			Status:    "default",
		})
		if err != nil {
			return err
		}

		err = p.userBalanceRepo.CreateBalanceRecordIdRel(ctx, recordId, play.Type, playGameTeamSortUserRel.ID)
		if err != nil {
			return err
		}

		return nil
	}); nil != err {
		return nil, err
	}

	return &v1.CreatePlayGameSortReply{PlayId: playGameTeamSortUserRel.PlayId}, nil
}

func (p *PlayUseCase) CreatePlayGameGoal(ctx context.Context, req *v1.CreatePlayGameGoalRequest) (*v1.CreatePlayGameGoalReply, error) {

	var (
		userId                  int64
		playGameTeamGoalUserRel *PlayGameTeamGoalUserRel
		play                    *Play
		pay                     int64
		userBalance             *UserBalance
		err                     error
		recordId                int64
		originPay               int64
		base                    int64 = 100000 // ????????????0.00001 todo ???????????????
		payLimit                int64 = 100    // ?????? todo ??????????????????
	)

	play, err = p.playRepo.GetPlayById(ctx, req.SendBody.PlayId) // ?????????
	if nil != err {
		return nil, err
	}

	if play.EndTime.Before(time.Now().UTC().Add(8 * time.Hour)) {
		return nil, errors.New(500, "TIME_ERROR", "???????????????")
	}

	if "game_team_goal_all" != req.SendBody.PlayType && "game_team_goal_up" != req.SendBody.PlayType && "game_team_goal_down" != req.SendBody.PlayType {
		return nil, errors.New(500, "PLAY_ERROR", "?????????????????????")
	}

	pay = req.SendBody.Pay * 100       // ???????????????????????????100???????????????*100
	if 0 != pay%payLimit || pay <= 0 { // ??????????????????
		return nil, errors.New(500, "PAY_ERROR", "??????????????????100")
	}

	userId, _, err = getUserFromJwt(ctx) // ????????????id
	if nil != err {
		return nil, errors.New(500, "USER_ERROR", "??????????????????")
	}
	userBalance, err = p.userBalanceRepo.GetUserBalance(ctx, userId) // ?????????
	if nil != err {
		return nil, err
	}
	pay = pay * base // todo ??????????????????????????????????????????????????????
	if pay > userBalance.Balance {
		return nil, errors.New(500, "USER_BALANCE_ERROR", "????????????")
	}
	originPay = pay

	/* todo ?????????
	 * ????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????,
	 * mysql??????innodb????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????0????????????
	 * ??????????????????update???????????????????????????????????????????????????????????????????????????
	 */
	if err = p.tx.ExecTx(ctx, func(ctx context.Context) error { // ??????
		recordId, err = p.userBalanceRepo.Pay(ctx, userId, pay) // ?????????????????????????????????????????????????????????????????????
		if err != nil {
			return err
		}

		fee := pay * 6 / 100 // ???????????????
		pay -= fee
		playGameTeamGoalUserRel, err = p.playGameTeamGoalUserRelRepo.CreatePlayGameTeamGoalUserRel(ctx, &PlayGameTeamGoalUserRel{
			ID:        0,
			UserId:    userId,
			PlayId:    play.ID,
			TeamId:    req.SendBody.TeamId,
			Type:      req.SendBody.PlayType,
			Goal:      req.SendBody.Goal,
			Pay:       pay,
			OriginPay: originPay,
			Status:    "default",
		})
		if err != nil {
			return err
		}

		err = p.userBalanceRepo.CreateBalanceRecordIdRel(ctx, recordId, play.Type, playGameTeamGoalUserRel.ID)
		if err != nil {
			return err
		}

		return nil
	}); nil != err {
		return nil, err
	}

	return &v1.CreatePlayGameGoalReply{PlayId: playGameTeamGoalUserRel.PlayId}, nil
}

func (p *PlayUseCase) PlayAmountTotal(ctx context.Context, req *v1.PlayAmountTotalRequest) (*v1.PlayAmountTotalReply, error) {

	var (
		play                           *Play
		playGameTeamGoalUserRelTotal   *PlayGameTeamGoalUserRelTotal
		playGameScoreUserRelTotal      *PlayGameScoreUserRelTotal
		playGameTeamSortUserRelTotal   *PlayGameTeamSortUserRelTotal
		playGameTeamResultUserRelTotal *PlayGameTeamResultUserRelTotal
		base                           int64 = 100000 // ????????????0.00001 todo ???????????????
		total                          int64
		err                            error
	)

	play, err = p.playRepo.GetPlayById(ctx, req.PlayId)
	if nil != err {
		return nil, err
	}

	if "game_team_goal_all" == play.Type || "game_team_goal_up" == play.Type || "game_team_goal_down" == play.Type {
		playGameTeamGoalUserRelTotal, err = p.playGameTeamGoalUserRelRepo.GetPlayGameTeamGoalUserRelByPlayIdTotal(ctx, play.ID)
		if nil != err {
			return nil, err
		}
		total = playGameTeamGoalUserRelTotal.Total
	} else if "game_score" == play.Type {
		playGameScoreUserRelTotal, err = p.playGameScoreUserRelRepo.GetPlayGameScoreUserRelByPlayIdTotal(ctx, play.ID)
		if nil != err {
			return nil, err
		}
		total = playGameScoreUserRelTotal.Total
	} else if "game_team_result" == play.Type {
		playGameTeamResultUserRelTotal, err = p.playGameTeamResultUserRelRepo.GetPlayGameTeamResultUserRelByPlayIdTotal(ctx, play.ID)
		if nil != err {
			return nil, err
		}
		total = playGameTeamResultUserRelTotal.Total
	} else if "team_sort_eight" == play.Type || "team_sort_three" == play.Type || "team_sort_sixteen" == play.Type {
		playGameTeamSortUserRelTotal, err = p.playGameTeamSortUserRelRepo.GetPlayGameTeamSortUserRelByPlayIdTotal(ctx, play.ID)
		if nil != err {
			return nil, err
		}
		total = playGameTeamSortUserRelTotal.Total
	}

	return &v1.PlayAmountTotalReply{TotalAmount: total / base}, err
}

func (p *PlayUseCase) PlayAmountTotalResult(ctx context.Context, req *v1.PlayAmountTotalResultRequest) (*v1.PlayAmountTotalResultReply, error) {
	var (
		play                      *Play
		playGameTeamResultUserRel []*PlayGameTeamResultUserRel
		base                      int64 = 100000 // ????????????0.00001 todo ???????????????
		total                     int64
		redTotal                  int64
		blueTotal                 int64
		drawTotal                 int64
		err                       error
		pool                      *LastTermPool
		poolTotal                 int64
	)

	play, err = p.playRepo.GetPlayById(ctx, req.PlayId)
	if nil != err {
		return nil, err
	}

	playGameTeamResultUserRel, _ = p.playGameTeamResultUserRelRepo.GetPlayGameTeamResultUserRelByPlayId(ctx, play.ID)

	for _, v := range playGameTeamResultUserRel {
		total += v.Pay
		if "red" == v.Content {
			redTotal += v.Pay
		}

		if "blue" == v.Content {
			blueTotal += v.Pay
		}

		if "draw" == v.Content {
			drawTotal += v.Pay
		}
	}

	pool, err = p.playRepo.GetLastTermPoolByPlayIdAndType(ctx, play.ID, play.Type)
	if nil != pool {
		poolTotal = pool.Total
	}

	return &v1.PlayAmountTotalResultReply{
		TotalAmount: fmt.Sprintf("%.2f", float64(total+poolTotal)/float64(base)),
		RedTotal:    fmt.Sprintf("%.2f", float64(redTotal)/float64(base)),
		DrawTotal:   fmt.Sprintf("%.2f", float64(drawTotal)/float64(base)),
		BlueTotal:   fmt.Sprintf("%.2f", float64(blueTotal)/float64(base)),
	}, nil

}

func (p *PlayUseCase) PlayAmountTotalScore(ctx context.Context, req *v1.PlayAmountTotalScoreRequest) (*v1.PlayAmountTotalScoreReply, error) {
	var (
		play                 *Play
		playGameScoreUserRel []*PlayGameScoreUserRel
		totalRes             map[string]int64
		base                 int64 = 100000 // ????????????0.00001 todo ???????????????
		total                int64
		err                  error
		pool                 *LastTermPool
		poolTotal            int64
	)

	play, err = p.playRepo.GetPlayById(ctx, req.PlayId)
	if nil != err {
		return nil, err
	}

	playGameScoreUserRel, _ = p.playGameScoreUserRelRepo.GetPlayGameScoreUserRelByPlayId(ctx, play.ID)
	totalRes = make(map[string]int64, 0)
	for _, v := range playGameScoreUserRel {
		total += v.Pay
		if _, ok := totalRes[v.Content]; ok {
			totalRes[v.Content] += v.Pay
		} else {
			totalRes[v.Content] = v.Pay
		}

	}

	pool, err = p.playRepo.GetLastTermPoolByPlayIdAndType(ctx, play.ID, play.Type)
	if nil != pool {
		poolTotal = pool.Total
	}

	res := &v1.PlayAmountTotalScoreReply{
		Total: fmt.Sprintf("%.2f", float64(total+poolTotal)/float64(base)),
		Items: nil,
	}
	res.Items = make([]*v1.PlayAmountTotalScoreReply_Item, 0)
	for k, v := range totalRes {
		res.Items = append(res.Items, &v1.PlayAmountTotalScoreReply_Item{
			Content: k,
			Total:   fmt.Sprintf("%.2f", float64(v)/float64(base)),
		})
	}
	return res, nil

}

func (p *PlayUseCase) PlayAmountTotalSort(ctx context.Context, req *v1.PlayAmountTotalSortRequest) (*v1.PlayAmountTotalSortReply, error) {
	var (
		play                    *Play
		playGameTeamSortUserRel []*PlayGameTeamSortUserRel
		totalFirstRes           map[string]int64
		totalSecondRes          map[string]int64
		totalThirdRes           map[string]int64
		base                    int64 = 100000 // ????????????0.00001 todo ???????????????
		total                   int64
		err                     error
		pool                    *LastTermPool
		poolTotal               int64
	)

	play, err = p.playRepo.GetPlayById(ctx, req.PlayId)
	if nil != err {
		return nil, err
	}

	playGameTeamSortUserRel, _ = p.playGameTeamSortUserRelRepo.GetPlayTeamSortUserRelByPlayId(ctx, play.ID)

	totalFirstRes = make(map[string]int64, 0)
	totalSecondRes = make(map[string]int64, 0)
	totalThirdRes = make(map[string]int64, 0)
	for _, v := range playGameTeamSortUserRel {
		total += v.Pay
		tmpContentSlice := strings.Split(v.Content, ":") // ??????
		for k, value := range tmpContentSlice {          //???????????????id
			if 0 < len(value) {
				if 0 == k {
					if _, ok := totalFirstRes[value]; ok {
						totalFirstRes[value] += v.Pay
					} else {
						totalFirstRes[value] = v.Pay
					}
				} else if 1 == k {
					if _, ok := totalSecondRes[value]; ok {
						totalSecondRes[value] += v.Pay
					} else {
						totalSecondRes[value] = v.Pay
					}
				} else if 2 == k {
					if _, ok := totalThirdRes[value]; ok {
						totalThirdRes[value] += v.Pay
					} else {
						totalThirdRes[value] = v.Pay
					}
				}
			}
		}
	}

	pool, err = p.playRepo.GetLastTermPoolByPlayIdAndType(ctx, play.ID, play.Type)
	if nil != pool {
		poolTotal = pool.Total
	}

	res := &v1.PlayAmountTotalSortReply{
		Total:       fmt.Sprintf("%.2f", float64(total+poolTotal)/float64(base)),
		FirstItems:  nil,
		SecondItems: nil,
		ThirdItems:  nil,
	}
	res.FirstItems = make([]*v1.PlayAmountTotalSortReply_First, 0)
	for k, v := range totalFirstRes {
		res.FirstItems = append(res.FirstItems, &v1.PlayAmountTotalSortReply_First{
			Content: k,
			Total:   fmt.Sprintf("%.2f", float64(v)/float64(base)),
		})
	}
	res.SecondItems = make([]*v1.PlayAmountTotalSortReply_Second, 0)
	for k, v := range totalSecondRes {
		res.SecondItems = append(res.SecondItems, &v1.PlayAmountTotalSortReply_Second{
			Content: k,
			Total:   fmt.Sprintf("%.2f", float64(v)/float64(base)),
		})
	}
	res.ThirdItems = make([]*v1.PlayAmountTotalSortReply_Third, 0)
	for k, v := range totalThirdRes {
		res.ThirdItems = append(res.ThirdItems, &v1.PlayAmountTotalSortReply_Third{
			Content: k,
			Total:   fmt.Sprintf("%.2f", float64(v)/float64(base)),
		})
	}
	return res, nil

}

func (p *PlayUseCase) PlayAmountTotalSortOther(ctx context.Context, req *v1.PlayAmountTotalSortOtherRequest) (*v1.PlayAmountTotalSortOtherReply, error) {
	var (
		play                    *Play
		playGameTeamSortUserRel []*PlayGameTeamSortUserRel
		totalRes                map[string]int64
		base                    int64 = 100000 // ????????????0.00001 todo ???????????????
		total                   int64
		err                     error
		pool                    *LastTermPool
		poolTotal               int64
	)

	play, err = p.playRepo.GetPlayById(ctx, req.PlayId)
	if nil != err {
		return nil, err
	}

	playGameTeamSortUserRel, _ = p.playGameTeamSortUserRelRepo.GetPlayTeamSortUserRelByPlayId(ctx, play.ID)
	totalRes = make(map[string]int64, 0)
	for _, v := range playGameTeamSortUserRel {
		total += v.Pay
		if _, ok := totalRes[v.Content]; ok {
			totalRes[v.Content] += v.Pay
		} else {
			totalRes[v.Content] = v.Pay
		}
	}

	pool, err = p.playRepo.GetLastTermPoolByPlayIdAndType(ctx, play.ID, play.Type)
	if nil != pool {
		poolTotal = pool.Total
	}

	res := &v1.PlayAmountTotalSortOtherReply{
		Total: fmt.Sprintf("%.2f", float64(total+poolTotal)/float64(base)),
		Items: nil,
	}
	res.Items = make([]*v1.PlayAmountTotalSortOtherReply_Item, 0)
	for k, v := range totalRes {
		res.Items = append(res.Items, &v1.PlayAmountTotalSortOtherReply_Item{
			Content: k,
			Total:   fmt.Sprintf("%.2f", float64(v)/float64(base)),
		})
	}
	return res, nil

}

func (p *PlayUseCase) PlayAmountTotalGoal(ctx context.Context, req *v1.PlayAmountTotalGoalRequest) (*v1.PlayAmountTotalGoalReply, error) {
	var (
		play                    *Play
		playGameTeamGoalUserRel []*PlayGameTeamGoalUserRel
		playGameRel             *PlayGameRel
		game                    *Game
		redTotalRes             map[int64]int64
		blueTotalRes            map[int64]int64
		base                    int64 = 100000 // ????????????0.00001 todo ???????????????
		total                   int64
		err                     error
		poolTotal               int64
		pool                    *LastTermPool
	)

	play, err = p.playRepo.GetPlayById(ctx, req.PlayId)
	if nil != err {
		return nil, err
	}
	playGameRel, err = p.playGameRelRepo.GetPlayGameRelByPlayId(ctx, req.PlayId)
	if nil != err {
		return nil, err
	}
	game, err = p.gameRepo.GetGameById(ctx, playGameRel.GameId)
	if nil != err {
		return nil, err
	}

	redTotalRes = make(map[int64]int64, 0)
	blueTotalRes = make(map[int64]int64, 0)
	playGameTeamGoalUserRel, _ = p.playGameTeamGoalUserRelRepo.GetPlayGameGoalUserRelByPlayId(ctx, play.ID)
	for _, v := range playGameTeamGoalUserRel {
		total += v.Pay
		if v.TeamId == game.RedTeamId {
			if _, ok := redTotalRes[v.Goal]; ok {
				redTotalRes[v.Goal] += v.Pay
			} else {
				redTotalRes[v.Goal] = v.Pay
			}
		} else if v.TeamId == game.BlueTeamId {
			if _, ok := blueTotalRes[v.Goal]; ok {
				blueTotalRes[v.Goal] += v.Pay
			} else {
				blueTotalRes[v.Goal] = v.Pay
			}
		}
	}

	pool, err = p.playRepo.GetLastTermPoolByPlayIdAndType(ctx, play.ID, play.Type)
	if nil != pool {
		poolTotal = pool.Total
	}

	res := &v1.PlayAmountTotalGoalReply{
		Total:     fmt.Sprintf("%.2f", float64(total+poolTotal)/float64(base)),
		RedItems:  nil,
		BlueItems: nil,
	}
	res.RedItems = make([]*v1.PlayAmountTotalGoalReply_RedItem, 0)
	for k, v := range redTotalRes {
		res.RedItems = append(res.RedItems, &v1.PlayAmountTotalGoalReply_RedItem{
			Content: k,
			Total:   fmt.Sprintf("%.2f", float64(v)/float64(base)),
		})
	}
	res.BlueItems = make([]*v1.PlayAmountTotalGoalReply_BlueItem, 0)
	for k, v := range blueTotalRes {
		res.BlueItems = append(res.BlueItems, &v1.PlayAmountTotalGoalReply_BlueItem{
			Content: k,
			Total:   fmt.Sprintf("%.2f", float64(v)/float64(base)),
		})
	}
	return res, nil

}
