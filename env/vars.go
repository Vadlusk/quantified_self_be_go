package env

import (
  "os"
)

func Port() string {
  port := os.Getenv("PORT")
  if port == "" { port = "3000" }
  return port
}

func Host() string {
  host := os.Getenv("HOST")
  if host == "" { host = "localhost" }
  return host
}

func User() string {
  user := os.Getenv("USER")
  if user == "" { user = "vadlusk" }
  return user
}

func DbPort() string {
  dbPort := os.Getenv("DB_PORT")
  if dbPort == "" { dbPort = "5432" }
  return dbPort
}

func DbName() string {
  dbName := os.Getenv("DB_NAME")
  if dbName == "" { dbName = "quantified_self_go_dev" }
  return dbName
}
