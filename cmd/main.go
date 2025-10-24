package main

import (
	"fmt"

	bootstrap "boilerplate/app"
	authController "boilerplate/app/modules/auth/controller"
	authService "boilerplate/app/modules/auth/service"
	hashBcryptService "boilerplate/app/modules/hash"
	jwtService "boilerplate/app/modules/jwt"

	userRepository "boilerplate/app/modules/user/repository"

	"boilerplate/infra/configuration"
	"boilerplate/infra/database"
)

func main() {
	app := bootstrap.New()

	cfg := configuration.New()
	client := database.New(
		cfg.FormatDatabaseURL(),
	)

	err := database.Migrate(client)

	if err != nil {
		fmt.Println("Failed to migrate database: ", err)
		return
	}

	fmt.Println("Database migrated successfully")

	userRepository := userRepository.New(client)

	jwtService := jwtService.New(cfg)
	hashService := hashBcryptService.New(cfg)
	authService := authService.New(hashService, jwtService, userRepository)

	authController := authController.New(authService)

	authController.Register(app)

	app.Listen(fmt.Sprintf(":%s", cfg.Server.Port))
}
