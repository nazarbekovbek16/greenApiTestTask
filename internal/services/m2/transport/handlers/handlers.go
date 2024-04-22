package handlers

import (
	"github.com/nazarbekovbek16/greenApiTestTask/internal/services/m2/config"
	"github.com/nazarbekovbek16/greenApiTestTask/internal/services/m2/service"
	"go.uber.org/zap"
)

type Handler struct {
	Test *TestHandler
	log  *zap.Logger
}

func NewHandler(l *zap.Logger, config *config.Config, service *service.Service) *Handler {
	return &Handler{
		Test: NewTestHandler(l, config, service),
		log:  l,
	}
}
