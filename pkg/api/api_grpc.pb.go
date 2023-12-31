// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0
// source: api/api.proto

package api

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

const (
	Api_StatV1_FullMethodName            = "/api.Api/StatV1"
	Api_AddCategoryV1_FullMethodName     = "/api.Api/AddCategoryV1"
	Api_CategoryListV1_FullMethodName    = "/api.Api/CategoryListV1"
	Api_DeleteCategoryV1_FullMethodName  = "/api.Api/DeleteCategoryV1"
	Api_WasteListV1_FullMethodName       = "/api.Api/WasteListV1"
	Api_AddWasteListV1_FullMethodName    = "/api.Api/AddWasteListV1"
	Api_DeleteWasteListV1_FullMethodName = "/api.Api/DeleteWasteListV1"
)

// ApiClient is the client API for Api service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiClient interface {
	StatV1(ctx context.Context, in *StatisticV1Request, opts ...grpc.CallOption) (*StatisticV1Response, error)
	AddCategoryV1(ctx context.Context, in *AddCategoryV1Request, opts ...grpc.CallOption) (*AddCategoryV1Response, error)
	CategoryListV1(ctx context.Context, in *CategoryListV1Request, opts ...grpc.CallOption) (*CategoryListV1Response, error)
	DeleteCategoryV1(ctx context.Context, in *DeleteCategoryV1Request, opts ...grpc.CallOption) (*DeleteCategoryV1Response, error)
	WasteListV1(ctx context.Context, in *WasteListV1Request, opts ...grpc.CallOption) (*WasteListV1Response, error)
	AddWasteListV1(ctx context.Context, in *AddWasteListV1Request, opts ...grpc.CallOption) (*AddWasteListV1Response, error)
	DeleteWasteListV1(ctx context.Context, in *DeleteWasteListV1Request, opts ...grpc.CallOption) (*DeleteWasteListV1Response, error)
}

type apiClient struct {
	cc grpc.ClientConnInterface
}

func NewApiClient(cc grpc.ClientConnInterface) ApiClient {
	return &apiClient{cc}
}

func (c *apiClient) StatV1(ctx context.Context, in *StatisticV1Request, opts ...grpc.CallOption) (*StatisticV1Response, error) {
	out := new(StatisticV1Response)
	err := c.cc.Invoke(ctx, Api_StatV1_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) AddCategoryV1(ctx context.Context, in *AddCategoryV1Request, opts ...grpc.CallOption) (*AddCategoryV1Response, error) {
	out := new(AddCategoryV1Response)
	err := c.cc.Invoke(ctx, Api_AddCategoryV1_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) CategoryListV1(ctx context.Context, in *CategoryListV1Request, opts ...grpc.CallOption) (*CategoryListV1Response, error) {
	out := new(CategoryListV1Response)
	err := c.cc.Invoke(ctx, Api_CategoryListV1_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) DeleteCategoryV1(ctx context.Context, in *DeleteCategoryV1Request, opts ...grpc.CallOption) (*DeleteCategoryV1Response, error) {
	out := new(DeleteCategoryV1Response)
	err := c.cc.Invoke(ctx, Api_DeleteCategoryV1_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) WasteListV1(ctx context.Context, in *WasteListV1Request, opts ...grpc.CallOption) (*WasteListV1Response, error) {
	out := new(WasteListV1Response)
	err := c.cc.Invoke(ctx, Api_WasteListV1_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) AddWasteListV1(ctx context.Context, in *AddWasteListV1Request, opts ...grpc.CallOption) (*AddWasteListV1Response, error) {
	out := new(AddWasteListV1Response)
	err := c.cc.Invoke(ctx, Api_AddWasteListV1_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) DeleteWasteListV1(ctx context.Context, in *DeleteWasteListV1Request, opts ...grpc.CallOption) (*DeleteWasteListV1Response, error) {
	out := new(DeleteWasteListV1Response)
	err := c.cc.Invoke(ctx, Api_DeleteWasteListV1_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiServer is the server API for Api service.
// All implementations must embed UnimplementedApiServer
// for forward compatibility
type ApiServer interface {
	StatV1(context.Context, *StatisticV1Request) (*StatisticV1Response, error)
	AddCategoryV1(context.Context, *AddCategoryV1Request) (*AddCategoryV1Response, error)
	CategoryListV1(context.Context, *CategoryListV1Request) (*CategoryListV1Response, error)
	DeleteCategoryV1(context.Context, *DeleteCategoryV1Request) (*DeleteCategoryV1Response, error)
	WasteListV1(context.Context, *WasteListV1Request) (*WasteListV1Response, error)
	AddWasteListV1(context.Context, *AddWasteListV1Request) (*AddWasteListV1Response, error)
	DeleteWasteListV1(context.Context, *DeleteWasteListV1Request) (*DeleteWasteListV1Response, error)
	mustEmbedUnimplementedApiServer()
}

// UnimplementedApiServer must be embedded to have forward compatible implementations.
type UnimplementedApiServer struct {
}

func (UnimplementedApiServer) StatV1(context.Context, *StatisticV1Request) (*StatisticV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StatV1 not implemented")
}
func (UnimplementedApiServer) AddCategoryV1(context.Context, *AddCategoryV1Request) (*AddCategoryV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCategoryV1 not implemented")
}
func (UnimplementedApiServer) CategoryListV1(context.Context, *CategoryListV1Request) (*CategoryListV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CategoryListV1 not implemented")
}
func (UnimplementedApiServer) DeleteCategoryV1(context.Context, *DeleteCategoryV1Request) (*DeleteCategoryV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCategoryV1 not implemented")
}
func (UnimplementedApiServer) WasteListV1(context.Context, *WasteListV1Request) (*WasteListV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WasteListV1 not implemented")
}
func (UnimplementedApiServer) AddWasteListV1(context.Context, *AddWasteListV1Request) (*AddWasteListV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddWasteListV1 not implemented")
}
func (UnimplementedApiServer) DeleteWasteListV1(context.Context, *DeleteWasteListV1Request) (*DeleteWasteListV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWasteListV1 not implemented")
}
func (UnimplementedApiServer) mustEmbedUnimplementedApiServer() {}

// UnsafeApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiServer will
// result in compilation errors.
type UnsafeApiServer interface {
	mustEmbedUnimplementedApiServer()
}

func RegisterApiServer(s grpc.ServiceRegistrar, srv ApiServer) {
	s.RegisterService(&Api_ServiceDesc, srv)
}

func _Api_StatV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatisticV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).StatV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_StatV1_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).StatV1(ctx, req.(*StatisticV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_AddCategoryV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCategoryV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).AddCategoryV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_AddCategoryV1_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).AddCategoryV1(ctx, req.(*AddCategoryV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_CategoryListV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryListV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).CategoryListV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_CategoryListV1_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).CategoryListV1(ctx, req.(*CategoryListV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_DeleteCategoryV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCategoryV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).DeleteCategoryV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_DeleteCategoryV1_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).DeleteCategoryV1(ctx, req.(*DeleteCategoryV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_WasteListV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WasteListV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).WasteListV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_WasteListV1_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).WasteListV1(ctx, req.(*WasteListV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_AddWasteListV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddWasteListV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).AddWasteListV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_AddWasteListV1_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).AddWasteListV1(ctx, req.(*AddWasteListV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_DeleteWasteListV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWasteListV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).DeleteWasteListV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Api_DeleteWasteListV1_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).DeleteWasteListV1(ctx, req.(*DeleteWasteListV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Api_ServiceDesc is the grpc.ServiceDesc for Api service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Api_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Api",
	HandlerType: (*ApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StatV1",
			Handler:    _Api_StatV1_Handler,
		},
		{
			MethodName: "AddCategoryV1",
			Handler:    _Api_AddCategoryV1_Handler,
		},
		{
			MethodName: "CategoryListV1",
			Handler:    _Api_CategoryListV1_Handler,
		},
		{
			MethodName: "DeleteCategoryV1",
			Handler:    _Api_DeleteCategoryV1_Handler,
		},
		{
			MethodName: "WasteListV1",
			Handler:    _Api_WasteListV1_Handler,
		},
		{
			MethodName: "AddWasteListV1",
			Handler:    _Api_AddWasteListV1_Handler,
		},
		{
			MethodName: "DeleteWasteListV1",
			Handler:    _Api_DeleteWasteListV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/api.proto",
}
