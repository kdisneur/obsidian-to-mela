package export

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/kdisneur/obsidiantomela/mela"
	"github.com/kdisneur/obsidiantomela/obsidian"
)

type ObsidianToMela struct {
	obsidianRecipes []obsidian.Recipe
	melaIDIndex     mela.IDIndex
}

func LoadObsidianFolder(obsidianFolder string) (*ObsidianToMela, error) {
	melaIDIndex := make(mela.IDIndex)
	var obsidianRecipes []obsidian.Recipe

	err := filepath.WalkDir(obsidianFolder, func(obsidianFilePath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if path.Ext(obsidianFilePath) != ".md" {
			return nil
		}

		recipe, err := obsidian.ParseFile(obsidianFilePath)
		if err != nil {
			return fmt.Errorf("failed to parse Obsidian recipe from %q: %v", obsidianFilePath, err)
		}

		obsidianRecipes = append(obsidianRecipes, recipe)
		melaIDIndex.AddName(recipe.Title)

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &ObsidianToMela{
		obsidianRecipes: obsidianRecipes,
		melaIDIndex:     melaIDIndex,
	}, nil
}

func (o *ObsidianToMela) Export(melaFolder string) error {
	for _, obsidianRecipe := range o.obsidianRecipes {
		identifier, found := o.melaIDIndex.ID(obsidianRecipe.Title)
		if !found {
			return fmt.Errorf("failed to find Mela ID for %q", obsidianRecipe.Title)
		}

		var ingredients []string
		for _, obsidianIngredient := range obsidianRecipe.Ingredients {
			recipeLink, hasRecipeLink := o.melaIDIndex.LinkName(obsidianIngredient.Title)

			ingredient := obsidianIngredient.Title
			if hasRecipeLink {
				ingredient = recipeLink
			}

			if obsidianIngredient.Quantity != "" {
				ingredient = fmt.Sprintf("%s %s", obsidianIngredient.Quantity, ingredient)
			}

			if obsidianIngredient.Additional != "" {
				ingredient = fmt.Sprintf("%s (%s)", ingredient, replaceLinks(o.melaIDIndex, obsidianIngredient.Additional))
			}

			ingredients = append(ingredients, ingredient)
		}

		var instructions []string
		for _, obsidianStep := range obsidianRecipe.Steps {
			instructions = append(instructions, replaceLinks(o.melaIDIndex, obsidianStep))
		}

		melaRecipe := mela.Recipe{
			ID:              identifier,
			Title:           obsidianRecipe.Title,
			Images:          []string{},
			Text:            "",
			Notes:           "",
			Nutrition:       "",
			Link:            "",
			Categories:      obsidianRecipe.Tags,
			Yield:           obsidianRecipe.Serving,
			PreparationTime: obsidianRecipe.PreparationDuration.String(),
			CookingTime:     obsidianRecipe.CookingDuration.String(),
			TotalTime:       obsidianRecipe.TotalDuration().String(),
			Ingredients:     strings.Join(ingredients, "\n"),
			Instructions:    strings.Join(instructions, "\n"),
		}

		encodedRecipe, err := json.Marshal(melaRecipe)
		if err != nil {
			return fmt.Errorf("failed to encode Mela recipe from %q: %v", obsidianRecipe.Title, err)
		}

		recipeFilePath := path.Join(melaFolder, fmt.Sprintf("%s.melarecipe", identifier))
		err = os.WriteFile(recipeFilePath, encodedRecipe, 0644)
		if err != nil {
			return fmt.Errorf("failed to write Mela recipe to %q: %v", recipeFilePath, err)
		}
	}

	return nil
}

func replaceLinks(melaIDIndex mela.IDIndex, s string) string {
	matches := regexp.MustCompile(`\[\[([^\]]+)\]\]`).FindAllStringSubmatch(s, -1)
	for _, name := range matches {
		newLink, hasLink := melaIDIndex.LinkName(name[1])

		replacement := name[1]
		if hasLink {
			replacement = newLink
		}

		s = strings.Replace(s, fmt.Sprintf("[[%s]]", name[1]), replacement, 1)
	}

	return s
}
