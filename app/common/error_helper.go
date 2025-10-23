package common

import (
	e "boilerplate/internal/errors"

	"github.com/gofiber/fiber/v2"
)

var HttpErrorMap = map[e.GlobalErrorCode]int{

	// 400
	e.BadRequestCode: fiber.StatusBadRequest,

	// 409
	e.ConflictErrorCode:     fiber.StatusConflict,
	e.UserAlreadyExistsCode: fiber.StatusConflict,

	// 500
	e.InternalServerErrorCode: fiber.StatusInternalServerError,
}
