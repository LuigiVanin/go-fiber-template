package guard

import (
	// "fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"

	"boilerplate/app/common"
	"boilerplate/app/models/dto"

	js "boilerplate/app/modules/jwt"
	ur "boilerplate/app/modules/user/repository"
	"boilerplate/infra/database/entity"
	e "boilerplate/infra/errors"
)

var _ common.IGuard = &AuthGuard{}

type AuthGuard struct {
	jwtService     js.IJwtService
	userRepository ur.IUserRepository
}

func NewAuthGuard(jwtModule *js.JwtModule, userRepository *ur.UserRepository) *AuthGuard {
	return &AuthGuard{
		jwtService:     jwtModule.JwtService,
		userRepository: userRepository,
	}
}

func (guard *AuthGuard) Activate(ctx *fiber.Ctx) error {

	type AuthHeader struct {
		Authorization string `reqHeader:"Authorization"`
	}

	header := AuthHeader{}

	if err := ctx.ReqHeaderParser(&header); err != nil {
		return err
	}

	if header.Authorization == "" {
		return e.ThorwUnauthorizedError("Empty authorization header")
	}

	token := strings.Split(header.Authorization, " ")

	if len(token) != 2 || token[0] != "Bearer" || token[1] == "" {
		return e.ThorwUnauthorizedError("Invalid authorization header")
	}

	jwtToken := token[1]

	payload, err := guard.jwtService.VerifyToken(jwtToken)

	if err != nil {
		return err
	}

	if payload.ExpiresAt.Before(time.Now()) {
		return e.ThrowTokenExpiredError("Token expired")
	}

	user, err := guard.userRepository.FindWhere(entity.User{ID: payload.UserId})

	if err != nil || user.ID != payload.UserId {
		return e.ThrowTokenExpiredError("User not found")
	}

	ctx.Locals("user", dto.User{
		ID:    payload.UserId,
		Name:  payload.Name,
		Email: payload.Email,
	})

	return ctx.Next()
}
