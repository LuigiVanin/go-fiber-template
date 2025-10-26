package middleware

import (
	"boilerplate/app/common"
	e "boilerplate/infra/errors"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func LogError(ctx *fiber.Ctx, message string, fields ...zap.Field) {
	defaultFields := []zap.Field{
		zap.String("method", ctx.Method()),
		zap.String("path", ctx.Path()),
		zap.String("ip", ctx.IP()),
	}

	allFields := append(defaultFields, fields...)
	common.Logger.Error(message, allFields...)
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	instance := ctx.OriginalURL()

	if appErr, ok := err.(*e.GlobalError); ok {
		problemDetail := appErr.IntoProblemDetail(instance)

		LogError(ctx, "Request error",
			zap.String("code", string(appErr.Code)),
			zap.Int("status", problemDetail.Status),
			zap.String("detail", appErr.Detail),
		)

		return ctx.
			Status(problemDetail.Status).
			JSON(problemDetail)
	}

	if validationErr, ok := err.(ValidationError); ok {
		LogError(ctx, "Validation error",
			zap.String("detail", validationErr.Error()),
		)

		return ctx.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"title":    "Validation error",
				"status":   fiber.StatusUnprocessableEntity,
				"detail":   validationErr.Error(),
				"instance": instance,
				"errors":   validationErr.List,
				"code":     string(e.BadRequestCode),
			},
		)
	}

	LogError(ctx, "Unexpected error",
		zap.Error(err),
	)

	return ctx.Status(fiber.StatusInternalServerError).JSON(
		e.NewProblemDetail(
			"about:blank",
			"Unexpected Internal Error",
			fiber.StatusInternalServerError,
			err.Error(),
			instance,
			"",
		),
	)
}
