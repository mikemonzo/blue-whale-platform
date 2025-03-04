// main.go (Orquestador)
package main

import (
	"fmt"
	"log"

	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/pkg/config"

	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/internal/infrastructure/http/server"
)

func main() {
	cfg := config.LoadConfig()

	fmt.Println("Starting Identity Provider service...")

	if err := server.Start(cfg); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
