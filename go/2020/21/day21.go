package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var input_day string

func fillIndex(list []string, index int32, mapIndex map[string]int32) int32 {
	for _, a := range list {
		if _, ok := mapIndex[a]; !ok {
			mapIndex[a] = index
			index++
		}
	}
	return index
}

func display(n int, iConstraint []BitVec, aConstraint []BitVec) {
	for i := 0; i < n; i++ {
		//for e := range iConstraint[i] {
		//	if iConstraint[i][e] {
		//		fmt.Print("X")
		//	} else {
		//		fmt.Print(" ")
		//	}
		//}
		//fmt.Print("|")
		//for e := range aConstraint[i] {
		//	if aConstraint[i][e] {
		//		fmt.Print("X")
		//	} else {
		//		fmt.Print(" ")
		//	}
		//}
		//fmt.Println()
		fmt.Printf("%v|%v", iConstraint[i], aConstraint[i])
		if iConstraint[i].Count() == 1 {
			fmt.Printf("(%d)", iConstraint[i].Next(0))
		}
		fmt.Println()
	}
}

type Constraints struct {
	n            int
	ingredrients []BitVec
	allergens    []BitVec
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	allergenIndex := make(map[string]int32)
	ingredientIndex := make(map[string]int32)
	var aIndex int32 = 0
	var iIndex int32 = 0

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

	iConstraint := make([]BitVec, len(lines))
	aConstraint := make([]BitVec, len(lines))
	for line := 0; line < len(lines); line++ {
		iConstraint[line] = New(iIndex)
		aConstraint[line] = New(aIndex)
		for _, i := range ingredients[line] {
			iConstraint[line].Set(ingredientIndex[i])
		}
		for _, a := range allergens[line] {
			aConstraint[line].Set(allergenIndex[a])
		}
	}

	display(len(lines), iConstraint, aConstraint)

	for i := 0; i < len(lines); i++ {
		if aConstraint[i].Count() == 1 {
			// intersect with lines which have the same allergen
			allergen := aConstraint[i].Next(0)
			for j := 0; j < len(lines); j++ {
				if aConstraint[j].Get(allergen) {
					iConstraint[i].And(iConstraint[i], iConstraint[j])
				}
			}
		}
	}
	fmt.Println("after intersection")
	display(len(lines), iConstraint, aConstraint)

	for i := 0; i < len(lines); i++ {
		if aConstraint[i].Count() == 1 && iConstraint[i].Count() == 1 {
			// propagate to other one-allergen lines
			ingredient := iConstraint[i].Next(0)
			for j := 0; j < len(lines); j++ {
				if j != i && aConstraint[j].Count() == 1 && aConstraint[i].Next(0) != aConstraint[j].Next(0) {
					iConstraint[j].Unset(ingredient)
				}
			}
		}
	}
	fmt.Println("after propagation")
	display(len(lines), iConstraint, aConstraint)

	set := New(iIndex)
	for i := 0; i < len(lines); i++ {
		if aConstraint[i].Count() == 1 {
			set.Or(set, iConstraint[i])
		}
	}
	res := 0
	for i := 0; i < iIndex; i++ {
		for j := 0; j < len(lines); j++ {

		}

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
