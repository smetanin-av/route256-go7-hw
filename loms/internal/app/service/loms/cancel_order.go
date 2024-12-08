package loms

import (
	"context"
	"fmt"

	"route256/loms/internal/app/domain"
)

func (s *Service) CancelOrder(ctx context.Context, orderID int64) error {
	const statusNew = domain.StatusCancelled

	err := s.txManager.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		statusOld, err := s.repository.GetStatus(ctx, orderID)
		if err != nil {
			return fmt.Errorf("get order status: %w", err)
		}

		if statusOld != domain.StatusAwaitingPayment {
			return ErrWrongOrderStatus
		}

		err = s.repository.CancelReserve(ctxTx, orderID)
		if err != nil {
			return fmt.Errorf("cancel reserve: %w", err)
		}

		err = s.repository.SetStatus(ctxTx, orderID, statusNew)
		if err != nil {
			return fmt.Errorf("set order status: %w", err)
		}

		err = s.SendStatusMessage(orderID, statusOld, statusNew)
		if err != nil {
			return fmt.Errorf("send message: %w", err)
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("run repeatable read: %w", err)
	}

	return nil
}
