package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"github.com/pemoreau/advent-of-code/go/utils/set"
	"slices"
	"strconv"
	"strings"
	"time"
)

//go:embed sample.txt
var inputTest string

func parse(input string) (map[string]set.Set[string], []map[string]set.Set[string], [][]string) {
	var parts = strings.Split(input, "\n\n")

	var orders = make(map[string]set.Set[string])
	var updates []map[string]set.Set[string]
	var lines [][]string
	for _, line := range strings.Split(parts[0], "\n") {
		var first = line[:2]
		var second = line[3:5]
		var order, ok = orders[first]
		if !ok {
			order = set.NewSet[string]()
			orders[first] = order
		}
		order.Add(second)
	}

	for _, line := range strings.Split(parts[1], "\n") {
		var elements = strings.Split(line, ",")
		var update = make(map[string]set.Set[string])
		lines = append(lines, elements)
		var toAdd = set.NewSet[string]()
		for i := len(elements) - 1; i >= 0; i-- {
			update[elements[i]] = toAdd
			toAdd = toAdd.Union(set.NewSet[string]())
			toAdd.Add(elements[i])
		}
		updates = append(updates, update)
	}
	return orders, updates, lines
}

func checkUpdate(orders map[string]set.Set[string], update map[string]set.Set[string], line []string) bool {
	for _, element := range line {
		nextUpdates := orders[element]
		if nextOrders, ok := update[element]; ok {
			if !nextUpdates.ContainsSet(nextOrders) {
				return false
			}
		}
	}
	return true
}

func Part1(input string) int {
	var orders, updates, lines = parse(input)
	var res int
	for i, update := range updates {
		if checkUpdate(orders, update, lines[i]) {
			var middle, _ = strconv.Atoi(lines[i][len(lines[i])/2])
			res += middle
		}
	}
	return res
}

func Part2(input string) int {
	var orders, updates, lines = parse(input)
	var cmp = func(a, b string) int {
		if orders[a].Contains(b) {
			return -1
		}
		if orders[b].Contains(a) {
			return 1
		}
		return 0
	}
	var res int
	for i, line := range lines {
		if checkUpdate(orders, updates[i], lines[i]) {
			continue
		}

		slices.SortFunc(line, cmp)
		var middle, _ = strconv.Atoi(lines[i][len(lines[i])/2])
		res += middle
	}
	return res
}

func main() {
	fmt.Println("--2024 day 05 solution--")
	var inputDay = utils.Input()
	//var inputDay = inputTest
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
