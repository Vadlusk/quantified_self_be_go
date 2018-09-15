package Meal

import (
  "github.com/vadlusk/quantified_self_be_go/models/Food"
  "github.com/vadlusk/quantified_self_be_go/db"
)

type Meal struct {
  ID    int         `json:"id"`
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
  if err != nil { panic(err) }
  defer rows.Close()
  for rows.Next() {
    if err := rows.Scan(&meal.ID, &meal.Name, &meal.Foods); err != nil {
      panic(err)
    }
    meals = append(meals, meal)
  }
  return meals
}
