package checkout

import (
	"context"
	"fmt"

	"route256/checkout/internal/app/domain"
	"route256/libs/tx_manager"
)

type Repository struct {
	provider tx_manager.DBProvider
}

func New(provider tx_manager.DBProvider) *Repository {
	return &Repository{provider: provider}
}

func (r *Repository) AddCartItem(ctx context.Context, change *domain.UpdateCart) error {
	db := r.provider.GetDB(ctx)

	_, err := db.Exec(ctx, `
INSERT INTO carts(user_id, sku, quantity)
VALUES ($1, $2, $3)
ON CONFLICT (user_id, sku)
DO UPDATE
SET quantity=carts.quantity+$3;
`,
		change.UserID, change.SKU, change.Count)
	if err != nil {
		return fmt.Errorf("exec upsert carts: %w", err)
	}

	return nil
}

func (r *Repository) CountOfSKU(ctx context.Context, userID int64, sku uint32) (uint16, error) {
	db := r.provider.GetDB(ctx)

	var count uint16
	err := db.QueryRow(ctx, `
SELECT quantity
FROM carts
WHERE user_id = $1 AND sku = $2;
`,
		userID, sku,
	).Scan(
		&count,
	)
	if err != nil {
		return 0, fmt.Errorf("query select carts: %w", err)
	}

	return count, nil
}

func (r *Repository) DecreaseCount(ctx context.Context, change *domain.UpdateCart) error {
	db := r.provider.GetDB(ctx)

	_, err := db.Exec(ctx, `
UPDATE carts
SET quantity = quantity - $3
WHERE user_id = $1 AND sku = $2;
`,
		change.UserID, change.SKU, change.Count,
	)
	if err != nil {
		return fmt.Errorf("exec update carts: %w", err)
	}

	return nil
}

func (r *Repository) DeleteCartItem(ctx context.Context, userID int64, sku uint32) error {
	db := r.provider.GetDB(ctx)

	_, err := db.Exec(ctx, `
DELETE
FROM carts
WHERE user_id = $1 AND sku = $2;
`,
		userID, sku,
	)
	if err != nil {
		return fmt.Errorf("exec delete carts: %w", err)
	}

	return nil
}

func (r *Repository) GetItemsInCart(ctx context.Context, userID int64) ([]*domain.OrderItem, error) {
	db := r.provider.GetDB(ctx)

	rows, err := db.Query(ctx, `
SELECT sku, quantity
FROM carts
WHERE user_id = $1
`,
		userID,
	)
	if err != nil {
		return nil, fmt.Errorf("query select carts: %w", err)
	}
	defer rows.Close()

	var items []*domain.OrderItem
	for rows.Next() {
		item := new(domain.OrderItem)
		err := rows.Scan(&item.SKU, &item.Count)
		if err != nil {
			return nil, fmt.Errorf("rows scan: %w", err)
		}
		items = append(items, item)
	}

	return items, nil
}

func (r *Repository) ClearCart(ctx context.Context, userID int64) error {
	db := r.provider.GetDB(ctx)

	_, err := db.Exec(ctx, `
DELETE
FROM carts
WHERE user_id = $1;
`,
		userID,
	)
	if err != nil {
		return fmt.Errorf("exec delete carts: %w", err)
	}

	return nil
}
