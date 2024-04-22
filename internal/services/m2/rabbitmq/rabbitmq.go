package rabbitmq

import (
	"context"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

func NewRabbitMQConn(connAddr string, ctx context.Context, l *zap.Logger) (*amqp.Connection, error) {
	var conn *amqp.Connection
	var err error

	l.Info("Connected to RabbitMQ")

	conn, err = amqp.Dial(connAddr)
	if err != nil {
		l.Error("Failed to connect to RabbitMQ")
		return nil, err
	}
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
