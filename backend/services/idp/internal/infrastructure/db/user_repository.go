package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/internal/domain/models"
	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/internal/domain/repositories"
	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/pkg/config"
)

type PostgresUserRepository struct {
	db     *sql.DB
	config config.Config
}

func NewPostgresUserRepository(db *sql.DB, cfg config.Config) repositories.UserRepository {
	return &PostgresUserRepository{db: db, config: cfg}
}

func (r *PostgresUserRepository) RunMigrations(migrationsPath string) error {
	m, err := migrate.New(
		fmt.Sprintf("file://%s", migrationsPath),
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			r.config.DbUser, r.config.DbPassword, r.config.DbHost, r.config.DbPort, r.config.DbName))
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to run migrations: %v", err)
	}

	log.Println("Migrations ran successfully")
	return nil
}

func (r *PostgresUserRepository) ListUsers(ctx context.Context) ([]*models.User, error) {
	query := `SELECT id, email, username, first_name, last_name, password, status, created_at, updated_at FROM users`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*models.User, 0)
	for rows.Next() {
		user := &models.User{}
		if err := rows.Scan(&user.ID, &user.Email, &user.Username, &user.FirstName, &user.LastName, &user.Password, &user.Status, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *PostgresUserRepository) CreateUser(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (id, email, username, first_name, last_name, password, status, created_at, updated_at)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	if user.ID == "" {
		user.ID = uuid.New().String()
	}

	fmt.Println("CreateUserDB > user: ", user)

	_, err := r.db.ExecContext(ctx, query, user.ID, user.Email, user.Username, user.FirstName, user.LastName, user.Password, user.Status, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return err
	}
	fmt.Println("CreateUserDB > user created")
	return nil
}

func (r *PostgresUserRepository) GetUser(ctx context.Context, id string) (*models.User, error) {
	query := `SELECT id, email, username, first_name, last_name, password, status, created_at, updated_at FROM users WHERE id = $1`

	row := r.db.QueryRowContext(ctx, query, id)
	user := &models.User{}
	if err := row.Scan(&user.ID, &user.Email, &user.Username, &user.FirstName, &user.LastName, &user.Password, &user.Status, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
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

func (r *PostgresUserRepository) UpdateUser(ctx context.Context, user *models.User) error {
	query := `UPDATE users SET email = $1, username = $2, first_name = $3, last_name = $4, password = $5, status = $6, updated_at = $7 WHERE id = $8`

	_, err := r.db.ExecContext(ctx, query, user.Email, user.Username, user.FirstName, user.LastName, user.Password, user.Status, user.UpdatedAt, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresUserRepository) DeleteUser(ctx context.Context, user *models.User) error {
	query := `DELETE FROM users WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, user.ID)
	if err != nil {
		return err
	}

	return nil
}
