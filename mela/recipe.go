package mela

type Recipe struct {
	ID              string   `json:"id"`
	Title           string   `json:"title"`
	Categories      []string `json:"categories"`
	Text            string   `json:"text"`
	Images          []string `json:"images"`
	Yield           string   `json:"yield"`
	PreparationTime string   `json:"prepTime"`
	CookingTime     string   `json:"cookTime"`
	TotalTime       string   `json:"totalTime"`
	Ingredients     string   `json:"ingredients"`
	Instructions    string   `json:"instructions"`
	Notes           string   `json:"notes"`
	Nutrition       string   `json:"nutrition"`
	Link            string   `json:"link"`
}
