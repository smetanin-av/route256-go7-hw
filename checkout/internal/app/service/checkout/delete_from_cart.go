package checkout

import (
	"context"
	"fmt"

	"route256/checkout/internal/app/domain"
)

func (s *Service) DeleteFromCart(ctx context.Context, change *domain.UpdateCart) error {
	err := s.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		count, err := s.repository.CountOfSKU(ctx, change.UserID, change.SKU)
		if err != nil {
			return fmt.Errorf("count by sku: %w", err)
		}

		if count < change.Count {
			return ErrCartInsufficient
		}

		if count > change.Count {
			err = s.repository.DecreaseCount(ctx, change)
			if err != nil {
				return fmt.Errorf("decrease count: %w", err)
			}
		} else {
			err = s.repository.DeleteCartItem(ctx, change.UserID, change.SKU)
			if err != nil {
				return fmt.Errorf("delete cart item: %w", err)
			}
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("run repeatable read: %w", err)
	}

	return nil
}
