package producer

import (
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

type Producer struct {
	client sarama.SyncProducer
}

func newSyncProducer(brokers []string) (sarama.SyncProducer, error) {
	cfg := sarama.NewConfig()

	cfg.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	cfg.Producer.RequiredAcks = sarama.WaitForAll

	cfg.Producer.Idempotent = true
	cfg.Net.MaxOpenRequests = 1

	cfg.Producer.CompressionLevel = sarama.CompressionLevelDefault
	cfg.Producer.Compression = sarama.CompressionGZIP

	cfg.Producer.Return.Successes = true
	cfg.Producer.Return.Errors = true

	producer, err := sarama.NewSyncProducer(brokers, cfg)
	if err != nil {
		return nil, fmt.Errorf("sarama new producer: %w", err)
	}

	return producer, nil
}

func New(brokers []string) (*Producer, error) {
	producer, err := newSyncProducer(brokers)
	if err != nil {
		return nil, err
	}

	return &Producer{
		client: producer,
	}, nil
}

func (p *Producer) SendSyncMessage(msg *sarama.ProducerMessage) (int32, int64, error) {
	return p.client.SendMessage(msg)
}

func (p *Producer) Close() {
	err := p.client.Close()
	if err != nil {
		log.Println("ERROR producer close:", err)
	}
}
