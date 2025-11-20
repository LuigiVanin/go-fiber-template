package repository

import (
	"boilerplate/infra/database/entity"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repository *UserRepository) FindAll() ([]entity.User, error) {
	users := []entity.User{}
	err := repository.db.Find(&users).Error
	return users, err
}

func (repository *UserRepository) FindWhere(model entity.User) (*entity.User, error) {
	user := entity.User{}
	err := repository.db.Where(model).First(&user).Error

	return &user, err
}

func (repository *UserRepository) Create(model entity.User) (uint, error) {
	err := repository.db.Create(&model).Error

	if err != nil {
		return 0, err
	}

	return model.ID, nil
}
