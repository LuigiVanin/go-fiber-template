package jwt

import "boilerplate/infra/configuration"

type JwtModule struct {
	JwtService IJwtService
}

func NewJwtModule(config configuration.Config) *JwtModule {
	return &JwtModule{
		JwtService: New(config),
	}
}
