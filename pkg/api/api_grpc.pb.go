// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: api.proto

package api

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AwesomeBotIIIClient is the client API for AwesomeBotIII service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AwesomeBotIIIClient interface {
	CreateUser(ctx context.Context, in *ChData, opts ...grpc.CallOption) (*empty.Empty, error)
	CreateSource(ctx context.Context, in *ChSrcData, opts ...grpc.CallOption) (*empty.Empty, error)
	GetSrcsByChat(ctx context.Context, in *ChData, opts ...grpc.CallOption) (*SrcSlcByChat, error)
	DeleteSource(ctx context.Context, in *ChSrcData, opts ...grpc.CallOption) (*empty.Empty, error)
	GetRSSBySource(ctx context.Context, in *ChSrcData, opts ...grpc.CallOption) (*RSSSlc, error)
}

type awesomeBotIIIClient struct {
	cc grpc.ClientConnInterface
}

func NewAwesomeBotIIIClient(cc grpc.ClientConnInterface) AwesomeBotIIIClient {
	return &awesomeBotIIIClient{cc}
}

func (c *awesomeBotIIIClient) CreateUser(ctx context.Context, in *ChData, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.awesomeBotIII/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *awesomeBotIIIClient) CreateSource(ctx context.Context, in *ChSrcData, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.awesomeBotIII/CreateSource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *awesomeBotIIIClient) GetSrcsByChat(ctx context.Context, in *ChData, opts ...grpc.CallOption) (*SrcSlcByChat, error) {
	out := new(SrcSlcByChat)
	err := c.cc.Invoke(ctx, "/api.awesomeBotIII/GetSrcsByChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *awesomeBotIIIClient) DeleteSource(ctx context.Context, in *ChSrcData, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.awesomeBotIII/DeleteSource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *awesomeBotIIIClient) GetRSSBySource(ctx context.Context, in *ChSrcData, opts ...grpc.CallOption) (*RSSSlc, error) {
	out := new(RSSSlc)
	err := c.cc.Invoke(ctx, "/api.awesomeBotIII/GetRSSBySource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AwesomeBotIIIServer is the server API for AwesomeBotIII service.
// All implementations must embed UnimplementedAwesomeBotIIIServer
// for forward compatibility
type AwesomeBotIIIServer interface {
	CreateUser(context.Context, *ChData) (*empty.Empty, error)
	CreateSource(context.Context, *ChSrcData) (*empty.Empty, error)
	GetSrcsByChat(context.Context, *ChData) (*SrcSlcByChat, error)
	DeleteSource(context.Context, *ChSrcData) (*empty.Empty, error)
	GetRSSBySource(context.Context, *ChSrcData) (*RSSSlc, error)
	mustEmbedUnimplementedAwesomeBotIIIServer()
}

// UnimplementedAwesomeBotIIIServer must be embedded to have forward compatible implementations.
type UnimplementedAwesomeBotIIIServer struct {
}

func (UnimplementedAwesomeBotIIIServer) CreateUser(context.Context, *ChData) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedAwesomeBotIIIServer) CreateSource(context.Context, *ChSrcData) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSource not implemented")
}
func (UnimplementedAwesomeBotIIIServer) GetSrcsByChat(context.Context, *ChData) (*SrcSlcByChat, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSrcsByChat not implemented")
}
func (UnimplementedAwesomeBotIIIServer) DeleteSource(context.Context, *ChSrcData) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSource not implemented")
}
func (UnimplementedAwesomeBotIIIServer) GetRSSBySource(context.Context, *ChSrcData) (*RSSSlc, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRSSBySource not implemented")
}
func (UnimplementedAwesomeBotIIIServer) mustEmbedUnimplementedAwesomeBotIIIServer() {}

// UnsafeAwesomeBotIIIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AwesomeBotIIIServer will
// result in compilation errors.
type UnsafeAwesomeBotIIIServer interface {
	mustEmbedUnimplementedAwesomeBotIIIServer()
}

func RegisterAwesomeBotIIIServer(s grpc.ServiceRegistrar, srv AwesomeBotIIIServer) {
	s.RegisterService(&AwesomeBotIII_ServiceDesc, srv)
}

func _AwesomeBotIII_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AwesomeBotIIIServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.awesomeBotIII/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AwesomeBotIIIServer).CreateUser(ctx, req.(*ChData))
	}
	return interceptor(ctx, in, info, handler)
}

func _AwesomeBotIII_CreateSource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChSrcData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AwesomeBotIIIServer).CreateSource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.awesomeBotIII/CreateSource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AwesomeBotIIIServer).CreateSource(ctx, req.(*ChSrcData))
	}
	return interceptor(ctx, in, info, handler)
}

func _AwesomeBotIII_GetSrcsByChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AwesomeBotIIIServer).GetSrcsByChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.awesomeBotIII/GetSrcsByChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AwesomeBotIIIServer).GetSrcsByChat(ctx, req.(*ChData))
	}
	return interceptor(ctx, in, info, handler)
}

func _AwesomeBotIII_DeleteSource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChSrcData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AwesomeBotIIIServer).DeleteSource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.awesomeBotIII/DeleteSource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AwesomeBotIIIServer).DeleteSource(ctx, req.(*ChSrcData))
	}
	return interceptor(ctx, in, info, handler)
}

func _AwesomeBotIII_GetRSSBySource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChSrcData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AwesomeBotIIIServer).GetRSSBySource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.awesomeBotIII/GetRSSBySource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AwesomeBotIIIServer).GetRSSBySource(ctx, req.(*ChSrcData))
	}
	return interceptor(ctx, in, info, handler)
}

// AwesomeBotIII_ServiceDesc is the grpc.ServiceDesc for AwesomeBotIII service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AwesomeBotIII_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.awesomeBotIII",
	HandlerType: (*AwesomeBotIIIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _AwesomeBotIII_CreateUser_Handler,
		},
		{
			MethodName: "CreateSource",
			Handler:    _AwesomeBotIII_CreateSource_Handler,
		},
		{
			MethodName: "GetSrcsByChat",
			Handler:    _AwesomeBotIII_GetSrcsByChat_Handler,
		},
		{
			MethodName: "DeleteSource",
			Handler:    _AwesomeBotIII_DeleteSource_Handler,
		},
		{
			MethodName: "GetRSSBySource",
			Handler:    _AwesomeBotIII_GetRSSBySource_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
