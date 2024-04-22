package handlers

import (
	"context"
	"github.com/nazarbekovbek16/greenApiTestTask/internal/services/m2/config"
	"github.com/nazarbekovbek16/greenApiTestTask/internal/services/m2/service"
	"go.uber.org/zap"
)

type TestHandler struct {
	testManager *service.Service
	config      *config.Config
	log         *zap.Logger
}

func NewTestHandler(logger *zap.Logger, config *config.Config, testManager *service.Service) *TestHandler {
	return &TestHandler{log: logger, config: config, testManager: testManager}
}

func (h *TestHandler) StartMQConsumer(ctx context.Context) error {
	err := h.testManager.Test.ConsumeTasks(ctx, h.config.TaskQueueName, h.config.ResponseQueueName)
	if err != nil {
		h.log.Error("Error occurred", zap.Error(err))
		return err
	}
	return err
}
