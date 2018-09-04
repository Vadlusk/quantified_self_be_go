package database

import (
  "database/sql"
  "fmt"
)

func InitializeDB(dbport int, host, user, dbname string) sql.DB {
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
  fmt.Println("Successfully connected to database!")
  return *db
}
