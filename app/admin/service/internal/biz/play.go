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

type PlayGameRel struct {
	ID     int64
	PlayId int64
	GameId int64
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
	ID      int64
	PlayId  int64
	UserId  int64
	Content string
	Pay     int64
	Status  string
}

type PlayGameTeamResultUserRel struct {
	ID      int64
	UserId  int64
	PlayId  int64
	Content string
	Pay     int64
	Status  string
}

type PlayGameTeamGoalUserRel struct {
	ID     int64
	UserId int64
	PlayId int64
	TeamId int64
	Type   string
	Goal   int64
	Pay    int64
	Status string
}

type PlayGameTeamSortUserRel struct {
	ID      int64
	UserId  int64
	PlayId  int64
	Content string
	SortId  int64
	Pay     int64
	Status  string
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
}

type SystemConfigRepo interface {
	GetSystemConfigList(ctx context.Context) ([]*SystemConfig, error)
	UpdateConfig(ctx context.Context, id int64, value int64) (bool, error)
}

type PlayRepo interface {
	GetPlayListByIds(ctx context.Context, ids ...int64) ([]*Play, error)
	GetPlayById(ctx context.Context, id int64) (*Play, error)
	CreatePlay(ctx context.Context, pc *Play) (*Play, error)
}

type PlayGameRelRepo interface {
	GetPlayGameRelByGameId(ctx context.Context, gameId int64) ([]*PlayGameRel, error)
	CreatePlayGameRel(ctx context.Context, rel *PlayGameRel) (*PlayGameRel, error)
}

type PlayRoomRelRepo interface {
	GetPlayRoomRelByRoomId(ctx context.Context, roomId int64) ([]*PlayRoomRel, error)
}

type PlaySortRelRepo interface {
	GetPlaySortRelBySortId(ctx context.Context, sortId int64) ([]*PlaySortRel, error)
	CreatePlaySortRel(ctx context.Context, rel *PlaySortRel) (*PlaySortRel, error)
}

type PlayGameScoreUserRelRepo interface {
	GetPlayGameScoreUserRelByPlayId(ctx context.Context, playId int64) ([]*PlayGameScoreUserRel, error)
	SetRewarded(ctx context.Context, id int64) error
	CreatePlayGameScoreUserRel(ctx context.Context, pr *PlayGameScoreUserRel) (*PlayGameScoreUserRel, error)
	GetPlayGameScoreUserRelByPlayIds(ctx context.Context, playIds ...int64) (map[int64][]*PlayGameScoreUserRel, error)
}

type PlayGameTeamResultUserRelRepo interface {
	GetPlayGameTeamResultUserRelByPlayIds(ctx context.Context, playIds ...int64) (map[int64][]*PlayGameTeamResultUserRel, error)
	SetRewarded(ctx context.Context, id int64) error
	CreatePlayGameTeamResultUserRel(ctx context.Context, pr *PlayGameTeamResultUserRel) (*PlayGameTeamResultUserRel, error)
	GetPlayGameTeamResultUserRelByPlayId(ctx context.Context, playId int64) ([]*PlayGameTeamResultUserRel, error)
}

type PlayGameTeamGoalUserRelRepo interface {
	GetPlayGameTeamGoalUserRelByPlayIds(ctx context.Context, playIds ...int64) (map[int64][]*PlayGameTeamGoalUserRel, error)
	SetRewarded(ctx context.Context, id int64) error
	GetPlayGameTeamGoalUserRelByPlayId(ctx context.Context, playId int64) ([]*PlayGameTeamGoalUserRel, error)
	CreatePlayGameTeamGoalUserRel(ctx context.Context, pr *PlayGameTeamGoalUserRel) (*PlayGameTeamGoalUserRel, error)
}

type PlayGameTeamSortUserRelRepo interface {
	GetPlayGameTeamSortUserRelByPlayIds(ctx context.Context, playIds ...int64) (map[int64][]*PlayGameTeamSortUserRel, error)
	SetRewarded(ctx context.Context, id int64) error
}

type UserBalanceRepo interface {
	TransferIntoUserGoalReward(ctx context.Context, userId int64, amount int64) error
	GetUserBalance(ctx context.Context, userId int64) (*UserBalance, error)
	GetUserBalanceRecord(ctx context.Context) ([]*UserBalanceRecord, error)
	TransferIntoUserGoalRecommendReward(ctx context.Context, userId int64, amount int64) error
	GetAddressEthBalanceByAddress(ctx context.Context, address string) (*AddressEthBalance, error)
	Deposit(ctx context.Context, userId int64, amount int64) (*UserBalance, error)
	UpdateEthBalanceByAddress(ctx context.Context, address string, balance string) (bool, error)
}

type UserProxyRepo interface {
	GetUserProxyAndDown(ctx context.Context) ([]*UserProxy, map[int64][]*UserProxy, error)
}

type UserInfoRepo interface {
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

	if !strings.EqualFold("end", game.Status) {
		return nil, errors.New(500, "TIME_ERROR", "比赛未结束")
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

	p.grantTypeGameScore(ctx, game, playGameScore)
	p.grantTypeGameResult(ctx, game, playGameResult)
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
	)

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

	for _, playUserRel := range playGameTeamSortUserRel {
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
		)

		// 解析中奖人
		for _, v := range playUserRel {
			tmpTeams := strings.Split(v.Content, ":") // 解析
			if "team_sort_three" == playSort.Type {
				amountBaseTmp := int64(0)
				for k, sv := range tmpTeams { //解析除队伍id
					tmp, _ := strconv.ParseInt(sv, 10, 64)
					if 0 >= tmp {
						break // 不符合规范，非正式输入的
					}
					if _, ok := contentTeamId[tmp]; ok && k == contentTeamId[tmp] { // 不存在
						if 0 == k {
							amountBaseTmp += 50
						} else if 1 == k {
							amountBaseTmp += 30
						} else if 2 == k {
							amountBaseTmp += 20
						}
					}
				}

				if 0 < amountBaseTmp {
					if 999999999 != v.UserId {
						winTotalAmount += v.Pay
					}
					if strings.EqualFold("no_rewarded", v.Status) {
						winNoRewardedPlayGameTeamResultUserRel = append(winNoRewardedPlayGameTeamResultUserRel, &struct {
							AmountBase int64
							Pay        int64
							UserId     int64
							Id         int64
						}{AmountBase: amountBaseTmp, Pay: v.Pay, UserId: v.UserId, Id: v.ID})
					}
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

				if (16 == num && 16 == len(tmpTeams) && "team_sort_sixteen" == playSort.Type) || (8 == num && 8 == len(tmpTeams) && "team_sort_eight" == playSort.Type) { // 16强或8强全部猜中并且没发奖励
					if 999999999 != v.UserId {
						winTotalAmount += v.Pay
					}
					if strings.EqualFold("no_rewarded", v.Status) {
						winNoRewardedPlayGameTeamResultUserRel = append(winNoRewardedPlayGameTeamResultUserRel, &struct {
							AmountBase int64
							Pay        int64
							UserId     int64
							Id         int64
						}{AmountBase: 100, Pay: v.Pay, UserId: v.UserId, Id: v.ID})
					}
				}

			}

			poolAmount += v.Pay
		}

		sizeofWin := int64(len(winNoRewardedPlayGameTeamResultUserRel))
		if 0 == sizeofWin {
			continue
		}

		poolAmount = poolAmount * rate / 100
		for _, winV := range winNoRewardedPlayGameTeamResultUserRel {
			var (
				recommendUserIds []int64
				userInfo         *UserInfo
			)
			perAmount := poolAmount * winV.Pay * winV.AmountBase / 100 / winTotalAmount // 加权分的钱

			userInfo, err = p.userInfoRepo.GetUserInfoByUserId(ctx, winV.UserId) // 获取推荐关系
			for _, ruv := range strings.Split(userInfo.RecommendCode, "GA") {    // 解析userId, 取前三代
				tmp, _ := strconv.ParseInt(ruv, 10, 64)
				if 0 < tmp {
					recommendUserIds = append(recommendUserIds, tmp)
				}
			}
			userIdsLen := len(recommendUserIds)
			if userIdsLen > 3 {
				recommendUserIds = recommendUserIds[userIdsLen-3:]
			}

			if err = p.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务

				for k, recommendUserId := range recommendUserIds { // 推荐人
					var tmpPerAmount int64
					if 0 == k {
						tmpPerAmount = perAmount * 2 / 1000
						if res := p.userBalanceRepo.TransferIntoUserGoalRecommendReward(ctx, recommendUserId, tmpPerAmount); nil != res {
							return res
						}
					} else if 1 == k {
						tmpPerAmount = perAmount * 3 / 1000
						if res := p.userBalanceRepo.TransferIntoUserGoalRecommendReward(ctx, recommendUserId, tmpPerAmount); nil != res {
							return res
						}
					} else if 2 == k {
						tmpPerAmount = perAmount * 5 / 1000
						tmpPerAmount += perAmount * 10 / 1000
						if res := p.userBalanceRepo.TransferIntoUserGoalRecommendReward(ctx, recommendUserId, tmpPerAmount); nil != res {
							return res
						}
					}
					perAmount -= tmpPerAmount
				}

				if res := p.userBalanceRepo.TransferIntoUserGoalReward(ctx, winV.UserId, perAmount); nil != res {
					return res
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
	)
	for _, v := range play {
		playIds = append(playIds, v.ID)
	}

	playGameScoreUserRel, err = p.playGameScoreUserRelRepo.GetPlayGameScoreUserRelByPlayIds(ctx, playIds...)
	if err != nil {
		return false
	}

	for _, playUserRel := range playGameScoreUserRel {
		// 每一场玩法，开始统计
		var (
			winNoRewardedPlayGameScoreUserRel []*PlayGameScoreUserRel // 猜中未发放奖励的用户
			poolAmount                        int64                   // 每个玩法的奖池
			winTotalAmount                    int64                   // 中奖人的钱总额
		)

		for _, v := range playUserRel {
			poolAmount += v.Pay
			if strings.EqualFold(game.Result, v.Content) { // 判断是否猜中
				if 999999999 != v.UserId {
					winTotalAmount += v.Pay
				}
				if strings.EqualFold("no_rewarded", v.Status) {
					winNoRewardedPlayGameScoreUserRel = append(winNoRewardedPlayGameScoreUserRel, v)
				}
			}
		}

		sizeofWin := int64(len(winNoRewardedPlayGameScoreUserRel))
		if 0 == sizeofWin {
			// todo 可以原路退回
			continue
		}

		for _, winV := range winNoRewardedPlayGameScoreUserRel {
			var (
				recommendUserIds []int64
				userInfo         *UserInfo
			)
			perAmount := poolAmount * winV.Pay / winTotalAmount // 加权分的钱

			userInfo, err = p.userInfoRepo.GetUserInfoByUserId(ctx, winV.UserId) // 获取推荐关系
			for _, ruv := range strings.Split(userInfo.RecommendCode, "GA") {    // 解析userId, 取前三代
				tmp, _ := strconv.ParseInt(ruv, 10, 64)
				if 0 < tmp {
					recommendUserIds = append(recommendUserIds, tmp)
				}
			}
			userIdsLen := len(recommendUserIds)
			if userIdsLen > 3 {
				recommendUserIds = recommendUserIds[userIdsLen-3:]
			}

			if err = p.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务

				for k, recommendUserId := range recommendUserIds { // 推荐人
					var tmpPerAmount int64
					if 0 == k {
						tmpPerAmount = perAmount * 2 / 1000
						if res := p.userBalanceRepo.TransferIntoUserGoalRecommendReward(ctx, recommendUserId, tmpPerAmount); nil != res {
							return res
						}
					} else if 1 == k {
						tmpPerAmount = perAmount * 3 / 1000
						if res := p.userBalanceRepo.TransferIntoUserGoalRecommendReward(ctx, recommendUserId, tmpPerAmount); nil != res {
							return res
						}
					} else if 2 == k {
						tmpPerAmount = perAmount * 5 / 1000
						tmpPerAmount += perAmount * 10 / 1000
						if res := p.userBalanceRepo.TransferIntoUserGoalRecommendReward(ctx, recommendUserId, tmpPerAmount); nil != res {
							return res
						}
					}
					perAmount -= tmpPerAmount
				}

				if res := p.userBalanceRepo.TransferIntoUserGoalReward(ctx, winV.UserId, perAmount); nil != res {
					return res
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

	for _, playUserRel := range playGameTeamResultUserRel {
		// 每一场玩法
		var (
			winNoRewardedPlayGameTeamResultUserRel []*PlayGameTeamResultUserRel // 猜中未发放奖励的用户
			poolAmount                             int64                        // 每个玩法的奖池
			winTotalAmount                         int64                        // 中奖人的钱总额
		)
		for _, v := range playUserRel {
			if strings.EqualFold(content, v.Content) { // 判断是否猜中
				if 999999999 != v.UserId {
					winTotalAmount += v.Pay
				}
				if strings.EqualFold("no_rewarded", v.Status) {
					winNoRewardedPlayGameTeamResultUserRel = append(winNoRewardedPlayGameTeamResultUserRel, v)
				}
				continue //赢钱的不加入奖池
			}
			poolAmount += v.Pay
		}

		sizeofWin := int64(len(winNoRewardedPlayGameTeamResultUserRel))
		if 0 == sizeofWin {
			continue
		}

		poolAmount = poolAmount * rate / 100
		for _, winV := range winNoRewardedPlayGameTeamResultUserRel {
			var (
				recommendUserIds []int64
				userInfo         *UserInfo
			)
			perAmount := poolAmount * winV.Pay / winTotalAmount // 加权分的钱

			userInfo, err = p.userInfoRepo.GetUserInfoByUserId(ctx, winV.UserId) // 获取推荐关系
			for _, ruv := range strings.Split(userInfo.RecommendCode, "GA") {    // 解析userId, 取前三代
				tmp, _ := strconv.ParseInt(ruv, 10, 64)
				if 0 < tmp {
					recommendUserIds = append(recommendUserIds, tmp)
				}
			}
			userIdsLen := len(recommendUserIds)
			if userIdsLen > 3 {
				recommendUserIds = recommendUserIds[userIdsLen-3:]
			}

			if err = p.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务

				for k, recommendUserId := range recommendUserIds { // 推荐人
					var tmpPerAmount int64
					if 0 == k {
						tmpPerAmount = perAmount * 2 / 1000
						if res := p.userBalanceRepo.TransferIntoUserGoalRecommendReward(ctx, recommendUserId, tmpPerAmount); nil != res {
							return res
						}
					} else if 1 == k {
						tmpPerAmount = perAmount * 3 / 1000
						if res := p.userBalanceRepo.TransferIntoUserGoalRecommendReward(ctx, recommendUserId, tmpPerAmount); nil != res {
							return res
						}
					} else if 2 == k {
						tmpPerAmount = perAmount * 5 / 1000
						tmpPerAmount += perAmount * 10 / 1000
						if res := p.userBalanceRepo.TransferIntoUserGoalRecommendReward(ctx, recommendUserId, tmpPerAmount); nil != res {
							return res
						}
					}
					perAmount -= tmpPerAmount
				}

				perAmount += winV.Pay // 押注的钱原路返回
				if res := p.userBalanceRepo.TransferIntoUserGoalReward(ctx, winV.UserId, perAmount); nil != res {
					return res
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
		rate                    int64 = 80 // 猜中分比率可后台设置
	)

	for _, v := range play {
		playIds = append(playIds, v.ID)
	}

	playGameTeamGoalUserRel, err = p.playGameTeamGoalUserRelRepo.GetPlayGameTeamGoalUserRelByPlayIds(ctx, playIds...)
	if err != nil {
		return false
	}

	for _, playUserRel := range playGameTeamGoalUserRel {
		// 每一场玩法
		var (
			winNoRewardedPlayGameTeamGoalUserRel []*PlayGameTeamGoalUserRel // 猜中未发放奖励的用户
			poolAmount                           int64                      // 每个玩法的奖池
			winTotalAmount                       int64                      // 中奖人的钱总额
		)

		for _, v := range playUserRel { // 当前玩法，全为上半场或下半场或全场
			if strings.EqualFold("game_team_goal_all", v.Type) { // 判断是否猜中
				if v.TeamId == game.RedTeamId && v.Goal == game.RedTeamUpGoal+game.RedTeamDownGoal {
					if 999999999 != v.UserId {
						winTotalAmount += v.Pay
					}
					if strings.EqualFold("no_rewarded", v.Status) {
						winNoRewardedPlayGameTeamGoalUserRel = append(winNoRewardedPlayGameTeamGoalUserRel, v)
					}
				} else if v.TeamId == game.BlueTeamId && v.Goal == game.BlueTeamUpGoal+game.BlueTeamDownGoal {
					if 999999999 != v.UserId {
						winTotalAmount += v.Pay
					}
					if strings.EqualFold("no_rewarded", v.Status) {
						winNoRewardedPlayGameTeamGoalUserRel = append(winNoRewardedPlayGameTeamGoalUserRel, v)
					}
				}

			} else if strings.EqualFold("game_team_goal_up", v.Type) {
				if v.TeamId == game.RedTeamId && v.Goal == game.RedTeamUpGoal {
					if 999999999 != v.UserId {
						winTotalAmount += v.Pay
					}
					if strings.EqualFold("no_rewarded", v.Status) {
						winNoRewardedPlayGameTeamGoalUserRel = append(winNoRewardedPlayGameTeamGoalUserRel, v)
					}
				} else if v.TeamId == game.BlueTeamId && v.Goal == game.BlueTeamUpGoal {
					if 999999999 != v.UserId {
						winTotalAmount += v.Pay
					}
					if strings.EqualFold("no_rewarded", v.Status) {
						winNoRewardedPlayGameTeamGoalUserRel = append(winNoRewardedPlayGameTeamGoalUserRel, v)
					}
				}

			} else if strings.EqualFold("game_team_goal_down", v.Type) {
				if v.TeamId == game.RedTeamId && v.Goal == game.RedTeamDownGoal {
					if 999999999 != v.UserId {
						winTotalAmount += v.Pay
					}
					if strings.EqualFold("no_rewarded", v.Status) {
						winNoRewardedPlayGameTeamGoalUserRel = append(winNoRewardedPlayGameTeamGoalUserRel, v)
					}
				} else if v.TeamId == game.BlueTeamId && v.Goal == game.BlueTeamDownGoal {
					if 999999999 != v.UserId {
						winTotalAmount += v.Pay
					}
					if strings.EqualFold("no_rewarded", v.Status) {
						winNoRewardedPlayGameTeamGoalUserRel = append(winNoRewardedPlayGameTeamGoalUserRel, v)
					}
				}

			}

			poolAmount += v.Pay
		}

		sizeofWin := int64(len(winNoRewardedPlayGameTeamGoalUserRel))
		if 0 == sizeofWin {
			// todo 未中奖处理
			continue
		}

		poolAmount = poolAmount * rate / 100
		for _, winV := range winNoRewardedPlayGameTeamGoalUserRel {
			var (
				recommendUserIds []int64
				userInfo         *UserInfo
			)
			perAmount := poolAmount * winV.Pay / winTotalAmount // 加权分的钱

			userInfo, err = p.userInfoRepo.GetUserInfoByUserId(ctx, winV.UserId) // 获取推荐关系
			for _, ruv := range strings.Split(userInfo.RecommendCode, "GA") {    // 解析userId, 取前三代
				tmp, _ := strconv.ParseInt(ruv, 10, 64)
				if 0 < tmp {
					recommendUserIds = append(recommendUserIds, tmp)
				}
			}
			userIdsLen := len(recommendUserIds)
			if userIdsLen > 3 {
				recommendUserIds = recommendUserIds[userIdsLen-3:]
			}

			if err = p.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务

				for k, recommendUserId := range recommendUserIds { // 推荐人
					var tmpPerAmount int64
					if 0 == k {
						tmpPerAmount = perAmount * 2 / 1000
						if res := p.userBalanceRepo.TransferIntoUserGoalRecommendReward(ctx, recommendUserId, tmpPerAmount); nil != res {
							return res
						}
					} else if 1 == k {
						tmpPerAmount = perAmount * 3 / 1000
						if res := p.userBalanceRepo.TransferIntoUserGoalRecommendReward(ctx, recommendUserId, tmpPerAmount); nil != res {
							return res
						}
					} else if 2 == k {
						tmpPerAmount = perAmount * 5 / 1000
						tmpPerAmount += perAmount * 10 / 1000
						if res := p.userBalanceRepo.TransferIntoUserGoalRecommendReward(ctx, recommendUserId, tmpPerAmount); nil != res {
							return res
						}
					}
					perAmount -= tmpPerAmount
				}

				if res := p.userBalanceRepo.TransferIntoUserGoalReward(ctx, winV.UserId, perAmount); nil != res {
					return res
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

// CreatePlayGame 创建一个比赛玩法等记录
func (p *PlayUseCase) CreatePlayGame(ctx context.Context, req *v1.CreatePlayGameRequest) (*v1.CreatePlayGameReply, error) {
	var (
		playGameRel *PlayGameRel
		play        *Play
		game        *Game
		err         error
		startTime   time.Time
		endTime     time.Time
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

	err = p.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
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
	})

	return &v1.CreatePlayGameReply{
		PlayId: play.ID,
	}, err
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

func (p *PlayUseCase) GetPlayList(ctx context.Context, req *v1.GetPlayListRequest) (*v1.GetPlayListReply, error) {
	var (
		play        []*Play
		playIds     []int64
		playGameRel []*PlayGameRel
		err         error
	)

	playGameRel, err = p.playGameRelRepo.GetPlayGameRelByGameId(ctx, req.GameId)
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
			Pay:     item.Pay,
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
