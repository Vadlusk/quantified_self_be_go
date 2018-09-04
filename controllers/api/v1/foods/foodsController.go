package foods

import (
  "net/http"
)

func Create(w http.ResponseWriter, r *http.Request)  {
  w.Header().Set("Content-Type", "application/json")
}

func Index(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
}

func Show(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
}

func Update(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
}

func Destroy(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
}
