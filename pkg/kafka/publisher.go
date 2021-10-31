package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
)

type PublisherConfig struct {
	Brokers []string
	Debug   bool
	Trace   bool
}

func NewPublisher(cfg *PublisherConfig) (*kafka.Publisher, error) {
	saramaCfg := kafka.DefaultSaramaSyncPublisherConfig()
	saramaCfg.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	saramaCfg.Producer.Return.Errors = true

	publisher, err := kafka.NewPublisher(
		kafka.PublisherConfig{
			Brokers:   cfg.Brokers,
			Marshaler: CustomMarshaler{},
		},
		watermill.NewStdLogger(cfg.Debug, cfg.Trace),
	)
	if err != nil {
		return nil, err
	}

	return publisher, nil
}
