package logger

import (
	"route256/libs/logging"
	"route256/notifications/internal/app/domain"
)

type Notifier struct {
}

func New() *Notifier {
	return &Notifier{}
}

func (n *Notifier) SendMessage(msg *domain.Message) error {
	logging.Infof("send message: %+v", msg)
	return nil
}
