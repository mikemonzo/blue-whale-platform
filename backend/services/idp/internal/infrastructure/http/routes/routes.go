package routes

import (
	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/internal/domain/repositories"
	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/internal/infrastructure/http/handlers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configura todas las rutas del servicio IdP
func SetupRoutes(router *gin.Engine, userRepo repositories.UserRepository) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/health", handlers.HealthCheckHandler)
		//v1.POST("/login", handlers.LoginHandler)
		//v1.POST("/register", handlers.RegisterHandler)
		//v1.POST("/refresh", handlers.RefreshTokenHandler)
		//v1.GET("/userinfo", handlers.UserInfoHandler)
	}
	users := v1.Group("/users")
	{
		users.POST("", handlers.CreateUserHandler(userRepo))
		users.GET("", handlers.ListUsersHandler(userRepo))
		users.GET("/:id", handlers.GetUserHandler(userRepo))
		users.GET("/email/:email", handlers.GetUserByEmailHandler(userRepo))
		users.PUT("/:id", handlers.UpdateUserHandler(userRepo))
		users.DELETE("/delete", handlers.DeleteUserHandler(userRepo))
	}
}
