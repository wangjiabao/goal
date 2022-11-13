package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "goal/api/admin/service/v1"
	"time"
)

type Game struct {
	ID               int64
	Name             string
	RedTeamId        int64
	BlueTeamId       int64
	WinTeamId        int64
	Result           string
	RedTeamUpGoal    int64
	BlueTeamUpGoal   int64
	RedTeamDownGoal  int64
	BlueTeamDownGoal int64
	Status           string
	UpEndTime        time.Time
	DownStartTime    time.Time
	StartTime        time.Time
	EndTime          time.Time
}

type DisplayGame struct {
	ID     int64
	GameId int64
	Type   string
}

type GameRepo interface {
	GetGameById(ctx context.Context, gameId int64) (*Game, error)
	CreateGame(ctx context.Context, gc *Game) (*Game, error)
	UpdateGame(ctx context.Context, gc *Game) (*Game, error)
	GetGameList(ctx context.Context) ([]*Game, error)
	GetDisplayGame() (*DisplayGame, error)
	GetDisplayGameList() ([]*DisplayGame, error)
	UpdateDisplayGame(ctx context.Context, displayGame *DisplayGame, gameId int64) (*DisplayGame, error)
	CreateDisplayGame(ctx context.Context, gameId int64) (*DisplayGame, error)
	DeleteDisplayGame(ctx context.Context, gameId int64) (bool, error)
}

type GameUseCase struct {
	gameRepo                      GameRepo
	playRepo                      PlayRepo
	playGameRelRepo               PlayGameRelRepo
	playGameScoreUserRelRepo      PlayGameScoreUserRelRepo
	playGameTeamGoalUserRelRepo   PlayGameTeamGoalUserRelRepo
	playGameTeamResultUserRelRepo PlayGameTeamResultUserRelRepo
	log                           *log.Helper
}

func NewGameUseCase(
	repo GameRepo,
	playGameRelRepo PlayGameRelRepo,
	playRepo PlayRepo,
	playGameScoreUserRelRepo PlayGameScoreUserRelRepo,
	playGameTeamGoalUserRelRepo PlayGameTeamGoalUserRelRepo,
	playGameTeamResultUserRelRepo PlayGameTeamResultUserRelRepo,
	logger log.Logger,
) *GameUseCase {
	return &GameUseCase{
		gameRepo: repo,

		playRepo:                      playRepo,
		playGameScoreUserRelRepo:      playGameScoreUserRelRepo,
		playGameTeamGoalUserRelRepo:   playGameTeamGoalUserRelRepo,
		playGameTeamResultUserRelRepo: playGameTeamResultUserRelRepo,
		playGameRelRepo:               playGameRelRepo,
		log:                           log.NewHelper(logger),
	}
}

func (g *GameUseCase) CreateGame(ctx context.Context, req *v1.CreateGameRequest) (*v1.CreateGameReply, error) {
	var (
		err           error
		game          *Game
		downStartTime time.Time
		upEndTime     time.Time
		startTime     time.Time
		endTime       time.Time
	)

	startTime, err = time.Parse("2006-01-02 15:04:05", req.SendBody.StartTime) // 时间进行格式校验
	if nil != err {
		return nil, err
	}
	endTime, err = time.Parse("2006-01-02 15:04:05", req.SendBody.EndTime) // 时间进行格式校验
	if nil != err {
		return nil, err
	}
	upEndTime, err = time.Parse("2006-01-02 15:04:05", req.SendBody.UpEndTime) // 时间进行格式校验
	if nil != err {
		return nil, err
	}
	downStartTime, err = time.Parse("2006-01-02 15:04:05", req.SendBody.DownStartTime) // 时间进行格式校验
	if nil != err {
		return nil, err
	}

	if game, err = g.gameRepo.CreateGame(ctx, &Game{
		Name:          req.SendBody.Name,      // todo 参数校验
		RedTeamId:     req.SendBody.RedTeamId, // todo 正确性校验
		BlueTeamId:    req.SendBody.BlueTeamId,
		UpEndTime:     upEndTime,
		DownStartTime: downStartTime,
		StartTime:     startTime,
		EndTime:       endTime,
	}); nil != err {
		return nil, err
	}

	return &v1.CreateGameReply{GameId: game.ID}, nil
}

func (g *GameUseCase) UpdateGame(ctx context.Context, req *v1.UpdateGameRequest) (*v1.UpdateGameReply, error) {

	var (
		err           error
		game          *Game
		downStartTime time.Time
		upEndTime     time.Time
		startTime     time.Time
		endTime       time.Time
	)

	startTime, err = time.Parse("2006-01-02 15:04:05", req.SendBody.StartTime) // 时间进行格式校验
	if nil != err {
		return nil, err
	}
	endTime, err = time.Parse("2006-01-02 15:04:05", req.SendBody.EndTime) // 时间进行格式校验
	if nil != err {
		return nil, err
	}
	upEndTime, err = time.Parse("2006-01-02 15:04:05", req.SendBody.UpEndTime) // 时间进行格式校验
	if nil != err {
		return nil, err
	}
	downStartTime, err = time.Parse("2006-01-02 15:04:05", req.SendBody.DownStartTime) // 时间进行格式校验
	if nil != err {
		return nil, err
	}

	if game, err = g.gameRepo.UpdateGame(ctx, &Game{
		ID:               req.SendBody.GameId,
		Name:             req.SendBody.Name, // todo 参数校验
		UpEndTime:        upEndTime,
		DownStartTime:    downStartTime,
		StartTime:        startTime,
		EndTime:          endTime,
		Status:           req.SendBody.Status,
		RedTeamDownGoal:  req.SendBody.RedTeamDownGoal,
		BlueTeamDownGoal: req.SendBody.BlueTeamDownGoal,
		RedTeamUpGoal:    req.SendBody.RedTeamUpGoal,
		BlueTeamUpGoal:   req.SendBody.BlueTeamUpGoal,
		WinTeamId:        req.SendBody.WinTeamId,
	}); nil != err {
		return nil, err
	}

	return &v1.UpdateGameReply{
		GameId: game.ID,
	}, nil
}

func (g *GameUseCase) GetGameList(ctx context.Context) (*v1.GetGameListReply, error) {
	var (
		game []*Game
		err  error
	)

	game, err = g.gameRepo.GetGameList(ctx)
	if err != nil {
		return nil, err
	}

	res := &v1.GetGameListReply{
		Items: make([]*v1.GetGameListReply_Item, 0),
	}
	for _, item := range game {
		res.Items = append(res.Items, &v1.GetGameListReply_Item{
			GameId:        item.ID,
			Name:          item.Name,
			StartTime:     item.StartTime.Format("2006-01-02 15:04:05"),
			EndTime:       item.EndTime.Format("2006-01-02 15:04:05"),
			UpEndTime:     item.UpEndTime.Format("2006-01-02 15:04:05"),
			DownStartTime: item.DownStartTime.Format("2006-01-02 15:04:05"),
		})
	}

	return res, nil
}

func (g *GameUseCase) GetGameByGameId(ctx context.Context, req *v1.GetGameRequest) (*v1.GetGameReply, error) {
	game, err := g.gameRepo.GetGameById(ctx, req.GameId)
	if err != nil {
		return nil, err
	}

	return &v1.GetGameReply{
		GameId:           game.ID,
		Name:             game.Name,
		RedTeamId:        game.RedTeamId,
		BlueTeamId:       game.BlueTeamId,
		WinTeamId:        game.WinTeamId,
		Result:           game.Result,
		RedTeamUpGoal:    game.RedTeamUpGoal,
		BlueTeamUpGoal:   game.BlueTeamUpGoal,
		RedTeamDownGoal:  game.RedTeamDownGoal,
		BlueTeamDownGoal: game.BlueTeamDownGoal,
		Status:           game.Status,
		UpEndTime:        game.UpEndTime.Format("2006-01-02 15:04:05"),
		DownStartTime:    game.DownStartTime.Format("2006-01-02 15:04:05"),
		StartTime:        game.StartTime.Format("2006-01-02 15:04:05"),
		EndTime:          game.EndTime.Format("2006-01-02 15:04:05"),
	}, nil
}

func (g *GameUseCase) GetDisplayGameIndex(ctx context.Context) (*v1.DisplayGameIndexReply, error) {
	disPlayGame, err := g.gameRepo.GetDisplayGameList()
	if err != nil {
		return nil, err
	}

	res := &v1.DisplayGameIndexReply{
		Items: make([]*v1.DisplayGameIndexReply_Item, 0),
	}

	for _, v := range disPlayGame {
		res.Items = append(res.Items, &v1.DisplayGameIndexReply_Item{
			GameId: v.GameId,
		})
	}
	return res, nil
}

func (g *GameUseCase) SaveDisplayGameIndex(ctx context.Context, req *v1.SaveDisplayGameIndexRequest) (*v1.SaveDisplayGameIndexReply, error) {
	disPlayGame, err := g.gameRepo.CreateDisplayGame(ctx, req.SendBody.GameId)
	if err != nil {
		return nil, err
	}
	return &v1.SaveDisplayGameIndexReply{
		GameId: disPlayGame.GameId,
	}, nil
}

func (g *GameUseCase) DeleteDisplayGameIndex(ctx context.Context, req *v1.DeleteDisplayGameIndexRequest) (*v1.DeleteDisplayGameIndexReply, error) {
	_, err := g.gameRepo.DeleteDisplayGame(ctx, req.SendBody.GameId)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteDisplayGameIndexReply{
		Result: "成功",
	}, nil
}

func (g *GameUseCase) GameIndexStatistics(ctx context.Context, req *v1.GameIndexStatisticsRequest) (*v1.GameIndexStatisticsReply, error) {
	var (
		err                               error
		game                              *Game
		playIds                           []int64
		play                              []*Play
		playGameRel                       []*PlayGameRel
		playGameScoreUserRelPlayId        int64
		playGameScoreUserRel              []*PlayGameScoreUserRel
		playGameTeamGoalDownUserRelPlayId int64
		playGameTeamGoalUpUserRelPlayId   int64
		playGameTeamGoalAllUserRelPlayId  int64
		playGameTeamGoalUpUserRel         []*PlayGameTeamGoalUserRel
		playGameTeamGoalDownUserRel       []*PlayGameTeamGoalUserRel
		playGameTeamGoalAllUserRel        []*PlayGameTeamGoalUserRel
		playGameTeamResultUserRelPlayId   int64
		playGameTeamResultUserRel         []*PlayGameTeamResultUserRel
		base                              int64 = 100000 // 基础精度0.00001 todo 加配置文件
	)

	if err != nil {
		return nil, err
	}
	game, err = g.gameRepo.GetGameById(ctx, req.GameId)
	if err != nil {
		return nil, err
	}

	playGameRel, _ = g.playGameRelRepo.GetPlayGameRelByGameId(ctx, req.GameId)
	for _, v := range playGameRel {
		playIds = append(playIds, v.PlayId)
	}

	play, _ = g.playRepo.GetPlayListByIds(ctx, playIds...)
	for _, v := range play {
		if "game_team_goal_all" == v.Type {
			playGameTeamGoalAllUserRelPlayId = v.ID
		} else if "game_team_goal_up" == v.Type {
			playGameTeamGoalUpUserRelPlayId = v.ID
		} else if "game_team_goal_down" == v.Type {
			playGameTeamGoalDownUserRelPlayId = v.ID
		} else if "game_team_result" == v.Type {
			playGameTeamResultUserRelPlayId = v.ID
		} else if "game_score" == v.Type {
			playGameScoreUserRelPlayId = v.ID
		}
	}

	res := &v1.GameIndexStatisticsReply{
		GameId:            game.ID,
		GameName:          game.Name,
		GoalAllPlayId:     0,
		GoalAllTotal:      0,
		GoalAllRedTotal:   0,
		GoalAllBlueTotal:  0,
		GoalUpPlayId:      0,
		GoalUpTotal:       0,
		GoalUpRedTotal:    0,
		GoalUpBlueTotal:   0,
		GoalDownPlayId:    0,
		GoalDownTotal:     0,
		GoalDownRedTotal:  0,
		GoalDownBlueTotal: 0,
		ResultPlayId:      0,
		ResultRedTotal:    0,
		ResultTotal:       0,
		ResultBlueTotal:   0,
		ResultDrawTotal:   0,
		ScorePlayId:       0,
		ScoreTotal:        0,
		GoalAllRed:        make([]*v1.GameIndexStatisticsReply_GoalAllRed, 0),
		GoalAllBlue:       make([]*v1.GameIndexStatisticsReply_GoalAllBlue, 0),
		GoalUpRed:         make([]*v1.GameIndexStatisticsReply_GoalUpRed, 0),
		GoalUpBlue:        make([]*v1.GameIndexStatisticsReply_GoalUpBlue, 0),
		GoalDownRed:       make([]*v1.GameIndexStatisticsReply_GoalDownRed, 0),
		GoalDownBlue:      make([]*v1.GameIndexStatisticsReply_GoalDownBlue, 0),
		Score:             make([]*v1.GameIndexStatisticsReply_Score, 0),
	}

	// 进球数
	// todo 查一次库用切片
	playGameTeamGoalUpUserRel, _ = g.playGameTeamGoalUserRelRepo.GetPlayGameTeamGoalUserRelByPlayId(ctx, playGameTeamGoalUpUserRelPlayId)
	res.GoalUpPlayId = playGameTeamGoalUpUserRelPlayId
	playGameTeamGoalRedUpUserRelMap := make(map[int64]struct {
		Goal  int64
		Total int64
	})
	playGameTeamGoalBlueUpUserRelMap := make(map[int64]struct {
		Goal  int64
		Total int64
	})

	for _, v := range playGameTeamGoalUpUserRel {
		if v.TeamId == game.RedTeamId {
			playGameTeamGoalRedUpUserRelMap[v.Goal] = struct {
				Goal  int64
				Total int64
			}{Goal: v.Goal, Total: v.Pay + playGameTeamGoalRedUpUserRelMap[v.Goal].Total}
		}

		if v.TeamId == game.BlueTeamId {
			playGameTeamGoalBlueUpUserRelMap[v.Goal] = struct {
				Goal  int64
				Total int64
			}{Goal: v.Goal, Total: v.Pay + playGameTeamGoalBlueUpUserRelMap[v.Goal].Total}
		}

	}

	for _, v := range playGameTeamGoalRedUpUserRelMap {
		tmpTotal := v.Total / base
		res.GoalUpRedTotal += tmpTotal
		res.GoalUpTotal += tmpTotal
		res.GoalUpRed = append(res.GoalUpRed, &v1.GameIndexStatisticsReply_GoalUpRed{
			Goal:  v.Goal,
			Total: tmpTotal,
		})
	}
	for _, v := range playGameTeamGoalBlueUpUserRelMap {
		tmpTotal := v.Total / base
		res.GoalUpBlueTotal += tmpTotal
		res.GoalUpTotal += tmpTotal
		res.GoalUpBlue = append(res.GoalUpBlue, &v1.GameIndexStatisticsReply_GoalUpBlue{
			Goal:  v.Goal,
			Total: tmpTotal,
		})
	}

	playGameTeamGoalDownUserRel, _ = g.playGameTeamGoalUserRelRepo.GetPlayGameTeamGoalUserRelByPlayId(ctx, playGameTeamGoalDownUserRelPlayId)
	res.GoalDownPlayId = playGameTeamGoalDownUserRelPlayId
	playGameTeamGoalRedDownUserRelMap := make(map[int64]struct {
		Goal  int64
		Total int64
	})
	playGameTeamGoalBlueDownUserRelMap := make(map[int64]struct {
		Goal  int64
		Total int64
	})

	for _, v := range playGameTeamGoalDownUserRel {
		if v.TeamId == game.RedTeamId {
			playGameTeamGoalRedDownUserRelMap[v.Goal] = struct {
				Goal  int64
				Total int64
			}{Goal: v.Goal, Total: v.Pay + playGameTeamGoalRedDownUserRelMap[v.Goal].Total}
		}

		if v.TeamId == game.BlueTeamId {
			playGameTeamGoalBlueDownUserRelMap[v.Goal] = struct {
				Goal  int64
				Total int64
			}{Goal: v.Goal, Total: v.Pay + playGameTeamGoalBlueDownUserRelMap[v.Goal].Total}
		}
	}

	for _, v := range playGameTeamGoalRedDownUserRelMap {
		tmpTotal := v.Total / base
		res.GoalDownRedTotal += tmpTotal
		res.GoalDownTotal += tmpTotal
		res.GoalDownRed = append(res.GoalDownRed, &v1.GameIndexStatisticsReply_GoalDownRed{
			Goal:  v.Goal,
			Total: tmpTotal,
		})
	}
	for _, v := range playGameTeamGoalBlueDownUserRelMap {
		tmpTotal := v.Total / base
		res.GoalDownBlueTotal += tmpTotal
		res.GoalDownTotal += tmpTotal
		res.GoalDownBlue = append(res.GoalDownBlue, &v1.GameIndexStatisticsReply_GoalDownBlue{
			Goal:  v.Goal,
			Total: tmpTotal,
		})
	}

	playGameTeamGoalAllUserRel, _ = g.playGameTeamGoalUserRelRepo.GetPlayGameTeamGoalUserRelByPlayId(ctx, playGameTeamGoalAllUserRelPlayId)
	res.GoalAllPlayId = playGameTeamGoalAllUserRelPlayId
	playGameTeamGoalRedAllUserRelMap := make(map[int64]struct {
		Goal  int64
		Total int64
	})
	playGameTeamGoalBlueAllUserRelMap := make(map[int64]struct {
		Goal  int64
		Total int64
	})

	for _, v := range playGameTeamGoalAllUserRel {
		if v.TeamId == game.RedTeamId {
			playGameTeamGoalRedAllUserRelMap[v.Goal] = struct {
				Goal  int64
				Total int64
			}{Goal: v.Goal, Total: v.Pay + playGameTeamGoalRedAllUserRelMap[v.Goal].Total}
		}

		if v.TeamId == game.BlueTeamId {
			playGameTeamGoalBlueAllUserRelMap[v.Goal] = struct {
				Goal  int64
				Total int64
			}{Goal: v.Goal, Total: v.Pay + playGameTeamGoalBlueAllUserRelMap[v.Goal].Total}
		}
	}

	for _, v := range playGameTeamGoalRedAllUserRelMap {
		tmpTotal := v.Total / base
		res.GoalAllRedTotal += tmpTotal
		res.GoalAllTotal += tmpTotal
		res.GoalAllRed = append(res.GoalAllRed, &v1.GameIndexStatisticsReply_GoalAllRed{
			Goal:  v.Goal,
			Total: tmpTotal,
		})
	}
	for _, v := range playGameTeamGoalBlueAllUserRelMap {
		tmpTotal := v.Total / base
		res.GoalAllBlueTotal += tmpTotal
		res.GoalAllTotal += tmpTotal
		res.GoalAllBlue = append(res.GoalAllBlue, &v1.GameIndexStatisticsReply_GoalAllBlue{
			Goal:  v.Goal,
			Total: tmpTotal,
		})
	}

	// 结果
	playGameTeamResultUserRel, _ = g.playGameTeamResultUserRelRepo.GetPlayGameTeamResultUserRelByPlayId(ctx, playGameTeamResultUserRelPlayId)
	res.ResultPlayId = playGameTeamResultUserRelPlayId
	for _, v := range playGameTeamResultUserRel {
		tmpTotal := v.Pay / base
		if "red" == v.Content {
			res.ResultRedTotal += tmpTotal
		} else if "blue" == v.Content {
			res.ResultBlueTotal += tmpTotal
		} else if "draw" == v.Content {
			res.ResultDrawTotal += tmpTotal
		}
		res.ResultTotal += tmpTotal
	}

	// 得分
	playGameScoreUserRel, _ = g.playGameScoreUserRelRepo.GetPlayGameScoreUserRelByPlayId(ctx, playGameScoreUserRelPlayId)
	res.ScorePlayId = playGameScoreUserRelPlayId
	playGameScoreUserRelMap := make(map[string]struct {
		Content string
		Total   int64
	})
	for _, v := range playGameScoreUserRel {
		playGameScoreUserRelMap[v.Content] = struct {
			Content string
			Total   int64
		}{Content: v.Content, Total: v.Pay + playGameScoreUserRelMap[v.Content].Total}
	}
	for _, v := range playGameScoreUserRelMap {
		tmpTotal := v.Total / base
		res.ScoreTotal += tmpTotal
		res.Score = append(res.Score, &v1.GameIndexStatisticsReply_Score{
			Content: v.Content,
			Total:   tmpTotal,
		})
	}

	return res, nil
}
