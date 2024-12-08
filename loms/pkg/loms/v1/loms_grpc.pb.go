// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: loms/v1/loms.proto

package loms_v1

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

// LomsServiceClient is the client API for LomsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LomsServiceClient interface {
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error)
	ListOrder(ctx context.Context, in *ListOrderRequest, opts ...grpc.CallOption) (*ListOrderResponse, error)
	OrderPayed(ctx context.Context, in *OrderPayedRequest, opts ...grpc.CallOption) (*OrderPayedResponse, error)
	CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*CancelOrderResponse, error)
	GetStocks(ctx context.Context, in *GetStocksRequest, opts ...grpc.CallOption) (*GetStocksResponse, error)
}

type lomsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLomsServiceClient(cc grpc.ClientConnInterface) LomsServiceClient {
	return &lomsServiceClient{cc}
}

func (c *lomsServiceClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error) {
	out := new(CreateOrderResponse)
	err := c.cc.Invoke(ctx, "/route256.api.v1.loms.LomsService/CreateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lomsServiceClient) ListOrder(ctx context.Context, in *ListOrderRequest, opts ...grpc.CallOption) (*ListOrderResponse, error) {
	out := new(ListOrderResponse)
	err := c.cc.Invoke(ctx, "/route256.api.v1.loms.LomsService/ListOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lomsServiceClient) OrderPayed(ctx context.Context, in *OrderPayedRequest, opts ...grpc.CallOption) (*OrderPayedResponse, error) {
	out := new(OrderPayedResponse)
	err := c.cc.Invoke(ctx, "/route256.api.v1.loms.LomsService/OrderPayed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lomsServiceClient) CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*CancelOrderResponse, error) {
	out := new(CancelOrderResponse)
	err := c.cc.Invoke(ctx, "/route256.api.v1.loms.LomsService/CancelOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lomsServiceClient) GetStocks(ctx context.Context, in *GetStocksRequest, opts ...grpc.CallOption) (*GetStocksResponse, error) {
	out := new(GetStocksResponse)
	err := c.cc.Invoke(ctx, "/route256.api.v1.loms.LomsService/GetStocks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LomsServiceServer is the server API for LomsService service.
// All implementations must embed UnimplementedLomsServiceServer
// for forward compatibility
type LomsServiceServer interface {
	CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error)
	ListOrder(context.Context, *ListOrderRequest) (*ListOrderResponse, error)
	OrderPayed(context.Context, *OrderPayedRequest) (*OrderPayedResponse, error)
	CancelOrder(context.Context, *CancelOrderRequest) (*CancelOrderResponse, error)
	GetStocks(context.Context, *GetStocksRequest) (*GetStocksResponse, error)
	mustEmbedUnimplementedLomsServiceServer()
}

// UnimplementedLomsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLomsServiceServer struct {
}

func (UnimplementedLomsServiceServer) CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedLomsServiceServer) ListOrder(context.Context, *ListOrderRequest) (*ListOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrder not implemented")
}
func (UnimplementedLomsServiceServer) OrderPayed(context.Context, *OrderPayedRequest) (*OrderPayedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderPayed not implemented")
}
func (UnimplementedLomsServiceServer) CancelOrder(context.Context, *CancelOrderRequest) (*CancelOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOrder not implemented")
}
func (UnimplementedLomsServiceServer) GetStocks(context.Context, *GetStocksRequest) (*GetStocksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStocks not implemented")
}
func (UnimplementedLomsServiceServer) mustEmbedUnimplementedLomsServiceServer() {}

// UnsafeLomsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LomsServiceServer will
// result in compilation errors.
type UnsafeLomsServiceServer interface {
	mustEmbedUnimplementedLomsServiceServer()
}

func RegisterLomsServiceServer(s grpc.ServiceRegistrar, srv LomsServiceServer) {
	s.RegisterService(&LomsService_ServiceDesc, srv)
}

func _LomsService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LomsServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/route256.api.v1.loms.LomsService/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LomsServiceServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LomsService_ListOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LomsServiceServer).ListOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/route256.api.v1.loms.LomsService/ListOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LomsServiceServer).ListOrder(ctx, req.(*ListOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LomsService_OrderPayed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderPayedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LomsServiceServer).OrderPayed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/route256.api.v1.loms.LomsService/OrderPayed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LomsServiceServer).OrderPayed(ctx, req.(*OrderPayedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LomsService_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LomsServiceServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/route256.api.v1.loms.LomsService/CancelOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LomsServiceServer).CancelOrder(ctx, req.(*CancelOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LomsService_GetStocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStocksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LomsServiceServer).GetStocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/route256.api.v1.loms.LomsService/GetStocks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LomsServiceServer).GetStocks(ctx, req.(*GetStocksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LomsService_ServiceDesc is the grpc.ServiceDesc for LomsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LomsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "route256.api.v1.loms.LomsService",
	HandlerType: (*LomsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _LomsService_CreateOrder_Handler,
		},
		{
			MethodName: "ListOrder",
			Handler:    _LomsService_ListOrder_Handler,
		},
		{
			MethodName: "OrderPayed",
			Handler:    _LomsService_OrderPayed_Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _LomsService_CancelOrder_Handler,
		},
		{
			MethodName: "GetStocks",
			Handler:    _LomsService_GetStocks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "loms/v1/loms.proto",
}
