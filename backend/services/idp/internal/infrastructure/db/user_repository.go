package db

import (
	"context"
	"database/sql"
	"errors"

	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/internal/domain/models"
	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/internal/domain/repositories"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) repositories.UserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) CreateUser(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (id, email, username, first_name, last_name, password, status, created_at, updated_at)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := r.db.ExecContext(ctx, query, user.ID, user.Email, user.Username, user.FirstName, user.LastName, user.Password, user.Status, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresUserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `SELECT id, email, username, first_name, last_name, password, status, created_at, updated_at FROM users WHERE email = $1`

	row := r.db.QueryRowContext(ctx, query, email)
	user := &models.User{}
	if err := row.Scan(&user.ID, &user.Email, &user.Username, &user.FirstName, &user.LastName, &user.Password, &user.Status, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (r *PostgresUserRepository) Update(ctx context.Context, user *models.User) error {
	query := `UPDATE users SET email = $1, username = $2, first_name = $3, last_name = $4, password = $5, status = $6, updated_at = $7 WHERE id = $8`

	_, err := r.db.ExecContext(ctx, query, user.Email, user.Username, user.FirstName, user.LastName, user.Password, user.Status, user.UpdatedAt, user.ID)
	if err != nil {
		return err
	}

	return nil
}
