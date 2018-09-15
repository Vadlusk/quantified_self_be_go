package MealFood

import (
  "fmt"

  "github.com/vadlusk/quantified_self_be_go/models/Meal"
  "github.com/vadlusk/quantified_self_be_go/models/Food"
  "github.com/vadlusk/quantified_self_be_go/db"
)

func Create(foodId, mealId string) string {
  var id string
  err := db.Instance().QueryRow(`INSERT INTO meal_foods (food_id, meal_id) VALUES ($1, $2) RETURNING id`, foodId, mealId).Scan(&id)
  if err != nil { panic(err) }
  message := fmt.Sprintf("Successfully added %v to %v", foodName(foodId), mealName(mealId))
  return message
}

func foodName(foodId string) string {
  food := Food.Find(foodId)
  name := food.Name
  return name
}

func mealName(mealId string) string {
  meal := Meal.Find(mealId)
  name := meal.Name
  return name
}
