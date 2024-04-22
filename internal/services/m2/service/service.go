package service

import (
	"fmt"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

type Service struct {
	rabbitMq *amqp.Connection
	Test     ITestService
}

func NewService(logger *zap.Logger, rabbitMq *amqp.Connection) (*Service, error) {
	var service Service

	if rabbitMq == nil {
		logger.Error("RabbitMQ pointer is empty")
		return nil, fmt.Errorf("RabbitMQ pointer is empty")
	}
	service.Test = NewTestService(rabbitMq, logger)
	return &service, nil
}
