package loms

import (
	"time"

	"github.com/google/uuid"
	"route256/loms/internal/app/domain"
)

func (s *Service) SendStatusMessage(orderID int64, statusOld, statusNew domain.Status) error {
	message := newStatusMessage(orderID, statusOld, statusNew)
	return s.sender.SendMessage(message.MessageID, message)
}

func newStatusMessage(orderID int64, statusOld, statusNew domain.Status) *domain.StatusMessage {
	return &domain.StatusMessage{
		MessageID: uuid.New(),
		CreatedAt: time.Now(),
		OrderID:   orderID,
		StatusOld: statusOld,
		StatusNew: statusNew,
	}
}
