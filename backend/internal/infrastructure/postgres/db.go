package db

import (
	"context"
	"fmt"

	"github.com/RyotaAbe1014/Pastapal/internal/infrastructure/postgres/db/dbgen"
	"github.com/RyotaAbe1014/Pastapal/pkg/get_env"
	"github.com/jackc/pgx/v5"
)

var (
	dbName     = get_env.Getenv("POSTGRES_DB", "postgres")
	dbUser     = get_env.Getenv("POSTGRES_USER", "postgres")
	dbPassword = get_env.Getenv("POSTGRES_PASSWORD", "password")
	dbHost     = get_env.Getenv("POSTGRES_HOST", "localhost")
	dbPort     = get_env.Getenv("POSTGRES_PORT", "5432")
)

func createConnection(ctx context.Context) (*pgx.Conn, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return conn, nil
}

func getQueries(ctx context.Context) (*dbgen.Queries, *pgx.Conn, error) {
	conn, err := createConnection(ctx)

	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	// connを返却するのは、呼び出し元でCloseするため
	return dbgen.New(conn), conn, nil
}