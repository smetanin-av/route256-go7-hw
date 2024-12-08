package checkout

import (
	"context"
	"fmt"

	"route256/checkout/internal/app/domain"
)

func (s *Service) Purchase(ctx context.Context, userID int64) (int64, error) {
	items, err := s.repository.GetItemsInCart(ctx, userID)
	if err != nil {
		return 0, fmt.Errorf("get items in cart: %w", err)
	}

	if len(items) == 0 {
		return 0, ErrCartIsEmpty
	}

	orderID, err := s.lomsCli.CreateOrder(ctx, &domain.OrderInfo{
		UserID: userID,
		Items:  items,
	})
	if err != nil {
		return 0, fmt.Errorf("call loms: %w", err)
	}

	err = s.repository.ClearCart(ctx, userID)
	if err != nil {
		return orderID, fmt.Errorf("clear cart: %w", err)
	}

	return orderID, nil
}
