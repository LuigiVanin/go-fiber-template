package user

import (
	"boilerplate/app/modules/user/controller"
	"boilerplate/app/modules/user/repository"
	"boilerplate/app/modules/user/service"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

var Module = fx.Module("user",
	fx.Provide(
		fx.Annotate(
			repository.NewUserRepository,
			fx.As(new(repository.IUserRepository)),
		),
	),

	fx.Provide(
		fx.Annotate(
			service.NewUserService,
			fx.As(new(service.IUserService)),
		),
	),

	fx.Provide(controller.NewUserController),

	fx.Invoke(func(server *fiber.App, controller *controller.UserController) {
		controller.Register(server)
	}),
)
