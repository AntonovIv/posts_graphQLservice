package postgres

import (
	"context"
	"embed"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"

	"github.com/AntonovIv/post_graphQlservice/internal/service"
)

//go:embed-migrations/*.sql
var migrations embed.FS

type QueryManager interface {
	DB(ctx context.Context) Querier
	WithTransaction(ctx context.Context, fn service.TxFunc) error
}

type Querier interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	Begin(ctx context.Context) (pgx.Tx, error)
}

type manager struct {
	db *pgxpool.Pool
}

type Config struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Hostname string `yaml:"hostname"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
}

// go get github.com/pressly/goose

func New(ctx context.Context, config Config) (*manager, error) {
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", config.User, config.Password, config.Hostname, config.Port, config.Name)
	conf, err := pgx.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}
	db := stdlib.OpenDB(*conf)
	defer db.Close()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("ping context: %w", err)
	}

	goose.SetBaseFS(migrations)
	err = goose.SetDialect("postgres")
	if err != nil {
		return nil, fmt.Errorf("set dialect: %w", err)
	}

	err = goose.Up(db, "migrations")
	if err != nil {
		return nil, fmt.Errorf("up migrations: %w", err)
	}

	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return nil, err
	}

	return &manager{
		db: pool,
	}, nil
}

const txContextKey = "txContextKey"

func (m *manager) WithTransaction(ctx context.Context, fn service.TxFunc) error {
	return pgx.BeginFunc(ctx, m.DB(ctx), func(tx pgx.Tx) error {
		return fn(context.WithValue(ctx, txContextKey, tx))
	})
}

func (m *manager) DB(ctx context.Context) Querier {
	tx, ok := ctx.Value(txContextKey).(Querier)
	if ok && tx != nil {
		return tx
	}

	return m.db
}
