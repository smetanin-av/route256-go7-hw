package checkout

import (
	"context"
	"fmt"

	"route256/checkout/internal/app/domain"
	"route256/libs/worker_pool"
)

const (
	poolSize = 5
)

func (s *Service) ListCart(ctx context.Context, userID int64) (*domain.ListCart, error) {
	itemsInCart, err := s.repository.GetItemsInCart(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get items in cart: %w", err)
	}

	poolCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	enrichedItems := worker_pool.Process(poolCtx, poolSize, itemsInCart, s.enrichCartItem)

	var cart domain.ListCart
	for dto := range enrichedItems {
		if dto.Err != nil {
			cancel()
			return nil, dto.Err
		}

		cart.Items = append(cart.Items, dto.Out)
		cart.TotalPrice += uint32(dto.Out.Count) * dto.Out.Price
	}

	return &cart, nil
}

func (s *Service) enrichCartItem(ctx context.Context, item *domain.OrderItem) (*domain.ListCartItem, error) {
	info, err := s.productCli.GetProduct(ctx, item.SKU)
	if err != nil {
		return nil, fmt.Errorf("get product: %w", err)
	}

	return &domain.ListCartItem{
		SKU:   item.SKU,
		Count: item.Count,
		Name:  info.Name,
		Price: info.Price,
	}, nil
}
