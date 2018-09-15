package foods

import (
  "encoding/json"
  "net/http"

  "github.com/vadlusk/quantified_self_be_go/models/Food"
  "github.com/gorilla/mux"
)

type FoodInfo struct {
  Food Food.Food `json:"food"`
}

func Create(w http.ResponseWriter, r *http.Request)  {
  w.Header().Set("Content-Type", "application/json")
  var foodInfo FoodInfo
  _ = json.NewDecoder(r.Body).Decode(&foodInfo)
  food := Food.Create(foodInfo.Food)
  json.NewEncoder(w).Encode(food)
}

func Index(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  foods := Food.All()
  json.NewEncoder(w).Encode(foods)
}

func Show(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  food   := Food.Find(params["id"])
  json.NewEncoder(w).Encode(food)
}

func Update(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  var foodInfo FoodInfo
  _ = json.NewDecoder(r.Body).Decode(&foodInfo)
  food := Food.Update(params["id"], foodInfo.Food)
  json.NewEncoder(w).Encode(food)
}

func Destroy(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  Food.Destroy(params["id"])
  w.WriteHeader(http.StatusNoContent)
}
