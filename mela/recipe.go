package mela

type Recipe struct {
	ID              string   `json:"id"`
	Title           string   `json:"title"`
	Categories      []string `json:"categories"`
	Yield           string   `json:"yield"`
	PreparationTime string   `json:"prepTime"`
	CookingTime     string   `json:"cookTime"`
	TotalTime       string   `json:"totalTime"`
	Ingredients     string   `json:"ingredients"`
	Instructions    string   `json:"instructions"`
}
