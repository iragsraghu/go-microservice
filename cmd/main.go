package main

import (
	"fmt"
	"github/iragsraghu/go-microservice/config"
	"github/iragsraghu/go-microservice/internal/api"
	"github/iragsraghu/go-microservice/internal/services"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database connection with the loaded config
	services.InitDB(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	// Set up the Gin router
	router := api.SetupRouter()

	// Start the server on the specified port
	router.Run(fmt.Sprintf(":%s", cfg.ServerPort))
}
