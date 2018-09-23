package db

import (
  "database/sql"
  "fmt"

  "github.com/vadlusk/quantified_self_be_go/env"
  m "github.com/vadlusk/quantified_self_be_go/db/migrations"
  s "github.com/vadlusk/quantified_self_be_go/db/seeds"
)

var dbInstance *sql.DB
var initialized bool

func Instance() *sql.DB {
  if !initialized {
    InitDB()
  }
  return dbInstance
}

func InitDB() sql.DB {
  psqlInfo := fmt.Sprintf(
    "host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
    env.Host(), env.DbPort(), env.User(), env.DbName(), env.Password())
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
