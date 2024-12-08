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

func TestService_ListCart(t *testing.T) {
	t.Parallel()

	const (
		userID     int64  = 1
		totalPrice uint32 = 5
	)

	var (
		ctx     = context.Background()
		item1   = &domain.OrderItem{SKU: 1, Count: 1}
		item2   = &domain.OrderItem{SKU: 2, Count: 2}
		items   = []*domain.OrderItem{item1, item2}
		info1   = &domain.ProductInfo{Name: "sku #1", Price: 1}
		info2   = &domain.ProductInfo{Name: "sku #2", Price: 2}
		testErr = errors.New("test error")
	)

	t.Run("happy path", func(t *testing.T) {
		t.Parallel()

		repo := mocks.NewRepository(t)
		repo.On("GetItemsInCart", mock.Anything, userID).
			Return(items, nil).
			Once()

		productCli := mocks.NewProductClient(t)
		productCli.On("GetProduct", mock.Anything, item1.SKU).
			Return(info1, nil).
			Once()
		productCli.On("GetProduct", mock.Anything, item2.SKU).
			Return(info2, nil).
			Once()

		s := New(nil, productCli, nil, repo)
		res, err := s.ListCart(ctx, userID)

		assert.NoError(t, err)
		assert.ElementsMatch(t, res.Items, []*domain.ListCartItem{
			{
				SKU:   item1.SKU,
				Count: item1.Count,
				Name:  info1.Name,
				Price: info1.Price,
			},
			{
				SKU:   item2.SKU,
				Count: item2.Count,
				Name:  info2.Name,
				Price: info2.Price,
			},
		})
		assert.Equal(t, totalPrice, res.TotalPrice)
	})

	t.Run("get items in cart failed", func(t *testing.T) {
		t.Parallel()

		repo := mocks.NewRepository(t)
		repo.On("GetItemsInCart", mock.Anything, userID).
			Return(nil, testErr).
			Once()

		s := New(nil, nil, nil, repo)
		res, err := s.ListCart(ctx, userID)

		assert.ErrorIs(t, err, testErr)
		assert.Nil(t, res)
	})

	t.Run("get product failed", func(t *testing.T) {
		t.Parallel()

		repo := mocks.NewRepository(t)
		repo.On("GetItemsInCart", mock.Anything, userID).
			Return(items, nil).
			Once()

		productCli := mocks.NewProductClient(t)
		productCli.On("GetProduct", mock.Anything, item1.SKU).
			Return(info1, nil).
			Once()
		productCli.On("GetProduct", mock.Anything, item2.SKU).
			Return(nil, testErr).
			Once()

		s := New(nil, productCli, nil, repo)
		res, err := s.ListCart(ctx, userID)

		assert.ErrorIs(t, err, testErr)
		assert.Nil(t, res)
	})

}

func TestService_enrichCartItem(t *testing.T) {
	t.Parallel()

	var (
		ctx  = context.Background()
		item = &domain.OrderItem{
			SKU:   1,
			Count: 1,
		}
		info = &domain.ProductInfo{
			Name:  "sku #1",
			Price: 1,
		}
		testErr = errors.New("test error")
	)

	t.Run("happy path", func(t *testing.T) {
		t.Parallel()

		productCli := mocks.NewProductClient(t)
		productCli.On("GetProduct", mock.Anything, item.SKU).
			Return(info, nil).
			Once()

		s := New(nil, productCli, nil, nil)
		res, err := s.enrichCartItem(ctx, item)

		assert.NoError(t, err)
		assert.Equal(t, res, &domain.ListCartItem{
			SKU: item.SKU, Count: item.Count, Name: info.Name, Price: info.Price,
		})
	})

	t.Run("get product failed", func(t *testing.T) {
		t.Parallel()

		productCli := mocks.NewProductClient(t)
		productCli.On("GetProduct", mock.Anything, item.SKU).
			Return(nil, testErr).
			Once()

		s := New(nil, productCli, nil, nil)
		res, err := s.enrichCartItem(ctx, item)

		assert.ErrorIs(t, err, testErr)
		assert.Nil(t, res)
	})
}
