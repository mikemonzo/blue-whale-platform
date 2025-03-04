package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/pkg/config"
)

// NewDBConnection configura y devuelve una nueva conexión a la base de datos
func NewDBConnection(cfg config.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=%s",
		cfg.DbUser, cfg.DbPassword, cfg.DbName, cfg.DbHost, cfg.DbPort, "disable")
	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening DB connection: %v", err)
	}

	// Verificar la conexión
	if err := dbConn.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging DB: %v", err)
	}

	return dbConn, nil
}
