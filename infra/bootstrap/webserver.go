package bootstrap

import (
	"boilerplate/app/middleware"
	"boilerplate/infra/configuration"

	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewHttpServer(logger *zap.Logger, config configuration.Config) *fiber.App {

	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.NewErrorHandler(logger),
		AppName:      config.AppName,
		ServerHeader: config.AppName,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))
	app.Use(middleware.Json)
	app.Use(middleware.NewRequestLogger(logger))

	return app
}

func Start(
	lifecycle fx.Lifecycle,
	server *fiber.App,
	config configuration.Config,
	client *gorm.DB,
	log *zap.Logger,
) {

	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {

				log.Info("Connecting to database...")
				log.Info("Migrating database...")

				err := Migrate(client)

				if err != nil {
					log.Error("Failed to migrate database", zap.Error(err))
					return err
				}

				log.Info("Database migrated successfully!")

				log.Info("Starting server...",
					zap.String("port", config.Server.Port),
					zap.String("env", config.Env),
				)
				go func() {
					err := server.Listen(fmt.Sprintf(":%s", config.Server.Port))
					if err != nil {
						log.Error("Error starting server", zap.Error(err))
					}
				}()

				return nil
			},
			OnStop: func(ctx context.Context) error {
				log.Info("Shutting down server...")
				return server.Shutdown()
			},
		},
	)
}
