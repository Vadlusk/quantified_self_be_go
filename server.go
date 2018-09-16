package main

import (
  "fmt"
  "log"
  "net/http"
  "os"

  "github.com/rs/cors"
  "github.com/gorilla/mux"
)

func initServer(router *mux.Router) {
  c := cors.New(cors.Options{
    AllowedOrigins: []string{"*"},
    AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders: []string{"*"},
    Debug: true,
  })
  handler := c.Handler(router)
  port    := os.Getenv("PORT")
  if port == "" {
    port = "3000"
  }
  fmt.Println("Listening on port "+port)
  log.Fatal(http.ListenAndServe(":"+port, handler))
}
