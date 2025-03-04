package routes

import (
	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/internal/infrastructure/http/handlers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configura todas las rutas del servicio IdP
func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/health", handlers.HealthCheckHandler)
		//api.POST("/login", handlers.LoginHandler)
		//api.POST("/register", handlers.RegisterHandler)
		//api.POST("/refresh", handlers.RefreshTokenHandler)
		//api.GET("/userinfo", handlers.UserInfoHandler)
	}
}
