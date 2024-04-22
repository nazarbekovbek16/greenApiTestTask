package service

import (
	"context"
	"encoding/json"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

type TestService struct {
	rabbitMQ *amqp.Connection
	l        *zap.Logger
}

func NewTestService(rabbitMQ *amqp.Connection, l *zap.Logger) *TestService {
	return &TestService{rabbitMQ: rabbitMQ, l: l}
}

type ITestService interface {
	Double(ctx context.Context, number, taskQueueName string) (string, error)
}

func (s TestService) Double(ctx context.Context, param, taskQueueName string) (string, error) {
	var result string
	ch, err := s.rabbitMQ.Channel()
	if err != nil {
		s.l.Error("Failed to open a channel", zap.Error(err))
		return result, err
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
		s.l.Error("Failed to declare a queue", zap.Error(err))
		return result, err
	}

	body, err := json.Marshal(map[string]string{"param": param})
	if err != nil {
		s.l.Error("Failed to marshal JSON: %v", zap.Error(err))
		return result, err
	}

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		s.l.Error("Failed to publish a message: %v", zap.Error(err))
		return result, err
	}

	s.l.Info("Published message: " + string(body))

	// Ожидаем ответа от микросервиса М2
	msgs, err := ch.Consume(
		"response_queue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		s.l.Error("Failed to register a consumer: ", zap.Error(err))
	}

	// Получаем и обрабатываем ответ от микросервиса М2
	for d := range msgs {
		result = string(d.Body)

		// Возвращаем результат HTTP запроса
		return result, nil
	}
	return result, nil
}
