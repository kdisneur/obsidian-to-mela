package obsidian

type Recipe struct {
	Title               string
	Tags                []string
	PreparationDuration string
	RestingDuration     string
	CookingDuration     string
	Serving             string
	Ingredients         []Ingredient
	Steps               []string
}

type Ingredient struct {
	Title      string
	Quantity   string
	Additional string
}
