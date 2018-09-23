package Meal

import (
  "fmt"
  "github.com/vadlusk/quantified_self_be_go/models/Food"
  "github.com/vadlusk/quantified_self_be_go/db"
  "github.com/lib/pq"
)

type Meal struct {
  ID    string      `json:"id"`
  Name  string      `json:"name"`
  Foods []Food.Food `json:"foods"`
}

func All() []Meal {
  var (
    meal Meal
    meals []Meal
  )
  rows, err := db.Instance().Query(`SELECT meals.*,
       COALESCE(json_agg(foods.* ORDER BY foods.id)
       FILTER (WHERE foods.id IS NOT NULL), '[]') AS foods
       FROM meals
       LEFT JOIN meal_foods ON meals.id = meal_foods.meal_id
       LEFT JOIN foods ON foods.id = meal_foods.food_id
       GROUP BY meals.id`)
  fmt.Println(rows)
  if err != nil { panic(err) }
  defer rows.Close()
  for rows.Next() {
    if err := rows.Scan(&meal.ID, &meal.Name, pq.Array(&meal.Foods)); err != nil {
      panic(err)
    }
    meals = append(meals, meal)
  }
  return meals
}

func Find(id string) Meal {
  var meal Meal
  query := "SELECT * FROM meals WHERE id=$1"
  err := db.Instance().QueryRow(query, id).Scan(&meal.ID, &meal.Name)
  if err != nil { panic(err) }
  return meal
}
