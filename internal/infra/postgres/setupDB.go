package postgres

import (
  "database/sql"
  "sync"

  _ "github.com/lib/pq"
)

var (
  once sync.Once
  db *sql.DB
)

func GetConnection() (*sql.DB, error) {
  var err error

  once.Do(func() {
    db, err = sql.Open("postgres", "host=localhost port=5432 user=root dbname=rinha password=root sslmode=disable")
    if err != nil {
      panic(err)
    }

    err = db.Ping()
    if err != nil {
      panic(err)
    }
  })

  return db, err
}

