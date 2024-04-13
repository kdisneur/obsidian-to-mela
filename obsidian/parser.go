package obsidian

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
)

func ParseFile(filepath string) (Recipe, error) {
	title := strings.TrimSuffix(path.Base(filepath), ".md")

	content, err := os.ReadFile(filepath)
	if err != nil {
		return Recipe{}, fmt.Errorf("failed to read file %q: %v", filepath, err)
	}

	fileContent := strings.Split(string(content), "\n")

	tags, fileContent, err := parseTags(fileContent)
	if err != nil {
		return Recipe{}, fmt.Errorf("failed to parse tags of file %q: %v", filepath, err)
	}

	preparationDuration, fileContent, err := parseDuration(fileContent, "Préparation")
	if err != nil {
		return Recipe{}, fmt.Errorf("failed to parse preparation duration of file %q: %v", filepath, err)
	}

	restingDuration, fileContent, err := parseDuration(fileContent, "Repos")
	if err != nil {
		return Recipe{}, fmt.Errorf("failed to parse resting duration of file %q: %v", filepath, err)
	}

	cookingDuration, fileContent, err := parseDuration(fileContent, "Cuisson")
	if err != nil {
		return Recipe{}, fmt.Errorf("failed to parse cooking duration of file %q: %v", filepath, err)
	}

	serving, fileContent, err := parseServing(fileContent)
	if err != nil {
		return Recipe{}, fmt.Errorf("failed to parse serving of file %q: %v", filepath, err)
	}

	ingredients, fileContent, err := parseIngredients(fileContent)
	if err != nil {
		return Recipe{}, fmt.Errorf("failed to parse ingredients of file %q: %v", filepath, err)
	}

	steps, err := parseSteps(fileContent)
	if err != nil {
		return Recipe{}, fmt.Errorf("failed to parse steps of file %q: %v", filepath, err)
	}

	return Recipe{
		Title:               title,
		Tags:                tags,
		PreparationDuration: preparationDuration,
		RestingDuration:     restingDuration,
		CookingDuration:     cookingDuration,
		Serving:             serving,
		Ingredients:         ingredients,
		Steps:               steps,
	}, nil
}

func parseTags(lines []string) ([]string, []string, error) {
	lines = skipUntil(lines, regexp.MustCompile("^#[a-zA-Z0-9]"))

	var tags []string

	line := lines[0]
	for _, tag := range strings.Split(line, " ") {
		tags = append(tags, strings.TrimSpace(strings.TrimPrefix(tag, "#")))
	}

	return tags, lines[1:], nil
}

func parseDuration(lines []string, durationKind string) (string, []string, error) {
	lines = skipUntil(lines, regexp.MustCompile(durationKind))

	line := lines[0]

	splittedLine := strings.Split(line, ":")
	if len(splittedLine) != 2 {
		return "", nil, fmt.Errorf("expected a key value for the duration but got %q", line)
	}

	return strings.TrimSpace(splittedLine[1]), lines[1:], nil
}

func parseServing(lines []string) (string, []string, error) {
	lines = skipUntil(lines, regexp.MustCompile("Ingrédients "))

	line := lines[0]

	start := strings.Index(line, "(")
	end := strings.Index(line, ")")

	return line[start+1 : end], lines[1:], nil
}

func parseIngredients(lines []string) ([]Ingredient, []string, error) {
	lines = skipUntil(lines, regexp.MustCompile("^- "))

	var ingredients []Ingredient
	for i := range lines {
		if !strings.HasPrefix(lines[i], "-") {
			break
		}

		var title, quantity, additionalInformation string
		matches := regexp.
			MustCompile(`^-\s+\[\[(.*)\]\](\s*\|\s*([^|]+)(\s*\|\s*(.*))?)?`).
			FindStringSubmatch(lines[i])

		if matches == nil {
			break
		}

		title = matches[1]
		quantity = matches[3]
		additionalInformation = matches[5]

		ingredients = append(ingredients, Ingredient{
			Title:      strings.TrimSpace(title),
			Quantity:   strings.TrimSpace(quantity),
			Additional: strings.TrimSpace(additionalInformation),
		})
	}

	return ingredients, lines[len(ingredients):], nil
}

func parseSteps(lines []string) ([]string, error) {
	lines = skipUntil(lines, regexp.MustCompile("Étapes"))
	lines = lines[1:]

	var steps []string

	for i := range lines {
		trimmedLine := strings.TrimSpace(lines[i])
		if trimmedLine == "" {
			continue
		}

		steps = append(steps, trimmedLine)
	}

	return steps, nil
}

func skipUntil(lines []string, r *regexp.Regexp) []string {
	for i := range lines {
		if r.MatchString(lines[i]) {
			return lines[i:]
		}
	}

	return []string{""}
}
