package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Db *pgxpool.Pool

func SetupDatabase() (*pgxpool.Pool, error) {
	dbPool, err := pgxpool.New(context.Background(), "postgres://root:root@db:5432/rinha")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	Db = dbPool

	return Db, nil
}
