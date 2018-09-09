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

const (
  dbport = 5432
  host   = "localhost"
  user   = "vadlusk"
  dbname = "quantified_self_go_dev"
)

func main() {
  db := db.InitDB(dbport, host, user, dbname)
  defer db.Close()
  // start server
  router := InitRoutes()
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
