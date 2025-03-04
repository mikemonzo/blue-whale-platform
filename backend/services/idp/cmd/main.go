// main.go (Orquestador)
package main

import (
	"fmt"
	"idp/internal/infrastructure/http/server"
	"idp/pkg/config"
	"log"
)

func main() {
	cfg := config.LoadConfig()

	fmt.Println("Starting Identity Provider service...")

	if err := server.Start(cfg); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
