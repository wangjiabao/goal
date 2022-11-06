// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.0
// - protoc             v3.21.7
// source: api/game/service/v1/game.proto

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

const OperationGameDisplayGame = "/api.game.service.v1.Game/DisplayGame"
const OperationGameGetGameList = "/api.game.service.v1.Game/GetGameList"
const OperationGameGetGameSortList = "/api.game.service.v1.Game/GetGameSortList"

type GameHTTPServer interface {
	DisplayGame(context.Context, *DisplayGameRequest) (*DisplayGameReply, error)
	GetGameList(context.Context, *GetGameListRequest) (*GetGameListReply, error)
	GetGameSortList(context.Context, *GetGameSortListRequest) (*GetGameSortListReply, error)
}

func RegisterGameHTTPServer(s *http.Server, srv GameHTTPServer) {
	r := s.Route("/")
	r.GET("/api/display_game/{type}", _Game_DisplayGame0_HTTP_Handler(srv))
	r.GET("/api/games", _Game_GetGameList0_HTTP_Handler(srv))
	r.GET("/api/game/sorts", _Game_GetGameSortList0_HTTP_Handler(srv))
}

func _Game_DisplayGame0_HTTP_Handler(srv GameHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DisplayGameRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGameDisplayGame)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DisplayGame(ctx, req.(*DisplayGameRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DisplayGameReply)
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

type GameHTTPClient interface {
	DisplayGame(ctx context.Context, req *DisplayGameRequest, opts ...http.CallOption) (rsp *DisplayGameReply, err error)
	GetGameList(ctx context.Context, req *GetGameListRequest, opts ...http.CallOption) (rsp *GetGameListReply, err error)
	GetGameSortList(ctx context.Context, req *GetGameSortListRequest, opts ...http.CallOption) (rsp *GetGameSortListReply, err error)
}

type GameHTTPClientImpl struct {
	cc *http.Client
}

func NewGameHTTPClient(client *http.Client) GameHTTPClient {
	return &GameHTTPClientImpl{client}
}

func (c *GameHTTPClientImpl) DisplayGame(ctx context.Context, in *DisplayGameRequest, opts ...http.CallOption) (*DisplayGameReply, error) {
	var out DisplayGameReply
	pattern := "/api/display_game/{type}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationGameDisplayGame))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GameHTTPClientImpl) GetGameList(ctx context.Context, in *GetGameListRequest, opts ...http.CallOption) (*GetGameListReply, error) {
	var out GetGameListReply
	pattern := "/api/games"
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
	pattern := "/api/game/sorts"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationGameGetGameSortList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
