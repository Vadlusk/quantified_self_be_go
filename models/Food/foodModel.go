package Food

import "github.com/vadlusk/quantified_self_be_go/db"

type Food struct {
  ID       string `json:"id"`
  Name     string `json:"name"`
  Calories string `json:"calories"`
}

func Create(info Food) Food {
  var created Food
  err := db.Instance().QueryRow(`INSERT INTO foods (name, calories) VALUES ($1, $2) RETURNING *`, info.Name, info.Calories).Scan(&created.ID, &created.Name, &created.Calories)
  if err != nil { panic(err) }
  return created
}

func All() []Food {
  var (
    food Food
    foods []Food
  )
  rows, err := db.Instance().Query(`SELECT * FROM foods`)
  if err != nil { panic(err) }
  defer rows.Close()
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

func Update(id string, info Food) Food {
  var food Food
  err := db.Instance().QueryRow(`UPDATE foods SET name=$1, calories=$2 WHERE id=$3 RETURNING *`, info.Name, info.Calories, id).Scan(&food.ID, &food.Name, &food.Calories)
  if err != nil { panic(err) }
  return food
}

func Destroy(id string) bool {
  err := db.Instance().QueryRow(`DELETE FROM foods WHERE id=$1`, id)
  if err != nil { panic(err) }
  return true
}
