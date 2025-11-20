package jwt

import (
	"errors"
	"strconv"

	"boilerplate/app/models/dto"
	"boilerplate/infra/configuration"
	e "boilerplate/infra/errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService struct {
	cfg configuration.Config
}

func NewJwtService(cfg configuration.Config) *JwtService {
	return &JwtService{
		cfg: cfg,
	}
}

func (service *JwtService) GenerateToken(payload dto.JwtPayload) (string, error) {

	expTime, err := strconv.Atoi(service.cfg.JwtExpTime)

	if err != nil {
		expTime = 3600
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, dto.JwtPayload{
		UserId: payload.UserId,
		Name:   payload.Name,
		Email:  payload.Email,
		Time:   payload.Time,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(
				payload.Time+int64(expTime),
				0,
			)),
			IssuedAt: jwt.NewNumericDate(time.Unix(payload.Time, 0)),
		},
	})

	tokenString, err := token.SignedString([]byte(service.cfg.JwtSecret))

	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return tokenString, nil
}

func (service *JwtService) VerifyToken(token string) (*dto.JwtPayload, error) {

	parsedToken, err := jwt.ParseWithClaims(token, &dto.JwtPayload{}, func(token *jwt.Token) (interface{}, error) {

		// Validate if the signed method is not NONE
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, e.ThorwUnauthorizedError("Invalid token: missing signing method verification")
		}

		return []byte(service.cfg.JwtSecret), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, e.ThrowTokenExpiredError("Token expired")
		}

		return nil, e.ThorwUnauthorizedError("Invalid token: " + err.Error())
	}

	if claims, ok := parsedToken.Claims.(*dto.JwtPayload); ok && parsedToken.Valid {
		return claims, nil
	}

	return nil, e.ThorwUnauthorizedError("Invalid token")
}
