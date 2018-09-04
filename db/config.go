package db

import (
  "database/sql"
  "fmt"

  m "github.com/vadlusk/quantified_self_be_go/db/migrations"
  // "github.com/vadlusk/quantified_self_be_go/db/seeds"
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
  fmt.Println("Successfully connected to database!")
  m.Migrate(db)
  db.Exec(`INSERT INTO meals (id, name)
           VALUES (1, 'Breakfast'), (2, 'Snack'), (3, 'Lunch'), (4, 'Dinner')
  `)
  fmt.Println("Successfully seeded!")
  return *db
}
