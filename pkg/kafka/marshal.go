package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
)

type CustomMarshaler struct{}

func (CustomMarshaler) Marshal(topic string, msg *message.Message) (*sarama.ProducerMessage, error) {
	saramaMsg, err := kafka.DefaultMarshaler.Marshal(kafka.DefaultMarshaler{}, topic, msg)
	if err != nil {
		return nil, err
	}

	switch {
	case msg.UUID != "":
		saramaMsg.Key = sarama.StringEncoder(msg.UUID)
	default:
		saramaMsg.Key = sarama.StringEncoder(watermill.NewUUID())
	}

	return saramaMsg, nil
}

func (CustomMarshaler) Unmarshal(kafkaMsg *sarama.ConsumerMessage) (*message.Message, error) {
	return kafka.DefaultMarshaler.Unmarshal(kafka.DefaultMarshaler{}, kafkaMsg)
}
