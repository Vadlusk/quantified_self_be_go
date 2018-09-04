package main

import (
  "github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
  r := mux.NewRouter()
  r.HandleFunc("/api/v1/foods/", CreateFood).Methods("POST")
  r.HandleFunc("/api/v1/foods/", GetFoods).Methods("GET")
  r.HandleFunc("/api/v1/foods/{id}", GetFood).Methods("GET")
  r.HandleFunc("/api/v1/foods/{id}", UpdateFood).Methods("PUT")
  r.HandleFunc("/api/v1/foods/{id}", DeleteFood).Methods("DELETE")
  r.HandleFunc("/api/v1/meals/", GetMeals).Methods("GET")
  r.HandleFunc("/api/v1/meals/{meal_id}/foods", GetMeal).Methods("GET")
  r.HandleFunc("/api/v1/meals/{meal_id}/foods/{id}", CreateMealFood).Methods("POST")
  r.HandleFunc("/api/v1/meals/{meal_id}/foods/{id}", DeleteMealFood).Methods("DELETE")
  return r
}
