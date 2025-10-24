package auth

import (
	"boilerplate/app/models/dto"
	hs "boilerplate/app/modules/hash"
	"boilerplate/app/modules/jwt"
	ur "boilerplate/app/modules/user/repository"
	"boilerplate/infra/database/entity"
	e "boilerplate/infra/errors"
	"fmt"
	"time"
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

func (service *AuthService) SignIn(payload dto.LoginPaylod) error {

	user, err := service.userRepository.FindWhere(entity.User{Email: payload.Email})

	if err != nil || user == nil || user.ID == 0 {
		return e.ThrowNotFound(fmt.Sprintf("%s Email not found", payload.Email))
	}

	if !service.hashService.ComparePassword(user.Password, payload.Password) {
		return e.ThorwUnauthorizedError("Invalid password")
	}

	token, err := service.jwtService.GenerateToken(dto.JwtPayload{
		UserId: user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Time:   time.Now().Unix(),
	})

	if err != nil {
		return e.ThrowInternalServerError(err.Error())
	}

	fmt.Println("Token", token)

	return nil
}

func (service *AuthService) SignUp(payload dto.SignUpPaylod) (uint, error) {
	fmt.Println("Payload Email: ", payload.Email)

	user, err := service.userRepository.FindWhere(entity.User{Email: payload.Email})

	fmt.Println("User: ", user)

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

	fmt.Println("User created: ", userId)

	return userId, nil
}
