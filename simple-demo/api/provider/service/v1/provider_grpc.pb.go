// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: api/provider/service/v1/provider.proto

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

// ProviderClient is the client API for Provider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProviderClient interface {
	CreateProvider(ctx context.Context, in *CreateProviderRequest, opts ...grpc.CallOption) (*CreateProviderReply, error)
	UpdateProvider(ctx context.Context, in *UpdateProviderRequest, opts ...grpc.CallOption) (*UpdateProviderReply, error)
	DeleteProvider(ctx context.Context, in *DeleteProviderRequest, opts ...grpc.CallOption) (*DeleteProviderReply, error)
	GetProvider(ctx context.Context, in *GetProviderRequest, opts ...grpc.CallOption) (*GetProviderReply, error)
	ListProvider(ctx context.Context, in *ListProviderRequest, opts ...grpc.CallOption) (*ListProviderReply, error)
}

type providerClient struct {
	cc grpc.ClientConnInterface
}

func NewProviderClient(cc grpc.ClientConnInterface) ProviderClient {
	return &providerClient{cc}
}

func (c *providerClient) CreateProvider(ctx context.Context, in *CreateProviderRequest, opts ...grpc.CallOption) (*CreateProviderReply, error) {
	out := new(CreateProviderReply)
	err := c.cc.Invoke(ctx, "/api.provider.service.v1.Provider/CreateProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerClient) UpdateProvider(ctx context.Context, in *UpdateProviderRequest, opts ...grpc.CallOption) (*UpdateProviderReply, error) {
	out := new(UpdateProviderReply)
	err := c.cc.Invoke(ctx, "/api.provider.service.v1.Provider/UpdateProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerClient) DeleteProvider(ctx context.Context, in *DeleteProviderRequest, opts ...grpc.CallOption) (*DeleteProviderReply, error) {
	out := new(DeleteProviderReply)
	err := c.cc.Invoke(ctx, "/api.provider.service.v1.Provider/DeleteProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerClient) GetProvider(ctx context.Context, in *GetProviderRequest, opts ...grpc.CallOption) (*GetProviderReply, error) {
	out := new(GetProviderReply)
	err := c.cc.Invoke(ctx, "/api.provider.service.v1.Provider/GetProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerClient) ListProvider(ctx context.Context, in *ListProviderRequest, opts ...grpc.CallOption) (*ListProviderReply, error) {
	out := new(ListProviderReply)
	err := c.cc.Invoke(ctx, "/api.provider.service.v1.Provider/ListProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProviderServer is the server API for Provider service.
// All implementations must embed UnimplementedProviderServer
// for forward compatibility
type ProviderServer interface {
	CreateProvider(context.Context, *CreateProviderRequest) (*CreateProviderReply, error)
	UpdateProvider(context.Context, *UpdateProviderRequest) (*UpdateProviderReply, error)
	DeleteProvider(context.Context, *DeleteProviderRequest) (*DeleteProviderReply, error)
	GetProvider(context.Context, *GetProviderRequest) (*GetProviderReply, error)
	ListProvider(context.Context, *ListProviderRequest) (*ListProviderReply, error)
	mustEmbedUnimplementedProviderServer()
}

// UnimplementedProviderServer must be embedded to have forward compatible implementations.
type UnimplementedProviderServer struct {
}

func (UnimplementedProviderServer) CreateProvider(context.Context, *CreateProviderRequest) (*CreateProviderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProvider not implemented")
}
func (UnimplementedProviderServer) UpdateProvider(context.Context, *UpdateProviderRequest) (*UpdateProviderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProvider not implemented")
}
func (UnimplementedProviderServer) DeleteProvider(context.Context, *DeleteProviderRequest) (*DeleteProviderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProvider not implemented")
}
func (UnimplementedProviderServer) GetProvider(context.Context, *GetProviderRequest) (*GetProviderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProvider not implemented")
}
func (UnimplementedProviderServer) ListProvider(context.Context, *ListProviderRequest) (*ListProviderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProvider not implemented")
}
func (UnimplementedProviderServer) mustEmbedUnimplementedProviderServer() {}

// UnsafeProviderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProviderServer will
// result in compilation errors.
type UnsafeProviderServer interface {
	mustEmbedUnimplementedProviderServer()
}

func RegisterProviderServer(s grpc.ServiceRegistrar, srv ProviderServer) {
	s.RegisterService(&Provider_ServiceDesc, srv)
}

func _Provider_CreateProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).CreateProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.provider.service.v1.Provider/CreateProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).CreateProvider(ctx, req.(*CreateProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provider_UpdateProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).UpdateProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.provider.service.v1.Provider/UpdateProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).UpdateProvider(ctx, req.(*UpdateProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provider_DeleteProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).DeleteProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.provider.service.v1.Provider/DeleteProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).DeleteProvider(ctx, req.(*DeleteProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provider_GetProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).GetProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.provider.service.v1.Provider/GetProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).GetProvider(ctx, req.(*GetProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provider_ListProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServer).ListProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.provider.service.v1.Provider/ListProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServer).ListProvider(ctx, req.(*ListProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Provider_ServiceDesc is the grpc.ServiceDesc for Provider service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Provider_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.provider.service.v1.Provider",
	HandlerType: (*ProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProvider",
			Handler:    _Provider_CreateProvider_Handler,
		},
		{
			MethodName: "UpdateProvider",
			Handler:    _Provider_UpdateProvider_Handler,
		},
		{
			MethodName: "DeleteProvider",
			Handler:    _Provider_DeleteProvider_Handler,
		},
		{
			MethodName: "GetProvider",
			Handler:    _Provider_GetProvider_Handler,
		},
		{
			MethodName: "ListProvider",
			Handler:    _Provider_ListProvider_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/provider/service/v1/provider.proto",
}
