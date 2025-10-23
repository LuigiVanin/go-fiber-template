package auth

import (
	"boilerplate/app/models/dto"
	u "boilerplate/app/module/user/repository"
	"boilerplate/internal/database/models"
	e "boilerplate/internal/errors"
	"fmt"
)

type AuthService struct {
	userRepository u.IUserRepository
}

func New(userRepository u.IUserRepository) *AuthService {
	return &AuthService{
		userRepository: userRepository,
	}
}

func (service *AuthService) SignIn(payload dto.LoginPaylod) error {

	return nil
}

func (service *AuthService) SignUp(payload dto.SignUpPaylod) error {
	fmt.Println("Payload Email: ", payload.Email)

	user, err := service.userRepository.FindWhere(models.User{Email: payload.Email})

	fmt.Println("User: ", user)

	if err == nil {
		return e.ThrowUserAlreadyExists(fmt.Sprintf("%s Email already in use", payload.Email))
	}

	userId, err := service.userRepository.Create(
		models.User{
			Name:     payload.Name,
			Email:    payload.Email,
			Password: payload.Password,
		})

	if err != nil {
		return e.ThrowInternalServerError(err.Error())
	}

	fmt.Println("User created: ", userId)

	return nil
}
