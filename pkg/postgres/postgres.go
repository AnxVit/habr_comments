package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/AnxVit/ozon_1/internal/logger"

	"github.com/AnxVit/ozon_1/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresClient(cfg *config.Config) (*pgxpool.Pool, error) {
	dbConfig, err := pgxpool.ParseConfig(cfg.PostgreSQL.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse PostgreSQL config: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn, err := pgxpool.NewWithConfig(ctx, dbConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}

	// Проверка подключения
	if err = conn.Ping(ctx); err != nil {
		conn.Close()
		return nil, fmt.Errorf("could not acquire connection from PostgreSQL pool: %w", err)
	}

	logger.Get().Info("Connected to PostgreSQL")
	return conn, nil
}

func ClosePostgresClient(conn *pgxpool.Pool) {
	logger.Get().Info("Closing PostgreSQL connection")
	conn.Close()
}
