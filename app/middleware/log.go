package middleware

import (
	"boilerplate/app/common"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func RequestLogger(ctx *fiber.Ctx) error {
	common.Logger.Info(
		"Request received",
		zap.String("method", ctx.Method()),
		zap.String("path", ctx.Path()),
		zap.String("ip", ctx.IP()),
	)
	return ctx.Next()
}
