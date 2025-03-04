package repository

import (
	"context"

	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/internal/domain/model"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	// GetUserByID(ctx context.Context, id string) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
}
