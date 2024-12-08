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

func TestService_DeleteFromCart(t *testing.T) {
	t.Parallel()

	var (
		ctx    = context.Background()
		change = &domain.UpdateCart{
			UserID: 1,
			SKU:    1,
			Count:  2,
		}
		testErr = errors.New("test error")
	)

	t.Run("happy path delete part", func(t *testing.T) {
		t.Parallel()

		repo := mocks.NewRepository(t)
		repo.On("CountOfSKU", mock.Anything, change.UserID, change.SKU).
			Return(uint16(3), nil).
			Once()
		repo.On("DecreaseCount", mock.Anything, change).
			Return(nil).
			Once()

		s := New(nil, nil, mocks.NewProxyTxManager(t), repo)
		err := s.DeleteFromCart(ctx, change)

		assert.NoError(t, err)
	})

	t.Run("happy path delete all", func(t *testing.T) {
		t.Parallel()

		repo := mocks.NewRepository(t)
		repo.On("CountOfSKU", mock.Anything, change.UserID, change.SKU).
			Return(uint16(2), nil).
			Once()
		repo.On("DeleteCartItem", mock.Anything, change.UserID, change.SKU).
			Return(nil).
			Once()

		s := New(nil, nil, mocks.NewProxyTxManager(t), repo)
		err := s.DeleteFromCart(ctx, change)

		assert.NoError(t, err)
	})

	t.Run("cart insufficient", func(t *testing.T) {
		t.Parallel()

		repo := mocks.NewRepository(t)
		repo.On("CountOfSKU", mock.Anything, change.UserID, change.SKU).
			Return(uint16(1), nil).
			Once()

		s := New(nil, nil, mocks.NewProxyTxManager(t), repo)
		err := s.DeleteFromCart(ctx, change)

		assert.ErrorIs(t, err, ErrCartInsufficient)
	})

	t.Run("count of sku failed", func(t *testing.T) {
		t.Parallel()

		repo := mocks.NewRepository(t)
		repo.On("CountOfSKU", mock.Anything, change.UserID, change.SKU).
			Return(uint16(0), testErr).
			Once()

		s := New(nil, nil, mocks.NewProxyTxManager(t), repo)
		err := s.DeleteFromCart(ctx, change)

		assert.ErrorIs(t, err, testErr)
	})

	t.Run("decrease count failed", func(t *testing.T) {
		t.Parallel()

		repo := mocks.NewRepository(t)
		repo.On("CountOfSKU", mock.Anything, change.UserID, change.SKU).
			Return(uint16(3), nil).
			Once()
		repo.On("DecreaseCount", mock.Anything, change).
			Return(testErr).
			Once()

		s := New(nil, nil, mocks.NewProxyTxManager(t), repo)
		err := s.DeleteFromCart(ctx, change)

		assert.ErrorIs(t, err, testErr)
	})

	t.Run("delete cart item failed", func(t *testing.T) {
		t.Parallel()

		repo := mocks.NewRepository(t)
		repo.On("CountOfSKU", mock.Anything, change.UserID, change.SKU).
			Return(uint16(2), nil).
			Once()
		repo.On("DeleteCartItem", mock.Anything, change.UserID, change.SKU).
			Return(testErr).
			Once()

		s := New(nil, nil, mocks.NewProxyTxManager(t), repo)
		err := s.DeleteFromCart(ctx, change)

		assert.ErrorIs(t, err, testErr)
	})
}
