package loms

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"route256/checkout/internal/app/domain"
	"route256/checkout/internal/pb/route256/loms/v1"
)

type Client struct {
	conn *grpc.ClientConn
	impl loms_v1.LomsServiceClient
}

func New(ctx context.Context, host string) (*Client, error) {
	conn, err := grpc.DialContext(
		ctx,
		host,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("grpc dial: %w", err)
	}

	return &Client{
		conn: conn,
		impl: loms_v1.NewLomsServiceClient(conn),
	}, nil
}

func (c *Client) Close() {
	if c.conn == nil {
		return
	}
	err := c.conn.Close()
	if err != nil {
		log.Println("ERROR close connection:", err)
	}
}

func (c *Client) CreateOrder(ctx context.Context, mdl *domain.OrderInfo) (int64, error) {
	res, err := c.impl.CreateOrder(ctx, toCreateOrderRequest(mdl))
	if err != nil {
		return 0, fmt.Errorf("send request: %w", err)
	}

	return res.OrderId, nil
}

func (c *Client) GetStocks(ctx context.Context, sku uint32) ([]domain.StockInfo, error) {
	res, err := c.impl.GetStocks(ctx, &loms_v1.GetStocksRequest{Sku: sku})
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}

	return fromGetStocksResponse(res), nil
}
