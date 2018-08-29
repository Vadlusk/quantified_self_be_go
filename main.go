package main

import (
  "database/sql"
  "fmt"
  "os"
  _ "encoding/json"
  "log"
  "net/http"
  "github.com/gorilla/mux"
  _ "github.com/lib/pq"
)

const (
  host   = "localhost"
  dbport = 5432
  user   = "vadlusk"
  dbname = "quantified_self_go_dev"
)

func main() {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
    host, dbport, user, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()
  err = db.Ping()
  if err != nil {
    panic(err)
  }
  fmt.Println("Successfully connected to database!")
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
  port := os.Getenv("PORT")
  if port == "" {
    port = "3000"
  }
  fmt.Println("Listening on port "+port)
  log.Fatal(http.ListenAndServe(":"+port, router))
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
