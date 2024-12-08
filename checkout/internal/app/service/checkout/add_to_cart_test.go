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

func Test_isStockInsufficient(t *testing.T) {
	t.Parallel()

	type args struct {
		stocks         []domain.StockInfo
		requestedCount uint64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "warehouses: 1, total > count",
			args: args{
				stocks: []domain.StockInfo{
					{Count: 2},
				},
				requestedCount: 1,
			},
			want: false,
		},
		{
			name: "warehouses: 1, total = count",
			args: args{
				stocks: []domain.StockInfo{
					{Count: 1},
				},
				requestedCount: 1,
			},
			want: false,
		},
		{
			name: "warehouses: 1, total < count",
			args: args{
				stocks: []domain.StockInfo{
					{Count: 1},
				},
				requestedCount: 2,
			},
			want: true,
		},
		{
			name: "warehouses: 2, total > count",
			args: args{
				stocks: []domain.StockInfo{
					{Count: 1},
					{Count: 2},
				},
				requestedCount: 2,
			},
			want: false,
		},
		{
			name: "warehouses: 2, total = count",
			args: args{
				stocks: []domain.StockInfo{
					{Count: 1},
					{Count: 1},
				},
				requestedCount: 2,
			},
			want: false,
		},
		{
			name: "warehouses: 2, total < count",
			args: args{
				stocks: []domain.StockInfo{
					{Count: 1},
					{Count: 1},
				},
				requestedCount: 3,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := isStockInsufficient(tt.args.stocks, tt.args.requestedCount)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestService_AddToCart(t *testing.T) {
	t.Parallel()

	var (
		ctx    = context.Background()
		change = &domain.UpdateCart{
			SKU:   1,
			Count: 3,
		}
		stocks = []domain.StockInfo{
			{Count: 2},
			{Count: 2},
		}
		testErr = errors.New("test error")
	)

	t.Run("happy path", func(t *testing.T) {
		t.Parallel()

		lomsCli := mocks.NewLomsClient(t)
		lomsCli.On("GetStocks", mock.Anything, change.SKU).
			Return(stocks, nil).
			Once()

		repo := mocks.NewRepository(t)
		repo.On("AddCartItem", mock.Anything, change).
			Return(nil).
			Once()

		s := New(lomsCli, nil, nil, repo)
		err := s.AddToCart(ctx, change)

		assert.NoError(t, err)
	})

	t.Run("stock insufficient", func(t *testing.T) {
		t.Parallel()

		lomsCli := mocks.NewLomsClient(t)
		lomsCli.On("GetStocks", mock.Anything, change.SKU).
			Return(stocks, nil).
			Once()

		s := New(lomsCli, nil, nil, nil)
		err := s.AddToCart(ctx, &domain.UpdateCart{
			SKU:   change.SKU,
			Count: 5,
		})

		assert.ErrorIs(t, err, ErrStockInsufficient)
	})

	t.Run("get stocks failed", func(t *testing.T) {
		t.Parallel()

		lomsCli := mocks.NewLomsClient(t)
		lomsCli.On("GetStocks", mock.Anything, change.SKU).
			Return(nil, testErr).
			Once()

		s := New(lomsCli, nil, nil, nil)
		err := s.AddToCart(ctx, change)

		assert.ErrorIs(t, err, testErr)
	})

	t.Run("add cart item failed", func(t *testing.T) {
		t.Parallel()

		lomsCli := mocks.NewLomsClient(t)
		lomsCli.On("GetStocks", mock.Anything, change.SKU).
			Return(stocks, nil).
			Once()

		repo := mocks.NewRepository(t)
		repo.On("AddCartItem", mock.Anything, change).
			Return(testErr).
			Once()

		s := New(lomsCli, nil, nil, repo)
		err := s.AddToCart(ctx, change)

		assert.ErrorIs(t, err, testErr)
	})
}
