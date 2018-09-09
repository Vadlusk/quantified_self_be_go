package main

import (
  "database/sql"
  "fmt"
  "log"
  "net/http"
  "os"

  "github.com/vadlusk/quantified_self_be_go/db"
  "github.com/rs/cors"
  _ "github.com/lib/pq"
)

const (
  dbport = 5432
  host   = "localhost"
  user   = "vadlusk"
  dbname = "quantified_self_go_dev"
)

type Env struct {
  db *sql.DB
}

func main() {
  db := db.InitDB(dbport, host, user, dbname)
  defer db.Close()
  env := &Env{db: &db}
  // start server
  router := InitRoutes(env)
  c := cors.New(cors.Options{
    AllowedOrigins: []string{"*"},
    AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders: []string{"*"},
    Debug: true,
  })
  handler := c.Handler(router)
  port := os.Getenv("PORT")
  if port == "" {
    port = "3000"
  }
  fmt.Println("Listening on port "+port)
  log.Fatal(http.ListenAndServe(":"+port, handler))
}
