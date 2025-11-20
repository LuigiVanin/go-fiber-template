package middleware

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func NewRequestLogger(logger *zap.Logger) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		logger.Info(
			"Request received",
			zap.String("method", ctx.Method()),
			zap.String("path", ctx.Path()),
			zap.String("ip", ctx.IP()),
		)
		return ctx.Next()
	}
}
