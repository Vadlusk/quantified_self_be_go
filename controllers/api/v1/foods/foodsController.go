package foods

import (
  "encoding/json"
  "net/http"

  "github.com/vadlusk/quantified_self_be_go/models/Food"
)

func Create(w http.ResponseWriter, r *http.Request)  {
  w.Header().Set("Content-Type", "application/json")
}

func Index(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  foods := Food.All()
  json.NewEncoder(w).Encode(foods)
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
