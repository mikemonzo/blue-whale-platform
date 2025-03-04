package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/internal/domain/models"
	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/internal/domain/repositories"
)

type CreateUserRequest struct {
	Email     string `json:"email" binding:"required,email"`
	Username  string `json:"username" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

func CreateUserHandler(userRepo repositories.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateUserRequest
		fmt.Println("CreateUserHandler")
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("CreateUserHandler > user")
		user := &models.User{
			Email:     req.Email,
			Username:  req.Username,
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Password:  req.Password,
			Status:    models.UserStatusInactive,
		}
		fmt.Println("CreateUserHandler > user: ", user)
		if err := userRepo.CreateUser(context.Background(), user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("CreateUserHandler > userRepo.CreateUser")
		// send welcome email (implementation omitted)

		c.JSON(http.StatusCreated, user)
	}
}
