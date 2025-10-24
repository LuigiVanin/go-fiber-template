package hash

type IHashService interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword string, password string) bool
}
