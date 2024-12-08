package loms

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"route256/loms/internal/app/domain"
	"route256/loms/internal/app/service/loms/mocks"
)

func TestService_ListOrder(t *testing.T) {
	t.Parallel()

	const (
		orderID int64 = 1
	)

	var (
		ctx       = context.Background()
		orderInfo = &domain.OrderInfo{
			UserID: 1,
			Status: domain.StatusNew,
		}
		testErr = errors.New("test error")
	)

	t.Run("happy path", func(t *testing.T) {
		t.Parallel()

		repo := mocks.NewRepository(t)
		repo.On("ListOrder", mock.Anything, orderID).
			Return(orderInfo, nil).
			Once()

		s := New(nil, repo, nil)
		res, err := s.ListOrder(ctx, orderID)

		assert.NoError(t, err)
		assert.Equal(t, res, orderInfo)
	})

	t.Run("list order failed", func(t *testing.T) {
		t.Parallel()

		repo := mocks.NewRepository(t)
		repo.On("ListOrder", mock.Anything, orderID).
			Return(nil, testErr).
			Once()

		s := New(nil, repo, nil)
		res, err := s.ListOrder(ctx, orderID)

		assert.ErrorIs(t, err, testErr)
		assert.Nil(t, res)
	})
}
