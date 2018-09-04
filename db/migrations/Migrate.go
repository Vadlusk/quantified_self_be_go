package migrations

import "database/sql"

func Migrate(db *sql.DB) {
  createMeals(db)
  createFoods(db)
  createMealFoods(db)
}
