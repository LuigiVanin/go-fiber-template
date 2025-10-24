package jwt

import (
	"errors"

	"boilerplate/app/models/dto"
	"boilerplate/infra/configuration"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService struct {
	cfg configuration.Config
}

func New(cfg configuration.Config) *JwtService {
	return &JwtService{
		cfg: cfg,
	}
}

func (service *JwtService) GenerateToken(payload dto.JwtPayload) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, dto.JwtPayload{
		UserId: payload.UserId,
		Time:   payload.Time,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(payload.Time+3600, 0)),
			IssuedAt:  jwt.NewNumericDate(time.Unix(payload.Time, 0)),
		},
	})

	tokenString, err := token.SignedString([]byte(service.cfg.JwtSecret))

	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return tokenString, nil
}

func (service *JwtService) VerifyToken(token string) (*dto.JwtPayload, error) {
	return nil, nil
}
