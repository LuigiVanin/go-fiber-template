package common

import "github.com/gofiber/fiber/v2"

type IController interface {
	Register(app *fiber.App)
}
