package obsidian_test

import (
	"testing"

	"github.com/kdisneur/obsidiantomela/obsidian"
	"github.com/kdisneur/obsidiantomela/testutil/require"
)

func TestParseFileSuccess(t *testing.T) {
	expectedRecipe := obsidian.Recipe{
		Title: "Shakshuka",
		Tags: []string{
			"végétarien-oui",
			"réchauffable-non",
			"difficulté-facile",
			"coût-économique",
		},
		PreparationDuration: "15m0s",
		RestingDuration:     "0s",
		CookingDuration:     "20m0s",
		Serving:             "6 personnes",
		Ingredients: []obsidian.Ingredient{
			{Title: "Oignon", Quantity: "1"},
			{Title: "Poivron", Quantity: "1"},
			{Title: "Ail", Quantity: "4"},
			{Title: "Paprika", Quantity: "2cc"},
			{Title: "Poudre de Cumin", Quantity: "1cc"},
			{Title: "Tomate pelée", Quantity: "800g", Additional: "some information"},
			{Title: "Oeuf", Quantity: "6"},
			{Title: "Coriandre", Quantity: "1cs"},
			{Title: "Persil", Quantity: "1cs"},
			{Title: "Sel"},
			{Title: "Poivre"},
			{Title: "Huile d'olive"},
		},
		Steps: []string{
			"Couper le poivron en dés",
			"Ciseler l'oignon",
			"Ciseler l'ail finement",
			"Ciseler la coriandre finement",
			"Ciseler le persil finement",
			"Chauffer l'huile d'olive dans une sauteuse à feu moyen (4)",
			"Ajouter le poivron, et l'oignon",
			"Laisser cuire 5m, jusqu'à ce que les oignons deviennent translucides",
			"Ajouter l'ail et les épices",
			"Laisser cuire 1m de plus",
			"Ajouter les tomates pelées et le jus dans la sauteuse",
			"Morceler les tomates",
			"Assaisonner avec du sel et du poivre",
			"Pousser la sauce à fine ébullition",
			"Faire des trous dans la sauce pour y déposer un oeuf",
			"Couvrir la sauteuse",
			"Laisser cuire pendant 5m à 8m",
			"Garnir de coriandre et persil juste avant de servir",
		},
	}

	actualRecipe, err := obsidian.ParseFile("testdata/Shakshuka.md")
	require.NoError(t, err, "should be able to parse recipe file")

	require.Equal(t, expectedRecipe.Title, actualRecipe.Title, "should have same title")
	require.DeepEqual(t, expectedRecipe.Tags, actualRecipe.Tags, "should have same tags")
	require.Equal(t, expectedRecipe.PreparationDuration, actualRecipe.PreparationDuration, "should have same preparation duration")
	require.Equal(t, expectedRecipe.RestingDuration, actualRecipe.RestingDuration, "should have same resting duration")
	require.Equal(t, expectedRecipe.CookingDuration, actualRecipe.CookingDuration, "should have same cooking duration")
	require.Equal(t, expectedRecipe.Serving, actualRecipe.Serving, "should have same serving duration")
	require.DeepEqual(t, expectedRecipe.Ingredients, actualRecipe.Ingredients, "should have same ingredients")
	require.DeepEqual(t, expectedRecipe.Steps, actualRecipe.Steps, "should have same steps")
}
