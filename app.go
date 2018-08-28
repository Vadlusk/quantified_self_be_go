package main

import (
  "encoding/json"
  "log"
  "net/http"
  "github.com/gorilla/mux"
)

func main() {
  router := mux.NewRouter()
  log.Fatal(http.ListenAndServe(":3000", router))
}
