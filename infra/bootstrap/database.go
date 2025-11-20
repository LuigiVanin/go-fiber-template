package bootstrap

import (
	"fmt"

	_ "github.com/lib/pq"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"boilerplate/infra/configuration"
	"boilerplate/infra/database/entity"
)

func CreateConnection(url string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		fmt.Printf("Failed to connect to database: %v", err)
	}

	return db, err
}

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&entity.User{})

	if err != nil {
		fmt.Printf("Failed to migrate database: %v", err)
	}

	return err
}

func NewDatabaseClient(config configuration.Config) *gorm.DB {
	client, err := CreateConnection(config.FormatDatabaseURL())

	if err != nil {
		fmt.Printf("Failed to connect to database: %v", err)
		panic(err)
	}

	db, err := client.DB()

	if err != nil {
		fmt.Printf("Failed to get database: %v", err)
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		fmt.Printf("Failed to ping database: %v", err)
		panic(err)
	}

	return client
}
