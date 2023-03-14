package infrastructure

import (
	"github.com/exepirit/sd-gallery/internal/config"
	"go.uber.org/zap"
)

func NewLogger(cfg *config.Config) *zap.Logger {
	logger, _ := zap.NewDevelopment()
	return logger
}
