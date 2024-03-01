package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Db *pgxpool.Pool

func SetupDatabase() (*pgxpool.Pool, error) {
  dbConfig, err := pgxpool.ParseConfig("postgres://root:root@127.0.0.1:5432/rinha")
  if err != nil {
    log.Fatalf("Unable to parse database config: %v\n", err)
  }
  dbConfig.MaxConns = 100
  dbConfig.MinConns = 10

	dbPool, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	Db = dbPool

	return Db, nil
}
