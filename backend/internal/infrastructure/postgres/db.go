package db

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/RyotaAbe1014/Pastapal/internal/infrastructure/postgres/db/dbgen"
	"github.com/RyotaAbe1014/Pastapal/pkg/get_env"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	dbName     = get_env.Getenv("POSTGRES_DB", "postgres")
	dbUser     = get_env.Getenv("POSTGRES_USER", "postgres")
	dbPassword = get_env.Getenv("POSTGRES_PASSWORD", "postgres")
	dbHost     = get_env.Getenv("POSTGRES_HOST", "postgres")
	dbPort     = get_env.Getenv("POSTGRES_PORT", "5432")
)

var (
	once     sync.Once
	readOnce sync.Once
	query    *dbgen.Queries
	dbcon    *pgxpool.Pool
)

// contextからQueriesを取得する。contextにQueriesが存在しない場合は、パッケージ変数からQueriesを取得する
func GetQuery(ctx context.Context) *dbgen.Queries {
	txq := getQueriesWithContext(ctx)
	if txq != nil {
		return txq
	}
	return query
}

func SetQuery(q *dbgen.Queries) {
	query = q
}
func GetDB() *pgxpool.Pool {
	return dbcon
}

func SetDB(d *pgxpool.Pool) {
	dbcon = d
}

func NewMainDB() {
	once.Do(func() {
		dbcon, err := createConnection(context.Background())
		if err != nil {
			panic(err)
		}
		q := dbgen.New(dbcon)
		SetQuery(q)
		SetDB(dbcon)
	})
}

func createConnection(ctx context.Context) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	maxRetries := 5
	delay := 2 * time.Second

	for i := 0; i < maxRetries; i++ {
		conn, err := pgxpool.New(ctx, connStr)
		if err == nil {
			return conn, nil
		}

		fmt.Printf("Database not ready, retrying in %v... (%d/%d)\n", delay, i+1, maxRetries)
		time.Sleep(delay)
	}

	return nil, fmt.Errorf("failed to connect to database after %d retries", maxRetries)
}

func getQueries(ctx context.Context) (*dbgen.Queries, *pgxpool.Pool, error) {
	conn, err := createConnection(ctx)

	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	// connを返却するのは、呼び出し元でCloseするため
	return dbgen.New(conn), conn, nil
}

type CtxKey string

const (
	QueriesKey CtxKey = "queries"
)

func WithQueries(ctx context.Context, q *dbgen.Queries) context.Context {
	return context.WithValue(ctx, QueriesKey, q)
}

func getQueriesWithContext(ctx context.Context) *dbgen.Queries {
	queries, ok := ctx.Value(QueriesKey).(*dbgen.Queries)
	if !ok {
		return nil
	}
	return queries
}
