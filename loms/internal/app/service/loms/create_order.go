package loms

import (
	"context"
	"fmt"

	"route256/loms/internal/app/domain"
)

func (s *Service) CreateOrder(ctx context.Context, userID int64, items []domain.OrderItem) (int64, error) {
	var orderID int64

	err := s.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		tempID, err := s.repository.AddOrder(ctxTx, userID, items)
		if err != nil {
			return fmt.Errorf("add order: %w", err)
		}

		err = s.reserveItems(ctxTx, tempID, items)
		if err != nil {
			return fmt.Errorf("reserve items: %w", err)
		}

		err = s.SendStatusMessage(tempID, domain.StatusNew, domain.StatusAwaitingPayment)
		if err != nil {
			return fmt.Errorf("send message: %w", err)
		}

		orderID = tempID
		return nil
	})

	if err != nil {
		return 0, fmt.Errorf("run repeatable read: %w", err)
	}

	return orderID, nil
}

func (s *Service) reserveItems(ctx context.Context, orderID int64, items []domain.OrderItem) error {
	for _, item := range items {
		stocks, err := s.repository.GetStocks(ctx, item.SKU)
		if err != nil {
			return fmt.Errorf("get stocks: %w", err)
		}

		reserve := &domain.ReserveInfo{
			OrderID: orderID,
			SKU:     item.SKU,
		}
		rest := uint64(item.Count)
		for i := 0; i < len(stocks) && rest > 0; i++ {
			stock := stocks[i]
			var count uint64
			if stock.Count > rest {
				count = rest
			} else {
				count = stock.Count
			}

			reserve.Items = append(reserve.Items, &domain.StockInfo{
				WarehouseID: stock.WarehouseID,
				Count:       count,
			})
			rest -= count
		}

		if rest > 0 {
			errorMsg := fmt.Sprintf("sku: %d max: %d", item.SKU, uint64(item.Count)-rest)
			return fmt.Errorf("%w: %s", ErrStockInsufficient, errorMsg)
		}

		err = s.repository.ReserveStocks(ctx, reserve)
		if err != nil {
			return fmt.Errorf("reserve stocks: %w", err)
		}
	}

	return nil
}
