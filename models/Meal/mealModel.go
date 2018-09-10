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
  var meals []Meal
  return meals
}
