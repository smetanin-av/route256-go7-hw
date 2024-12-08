package service

import (
	"context"
	"fmt"

	"route256/notifications/internal/app/domain"
)

type TxManager interface {
	RunRepeatableRead(ctx context.Context, fn func(ctxTx context.Context) error) error
}

type Repository interface {
	SaveMessage(ctx context.Context, msg *domain.Message) error
}

type Notifier interface {
	SendMessage(msg *domain.Message) error
}

type Service struct {
	txMan    TxManager
	repo     Repository
	notifier Notifier
}

func New(txMan TxManager, repo Repository, notifier Notifier) *Service {
	return &Service{
		txMan:    txMan,
		repo:     repo,
		notifier: notifier,
	}
}

func (s *Service) ProcessMsg(ctx context.Context, msg *domain.Message) error {
	err := s.txMan.RunRepeatableRead(ctx, func(ctxTx context.Context) error {
		err := s.repo.SaveMessage(ctxTx, msg)
		if err != nil {
			return fmt.Errorf("save message: %w", err)
		}

		err = s.notifier.SendMessage(msg)
		if err != nil {
			return fmt.Errorf("send message: %w", err)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("run repeatable read: %w", err)
	}

	return nil
}
