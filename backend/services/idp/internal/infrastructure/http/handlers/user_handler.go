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

type UpdateUserRequest struct {
	Email     string `json:"email" binding:"required,email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	Status    string `json:"status"`
}

func ListUsersHandler(userRepo repositories.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := userRepo.ListUsers(context.Background())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	}
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

func GetUserHandler(userRepo repositories.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		user, err := userRepo.GetUser(context.Background(), id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

func GetUserByEmailHandler(userRepo repositories.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Param("email")
		user, err := userRepo.GetUserByEmail(context.Background(), email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

func UpdateUserHandler(userRepo repositories.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		user, err := userRepo.GetUser(context.Background(), id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var req UpdateUserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validateUpdateUserRequest(user, &req)

		if err := userRepo.UpdateUser(context.Background(), user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

func validateUpdateUserRequest(user *models.User, req *UpdateUserRequest) {
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.FirstName != "" {
		user.FirstName = req.FirstName
	}
	if req.LastName != "" {
		user.LastName = req.LastName
	}
	if req.Password != "" {
		user.Password = req.Password
	}

	if req.Status != "" {
		switch req.Status {
		case "active":
			user.Status = models.UserStatusActive
		case "inactive":
			user.Status = models.UserStatusInactive
		case "blocked":
			user.Status = models.UserStatusBlocked
		}
	}
}

func DeleteUserHandler(userRepo repositories.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		// id := c.Param("id")
		// user, err := userRepo.GetUser(context.Background(), id)
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		// 	return
		// }
		// c.JSON(http.StatusOK, user)
	}
}
