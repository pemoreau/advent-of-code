package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input_test.txt
var input_day string

type Allergen string
type Ingredient string

func fillIndex(list []string, index int, mapIndex map[string]int) int {
	for _, a := range list {
		if _, ok := mapIndex[a]; !ok {
			mapIndex[a] = index
			index++
		}
	}
	return index
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	allergenIndex := make(map[string]int)
	ingredientIndex := make(map[string]int)
	aIndex := 0
	iIndex := 0

	ingredients := make([][]string, 0, len(lines))
	allergens := make([][]string, 0, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, " (contains ")
		iList := strings.Split(parts[0], " ")
		parts[1] = strings.TrimSuffix(parts[1], ")")
		aList := strings.Split(parts[1], ", ")

		ingredients = append(ingredients, iList)
		allergens = append(allergens, aList)

		iIndex = fillIndex(iList, iIndex, ingredientIndex)
		aIndex = fillIndex(aList, aIndex, allergenIndex)
	}
	//fmt.Println("ingredients", ingredients)
	//fmt.Println("allergens", allergens)

	iConstraint := make([][]bool, len(lines))
	aConstraint := make([][]bool, len(lines))
	for line := 0; line < len(lines); line++ {
		iConstraint[line] = make([]bool, iIndex)
		aConstraint[line] = make([]bool, aIndex)
		for _, i := range ingredients[line] {
			iConstraint[line][ingredientIndex[i]] = true
		}
		for _, a := range allergens[line] {
			aConstraint[line][allergenIndex[a]] = true
		}
	}

	fmt.Println(ingredientIndex)
	fmt.Println(allergenIndex)

	for i := 0; i < len(lines); i++ {
		fmt.Println(iConstraint[i], aConstraint[i])
	}

	return 0
}

func Part2(input string) int {
	// input = strings.TrimSuffix(input, "\n")
	// lines := strings.Split(input, "\n")
	return 0

}

func main() {
	fmt.Println("--2020 day 21 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(input_day))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(input_day))
	fmt.Println(time.Since(start))
}
