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

type PlayGameScoreUserRel struct {
	ID      int64
	PlayId  int64
	UserId  int64
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
}

type PlayGameRelRepo interface {
	GetPlayGameRelByGameId(ctx context.Context, gameId int64) ([]*PlayGameRel, error)
}

type PlayGameScoreUserRelRepo interface {
	SetRewarded(ctx context.Context, userId int64) error
	GetPlayGameScoreUserRelByPlayIds(ctx context.Context, playIds ...int64) (map[int64][]*PlayGameScoreUserRel, error)
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
	playRepo                 PlayRepo
	gameRepo                 GameRepo
	playGameRelRepo          PlayGameRelRepo
	playGameScoreUserRelRepo PlayGameScoreUserRelRepo
	userBalanceRepo          UserBalanceRepo
	userProxyRepo            UserProxyRepo
	userInfoRepo             UserInfoRepo
	tx                       Transaction
	log                      *log.Helper
}

func NewPlayUseCase(
	repo PlayRepo,
	playGameRelRepo PlayGameRelRepo,
	playGameScoreUserRelRepo PlayGameScoreUserRelRepo,
	gameRepo GameRepo,
	userBalanceRepo UserBalanceRepo,
	userProxyRepo UserProxyRepo,
	userInfoRepo UserInfoRepo,
	tx Transaction,
	logger log.Logger) *PlayUseCase {
	return &PlayUseCase{
		playRepo:                 repo,
		gameRepo:                 gameRepo,
		playGameRelRepo:          playGameRelRepo,
		playGameScoreUserRelRepo: playGameScoreUserRelRepo,
		userBalanceRepo:          userBalanceRepo,
		userProxyRepo:            userProxyRepo,
		userInfoRepo:             userInfoRepo,
		tx:                       tx,
		log:                      log.NewHelper(logger),
	}
}

func (p *PlayUseCase) GamePlayGrant(ctx context.Context, req *v1.GamePlayGrantRequest) (*v1.GamePlayGrantReply, error) {
	var (
		game          *Game
		playGameRel   []*PlayGameRel
		playIds       []int64
		play          []*Play
		playGameScore []*Play
		err           error
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
		}
	}

	if !p.grantTypeGameScore(ctx, game, playGameScore) {
		return nil, errors.New(500, "HANDLE_ERROR", "处理失败")
	}

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
		// 每一场玩法
		for _, v := range playUserRel {
			var (
				winNoRewardedPlayGameScoreUserRel []*PlayGameScoreUserRel // 猜中未发放奖励的用户
				poolAmount                        int64                   // 每个玩法的奖池
			)
			poolAmount += v.Pay
			if strings.EqualFold(game.Result, v.Content) && strings.EqualFold("no_rewarded", v.Status) { // 判断是否猜中
				winNoRewardedPlayGameScoreUserRel = append(winNoRewardedPlayGameScoreUserRel, v)
			}
			sizeofWin := int64(len(winNoRewardedPlayGameScoreUserRel))
			if 0 == sizeofWin {
				continue
			}

			perAmount := poolAmount / sizeofWin
			for _, winV := range winNoRewardedPlayGameScoreUserRel {
				var (
					recommendUserIds []int64
					userInfo         *UserInfo
				)
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
	}

	return true
}
