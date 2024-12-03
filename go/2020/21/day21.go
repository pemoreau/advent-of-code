package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"github.com/pemoreau/advent-of-code/go/utils/set"
	"sort"
	"strings"
	"time"
)

//go:embed sample.txt
var inputTest string

func parse(input string) (allergens map[string]set.Set[string], icount map[string]int) {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	icount = map[string]int{}
	allergens = map[string]set.Set[string]{} // allergen -> ingredients

	for _, line := range lines {
		ingredients := set.Set[string]{}

		parts := strings.Split(line, " (contains ")
		iList := strings.Split(parts[0], " ")
		parts[1] = strings.TrimSuffix(parts[1], ")")
		aList := strings.Split(parts[1], ", ")

		for _, i := range iList {
			ingredients.Add(i)
			icount[i]++
		}

		for _, a := range aList {
			if _, ok := allergens[a]; !ok {
				allergens[a] = set.Set[string]{}
				for i := range ingredients {
					allergens[a].Add(i)
				}
			} else {
				for i := range allergens[a] {
					if !ingredients.Contains(i) {
						allergens[a].Remove(i)
					}
				}
			}
		}
	}
	return
}

func Part1(input string) int {
	allergens, icount := parse(input)
	res := 0
next:
	for i := range icount {
		for _, aset := range allergens {
			if aset.Contains(i) {
				continue next
			}
		}
		res += icount[i]
	}
	return res
}

func Part2(input string) string {
	allergens, _ := parse(input)

	containAllergen := map[string]string{} // ingredient -> allergen

	ingredients := []string{}
	for len(allergens) > 0 {
		for a, aset := range allergens {
			if len(aset) == 1 {
				i := aset.Element()
				for a := range allergens {
					allergens[a].Remove(i)
				}
				delete(allergens, a)
				containAllergen[i] = a
				ingredients = append(ingredients, i)
			}
		}
	}

	sort.Slice(ingredients, func(i, j int) bool { return containAllergen[ingredients[i]] < containAllergen[ingredients[j]] })
	return strings.Join(ingredients, ",")
}

func main() {
	fmt.Println("--2020 day 21 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
