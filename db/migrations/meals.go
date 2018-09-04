package migrations

import "database/sql"

func createMeals(db *sql.DB) {
  db.Exec(`CREATE TABLE IF NOT EXISTS meals (
            id SERIAL PRIMARY KEY NOT NULL,
            name TEXT NOT NULL
  )`)
}
