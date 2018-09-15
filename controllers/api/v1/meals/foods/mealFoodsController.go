package foods

import (
  "encoding/json"
  "net/http"

  "github.com/vadlusk/quantified_self_be_go/models/MealFood"
  "github.com/gorilla/mux"
)

func Create(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params   := mux.Vars(r)
  mealFood := MealFood.Create(params["id"], params["meal_id"])
  json.NewEncoder(w).Encode(mealFood)
}

func Destroy(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
}
