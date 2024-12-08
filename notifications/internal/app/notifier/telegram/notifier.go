package telegram

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"route256/notifications/internal/app/domain"
)

type Notifier struct {
	chatId int64
	api    *tgbotapi.BotAPI
}

func New(token string, chatId int64) (*Notifier, error) {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("tg new bot api: %w", err)
	}

	return &Notifier{
		chatId: chatId,
		api:    api,
	}, nil
}

func (n *Notifier) SendMessage(msg *domain.Message) error {
	text := fmt.Sprintf(`
order status changed
date: %v
order: %d
from: %s
to: %s
`,
		msg.CreatedAt, msg.OrderID, msg.StatusOld, msg.StatusNew,
	)

	res, err := n.api.Send(tgbotapi.NewMessage(n.chatId, text))
	if err != nil {
		return fmt.Errorf("send message: %w", err)
	}
	log.Printf("send response: %+v\n", res)

	return nil
}
