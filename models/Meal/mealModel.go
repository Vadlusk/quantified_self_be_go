package Meal

import (
  "github.com/vadlusk/quantified_self_be_go/models/Food"
  "github.com/vadlusk/quantified_self_be_go/db"
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
  query := `SELECT * FROM meals`
  rows, err := db.Instance().Query(query)
  if err != nil { panic(err) }
  defer rows.Close()
  for rows.Next() {
    if err := rows.Scan(&meal.ID, &meal.Name); err != nil {
      panic(err)
    }
    meals = append(meals, meal)
  }
  for i := 0; i <= 3; i++ {
    meals[i].Foods = mealFoods(i + 1)
  }
  return meals
}

func mealFoods(id int) []Food.Food {
  var (
    food Food.Food
    foods []Food.Food
  )
  query := `SELECT f.* from foods f
            INNER JOIN meal_foods mf ON f.id = mf.food_id
            WHERE mf.meal_id=$1`
  rows, err := db.Instance().Query(query)
  if err != nil { panic(err) }
  defer rows.Close()
  for rows.Next() {
    if err := rows.Scan(&food.ID, &food.Name, &food.Calories); err != nil {
      panic(err)
    }
    foods = append(foods, food)
  }
  return foods
}

func Find(id string) Meal {
  var meal Meal
  query := "SELECT * FROM meals WHERE id=$1"
  err := db.Instance().QueryRow(query, id).Scan(&meal.ID, &meal.Name)
  if err != nil { panic(err) }
  return meal
}
