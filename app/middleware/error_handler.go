package middleware

import (
	"fmt"

	e "boilerplate/infra/errors"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	instance := ctx.OriginalURL()
	fmt.Println("Error: ", err)
	if appErr, ok := err.(*e.GlobalError); ok {

		problemDetail := appErr.IntoProblemDetail(instance)

		return ctx.
			Status(problemDetail.Status).
			JSON(problemDetail)
	}

	if validationErr, ok := err.(ValidationError); ok {
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
