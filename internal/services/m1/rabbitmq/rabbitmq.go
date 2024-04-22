package rabbitmq

import (
	"context"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

func NewRabbitMQConn(connAddr string, ctx context.Context, l *zap.Logger) (*amqp.Connection, error) {
	conn, err := amqp.Dial(connAddr)
	if err != nil {
		l.Error("Failed to connect to RabbitMQ")
		return nil, err
	}
	l.Info("Connected to RabbitMQ")
	go func() {
		select {
		case <-ctx.Done():
			err := conn.Close()
			if err != nil {
				l.Error("Failed to close RabbitMQ connection")
			}
			l.Info("RabbitMQ connection is closed")
		}
	}()

	return conn, err
}
