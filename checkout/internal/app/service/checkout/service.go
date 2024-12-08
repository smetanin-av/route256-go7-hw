package checkout

//go:generate mockery --case underscore --name LomsClient
//go:generate mockery --case underscore --name ProductClient
//go:generate mockery --case underscore --name TxManager
//go:generate mockery --case underscore --name Repository

import (
	"context"

	"route256/checkout/internal/app/domain"
)

type LomsClient interface {
	GetStocks(ctx context.Context, sku uint32) ([]domain.StockInfo, error)
	CreateOrder(ctx context.Context, mdl *domain.OrderInfo) (int64, error)
}

type ProductClient interface {
	GetProduct(ctx context.Context, sku uint32) (*domain.ProductInfo, error)
	ListSkus(ctx context.Context, afterSku uint32, count uint32) ([]uint32, error)
}

type TxManager interface {
	RunRepeatableRead(ctx context.Context, fn func(ctxTx context.Context) error) error
}

type Repository interface {
	AddCartItem(ctx context.Context, change *domain.UpdateCart) error
	CountOfSKU(ctx context.Context, userID int64, sku uint32) (uint16, error)
	DecreaseCount(ctx context.Context, change *domain.UpdateCart) error
	DeleteCartItem(ctx context.Context, userID int64, sku uint32) error
	GetItemsInCart(ctx context.Context, userID int64) ([]*domain.OrderItem, error)
	ClearCart(ctx context.Context, userID int64) error
}

type Service struct {
	lomsCli    LomsClient
	productCli ProductClient
	txManager  TxManager
	repository Repository
}

func New(lomsCli LomsClient, productCli ProductClient, txMan TxManager, repo Repository) *Service {
	return &Service{
		lomsCli:    lomsCli,
		productCli: productCli,
		txManager:  txMan,
		repository: repo,
	}
}
