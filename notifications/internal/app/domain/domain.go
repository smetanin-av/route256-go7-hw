package domain

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	MessageID uuid.UUID
	CreatedAt time.Time
	OrderID   int64
	StatusOld string
	StatusNew string
}
