package database

import (
	"context"
	"fmt"
	"os"

	"github.com/buemura/event-driven-commerce/packages/tracing"
	"github.com/buemura/event-driven-commerce/svc-product/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	Conn *pgxpool.Pool
)

func Connect() {
	dbConfig, err := pgxpool.ParseConfig(config.DATABASE_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create pool config: %v\n", err)
		os.Exit(1)
	}

	dbConfig.ConnConfig.Tracer = tracing.NewQueryOnlyTracer()

	pool, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	Conn = pool
}
