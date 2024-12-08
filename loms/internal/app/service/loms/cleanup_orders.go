package loms

import (
	"context"
	"time"

	"route256/libs/logging"
)

func (s *Service) CleanUpOrders(ctx context.Context, delta time.Duration) {
	ordersIDs, err := s.repository.GetOrdersToCleanUp(ctx, delta)
	if err != nil {
		logging.Errorf("get orders to cleanup: %v", err)
	}

	for _, orderID := range ordersIDs {
		err = s.CancelOrder(ctx, orderID)
		if err != nil {
			logging.Errorf("cancel order: %v", err)
		} else {
			logging.Warnf("order %d canceled", orderID)
		}
	}
}
