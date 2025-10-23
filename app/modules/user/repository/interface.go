package user

import "boilerplate/infra/database/entity"

type IUserRepository interface {
	FindAll() ([]entity.User, error)
	FindWhere(model entity.User) (*entity.User, error)
	Create(model entity.User) (uint, error)
}
