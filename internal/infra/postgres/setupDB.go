package postgres

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

type Database struct {
  Pool *pgxpool.Pool
}

var (
  NewDatabaseInstance *Database
)

func GetPGXPoolConfig() (*pgxpool.Config) {
  const (
    defaultMaxConnections = int32(10000)
    defaultMinConnections = int32(1)
    defaultMaxConnLifetime = time.Hour
    defaultMaxConnIdleTime = time.Minute * 30
    defaultHealthCheckPeriod = time.Minute
    defaultConnectionTimeout = time.Second * 5

    connectionString = "postgres://root:root@localhost:5432/rinha?"
  )

  config, err := pgxpool.ParseConfig(connectionString)
  if err != nil {
    log.Fatal("error configuring the database: ", err)
  }

  config.MaxConns = defaultMaxConnections
  config.MinConns = defaultMinConnections
  config.MaxConnLifetime = defaultMaxConnLifetime
  config.MaxConnIdleTime = defaultMaxConnIdleTime
  config.HealthCheckPeriod = defaultHealthCheckPeriod
  config.ConnConfig.ConnectTimeout = defaultConnectionTimeout

  return config
}

func NewDatabase() *Database {
  pool, err := pgxpool.NewWithConfig(context.Background(), GetPGXPoolConfig())
  if err != nil {
    log.Fatal("error connecting to the database: ", err)
  }

  NewDatabaseInstance = &Database{
    Pool: pool,
  }

  return &Database{
    Pool: pool,
  }
}
