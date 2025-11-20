package auth

import (
	controller "boilerplate/app/modules/auth/controller"
	"boilerplate/app/modules/auth/service"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

var Module = fx.Module("auth",

	fx.Provide(
		fx.Annotate(
			service.New,
			fx.As(new(service.IAuthService)),
		),
	),

	fx.Provide(controller.NewAuthController),

	fx.Invoke(func(server *fiber.App, controller *controller.AuthController) {
		controller.Register(server)
	}),
)
