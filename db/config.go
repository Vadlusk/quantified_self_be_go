package db

import (
  "database/sql"
  "fmt"

  // "github.com/vadlusk/quantified_self_be_go/db/migrations"
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
  // create tables and seed meals
  createMeals(db)
  // db.Exec(`CREATE TABLE IF NOT EXISTS meals (
  //           id SERIAL PRIMARY KEY NOT NULL,
  //           name TEXT NOT NULL
  // )`)
  db.Exec(`CREATE TABLE IF NOT EXISTS foods (
            id SERIAL PRIMARY KEY NOT NULL,
            name TEXT NOT NULL,
            calories INT NOT NULL
  )`)
  db.Exec(`CREATE TABLE IF NOT EXISTS meal_foods (
            id SERIAL PRIMARY KEY NOT NULL,
            meal_id INT REFERENCES meals ON DELETE CASCADE,
            food_id INT REFERENCES foods ON DELETE CASCADE
  )`)
  db.Exec(`INSERT INTO meals (id, name)
           VALUES (1, 'Breakfast'), (2, 'Snack'), (3, 'Lunch'), (4, 'Dinner')
  `)
  fmt.Println("Successfully seeded!")
  return *db
}

func createMeals(db *sql.DB) {
  db.Exec(`CREATE TABLE IF NOT EXISTS meals (
            id SERIAL PRIMARY KEY NOT NULL,
            name TEXT NOT NULL
  )`)
}
