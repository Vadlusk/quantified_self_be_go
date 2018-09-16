package main

import (
  "github.com/vadlusk/quantified_self_be_go/db"
  _ "github.com/lib/pq"
)

func main() {
  db     := db.InitDB()
  router := initRoutes()
  initServer(router)
  defer db.Close()
}
