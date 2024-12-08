package mocks

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
)

func NewProxyTxManager(t *testing.T) *TxManager {
	manager := NewTxManager(t)

	manager.On("RunRepeatableRead", mock.Anything, mock.Anything).
		Return(func(ctx context.Context, fn func(ctxTx context.Context) error) error {
			return fn(ctx)
		}).
		Once()

	return manager
}
