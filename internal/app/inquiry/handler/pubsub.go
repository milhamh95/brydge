package handler

import (
	"context"
	"encoding/json"
	"time"

	"github.com/ThreeDotsLabs/watermill/message"
)

type pubsubHandler struct {
	usecase useCase
	topic   string
}

func NewPubsub(inquiryUseCase useCase, topic string) *pubsubHandler {
	return &pubsubHandler{
		usecase: inquiryUseCase,
		topic:   topic,
	}
}

func (h *pubsubHandler) StartSubscribe(messages <-chan *message.Message, ctxTimeout int) {
	for msg := range messages {
		go func(msg *message.Message, ctxTimeout int) {
			ctx, cancel := context.WithTimeout(
				context.Background(),
				time.Duration(ctxTimeout)*time.Second,
			)

			defer cancel()
			h.Inquiry(ctx, msg)
		}(msg, ctxTimeout)
	}
}

func (h *pubsubHandler) Inquiry(ctx context.Context, msg *message.Message) error {
	var req request

	err := json.Unmarshal(msg.Payload, &req)
	if err != nil {
		return err
	}

	h.usecase.Inquiry(ctx)
	return nil
}
