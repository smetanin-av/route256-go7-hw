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

func TestService_GetStocks(t *testing.T) {
	t.Parallel()

	const (
		sku uint32 = 1
	)

	var (
		ctx    = context.Background()
		stocks = []*domain.StockInfo{
			{
				WarehouseID: 1,
				Count:       1,
			},
		}
		testErr = errors.New("test error")
	)

	t.Run("happy path", func(t *testing.T) {
		t.Parallel()

		repo := mocks.NewRepository(t)
		repo.On("GetStocks", mock.Anything, sku).
			Return(stocks, nil).
			Once()

		s := New(nil, repo, nil)
		res, err := s.GetStocks(ctx, sku)

		assert.NoError(t, err)
		assert.Equal(t, res, stocks)
	})

	t.Run("get stocks failed", func(t *testing.T) {
		t.Parallel()

		repo := mocks.NewRepository(t)
		repo.On("GetStocks", mock.Anything, sku).
			Return(nil, testErr).
			Once()

		s := New(nil, repo, nil)
		res, err := s.GetStocks(ctx, sku)

		assert.ErrorIs(t, err, testErr)
		assert.Nil(t, res)
	})
}
