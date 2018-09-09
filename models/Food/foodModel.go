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
