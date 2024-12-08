package checkout

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"route256/checkout/internal/app/service/checkout"
	. "route256/checkout/pkg/checkout/v1"
)

type Server struct {
	UnimplementedCheckoutServiceServer
	impl *checkout.Service
}

func New(impl *checkout.Service) *Server {
	return &Server{impl: impl}
}

func (s *Server) AddToCart(ctx context.Context, req *AddToCartRequest) (*AddToCartResponse, error) {
	err := s.impl.AddToCart(ctx, fromAddToCartRequest(req))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &AddToCartResponse{}, nil
}

func (s *Server) DeleteFromCart(ctx context.Context, req *DeleteFromCartRequest) (*DeleteFromCartResponse, error) {
	err := s.impl.DeleteFromCart(ctx, fromDeleteFromCartRequest(req))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &DeleteFromCartResponse{}, nil
}

func (s *Server) ListCart(ctx context.Context, req *ListCartRequest) (*ListCartResponse, error) {
	cartInfo, err := s.impl.ListCart(ctx, req.User)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return toListCartResponse(cartInfo), nil
}

func (s *Server) Purchase(ctx context.Context, req *PurchaseRequest) (*PurchaseResponse, error) {
	orderID, err := s.impl.Purchase(ctx, req.User)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &PurchaseResponse{OrderId: orderID}, nil
}
