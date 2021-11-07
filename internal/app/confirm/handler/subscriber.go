package handler

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
)

type subscriberHandler struct {
	usecase         useCase
	kafkaSubscriber *kafka.Subscriber
	topic           string
}

func NewSubscriber(confirmUseCase useCase, kafkaSubscriber *kafka.Subscriber, topic string) *subscriberHandler {
	return &subscriberHandler{
		usecase:         confirmUseCase,
		kafkaSubscriber: kafkaSubscriber,
		topic:           topic,
	}
}

func (h *subscriberHandler) StartSubscribe(ctxTimeout int) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(ctxTimeout))
	defer cancel()

	messages, err := h.kafkaSubscriber.Subscribe(ctx, h.topic)
	if err != nil {
		log.Fatalf("subscribe message topic:%s, err:%s", h.topic, err)
	}

	for msg := range messages {
		go func(msg *message.Message, ctxTimeout int) {
			ctx, cancel := context.WithTimeout(
				context.Background(),
				time.Duration(ctxTimeout)*time.Second,
			)

			defer cancel()
			h.Confirm(ctx, msg)
		}(msg, ctxTimeout)
	}
}

func (h *subscriberHandler) Confirm(ctx context.Context, msg *message.Message) error {
	var req request

	err := json.Unmarshal(msg.Payload, &req)
	if err != nil {
		return err
	}

	h.usecase.Confirm(ctx)
	return nil
}
