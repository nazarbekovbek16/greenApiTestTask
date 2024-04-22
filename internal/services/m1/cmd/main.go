package main

import (
	"context"
	"fmt"
	"github.com/nazarbekovbek16/greenApiTestTask/internal/services/m1/config"
	"github.com/nazarbekovbek16/greenApiTestTask/internal/services/m1/logger"
	"github.com/nazarbekovbek16/greenApiTestTask/internal/services/m1/rabbitmq"
	"github.com/nazarbekovbek16/greenApiTestTask/internal/services/m1/service"
	"github.com/nazarbekovbek16/greenApiTestTask/internal/services/m1/transport"
	"github.com/nazarbekovbek16/greenApiTestTask/internal/services/m1/transport/handlers"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	conf := config.NewConfig()

	ctx, cancel := context.WithCancel(context.Background())

	l, err := logger.Init(conf.Level)

	rabbitMQ, err := rabbitmq.NewRabbitMQConn(conf.RabbitMQ, ctx, l)

	if err != nil {
		return fmt.Errorf("cannot init logger: %w", err)
	}
	defer func(l *zap.Logger) {
		err = l.Sync()
		if err != nil {
			log.Fatalln(err)
		}
	}(l)

	defer cancel()

	gracefulShutdown(cancel, l)

	service, err := service.NewService(l, rabbitMQ)
	if err != nil {
		return err
	}

	handler := handlers.NewHandler(l, conf, service)

	srv := transport.NewServer(handler, conf)

	l.Info("Start server")
	err = srv.Run(ctx)
	if err != nil {
		return err
	}

	return nil
}

func gracefulShutdown(ctx context.CancelFunc, l *zap.Logger) {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)
	go func() {
		sig := <-done
		l.Info("Received signal", zap.String("signal", sig.String()))
		l.Info("Gracefully shutdown")
		ctx()
	}()
}
