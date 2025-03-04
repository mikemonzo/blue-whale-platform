package main

import (
	"fmt"
	"log"

	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/internal/infrastructure/db"
	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/internal/infrastructure/http/server"
	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/pkg/config"
)

func main() {
	cfg := config.LoadConfig()

	// Configure DB connection
	dbConn, err := db.NewDBConnection(cfg)
	if err != nil {
		log.Fatalf("Error opening DB connection: %v", err)
	}
	defer dbConn.Close()

	// Initialize repositories
	userRepo := db.NewPostgresUserRepository(dbConn, cfg)

	// Run migrations
	migrationsPath := "internal/infrastructure/db/migrations"
	if err := userRepo.RunMigrations(migrationsPath); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	fmt.Println("Starting Identity Provider service...")

	if err := server.Start(cfg, userRepo); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
