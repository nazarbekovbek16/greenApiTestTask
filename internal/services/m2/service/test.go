package service

import (
	"context"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"time"
)

type TestService struct {
	rabbitMQ *amqp.Connection
	l        *zap.Logger
}

func NewTestService(rabbitMQ *amqp.Connection, l *zap.Logger) *TestService {
	return &TestService{rabbitMQ: rabbitMQ, l: l}
}

type ITestService interface {
	ConsumeTasks(ctx context.Context, taskQueueName, responseQueueName string) error
}

func (s TestService) ConsumeTasks(ctx context.Context, taskQueueName, responseQueueName string) error {
	ch, err := s.rabbitMQ.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		taskQueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	for d := range msgs {
		response, err := handleMQ(d.Body)
		if err != nil {
			s.l.Error("Failed to handle message: %v", zap.Error(err))
			continue
		}

		time.Sleep(5 * time.Second)

		responseQueue, err := ch.QueueDeclare(
			responseQueueName,
			true,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			s.l.Error("Failed to declare response queue: %v", zap.Error(err))
			continue
		}

		err = ch.Publish(
			"",
			responseQueue.Name,
			false,
			false,
			amqp.Publishing{
				ContentType:   "text/plain",
				CorrelationId: d.CorrelationId,
				Body:          response,
			},
		)
		if err != nil {
			s.l.Error("Failed to publish response: %v", zap.Error(err))
			continue
		}

		s.l.Info("Processed message: " + string(d.Body))
		d.Ack(false)
	}

	return nil
}
