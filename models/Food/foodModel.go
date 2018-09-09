package Food

import "github.com/vadlusk/quantified_self_be_go/db"

type Food struct {
  ID       string `json:"id"`
  Name     string `json:"name"`
  Calories int    `json:"calories"`
}

func All() []Food {
  rows, err := db.Instance().Query(`SELECT * FROM foods`)
  if err != nil { panic(err) }
  defer rows.Close()
  var (
    food Food
    foods []Food
  )
  for rows.Next() {
    if err := rows.Scan(&food.ID, &food.Name, &food.Calories); err != nil {
      panic(err)
    }
    foods = append(foods, food)
  }
  return foods
}

func Find(id string) Food {
  var food Food
  err := db.Instance().QueryRow(`SELECT * FROM foods WHERE id=$1`, id).Scan(&food.ID, &food.Name, &food.Calories)
  if err != nil { panic(err) }
  return food
}
