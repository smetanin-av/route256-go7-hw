package tx_manager

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"route256/libs/logging"
)

var txKey struct{}

type Manager struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Manager {
	return &Manager{pool: pool}
}

type DBProvider interface {
	GetDB(ctx context.Context) Querier
}

type Querier interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, optionsAndArgs ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, optionsAndArgs ...any) pgx.Row
}

func (m *Manager) RunRepeatableRead(ctx context.Context, fn func(ctxTx context.Context) error) error {
	tx, err := m.pool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.RepeatableRead,
	})
	if err != nil {
		return fmt.Errorf("start tx: %w", err)
	}
	defer func() {
		err := tx.Rollback(ctx)
		if err != nil && !errors.Is(err, pgx.ErrTxClosed) {
			logging.Errorf("rollback transaction: %v", err)
		}
	}()

	ctxTx := context.WithValue(ctx, txKey, tx)
	if err = fn(ctxTx); err != nil {
		return fmt.Errorf("exec body: %w", err)
	}

	if err = tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}

	return nil
}

func (m *Manager) GetDB(ctx context.Context) Querier {
	tx, ok := ctx.Value(txKey).(Querier)
	if ok {
		return tx
	}

	return m.pool
}
