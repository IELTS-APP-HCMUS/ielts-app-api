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
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}

	// Create the PostgreSQL DSN (Data Source Name) using the loaded config
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=require",
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
	service := services.NewService(userRepo)

	// Initialize the Gin router and register routes
	router := gin.Default()
	handler := handlers.NewHandler(service)
	handler.RegisterRoutes(router)

	// Start the server
	router.Run(":8080")
}
