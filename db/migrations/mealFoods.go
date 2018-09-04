package migrations

import "database/sql"

func createMealFoods(db *sql.DB) {
  db.Exec(`CREATE TABLE IF NOT EXISTS meal_foods (
            id SERIAL PRIMARY KEY NOT NULL,
            meal_id INT REFERENCES meals ON DELETE CASCADE,
            food_id INT REFERENCES foods ON DELETE CASCADE
  )`)
}
