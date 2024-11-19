package db

import (
	"context"
	"fmt"
	"sync"

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

var (
	once     sync.Once
	readOnce sync.Once
	query    *dbgen.Queries
	dbcon    *pgx.Conn
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
func GetDB() *pgx.Conn {
	return dbcon
}
func SetDB(d *pgx.Conn) {
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
