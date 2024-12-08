package domain

import (
	"time"

	"github.com/google/uuid"
)

type OrderItem struct {
	SKU   uint32
	Count uint16
}

type Status uint32

type OrderInfo struct {
	UserID int64
	Items  []*OrderItem
	Status Status
}

type StockInfo struct {
	WarehouseID int64
	Count       uint64
}

type ReserveInfo struct {
	OrderID int64
	SKU     uint32
	Items   []*StockInfo
}

type StatusMessage struct {
	MessageID uuid.UUID
	CreatedAt time.Time
	OrderID   int64
	StatusOld Status
	StatusNew Status
}
