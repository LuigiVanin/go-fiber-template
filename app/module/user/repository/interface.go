package user

import "boilerplate/internal/database/models"

type IUserRepository interface {
	FindAll() ([]models.User, error)
	FindWhere(model models.User) (*models.User, error)
	Create(model models.User) (uint, error)
}
