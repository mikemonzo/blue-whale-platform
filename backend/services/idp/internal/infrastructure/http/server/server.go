package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/internal/infrastructure/http/routes"
	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/pkg/config"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Start inicia el servidor HTTP
func Start(cfg config.Config) error {
	// Configurar el enrutador HTTP con Gin
	r := gin.Default()
	routes.SetupRoutes(r)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	// Canal para manejar señales del sistema
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		logrus.Infof("Starting server on port %d", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("Server failed: %v", err)
		}
	}()

	// Esperar señal para apagar el servidor
	<-quit
	logrus.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logrus.Errorf("Server forced to shutdown: %v", err)
	}

	logrus.Info("Server exited properly")
	return nil
}
