package loms

import (
	"context"
	"fmt"
	"time"

	"github.com/georgysavva/scany/pgxscan"
	"route256/libs/tx_manager"
	"route256/loms/internal/app/domain"
)

type Repository struct {
	provider tx_manager.DBProvider
}

func New(provider tx_manager.DBProvider) *Repository {
	return &Repository{provider: provider}
}

func (r Repository) AddOrder(ctx context.Context, userID int64, items []domain.OrderItem) (int64, error) {
	db := r.provider.GetDB(ctx)

	var orderID int64
	err := db.QueryRow(ctx, `
INSERT INTO orders(user_id, status, created_at)
VALUES ($1, $2, $3)
RETURNING order_id;
`,
		userID, domain.StatusAwaitingPayment, time.Now(),
	).Scan(
		&orderID,
	)
	if err != nil {
		return 0, fmt.Errorf("query insert orders: %w", err)
	}

	for _, item := range items {
		_, err := db.Exec(ctx, `
INSERT INTO items(order_id, sku, quantity)
VALUES ($1, $2, $3)
`,
			orderID, item.SKU, item.Count,
		)
		if err != nil {
			return 0, fmt.Errorf("exec insert items: %w", err)
		}
	}

	return orderID, nil
}

func (r Repository) ListOrder(ctx context.Context, orderID int64) (*domain.OrderInfo, error) {
	db := r.provider.GetDB(ctx)

	var order domain.OrderInfo
	err := db.QueryRow(ctx, `
SELECT user_id, status
FROM orders
WHERE order_id = $1
`,
		orderID,
	).Scan(
		&order.UserID, &order.Status,
	)
	if err != nil {
		return nil, fmt.Errorf("query select orders: %w", err)
	}

	rows, err := db.Query(ctx, `
SELECT sku, quantity
FROM items
WHERE order_id = $1
`,
		orderID,
	)
	if err != nil {
		return nil, fmt.Errorf("query select items: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		item := new(domain.OrderItem)
		err := rows.Scan(&item.SKU, &item.Count)
		if err != nil {
			return nil, fmt.Errorf("rows scan: %w", err)
		}
		order.Items = append(order.Items, item)
	}

	return &order, nil
}

func (r Repository) GetStatus(ctx context.Context, orderID int64) (domain.Status, error) {
	db := r.provider.GetDB(ctx)

	var status domain.Status
	err := db.QueryRow(ctx, `
SELECT status
FROM orders
WHERE order_id = $1
`,
		orderID,
	).Scan(
		&status,
	)
	if err != nil {
		return domain.StatusInvalid, fmt.Errorf("query select orders: %w", err)
	}

	return status, nil
}

func (r Repository) SetStatus(ctx context.Context, orderID int64, status domain.Status) error {
	db := r.provider.GetDB(ctx)

	_, err := db.Exec(ctx, `
UPDATE orders
SET status = $1
WHERE order_id = $2;
`,
		status, orderID,
	)
	if err != nil {
		return fmt.Errorf("exec update orders: %w", err)
	}

	return nil
}

func (r Repository) GetStocks(ctx context.Context, sku uint32) ([]*domain.StockInfo, error) {
	db := r.provider.GetDB(ctx)

	rows, err := db.Query(ctx, `
SELECT warehouse_id, quantity 
FROM stocks
WHERE sku = $1 AND quantity > 0
`,
		sku,
	)
	if err != nil {
		return nil, fmt.Errorf("query select stocks: %w", err)
	}
	defer rows.Close()

	var stocks []*domain.StockInfo
	for rows.Next() {
		stock := new(domain.StockInfo)
		err := rows.Scan(&stock.WarehouseID, &stock.Count)
		if err != nil {
			return nil, fmt.Errorf("rows scan: %w", err)
		}
		stocks = append(stocks, stock)
	}

	return stocks, nil
}

func (r Repository) ReserveStocks(ctx context.Context, info *domain.ReserveInfo) error {
	db := r.provider.GetDB(ctx)

	for _, item := range info.Items {
		_, err := db.Exec(ctx, `
INSERT INTO reserved(order_id, sku, warehouse_id, quantity)
VALUES ($1, $2, $3, $4)
`,
			info.OrderID, info.SKU, item.WarehouseID, item.Count,
		)
		if err != nil {
			return fmt.Errorf("exec insert reserved: %w", err)
		}

		_, err = db.Exec(ctx, `
UPDATE stocks
SET quantity = quantity - $1
WHERE sku = $2 AND warehouse_id = $3
`,
			item.Count, info.SKU, item.WarehouseID,
		)
		if err != nil {
			return fmt.Errorf("exec update stocks: %w", err)
		}
	}

	return nil
}

func (r Repository) CancelReserve(ctx context.Context, orderID int64) error {
	db := r.provider.GetDB(ctx)

	var items []struct {
		SKU         int64 `db:"sku"`
		WarehouseID int64 `db:"warehouse_id"`
		Count       int64 `db:"quantity"`
	}
	err := pgxscan.Select(ctx, db, &items, `
SELECT sku, warehouse_id, quantity
FROM reserved
WHERE order_id = $1
`,
		orderID,
	)
	if err != nil {
		return fmt.Errorf("query select reserved: %w", err)
	}

	for _, item := range items {
		_, err := db.Exec(ctx, `
INSERT INTO stocks(sku, warehouse_id, quantity)
VALUES ($1, $2, $3)
ON CONFLICT (sku, warehouse_id)
DO UPDATE 
SET quantity=stocks.quantity+$3
`,
			item.SKU, item.WarehouseID, item.Count,
		)
		if err != nil {
			return fmt.Errorf("exec upsert stocks: %w", err)
		}
	}

	_, err = db.Exec(ctx, `
DELETE
FROM reserved
WHERE order_id = $1;
`,
		orderID,
	)
	if err != nil {
		return fmt.Errorf("exec delete reserved: %w", err)
	}

	return nil
}

func (r Repository) MarkAsSold(ctx context.Context, orderID int64) error {
	db := r.provider.GetDB(ctx)

	_, err := db.Exec(ctx, `
INSERT INTO sold(order_id, sku, warehouse_id, quantity)
SELECT order_id, sku, warehouse_id, quantity
FROM reserved
WHERE order_id = $1;
`,
		orderID,
	)
	if err != nil {
		return fmt.Errorf("exec insert sold: %w", err)
	}

	_, err = db.Exec(ctx, `
DELETE
FROM reserved
WHERE order_id = $1;
`,
		orderID,
	)
	if err != nil {
		return fmt.Errorf("exec delete reserved: %w", err)
	}

	_, err = db.Exec(ctx, `
DELETE
FROM stocks
WHERE quantity = 0;
`)
	if err != nil {
		return fmt.Errorf("exec delete stocks: %w", err)
	}

	return nil
}

func (r Repository) GetOrdersToCleanUp(ctx context.Context, delta time.Duration) ([]int64, error) {
	db := r.provider.GetDB(ctx)

	rows, err := db.Query(ctx, `
SELECT order_id 
FROM orders
WHERE created_at < $1::timestamp AND status = $2;
`,
		time.Now().Add(-delta), domain.StatusAwaitingPayment,
	)
	if err != nil {
		return nil, fmt.Errorf("query select orders: %w", err)
	}
	defer rows.Close()

	var ordersIDs []int64
	for rows.Next() {
		var orderID int64
		err := rows.Scan(&orderID)
		if err != nil {
			return nil, fmt.Errorf("rows scan: %w", err)
		}
		ordersIDs = append(ordersIDs, orderID)
	}

	return ordersIDs, nil
}
