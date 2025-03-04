package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheckHandler verifica el estado del servicio
func HealthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Identity Provider service is running",
	})
}
