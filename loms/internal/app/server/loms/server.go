package loms

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"route256/loms/internal/app/service/loms"
	. "route256/loms/pkg/loms/v1"
)

type Server struct {
	UnimplementedLomsServiceServer
	impl *loms.Service
}

func New(impl *loms.Service) *Server {
	return &Server{impl: impl}
}

func (s *Server) CreateOrder(ctx context.Context, req *CreateOrderRequest) (*CreateOrderResponse, error) {
	orderID, err := s.impl.CreateOrder(ctx, req.User, fromCreateOrderRequest(req))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &CreateOrderResponse{OrderId: orderID}, nil
}

func (s *Server) ListOrder(ctx context.Context, req *ListOrderRequest) (*ListOrderResponse, error) {
	order, err := s.impl.ListOrder(ctx, req.OrderId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return toListOrderResponse(order), nil
}

func (s *Server) OrderPayed(ctx context.Context, req *OrderPayedRequest) (*OrderPayedResponse, error) {
	err := s.impl.OrderPayed(ctx, req.OrderId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &OrderPayedResponse{}, nil
}

func (s *Server) CancelOrder(ctx context.Context, req *CancelOrderRequest) (*CancelOrderResponse, error) {
	err := s.impl.CancelOrder(ctx, req.OrderId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &CancelOrderResponse{}, nil
}

func (s *Server) GetStocks(ctx context.Context, req *GetStocksRequest) (*GetStocksResponse, error) {
	stocks, err := s.impl.GetStocks(ctx, req.Sku)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return toGetStocksResponse(stocks), nil
}
