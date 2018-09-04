package main

import (
  "fmt"
  "log"
  "net/http"
  "os"

  "github.com/vadlusk/quantified_self_be_go/db"
  "github.com/rs/cors"
  _ "github.com/lib/pq"
)

type FoodStruct struct {
  Food *Food `json:"food"`
}

type Food struct {
  Name     string `json:"name"`
  Calories string `json:"calories"`
}

const (
  dbport = 5432
  host   = "localhost"
  user   = "vadlusk"
  dbname = "quantified_self_go_dev"
)

func main() {
  db := db.InitDB(dbport, host, user, dbname)
  defer db.Close()
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
  db.Exec(`INSERT INTO meals (id, name)
           VALUES (1, 'Breakfast'), (2, 'Snack'), (3, 'Lunch'), (4, 'Dinner')
  `)
  fmt.Println("Successfully seeded!")
  // start server
  c := cors.New(cors.Options{
    AllowedOrigins: []string{"*"},
    AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders: []string{"*"},
    Debug: true,
  })
  router  := InitRoutes()
  handler := c.Handler(router)
  port    := os.Getenv("PORT")
  if port == "" {
    port = "3000"
  }
  fmt.Println("Listening on port "+port)
  log.Fatal(http.ListenAndServe(":"+port, handler))
}

// func CreateFood(w http.ResponseWriter, r *http.Request) {
//   w.Header().Set("Content-Type", "application/json")
//   var f FoodStruct
//   err := json.NewDecoder(r.Body).Decode(&f)
//   if err != nil {
//     panic(err)
//   }
//   // db.Exec(`INSERT INTO foods (name, calories)
//   //          VALUES (?, ?)
//   //          RETURNING `,
//   //         [f.Food.Name, f.Food.Calories])
// }

// func GetFoods(w http.ResponseWriter, r *http.Request) {
//   w.Header().Set("Content-Type", "application/json")
// }
//
// func GetFood(w http.ResponseWriter, r *http.Request) {
//   w.Header().Set("Content-Type", "application/json")
// }
//
// func UpdateFood(w http.ResponseWriter, r *http.Request) {
//   w.Header().Set("Content-Type", "application/json")
// }
//
// func DeleteFood(w http.ResponseWriter, r *http.Request) {
//   w.Header().Set("Content-Type", "application/json")
//
// }
// func GetMeals(w http.ResponseWriter, r *http.Request) {
//   w.Header().Set("Content-Type", "application/json")
// }
//
// func GetMeal(w http.ResponseWriter, r *http.Request) {
//   w.Header().Set("Content-Type", "application/json")
// }

// func CreateMealFood(w http.ResponseWriter, r *http.Request) {
//   w.Header().Set("Content-Type", "application/json")
// }
//
// func DeleteMealFood(w http.ResponseWriter, r *http.Request) {
//   w.Header().Set("Content-Type", "application/json")
// }
