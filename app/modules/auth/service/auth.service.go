package auth

import (
	"boilerplate/app/models/dto"
	ur "boilerplate/app/modules/user/repository"
	"boilerplate/infra/database/entity"
	e "boilerplate/infra/errors"
	"fmt"
)

type AuthService struct {
	userRepository ur.IUserRepository
}

func New(userRepository ur.IUserRepository) *AuthService {
	return &AuthService{
		userRepository: userRepository,
	}
}

func (service *AuthService) SignIn(payload dto.LoginPaylod) error {

	return nil
}

func (service *AuthService) SignUp(payload dto.SignUpPaylod) error {
	fmt.Println("Payload Email: ", payload.Email)

	user, err := service.userRepository.FindWhere(entity.User{Email: payload.Email})

	fmt.Println("User: ", user)

	if err == nil {
		return e.ThrowUserAlreadyExists(fmt.Sprintf("%s Email already in use", payload.Email))
	}

	userId, err := service.userRepository.Create(
		entity.User{
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
