package kafka

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Shopify/sarama"
	"route256/libs/kafka/consumer"
)

type CallbackFn[T any] func(ctx context.Context, msg *T) error

func Subscribe[T any](ctx context.Context, impl *consumer.Consumer, topic string, callback CallbackFn[T]) {
	impl.Subscribe(ctx, topic, func(ctx context.Context, msg *sarama.ConsumerMessage) error {
		var model T
		err := json.Unmarshal(msg.Value, &model)
		if err != nil {
			return fmt.Errorf("unmarshal message: %w", err)
		}

		err = callback(ctx, &model)
		if err != nil {
			return fmt.Errorf("handle message: %w", err)
		}

		return nil
	})
}
