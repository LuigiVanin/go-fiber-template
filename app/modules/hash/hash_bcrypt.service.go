package hash

import (
	"boilerplate/infra/configuration"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type HashBcryptService struct {
	cfg configuration.Config
}

func New(cfg configuration.Config) *HashBcryptService {
	return &HashBcryptService{
		cfg: cfg,
	}
}

func (service *HashBcryptService) HashPassword(password string) (string, error) {
	salt, err := strconv.Atoi(service.cfg.HashSalt)

	if err != nil {
		salt = 10
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	return string(bytes), err
}

func (service *HashBcryptService) ComparePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
