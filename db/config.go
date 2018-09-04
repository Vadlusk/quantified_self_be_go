package db

import (
  "database/sql"
  "fmt"

  m "github.com/vadlusk/quantified_self_be_go/db/migrations"
  s "github.com/vadlusk/quantified_self_be_go/db/seeds"
)

func InitDB(dbport int, host, user, dbname string) sql.DB {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
    host, dbport, user, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  err = db.Ping()
  if err != nil {
    panic(err)
  }
  m.Migrate(db)
  s.Seed(db)
  return *db
}
