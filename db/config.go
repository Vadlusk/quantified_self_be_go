package db

import (
  "database/sql"
  "fmt"

  m "github.com/vadlusk/quantified_self_be_go/db/migrations"
  s "github.com/vadlusk/quantified_self_be_go/db/seeds"
)

var dbInstance *sql.DB
var initialized bool

const (
  dbport = 5432
  host   = "localhost"
  user   = "vadlusk"
  dbname = "quantified_self_go_dev"
)

func Instance() *sql.DB {
  if !initialized {
    InitDB()
  }
  return dbInstance
}

func InitDB() sql.DB {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
    host, dbport, user, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil { panic(err) }
  err = db.Ping()
  if err != nil { panic(err) }
  initialized = true
  dbInstance = db
  m.Migrate(db)
  s.Seed(db)
  return *db
}
