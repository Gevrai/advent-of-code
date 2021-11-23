package main

import (
	"fmt"
	"sort"
	"strings"

	. "advent-of-code-2020/utils"
)

type Ingredient string
type Allergen string

type Food struct {
	Ingredients map[Ingredient]bool
	Allergens   map[Allergen]bool
}

func NewFood(input string) (f Food) {
	f = Food{
		Ingredients: map[Ingredient]bool{},
		Allergens:   map[Allergen]bool{},
	}
	parts := strings.Split(input, "(")
	for _, ing := range strings.Split(parts[0], " ") {
		ing := strings.TrimSpace(ing)
		if ing != "" {
			if _, ok := f.Ingredients[Ingredient(ing)]; ok {
				panic(fmt.Sprintf("ingredient %s appears many times", ing))
			}
			f.Ingredients[Ingredient(ing)] = true
		}
	}
	if len(parts) == 1 {
		return f
	}

	allergens := strings.TrimLeft(parts[1], "contains")
	allergens = strings.TrimRight(allergens, ")")
	for _, a := range strings.Split(allergens, ",") {
		a := strings.TrimSpace(a)
		if a != "" {
			if _, ok := f.Allergens[Allergen(a)]; ok {
				panic(fmt.Sprintf("allergen %s appears many times", a))
			}
			f.Allergens[Allergen(a)] = true
		}
	}
	return f
}

func main() {
	DownloadDayInput(2020, 21, false)
	input := SplitNewLine(ReadInputFileRelative())

	var foods []Food
	for _, l := range input {
		foods = append(foods, NewFood(l))
	}

	allergens := getPotentialAllergenIngredients(foods)

	ingredientsAllergenic := map[Ingredient]bool{}
	for _, ings := range allergens {
		for ing := range ings {
			ingredientsAllergenic[ing] = true
		}
	}

	count := 0
	innertIngredients := map[Ingredient]bool{}
	for _, f := range foods {
		for ing := range f.Ingredients {
			if !ingredientsAllergenic[ing] {
				innertIngredients[ing] = true
				count++
			}
		}
	}
	println("Part 1:", count)

	for ing := range innertIngredients {
		removeIngredient(allergens, ing)
	}

	realAllergens := map[Allergen]Ingredient{}
	for {
		allUnique := true
		for all, ings := range allergens {
			if len(ings) == 1 {
				for ing := range ings {
					realAllergens[all] = ing
					removeIngredient(allergens, ing)
				}
				delete(allergens, all)
				continue
			}
			allUnique = false
		}
		if allUnique {
			break
		}
	}

	var allergensSorted []Allergen
	for k := range realAllergens {
		allergensSorted = append(allergensSorted, k)
	}
	sort.SliceStable(allergensSorted, func(i, j int) bool { return allergensSorted[i] < allergensSorted[j] })

	var ingredientsSortedByAllergens []string
	for _, v := range allergensSorted {
		ingredientsSortedByAllergens = append(ingredientsSortedByAllergens, string(realAllergens[v]))
	}
	println("Part 2:", strings.Join(ingredientsSortedByAllergens, ","))
}

func removeIngredient(allergens map[Allergen]map[Ingredient]bool, ingredient Ingredient) {
	for all := range allergens {
		delete(allergens[all], ingredient)
	}
}

func getPotentialAllergenIngredients(foods []Food) map[Allergen]map[Ingredient]bool {
	allergens := make(map[Allergen]map[Ingredient]bool)
	for _, f := range foods {
		for all := range f.Allergens {
			if allergens[all] == nil {
				// We didn't see this allergen yet, add all current ingredients as potential allergen
				allergens[all] = make(map[Ingredient]bool)
				for ing := range f.Ingredients {
					allergens[all][ing] = true
				}
				continue
			}
			// Remove all ingredients NOT applicable to this allergen
			allergens[all] = intersection(allergens[all], f.Ingredients)
		}
	}
	return allergens
}

func intersection(a, b map[Ingredient]bool) map[Ingredient]bool {
	inter := map[Ingredient]bool{}
	for s := range a {
		if b[s] {
			inter[s] = true
		}
	}
	return inter
}
