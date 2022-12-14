// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: api/user/service/v1/user.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	EthAuthorize(ctx context.Context, in *EthAuthorizeRequest, opts ...grpc.CallOption) (*EthAuthorizeReply, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserReply, error)
	Deposit(ctx context.Context, in *DepositRequest, opts ...grpc.CallOption) (*DepositReply, error)
	Withdraw(ctx context.Context, in *WithdrawRequest, opts ...grpc.CallOption) (*WithdrawReply, error)
	GetUserRecommendList(ctx context.Context, in *GetUserRecommendListRequest, opts ...grpc.CallOption) (*GetUserRecommendListReply, error)
	GetUserDepositList(ctx context.Context, in *GetUserDepositListRequest, opts ...grpc.CallOption) (*GetUserDepositListReply, error)
	GetUserWithdrawList(ctx context.Context, in *GetUserWithdrawListRequest, opts ...grpc.CallOption) (*GetUserWithdrawListReply, error)
	CreateProxy(ctx context.Context, in *CreateProxyRequest, opts ...grpc.CallOption) (*CreateProxyReply, error)
	CreateDownProxy(ctx context.Context, in *CreateDownProxyRequest, opts ...grpc.CallOption) (*CreateDownProxyReply, error)
	GetUserProxyList(ctx context.Context, in *GetUserProxyListRequest, opts ...grpc.CallOption) (*GetUserProxyListReply, error)
	GetUserProxyConfigList(ctx context.Context, in *GetUserProxyConfigListRequest, opts ...grpc.CallOption) (*GetUserProxyConfigListReply, error)
	UserDeposit(ctx context.Context, in *UserDepositRequest, opts ...grpc.CallOption) (*UserDepositReply, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) EthAuthorize(ctx context.Context, in *EthAuthorizeRequest, opts ...grpc.CallOption) (*EthAuthorizeReply, error) {
	out := new(EthAuthorizeReply)
	err := c.cc.Invoke(ctx, "/api.user.service.v1.User/EthAuthorize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserReply, error) {
	out := new(GetUserReply)
	err := c.cc.Invoke(ctx, "/api.user.service.v1.User/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Deposit(ctx context.Context, in *DepositRequest, opts ...grpc.CallOption) (*DepositReply, error) {
	out := new(DepositReply)
	err := c.cc.Invoke(ctx, "/api.user.service.v1.User/Deposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Withdraw(ctx context.Context, in *WithdrawRequest, opts ...grpc.CallOption) (*WithdrawReply, error) {
	out := new(WithdrawReply)
	err := c.cc.Invoke(ctx, "/api.user.service.v1.User/Withdraw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserRecommendList(ctx context.Context, in *GetUserRecommendListRequest, opts ...grpc.CallOption) (*GetUserRecommendListReply, error) {
	out := new(GetUserRecommendListReply)
	err := c.cc.Invoke(ctx, "/api.user.service.v1.User/GetUserRecommendList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserDepositList(ctx context.Context, in *GetUserDepositListRequest, opts ...grpc.CallOption) (*GetUserDepositListReply, error) {
	out := new(GetUserDepositListReply)
	err := c.cc.Invoke(ctx, "/api.user.service.v1.User/GetUserDepositList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserWithdrawList(ctx context.Context, in *GetUserWithdrawListRequest, opts ...grpc.CallOption) (*GetUserWithdrawListReply, error) {
	out := new(GetUserWithdrawListReply)
	err := c.cc.Invoke(ctx, "/api.user.service.v1.User/GetUserWithdrawList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) CreateProxy(ctx context.Context, in *CreateProxyRequest, opts ...grpc.CallOption) (*CreateProxyReply, error) {
	out := new(CreateProxyReply)
	err := c.cc.Invoke(ctx, "/api.user.service.v1.User/CreateProxy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) CreateDownProxy(ctx context.Context, in *CreateDownProxyRequest, opts ...grpc.CallOption) (*CreateDownProxyReply, error) {
	out := new(CreateDownProxyReply)
	err := c.cc.Invoke(ctx, "/api.user.service.v1.User/CreateDownProxy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserProxyList(ctx context.Context, in *GetUserProxyListRequest, opts ...grpc.CallOption) (*GetUserProxyListReply, error) {
	out := new(GetUserProxyListReply)
	err := c.cc.Invoke(ctx, "/api.user.service.v1.User/GetUserProxyList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserProxyConfigList(ctx context.Context, in *GetUserProxyConfigListRequest, opts ...grpc.CallOption) (*GetUserProxyConfigListReply, error) {
	out := new(GetUserProxyConfigListReply)
	err := c.cc.Invoke(ctx, "/api.user.service.v1.User/GetUserProxyConfigList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserDeposit(ctx context.Context, in *UserDepositRequest, opts ...grpc.CallOption) (*UserDepositReply, error) {
	out := new(UserDepositReply)
	err := c.cc.Invoke(ctx, "/api.user.service.v1.User/UserDeposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	EthAuthorize(context.Context, *EthAuthorizeRequest) (*EthAuthorizeReply, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserReply, error)
	Deposit(context.Context, *DepositRequest) (*DepositReply, error)
	Withdraw(context.Context, *WithdrawRequest) (*WithdrawReply, error)
	GetUserRecommendList(context.Context, *GetUserRecommendListRequest) (*GetUserRecommendListReply, error)
	GetUserDepositList(context.Context, *GetUserDepositListRequest) (*GetUserDepositListReply, error)
	GetUserWithdrawList(context.Context, *GetUserWithdrawListRequest) (*GetUserWithdrawListReply, error)
	CreateProxy(context.Context, *CreateProxyRequest) (*CreateProxyReply, error)
	CreateDownProxy(context.Context, *CreateDownProxyRequest) (*CreateDownProxyReply, error)
	GetUserProxyList(context.Context, *GetUserProxyListRequest) (*GetUserProxyListReply, error)
	GetUserProxyConfigList(context.Context, *GetUserProxyConfigListRequest) (*GetUserProxyConfigListReply, error)
	UserDeposit(context.Context, *UserDepositRequest) (*UserDepositReply, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) EthAuthorize(context.Context, *EthAuthorizeRequest) (*EthAuthorizeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EthAuthorize not implemented")
}
func (UnimplementedUserServer) GetUser(context.Context, *GetUserRequest) (*GetUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServer) Deposit(context.Context, *DepositRequest) (*DepositReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deposit not implemented")
}
func (UnimplementedUserServer) Withdraw(context.Context, *WithdrawRequest) (*WithdrawReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Withdraw not implemented")
}
func (UnimplementedUserServer) GetUserRecommendList(context.Context, *GetUserRecommendListRequest) (*GetUserRecommendListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserRecommendList not implemented")
}
func (UnimplementedUserServer) GetUserDepositList(context.Context, *GetUserDepositListRequest) (*GetUserDepositListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserDepositList not implemented")
}
func (UnimplementedUserServer) GetUserWithdrawList(context.Context, *GetUserWithdrawListRequest) (*GetUserWithdrawListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserWithdrawList not implemented")
}
func (UnimplementedUserServer) CreateProxy(context.Context, *CreateProxyRequest) (*CreateProxyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProxy not implemented")
}
func (UnimplementedUserServer) CreateDownProxy(context.Context, *CreateDownProxyRequest) (*CreateDownProxyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDownProxy not implemented")
}
func (UnimplementedUserServer) GetUserProxyList(context.Context, *GetUserProxyListRequest) (*GetUserProxyListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProxyList not implemented")
}
func (UnimplementedUserServer) GetUserProxyConfigList(context.Context, *GetUserProxyConfigListRequest) (*GetUserProxyConfigListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProxyConfigList not implemented")
}
func (UnimplementedUserServer) UserDeposit(context.Context, *UserDepositRequest) (*UserDepositReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserDeposit not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_EthAuthorize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EthAuthorizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).EthAuthorize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.service.v1.User/EthAuthorize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).EthAuthorize(ctx, req.(*EthAuthorizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.service.v1.User/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Deposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Deposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.service.v1.User/Deposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Deposit(ctx, req.(*DepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Withdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithdrawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Withdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.service.v1.User/Withdraw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Withdraw(ctx, req.(*WithdrawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserRecommendList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRecommendListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserRecommendList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.service.v1.User/GetUserRecommendList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserRecommendList(ctx, req.(*GetUserRecommendListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserDepositList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserDepositListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserDepositList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.service.v1.User/GetUserDepositList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserDepositList(ctx, req.(*GetUserDepositListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserWithdrawList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserWithdrawListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserWithdrawList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.service.v1.User/GetUserWithdrawList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserWithdrawList(ctx, req.(*GetUserWithdrawListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_CreateProxy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProxyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CreateProxy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.service.v1.User/CreateProxy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CreateProxy(ctx, req.(*CreateProxyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_CreateDownProxy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDownProxyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CreateDownProxy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.service.v1.User/CreateDownProxy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CreateDownProxy(ctx, req.(*CreateDownProxyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserProxyList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserProxyListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserProxyList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.service.v1.User/GetUserProxyList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserProxyList(ctx, req.(*GetUserProxyListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserProxyConfigList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserProxyConfigListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserProxyConfigList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.service.v1.User/GetUserProxyConfigList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserProxyConfigList(ctx, req.(*GetUserProxyConfigListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserDeposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserDeposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.service.v1.User/UserDeposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserDeposit(ctx, req.(*UserDepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.user.service.v1.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EthAuthorize",
			Handler:    _User_EthAuthorize_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _User_GetUser_Handler,
		},
		{
			MethodName: "Deposit",
			Handler:    _User_Deposit_Handler,
		},
		{
			MethodName: "Withdraw",
			Handler:    _User_Withdraw_Handler,
		},
		{
			MethodName: "GetUserRecommendList",
			Handler:    _User_GetUserRecommendList_Handler,
		},
		{
			MethodName: "GetUserDepositList",
			Handler:    _User_GetUserDepositList_Handler,
		},
		{
			MethodName: "GetUserWithdrawList",
			Handler:    _User_GetUserWithdrawList_Handler,
		},
		{
			MethodName: "CreateProxy",
			Handler:    _User_CreateProxy_Handler,
		},
		{
			MethodName: "CreateDownProxy",
			Handler:    _User_CreateDownProxy_Handler,
		},
		{
			MethodName: "GetUserProxyList",
			Handler:    _User_GetUserProxyList_Handler,
		},
		{
			MethodName: "GetUserProxyConfigList",
			Handler:    _User_GetUserProxyConfigList_Handler,
		},
		{
			MethodName: "UserDeposit",
			Handler:    _User_UserDeposit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/user/service/v1/user.proto",
}
