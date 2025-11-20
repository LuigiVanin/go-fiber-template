package main

import (
	"context"

	"boilerplate/app/middleware/guard"
	"boilerplate/app/modules/auth"
	"boilerplate/app/modules/hash"
	"boilerplate/app/modules/jwt"
	"boilerplate/app/modules/user"
	userRepository "boilerplate/app/modules/user/repository"
	"boilerplate/infra/bootstrap"
	cfg "boilerplate/infra/configuration"

	"go.uber.org/fx"
)

func main() {
	// Create config and logger outside of fx to reuse for fx.WithLogger
	config := cfg.New()
	logger := bootstrap.NewZapLogger(config.Env)

	fx.New(
		// fx.WithLogger(func() fxevent.Logger {
		// 	return &fxevent.ZapLogger{Logger: zapLogger}
		// }),

		// Resources
		fx.Supply(config),
		fx.Supply(logger),
		fx.Provide(bootstrap.NewDatabaseClient),
		fx.Provide(bootstrap.NewHttpServer),

		// Services Modules
		fx.Provide(hash.NewHashBcryptModule),
		fx.Provide(jwt.NewJwtModule),

		// Global Assets
		fx.Provide(userRepository.New),
		fx.Provide(guard.NewAuthGuard),

		// Modules
		fx.Provide(user.NewUserModule),
		fx.Provide(auth.NewAuthModule),

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
