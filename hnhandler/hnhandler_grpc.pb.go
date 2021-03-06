// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package hnprotobuf

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

// UserGetterClient is the client API for UserGetter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserGetterClient interface {
	GetUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserReply, error)
}

type userGetterClient struct {
	cc grpc.ClientConnInterface
}

func NewUserGetterClient(cc grpc.ClientConnInterface) UserGetterClient {
	return &userGetterClient{cc}
}

func (c *userGetterClient) GetUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, "/hnhandler.UserGetter/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserGetterServer is the server API for UserGetter service.
// All implementations must embed UnimplementedUserGetterServer
// for forward compatibility
type UserGetterServer interface {
	GetUser(context.Context, *UserRequest) (*UserReply, error)
	mustEmbedUnimplementedUserGetterServer()
}

// UnimplementedUserGetterServer must be embedded to have forward compatible implementations.
type UnimplementedUserGetterServer struct {
}

func (UnimplementedUserGetterServer) GetUser(context.Context, *UserRequest) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserGetterServer) mustEmbedUnimplementedUserGetterServer() {}

// UnsafeUserGetterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserGetterServer will
// result in compilation errors.
type UnsafeUserGetterServer interface {
	mustEmbedUnimplementedUserGetterServer()
}

func RegisterUserGetterServer(s grpc.ServiceRegistrar, srv UserGetterServer) {
	s.RegisterService(&UserGetter_ServiceDesc, srv)
}

func _UserGetter_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserGetterServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hnhandler.UserGetter/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserGetterServer).GetUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserGetter_ServiceDesc is the grpc.ServiceDesc for UserGetter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserGetter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hnhandler.UserGetter",
	HandlerType: (*UserGetterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _UserGetter_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hnhandler/hnhandler.proto",
}

// ItemGetterClient is the client API for ItemGetter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ItemGetterClient interface {
	GetItem(ctx context.Context, in *ItemRequest, opts ...grpc.CallOption) (*ItemReply, error)
}

type itemGetterClient struct {
	cc grpc.ClientConnInterface
}

func NewItemGetterClient(cc grpc.ClientConnInterface) ItemGetterClient {
	return &itemGetterClient{cc}
}

func (c *itemGetterClient) GetItem(ctx context.Context, in *ItemRequest, opts ...grpc.CallOption) (*ItemReply, error) {
	out := new(ItemReply)
	err := c.cc.Invoke(ctx, "/hnhandler.ItemGetter/GetItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ItemGetterServer is the server API for ItemGetter service.
// All implementations must embed UnimplementedItemGetterServer
// for forward compatibility
type ItemGetterServer interface {
	GetItem(context.Context, *ItemRequest) (*ItemReply, error)
	mustEmbedUnimplementedItemGetterServer()
}

// UnimplementedItemGetterServer must be embedded to have forward compatible implementations.
type UnimplementedItemGetterServer struct {
}

func (UnimplementedItemGetterServer) GetItem(context.Context, *ItemRequest) (*ItemReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItem not implemented")
}
func (UnimplementedItemGetterServer) mustEmbedUnimplementedItemGetterServer() {}

// UnsafeItemGetterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ItemGetterServer will
// result in compilation errors.
type UnsafeItemGetterServer interface {
	mustEmbedUnimplementedItemGetterServer()
}

func RegisterItemGetterServer(s grpc.ServiceRegistrar, srv ItemGetterServer) {
	s.RegisterService(&ItemGetter_ServiceDesc, srv)
}

func _ItemGetter_GetItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemGetterServer).GetItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hnhandler.ItemGetter/GetItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemGetterServer).GetItem(ctx, req.(*ItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ItemGetter_ServiceDesc is the grpc.ServiceDesc for ItemGetter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ItemGetter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hnhandler.ItemGetter",
	HandlerType: (*ItemGetterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetItem",
			Handler:    _ItemGetter_GetItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hnhandler/hnhandler.proto",
}

// NewStoriesGetterClient is the client API for NewStoriesGetter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NewStoriesGetterClient interface {
	GetNewStories(ctx context.Context, in *NewStoriesRequest, opts ...grpc.CallOption) (*NewStoriesReply, error)
}

type newStoriesGetterClient struct {
	cc grpc.ClientConnInterface
}

func NewNewStoriesGetterClient(cc grpc.ClientConnInterface) NewStoriesGetterClient {
	return &newStoriesGetterClient{cc}
}

func (c *newStoriesGetterClient) GetNewStories(ctx context.Context, in *NewStoriesRequest, opts ...grpc.CallOption) (*NewStoriesReply, error) {
	out := new(NewStoriesReply)
	err := c.cc.Invoke(ctx, "/hnhandler.NewStoriesGetter/GetNewStories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NewStoriesGetterServer is the server API for NewStoriesGetter service.
// All implementations must embed UnimplementedNewStoriesGetterServer
// for forward compatibility
type NewStoriesGetterServer interface {
	GetNewStories(context.Context, *NewStoriesRequest) (*NewStoriesReply, error)
	mustEmbedUnimplementedNewStoriesGetterServer()
}

// UnimplementedNewStoriesGetterServer must be embedded to have forward compatible implementations.
type UnimplementedNewStoriesGetterServer struct {
}

func (UnimplementedNewStoriesGetterServer) GetNewStories(context.Context, *NewStoriesRequest) (*NewStoriesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNewStories not implemented")
}
func (UnimplementedNewStoriesGetterServer) mustEmbedUnimplementedNewStoriesGetterServer() {}

// UnsafeNewStoriesGetterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NewStoriesGetterServer will
// result in compilation errors.
type UnsafeNewStoriesGetterServer interface {
	mustEmbedUnimplementedNewStoriesGetterServer()
}

func RegisterNewStoriesGetterServer(s grpc.ServiceRegistrar, srv NewStoriesGetterServer) {
	s.RegisterService(&NewStoriesGetter_ServiceDesc, srv)
}

func _NewStoriesGetter_GetNewStories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewStoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewStoriesGetterServer).GetNewStories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hnhandler.NewStoriesGetter/GetNewStories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewStoriesGetterServer).GetNewStories(ctx, req.(*NewStoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NewStoriesGetter_ServiceDesc is the grpc.ServiceDesc for NewStoriesGetter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NewStoriesGetter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hnhandler.NewStoriesGetter",
	HandlerType: (*NewStoriesGetterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNewStories",
			Handler:    _NewStoriesGetter_GetNewStories_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hnhandler/hnhandler.proto",
}
