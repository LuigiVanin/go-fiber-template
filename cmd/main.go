package main

import (
	"fmt"

	bootstrap "boilerplate/app"
	authController "boilerplate/app/modules/auth/controller"
	authService "boilerplate/app/modules/auth/service"
	userRepository "boilerplate/app/modules/user/repository"

	"boilerplate/infra/configuration"
	"boilerplate/infra/database"
)

func main() {
	app := bootstrap.New()

	env := configuration.NewEnvironment()
	client := database.New(env.FormatDatabaseURL())

	err := database.Migrate(client)

	if err != nil {
		fmt.Println("Failed to migrate database: ", err)
		return
	}

	fmt.Println("Database migrated successfully")

	userRepository := userRepository.New(client)

	authService := authService.New(userRepository)

	authController := authController.New(authService)

	authController.Register(app)

	app.Listen(":3000")
}
