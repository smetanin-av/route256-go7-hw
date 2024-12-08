package repository

import (
	"context"
	"fmt"

	"route256/libs/tx_manager"
	"route256/notifications/internal/app/domain"
)

type Repository struct {
	provider tx_manager.DBProvider
}

func New(provider tx_manager.DBProvider) *Repository {
	return &Repository{provider: provider}
}

func (r *Repository) SaveMessage(ctx context.Context, msg *domain.Message) error {
	db := r.provider.GetDB(ctx)

	_, err := db.Exec(ctx, `
INSERT INTO messages(message_id, created_at, order_id, status_old, status_new)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (message_id)
DO NOTHING;
`,
		msg.MessageID, msg.CreatedAt, msg.OrderID, msg.StatusOld, msg.StatusNew)
	if err != nil {
		return fmt.Errorf("exec insert messages: %w", err)
	}

	return nil
}
