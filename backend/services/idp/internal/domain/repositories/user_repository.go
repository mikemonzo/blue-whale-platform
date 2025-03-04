package repositories

import (
	"context"

	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/internal/domain/models"
)

type UserRepository interface {
	RunMigrations(migrationsPath string) error
	CreateUser(ctx context.Context, user *models.User) error
	// GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	// GetUserByID(ctx context.Context, id string) (*models.User, error)
	// Update(ctx context.Context, user *models.User) error
}
