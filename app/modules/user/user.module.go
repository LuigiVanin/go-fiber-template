package user

import (
	"boilerplate/app/middleware/guard"
	"boilerplate/app/modules/user/controller"
	"boilerplate/app/modules/user/repository"
	"boilerplate/app/modules/user/service"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type UserModule struct {
	server         *fiber.App
	userController *controller.UserController
	userService    service.IUserService
}

func NewUserModule(server *fiber.App, userRepository *repository.UserRepository, authGuard *guard.AuthGuard, logger *zap.Logger) *UserModule {
	userService := service.New(userRepository)
	userController := controller.New(authGuard, userService, logger)

	return &UserModule{
		server:         server,
		userController: userController,
		userService:    userService,
	}
}

func (module *UserModule) Register() {
	module.userController.Register(module.server)
}
