package coresharedgorm

import (
	"fmt"
	"log"
	coresharedconfig "wishlist-wrangler-api/src/core-shared/core-shared-config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectWithGorm(config coresharedconfig.ConfigDto) (*gorm.DB, error) {
	host := config.Host
	port := config.Port
	database := config.Database
	user := config.Username
	password := config.Password
	sslmode := config.Sslmode

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host,
		user,
		password,
		database,
		port,
		sslmode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return db, nil
}
