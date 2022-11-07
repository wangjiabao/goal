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

type UserBalance struct {
	ID      int64
	UserId  int64
	Balance int64
}

type UserProxy struct {
	ID       int64
	UserId   int64
	UpUserId int64
	Rate     int64
}

type UserInfo struct {
	ID              int64
	UserId          int64
	Name            string
	Avatar          string
	RecommendCode   string
	MyRecommendCode string
	Code            string
}

type PlayRepo interface {
	GetPlayListByIds(ctx context.Context, ids ...int64) ([]*Play, error)
	CreatePlay(ctx context.Context, pc *Play) (*Play, error)
}

type PlayGameRelRepo interface {
	GetPlayGameRelByGameId(ctx context.Context, gameId int64) ([]*PlayGameRel, error)
	CreatePlayGameRel(ctx context.Context, rel *PlayGameRel) (*PlayGameRel, error)
}

type PlaySortRelRepo interface {
	CreatePlaySortRel(ctx context.Context, rel *PlaySortRel) (*PlaySortRel, error)
}
type PlayGameScoreUserRelRepo interface {
	SetRewarded(ctx context.Context, userId int64) error
	GetPlayGameScoreUserRelByPlayIds(ctx context.Context, playIds ...int64) (map[int64][]*PlayGameScoreUserRel, error)
}

type PlayGameTeamResultUserRelRepo interface {
	GetPlayGameScoreUserRelByPlayIds(ctx context.Context, playIds ...int64) (map[int64][]*PlayGameTeamResultUserRel, error)
	SetRewarded(ctx context.Context, userId int64) error
}

type UserBalanceRepo interface {
	TransferInto(ctx context.Context, userId int64, amount int64) error
	GetUserBalance(ctx context.Context, userId int64) (*UserBalance, error)
}

type UserProxyRepo interface {
	GetUserProxyAndDown(ctx context.Context) ([]*UserProxy, map[int64][]*UserProxy, error)
}

type UserInfoRepo interface {
	GetUserInfoByUserId(ctx context.Context, userId int64) (*UserInfo, error)
}

type PlayUseCase struct {
	playRepo                      PlayRepo
	gameRepo                      GameRepo
	playGameRelRepo               PlayGameRelRepo
	playSortRelRepo               PlaySortRelRepo
	playGameScoreUserRelRepo      PlayGameScoreUserRelRepo
	playGameTeamResultUserRelRepo PlayGameTeamResultUserRelRepo
	userBalanceRepo               UserBalanceRepo
	userProxyRepo                 UserProxyRepo
	userInfoRepo                  UserInfoRepo
	sortRepo                      SortRepo
	tx                            Transaction
	log                           *log.Helper
}

func NewPlayUseCase(
	repo PlayRepo,
	playGameRelRepo PlayGameRelRepo,
	playSortRelRepo PlaySortRelRepo,
	playGameScoreUserRelRepo PlayGameScoreUserRelRepo,

	playGameTeamResultUserRelRepo PlayGameTeamResultUserRelRepo,
	gameRepo GameRepo,
	sortRepo SortRepo,
	userBalanceRepo UserBalanceRepo,
	userProxyRepo UserProxyRepo,
	userInfoRepo UserInfoRepo,
	tx Transaction,
	logger log.Logger) *PlayUseCase {
	return &PlayUseCase{
		playRepo:                 repo,
		gameRepo:                 gameRepo,
		sortRepo:                 sortRepo,
		playGameRelRepo:          playGameRelRepo,
		playSortRelRepo:          playSortRelRepo,
		playGameScoreUserRelRepo: playGameScoreUserRelRepo,

		playGameTeamResultUserRelRepo: playGameTeamResultUserRelRepo,
		userBalanceRepo:               userBalanceRepo,
		userProxyRepo:                 userProxyRepo,
		userInfoRepo:                  userInfoRepo,
		tx:                            tx,
		log:                           log.NewHelper(logger),
	}
}

func (p *PlayUseCase) GamePlayGrant(ctx context.Context, req *v1.GamePlayGrantRequest) (*v1.GamePlayGrantReply, error) {
	var (
		game           *Game
		playGameRel    []*PlayGameRel
		playIds        []int64
		play           []*Play
		playGameScore  []*Play
		playGameResult []*Play
		err            error
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
		}
	}

	p.grantTypeGameScore(ctx, game, playGameScore)
	p.grantTypeGameResult(ctx, game, playGameResult)

	return &v1.GamePlayGrantReply{
		Result: "处理完成",
	}, nil
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
				winTotalAmount += v.Pay
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
			for _, ruv := range strings.Split(userInfo.RecommendCode, "GA") {    // 解析除userId, 取前三代
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
						if res := p.userBalanceRepo.TransferInto(ctx, recommendUserId, tmpPerAmount); nil != res {
							return res
						}
					} else if 1 == k {
						tmpPerAmount = perAmount * 3 / 1000
						if res := p.userBalanceRepo.TransferInto(ctx, recommendUserId, tmpPerAmount); nil != res {
							return res
						}
					} else if 2 == k {
						tmpPerAmount = perAmount * 5 / 1000
						tmpPerAmount += perAmount * 10 / 1000
						if res := p.userBalanceRepo.TransferInto(ctx, recommendUserId, tmpPerAmount); nil != res {
							return res
						}
					}
					perAmount -= tmpPerAmount
				}

				if res := p.userBalanceRepo.TransferInto(ctx, winV.UserId, perAmount); nil != res {
					return res
				}

				if res := p.playGameScoreUserRelRepo.SetRewarded(ctx, winV.UserId); nil != res {
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

	playGameTeamResultUserRel, err = p.playGameTeamResultUserRelRepo.GetPlayGameScoreUserRelByPlayIds(ctx, playIds...)
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
				winTotalAmount += v.Pay
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
			for _, ruv := range strings.Split(userInfo.RecommendCode, "GA") {    // 解析除userId, 取前三代
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
						if res := p.userBalanceRepo.TransferInto(ctx, recommendUserId, tmpPerAmount); nil != res {
							return res
						}
					} else if 1 == k {
						tmpPerAmount = perAmount * 3 / 1000
						if res := p.userBalanceRepo.TransferInto(ctx, recommendUserId, tmpPerAmount); nil != res {
							return res
						}
					} else if 2 == k {
						tmpPerAmount = perAmount * 5 / 1000
						tmpPerAmount += perAmount * 10 / 1000
						if res := p.userBalanceRepo.TransferInto(ctx, recommendUserId, tmpPerAmount); nil != res {
							return res
						}
					}
					perAmount -= tmpPerAmount
				}

				perAmount += winV.Pay // 押注的钱原路返回
				if res := p.userBalanceRepo.TransferInto(ctx, winV.UserId, perAmount); nil != res {
					return res
				}

				if res := p.playGameTeamResultUserRelRepo.SetRewarded(ctx, winV.UserId); nil != res {
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
