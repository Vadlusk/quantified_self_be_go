package main

import (
  "github.com/vadlusk/quantified_self_be_go/controllers/api/v1/foods"
  "github.com/vadlusk/quantified_self_be_go/controllers/api/v1/meals"
  mealFoods "github.com/vadlusk/quantified_self_be_go/controllers/api/v1/meals/foods"
  "github.com/gorilla/mux"
)

func initRoutes() *mux.Router {
  r := mux.NewRouter().StrictSlash(true)
  s := r.PathPrefix("/api/v1").Subrouter()
  f := s.PathPrefix("/foods").Subrouter()
  f.HandleFunc("/", foods.Create).Methods("POST")
  f.HandleFunc("/", foods.Index).Methods("GET")
  f.HandleFunc("/{id}", foods.Show).Methods("GET")
  f.HandleFunc("/{id}", foods.Update).Methods("PUT")
  f.HandleFunc("/{id}", foods.Destroy).Methods("DELETE")
  m := s.PathPrefix("/meals").Subrouter()
  m.HandleFunc("/", meals.Index).Methods("GET")
  m.HandleFunc("/{meal_id}/foods", meals.Show).Methods("GET")
  m.HandleFunc("/{meal_id}/foods/{id}", mealFoods.Create).Methods("POST")
  m.HandleFunc("/{meal_id}/foods/{id}", mealFoods.Destroy).Methods("DELETE")
  return r
}
