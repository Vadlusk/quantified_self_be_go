package main

import (
  "encoding/json"
  "log"
  "net/http"
  "github.com/gorilla/mux"
)

func main() {
  router := mux.NewRouter()
  router.HandleFunc("/api/v1/foods", CreateFood).Methods("POST")
  router.HandleFunc("/api/v1/foods", GetFoods).Methods("GET")
  router.HandleFunc("/api/v1/foods/{id}", GetFood).Methods("GET")
  router.HandleFunc("/api/v1/foods/{id}", UpdateFood).Methods("PUT")
  router.HandleFunc("/api/v1/foods/{id}", DeleteFood).Methods("DELETE")
  router.HandleFunc("/api/v1/meals", GetMeals).Methods("GET")
  router.HandleFunc("/api/v1/meals/{meal_id}/foods", GetMeal).Methods("GET")
  router.HandleFunc("/api/v1/meals/{meal_id}/foods/{id}", CreateMealFood).Methods("POST")
  router.HandleFunc("/api/v1/meals/{meal_id}/foods/{id}", DeleteMealFood).Methods("DELETE")
  log.Fatal(http.ListenAndServe(":3000", router))
}

func CreateFood(w http.ResponseWriter, r *http.Request) {}
func GetFoods(w http.ResponseWriter, r *http.Request) {}
func GetFood(w http.ResponseWriter, r *http.Request) {}
func UpdateFood(w http.ResponseWriter, r *http.Request) {}
func DeleteFood(w http.ResponseWriter, r *http.Request) {}
func GetMeals(w http.ResponseWriter, r *http.Request) {}
func GetMeal(w http.ResponseWriter, r *http.Request) {}
func CreateMealFood(w http.ResponseWriter, r *http.Request) {}
func DeleteMealFood(w http.ResponseWriter, r *http.Request) {}
