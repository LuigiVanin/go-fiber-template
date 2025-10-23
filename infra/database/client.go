package database

import (
	"boilerplate/infra/database/entity"
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

func New(url string) *gorm.DB {
	client, err := CreateConnection(url)

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
