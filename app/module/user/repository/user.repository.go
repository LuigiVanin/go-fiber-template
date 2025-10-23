package user

import (
	"boilerplate/internal/database/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repository *UserRepository) FindAll() ([]models.User, error) {
	users := []models.User{}
	err := repository.db.Find(&users).Error
	return users, err
}

func (repository *UserRepository) FindWhere(model models.User) (*models.User, error) {
	user := models.User{}
	err := repository.db.Where(model).First(&user).Error

	return &user, err
}

func (repository *UserRepository) Create(model models.User) (uint, error) {
	err := repository.db.Create(&model).Error

	if err != nil {
		return 0, err
	}

	return model.ID, nil
}
