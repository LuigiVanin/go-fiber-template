package main

import (
	"context"

	"boilerplate/app/middleware/guard"
	"boilerplate/app/modules/auth"
	"boilerplate/app/modules/hash"
	"boilerplate/app/modules/jwt"
	"boilerplate/app/modules/user"
	"boilerplate/infra/bootstrap"
	cfg "boilerplate/infra/configuration"

	"go.uber.org/fx"
)

func main() {
	config := cfg.New()
	logger := bootstrap.NewZapLogger(config.Env)

	fx.New(
		fx.Supply(config),
		fx.Supply(logger),
		fx.Provide(bootstrap.NewDatabaseClient),
		fx.Provide(bootstrap.NewHttpServer),

		// Guards
		fx.Provide(guard.NewAuthGuard),

		// Modules
		hash.Module,
		jwt.Module,
		user.Module,
		auth.Module,

		fx.Invoke(bootstrap.Start),
		fx.Invoke(func(lc fx.Lifecycle) {
			lc.Append(fx.Hook{
				OnStop: func(ctx context.Context) error {
					return logger.Sync()
				},
			})
		}),
	).Run()
}
