package product

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redis_rate/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"route256/checkout/internal/app/domain"
	"route256/checkout/internal/pb/route256/product"
)

const limiterKey = "product_service_rate_limit"

type Client struct {
	conn    *grpc.ClientConn
	impl    product.ProductServiceClient
	token   string
	limiter *redis_rate.Limiter
	limit   redis_rate.Limit
}

func New(ctx context.Context, host string, token string, rcli *redis.Client) (*Client, error) {
	conn, err := grpc.DialContext(
		ctx,
		host,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("grpc dial: %w", err)
	}

	return &Client{
		conn:    conn,
		impl:    product.NewProductServiceClient(conn),
		token:   token,
		limiter: redis_rate.NewLimiter(rcli),
		limit:   redis_rate.PerSecond(10),
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

func (c *Client) GetProduct(ctx context.Context, sku uint32) (*domain.ProductInfo, error) {
	err := c.waitUntilRequestAllowed(ctx)
	if err != nil {
		return nil, fmt.Errorf("check limiter: %w", err)
	}

	req := product.GetProductRequest{
		Token: c.token,
		Sku:   sku,
	}

	res, err := c.impl.GetProduct(ctx, &req)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}

	return &domain.ProductInfo{
		Name:  res.Name,
		Price: res.Price,
	}, nil
}

func (c *Client) ListSkus(ctx context.Context, afterSku uint32, count uint32) ([]uint32, error) {
	req := product.ListSkusRequest{
		Token:         c.token,
		StartAfterSku: afterSku,
		Count:         count,
	}

	res, err := c.impl.ListSkus(ctx, &req)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}

	return res.Skus, nil
}

func (c *Client) waitUntilRequestAllowed(ctx context.Context) error {
	for {
		res, err := c.limiter.Allow(ctx, limiterKey, c.limit)
		if err != nil {
			return fmt.Errorf("limiter allow: %w", err)
		}

		if res.Allowed > 0 {
			return nil
		}
		time.Sleep(res.RetryAfter)
	}
}
