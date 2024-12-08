package loms

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"route256/loms/internal/app/domain"
	"route256/loms/internal/app/service/loms/mocks"
)

func TestService_CancelOrder(t *testing.T) {
	t.Parallel()

	const (
		orderID int64 = 1
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
		err := s.CancelOrder(ctx, orderID)

		assert.NoError(t, err)
	})

	t.Run("get status failed", func(t *testing.T) {
		t.Parallel()

		repo := mocks.NewRepository(t)
		repo.On("GetStatus", mock.Anything, orderID).
			Return(domain.StatusInvalid, testErr).
			Once()

		s := New(mocks.NewProxyTxManager(t), repo, nil)
		err := s.CancelOrder(ctx, orderID)

		assert.ErrorIs(t, err, testErr)
	})

	for _, status := range []domain.Status{
		domain.StatusInvalid,
		domain.StatusNew,
		domain.StatusFailed,
		domain.StatusPayed,
		domain.StatusCancelled,
	} {
		status := status

		t.Run(fmt.Sprintf("wrong status %v", status), func(t *testing.T) {
			t.Parallel()

			repo := mocks.NewRepository(t)
			repo.On("GetStatus", mock.Anything, orderID).
				Return(status, nil).
				Once()

			s := New(mocks.NewProxyTxManager(t), repo, nil)
			err := s.CancelOrder(ctx, orderID)

			assert.ErrorIs(t, err, ErrWrongOrderStatus)
		})
	}

	t.Run("cancel reserve failed", func(t *testing.T) {
		t.Parallel()

		repo := mocks.NewRepository(t)
		repo.On("GetStatus", mock.Anything, orderID).
			Return(domain.StatusAwaitingPayment, nil).
			Once()
		repo.On("CancelReserve", mock.Anything, orderID).
			Return(testErr).
			Once()

		s := New(mocks.NewProxyTxManager(t), repo, nil)
		err := s.CancelOrder(ctx, orderID)

		assert.ErrorIs(t, err, testErr)
	})

	t.Run("set status failed", func(t *testing.T) {
		t.Parallel()

		repo := mocks.NewRepository(t)
		repo.On("GetStatus", mock.Anything, orderID).
			Return(domain.StatusAwaitingPayment, nil).
			Once()
		repo.On("CancelReserve", mock.Anything, orderID).
			Return(nil).
			Once()
		repo.On("SetStatus", mock.Anything, orderID, domain.StatusCancelled).
			Return(testErr).
			Once()

		s := New(mocks.NewProxyTxManager(t), repo, nil)
		err := s.CancelOrder(ctx, orderID)

		assert.ErrorIs(t, err, testErr)
	})
}
