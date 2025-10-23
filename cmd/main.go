package main

import (
	"fmt"

	bootstrap "boilerplate/app"
	authController "boilerplate/app/module/auth/controller"
	authService "boilerplate/app/module/auth/service"
	userRepository "boilerplate/app/module/user/repository"

	"boilerplate/internal/configuration"
	database "boilerplate/internal/database"
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
