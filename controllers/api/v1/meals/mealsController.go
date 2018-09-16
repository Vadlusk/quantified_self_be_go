package meals

import (
  "encoding/json"
  "net/http"

  "github.com/vadlusk/quantified_self_be_go/models/Meal"
  "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  meals := Meal.All()
  json.NewEncoder(w).Encode(meals)
}

func Show(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  meal   := Meal.Find(params["meal_id"])
  json.NewEncoder(w).Encode(meal)
}
