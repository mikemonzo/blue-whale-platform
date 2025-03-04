package repositories

import (
	"context"

	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/internal/domain/models"
)

type UserRepository interface {
	RunMigrations(migrationsPath string) error
	ListUsers(ctx context.Context) ([]*models.User, error)
	CreateUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, id string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, user *models.User) error
}
