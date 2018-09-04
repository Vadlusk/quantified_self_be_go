package main

import (
  "github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
  r := mux.NewRouter()
  s := r.PathPrefix("/api/v1").Subrouter()
  s.HandleFunc("/foods/", CreateFood).Methods("POST")
  s.HandleFunc("/foods/", GetFoods).Methods("GET")
  s.HandleFunc("/foods/{id}", GetFood).Methods("GET")
  s.HandleFunc("/foods/{id}", UpdateFood).Methods("PUT")
  s.HandleFunc("/foods/{id}", DeleteFood).Methods("DELETE")
  s.HandleFunc("/meals/", GetMeals).Methods("GET")
  s.HandleFunc("/meals/{meal_id}/foods", GetMeal).Methods("GET")
  s.HandleFunc("/meals/{meal_id}/foods/{id}", CreateMealFood).Methods("POST")
  s.HandleFunc("/meals/{meal_id}/foods/{id}", DeleteMealFood).Methods("DELETE")
  return r
}
