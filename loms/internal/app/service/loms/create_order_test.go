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

func TestService_CreateOrder(t *testing.T) {
	t.Parallel()

	const (
		userID  int64 = 1
		orderID int64 = 1
	)

	var (
		ctx  = context.Background()
		item = domain.OrderItem{
			SKU:   1,
			Count: 1,
		}
		items = []domain.OrderItem{
			item,
		}
		stock = &domain.StockInfo{
			WarehouseID: 1,
			Count:       2,
		}
		stocks = []*domain.StockInfo{
			stock,
		}
		reserve = &domain.ReserveInfo{
			OrderID: orderID,
			SKU:     item.SKU,
			Items: []*domain.StockInfo{
				{
					WarehouseID: stock.WarehouseID,
					Count:       uint64(item.Count),
				},
			},
		}
		testErr    = errors.New("test error")
		isMsgMatch = func(msg *domain.StatusMessage) bool {
			return msg.OrderID == orderID && msg.StatusOld == domain.StatusNew && msg.StatusNew == domain.StatusAwaitingPayment
		}
	)

	t.Run("happy path", func(t *testing.T) {
		t.Parallel()

		repo := mocks.NewRepository(t)
		repo.On("AddOrder", mock.Anything, userID, items).
			Return(orderID, nil).
			Once()
		repo.On("GetStocks", mock.Anything, item.SKU).
			Return(stocks, nil).
			Once()
		repo.On("ReserveStocks", mock.Anything, reserve).
			Return(nil).
			Once()

		sender := mocks.NewMsgSender(t)
		sender.On("SendMessage", mock.AnythingOfType("uuid.UUID"), mock.MatchedBy(isMsgMatch)).
			Return(nil).
			Once()

		s := New(mocks.NewProxyTxManager(t), repo, sender)
		res, err := s.CreateOrder(ctx, userID, items)

		assert.NoError(t, err)
		assert.Equal(t, res, orderID)
	})

	t.Run("add order failed", func(t *testing.T) {
		t.Parallel()

		repo := mocks.NewRepository(t)
		repo.On("AddOrder", mock.Anything, userID, items).
			Return(int64(0), testErr).
			Once()

		s := New(mocks.NewProxyTxManager(t), repo, nil)
		res, err := s.CreateOrder(ctx, userID, items)

		assert.ErrorIs(t, err, testErr)
		assert.Empty(t, res)
	})

	t.Run("get stocks failed", func(t *testing.T) {
		t.Parallel()

		repo := mocks.NewRepository(t)
		repo.On("AddOrder", mock.Anything, userID, items).
			Return(orderID, nil).
			Once()
		repo.On("GetStocks", mock.Anything, item.SKU).
			Return(nil, testErr).
			Once()

		s := New(mocks.NewProxyTxManager(t), repo, nil)
		res, err := s.CreateOrder(ctx, userID, items)

		assert.ErrorIs(t, err, testErr)
		assert.Empty(t, res)
	})

	t.Run("reserve stocks failed", func(t *testing.T) {
		t.Parallel()

		repo := mocks.NewRepository(t)
		repo.On("AddOrder", mock.Anything, userID, items).
			Return(orderID, nil).
			Once()
		repo.On("GetStocks", mock.Anything, item.SKU).
			Return(stocks, nil).
			Once()
		repo.On("ReserveStocks", mock.Anything, reserve).
			Return(testErr).
			Once()

		s := New(mocks.NewProxyTxManager(t), repo, nil)
		res, err := s.CreateOrder(ctx, userID, items)

		assert.ErrorIs(t, err, testErr)
		assert.Empty(t, res)
	})
}

func TestService_reserveItems(t *testing.T) {
	t.Parallel()

	const (
		orderID     int64 = 1
		warehouseID int64 = 1
	)

	var (
		ctx  = context.Background()
		item = domain.OrderItem{
			SKU:   1,
			Count: 2,
		}
		items = []domain.OrderItem{
			item,
		}
		reserve = &domain.ReserveInfo{
			OrderID: orderID,
			SKU:     item.SKU,
			Items: []*domain.StockInfo{
				{
					WarehouseID: warehouseID,
					Count:       uint64(item.Count),
				},
			},
		}
	)

	t.Run("stocks gt reserve", func(t *testing.T) {
		t.Parallel()

		stock := &domain.StockInfo{
			WarehouseID: warehouseID,
			Count:       3,
		}

		repo := mocks.NewRepository(t)
		repo.On("GetStocks", mock.Anything, item.SKU).
			Return([]*domain.StockInfo{stock}, nil).
			Once()
		repo.On("ReserveStocks", mock.Anything, reserve).
			Return(nil).
			Maybe()

		s := New(nil, repo, nil)
		res := s.reserveItems(ctx, orderID, items)

		assert.NoError(t, res)
	})

	t.Run("stocks eq reserve", func(t *testing.T) {
		t.Parallel()

		stock := &domain.StockInfo{
			WarehouseID: warehouseID,
			Count:       2,
		}

		repo := mocks.NewRepository(t)
		repo.On("GetStocks", mock.Anything, item.SKU).
			Return([]*domain.StockInfo{stock}, nil).
			Once()
		repo.On("ReserveStocks", mock.Anything, reserve).
			Return(nil).
			Maybe()

		s := New(nil, repo, nil)
		res := s.reserveItems(ctx, orderID, items)

		assert.NoError(t, res)
	})

	t.Run("stocks insufficient", func(t *testing.T) {
		t.Parallel()

		stock := &domain.StockInfo{
			WarehouseID: warehouseID,
			Count:       1,
		}

		repo := mocks.NewRepository(t)
		repo.On("GetStocks", mock.Anything, item.SKU).
			Return([]*domain.StockInfo{stock}, nil).
			Once()

		s := New(nil, repo, nil)
		res := s.reserveItems(ctx, orderID, items)

		assert.ErrorIs(t, res, ErrStockInsufficient)
	})
}
