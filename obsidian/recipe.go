package obsidian

import "time"

type Recipe struct {
	Title               string
	Tags                []string
	PreparationDuration time.Duration
	RestingDuration     time.Duration
	CookingDuration     time.Duration
	Serving             string
	Ingredients         []Ingredient
	Steps               []string
}

func (r Recipe) TotalDuration() time.Duration {
	return r.PreparationDuration + r.RestingDuration + r.CookingDuration
}

type Ingredient struct {
	Title      string
	Quantity   string
	Additional string
}
