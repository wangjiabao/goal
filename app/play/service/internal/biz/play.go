package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "goal/api/play/service/v1"
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

type PlayRoomRel struct {
	ID     int64
	RoomId int64
	PlayId int64
}

type PlayGameScoreUserRel struct {
	ID      int64
	UserId  int64
	PlayId  int64
	Content string
	Pay     int64
	Status  string
}

type PlayGameTeamSortUserRel struct {
	ID      int64
	UserId  int64
	PlayId  int64
	SortId  int64
	Status  string
	Content string
	Pay     int64
}

type PlayGameTeamGoalUserRel struct {
	ID     int64
	UserId int64
	PlayId int64
	TeamId int64
	Type   string
	Pay    int64
	Goal   int64
	Status string
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

type PlayRepo interface {
	GetAdminCreatePlayList(ctx context.Context) ([]*Play, error)
	GetAdminCreatePlayListByIds(ctx context.Context, ids ...int64) ([]*Play, error)
	GetPlayListByIds(ctx context.Context, ids ...int64) ([]*Play, error)
	CreatePlay(ctx context.Context, pc *Play) (*Play, error)
	GetAdminCreatePlayListByType(ctx context.Context, playType string) ([]*Play, error)
	GetPlayById(ctx context.Context, playId int64) (*Play, error)
}

type PlayGameRelRepo interface {
	GetPlayGameRelByGameId(ctx context.Context, gameId int64) ([]*PlayGameRel, error)
	CreatePlayGameRel(ctx context.Context, rel *PlayGameRel) (*PlayGameRel, error)
}

type PlaySortRelRepo interface {
	GetPlaySortRelBySortIds(ctx context.Context, sortIds ...int64) ([]*PlaySortRel, error)
	CreatePlaySortRel(ctx context.Context, rel *PlaySortRel) (*PlaySortRel, error)
}

type PlayRoomRelRepo interface {
	GetPlayRoomRelByRoomId(ctx context.Context, roomId int64) ([]*PlayRoomRel, error)
	CreatePlayRoomRel(ctx context.Context, pc *PlayRoomRel) (*PlayRoomRel, error)
}

type UserBalanceRepo interface {
	Pay(ctx context.Context, userId int64, pay int64) error
	TransferInto(ctx context.Context, userId int64, amount int64) error
	GetUserBalance(ctx context.Context, userId int64) (*UserBalance, error)
}

type UserProxyRepo interface {
	GetUserProxyAndDown(ctx context.Context) ([]*UserProxy, map[int64][]*UserProxy, error)
}

type PlayGameTeamResultUserRelRepo interface {
	CreatePlayGameTeamResultUserRel(ctx context.Context, pr *PlayGameTeamResultUserRel) (*PlayGameTeamResultUserRel, error)
}

type PlayGameTeamGoalUserRelRepo interface {
	CreatePlayGameTeamGoalUserRel(ctx context.Context, pr *PlayGameTeamGoalUserRel) (*PlayGameTeamGoalUserRel, error)
}

type PlayGameTeamSortUserRelRepo interface {
	CreatePlayGameTeamSortUserRel(ctx context.Context, pr *PlayGameTeamSortUserRel) (*PlayGameTeamSortUserRel, error)
}

type PlayGameScoreUserRelRepo interface {
	CreatePlayGameScoreUserRel(ctx context.Context, pr *PlayGameScoreUserRel) (*PlayGameScoreUserRel, error)
}

type PlayUseCase struct {
	playRepo                      PlayRepo
	playGameRelRepo               PlayGameRelRepo
	playSortRelRepo               PlaySortRelRepo
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
	playSortRelRepo PlaySortRelRepo,
	playRoomRelRepo PlayRoomRelRepo,
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
		playGameRelRepo:               playGameRelRepo,
		playSortRelRepo:               playSortRelRepo,
		playRoomRelRepo:               playRoomRelRepo,
		playGameScoreUserRelRepo:      playGameScoreUserRelRepo,
		playGameTeamSortUserRelRepo:   playGameTeamSortUserRelRepo,
		playGameTeamGoalUserRelRepo:   playGameTeamGoalUserRelRepo,
		playGameTeamResultUserRelRepo: playGameTeamResultUserRelRepo,
		userBalanceRepo:               userBalanceRepo,
		userProxyRepo:                 userProxyRepo,
		tx:                            tx,
		log:                           log.NewHelper(logger)}
}

// GetAdminCreateGameAndSortPlayList 获取指定比赛竞猜和一些排名竞猜的玩法的列表
func (p *PlayUseCase) GetAdminCreateGameAndSortPlayList(ctx context.Context, gameId int64, sortIds ...int64) (*v1.AllowedPlayListReply, error) {
	var (
		playIds     []int64 // todo 根据业务情况切片可能过大，不知道查询时会不会有问题，暂时这么处理
		plays       []*Play
		playGameRel []*PlayGameRel
		playSortRel []*PlaySortRel
		err         error
	)

	playGameRel, err = p.playGameRelRepo.GetPlayGameRelByGameId(ctx, gameId) // 获取比赛的玩法记录
	if err != nil {
		return nil, err
	}
	for _, v := range playGameRel {
		playIds = append(playIds, v.PlayId)
	}

	playSortRel, err = p.playSortRelRepo.GetPlaySortRelBySortIds(ctx, sortIds...) // 获取排名的玩法记录
	if err != nil {
		return nil, err
	}
	for _, v := range playSortRel {
		playIds = append(playIds, v.PlayId)
	}

	plays, err = p.playRepo.GetAdminCreatePlayListByIds(ctx, playIds...) // 获取admin创建的玩法
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

// GetRoomGameAndSortPlayList 获取房间内竞猜和一些排名竞猜的玩法的列表
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

	plays, err = p.playRepo.GetPlayListByIds(ctx, playIds...) // 获取admin创建的玩法
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

// CreatePlayGame 创建一个比赛玩法等记录
func (r *RoomUseCase) CreatePlayGame(ctx context.Context, req *v1.CreatePlayGameRequest) (*v1.CreatePlayGameReply, error) {
	var (
		userId      int64
		userType    string
		room        *Room
		playRoomRel *PlayRoomRel
		playGameRel *PlayGameRel
		play        *Play
		game        *Game
		err         error
		startTime   time.Time
		endTime     time.Time
	)

	game, err = r.gameRepo.GetGameById(ctx, req.SendBody.GameId) // 获取比赛信息以校验创建的玩法
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

	room, err = r.roomRepo.GetRoomByID(ctx, req.SendBody.RoomId) // 校验房间号 todo 类型

	if "game_team_goal_all" != req.SendBody.PlayType && // 验证type类型
		"game_score" != req.SendBody.PlayType &&
		"game_team_result" != req.SendBody.PlayType &&
		"game_team_goal_up" != req.SendBody.PlayType &&
		"game_team_goal_down" != req.SendBody.PlayType {
		return nil, errors.New(500, "TIME_ERROR", "玩法类型输入错误")
	}

	userId, userType, err = getUserFromJwt(ctx) // 获取用户id
	if nil != err {
		return nil, err
	}
	if "user" != userType && "admin" != userType {
		return nil, errors.New(500, "TIME_ERROR", "用户身份错误")
	}

	err = r.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		play, err = r.playRepo.CreatePlay(ctx, &Play{ // 新增玩法
			CreateUserId:   userId,
			CreateUserType: userType,
			Type:           req.SendBody.PlayType, // todo 用户输入参数未验证
			StartTime:      startTime,
			EndTime:        endTime,
		})
		if err != nil {
			return err
		}

		playRoomRel, err = r.playRoomRelRepo.CreatePlayRoomRel(ctx, &PlayRoomRel{ // 新增房间和玩法关系
			PlayId: play.ID,
			RoomId: room.ID,
		})
		if err != nil {
			return err
		}

		playGameRel, err = r.playGameRelRepo.CreatePlayGameRel(ctx, &PlayGameRel{ // 新增比赛和玩法关系
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
func (r *RoomUseCase) CreatePlaySort(ctx context.Context, req *v1.CreatePlaySortRequest) (*v1.CreatePlaySortReply, error) {
	var (
		userId      int64
		userType    string
		room        *Room
		playRoomRel *PlayRoomRel
		playSortRel *PlaySortRel
		play        *Play
		sort        *Sort
		err         error
		startTime   time.Time
		endTime     time.Time
	)

	sort, err = r.sortRepo.GetGameSortById(ctx, req.SendBody.SortId) // 获取排名截至日期以校验创建的玩法
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

	room, err = r.roomRepo.GetRoomByID(ctx, req.SendBody.RoomId) // 校验房间号 todo 类型

	userId, userType, err = getUserFromJwt(ctx) // 获取用户id
	if nil != err {
		return nil, err
	}
	if "user" != userType && "admin" != userType {
		return nil, errors.New(500, "TIME_ERROR", "用户身份错误")
	}

	if err = r.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		play, err = r.playRepo.CreatePlay(ctx, &Play{ // 新增玩法
			CreateUserId:   userId,
			CreateUserType: userType,
			Type:           sort.Type,
			StartTime:      startTime,
			EndTime:        endTime,
		})
		if err != nil {
			return err
		}

		playRoomRel, err = r.playRoomRelRepo.CreatePlayRoomRel(ctx, &PlayRoomRel{ // 新增房间和玩法关系
			PlayId: play.ID,
			RoomId: room.ID,
		})
		if err != nil {
			return err
		}

		playSortRel, err = r.playSortRelRepo.CreatePlaySortRel(ctx, &PlaySortRel{ // 新增排名和玩法关系
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

func (p *PlayUseCase) CreatePlayGameScore(ctx context.Context, req *v1.CreatePlayGameScoreRequest) (*v1.CreatePlayGameScoreReply, error) {

	var (
		userId               int64
		playGameScoreUserRel *PlayGameScoreUserRel
		play                 *Play
		pay                  int64
		userBalance          *UserBalance
		upUserProxy          []*UserProxy
		downUserProxy        map[int64][]*UserProxy
		err                  error
		feeRate              int64 = 5      // 根据base运算，意味着百分之十 todo 后台可以设置
		base                 int64 = 100000 // 基础精度0.00001 todo 加配置文件
		payLimit             int64 = 100    // 限额 todo 后台可以设置
	)
	// todo 参数真实验证，房间人员验证
	play, err = p.playRepo.GetPlayById(ctx, req.SendBody.PlayId) // 查玩法
	if nil != err {
		return nil, err
	}
	if "game_score" != play.Type {
		return nil, errors.New(500, "PLAY_ERROR", "玩法类型不匹配")
	}

	pay = req.SendBody.Pay * 100       // 基础的数是注，每注100在玩法这里*100
	if 0 != pay%payLimit || pay <= 0 { // 限制的整数倍
		return nil, errors.New(500, "PAY_ERROR", "玩法最低限额100")
	}

	userId, _, err = getUserFromJwt(ctx) // 获取用户id
	if nil != err {
		return nil, errors.New(500, "USER_ERROR", "用户信息错误")
	}
	userBalance, err = p.userBalanceRepo.GetUserBalance(ctx, userId) // 查余额
	if nil != err {
		return nil, err
	}
	pay = pay * base // todo 这样的数学计算不知道会不会有什么问题
	if pay > userBalance.Balance {
		return nil, errors.New(500, "USER_BALANCE_ERROR", "余额不足")
	}

	// 查找代理
	upUserProxy, downUserProxy, err = p.userProxyRepo.GetUserProxyAndDown(ctx)
	if nil != err {
		return nil, err
	}

	/* todo 优化点
	 * 考虑项目规模和业务场景，同一个人正确使用自己余额并且并发的情况较少，为了简单和有效应对恶意的并发请求，代码逻辑上加个简单的乐观锁,
	 * mysql使用innodb引擎隔离级别是读一致性读，读可加行锁，写会自动加排他锁，底层默认支持的情况更为见效，但是解决不了余额小于0的问题，
	 * 因此在事务中update余额时在条件中多加一条保证大于等于扣减的数额即可。
	 */
	if err = p.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = p.userBalanceRepo.Pay(ctx, userId, pay) // 余额扣除，先扣后加以防，给用户是代理余额增长了
		if err != nil {
			return err
		}

		fee := pay / feeRate // 扣除手续费
		pay -= fee
		for _, uv := range upUserProxy {
			var uvFee int64
			if 0 >= uv.Rate {
				continue
			}
			uvFee = fee / (100 / uv.Rate)
			if dv, ok := downUserProxy[uv.UserId]; ok {
				for _, v := range dv {
					var vFee int64
					if 0 >= v.Rate {
						continue
					}
					vFee = uvFee / (100 / v.Rate)                             // 本次转给小代理金额
					err = p.userBalanceRepo.TransferInto(ctx, v.UserId, vFee) // 转给小代理
					if err != nil {
						return err
					}
					uvFee -= vFee
					if uvFee < 0 { // 不足分
						break
					}
				}
			}
			err = p.userBalanceRepo.TransferInto(ctx, uv.UserId, uvFee) // 转给大代理
			if err != nil {
				return err
			}

			fee -= uvFee
			if fee < 0 {
				break // 分红余额已经不足
			}
		}

		playGameScoreUserRel, err = p.playGameScoreUserRelRepo.CreatePlayGameScoreUserRel(ctx, &PlayGameScoreUserRel{
			ID:      0,
			UserId:  userId,
			PlayId:  play.ID,
			Content: strconv.FormatInt(req.SendBody.RedScore, 10) + ":" + strconv.FormatInt(req.SendBody.BlueScore, 10),
			Pay:     pay,
			Status:  "no_rewarded",
		})
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
		upUserProxy               []*UserProxy
		downUserProxy             map[int64][]*UserProxy
		err                       error
		feeRate                   int64 = 5      // 根据base运算，意味着百分之十 todo 后台可以设置
		base                      int64 = 100000 // 基础精度0.00001 todo 加配置文件
		payLimit                  int64 = 100    // 限额 todo 后台可以设置
	)

	if strings.EqualFold("red", req.SendBody.Result) {
		gameResult = "red"
	} else if strings.EqualFold("blue", req.SendBody.Result) {
		gameResult = "blue"
	} else if strings.EqualFold("draw", req.SendBody.Result) {
		gameResult = "draw"
	} else {
		return nil, errors.New(500, "RESULT_ERROR", "比赛结果不匹配")
	}

	// todo 限制只能参加一次 todo 参数真实验证，房间人员验证

	play, err = p.playRepo.GetPlayById(ctx, req.SendBody.PlayId) // 查玩法
	if nil != err {
		return nil, err
	}
	if "game_team_result" != play.Type {
		return nil, errors.New(500, "PLAY_ERROR", "玩法类型不匹配")
	}

	pay = req.SendBody.Pay * 100       // 基础的数是注，每注100在玩法这里*100
	if 0 != pay%payLimit || pay <= 0 { // 限制的整数倍
		return nil, errors.New(500, "PAY_ERROR", "玩法最低限额100")
	}

	userId, _, err = getUserFromJwt(ctx) // 获取用户id
	if nil != err {
		return nil, errors.New(500, "USER_ERROR", "用户信息错误")
	}
	userBalance, err = p.userBalanceRepo.GetUserBalance(ctx, userId) // 查余额
	if nil != err {
		return nil, err
	}
	pay = pay * base // todo 这样的数学计算不知道会不会有什么问题
	if pay > userBalance.Balance {
		return nil, errors.New(500, "USER_BALANCE_ERROR", "余额不足")
	}

	// 查找代理
	upUserProxy, downUserProxy, err = p.userProxyRepo.GetUserProxyAndDown(ctx)
	if nil != err {
		return nil, err
	}

	/* todo 优化点
	 * 考虑项目规模和业务场景，同一个人正确使用自己余额并且并发的情况较少，为了简单和有效应对恶意的并发请求，代码逻辑上加个简单的乐观锁,
	 * mysql使用innodb引擎隔离级别是读一致性读，读可加行锁，写会自动加排他锁，底层默认支持的情况更为见效，但是解决不了余额小于0的问题，
	 * 因此在事务中update余额时在条件中多加一条保证大于等于扣减的数额即可。
	 */
	if err = p.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = p.userBalanceRepo.Pay(ctx, userId, pay) // 余额扣除，先扣后加以防，给用户是代理余额增长了
		if err != nil {
			return err
		}

		fee := pay / feeRate // 扣除手续费
		for _, uv := range upUserProxy {
			var uvFee int64
			if 0 >= uv.Rate {
				continue
			}
			uvFee = fee / (100 / uv.Rate)
			if dv, ok := downUserProxy[uv.UserId]; ok {
				for _, v := range dv {
					var vFee int64
					if 0 >= v.Rate {
						continue
					}
					vFee = uvFee / (100 / v.Rate)                             // 本次转给小代理金额
					err = p.userBalanceRepo.TransferInto(ctx, v.UserId, vFee) // 转给小代理
					if err != nil {
						return err
					}
					uvFee -= vFee
					if uvFee < 0 { // 不足分
						break
					}
				}
			}
			err = p.userBalanceRepo.TransferInto(ctx, uv.UserId, uvFee) // 转给大代理
			if err != nil {
				return err
			}

			fee -= uvFee
			if fee < 0 {
				break // 分红余额已经不足
			}
		}

		playGameTeamResultUserRel, err = p.playGameTeamResultUserRelRepo.CreatePlayGameTeamResultUserRel(ctx, &PlayGameTeamResultUserRel{
			ID:      0,
			UserId:  userId,
			PlayId:  play.ID,
			Content: gameResult,
			Pay:     pay,
			Status:  "no_rewarded",
		})
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
		upUserProxy             []*UserProxy
		downUserProxy           map[int64][]*UserProxy
		err                     error
		feeRate                 int64 = 5      // 根据base运算，意味着百分之十 todo 后台可以设置
		base                    int64 = 100000 // 基础精度0.00001 todo 加配置文件
		payLimit                int64 = 100    // 限额 todo 后台可以设置
	)

	// todo 参数真实验证，房间人员验证
	play, err = p.playRepo.GetPlayById(ctx, req.SendBody.PlayId) // 查玩法
	if nil != err {
		return nil, err
	}

	if "team_sort_eight" != play.Type && "team_sort_three" != play.Type && "team_sort_sixteen" != play.Type {
		return nil, errors.New(500, "PLAY_ERROR", "玩法类型不匹配")
	}

	pay = req.SendBody.Pay * 100       // 基础的数是注，每注100在玩法这里*100
	if 0 != pay%payLimit || pay <= 0 { // 限制的整数倍
		return nil, errors.New(500, "PAY_ERROR", "玩法最低限额100")
	}

	userId, _, err = getUserFromJwt(ctx) // 获取用户id
	if nil != err {
		return nil, errors.New(500, "USER_ERROR", "用户信息错误")
	}
	userBalance, err = p.userBalanceRepo.GetUserBalance(ctx, userId) // 查余额
	if nil != err {
		return nil, err
	}
	pay = pay * base // todo 这样的数学计算不知道会不会有什么问题
	if pay > userBalance.Balance {
		return nil, errors.New(500, "USER_BALANCE_ERROR", "余额不足")
	}

	// 查找代理
	upUserProxy, downUserProxy, err = p.userProxyRepo.GetUserProxyAndDown(ctx)
	if nil != err {
		return nil, err
	}

	/* todo 优化点
	 * 考虑项目规模和业务场景，同一个人正确使用自己余额并且并发的情况较少，为了简单和有效应对恶意的并发请求，代码逻辑上加个简单的乐观锁,
	 * mysql使用innodb引擎隔离级别是读一致性读，读可加行锁，写会自动加排他锁，底层默认支持的情况更为见效，但是解决不了余额小于0的问题，
	 * 因此在事务中update余额时在条件中多加一条保证大于等于扣减的数额即可。
	 */
	if err = p.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = p.userBalanceRepo.Pay(ctx, userId, pay) // 余额扣除，先扣后加以防，给用户是代理余额增长了
		if err != nil {
			return err
		}

		fee := pay / feeRate // 扣除手续费
		for _, uv := range upUserProxy {
			var uvFee int64
			if 0 >= uv.Rate {
				continue
			}
			uvFee = fee / (100 / uv.Rate)
			if dv, ok := downUserProxy[uv.UserId]; ok {
				for _, v := range dv {
					var vFee int64
					if 0 >= v.Rate {
						continue
					}
					vFee = uvFee / (100 / v.Rate)                             // 本次转给小代理金额
					err = p.userBalanceRepo.TransferInto(ctx, v.UserId, vFee) // 转给小代理
					if err != nil {
						return err
					}
					uvFee -= vFee
					if uvFee < 0 { // 不足分
						break
					}
				}
			}
			err = p.userBalanceRepo.TransferInto(ctx, uv.UserId, uvFee) // 转给大代理
			if err != nil {
				return err
			}

			fee -= uvFee
			if fee < 0 {
				break // 分红余额已经不足
			}
		}

		playGameTeamSortUserRel, err = p.playGameTeamSortUserRelRepo.CreatePlayGameTeamSortUserRel(ctx, &PlayGameTeamSortUserRel{
			ID:      0,
			UserId:  userId,
			PlayId:  play.ID,
			SortId:  req.SendBody.SortId,
			Content: req.SendBody.Content,
			Pay:     pay,
			Status:  "no_rewarded",
		})
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
		upUserProxy             []*UserProxy
		downUserProxy           map[int64][]*UserProxy
		err                     error
		feeRate                 int64 = 5      // 根据base运算，意味着百分之十 todo 后台可以设置
		base                    int64 = 100000 // 基础精度0.00001 todo 加配置文件
		payLimit                int64 = 100    // 限额 todo 后台可以设置
	)

	play, err = p.playRepo.GetPlayById(ctx, req.SendBody.PlayId) // 查玩法
	if nil != err {
		return nil, err
	}

	if "game_team_goal_all" != req.SendBody.PlayType && "game_team_goal_up" != req.SendBody.PlayType && "game_team_goal_down" != req.SendBody.PlayType {
		return nil, errors.New(500, "PLAY_ERROR", "玩法类型不匹配")
	}

	pay = req.SendBody.Pay * 100       // 基础的数是注，每注100在玩法这里*100
	if 0 != pay%payLimit || pay <= 0 { // 限制的整数倍
		return nil, errors.New(500, "PAY_ERROR", "玩法最低限额100")
	}

	userId, _, err = getUserFromJwt(ctx) // 获取用户id
	if nil != err {
		return nil, errors.New(500, "USER_ERROR", "用户信息错误")
	}
	userBalance, err = p.userBalanceRepo.GetUserBalance(ctx, userId) // 查余额
	if nil != err {
		return nil, err
	}
	pay = pay * base // todo 这样的数学计算不知道会不会有什么问题
	if pay > userBalance.Balance {
		return nil, errors.New(500, "USER_BALANCE_ERROR", "余额不足")
	}

	// 查找代理
	upUserProxy, downUserProxy, err = p.userProxyRepo.GetUserProxyAndDown(ctx)
	if nil != err {
		return nil, err
	}

	/* todo 优化点
	 * 考虑项目规模和业务场景，同一个人正确使用自己余额并且并发的情况较少，为了简单和有效应对恶意的并发请求，代码逻辑上加个简单的乐观锁,
	 * mysql使用innodb引擎隔离级别是读一致性读，读可加行锁，写会自动加排他锁，底层默认支持的情况更为见效，但是解决不了余额小于0的问题，
	 * 因此在事务中update余额时在条件中多加一条保证大于等于扣减的数额即可。
	 */
	if err = p.tx.ExecTx(ctx, func(ctx context.Context) error { // 事务
		err = p.userBalanceRepo.Pay(ctx, userId, pay) // 余额扣除，先扣后加以防，给用户是代理余额增长了
		if err != nil {
			return err
		}

		fee := pay / feeRate // 扣除手续费
		for _, uv := range upUserProxy {
			var uvFee int64
			if 0 >= uv.Rate {
				continue
			}
			uvFee = fee / (100 / uv.Rate)
			if dv, ok := downUserProxy[uv.UserId]; ok {
				for _, v := range dv {
					var vFee int64
					if 0 >= v.Rate {
						continue
					}
					vFee = uvFee / (100 / v.Rate)                             // 本次转给小代理金额
					err = p.userBalanceRepo.TransferInto(ctx, v.UserId, vFee) // 转给小代理
					if err != nil {
						return err
					}
					uvFee -= vFee
					if uvFee < 0 { // 不足分
						break
					}
				}
			}
			err = p.userBalanceRepo.TransferInto(ctx, uv.UserId, uvFee) // 转给大代理
			if err != nil {
				return err
			}

			fee -= uvFee
			if fee < 0 {
				break // 分红余额已经不足
			}
		}

		playGameTeamGoalUserRel, err = p.playGameTeamGoalUserRelRepo.CreatePlayGameTeamGoalUserRel(ctx, &PlayGameTeamGoalUserRel{
			ID:     0,
			UserId: userId,
			PlayId: play.ID,
			TeamId: req.SendBody.TeamId,
			Type:   req.SendBody.PlayType,
			Goal:   req.SendBody.Goal,
			Pay:    pay,
			Status: "no_rewarded",
		})
		if err != nil {
			return err
		}
		return nil
	}); nil != err {
		return nil, err
	}

	return &v1.CreatePlayGameGoalReply{PlayId: playGameTeamGoalUserRel.PlayId}, nil
}
