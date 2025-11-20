package hash

import (
	"go.uber.org/fx"
)

var Module = fx.Module("hash",
	fx.Provide(
		fx.Annotate(
			NewHashBcryptService,
			fx.As(new(IHashService)),
		),
	),
)
