package Food

type Food struct {
  ID       string `json:"id"`
  Name     string `json:"name"`
  Calories int    `json:"calories"`
}

func All() []Food {
  return
}
