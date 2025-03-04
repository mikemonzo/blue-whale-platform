// main.go (Orquestador)
package main

import (
	"fmt"
	"log"

	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/pkg/config"

	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/internal/infrastructure/db"
	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/internal/infrastructure/http/server"
)

func main() {
	cfg := config.LoadConfig()

	// Configure DB connection
	dbConn, err := db.NewDBConnection(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer dbConn.Close()

	userRepo := db.NewPostgresUserRepository(dbConn)

	fmt.Println("Starting Identity Provider service...")

	if err := server.Start(cfg, userRepo); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
