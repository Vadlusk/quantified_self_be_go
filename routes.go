package main

import (
  "github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
  r := mux.NewRouter()
  s := r.PathPrefix("/api/v1").Subrouter()
  f := s.PathPrefix("/foods").Subrouter()
  m := s.PathPrefix("/meals").Subrouter()
  f.HandleFunc("/", CreateFood).Methods("POST")
  f.HandleFunc("/", GetFoods).Methods("GET")
  f.HandleFunc("/{id}", GetFood).Methods("GET")
  f.HandleFunc("/{id}", UpdateFood).Methods("PUT")
  f.HandleFunc("/{id}", DeleteFood).Methods("DELETE")
  m.HandleFunc("/", GetMeals).Methods("GET")
  m.HandleFunc("/{meal_id}/foods", GetMeal).Methods("GET")
  m.HandleFunc("/{meal_id}/foods/{id}", CreateMealFood).Methods("POST")
  m.HandleFunc("/{meal_id}/foods/{id}", DeleteMealFood).Methods("DELETE")
  return r
}
