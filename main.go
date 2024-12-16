package main

import (
	"log"
	coresharedconfig "wishlist-wrangler-api/src/core-shared/core-shared-config"
	coresharedgorm "wishlist-wrangler-api/src/core-shared/core-shared-gorm"
	coresharedlogger "wishlist-wrangler-api/src/core-shared/core-shared-logger"
	coresharedrestservice "wishlist-wrangler-api/src/core-shared/core-shared-rest-service"
	usercontroller "wishlist-wrangler-api/src/user-controller"
	usersharedrepository "wishlist-wrangler-api/src/user-shared/user-shared-repository"
	usersharedservice "wishlist-wrangler-api/src/user-shared/user-shared-service"
)

// @title           User Service
// @version         1.0
// @description     User Service API
func main() {
	appName := "user-service"

	config, err := coresharedconfig.LoadConfig(appName)
	if err != nil {
		log.Printf("Failed to get config: %v", err)
	}

	factory := coresharedlogger.NewPatternLoggerFactory(config.LoggingPattern)
	logger := factory.CreateLogger(appName)

	logger.Info("Connecting to database")
	db, err := coresharedgorm.ConnectWithGorm(config)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	userRepository := usersharedrepository.NewUserRepository(db)
	userService := usersharedservice.NewUserService(userRepository, logger)
	userController := usercontroller.NewUserController(userService)

	userRoutes := userController.RegisterRoutes()

	logger.Info("Starting service")
	coresharedrestservice.StartServiceWithRoutes(&config, appName, userRoutes)
}
