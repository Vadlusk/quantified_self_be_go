package main

import (
  "fmt"
  "log"
  "net/http"

  "github.com/vadlusk/quantified_self_be_go/env"
  "github.com/rs/cors"
  "github.com/gorilla/mux"
)

func initServer(router *mux.Router) {
  cors := cors.New(cors.Options{
    AllowedOrigins: []string{"*"},
    AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders: []string{"*"},
    Debug: true,
  })
  handler := cors.Handler(router)
  port    := env.Port()
  fmt.Println("Listening on port "+port)
  log.Fatal(http.ListenAndServe(":"+port, handler))
}
