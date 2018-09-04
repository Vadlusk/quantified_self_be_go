package seeds

import "database/sql"

func seedMeals(db *sql.DB) {
  db.Exec(`INSERT INTO meals (id, name)
           VALUES (1, 'Breakfast'), (2, 'Snack'), (3, 'Lunch'), (4, 'Dinner')
  `)
}
