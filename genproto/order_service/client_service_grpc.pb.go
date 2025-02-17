// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.0
// source: client_service.proto

package order_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ClientServiceClient is the client API for ClientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientServiceClient interface {
	CreateMerchant(ctx context.Context, in *CreateMerchantReq, opts ...grpc.CallOption) (*CreateMerchantResponse, error)
	GetSingleMerchant(ctx context.Context, in *MerchantId, opts ...grpc.CallOption) (*GetSingleMerchantResponse, error)
	GetListMerchants(ctx context.Context, in *GetListMerchantsReq, opts ...grpc.CallOption) (*GetListMerchantsResponse, error)
	UpdateMerchant(ctx context.Context, in *UpdateMerchantReq, opts ...grpc.CallOption) (*UpdateMerchantResponse, error)
	DeleteMerchant(ctx context.Context, in *MerchantId, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*CreateUserResponse, error)
	GetSingleUser(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*GetSingleUserResponse, error)
	GetListUsers(ctx context.Context, in *GetListUsersReq, opts ...grpc.CallOption) (*GetListUsersResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	DeleteUser(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type clientServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientServiceClient(cc grpc.ClientConnInterface) ClientServiceClient {
	return &clientServiceClient{cc}
}

func (c *clientServiceClient) CreateMerchant(ctx context.Context, in *CreateMerchantReq, opts ...grpc.CallOption) (*CreateMerchantResponse, error) {
	out := new(CreateMerchantResponse)
	err := c.cc.Invoke(ctx, "/order_service.ClientService/CreateMerchant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) GetSingleMerchant(ctx context.Context, in *MerchantId, opts ...grpc.CallOption) (*GetSingleMerchantResponse, error) {
	out := new(GetSingleMerchantResponse)
	err := c.cc.Invoke(ctx, "/order_service.ClientService/GetSingleMerchant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) GetListMerchants(ctx context.Context, in *GetListMerchantsReq, opts ...grpc.CallOption) (*GetListMerchantsResponse, error) {
	out := new(GetListMerchantsResponse)
	err := c.cc.Invoke(ctx, "/order_service.ClientService/GetListMerchants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) UpdateMerchant(ctx context.Context, in *UpdateMerchantReq, opts ...grpc.CallOption) (*UpdateMerchantResponse, error) {
	out := new(UpdateMerchantResponse)
	err := c.cc.Invoke(ctx, "/order_service.ClientService/UpdateMerchant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) DeleteMerchant(ctx context.Context, in *MerchantId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/order_service.ClientService/DeleteMerchant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) CreateUser(ctx context.Context, in *CreateUserReq, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/order_service.ClientService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) GetSingleUser(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*GetSingleUserResponse, error) {
	out := new(GetSingleUserResponse)
	err := c.cc.Invoke(ctx, "/order_service.ClientService/GetSingleUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) GetListUsers(ctx context.Context, in *GetListUsersReq, opts ...grpc.CallOption) (*GetListUsersResponse, error) {
	out := new(GetListUsersResponse)
	err := c.cc.Invoke(ctx, "/order_service.ClientService/GetListUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, "/order_service.ClientService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) DeleteUser(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/order_service.ClientService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientServiceServer is the server API for ClientService service.
// All implementations must embed UnimplementedClientServiceServer
// for forward compatibility
type ClientServiceServer interface {
	CreateMerchant(context.Context, *CreateMerchantReq) (*CreateMerchantResponse, error)
	GetSingleMerchant(context.Context, *MerchantId) (*GetSingleMerchantResponse, error)
	GetListMerchants(context.Context, *GetListMerchantsReq) (*GetListMerchantsResponse, error)
	UpdateMerchant(context.Context, *UpdateMerchantReq) (*UpdateMerchantResponse, error)
	DeleteMerchant(context.Context, *MerchantId) (*emptypb.Empty, error)
	CreateUser(context.Context, *CreateUserReq) (*CreateUserResponse, error)
	GetSingleUser(context.Context, *UserId) (*GetSingleUserResponse, error)
	GetListUsers(context.Context, *GetListUsersReq) (*GetListUsersResponse, error)
	UpdateUser(context.Context, *UpdateUserReq) (*UpdateUserResponse, error)
	DeleteUser(context.Context, *UserId) (*emptypb.Empty, error)
	mustEmbedUnimplementedClientServiceServer()
}

// UnimplementedClientServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClientServiceServer struct {
}

func (UnimplementedClientServiceServer) CreateMerchant(context.Context, *CreateMerchantReq) (*CreateMerchantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMerchant not implemented")
}
func (UnimplementedClientServiceServer) GetSingleMerchant(context.Context, *MerchantId) (*GetSingleMerchantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSingleMerchant not implemented")
}
func (UnimplementedClientServiceServer) GetListMerchants(context.Context, *GetListMerchantsReq) (*GetListMerchantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListMerchants not implemented")
}
func (UnimplementedClientServiceServer) UpdateMerchant(context.Context, *UpdateMerchantReq) (*UpdateMerchantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMerchant not implemented")
}
func (UnimplementedClientServiceServer) DeleteMerchant(context.Context, *MerchantId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMerchant not implemented")
}
func (UnimplementedClientServiceServer) CreateUser(context.Context, *CreateUserReq) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedClientServiceServer) GetSingleUser(context.Context, *UserId) (*GetSingleUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSingleUser not implemented")
}
func (UnimplementedClientServiceServer) GetListUsers(context.Context, *GetListUsersReq) (*GetListUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListUsers not implemented")
}
func (UnimplementedClientServiceServer) UpdateUser(context.Context, *UpdateUserReq) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedClientServiceServer) DeleteUser(context.Context, *UserId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedClientServiceServer) mustEmbedUnimplementedClientServiceServer() {}

// UnsafeClientServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientServiceServer will
// result in compilation errors.
type UnsafeClientServiceServer interface {
	mustEmbedUnimplementedClientServiceServer()
}

func RegisterClientServiceServer(s grpc.ServiceRegistrar, srv ClientServiceServer) {
	s.RegisterService(&ClientService_ServiceDesc, srv)
}

func _ClientService_CreateMerchant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMerchantReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).CreateMerchant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order_service.ClientService/CreateMerchant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).CreateMerchant(ctx, req.(*CreateMerchantReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_GetSingleMerchant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MerchantId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).GetSingleMerchant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order_service.ClientService/GetSingleMerchant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).GetSingleMerchant(ctx, req.(*MerchantId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_GetListMerchants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListMerchantsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).GetListMerchants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order_service.ClientService/GetListMerchants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).GetListMerchants(ctx, req.(*GetListMerchantsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_UpdateMerchant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMerchantReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).UpdateMerchant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order_service.ClientService/UpdateMerchant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).UpdateMerchant(ctx, req.(*UpdateMerchantReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_DeleteMerchant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MerchantId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).DeleteMerchant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order_service.ClientService/DeleteMerchant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).DeleteMerchant(ctx, req.(*MerchantId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order_service.ClientService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).CreateUser(ctx, req.(*CreateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_GetSingleUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).GetSingleUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order_service.ClientService/GetSingleUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).GetSingleUser(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_GetListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListUsersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).GetListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order_service.ClientService/GetListUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).GetListUsers(ctx, req.(*GetListUsersReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order_service.ClientService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).UpdateUser(ctx, req.(*UpdateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order_service.ClientService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).DeleteUser(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientService_ServiceDesc is the grpc.ServiceDesc for ClientService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order_service.ClientService",
	HandlerType: (*ClientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMerchant",
			Handler:    _ClientService_CreateMerchant_Handler,
		},
		{
			MethodName: "GetSingleMerchant",
			Handler:    _ClientService_GetSingleMerchant_Handler,
		},
		{
			MethodName: "GetListMerchants",
			Handler:    _ClientService_GetListMerchants_Handler,
		},
		{
			MethodName: "UpdateMerchant",
			Handler:    _ClientService_UpdateMerchant_Handler,
		},
		{
			MethodName: "DeleteMerchant",
			Handler:    _ClientService_DeleteMerchant_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _ClientService_CreateUser_Handler,
		},
		{
			MethodName: "GetSingleUser",
			Handler:    _ClientService_GetSingleUser_Handler,
		},
		{
			MethodName: "GetListUsers",
			Handler:    _ClientService_GetListUsers_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _ClientService_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _ClientService_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "client_service.proto",
}
