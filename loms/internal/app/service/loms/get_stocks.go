package loms

import (
	"context"

	"route256/loms/internal/app/domain"
)

func (s *Service) GetStocks(ctx context.Context, sku uint32) ([]*domain.StockInfo, error) {
	return s.repository.GetStocks(ctx, sku)
}
