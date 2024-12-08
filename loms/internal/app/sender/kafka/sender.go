package kafka

import (
	"encoding/json"
	"fmt"

	"github.com/Shopify/sarama"
	"route256/libs/kafka/producer"
)

type Sender struct {
	impl  *producer.Producer
	topic string
}

func New(impl *producer.Producer, topic string) *Sender {
	return &Sender{
		impl:  impl,
		topic: topic,
	}
}

func (s *Sender) SendMessage(key fmt.Stringer, msg any) error {
	value, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("json marshal: %w", err)
	}

	_, _, err = s.impl.SendSyncMessage(&sarama.ProducerMessage{
		Topic:     s.topic,
		Value:     sarama.ByteEncoder(value),
		Partition: -1,
		Key:       sarama.StringEncoder(key.String()),
	})
	if err != nil {
		return fmt.Errorf("send sync message: %w", err)
	}

	return nil
}
