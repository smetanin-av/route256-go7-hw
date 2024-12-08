package consumer

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Shopify/sarama"
)

type CallbackFn func(ctx context.Context, msg *sarama.ConsumerMessage) error

type Consumer struct {
	client sarama.ConsumerGroup
}

func New(brokers []string, groupID string) (*Consumer, error) {
	cfg := sarama.NewConfig()
	cfg.Version = sarama.MaxVersion

	cfg.Consumer.Offsets.Initial = sarama.OffsetOldest

	cfg.Consumer.Group.ResetInvalidOffsets = true
	cfg.Consumer.Group.Heartbeat.Interval = 3 * time.Second
	cfg.Consumer.Group.Session.Timeout = 60 * time.Second
	cfg.Consumer.Group.Rebalance.Timeout = 60 * time.Second

	client, err := sarama.NewConsumerGroup(brokers, groupID, cfg)
	if err != nil {
		return nil, fmt.Errorf("sarama new consumer: %w", err)
	}

	return &Consumer{
		client: client,
	}, err
}

func (c *Consumer) Subscribe(ctx context.Context, topic string, callback CallbackFn) {
	instance := &handler{callback: callback}
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				break
			}

			err := c.client.Consume(ctx, []string{topic}, instance)
			if err != nil {
				log.Println("ERROR init consume:", err)
			}
		}
	}()
}

func (c *Consumer) Close() {
	err := c.client.Close()
	if err != nil {
		log.Println("ERROR consumer close:", err)
	}
}

type handler struct {
	callback CallbackFn
}

func (c *handler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (c *handler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (c *handler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		ctx := session.Context()
		select {
		case <-ctx.Done():
			return nil
		case msg := <-claim.Messages():
			err := c.callback(ctx, msg)
			if err != nil {
				log.Println("ERROR process message:", err)
				return err
			}
			session.MarkMessage(msg, "")
		}
	}
}
