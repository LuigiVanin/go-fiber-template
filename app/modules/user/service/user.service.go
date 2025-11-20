package service

import (
	"fmt"

	"boilerplate/app/models/dto"
	ur "boilerplate/app/modules/user/repository"
	"boilerplate/infra/database/entity"
	e "boilerplate/infra/errors"
)

type UserService struct {
	userRepository ur.IUserRepository
}

func NewUserService(userRepository ur.IUserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (service *UserService) FindById(userId uint) (dto.User, error) {
	user, err := service.userRepository.FindWhere(entity.User{ID: userId})

	if err != nil {
		return dto.User{}, e.ThrowNotFound(fmt.Sprintf("User not found: %s", err.Error()))
	}

	return dto.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
