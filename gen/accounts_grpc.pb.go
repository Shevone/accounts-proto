// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: accounts.proto

package accounts

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

// AccountsClient is the client API for Accounts service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountsClient interface {
	EditProfile(ctx context.Context, in *EditRequest, opts ...grpc.CallOption) (*EditResponse, error)
	GetUserData(ctx context.Context, in *GetUserRequests, opts ...grpc.CallOption) (*GetDataResponse, error)
	GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error)
}

type accountsClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountsClient(cc grpc.ClientConnInterface) AccountsClient {
	return &accountsClient{cc}
}

func (c *accountsClient) EditProfile(ctx context.Context, in *EditRequest, opts ...grpc.CallOption) (*EditResponse, error) {
	out := new(EditResponse)
	err := c.cc.Invoke(ctx, "/accounts.Accounts/EditProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) GetUserData(ctx context.Context, in *GetUserRequests, opts ...grpc.CallOption) (*GetDataResponse, error) {
	out := new(GetDataResponse)
	err := c.cc.Invoke(ctx, "/accounts.Accounts/GetUserData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsClient) GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error) {
	out := new(GetUsersResponse)
	err := c.cc.Invoke(ctx, "/accounts.Accounts/GetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountsServer is the server API for Accounts service.
// All implementations must embed UnimplementedAccountsServer
// for forward compatibility
type AccountsServer interface {
	EditProfile(context.Context, *EditRequest) (*EditResponse, error)
	GetUserData(context.Context, *GetUserRequests) (*GetDataResponse, error)
	GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error)
	mustEmbedUnimplementedAccountsServer()
}

// UnimplementedAccountsServer must be embedded to have forward compatible implementations.
type UnimplementedAccountsServer struct {
}

func (UnimplementedAccountsServer) EditProfile(context.Context, *EditRequest) (*EditResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditProfile not implemented")
}
func (UnimplementedAccountsServer) GetUserData(context.Context, *GetUserRequests) (*GetDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserData not implemented")
}
func (UnimplementedAccountsServer) GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedAccountsServer) mustEmbedUnimplementedAccountsServer() {}

// UnsafeAccountsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountsServer will
// result in compilation errors.
type UnsafeAccountsServer interface {
	mustEmbedUnimplementedAccountsServer()
}

func RegisterAccountsServer(s grpc.ServiceRegistrar, srv AccountsServer) {
	s.RegisterService(&Accounts_ServiceDesc, srv)
}

func _Accounts_EditProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).EditProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accounts.Accounts/EditProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).EditProfile(ctx, req.(*EditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_GetUserData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequests)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).GetUserData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accounts.Accounts/GetUserData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).GetUserData(ctx, req.(*GetUserRequests))
	}
	return interceptor(ctx, in, info, handler)
}

func _Accounts_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accounts.Accounts/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).GetUsers(ctx, req.(*GetUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Accounts_ServiceDesc is the grpc.ServiceDesc for Accounts service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Accounts_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "accounts.Accounts",
	HandlerType: (*AccountsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EditProfile",
			Handler:    _Accounts_EditProfile_Handler,
		},
		{
			MethodName: "GetUserData",
			Handler:    _Accounts_GetUserData_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _Accounts_GetUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accounts.proto",
}
