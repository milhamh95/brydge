package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"

	confirmHandler "github.com/milhamh95/brydge/internal/app/confirm/handler"
	inquiryHandler "github.com/milhamh95/brydge/internal/app/inquiry/handler"
	kafkaPkg "github.com/milhamh95/brydge/pkg/kafka"
)

type App struct{}

func NewApp() *App {
	return &App{}
}

func (a *App) Start() {
	e := echo.New()
	e.GET("/ping", func(c echo.Context) error {
		type response struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}
		res := response{
			Status:  "success",
			Message: "ping",
		}

		return c.JSON(http.StatusOK, res)
	})

	kafkaPubsub, err := kafkaPkg.NewSubscriber(&kafkaPkg.SubscriberConfig{
		Brokers:       []string{},
		ConsumerGroup: "",
		Debug:         false,
		Trace:         false,
		FetchMinBytes: 10e3,
		FetchMaxBytes: 10e6,
		MaxWaitTime:   120,
	})
	if err != nil {
		log.Fatal(err)
	}

	// REST API
	restAPI := e.Group("/api")
	inquiryHandler.NewRest(nil).InitRoutes(restAPI)

	// should subscribe in separate go routine
	// so it will not blocking
	inquiryMessages, err := kafkaPubsub.Subscribe(context.Background(), "kafka-inquiry-requested")
	if err != nil {
		log.Fatal(err)
	}

	inquiryPubSub := inquiryHandler.NewPubsub(nil, "kafka-inquiry-requested")
	go inquiryPubSub.StartSubscribe(inquiryMessages, 10)

	confirmSubscriber := confirmHandler.NewSubscriber(nil, kafkaPubsub, "kafka-confirm-requested")
	go confirmSubscriber.StartSubscribe(10)

	go func() {
		err := e.Start(":8000")
		if err != nil {
			log.Fatal(err)
		}
	}()

	// graceful shutdown if there is an interruption
	// ex: ctrl + c
	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, os.Interrupt)
	<-quitCh

	log.Print("OS interrupt app")

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(3)*time.Second,
	)
	defer cancel()

	log.Print("shutting down app")
	err = e.Shutdown(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
