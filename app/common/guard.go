package common

import "github.com/gofiber/fiber/v2"

type IGuard interface {
	Activate(ctx *fiber.Ctx) error
}
