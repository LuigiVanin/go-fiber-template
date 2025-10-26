package auth

import (
	"fmt"
	"time"

	"boilerplate/app/models/dto"
	hs "boilerplate/app/modules/hash"
	"boilerplate/app/modules/jwt"
	ur "boilerplate/app/modules/user/repository"
	"boilerplate/infra/database/entity"
	e "boilerplate/infra/errors"
)

type AuthService struct {
	userRepository ur.IUserRepository
	hashService    hs.IHashService
	jwtService     jwt.IJwtService
}

func New(hashService hs.IHashService, jwtService jwt.IJwtService, userRepository ur.IUserRepository) *AuthService {
	return &AuthService{
		userRepository: userRepository,
		hashService:    hashService,
		jwtService:     jwtService,
	}
}

func (service *AuthService) SignIn(payload dto.LoginPaylod) (dto.SignInResponse, error) {

	user, err := service.userRepository.FindWhere(entity.User{Email: payload.Email})

	if err != nil || user == nil || user.ID == 0 {
		return dto.SignInResponse{}, e.ThrowNotFound(fmt.Sprintf("%s Email not found", payload.Email))
	}

	if !service.hashService.ComparePassword(user.Password, payload.Password) {
		return dto.SignInResponse{}, e.ThorwUnauthorizedError("Invalid password")
	}

	token, err := service.jwtService.GenerateToken(dto.JwtPayload{
		UserId: user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Time:   time.Now().Unix(),
	})

	if err != nil {
		return dto.SignInResponse{}, e.ThrowInternalServerError(err.Error())
	}

	return dto.SignInResponse{
		Token: token,
		User: dto.User{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil
}

func (service *AuthService) SignUp(payload dto.SignUpPaylod) (uint, error) {

	_, err := service.userRepository.FindWhere(entity.User{Email: payload.Email})

	if err == nil {
		return 0, e.ThrowUserAlreadyExists(fmt.Sprintf("%s Email already in use", payload.Email))
	}

	hashedPassword, err := service.hashService.HashPassword(payload.Password)

	if err != nil {
		return 0, e.ThrowInternalServerError(err.Error())
	}

	userId, err := service.userRepository.Create(
		entity.User{
			Name:     payload.Name,
			Email:    payload.Email,
			Password: hashedPassword,
		})

	if err != nil || userId == 0 {
		return 0, e.ThrowInternalServerError(err.Error())
	}

	return userId, nil
}
