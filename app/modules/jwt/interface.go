package jwt

import (
	"boilerplate/app/models/dto"
)

type IJwtService interface {
	GenerateToken(data dto.JwtPayload) (string, error)
	VerifyToken(token string) (*dto.JwtPayload, error)
}
