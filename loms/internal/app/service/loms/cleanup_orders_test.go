package loms

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"route256/loms/internal/app/domain"
	"route256/loms/internal/app/service/loms/mocks"
)

func TestService_CleanUpOrders(t *testing.T) {
	t.Parallel()

	const (
		orderID int64 = 1
		delta         = time.Second
	)

	var (
		ctx        = context.Background()
		testErr    = errors.New("test error")
		statusOld  = domain.StatusAwaitingPayment
		statusNew  = domain.StatusCancelled
		isMsgMatch = func(msg *domain.StatusMessage) bool {
			return msg.OrderID == orderID && msg.StatusOld == statusOld && msg.StatusNew == statusNew
		}
	)

	t.Run("happy path", func(t *testing.T) {
		t.Parallel()

		repo := mocks.NewRepository(t)
		repo.On("GetOrdersToCleanUp", mock.Anything, delta).
			Return([]int64{orderID}, nil).
			Once()
		repo.On("GetStatus", mock.Anything, orderID).
			Return(statusOld, nil).
			Once()
		repo.On("CancelReserve", mock.Anything, orderID).
			Return(nil).
			Once()
		repo.On("SetStatus", mock.Anything, orderID, statusNew).
			Return(nil).
			Once()

		sender := mocks.NewMsgSender(t)
		sender.On("SendMessage", mock.AnythingOfType("uuid.UUID"), mock.MatchedBy(isMsgMatch)).
			Return(nil).
			Once()

		s := New(mocks.NewProxyTxManager(t), repo, sender)
		s.CleanUpOrders(ctx, delta)
	})

	t.Run("get orders failed", func(t *testing.T) {
		t.Parallel()

		repo := mocks.NewRepository(t)
		repo.On("GetOrdersToCleanUp", mock.Anything, delta).
			Return(nil, testErr).
			Once()

		s := New(nil, repo, nil)
		s.CleanUpOrders(ctx, delta)
	})

	t.Run("cancel order failed", func(t *testing.T) {
		t.Parallel()

		repo := mocks.NewRepository(t)
		repo.On("GetOrdersToCleanUp", mock.Anything, delta).
			Return([]int64{orderID}, nil).
			Once()
		repo.On("GetStatus", mock.Anything, orderID).
			Return(domain.StatusInvalid, testErr).
			Once()

		s := New(mocks.NewProxyTxManager(t), repo, nil)
		s.CleanUpOrders(ctx, delta)
	})
}
