package loms

import (
	"context"

	"route256/loms/internal/app/domain"
)

func (s *Service) ListOrder(ctx context.Context, orderID int64) (*domain.OrderInfo, error) {
	return s.repository.ListOrder(ctx, orderID)
}
