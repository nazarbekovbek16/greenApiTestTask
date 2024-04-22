package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/nazarbekovbek16/greenApiTestTask/internal/services/m1/config"
	"github.com/nazarbekovbek16/greenApiTestTask/internal/services/m1/service"
	"go.uber.org/zap"
	"net/http"
)

type TestHandler struct {
	testManager *service.Service
	config      *config.Config
	log         *zap.Logger
}

func NewTestHandler(logger *zap.Logger, config *config.Config, testManager *service.Service) *TestHandler {
	return &TestHandler{log: logger, config: config, testManager: testManager}
}
func (h *TestHandler) DoubleInteger(c echo.Context) error {
	param := c.FormValue("param")
	result, err := h.testManager.Test.Double(c.Request().Context(), param, h.config.TaskQueueName)
	if err != nil {
		h.log.Error("Error occurred", zap.Error(err))
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.String(http.StatusOK, fmt.Sprintf("Result: %s", result))
}
