// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.0
// - protoc             v3.21.7
// source: api/admin/service/v1/admin.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationAdminCreatePlayGame = "/api.admin.service.v1.Admin/CreatePlayGame"
const OperationAdminCreatePlaySort = "/api.admin.service.v1.Admin/CreatePlaySort"
const OperationAdminGamePlayGrant = "/api.admin.service.v1.Admin/GamePlayGrant"
const OperationAdminGetPlayList = "/api.admin.service.v1.Admin/GetPlayList"
const OperationAdminGetPlayRelList = "/api.admin.service.v1.Admin/GetPlayRelList"
const OperationAdminGetRoomList = "/api.admin.service.v1.Admin/GetRoomList"
const OperationAdminGetRoomPlayList = "/api.admin.service.v1.Admin/GetRoomPlayList"
const OperationAdminSortPlayGrant = "/api.admin.service.v1.Admin/SortPlayGrant"

type AdminHTTPServer interface {
	CreatePlayGame(context.Context, *CreatePlayGameRequest) (*CreatePlayGameReply, error)
	CreatePlaySort(context.Context, *CreatePlaySortRequest) (*CreatePlaySortReply, error)
	GamePlayGrant(context.Context, *GamePlayGrantRequest) (*GamePlayGrantReply, error)
	GetPlayList(context.Context, *GetPlayListRequest) (*GetPlayListReply, error)
	GetPlayRelList(context.Context, *GetPlayRelListRequest) (*GetPlayRelListReply, error)
	GetRoomList(context.Context, *GetRoomListRequest) (*GetRoomListReply, error)
	GetRoomPlayList(context.Context, *GetRoomPlayListRequest) (*GetRoomPlayListReply, error)
	SortPlayGrant(context.Context, *SortPlayGrantRequest) (*SortPlayGrantReply, error)
}

func RegisterAdminHTTPServer(s *http.Server, srv AdminHTTPServer) {
	r := s.Route("/")
	r.POST("/api/goal_admin/play/game/grant", _Admin_GamePlayGrant0_HTTP_Handler(srv))
	r.POST("/api/goal_admin/play/sort/grant", _Admin_SortPlayGrant0_HTTP_Handler(srv))
	r.POST("/api/goal_admin/play/game", _Admin_CreatePlayGame0_HTTP_Handler(srv))
	r.POST("/api/goal_admin/play/sort", _Admin_CreatePlaySort0_HTTP_Handler(srv))
	r.GET("/api/goal_admin/game/{game_id}/play", _Admin_GetPlayList0_HTTP_Handler(srv))
	r.GET("/api/goal_admin/play/{play_id}/play_rel", _Admin_GetPlayRelList0_HTTP_Handler(srv))
	r.GET("/api/goal_admin/room_list", _Admin_GetRoomList0_HTTP_Handler(srv))
	r.GET("/api/goal_admin/room/{room_id}/play_list", _Admin_GetRoomPlayList0_HTTP_Handler(srv))
}

func _Admin_GamePlayGrant0_HTTP_Handler(srv AdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GamePlayGrantRequest
		if err := ctx.Bind(&in.SendBody); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAdminGamePlayGrant)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GamePlayGrant(ctx, req.(*GamePlayGrantRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GamePlayGrantReply)
		return ctx.Result(200, reply)
	}
}

func _Admin_SortPlayGrant0_HTTP_Handler(srv AdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SortPlayGrantRequest
		if err := ctx.Bind(&in.SendBody); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAdminSortPlayGrant)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SortPlayGrant(ctx, req.(*SortPlayGrantRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SortPlayGrantReply)
		return ctx.Result(200, reply)
	}
}

func _Admin_CreatePlayGame0_HTTP_Handler(srv AdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreatePlayGameRequest
		if err := ctx.Bind(&in.SendBody); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAdminCreatePlayGame)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreatePlayGame(ctx, req.(*CreatePlayGameRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreatePlayGameReply)
		return ctx.Result(200, reply)
	}
}

func _Admin_CreatePlaySort0_HTTP_Handler(srv AdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreatePlaySortRequest
		if err := ctx.Bind(&in.SendBody); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAdminCreatePlaySort)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreatePlaySort(ctx, req.(*CreatePlaySortRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreatePlaySortReply)
		return ctx.Result(200, reply)
	}
}

func _Admin_GetPlayList0_HTTP_Handler(srv AdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetPlayListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAdminGetPlayList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetPlayList(ctx, req.(*GetPlayListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetPlayListReply)
		return ctx.Result(200, reply)
	}
}

func _Admin_GetPlayRelList0_HTTP_Handler(srv AdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetPlayRelListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAdminGetPlayRelList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetPlayRelList(ctx, req.(*GetPlayRelListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetPlayRelListReply)
		return ctx.Result(200, reply)
	}
}

func _Admin_GetRoomList0_HTTP_Handler(srv AdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetRoomListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAdminGetRoomList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetRoomList(ctx, req.(*GetRoomListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetRoomListReply)
		return ctx.Result(200, reply)
	}
}

func _Admin_GetRoomPlayList0_HTTP_Handler(srv AdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetRoomPlayListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAdminGetRoomPlayList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetRoomPlayList(ctx, req.(*GetRoomPlayListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetRoomPlayListReply)
		return ctx.Result(200, reply)
	}
}

type AdminHTTPClient interface {
	CreatePlayGame(ctx context.Context, req *CreatePlayGameRequest, opts ...http.CallOption) (rsp *CreatePlayGameReply, err error)
	CreatePlaySort(ctx context.Context, req *CreatePlaySortRequest, opts ...http.CallOption) (rsp *CreatePlaySortReply, err error)
	GamePlayGrant(ctx context.Context, req *GamePlayGrantRequest, opts ...http.CallOption) (rsp *GamePlayGrantReply, err error)
	GetPlayList(ctx context.Context, req *GetPlayListRequest, opts ...http.CallOption) (rsp *GetPlayListReply, err error)
	GetPlayRelList(ctx context.Context, req *GetPlayRelListRequest, opts ...http.CallOption) (rsp *GetPlayRelListReply, err error)
	GetRoomList(ctx context.Context, req *GetRoomListRequest, opts ...http.CallOption) (rsp *GetRoomListReply, err error)
	GetRoomPlayList(ctx context.Context, req *GetRoomPlayListRequest, opts ...http.CallOption) (rsp *GetRoomPlayListReply, err error)
	SortPlayGrant(ctx context.Context, req *SortPlayGrantRequest, opts ...http.CallOption) (rsp *SortPlayGrantReply, err error)
}

type AdminHTTPClientImpl struct {
	cc *http.Client
}

func NewAdminHTTPClient(client *http.Client) AdminHTTPClient {
	return &AdminHTTPClientImpl{client}
}

func (c *AdminHTTPClientImpl) CreatePlayGame(ctx context.Context, in *CreatePlayGameRequest, opts ...http.CallOption) (*CreatePlayGameReply, error) {
	var out CreatePlayGameReply
	pattern := "/api/goal_admin/play/game"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAdminCreatePlayGame))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.SendBody, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AdminHTTPClientImpl) CreatePlaySort(ctx context.Context, in *CreatePlaySortRequest, opts ...http.CallOption) (*CreatePlaySortReply, error) {
	var out CreatePlaySortReply
	pattern := "/api/goal_admin/play/sort"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAdminCreatePlaySort))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.SendBody, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AdminHTTPClientImpl) GamePlayGrant(ctx context.Context, in *GamePlayGrantRequest, opts ...http.CallOption) (*GamePlayGrantReply, error) {
	var out GamePlayGrantReply
	pattern := "/api/goal_admin/play/game/grant"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAdminGamePlayGrant))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.SendBody, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AdminHTTPClientImpl) GetPlayList(ctx context.Context, in *GetPlayListRequest, opts ...http.CallOption) (*GetPlayListReply, error) {
	var out GetPlayListReply
	pattern := "/api/goal_admin/game/{game_id}/play"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAdminGetPlayList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AdminHTTPClientImpl) GetPlayRelList(ctx context.Context, in *GetPlayRelListRequest, opts ...http.CallOption) (*GetPlayRelListReply, error) {
	var out GetPlayRelListReply
	pattern := "/api/goal_admin/play/{play_id}/play_rel"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAdminGetPlayRelList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AdminHTTPClientImpl) GetRoomList(ctx context.Context, in *GetRoomListRequest, opts ...http.CallOption) (*GetRoomListReply, error) {
	var out GetRoomListReply
	pattern := "/api/goal_admin/room_list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAdminGetRoomList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AdminHTTPClientImpl) GetRoomPlayList(ctx context.Context, in *GetRoomPlayListRequest, opts ...http.CallOption) (*GetRoomPlayListReply, error) {
	var out GetRoomPlayListReply
	pattern := "/api/goal_admin/room/{room_id}/play_list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAdminGetRoomPlayList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AdminHTTPClientImpl) SortPlayGrant(ctx context.Context, in *SortPlayGrantRequest, opts ...http.CallOption) (*SortPlayGrantReply, error) {
	var out SortPlayGrantReply
	pattern := "/api/goal_admin/play/sort/grant"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAdminSortPlayGrant))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.SendBody, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

const OperationUserGetUser = "/api.admin.service.v1.User/GetUser"
const OperationUserGetUserBalanceRecord = "/api.admin.service.v1.User/GetUserBalanceRecord"
const OperationUserGetUserList = "/api.admin.service.v1.User/GetUserList"
const OperationUserGetUserProxyList = "/api.admin.service.v1.User/GetUserProxyList"
const OperationUserGetUserRecommendList = "/api.admin.service.v1.User/GetUserRecommendList"
const OperationUserUserDeposit = "/api.admin.service.v1.User/UserDeposit"

type UserHTTPServer interface {
	GetUser(context.Context, *GetUserRequest) (*GetUserReply, error)
	GetUserBalanceRecord(context.Context, *GetUserBalanceRecordRequest) (*GetUserBalanceRecordReply, error)
	GetUserList(context.Context, *GetUserListRequest) (*GetUserListReply, error)
	GetUserProxyList(context.Context, *GetUserProxyListRequest) (*GetUserProxyListReply, error)
	GetUserRecommendList(context.Context, *GetUserRecommendListRequest) (*GetUserRecommendListReply, error)
	UserDeposit(context.Context, *UserDepositRequest) (*UserDepositReply, error)
}

func RegisterUserHTTPServer(s *http.Server, srv UserHTTPServer) {
	r := s.Route("/")
	r.POST("/api/goal_admin/user/deposit", _User_UserDeposit0_HTTP_Handler(srv))
	r.GET("/api/goal_admin/user_list", _User_GetUserList0_HTTP_Handler(srv))
	r.GET("/api/goal_admin/user_proxy_list", _User_GetUserProxyList0_HTTP_Handler(srv))
	r.GET("/api/goal_admin/user/{user_id}/recommend_user_list", _User_GetUserRecommendList0_HTTP_Handler(srv))
	r.GET("/api/goal_admin/user/{user_id}", _User_GetUser0_HTTP_Handler(srv))
	r.GET("/api/goal_admin/user_balance_record", _User_GetUserBalanceRecord0_HTTP_Handler(srv))
}

func _User_UserDeposit0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UserDepositRequest
		if err := ctx.Bind(&in.SendBody); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserUserDeposit)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UserDeposit(ctx, req.(*UserDepositRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserDepositReply)
		return ctx.Result(200, reply)
	}
}

func _User_GetUserList0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserGetUserList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserList(ctx, req.(*GetUserListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserListReply)
		return ctx.Result(200, reply)
	}
}

func _User_GetUserProxyList0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserProxyListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserGetUserProxyList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserProxyList(ctx, req.(*GetUserProxyListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserProxyListReply)
		return ctx.Result(200, reply)
	}
}

func _User_GetUserRecommendList0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserRecommendListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserGetUserRecommendList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserRecommendList(ctx, req.(*GetUserRecommendListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserRecommendListReply)
		return ctx.Result(200, reply)
	}
}

func _User_GetUser0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserGetUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUser(ctx, req.(*GetUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserReply)
		return ctx.Result(200, reply)
	}
}

func _User_GetUserBalanceRecord0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserBalanceRecordRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserGetUserBalanceRecord)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserBalanceRecord(ctx, req.(*GetUserBalanceRecordRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserBalanceRecordReply)
		return ctx.Result(200, reply)
	}
}

type UserHTTPClient interface {
	GetUser(ctx context.Context, req *GetUserRequest, opts ...http.CallOption) (rsp *GetUserReply, err error)
	GetUserBalanceRecord(ctx context.Context, req *GetUserBalanceRecordRequest, opts ...http.CallOption) (rsp *GetUserBalanceRecordReply, err error)
	GetUserList(ctx context.Context, req *GetUserListRequest, opts ...http.CallOption) (rsp *GetUserListReply, err error)
	GetUserProxyList(ctx context.Context, req *GetUserProxyListRequest, opts ...http.CallOption) (rsp *GetUserProxyListReply, err error)
	GetUserRecommendList(ctx context.Context, req *GetUserRecommendListRequest, opts ...http.CallOption) (rsp *GetUserRecommendListReply, err error)
	UserDeposit(ctx context.Context, req *UserDepositRequest, opts ...http.CallOption) (rsp *UserDepositReply, err error)
}

type UserHTTPClientImpl struct {
	cc *http.Client
}

func NewUserHTTPClient(client *http.Client) UserHTTPClient {
	return &UserHTTPClientImpl{client}
}

func (c *UserHTTPClientImpl) GetUser(ctx context.Context, in *GetUserRequest, opts ...http.CallOption) (*GetUserReply, error) {
	var out GetUserReply
	pattern := "/api/goal_admin/user/{user_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserGetUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) GetUserBalanceRecord(ctx context.Context, in *GetUserBalanceRecordRequest, opts ...http.CallOption) (*GetUserBalanceRecordReply, error) {
	var out GetUserBalanceRecordReply
	pattern := "/api/goal_admin/user_balance_record"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserGetUserBalanceRecord))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) GetUserList(ctx context.Context, in *GetUserListRequest, opts ...http.CallOption) (*GetUserListReply, error) {
	var out GetUserListReply
	pattern := "/api/goal_admin/user_list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserGetUserList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) GetUserProxyList(ctx context.Context, in *GetUserProxyListRequest, opts ...http.CallOption) (*GetUserProxyListReply, error) {
	var out GetUserProxyListReply
	pattern := "/api/goal_admin/user_proxy_list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserGetUserProxyList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) GetUserRecommendList(ctx context.Context, in *GetUserRecommendListRequest, opts ...http.CallOption) (*GetUserRecommendListReply, error) {
	var out GetUserRecommendListReply
	pattern := "/api/goal_admin/user/{user_id}/recommend_user_list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserGetUserRecommendList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) UserDeposit(ctx context.Context, in *UserDepositRequest, opts ...http.CallOption) (*UserDepositReply, error) {
	var out UserDepositReply
	pattern := "/api/goal_admin/user/deposit"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserUserDeposit))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.SendBody, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

const OperationGameCreateGame = "/api.admin.service.v1.Game/CreateGame"
const OperationGameCreateSort = "/api.admin.service.v1.Game/CreateSort"
const OperationGameCreateTeam = "/api.admin.service.v1.Game/CreateTeam"
const OperationGameDisplayGameIndex = "/api.admin.service.v1.Game/DisplayGameIndex"
const OperationGameGetGame = "/api.admin.service.v1.Game/GetGame"
const OperationGameGetGameList = "/api.admin.service.v1.Game/GetGameList"
const OperationGameGetGameSortList = "/api.admin.service.v1.Game/GetGameSortList"
const OperationGameGetTeamList = "/api.admin.service.v1.Game/GetTeamList"
const OperationGameSaveDisplayGameIndex = "/api.admin.service.v1.Game/SaveDisplayGameIndex"
const OperationGameUpdateGame = "/api.admin.service.v1.Game/UpdateGame"

type GameHTTPServer interface {
	CreateGame(context.Context, *CreateGameRequest) (*CreateGameReply, error)
	CreateSort(context.Context, *CreateSortRequest) (*CreateSortReply, error)
	CreateTeam(context.Context, *CreateTeamRequest) (*CreateTeamReply, error)
	DisplayGameIndex(context.Context, *DisplayGameIndexRequest) (*DisplayGameIndexReply, error)
	GetGame(context.Context, *GetGameRequest) (*GetGameReply, error)
	GetGameList(context.Context, *GetGameListRequest) (*GetGameListReply, error)
	GetGameSortList(context.Context, *GetGameSortListRequest) (*GetGameSortListReply, error)
	GetTeamList(context.Context, *GetTeamListRequest) (*GetTeamListReply, error)
	SaveDisplayGameIndex(context.Context, *SaveDisplayGameIndexRequest) (*SaveDisplayGameIndexReply, error)
	UpdateGame(context.Context, *UpdateGameRequest) (*UpdateGameReply, error)
}

func RegisterGameHTTPServer(s *http.Server, srv GameHTTPServer) {
	r := s.Route("/")
	r.POST("/api/goal_admin/game/create", _Game_CreateGame0_HTTP_Handler(srv))
	r.POST("/api/goal_admin/game/update", _Game_UpdateGame0_HTTP_Handler(srv))
	r.GET("/api/goal_admin/game/{game_id}", _Game_GetGame0_HTTP_Handler(srv))
	r.GET("/api/goal_admin/games", _Game_GetGameList0_HTTP_Handler(srv))
	r.GET("/api/goal_admin/display_game/index", _Game_DisplayGameIndex0_HTTP_Handler(srv))
	r.POST("/api/goal_admin/display_game/index", _Game_SaveDisplayGameIndex0_HTTP_Handler(srv))
	r.POST("/api/goal_admin/sort/create", _Game_CreateSort0_HTTP_Handler(srv))
	r.GET("/api/goal_admin/sorts", _Game_GetGameSortList0_HTTP_Handler(srv))
	r.GET("/api/goal_admin/team/list", _Game_GetTeamList0_HTTP_Handler(srv))
	r.POST("/api/goal_admin/team/create", _Game_CreateTeam0_HTTP_Handler(srv))
}

func _Game_CreateGame0_HTTP_Handler(srv GameHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateGameRequest
		if err := ctx.Bind(&in.SendBody); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGameCreateGame)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateGame(ctx, req.(*CreateGameRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateGameReply)
		return ctx.Result(200, reply)
	}
}

func _Game_UpdateGame0_HTTP_Handler(srv GameHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateGameRequest
		if err := ctx.Bind(&in.SendBody); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGameUpdateGame)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateGame(ctx, req.(*UpdateGameRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateGameReply)
		return ctx.Result(200, reply)
	}
}

func _Game_GetGame0_HTTP_Handler(srv GameHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetGameRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGameGetGame)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetGame(ctx, req.(*GetGameRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetGameReply)
		return ctx.Result(200, reply)
	}
}

func _Game_GetGameList0_HTTP_Handler(srv GameHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetGameListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGameGetGameList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetGameList(ctx, req.(*GetGameListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetGameListReply)
		return ctx.Result(200, reply)
	}
}

func _Game_DisplayGameIndex0_HTTP_Handler(srv GameHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DisplayGameIndexRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGameDisplayGameIndex)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DisplayGameIndex(ctx, req.(*DisplayGameIndexRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DisplayGameIndexReply)
		return ctx.Result(200, reply)
	}
}

func _Game_SaveDisplayGameIndex0_HTTP_Handler(srv GameHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SaveDisplayGameIndexRequest
		if err := ctx.Bind(&in.SendBody); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGameSaveDisplayGameIndex)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SaveDisplayGameIndex(ctx, req.(*SaveDisplayGameIndexRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SaveDisplayGameIndexReply)
		return ctx.Result(200, reply)
	}
}

func _Game_CreateSort0_HTTP_Handler(srv GameHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateSortRequest
		if err := ctx.Bind(&in.SendBody); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGameCreateSort)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateSort(ctx, req.(*CreateSortRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateSortReply)
		return ctx.Result(200, reply)
	}
}

func _Game_GetGameSortList0_HTTP_Handler(srv GameHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetGameSortListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGameGetGameSortList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetGameSortList(ctx, req.(*GetGameSortListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetGameSortListReply)
		return ctx.Result(200, reply)
	}
}

func _Game_GetTeamList0_HTTP_Handler(srv GameHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetTeamListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGameGetTeamList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetTeamList(ctx, req.(*GetTeamListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetTeamListReply)
		return ctx.Result(200, reply)
	}
}

func _Game_CreateTeam0_HTTP_Handler(srv GameHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateTeamRequest
		if err := ctx.Bind(&in.SendBody); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGameCreateTeam)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateTeam(ctx, req.(*CreateTeamRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateTeamReply)
		return ctx.Result(200, reply)
	}
}

type GameHTTPClient interface {
	CreateGame(ctx context.Context, req *CreateGameRequest, opts ...http.CallOption) (rsp *CreateGameReply, err error)
	CreateSort(ctx context.Context, req *CreateSortRequest, opts ...http.CallOption) (rsp *CreateSortReply, err error)
	CreateTeam(ctx context.Context, req *CreateTeamRequest, opts ...http.CallOption) (rsp *CreateTeamReply, err error)
	DisplayGameIndex(ctx context.Context, req *DisplayGameIndexRequest, opts ...http.CallOption) (rsp *DisplayGameIndexReply, err error)
	GetGame(ctx context.Context, req *GetGameRequest, opts ...http.CallOption) (rsp *GetGameReply, err error)
	GetGameList(ctx context.Context, req *GetGameListRequest, opts ...http.CallOption) (rsp *GetGameListReply, err error)
	GetGameSortList(ctx context.Context, req *GetGameSortListRequest, opts ...http.CallOption) (rsp *GetGameSortListReply, err error)
	GetTeamList(ctx context.Context, req *GetTeamListRequest, opts ...http.CallOption) (rsp *GetTeamListReply, err error)
	SaveDisplayGameIndex(ctx context.Context, req *SaveDisplayGameIndexRequest, opts ...http.CallOption) (rsp *SaveDisplayGameIndexReply, err error)
	UpdateGame(ctx context.Context, req *UpdateGameRequest, opts ...http.CallOption) (rsp *UpdateGameReply, err error)
}

type GameHTTPClientImpl struct {
	cc *http.Client
}

func NewGameHTTPClient(client *http.Client) GameHTTPClient {
	return &GameHTTPClientImpl{client}
}

func (c *GameHTTPClientImpl) CreateGame(ctx context.Context, in *CreateGameRequest, opts ...http.CallOption) (*CreateGameReply, error) {
	var out CreateGameReply
	pattern := "/api/goal_admin/game/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationGameCreateGame))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.SendBody, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GameHTTPClientImpl) CreateSort(ctx context.Context, in *CreateSortRequest, opts ...http.CallOption) (*CreateSortReply, error) {
	var out CreateSortReply
	pattern := "/api/goal_admin/sort/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationGameCreateSort))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.SendBody, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GameHTTPClientImpl) CreateTeam(ctx context.Context, in *CreateTeamRequest, opts ...http.CallOption) (*CreateTeamReply, error) {
	var out CreateTeamReply
	pattern := "/api/goal_admin/team/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationGameCreateTeam))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.SendBody, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GameHTTPClientImpl) DisplayGameIndex(ctx context.Context, in *DisplayGameIndexRequest, opts ...http.CallOption) (*DisplayGameIndexReply, error) {
	var out DisplayGameIndexReply
	pattern := "/api/goal_admin/display_game/index"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationGameDisplayGameIndex))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GameHTTPClientImpl) GetGame(ctx context.Context, in *GetGameRequest, opts ...http.CallOption) (*GetGameReply, error) {
	var out GetGameReply
	pattern := "/api/goal_admin/game/{game_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationGameGetGame))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GameHTTPClientImpl) GetGameList(ctx context.Context, in *GetGameListRequest, opts ...http.CallOption) (*GetGameListReply, error) {
	var out GetGameListReply
	pattern := "/api/goal_admin/games"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationGameGetGameList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GameHTTPClientImpl) GetGameSortList(ctx context.Context, in *GetGameSortListRequest, opts ...http.CallOption) (*GetGameSortListReply, error) {
	var out GetGameSortListReply
	pattern := "/api/goal_admin/sorts"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationGameGetGameSortList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GameHTTPClientImpl) GetTeamList(ctx context.Context, in *GetTeamListRequest, opts ...http.CallOption) (*GetTeamListReply, error) {
	var out GetTeamListReply
	pattern := "/api/goal_admin/team/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationGameGetTeamList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GameHTTPClientImpl) SaveDisplayGameIndex(ctx context.Context, in *SaveDisplayGameIndexRequest, opts ...http.CallOption) (*SaveDisplayGameIndexReply, error) {
	var out SaveDisplayGameIndexReply
	pattern := "/api/goal_admin/display_game/index"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationGameSaveDisplayGameIndex))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.SendBody, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GameHTTPClientImpl) UpdateGame(ctx context.Context, in *UpdateGameRequest, opts ...http.CallOption) (*UpdateGameReply, error) {
	var out UpdateGameReply
	pattern := "/api/goal_admin/game/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationGameUpdateGame))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.SendBody, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
