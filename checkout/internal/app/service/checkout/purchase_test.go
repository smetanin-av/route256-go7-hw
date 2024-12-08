package checkout

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"route256/checkout/internal/app/domain"
	"route256/checkout/internal/app/service/checkout/mocks"
)

func TestService_Purchase(t *testing.T) {
	t.Parallel()

	const (
		userID  int64 = 1
		orderID int64 = 1
	)

	var (
		ctx   = context.Background()
		item1 = &domain.OrderItem{
			SKU:   1,
			Count: 1,
		}
		item2 = &domain.OrderItem{
			SKU:   2,
			Count: 2,
		}
		items = []*domain.OrderItem{
			item1,
			item2,
		}
		orderInfo = &domain.OrderInfo{
			UserID: userID,
			Items:  items,
		}
		testErr = errors.New("test error")
	)

	t.Run("happy path", func(t *testing.T) {
		t.Parallel()

		repo := mocks.NewRepository(t)
		repo.On("GetItemsInCart", mock.Anything, userID).
			Return(items, nil).
			Once()
		repo.On("ClearCart", mock.Anything, userID).
			Return(nil).
			Once()

		lomsCli := mocks.NewLomsClient(t)
		lomsCli.On("CreateOrder", mock.Anything, orderInfo).
			Return(orderID, nil).
			Once()

		s := New(lomsCli, nil, nil, repo)
		res, err := s.Purchase(ctx, userID)

		assert.NoError(t, err)
		assert.Equal(t, res, orderID)
	})

	t.Run("get items in cart failed", func(t *testing.T) {
		t.Parallel()

		repo := mocks.NewRepository(t)
		repo.On("GetItemsInCart", mock.Anything, userID).
			Return(nil, testErr).
			Once()

		s := New(nil, nil, nil, repo)
		res, err := s.Purchase(ctx, userID)

		assert.ErrorIs(t, err, testErr)
		assert.Empty(t, res)
	})

	t.Run("empty cart", func(t *testing.T) {
		t.Parallel()

		repo := mocks.NewRepository(t)
		repo.On("GetItemsInCart", mock.Anything, userID).
			Return([]*domain.OrderItem{}, nil).
			Once()

		s := New(nil, nil, nil, repo)
		res, err := s.Purchase(ctx, userID)

		assert.ErrorIs(t, err, ErrCartIsEmpty)
		assert.Empty(t, res)
	})

	t.Run("create order failed", func(t *testing.T) {
		t.Parallel()

		repo := mocks.NewRepository(t)
		repo.On("GetItemsInCart", mock.Anything, userID).
			Return(items, nil).
			Once()

		lomsCli := mocks.NewLomsClient(t)
		lomsCli.On("CreateOrder", mock.Anything, orderInfo).
			Return(int64(0), testErr).
			Once()

		s := New(lomsCli, nil, nil, repo)
		res, err := s.Purchase(ctx, userID)

		assert.ErrorIs(t, err, testErr)
		assert.Empty(t, res)
	})

	t.Run("clear cart failed", func(t *testing.T) {
		t.Parallel()

		repo := mocks.NewRepository(t)
		repo.On("GetItemsInCart", mock.Anything, userID).
			Return(items, nil).
			Once()
		repo.On("ClearCart", mock.Anything, userID).
			Return(testErr).
			Once()

		lomsCli := mocks.NewLomsClient(t)
		lomsCli.On("CreateOrder", mock.Anything, orderInfo).
			Return(orderID, nil).
			Once()

		s := New(lomsCli, nil, nil, repo)
		res, err := s.Purchase(ctx, userID)

		assert.ErrorIs(t, err, testErr)
		assert.Equal(t, res, orderID)
	})
}
