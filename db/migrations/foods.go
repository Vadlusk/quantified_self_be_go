package migrations

import "database/sql"

func createFoods(db *sql.DB) {
  db.Exec(`CREATE TABLE IF NOT EXISTS foods (
            id SERIAL PRIMARY KEY NOT NULL,
            name TEXT NOT NULL,
            calories INT NOT NULL
  )`)
}
