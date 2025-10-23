package middleware

import (
	"fmt"

	"boilerplate/app/common"
	e "boilerplate/infra/errors"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	instance := ctx.OriginalURL()
	type_ := "about:blank"

	fmt.Println("Error: ", err)
	if appErr, ok := err.(*e.GlobalError); ok {

		if appErr.Type != "" {
			type_ = appErr.Type
		}

		status := common.HttpErrorMap[appErr.Code]
		if status == 0 {
			status = fiber.StatusInternalServerError
		}
		problemDetail := e.NewProblemDetail(
			type_,
			appErr.Title,
			status,
			appErr.Detail,
			instance,
		)
		return ctx.Status(status).JSON(problemDetail)
	}

	if validationErr, ok := err.(ValidationError); ok {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"title":    "Validation error",
				"status":   fiber.StatusUnprocessableEntity,
				"detail":   validationErr.Error(),
				"instance": instance,
				"errors":   validationErr.List,
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
		),
	)
}
