package loms

//go:generate mockery --case underscore --name TxManager
//go:generate mockery --case underscore --name Repository
//go:generate mockery --case underscore --name MsgSender

import (
	"context"
	"fmt"
	"time"

	"route256/loms/internal/app/domain"
)

type TxManager interface {
	RunRepeatableRead(ctx context.Context, fn func(ctxTx context.Context) error) error
}

type Repository interface {
	AddOrder(ctx context.Context, userID int64, items []domain.OrderItem) (int64, error)
	ListOrder(ctx context.Context, orderID int64) (*domain.OrderInfo, error)
	GetStatus(ctx context.Context, orderID int64) (domain.Status, error)
	SetStatus(ctx context.Context, orderID int64, status domain.Status) error
	GetStocks(ctx context.Context, sku uint32) ([]*domain.StockInfo, error)
	ReserveStocks(ctx context.Context, info *domain.ReserveInfo) error
	CancelReserve(ctx context.Context, orderID int64) error
	MarkAsSold(ctx context.Context, orderID int64) error
	GetOrdersToCleanUp(ctx context.Context, delta time.Duration) ([]int64, error)
}

type MsgSender interface {
	SendMessage(key fmt.Stringer, msg any) error
}

type Service struct {
	txManager  TxManager
	repository Repository
	sender     MsgSender
}

func New(txMan TxManager, repo Repository, sender MsgSender) *Service {
	return &Service{
		txManager:  txMan,
		repository: repo,
		sender:     sender,
	}
}
