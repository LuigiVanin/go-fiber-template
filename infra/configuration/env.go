package configuration

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig

	HashSalt  string
	JwtSecret string
}

func New() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found or error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	hashSalt := os.Getenv("HASH_SALT")
	jwtSecret := os.Getenv("JWT_SECRET")

	serverPort := os.Getenv("SERVER_PORT")

	if hashSalt == "" {
		hashSalt = "10"
	}

	if serverPort == "" {
		serverPort = "3000"
	}

	return Config{
		Database: DatabaseConfig{
			Host:     dbHost,
			Port:     dbPort,
			User:     dbUser,
			Password: dbPassword,
			Name:     dbName,
			SSLMode:  dbSSLMode,
		},
		Server: ServerConfig{
			Port: serverPort,
		},

		HashSalt:  hashSalt,
		JwtSecret: jwtSecret,
	}
}

func (env *Config) FormatDatabaseURL() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		env.Database.User,
		env.Database.Password,
		env.Database.Host,
		env.Database.Port,
		env.Database.Name,
		env.Database.SSLMode,
	)
}
