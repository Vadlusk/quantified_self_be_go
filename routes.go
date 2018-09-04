package main

import (
  "github.com/vadlusk/quantified_self_be_go/controllers/api/v1/foods"
  // "github.com/vadlusk/quantified_self_be_go/controllers/api/v1/meals"
  // mealFoods "github.com/vadlusk/quantified_self_be_go/controllers/api/v1/meals/foods"
  "github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
  r := mux.NewRouter().StrictSlash(true)
  s := r.PathPrefix("/api/v1").Subrouter()
  f := s.PathPrefix("/foods").Subrouter()
  m := s.PathPrefix("/meals").Subrouter()
  f.HandleFunc("/", foods.Create).Methods("POST")
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
