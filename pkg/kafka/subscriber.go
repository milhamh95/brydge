package kafka

import (
	"time"

	"github.com/Shopify/sarama"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
)

// // SubscriberHandler is a function signature for kafka subscriber message handler
// type SubscriberHandler func(ctx context.Context, msg *message.Message)

// type Subscriber struct {

// }

// SubscriberConfig is a config for kafka subscriber
type SubscriberConfig struct {
	Brokers       []string
	ConsumerGroup string
	Debug         bool
	Trace         bool
	FetchMinBytes int32
	FetchMaxBytes int32
	MaxWaitTime   int
}

// NewSubscriber will return initiate kafka subscriber
func NewSubscriber(cfg *SubscriberConfig) (*kafka.Subscriber, error) {
	saramaCfg := kafka.DefaultSaramaSubscriberConfig()

	// initial offset means subscriber will receive messages
	// from the oldest offset or latest offset.
	// need to set as latest offset, so subscriber doesn't
	// receive message from the beginning of offset
	saramaCfg.Consumer.Offsets.Initial = sarama.OffsetNewest

	saramaCfg.Consumer.Fetch.Min = 1 // in bytes
	if cfg.FetchMinBytes >= 1 {
		saramaCfg.Consumer.Fetch.Min = cfg.FetchMinBytes
	}

	saramaCfg.Consumer.Fetch.Max = 10e6 // in bytes
	if cfg.FetchMaxBytes >= 1 {
		saramaCfg.Consumer.Fetch.Max = cfg.FetchMaxBytes
	}

	if cfg.MaxWaitTime >= 1 {
		// in millisecond
		saramaCfg.Consumer.MaxWaitTime = time.Duration(cfg.MaxWaitTime) * time.Millisecond
	}

	subscriber, err := kafka.NewSubscriber(
		kafka.SubscriberConfig{
			Brokers:               cfg.Brokers,
			ConsumerGroup:         cfg.ConsumerGroup,
			OverwriteSaramaConfig: saramaCfg,
			Unmarshaler:           CustomMarshaler{},
		},
		watermill.NewStdLogger(cfg.Debug, cfg.Trace),
	)
	if err != nil {
		return nil, err
	}

	return subscriber, nil
}
