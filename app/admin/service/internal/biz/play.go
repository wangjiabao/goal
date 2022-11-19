package biz

import (
	"bytes"
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "goal/api/admin/service/v1"
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

type PlayGameRel struct {
	ID     int64
	PlayId int64
	GameId int64
}

type Admin struct {
	ID       int64
	Account  string
	Password string
}

type PlayRoomRel struct {
	ID     int64
	PlayId int64
	RoomId int64
}

type PlaySortRel struct {
	ID     int64
	PlayId int64
	SortId int64
}

type PlayGameScoreUserRel struct {
	ID        int64
	PlayId    int64
	UserId    int64
	Content   string
	Pay       int64
	OriginPay int64
	Status    string
}

type PlayGameTeamResultUserRel struct {
	ID        int64
	UserId    int64
	PlayId    int64
	Content   string
	Pay       int64
	OriginPay int64
	Status    string
}

type PlayGameTeamGoalUserRel struct {
	ID        int64
	UserId    int64
	PlayId    int64
	TeamId    int64
	Type      string
	Goal      int64
	Pay       int64
	OriginPay int64
	Status    string
}

type PlayGameTeamSortUserRel struct {
	ID        int64
	UserId    int64
	PlayId    int64
	Content   string
	SortId    int64
	OriginPay int64
	Pay       int64
	Status    string
}

type UserBalance struct {
	ID      int64
	UserId  int64
	Balance int64
}

type UserBalanceRecord struct {
	ID        int64
	UserId    int64
	Balance   int64
	Type      string
	Amount    int64
	Reason    string
	CreatedAt time.Time
}

type UserProxy struct {
	ID        int64
	UserId    int64
	UpUserId  int64
	Rate      int64
	CreatedAt time.Time
}

type SystemConfig struct {
	ID    int64
	Name  string
	Value int64
}

type Room struct {
	ID        int64
	Account   string
	Type      string
	CreatedAt time.Time
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

type RoomRepo interface {
	GetRoomList(ctx context.Context) ([]*Room, error)
	GetPlayRoomByPlayId(ctx context.Context, playId int64) (*PlayRoomRel, error)
}

type SystemConfigRepo interface {
	GetSystemConfigList(ctx context.Context) ([]*SystemConfig, error)
	GetSystemConfigByNames(ctx context.Context, name ...string) (map[string]*SystemConfig, error)
	GetSystemConfigByName(ctx context.Context, name string) (*SystemConfig, error)
	UpdateConfig(ctx context.Context, id int64, value int64) (bool, error)
}

type PlayRepo interface {
	GetAdmin(ctx context.Context, account string, password string) (*Admin, error)
	GetPlayListByIds(ctx context.Context, ids ...int64) ([]*Play, error)
	GetPlayById(ctx context.Context, id int64) (*Play, error)
	CreatePlay(ctx context.Context, pc *Play) (*Play, error)
	DeletePlayById(ctx context.Context, id int64) (bool, error)
	GetAdminCreatePlayByType(ctx context.Context, playType string) ([]*Play, error)
	GetAdminCreatePlay(ctx context.Context) ([]*Play, error)
	GetAdminCreatePlayBySortType(ctx context.Context, playType string) (*Play, error)
	GetLastTermPoolByPlayIdAndType(ctx context.Context, playId int64, playType string) (*LastTermPool, error)
	CreateLastTermPool(ctx context.Context, lastTermPool *LastTermPool) (*LastTermPool, error)
}

type PlayGameRelRepo interface {
	GetPlayGameRelByGameId(ctx context.Context, gameId int64) ([]*PlayGameRel, error)
	GetPlayGameRelListByGameIdAndPlayIds(ctx context.Context, gameId int64, playId ...int64) ([]*PlayGameRel, error)
	GetPlayGameRelByGameIdAndPlayIds(ctx context.Context, gameId int64, playIds ...int64) (*PlayGameRel, error)
	CreatePlayGameRel(ctx context.Context, rel *PlayGameRel) (*PlayGameRel, error)
	DeletePlayGameRelByPlayId(ctx context.Context, playId int64) (bool, error)
}

type PlayRoomRelRepo interface {
	GetPlayRoomRelByRoomId(ctx context.Context, roomId int64) ([]*PlayRoomRel, error)
}

type PlaySortRelRepo interface {
	GetPlaySortRelBySortIdAndPlayIds(ctx context.Context, sortId int64, playId ...int64) (*PlaySortRel, error)
	GetPlaySortRelListBySortIdAndPlayIds(ctx context.Context, sortId int64, playIds ...int64) ([]*PlaySortRel, error)
	GetPlaySortRelBySortId(ctx context.Context, sortId int64) ([]*PlaySortRel, error)
	CreatePlaySortRel(ctx context.Context, rel *PlaySortRel) (*PlaySortRel, error)
	DeletePlaySortRelByPlayId(ctx context.Context, playId int64) (bool, error)
}

type PlayGameScoreUserRelRepo interface {
	GetPlayGameScoreUserRelByPlayId(ctx context.Context, playId int64) ([]*PlayGameScoreUserRel, error)
	SetRewarded(ctx context.Context, id int64) error
	SetNoRewarded(ctx context.Context, id int64) error
	CreatePlayGameScoreUserRel(ctx context.Context, pr *PlayGameScoreUserRel) (*PlayGameScoreUserRel, error)
	GetPlayGameScoreUserRelByPlayIds(ctx context.Context, playIds ...int64) (map[int64][]*PlayGameScoreUserRel, error)
}

type PlayGameTeamResultUserRelRepo interface {
	GetPlayGameTeamResultUserRelByPlayIds(ctx context.Context, playIds ...int64) (map[int64][]*PlayGameTeamResultUserRel, error)
	SetRewarded(ctx context.Context, id int64) error
	SetNoRewarded(ctx context.Context, id int64) error
	CreatePlayGameTeamResultUserRel(ctx context.Context, pr *PlayGameTeamResultUserRel) (*PlayGameTeamResultUserRel, error)
	GetPlayGameTeamResultUserRelByPlayId(ctx context.Context, playId int64) ([]*PlayGameTeamResultUserRel, error)
}

type PlayGameTeamGoalUserRelRepo interface {
	GetPlayGameTeamGoalUserRelByPlayIdsAndType(ctx context.Context, playType string, playIds ...int64) (map[int64][]*PlayGameTeamGoalUserRel, error)
	SetRewarded(ctx context.Context, id int64) error
	SetNoRewarded(ctx context.Context, id int64) error
	GetPlayGameTeamGoalUserRelByPlayId(ctx context.Context, playId int64) ([]*PlayGameTeamGoalUserRel, error)
	CreatePlayGameTeamGoalUserRel(ctx context.Context, pr *PlayGameTeamGoalUserRel) (*PlayGameTeamGoalUserRel, error)
}

type PlayGameTeamSortUserRelRepo interface {
	GetPlayGameTeamSortUserRelByPlayIds(ctx context.Context, playIds ...int64) (map[int64][]*PlayGameTeamSortUserRel, error)
	SetRewarded(ctx context.Context, id int64) error
	SetNoRewarded(ctx context.Context, id int64) error
}

type UserBalanceRepo interface {
	TransferIntoUserGoalReward(ctx context.Context, userId int64, amount int64) (int64, error)
	TransferIntoUserBack(ctx context.Context, userId int64, amount int64) (int64, error)
	CreateBalanceRecordIdRel(ctx context.Context, recordId int64, relType string, id int64) error
	GetUserBalance(ctx context.Context, userId int64) (*UserBalance, error)
	GetUserBalanceRecordTotal(ctx context.Context, recordType string, today bool) (*UserBalanceRecordTotal, error)
	GetUserBalanceTotal(ctx context.Context) (*UserBalanceTotal, error)
	GetUserBalanceRecord(ctx context.Context, reason string, b *Pagination, userIds ...int64) ([]*UserBalanceRecord, error, int64)
	TransferIntoUserGoalRecommendReward(ctx context.Context, userId int64, amount int64) (int64, error)
	GetAddressEthBalanceByAddress(ctx context.Context, address string) (*AddressEthBalance, error)
	UpdateUserBalance(ctx context.Context, userId int64, amount int64) (bool, error)
	Withdraw(ctx context.Context, userId int64, amount int64) error
	Deposit(ctx context.Context, userId int64, amount int64) (*UserBalance, error)
	GetAddressEthBalance(ctx context.Context) ([]*AddressEthBalance, error)
	WithdrawList(ctx context.Context, status string, b *Pagination, userIds ...int64) ([]*UserWithdraw, error, int64)
	WithdrawById(ctx context.Context, id int64) (*UserWithdraw, error)
	GetUserByToAddress(ctx context.Context, address string) (*User, error)
	UpdateWithdraw(ctx context.Context, Id int64, status string, tx string) error
	UpdateEthBalanceByAddress(ctx context.Context, address string, balance string) (bool, error)
}

type UserProxyRepo interface {
	GetUserProxyAndDown(ctx context.Context) ([]*UserProxy, map[int64][]*UserProxy, error)
}

type UserInfoRepo interface {
	GetUserInfoByMyRecommendCode(ctx context.Context, myRecommendCode string) (*UserInfo, error)
	GetUserInfoByUserId(ctx context.Context, userId int64) (*UserInfo, error)
	GetUserInfoListByRecommendCode(ctx context.Context, recommendCode string) ([]*UserInfo, error)
}

type PlayUseCase struct {
	uRepo                         UserRepo
	systemConfigRepo              SystemConfigRepo
	roomRepo                      RoomRepo
	playRepo                      PlayRepo
	gameRepo                      GameRepo
	playGameRelRepo               PlayGameRelRepo
	playRoomRelRepo               PlayRoomRelRepo
	playSortRelRepo               PlaySortRelRepo
	playGameScoreUserRelRepo      PlayGameScoreUserRelRepo
	playGameTeamResultUserRelRepo PlayGameTeamResultUserRelRepo
	playGameTeamGoalUserRelRepo   PlayGameTeamGoalUserRelRepo
	playGameTeamSortUserRelRepo   PlayGameTeamSortUserRelRepo
	userBalanceRepo               UserBalanceRepo
	userProxyRepo                 UserProxyRepo
	userInfoRepo                  UserInfoRepo
	sortRepo                      SortRepo
	tx                            Transaction
	log                           *log.Helper
}

func NewPlayUseCase(
	uRepo UserRepo,
	repo PlayRepo,
	roomRepo RoomRepo,
	systemConfigRepo SystemConfigRepo,
	playGameRelRepo PlayGameRelRepo,
	playSortRelRepo PlaySortRelRepo,
	playRoomRelRepo PlayRoomRelRepo,
	playGameScoreUserRelRepo PlayGameScoreUserRelRepo,
	playGameTeamGoalUserRelRepo PlayGameTeamGoalUserRelRepo,
	playGameTeamResultUserRelRepo PlayGameTeamResultUserRelRepo,
	playGameTeamSortUserRelRepo PlayGameTeamSortUserRelRepo,
	gameRepo GameRepo,
	sortRepo SortRepo,
	userBalanceRepo UserBalanceRepo,
	userProxyRepo UserProxyRepo,
	userInfoRepo UserInfoRepo,
	tx Transaction,
	logger log.Logger) *PlayUseCase {
	return &PlayUseCase{
		uRepo:                         uRepo,
		roomRepo:                      roomRepo,
		systemConfigRepo:              systemConfigRepo,
		playRepo:                      repo,
		gameRepo:                      gameRepo,
		sortRepo:                      sortRepo,
		playGameRelRepo:               playGameRelRepo,
		playRoomRelRepo:               playRoomRelRepo,
		playSortRelRepo:               playSortRelRepo,
		playGameScoreUserRelRepo:      playGameScoreUserRelRepo,
		playGameTeamGoalUserRelRepo:   playGameTeamGoalUserRelRepo,
		playGameTeamResultUserRelRepo: playGameTeamResultUserRelRepo,
		playGameTeamSortUserRelRepo:   playGameTeamSortUserRelRepo,
		userBalanceRepo:               userBalanceRepo,
		userProxyRepo:                 userProxyRepo,
		userInfoRepo:                  userInfoRepo,
		tx:                            tx,
		log:                           log.NewHelper(logger),
	}
}

func (p *PlayUseCase) GamePlayGrant(ctx context.Context, req *v1.GamePlayGrantRequest) (*v1.GamePlayGrantReply, error) {
	var (
		game                                        *Game
		playGameRel                                 []*PlayGameRel
		playIds                                     []int64
		play                                        []*Play
		playGameScore, playGameResult, playGameGoal []*Play
		err                                         error
	)

	game, err = p.gameRepo.GetGameById(ctx, req.SendBody.GameId)
	if err != nil {
		return nil, err
	}

	playGameRel, err = p.playGameRelRepo.GetPlayGameRelByGameId(ctx, game.ID)
	if err != nil {
		return nil, err
	}
	for _, v := range playGameRel {
		playIds = append(playIds, v.PlayId)
	}

	play, err = p.playRepo.GetPlayListByIds(ctx, playIds...) // 获取玩法
	if err != nil {
		return nil, err
	}

	for _, v := range play {
		if bytes.HasPrefix([]byte(v.Type), []byte("game_score")) { // 处理game_score系列类型
			playGameScore = append(playGameScore, v)
		} else if bytes.HasPrefix([]byte(v.Type), []byte("game_team_result")) { // 处理game_result系列类型
			playGameResult = append(playGameResult, v)
		} else if bytes.HasPrefix([]byte(v.Type), []byte("game_team_goal")) {
			playGameGoal = append(playGameGoal, v)
		}
	}

	if strings.EqualFold("end", game.Status) && game.EndTime.Before(time.Now().UTC().Add(8*time.Hour)) {
		p.grantTypeGameScore(ctx, game, playGameScore)
		p.grantTypeGameResult(ctx, game, playGameResult)
	}
	p.grantTypeGameGoal(ctx, game, playGameGoal)

	return &v1.GamePlayGrantReply{
		Result: "处理完成",
	}, nil
}

func (p *PlayUseCase) SortPlayGrant(ctx context.Context, req *v1.SortPlayGrantRequest) (*v1.SortPlayGrantReply, error) {
	var (
		playSort     *Sort
		playSortRel  []*PlaySortRel
		playIds      []int64
		playGameSort []*Play
		err          error
	)

	playSort, err = p.sortRepo.GetGameSortById(ctx, req.SendBody.SortId) // 获取排名截至日期以校验创建的玩法
	if nil != err {
		return nil, err
	}

	if !strings.EqualFold("end", playSort.Status) {
		return nil, errors.New(500, "TIME_ERROR", "比赛排名未结束")
	}

	playSortRel, err = p.playSortRelRepo.GetPlaySortRelBySortId(ctx, playSort.ID)
	if err != nil {
		return nil, err
	}
	for _, v := range playSortRel {
		playIds = append(playIds, v.PlayId)
	}

	playGameSort, err = p.playRepo.GetPlayListByIds(ctx, playIds...) // 获取玩法
	if err != nil {
		return nil, err
	}

	p.grantTypeGameSort(ctx, playSort, playGameSort)

	return &v1.SortPlayGrantReply{
		Result: "处理完成",
	}, nil

}

func (p *PlayUseCase) grantTypeGameSort(ctx context.Context, playSort *Sort, play []*Play) bool {
	var (
		playIds                 []int64
		playGameTeamSortUserRel map[int64][]*PlayGameTeamSortUserRel
		err                     error
		rate                    int64 = 80 // 猜中分比率可后台设置
		systemConfig            map[string]*SystemConfig
		ok                      bool
		goalBalanceRecordId     int64
	)

	systemConfig, err = p.systemConfigRepo.GetSystemConfigByNames(ctx, "sort_play_rate")
	if _, ok = systemConfig["sort_play_rate"]; !ok {
		return false
	}
	rate = systemConfig["sort_play_rate"].Value

	for _, v := range play {
		playIds = append(playIds, v.ID)
	}

	playGameTeamSortUserRel, err = p.playGameTeamSortUserRelRepo.GetPlayGameTeamSortUserRelByPlayIds(ctx, playIds...)
	if err != nil {
		return false
	}

	contentTeamId := make(map[int64]int)
	for k, v := range strings.Split(playSort.Content, ":") { //解析除队伍id
		if tmp, ok := strconv.ParseInt(v, 10, 64); 0 < tmp && nil == ok {
			contentTeamId[tmp] = k
		}
	}

	for playId, playUserRel := range playGameTeamSortUserRel {
		// 每一场玩法，数据都是一个玩法类型
		var (
			winNoRewardedPlayGameTeamResultUserRel []*struct {
				AmountBase int64
				Pay        int64
				UserId     int64
				Id         int64
			} // 猜中未发放奖励的用户
			poolAmount     int64 // 每个玩法的奖池
			winTotalAmount int64 // 中奖人的钱总额
			kPlay          *Play
		)

		kPlay, err = p.playRepo.GetPlayById(ctx, playId)
		if nil == play {
			continue
		}

		// 不足两人退钱
		tmpTotalUserIdMap := make(map[int64]int64, 0)
		var tmpBackedUser []*PlayGameTeamSortUserRel
		for _, v := range playUserRel {
			if 999999999 != v.UserId {
				tmpTotalUserIdMap[v.UserId] = v.UserId
			}
			if strings.EqualFold("default", v.Status) {
				tmpBackedUser = append(tmpBackedUser, v)
			}
		}
		if 2 > len(tmpTotalUserIdMap) {
			for _, v := range tmpBackedUser {
				if err = p.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
					if goalBalanceRecordId, err = p.userBalanceRepo.TransferIntoUserBack(ctx, v.UserId, v.OriginPay); nil != err {
						return err
					}
					if err = p.userBalanceRepo.CreateBalanceRecordIdRel(ctx, goalBalanceRecordId, playSort.Type, v.ID); nil != err {
						return err
					}
					if res := p.playGameTeamSortUserRelRepo.SetRewarded(ctx, v.ID); nil != res {
						return res
					}
					return nil
				}); nil != err {
					continue
				}
			}

			continue
		}

		// 解析中奖人
		for _, v := range playUserRel {
			tmpTeams := strings.Split(v.Content, ":") // 解析
			if "team_sort_three" == playSort.Type {
				amountBaseTmp := int64(0)
				for k, sv := range tmpTeams { //解析除队伍id
					tmp, _ := strconv.ParseInt(sv, 10, 64)
					if 0 >= tmp {
						continue // 不符合规范，非正式输入的
					}
					if _, ok := contentTeamId[tmp]; ok && k == contentTeamId[tmp] { // 不存在
						if 0 == k {
							amountBaseTmp += 50
							break
						} else if 1 == k {
							amountBaseTmp += 30
							break
						} else if 2 == k {
							amountBaseTmp += 20
							break
						}
					}
				}

				if 0 < amountBaseTmp {
					if 999999999 != v.UserId {
						winTotalAmount += v.Pay
					}
					if strings.EqualFold("default", v.Status) {
						winNoRewardedPlayGameTeamResultUserRel = append(winNoRewardedPlayGameTeamResultUserRel, &struct {
							AmountBase int64
							Pay        int64
							UserId     int64
							Id         int64
						}{AmountBase: amountBaseTmp, Pay: v.Pay, UserId: v.UserId, Id: v.ID})
					}
				} else {
					_ = p.playGameTeamSortUserRelRepo.SetNoRewarded(ctx, v.ID)
				}

			} else { // 非冠亚军
				num := 0
				for _, sv := range tmpTeams { //解析除队伍id
					tmp, _ := strconv.ParseInt(sv, 10, 64)
					if 0 >= tmp {
						break // 不符合规范，非正式输入的
					}
					if _, ok := contentTeamId[tmp]; !ok { // 不存在
						break
					}
					num++
				}

				if 16 == len(tmpTeams) && "team_sort_eight" == playSort.Type { // 16强或8强全部猜中并且没发奖励
					amountBaseTmp := int64(0)
					if 8 == num {
						amountBaseTmp += 100
					} else if 7 == num {
						amountBaseTmp += 70
					} else if 6 == num {
						amountBaseTmp += 10
					}

					if 0 < amountBaseTmp {
						if 999999999 != v.UserId {
							winTotalAmount += v.Pay
						}
						if strings.EqualFold("default", v.Status) {
							winNoRewardedPlayGameTeamResultUserRel = append(winNoRewardedPlayGameTeamResultUserRel, &struct {
								AmountBase int64
								Pay        int64
								UserId     int64
								Id         int64
							}{AmountBase: amountBaseTmp, Pay: v.Pay, UserId: v.UserId, Id: v.ID})

						}
					} else {
						_ = p.playGameTeamSortUserRelRepo.SetNoRewarded(ctx, v.ID)
					}

				} else if 8 == len(tmpTeams) && "team_sort_sixteen" == playSort.Type {
					amountBaseTmp := int64(0)
					if 10 == num {
						amountBaseTmp += 1
					} else if 11 == num {
						amountBaseTmp += 3
					} else if 12 == num {
						amountBaseTmp += 6
					} else if 13 == num {
						amountBaseTmp += 12
					} else if 14 == num {
						amountBaseTmp += 22
					} else if 15 == num {
						amountBaseTmp += 42
					} else if 16 == num {
						amountBaseTmp += 100
					}

					if 0 < amountBaseTmp {
						if 999999999 != v.UserId {
							winTotalAmount += v.Pay
						}
						if strings.EqualFold("default", v.Status) {
							winNoRewardedPlayGameTeamResultUserRel = append(winNoRewardedPlayGameTeamResultUserRel, &struct {
								AmountBase int64
								Pay        int64
								UserId     int64
								Id         int64
							}{AmountBase: amountBaseTmp, Pay: v.Pay, UserId: v.UserId, Id: v.ID})

						}
					} else {
						_ = p.playGameTeamSortUserRelRepo.SetNoRewarded(ctx, v.ID)
					}

				} else {
					_ = p.playGameTeamSortUserRelRepo.SetNoRewarded(ctx, v.ID)
				}

			}

			poolAmount += v.Pay
		}

		// 加入下一期奖池，注意后续再次结算再有中奖的，不能撤回加入下一期池子
		if 0 == winTotalAmount {
			var playRoomRel *PlayRoomRel

			// 开房间的退回
			var tmpRoomBackedUser []*PlayGameTeamSortUserRel
			playRoomRel, err = p.roomRepo.GetPlayRoomByPlayId(ctx, kPlay.ID)
			if nil != playRoomRel && strings.EqualFold("admin", kPlay.CreateUserType) {
				for _, v := range playUserRel {
					if 999999999 != v.UserId {
						if strings.EqualFold("default", v.Status) || strings.EqualFold("no_rewarded", v.Status) {
							tmpRoomBackedUser = append(tmpRoomBackedUser, v)
						}
					}
				}
				for _, v := range tmpRoomBackedUser {
					if err = p.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
						if goalBalanceRecordId, err = p.userBalanceRepo.TransferIntoUserBack(ctx, v.UserId, v.Pay); nil != err {
							return err
						}
						if err = p.userBalanceRepo.CreateBalanceRecordIdRel(ctx, goalBalanceRecordId, playSort.Type, v.ID); nil != err {
							return err
						}

						if res := p.playGameTeamSortUserRelRepo.SetRewarded(ctx, v.ID); nil != res {
							return res
						}
						return nil
					}); nil != err {
						continue
					}
				}
			} else {
				var (
					lastTermPool           *LastTermPool
					nextSort               *Sort
					adminCreatePlay        []*Play
					adminCreatePlaySortRel *PlaySortRel
					adminCreatePlayIds     []int64
				)
				if strings.EqualFold("team_sort_sixteen", kPlay.Type) {
					nextSort, err = p.sortRepo.GetNexGameSort(ctx, playSort.EndTime)
					if nil == nextSort || nil != err {
						return false
					}
					adminCreatePlay, err = p.playRepo.GetAdminCreatePlayByType(ctx, "team_sort_eight")
					for _, v := range adminCreatePlay {
						adminCreatePlayIds = append(adminCreatePlayIds, v.ID)
					}
					adminCreatePlaySortRel, err = p.playSortRelRepo.GetPlaySortRelBySortIdAndPlayIds(ctx, nextSort.ID, adminCreatePlayIds...)
					if nil == adminCreatePlaySortRel || nil != err {
						return false
					}
					if nil != nextSort {
						var nextTermPool *LastTermPool
						nextTermPool, err = p.playRepo.GetLastTermPoolByPlayIdAndType(ctx, adminCreatePlaySortRel.PlayId, "team_sort_eight")
						if nil == nextTermPool {
							_, err = p.playRepo.CreateLastTermPool(ctx, &LastTermPool{
								GameId:         adminCreatePlaySortRel.SortId,
								OriginGameId:   playSort.ID,
								PlayId:         adminCreatePlaySortRel.PlayId,
								OriginPlayId:   playId,
								Total:          poolAmount,
								PlayType:       "team_sort_eight",
								OriginPlayType: "team_sort_sixteen",
							})
							if nil != err {
								return false
							}
						}
					}

				} else if strings.EqualFold("team_sort_eight", kPlay.Type) {
					lastTermPool, err = p.playRepo.GetLastTermPoolByPlayIdAndType(ctx, playId, "team_sort_eight")
					if nil != lastTermPool {
						poolAmount += lastTermPool.Total
					}
					nextSort, err = p.sortRepo.GetNexGameSort(ctx, playSort.EndTime)
					if nil == nextSort || nil != err {
						return false
					}
					adminCreatePlay, err = p.playRepo.GetAdminCreatePlayByType(ctx, "team_sort_three")
					for _, v := range adminCreatePlay {
						adminCreatePlayIds = append(adminCreatePlayIds, v.ID)
					}
					adminCreatePlaySortRel, err = p.playSortRelRepo.GetPlaySortRelBySortIdAndPlayIds(ctx, nextSort.ID, adminCreatePlayIds...)
					if nil == adminCreatePlaySortRel || nil != err {
						return false
					}

					var nextTermPool *LastTermPool
					nextTermPool, err = p.playRepo.GetLastTermPoolByPlayIdAndType(ctx, adminCreatePlaySortRel.PlayId, "team_sort_three")
					if nil == nextTermPool {
						_, err = p.playRepo.CreateLastTermPool(ctx, &LastTermPool{
							GameId:         adminCreatePlaySortRel.SortId,
							OriginGameId:   playSort.ID,
							PlayId:         adminCreatePlaySortRel.PlayId,
							OriginPlayId:   playId,
							Total:          poolAmount,
							PlayType:       "team_sort_three",
							OriginPlayType: "team_sort_eight",
						})
						if nil != err {
							return false
						}
					}
				}
			}

			continue
		}

		sizeofWin := int64(len(winNoRewardedPlayGameTeamResultUserRel))
		if 0 == sizeofWin {
			continue
		}

		// 上一期的奖池
		var lastTermPool *LastTermPool
		lastTermPool, err = p.playRepo.GetLastTermPoolByPlayIdAndType(ctx, playId, kPlay.Type)
		if nil != lastTermPool {
			poolAmount += lastTermPool.Total
		}

		poolAmount = poolAmount * rate / 100
		for _, winV := range winNoRewardedPlayGameTeamResultUserRel {
			perAmount := poolAmount * winV.Pay * winV.AmountBase / 100 / winTotalAmount // 加权分的钱

			if err = p.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				if goalBalanceRecordId, err = p.userBalanceRepo.TransferIntoUserGoalReward(ctx, winV.UserId, perAmount); nil != err {
					return err
				}
				if err = p.userBalanceRepo.CreateBalanceRecordIdRel(ctx, goalBalanceRecordId, playSort.Type, winV.Id); nil != err {
					return err
				}
				if res := p.playGameTeamSortUserRelRepo.SetRewarded(ctx, winV.Id); nil != res {
					return res
				}

				return nil
			}); nil != err {
				continue
			}

		}
	}

	return true
}

func (p *PlayUseCase) grantTypeGameScore(ctx context.Context, game *Game, play []*Play) bool {
	var (
		playIds              []int64
		playGameScoreUserRel map[int64][]*PlayGameScoreUserRel
		err                  error
		goalBalanceRecordId  int64
	)
	for _, v := range play {
		playIds = append(playIds, v.ID)
	}

	playGameScoreUserRel, err = p.playGameScoreUserRelRepo.GetPlayGameScoreUserRelByPlayIds(ctx, playIds...)
	if err != nil {
		return false
	}

	for playId, playUserRel := range playGameScoreUserRel {
		// 每一场玩法，开始统计
		var (
			winNoRewardedPlayGameScoreUserRel []*PlayGameScoreUserRel // 猜中未发放奖励的用户
			poolAmount                        int64                   // 每个玩法的奖池
			winTotalAmount                    int64                   // 中奖人的钱总额
			kPlay                             *Play
		)

		kPlay, err = p.playRepo.GetPlayById(ctx, playId)
		if nil == play {
			continue
		}

		// 不足两人退钱，排除加池子
		tmpTotalUserIdMap := make(map[int64]int64, 0)
		var tmpBackedUser []*PlayGameScoreUserRel
		for _, v := range playUserRel {
			if 999999999 != v.UserId {
				tmpTotalUserIdMap[v.UserId] = v.UserId
			}
			if strings.EqualFold("default", v.Status) {
				tmpBackedUser = append(tmpBackedUser, v)
			}
		}

		if 2 > len(tmpTotalUserIdMap) {
			for _, v := range tmpBackedUser {
				if err = p.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
					if goalBalanceRecordId, err = p.userBalanceRepo.TransferIntoUserBack(ctx, v.UserId, v.OriginPay); nil != err {
						return err
					}
					if err = p.userBalanceRepo.CreateBalanceRecordIdRel(ctx, goalBalanceRecordId, "game_score", v.ID); nil != err {
						return err
					}
					if res := p.playGameScoreUserRelRepo.SetRewarded(ctx, v.ID); nil != res {
						return res
					}
					return nil
				}); nil != err {
					continue
				}
			}

			continue
		}

		for _, v := range playUserRel {
			poolAmount += v.Pay
			if strings.EqualFold(game.Result, v.Content) { // 判断是否猜中
				if 999999999 != v.UserId {
					winTotalAmount += v.Pay
				}
				if strings.EqualFold("default", v.Status) {
					winNoRewardedPlayGameScoreUserRel = append(winNoRewardedPlayGameScoreUserRel, v)
				}
			} else {
				_ = p.playGameScoreUserRelRepo.SetNoRewarded(ctx, v.ID)
			}
		}

		// 加入下一期奖池，注意后续再次结算再有中奖的，不能撤回加入下一期池子
		if 0 == winTotalAmount {
			var playRoomRel *PlayRoomRel

			// 开房间的退回
			var tmpRoomBackedUser []*PlayGameScoreUserRel
			playRoomRel, err = p.roomRepo.GetPlayRoomByPlayId(ctx, kPlay.ID)
			if nil != playRoomRel && strings.EqualFold("admin", kPlay.CreateUserType) {
				for _, v := range playUserRel {
					if 999999999 != v.UserId {
						if strings.EqualFold("default", v.Status) || strings.EqualFold("no_rewarded", v.Status) {
							tmpRoomBackedUser = append(tmpRoomBackedUser, v)
						}
					}
				}
				for _, v := range tmpRoomBackedUser {
					if err = p.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
						if goalBalanceRecordId, err = p.userBalanceRepo.TransferIntoUserBack(ctx, v.UserId, v.Pay); nil != err {
							return err
						}
						if err = p.userBalanceRepo.CreateBalanceRecordIdRel(ctx, goalBalanceRecordId, "game_sore", v.ID); nil != err {
							return err
						}

						if res := p.playGameScoreUserRelRepo.SetRewarded(ctx, v.ID); nil != res {
							return res
						}
						return nil
					}); nil != err {
						continue
					}
				}
			} else {
				var (
					lastTermPool           *LastTermPool
					nextGame               *Game
					adminCreatePlay        []*Play
					adminCreatePlayGameRel *PlayGameRel
					adminCreatePlayIds     []int64
				)
				if strings.EqualFold("game_score", kPlay.Type) {
					lastTermPool, err = p.playRepo.GetLastTermPoolByPlayIdAndType(ctx, playId, "game_score")
					if nil != lastTermPool {
						poolAmount += lastTermPool.Total
					}
					nextGame, err = p.gameRepo.GetNextGame(ctx, game.EndTime)
					if nil == nextGame || nil != err {
						return false
					}
					adminCreatePlay, err = p.playRepo.GetAdminCreatePlayByType(ctx, "game_score")
					for _, v := range adminCreatePlay {
						adminCreatePlayIds = append(adminCreatePlayIds, v.ID)
					}

					adminCreatePlayGameRel, err = p.playGameRelRepo.GetPlayGameRelByGameIdAndPlayIds(ctx, nextGame.ID, adminCreatePlayIds...)
					if nil == adminCreatePlayGameRel || nil != err {
						return false
					}
					if nil != nextGame {
						var nextTermPool *LastTermPool
						nextTermPool, err = p.playRepo.GetLastTermPoolByPlayIdAndType(ctx, adminCreatePlayGameRel.PlayId, kPlay.Type)
						if nil == nextTermPool {
							_, err = p.playRepo.CreateLastTermPool(ctx, &LastTermPool{
								GameId:         adminCreatePlayGameRel.GameId,
								OriginGameId:   game.ID,
								PlayId:         adminCreatePlayGameRel.PlayId,
								OriginPlayId:   playId,
								Total:          poolAmount,
								PlayType:       "game_score",
								OriginPlayType: "game_score",
							})
							if nil != err {
								return false
							}
						}
					}
				}
			}

			continue
		}

		sizeofWin := int64(len(winNoRewardedPlayGameScoreUserRel))
		if 0 == sizeofWin {
			continue
		}

		// 上一期的奖池
		var lastTermPool *LastTermPool
		lastTermPool, err = p.playRepo.GetLastTermPoolByPlayIdAndType(ctx, playId, kPlay.Type)
		if nil != lastTermPool {
			poolAmount += lastTermPool.Total
		}

		for _, winV := range winNoRewardedPlayGameScoreUserRel {
			perAmount := poolAmount * winV.Pay / winTotalAmount // 加权分的钱

			if err = p.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				if goalBalanceRecordId, err = p.userBalanceRepo.TransferIntoUserGoalReward(ctx, winV.UserId, perAmount); nil != err {
					return err
				}
				if err = p.userBalanceRepo.CreateBalanceRecordIdRel(ctx, goalBalanceRecordId, "game_score", winV.ID); nil != err {
					return err
				}

				if res := p.playGameScoreUserRelRepo.SetRewarded(ctx, winV.ID); nil != res {
					return res
				}

				return nil
			}); nil != err {
				continue
			}
		}
	}

	return true
}

func (p *PlayUseCase) grantTypeGameResult(ctx context.Context, game *Game, play []*Play) bool {
	var (
		playIds                   []int64
		playGameTeamResultUserRel map[int64][]*PlayGameTeamResultUserRel
		content                   string
		err                       error
		rate                      int64 = 80 // 猜中分比率可后台设置
		systemConfig              map[string]*SystemConfig
		ok                        bool
		goalBalanceRecordId       int64
	)

	for _, v := range play {
		playIds = append(playIds, v.ID)
	}

	playGameTeamResultUserRel, err = p.playGameTeamResultUserRelRepo.GetPlayGameTeamResultUserRelByPlayIds(ctx, playIds...)
	if err != nil {
		return false
	}

	if game.WinTeamId == game.RedTeamId {
		content = "red"
	} else if game.WinTeamId == game.BlueTeamId {
		content = "blue"
	} else {
		content = "draw"
	}

	systemConfig, err = p.systemConfigRepo.GetSystemConfigByNames(ctx, "result_play_rate")
	if _, ok = systemConfig["result_play_rate"]; !ok {
		return false
	}
	rate = systemConfig["result_play_rate"].Value

	for playId, playUserRel := range playGameTeamResultUserRel {
		// 每一场玩法
		var (
			winNoRewardedPlayGameTeamResultUserRel []*PlayGameTeamResultUserRel // 猜中未发放奖励的用户
			poolAmount                             int64                        // 每个玩法的奖池
			winTotalAmount                         int64                        // 中奖人的钱总额
			kPlay                                  *Play
		)

		kPlay, err = p.playRepo.GetPlayById(ctx, playId)
		if nil == play {
			continue
		}

		// 不足两人退钱
		tmpTotalUserIdMap := make(map[int64]int64, 0)
		var tmpBackedUser []*PlayGameTeamResultUserRel
		for _, v := range playUserRel {
			if 999999999 != v.UserId {
				tmpTotalUserIdMap[v.UserId] = v.UserId
			}
			if strings.EqualFold("default", v.Status) {
				tmpBackedUser = append(tmpBackedUser, v)
			}
		}
		if 2 > len(tmpTotalUserIdMap) {
			for _, v := range tmpBackedUser {
				if err = p.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
					if goalBalanceRecordId, err = p.userBalanceRepo.TransferIntoUserBack(ctx, v.UserId, v.OriginPay); nil != err {
						return err
					}
					if err = p.userBalanceRepo.CreateBalanceRecordIdRel(ctx, goalBalanceRecordId, "game_team_result", v.ID); nil != err {
						return err
					}
					if res := p.playGameTeamResultUserRelRepo.SetRewarded(ctx, v.ID); nil != res {
						return res
					}
					return nil
				}); nil != err {
					continue
				}
			}

			continue
		}

		for _, v := range playUserRel {
			if strings.EqualFold(content, v.Content) { // 判断是否猜中
				if 999999999 != v.UserId {
					winTotalAmount += v.Pay
				}
				if strings.EqualFold("default", v.Status) {
					winNoRewardedPlayGameTeamResultUserRel = append(winNoRewardedPlayGameTeamResultUserRel, v)
				}
				continue //赢钱的不加入奖池
			} else {
				_ = p.playGameTeamResultUserRelRepo.SetNoRewarded(ctx, v.ID)
			}
			poolAmount += v.Pay
		}

		// 加入下一期奖池，注意后续再次结算再有中奖的，不能撤回加入下一期池子
		if 0 == winTotalAmount {
			var playRoomRel *PlayRoomRel

			// 开房间的退回
			var tmpRoomBackedUser []*PlayGameTeamResultUserRel
			playRoomRel, err = p.roomRepo.GetPlayRoomByPlayId(ctx, kPlay.ID)
			if nil != playRoomRel && strings.EqualFold("admin", kPlay.CreateUserType) {
				for _, v := range playUserRel {
					if 999999999 != v.UserId {
						if strings.EqualFold("default", v.Status) || strings.EqualFold("no_rewarded", v.Status) {
							tmpRoomBackedUser = append(tmpRoomBackedUser, v)
						}
					}
				}
				for _, v := range tmpRoomBackedUser {
					if err = p.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
						if goalBalanceRecordId, err = p.userBalanceRepo.TransferIntoUserBack(ctx, v.UserId, v.Pay); nil != err {
							return err
						}
						if err = p.userBalanceRepo.CreateBalanceRecordIdRel(ctx, goalBalanceRecordId, "game_team_result", v.ID); nil != err {
							return err
						}

						if res := p.playGameTeamResultUserRelRepo.SetRewarded(ctx, v.ID); nil != res {
							return res
						}
						return nil
					}); nil != err {
						continue
					}
				}
			} else {
				var (
					lastTermPool           *LastTermPool
					nextGame               *Game
					adminCreatePlay        []*Play
					adminCreatePlayGameRel *PlayGameRel
					adminCreatePlayIds     []int64
				)
				if strings.EqualFold("game_team_result", kPlay.Type) {
					lastTermPool, err = p.playRepo.GetLastTermPoolByPlayIdAndType(ctx, playId, "game_team_result")
					if nil != lastTermPool {
						poolAmount += lastTermPool.Total
					}
					nextGame, err = p.gameRepo.GetNextGame(ctx, game.EndTime)
					if nil == nextGame || nil != err {
						return false
					}
					adminCreatePlay, err = p.playRepo.GetAdminCreatePlayByType(ctx, "game_team_result")
					for _, v := range adminCreatePlay {
						adminCreatePlayIds = append(adminCreatePlayIds, v.ID)
					}
					adminCreatePlayGameRel, err = p.playGameRelRepo.GetPlayGameRelByGameIdAndPlayIds(ctx, nextGame.ID, adminCreatePlayIds...)
					if nil == adminCreatePlayGameRel || nil != err {
						return false
					}
					if nil != nextGame {
						var nextTermPool *LastTermPool
						nextTermPool, err = p.playRepo.GetLastTermPoolByPlayIdAndType(ctx, adminCreatePlayGameRel.PlayId, kPlay.Type)
						if nil == nextTermPool {
							_, err = p.playRepo.CreateLastTermPool(ctx, &LastTermPool{
								GameId:         adminCreatePlayGameRel.GameId,
								OriginGameId:   game.ID,
								PlayId:         adminCreatePlayGameRel.PlayId,
								OriginPlayId:   playId,
								Total:          poolAmount,
								PlayType:       "game_team_result",
								OriginPlayType: "game_team_result",
							})
							if nil != err {
								return false
							}
						}
					}
				}
			}

			continue
		}

		sizeofWin := int64(len(winNoRewardedPlayGameTeamResultUserRel))
		if 0 == sizeofWin {
			continue
		}

		// 上一期的奖池
		var lastTermPool *LastTermPool
		lastTermPool, err = p.playRepo.GetLastTermPoolByPlayIdAndType(ctx, playId, kPlay.Type)
		if nil != lastTermPool {
			poolAmount += lastTermPool.Total
		}

		poolAmount = poolAmount * rate / 100
		for _, winV := range winNoRewardedPlayGameTeamResultUserRel {
			perAmount := poolAmount * winV.Pay / winTotalAmount // 加权分的钱

			if err = p.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				perAmount += winV.Pay // 押注的钱原路返回

				if goalBalanceRecordId, err = p.userBalanceRepo.TransferIntoUserGoalReward(ctx, winV.UserId, perAmount); nil != err {
					return err
				}
				if err = p.userBalanceRepo.CreateBalanceRecordIdRel(ctx, goalBalanceRecordId, "game_team_result", winV.ID); nil != err {
					return err
				}

				if res := p.playGameTeamResultUserRelRepo.SetRewarded(ctx, winV.ID); nil != res {
					return res
				}

				return nil
			}); nil != err {
				continue
			}

		}
	}

	return true
}

func (p *PlayUseCase) grantTypeGameGoal(ctx context.Context, game *Game, play []*Play) bool {

	var (
		playIds                 []int64
		playGameTeamGoalUserRel map[int64][]*PlayGameTeamGoalUserRel
		err                     error
		res                     bool
	)

	for _, v := range play {
		playIds = append(playIds, v.ID)
	}

	// 上半场
	if game.UpEndTime.Before(time.Now().UTC().Add(8 * time.Hour)) {
		playGameTeamGoalUserRel, err = p.playGameTeamGoalUserRelRepo.GetPlayGameTeamGoalUserRelByPlayIdsAndType(ctx, "game_team_goal_up", playIds...)
		if err != nil {
			return res
		}
		res = p.grantTypeGameGoalHandle(ctx, playGameTeamGoalUserRel, game)
		if !res {
			return res
		}
	}

	// 下半场
	if game.EndTime.Before(time.Now().UTC().Add(8 * time.Hour)) {
		playGameTeamGoalUserRel, err = p.playGameTeamGoalUserRelRepo.GetPlayGameTeamGoalUserRelByPlayIdsAndType(ctx, "game_team_goal_down", playIds...)
		if err != nil {
			return res
		}
		res = p.grantTypeGameGoalHandle(ctx, playGameTeamGoalUserRel, game)
		if !res {
			return res
		}

		// 全场
		playGameTeamGoalUserRel, err = p.playGameTeamGoalUserRelRepo.GetPlayGameTeamGoalUserRelByPlayIdsAndType(ctx, "game_team_goal_all", playIds...)
		if err != nil {
			return res
		}
		res = p.grantTypeGameGoalHandle(ctx, playGameTeamGoalUserRel, game)
		if !res {
			return res
		}
	}
	return true
}

func (p *PlayUseCase) grantTypeGameGoalHandle(ctx context.Context, playGameTeamGoalUserRel map[int64][]*PlayGameTeamGoalUserRel, game *Game) bool {
	var (
		err                 error
		goalBalanceRecordId int64
	)

	for playId, playUserRel := range playGameTeamGoalUserRel {
		// 每一场玩法
		var (
			winNoRewardedPlayGameTeamGoalUserRel []*PlayGameTeamGoalUserRel // 猜中未发放奖励的用户
			poolAmount                           int64                      // 每个玩法的奖池
			winTotalAmount                       int64                      // 中奖人的钱总额
			play                                 *Play
		)
		play, err = p.playRepo.GetPlayById(ctx, playId)
		if nil == play {
			continue
		}

		// 不足两人退钱
		tmpTotalUserIdMap := make(map[int64]int64, 0)
		var tmpBackedUser []*PlayGameTeamGoalUserRel
		for _, v := range playUserRel {
			if 999999999 != v.UserId {
				tmpTotalUserIdMap[v.UserId] = v.UserId
			}
			if strings.EqualFold("default", v.Status) {
				tmpBackedUser = append(tmpBackedUser, v)
			}
		}
		if 2 > len(tmpTotalUserIdMap) {
			for _, v := range tmpBackedUser {
				if err = p.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
					if goalBalanceRecordId, err = p.userBalanceRepo.TransferIntoUserBack(ctx, v.UserId, v.OriginPay); nil != err {
						return err
					}
					if err = p.userBalanceRepo.CreateBalanceRecordIdRel(ctx, goalBalanceRecordId, v.Type, v.ID); nil != err {
						return err
					}

					if res := p.playGameTeamGoalUserRelRepo.SetRewarded(ctx, v.ID); nil != res {
						return res
					}
					return nil
				}); nil != err {
					continue
				}
			}

			continue
		}

		for _, v := range playUserRel { // 当前玩法，全为上半场或下半场或全场
			if strings.EqualFold("game_team_goal_all", v.Type) { // 判断是否猜中
				if v.TeamId == game.RedTeamId && v.Goal == game.RedTeamUpGoal+game.RedTeamDownGoal {
					if 999999999 != v.UserId {
						winTotalAmount += v.Pay
					}
					if strings.EqualFold("default", v.Status) {
						winNoRewardedPlayGameTeamGoalUserRel = append(winNoRewardedPlayGameTeamGoalUserRel, v)
					}
				} else if v.TeamId == game.BlueTeamId && v.Goal == game.BlueTeamUpGoal+game.BlueTeamDownGoal {
					if 999999999 != v.UserId {
						winTotalAmount += v.Pay
					}
					if strings.EqualFold("default", v.Status) {
						winNoRewardedPlayGameTeamGoalUserRel = append(winNoRewardedPlayGameTeamGoalUserRel, v)
					}
				} else {
					_ = p.playGameTeamGoalUserRelRepo.SetNoRewarded(ctx, v.ID)
				}

			} else if strings.EqualFold("game_team_goal_up", v.Type) {
				if v.TeamId == game.RedTeamId && v.Goal == game.RedTeamUpGoal {
					if 999999999 != v.UserId {
						winTotalAmount += v.Pay
					}
					if strings.EqualFold("default", v.Status) {
						winNoRewardedPlayGameTeamGoalUserRel = append(winNoRewardedPlayGameTeamGoalUserRel, v)
					}
				} else if v.TeamId == game.BlueTeamId && v.Goal == game.BlueTeamUpGoal {
					if 999999999 != v.UserId {
						winTotalAmount += v.Pay
					}
					if strings.EqualFold("default", v.Status) {
						winNoRewardedPlayGameTeamGoalUserRel = append(winNoRewardedPlayGameTeamGoalUserRel, v)
					}
				} else {
					_ = p.playGameTeamGoalUserRelRepo.SetNoRewarded(ctx, v.ID)
				}

			} else if strings.EqualFold("game_team_goal_down", v.Type) {
				if v.TeamId == game.RedTeamId && v.Goal == game.RedTeamDownGoal {
					if 999999999 != v.UserId {
						winTotalAmount += v.Pay
					}
					if strings.EqualFold("default", v.Status) {
						winNoRewardedPlayGameTeamGoalUserRel = append(winNoRewardedPlayGameTeamGoalUserRel, v)
					}
				} else if v.TeamId == game.BlueTeamId && v.Goal == game.BlueTeamDownGoal {
					if 999999999 != v.UserId {
						winTotalAmount += v.Pay
					}
					if strings.EqualFold("default", v.Status) {
						winNoRewardedPlayGameTeamGoalUserRel = append(winNoRewardedPlayGameTeamGoalUserRel, v)
					}
				} else {
					_ = p.playGameTeamGoalUserRelRepo.SetNoRewarded(ctx, v.ID)
				}

			}

			poolAmount += v.Pay
		}

		// 加入下一期奖池，注意后续再次结算再有中奖的，不能撤回加入下一期池子
		if 0 == winTotalAmount {
			var playRoomRel *PlayRoomRel

			// 开房间的退回
			var tmpRoomBackedUser []*PlayGameTeamGoalUserRel
			playRoomRel, err = p.roomRepo.GetPlayRoomByPlayId(ctx, play.ID)
			if nil != playRoomRel && strings.EqualFold("admin", play.CreateUserType) {
				for _, v := range playUserRel {
					if 999999999 != v.UserId {
						if strings.EqualFold("default", v.Status) || strings.EqualFold("no_rewarded", v.Status) {
							tmpRoomBackedUser = append(tmpRoomBackedUser, v)
						}
					}
				}
				for _, v := range tmpRoomBackedUser {
					if err = p.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
						if goalBalanceRecordId, err = p.userBalanceRepo.TransferIntoUserBack(ctx, v.UserId, v.Pay); nil != err {
							return err
						}
						if err = p.userBalanceRepo.CreateBalanceRecordIdRel(ctx, goalBalanceRecordId, v.Type, v.ID); nil != err {
							return err
						}

						if res := p.playGameTeamGoalUserRelRepo.SetRewarded(ctx, v.ID); nil != res {
							return res
						}
						return nil
					}); nil != err {
						continue
					}
				}
			} else {
				var (
					lastTermPool           *LastTermPool
					nextGame               *Game
					adminCreatePlay        []*Play
					adminCreatePlayGameRel *PlayGameRel
					adminCreatePlayIds     []int64
				)
				if strings.EqualFold("game_team_goal_all", play.Type) {
					lastTermPool, err = p.playRepo.GetLastTermPoolByPlayIdAndType(ctx, playId, "game_team_goal_all")
					if nil != lastTermPool {
						poolAmount += lastTermPool.Total
					}
					nextGame, err = p.gameRepo.GetNextGame(ctx, game.EndTime)
					if nil == nextGame || nil != err {
						return false
					}
					adminCreatePlay, err = p.playRepo.GetAdminCreatePlayByType(ctx, "game_team_goal_up")
					for _, v := range adminCreatePlay {
						adminCreatePlayIds = append(adminCreatePlayIds, v.ID)
					}
					adminCreatePlayGameRel, err = p.playGameRelRepo.GetPlayGameRelByGameIdAndPlayIds(ctx, nextGame.ID, adminCreatePlayIds...)
					if nil == adminCreatePlayGameRel || nil != err {
						return false
					}
					if nil != nextGame {
						var nextTermPool *LastTermPool
						nextTermPool, err = p.playRepo.GetLastTermPoolByPlayIdAndType(ctx, adminCreatePlayGameRel.PlayId, "game_team_goal_up")
						if nil == nextTermPool {
							_, err = p.playRepo.CreateLastTermPool(ctx, &LastTermPool{
								GameId:         adminCreatePlayGameRel.GameId,
								OriginGameId:   game.ID,
								PlayId:         adminCreatePlayGameRel.PlayId,
								OriginPlayId:   playId,
								Total:          poolAmount,
								PlayType:       "game_team_goal_up",
								OriginPlayType: "game_team_goal_all",
							})
							if nil != err {
								return false
							}
						}
					}

				} else if strings.EqualFold("game_team_goal_up", play.Type) {
					lastTermPool, err = p.playRepo.GetLastTermPoolByPlayIdAndType(ctx, playId, "game_team_goal_up")
					if nil != lastTermPool {
						poolAmount += lastTermPool.Total
					}

					adminCreatePlay, err = p.playRepo.GetAdminCreatePlayByType(ctx, "game_team_goal_down")
					for _, v := range adminCreatePlay {
						adminCreatePlayIds = append(adminCreatePlayIds, v.ID)
					}
					adminCreatePlayGameRel, err = p.playGameRelRepo.GetPlayGameRelByGameIdAndPlayIds(ctx, game.ID, adminCreatePlayIds...)
					if nil == adminCreatePlayGameRel || nil != err {
						return false
					}

					var nextTermPool *LastTermPool
					nextTermPool, err = p.playRepo.GetLastTermPoolByPlayIdAndType(ctx, adminCreatePlayGameRel.PlayId, "game_team_goal_down")
					if nil == nextTermPool {
						_, err = p.playRepo.CreateLastTermPool(ctx, &LastTermPool{
							GameId:         adminCreatePlayGameRel.GameId,
							OriginGameId:   game.ID,
							PlayId:         adminCreatePlayGameRel.PlayId,
							OriginPlayId:   playId,
							Total:          poolAmount,
							PlayType:       "game_team_goal_down",
							OriginPlayType: "game_team_goal_up",
						})
						if nil != err {
							return false
						}
					}

				} else if strings.EqualFold("game_team_goal_down", play.Type) {
					lastTermPool, err = p.playRepo.GetLastTermPoolByPlayIdAndType(ctx, playId, "game_team_goal_down")
					if nil != lastTermPool {
						poolAmount += lastTermPool.Total
					}

					adminCreatePlay, err = p.playRepo.GetAdminCreatePlayByType(ctx, "game_team_goal_all")
					for _, v := range adminCreatePlay {
						adminCreatePlayIds = append(adminCreatePlayIds, v.ID)
					}
					adminCreatePlayGameRel, err = p.playGameRelRepo.GetPlayGameRelByGameIdAndPlayIds(ctx, game.ID, adminCreatePlayIds...)
					if nil == adminCreatePlayGameRel || nil != err {
						return false
					}

					var nextTermPool *LastTermPool
					nextTermPool, err = p.playRepo.GetLastTermPoolByPlayIdAndType(ctx, adminCreatePlayGameRel.PlayId, "game_team_goal_all")
					if nil == nextTermPool {
						_, err = p.playRepo.CreateLastTermPool(ctx, &LastTermPool{
							GameId:         adminCreatePlayGameRel.GameId,
							OriginGameId:   game.ID,
							PlayId:         adminCreatePlayGameRel.PlayId,
							OriginPlayId:   playId,
							Total:          poolAmount,
							PlayType:       "game_team_goal_all",
							OriginPlayType: "game_team_goal_down",
						})
						if nil != err {
							return false
						}
					}

				}
			}

			continue
		}

		sizeofWin := int64(len(winNoRewardedPlayGameTeamGoalUserRel))
		if 0 == sizeofWin {
			continue
		}

		// 上一期的奖池
		var lastTermPool *LastTermPool
		lastTermPool, err = p.playRepo.GetLastTermPoolByPlayIdAndType(ctx, playId, play.Type)
		if nil != lastTermPool {
			poolAmount += lastTermPool.Total
		}

		for _, winV := range winNoRewardedPlayGameTeamGoalUserRel {
			perAmount := poolAmount * winV.Pay / winTotalAmount // 加权分的钱

			if err = p.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
				if goalBalanceRecordId, err = p.userBalanceRepo.TransferIntoUserGoalReward(ctx, winV.UserId, perAmount); nil != err {
					return err
				}
				if err = p.userBalanceRepo.CreateBalanceRecordIdRel(ctx, goalBalanceRecordId, winV.Type, winV.ID); nil != err {
					return err
				}

				if res := p.playGameTeamGoalUserRelRepo.SetRewarded(ctx, winV.ID); nil != res {
					return res
				}

				return nil
			}); nil != err {
				continue
			}

		}
	}

	return true
}

func (p *PlayUseCase) Login(ctx context.Context, req *v1.LoginRequest) (*Admin, error) {
	return p.playRepo.GetAdmin(ctx, req.SendBody.Account, req.SendBody.Password)
}

// CreatePlayGame 创建一个比赛玩法等记录
func (p *PlayUseCase) CreatePlayGame(ctx context.Context, req *v1.CreatePlayGameRequest) (*v1.CreatePlayGameReply, error) {
	var (
		playGameRel            *PlayGameRel
		play                   *Play
		adminCreatePlay        []*Play
		adminCreatePlayIds     []int64
		adminCreatePlayGameRel *PlayGameRel
		game                   *Game
		err                    error
		startTime              time.Time
		endTime                time.Time
	)

	game, err = p.gameRepo.GetGameById(ctx, req.SendBody.GameId) // 获取比赛信息以校验创建的玩法
	if nil != err {
		return nil, err
	}

	startTime, err = time.Parse("2006-01-02 15:04:05", req.SendBody.StartTime) // 时间进行格式校验
	if nil != err {
		return nil, err
	}
	endTime, err = time.Parse("2006-01-02 15:04:05", req.SendBody.EndTime) // 时间进行格式校验
	if nil != err {
		return nil, err
	}
	if endTime.Before(startTime) || endTime.After(game.EndTime) {
		return nil, errors.New(500, "TIME_ERROR", "时间输入错误")
	}

	if "game_team_goal_all" != req.SendBody.PlayType && // 验证type类型
		"game_score" != req.SendBody.PlayType &&
		"game_team_result" != req.SendBody.PlayType &&
		"game_team_goal_up" != req.SendBody.PlayType &&
		"game_team_goal_down" != req.SendBody.PlayType {
		return nil, errors.New(500, "TIME_ERROR", "玩法类型输入错误")
	}

	adminCreatePlay, err = p.playRepo.GetAdminCreatePlayByType(ctx, req.SendBody.PlayType)
	for _, v := range adminCreatePlay {
		adminCreatePlayIds = append(adminCreatePlayIds, v.ID)
	}
	adminCreatePlayGameRel, err = p.playGameRelRepo.GetPlayGameRelByGameIdAndPlayIds(ctx, game.ID, adminCreatePlayIds...)
	if nil != adminCreatePlayGameRel {
		return nil, errors.New(500, "PLAY_ERROR", "已创建比赛下该玩法")
	}

	if err = p.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		play, err = p.playRepo.CreatePlay(ctx, &Play{ // 新增玩法
			CreateUserId:   1,
			CreateUserType: "admin",
			Type:           req.SendBody.PlayType, // todo 用户输入参数未验证
			StartTime:      startTime,
			EndTime:        endTime,
		})
		if err != nil {
			return err
		}

		playGameRel, err = p.playGameRelRepo.CreatePlayGameRel(ctx, &PlayGameRel{ // 新增比赛和玩法关系
			PlayId: play.ID,
			GameId: game.ID,
		})
		if err != nil {
			return err
		}

		return nil
	}); nil != err {
		return nil, err
	}

	return &v1.CreatePlayGameReply{
		PlayId: play.ID,
	}, err
}

func (p *PlayUseCase) DeletePlayGame(ctx context.Context, req *v1.DeletePlayGameRequest) (*v1.DeletePlayGameReply, error) {

	var err error
	if err = p.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		_, err = p.playRepo.DeletePlayById(ctx, req.SendBody.PlayId)
		if nil != err {
			return err
		}

		_, err = p.playGameRelRepo.DeletePlayGameRelByPlayId(ctx, req.SendBody.PlayId)
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		return nil, err
	}

	return &v1.DeletePlayGameReply{
		Result: "成功",
	}, nil
}

// CreatePlaySort  创建一个排名玩法等记录
func (p *PlayUseCase) CreatePlaySort(ctx context.Context, req *v1.CreatePlaySortRequest) (*v1.CreatePlaySortReply, error) {
	var (
		playSortRel *PlaySortRel
		play        *Play
		sort        *Sort
		err         error
		startTime   time.Time
		endTime     time.Time
	)

	sort, err = p.sortRepo.GetGameSortById(ctx, req.SendBody.SortId) // 获取排名截至日期以校验创建的玩法
	if nil != err {
		return nil, err
	}

	tmpPlaySort, tmpErr := p.playRepo.GetAdminCreatePlayBySortType(ctx, sort.Type)
	if nil == tmpErr || nil != tmpPlaySort {
		return nil, errors.New(500, "TIME_ERROR", "已存在玩法")
	}

	startTime, err = time.Parse("2006-01-02 15:04:05", req.SendBody.StartTime) // 时间进行格式校验
	if nil != err {
		return nil, err
	}
	endTime, err = time.Parse("2006-01-02 15:04:05", req.SendBody.EndTime)
	if nil != err {
		return nil, err
	}
	if endTime.Before(startTime) || endTime.After(sort.EndTime) {
		return nil, errors.New(500, "TIME_ERROR", "时间输入错误")
	}

	if err = p.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		play, err = p.playRepo.CreatePlay(ctx, &Play{ // 新增玩法
			CreateUserId:   1,
			CreateUserType: "admin",
			Type:           sort.Type,
			StartTime:      startTime,
			EndTime:        endTime,
		})
		if err != nil {
			return err
		}

		playSortRel, err = p.playSortRelRepo.CreatePlaySortRel(ctx, &PlaySortRel{ // 新增排名和玩法关系
			PlayId: play.ID,
			SortId: sort.ID,
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

func (p *PlayUseCase) DeletePlaySort(ctx context.Context, req *v1.DeletePlaySortRequest) (*v1.DeletePlaySortReply, error) {
	var err error
	if err = p.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		_, err = p.playRepo.DeletePlayById(ctx, req.SendBody.PlayId)
		if nil != err {
			return err
		}

		_, err = p.playSortRelRepo.DeletePlaySortRelByPlayId(ctx, req.SendBody.PlayId)
		if nil != err {
			return err
		}

		return nil
	}); nil != err {
		return nil, err
	}

	return &v1.DeletePlaySortReply{
		Result: "成功",
	}, nil
}

func (p *PlayUseCase) GetPlayList(ctx context.Context, req *v1.GetPlayListRequest) (*v1.GetPlayListReply, error) {
	var (
		play               []*Play
		playIds            []int64
		playGameRel        []*PlayGameRel
		adminCreatePlay    []*Play
		adminCreatePlayIds []int64
		err                error
	)
	adminCreatePlay, err = p.playRepo.GetAdminCreatePlay(ctx)
	for _, v := range adminCreatePlay {
		adminCreatePlayIds = append(adminCreatePlayIds, v.ID)
	}

	playGameRel, err = p.playGameRelRepo.GetPlayGameRelListByGameIdAndPlayIds(ctx, req.GameId, adminCreatePlayIds...)
	if err != nil {
		return nil, err
	}
	for _, v := range playGameRel {
		playIds = append(playIds, v.PlayId)
	}

	play, err = p.playRepo.GetPlayListByIds(ctx, playIds...)

	res := &v1.GetPlayListReply{
		Items: make([]*v1.GetPlayListReply_Item, 0),
	}
	for _, item := range play {
		res.Items = append(res.Items, &v1.GetPlayListReply_Item{
			PlayId:    item.ID,
			Type:      item.Type,
			StartTime: item.StartTime.Format("2006-01-02 15:04:05"),
			EndTime:   item.EndTime.Format("2006-01-02 15:04:05"),
		})
	}

	return res, nil
}

func (p *PlayUseCase) GetPlaySortList(ctx context.Context, req *v1.GetPlaySortListRequest) (*v1.GetPlaySortListReply, error) {
	var (
		play               []*Play
		playIds            []int64
		playSortRel        []*PlaySortRel
		adminCreatePlay    []*Play
		adminCreatePlayIds []int64
		err                error
	)

	adminCreatePlay, err = p.playRepo.GetAdminCreatePlay(ctx)
	for _, v := range adminCreatePlay {
		adminCreatePlayIds = append(adminCreatePlayIds, v.ID)
	}

	playSortRel, err = p.playSortRelRepo.GetPlaySortRelListBySortIdAndPlayIds(ctx, req.SortId, adminCreatePlayIds...)
	if err != nil {
		return nil, err
	}
	for _, v := range playSortRel {
		playIds = append(playIds, v.PlayId)
	}

	play, err = p.playRepo.GetPlayListByIds(ctx, playIds...)

	res := &v1.GetPlaySortListReply{
		Items: make([]*v1.GetPlaySortListReply_Item, 0),
	}
	for _, item := range play {
		res.Items = append(res.Items, &v1.GetPlaySortListReply_Item{
			PlayId:    item.ID,
			Type:      item.Type,
			StartTime: item.StartTime.Format("2006-01-02 15:04:05"),
			EndTime:   item.EndTime.Format("2006-01-02 15:04:05"),
		})
	}

	return res, nil
}

func (p *PlayUseCase) GetRoomPlayList(ctx context.Context, req *v1.GetRoomPlayListRequest) (*v1.GetRoomPlayListReply, error) {
	var (
		play        []*Play
		playIds     []int64
		playRoomRel []*PlayRoomRel
		err         error
	)

	playRoomRel, err = p.playRoomRelRepo.GetPlayRoomRelByRoomId(ctx, req.RoomId)
	if err != nil {
		return nil, err
	}
	for _, v := range playRoomRel {
		playIds = append(playIds, v.PlayId)
	}

	play, err = p.playRepo.GetPlayListByIds(ctx, playIds...)

	res := &v1.GetRoomPlayListReply{
		Items: make([]*v1.GetRoomPlayListReply_Item, 0),
	}
	for _, item := range play {
		res.Items = append(res.Items, &v1.GetRoomPlayListReply_Item{
			PlayId:    item.ID,
			Type:      item.Type,
			StartTime: item.StartTime.Format("2006-01-02 15:04:05"),
			EndTime:   item.EndTime.Format("2006-01-02 15:04:05"),
		})
	}

	return res, nil
}

func (p *PlayUseCase) GetPlayUserRelList(ctx context.Context, req *v1.GetPlayRelListRequest) (*v1.GetPlayRelListReply, error) {
	var (
		play        *Play
		userId      []int64
		base        int64 = 100000 // 基础精度0.00001 todo 加配置文件
		playUserRel []*struct {
			UserId int64
			Status string
			Pay    int64
		}
		user map[int64]*User
		err  error
	)

	play, err = p.playRepo.GetPlayById(ctx, req.PlayId)
	if nil != err {
		return nil, err
	}

	if "game_team_goal_all" == play.Type || "game_team_goal_up" == play.Type || "game_team_goal_down" == play.Type {
		playGameTeamGoalUserRel, _ := p.playGameTeamGoalUserRelRepo.GetPlayGameTeamGoalUserRelByPlayId(ctx, play.ID)
		for _, v := range playGameTeamGoalUserRel {
			playUserRel = append(playUserRel, &struct {
				UserId int64
				Status string
				Pay    int64
			}{UserId: v.UserId, Status: v.Status, Pay: v.Pay})
		}
	} else if "game_score" == play.Type {
		playGameScoreUserRel, _ := p.playGameScoreUserRelRepo.GetPlayGameScoreUserRelByPlayId(ctx, play.ID)
		for _, v := range playGameScoreUserRel {
			playUserRel = append(playUserRel, &struct {
				UserId int64
				Status string
				Pay    int64
			}{UserId: v.UserId, Status: v.Status, Pay: v.Pay})
		}
	} else if "game_team_result" == play.Type {
		playGameTeamResultUserRel, _ := p.playGameTeamResultUserRelRepo.GetPlayGameTeamResultUserRelByPlayId(ctx, play.ID)
		for _, v := range playGameTeamResultUserRel {
			playUserRel = append(playUserRel, &struct {
				UserId int64
				Status string
				Pay    int64
			}{UserId: v.UserId, Status: v.Status, Pay: v.Pay})
		}
	}

	for _, v := range playUserRel {
		userId = append(userId, v.UserId)
	}

	user, _ = p.uRepo.GetUserMap(ctx, userId...)

	res := &v1.GetPlayRelListReply{
		Items: make([]*v1.GetPlayRelListReply_Item, 0),
	}

	for _, item := range playUserRel {
		tempAddress := ""
		if v, ok := user[item.UserId]; ok {
			tempAddress = v.Address
		}
		res.Items = append(res.Items, &v1.GetPlayRelListReply_Item{
			Address: tempAddress,
			Pay:     item.Pay / base,
			Status:  item.Status,
		})
	}

	return res, nil
}

func (p *PlayUseCase) GetRooms(ctx context.Context) (*v1.GetRoomListReply, error) {
	var (
		room []*Room
		err  error
	)

	room, err = p.roomRepo.GetRoomList(ctx)
	if err != nil {
		return nil, err
	}

	res := &v1.GetRoomListReply{
		Items: make([]*v1.GetRoomListReply_Item, 0),
	}

	for _, item := range room {
		res.Items = append(res.Items, &v1.GetRoomListReply_Item{
			RoomId:    item.ID,
			Account:   item.Account,
			CreatedAt: item.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return res, nil
}

func (p *PlayUseCase) CreatePlayGameScore(ctx context.Context, req *v1.CreatePlayGameScoreRequest) (*v1.CreatePlayGameScoreReply, error) {

	var (
		playGameScoreUserRel *PlayGameScoreUserRel
		play                 *Play
		err                  error
		base                 int64 = 100000 // 基础精度0.00001 todo 加配置文件
	)

	play, err = p.playRepo.GetPlayById(ctx, req.SendBody.PlayId) // 查玩法
	if nil != err {
		return nil, err
	}

	playGameScoreUserRel, err = p.playGameScoreUserRelRepo.CreatePlayGameScoreUserRel(ctx, &PlayGameScoreUserRel{
		UserId:  999999999,
		PlayId:  play.ID,
		Content: "49:49",
		Pay:     req.SendBody.Pay * base,
		Status:  "rewarded",
	})
	if nil != err {
		return nil, err
	}

	return &v1.CreatePlayGameScoreReply{PlayId: playGameScoreUserRel.PlayId}, nil
}

func (p *PlayUseCase) CreatePlayGameResult(ctx context.Context, req *v1.CreatePlayGameResultRequest) (*v1.CreatePlayGameResultReply, error) {

	var (
		playGameTeamResultUserRel *PlayGameTeamResultUserRel
		play                      *Play
		err                       error
		base                      int64 = 100000 // 基础精度0.00001 todo 加配置文件
	)

	play, err = p.playRepo.GetPlayById(ctx, req.SendBody.PlayId) // 查玩法
	if nil != err {
		return nil, err
	}

	playGameTeamResultUserRel, err = p.playGameTeamResultUserRelRepo.CreatePlayGameTeamResultUserRel(ctx, &PlayGameTeamResultUserRel{
		UserId:  999999999,
		PlayId:  play.ID,
		Content: req.SendBody.Content,
		Pay:     req.SendBody.Pay * base,
		Status:  "rewarded",
	})
	if err != nil {
		return nil, err
	}

	return &v1.CreatePlayGameResultReply{PlayId: playGameTeamResultUserRel.PlayId}, nil
}

func (p *PlayUseCase) CreatePlayGameGoal(ctx context.Context, req *v1.CreatePlayGameGoalRequest) (*v1.CreatePlayGameGoalReply, error) {

	var (
		playGameTeamGoalUserRel *PlayGameTeamGoalUserRel
		play                    *Play
		game                    *Game
		err                     error
		teamId                  int64
		base                    int64 = 100000 // 基础精度0.00001 todo 加配置文件
	)

	play, err = p.playRepo.GetPlayById(ctx, req.SendBody.PlayId) // 查玩法
	if nil != err {
		return nil, err
	}

	game, err = p.gameRepo.GetGameById(ctx, req.SendBody.GameId)
	if err != nil {
		return nil, err
	}

	if "red" == req.SendBody.Team {
		teamId = game.RedTeamId
	} else {
		teamId = game.BlueTeamId
	}

	playGameTeamGoalUserRel, err = p.playGameTeamGoalUserRelRepo.CreatePlayGameTeamGoalUserRel(ctx, &PlayGameTeamGoalUserRel{
		UserId: 999999999,
		PlayId: play.ID,
		TeamId: teamId,
		Type:   play.Type,
		Goal:   req.SendBody.Goal,
		Pay:    req.SendBody.Pay * base,
		Status: "rewarded",
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreatePlayGameGoalReply{PlayId: playGameTeamGoalUserRel.PlayId}, nil
}

func (p *PlayUseCase) GetConfigList(ctx context.Context) (*v1.GetConfigListReply, error) {
	var (
		systemConfig []*SystemConfig
		err          error
	)

	systemConfig, err = p.systemConfigRepo.GetSystemConfigList(ctx)
	if err != nil {
		return nil, err
	}
	res := &v1.GetConfigListReply{
		Items: make([]*v1.GetConfigListReply_Item, 0),
	}
	for _, v := range systemConfig {
		res.Items = append(res.Items, &v1.GetConfigListReply_Item{
			Id:    v.ID,
			Name:  v.Name,
			Value: v.Value,
		})
	}

	return res, nil
}

func (p *PlayUseCase) UpdateConfig(ctx context.Context, req *v1.UpdateConfigRequest) (*v1.UpdateConfigReply, error) {
	var (
		err error
	)
	_, err = p.systemConfigRepo.UpdateConfig(ctx, req.SendBody.Id, req.SendBody.Value)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateConfigReply{Id: req.SendBody.Id}, nil
}
