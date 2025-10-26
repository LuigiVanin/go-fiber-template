package main

import (
	"fmt"

	bootstrap "boilerplate/app"
	"boilerplate/app/common"
	authService "boilerplate/app/modules/auth/service"
	hashBcryptService "boilerplate/app/modules/hash"
	jwtService "boilerplate/app/modules/jwt"
	logger "boilerplate/infra"

	authGuard "boilerplate/app/middleware/guard"

	userRepository "boilerplate/app/modules/user/repository"
	userService "boilerplate/app/modules/user/service"

	authController "boilerplate/app/modules/auth/controller"
	userController "boilerplate/app/modules/user/controller"

	"boilerplate/infra/configuration"
	"boilerplate/infra/database"

	"go.uber.org/zap"
)

func main() {
	cfg := configuration.New()
	common.Logger = logger.New(cfg.Env)

	app := bootstrap.New()

	client := database.New(
		cfg.FormatDatabaseURL(),
	)

	err := database.Migrate(client)

	if err != nil {
		common.Logger.Error("Failed to migrate database: ", zap.Error(err))
		return
	}

	common.Logger.Info("Database migrated successfully")

	userRepository := userRepository.New(client)

	jwtService := jwtService.New(cfg)
	hashService := hashBcryptService.New(cfg)
	authService := authService.New(hashService, jwtService, userRepository)
	userService := userService.New(userRepository)

	authGuard := authGuard.New(jwtService, userRepository)

	authController := authController.New(authService)
	userController := userController.New(authGuard, userService)

	authController.Register(app)
	userController.Register(app)

	app.Listen(fmt.Sprintf(":%s", cfg.Server.Port))
}
