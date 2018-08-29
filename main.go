package main

import (
  "database/sql"
  "fmt"
  "os"
  "encoding/json"
  "log"
  "net/http"

  "github.com/rs/cors"
  "github.com/gorilla/mux"
  _ "github.com/lib/pq"
)

type Food struct {
  ID       string `json:id`
  Name     string `json:name`
  Calories string `json:calories`
}

var foods []Food

type Meal struct {
  ID    string  `json:id`
  Name  string  `json:name`
}

var meals []Meal

const (
  host   = "localhost"
  dbport = 5432
  user   = "vadlusk"
  dbname = "quantified_self_go_dev"
)

func main() {
  // mock data
  meals = append(meals, Meal{ID: "1", Name: "Breakfast"})
  meals = append(meals, Meal{ID: "2", Name: "Snack"})
  meals = append(meals, Meal{ID: "3", Name: "Lunch"})
  meals = append(meals, Meal{ID: "4", Name: "Dinner"})
  foods = append(foods, Food{ID: "1", Name: "Banana", Calories: "45"})
  foods = append(foods, Food{ID: "2", Name: "Steak", Calories: "800"})
  foods = append(foods, Food{ID: "3", Name: "Apple", Calories: "50"})
  // connect to database
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
  // create tables and seed meals
  db.Exec(`CREATE TABLE IF NOT EXISTS meals (
    id SERIAL PRIMARY KEY NOT NULL,
    name TEXT NOT NULL
  )`)
  db.Exec(`CREATE TABLE IF NOT EXISTS foods (
    id SERIAL PRIMARY KEY NOT NULL,
    name TEXT NOT NULL,
    calories INT NOT NULL
  )`)
  db.Exec(`CREATE TABLE IF NOT EXISTS meal_foods (
    id SERIAL PRIMARY KEY NOT NULL,
    meal_id INT REFERENCES meals ON DELETE CASCADE,
    food_id INT REFERENCES foods ON DELETE CASCADE
  )`)
  fmt.Println("Successfully seeded!")
  // start server
  c := cors.New(cors.Options{
    AllowedOrigins: []string{"*"},
    AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders: []string{"*"},
    Debug: true,
  })
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
  handler := c.Handler(router)
  port := os.Getenv("PORT")
  if port == "" {
    port = "3000"
  }
  fmt.Println("Listening on port "+port)
  log.Fatal(http.ListenAndServe(":"+port, handler))
}

func CreateFood(w http.ResponseWriter, r *http.Request) {}

func GetFoods(w http.ResponseWriter, r *http.Request) {
  json.NewEncoder(w).Encode(foods)
}

func GetFood(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  for _, food := range foods {
    if food.ID == params["id"] {
      json.NewEncoder(w).Encode(food)
      return
    }
  }
}

func UpdateFood(w http.ResponseWriter, r *http.Request) {}
func DeleteFood(w http.ResponseWriter, r *http.Request) {}
func GetMeals(w http.ResponseWriter, r *http.Request) {
  json.NewEncoder(w).Encode(meals)
}
func GetMeal(w http.ResponseWriter, r *http.Request) {}
func CreateMealFood(w http.ResponseWriter, r *http.Request) {}
func DeleteMealFood(w http.ResponseWriter, r *http.Request) {}
