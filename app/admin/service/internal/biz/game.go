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
	UpdateDisplayGame(ctx context.Context, displayGame *DisplayGame, gameId int64) (*DisplayGame, error)
	CreateDisplayGame(ctx context.Context, gameId int64) (*DisplayGame, error)
}

type GameUseCase struct {
	gameRepo GameRepo
	log      *log.Helper
}

func NewGameUseCase(repo GameRepo, logger log.Logger) *GameUseCase {
	return &GameUseCase{gameRepo: repo, log: log.NewHelper(logger)}
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
	disPlayGame, err := g.gameRepo.GetDisplayGame()
	if err != nil {
		return nil, err
	}

	return &v1.DisplayGameIndexReply{
		GameId: disPlayGame.GameId,
	}, nil
}

func (g *GameUseCase) SaveDisplayGameIndex(ctx context.Context, req *v1.SaveDisplayGameIndexRequest) (*v1.SaveDisplayGameIndexReply, error) {
	disPlayGame, err := g.gameRepo.GetDisplayGame()
	if nil == disPlayGame {
		disPlayGame, err = g.gameRepo.CreateDisplayGame(ctx, req.SendBody.GameId)
	} else {
		disPlayGame, err = g.gameRepo.UpdateDisplayGame(ctx, disPlayGame, req.SendBody.GameId)
	}

	if err != nil {
		return nil, err
	}

	return &v1.SaveDisplayGameIndexReply{
		GameId: disPlayGame.GameId,
	}, nil
}
