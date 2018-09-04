package seeds

import "database/sql"

func Seed(db *sql.DB) {
  seedMeals(db)
}
