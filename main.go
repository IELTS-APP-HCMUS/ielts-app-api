package main

import (
	"fmt"
	"ielts-app-api/config"
	"ielts-app-api/internal/handlers"
	"ielts-app-api/internal/repositories"
	"ielts-app-api/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load configuration from environment variables
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}

	// Create the PostgreSQL DSN (Data Source Name) using the loaded config
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
		cfg.Postgres.Host,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Name,
		cfg.Postgres.Port,
	)

	// Initialize the database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	// Initialize the repository and service
	userRepo := repositories.NewUserRepository(db)
	targetRepo := repositories.NewTargetRepository(db)
	service := services.NewService(userRepo, targetRepo)
	// Initialize the Gin router and register routes
	router := gin.Default()
	handler := handlers.NewHandler(service)
	handler.RegisterRoutes(router)

	// Start the server
	router.Run("0.0.0.0:" + fmt.Sprintf("%d", 8080))
}
