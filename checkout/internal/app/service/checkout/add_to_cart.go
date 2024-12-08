package checkout

import (
	"context"
	"fmt"

	"route256/checkout/internal/app/domain"
)

func isStockInsufficient(stocks []domain.StockInfo, requestedCount uint64) bool {
	for _, stock := range stocks {
		if stock.Count >= requestedCount {
			return false
		}
		requestedCount -= stock.Count
	}

	return true
}

func (s *Service) AddToCart(ctx context.Context, change *domain.UpdateCart) error {
	stocks, err := s.lomsCli.GetStocks(ctx, change.SKU)
	if err != nil {
		return fmt.Errorf("get stocks: %w", err)
	}

	if isStockInsufficient(stocks, uint64(change.Count)) {
		return ErrStockInsufficient
	}

	return s.repository.AddCartItem(ctx, change)
}
