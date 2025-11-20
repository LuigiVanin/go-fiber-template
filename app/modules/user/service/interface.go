package service

import "boilerplate/app/models/dto"

type IUserService interface {
	FindById(userId uint) (dto.User, error)
}
