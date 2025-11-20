package auth

import (
	controller "boilerplate/app/modules/auth/controller"
	"boilerplate/app/modules/auth/service"
	"boilerplate/app/modules/hash"
	"boilerplate/app/modules/jwt"
	userRepository "boilerplate/app/modules/user/repository"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type AuthModule struct {
	server         *fiber.App
	authController *controller.AuthController
}

func NewAuthModule(
	server *fiber.App,
	hashModule *hash.HashBcryptModule,
	jwtModule *jwt.JwtModule,
	userRepository *userRepository.UserRepository,
	logger *zap.Logger,
) *AuthModule {

	return &AuthModule{
		server: server,
		authController: controller.New(
			service.New(
				hashModule.HashService,
				jwtModule.JwtService,
				userRepository,
			),
			logger,
		),
	}
}

func (module *AuthModule) Register() {
	module.authController.Register(module.server)
}
