package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"strconv"
	"strings"
	"time"
)

//go:embed input_test.txt
var input_day string

const (
	ORE      = 0
	CLAY     = 1
	OBSIDIAN = 2
	GEODE    = 3
)

type State struct {
	time      int
	product   [4]int
	robot     [4]int
	condition [4][3]int // needed ORE, CLAY, OBSIDIAN
}

type Production struct {
	time    int
	product [4]int
}

func (s State) String() string {
	return fmt.Sprintf("time: %d, product: %v, robot: %v\n", s.time, s.product, s.robot)
}

func neighbors(s State) []State {
	res := []State{}

	newState := s
	newState.time += 1
	for i := 0; i < 4; i++ {
		newState.product[i] += s.robot[i]
	}
	res = append(res, newState)

	for i := 0; i < 4; i++ {
		if s.product[0] >= s.condition[i][0] && s.product[1] >= s.condition[i][1] && s.product[2] >= s.condition[i][2] {
			newState := s
			newState.time += 1
			for j := 0; j < 4; j++ {
				newState.product[j] += s.robot[j]
			}
			newState.product[0] -= s.condition[i][0]
			newState.product[1] -= s.condition[i][1]
			newState.product[2] -= s.condition[i][2]
			newState.robot[i] += 1
			res = append(res, newState)
		}
	}
	return res
}

func smaller(a, b [4]int) bool {
	for i := 0; i < 4; i++ {
		if a[i] > b[i] {
			return false
		}
	}
	return a != b
}

func removeDuplicates(states []State) []State {
	if len(states) == 0 {
		return states
	}
	time := states[0].time
	robotToProduct := map[[4]int][4]int{}
	bag := utils.Set[State]{}
	// compute max product for a given robot configuration
	for _, s := range states {
		bag.Add(s)
		if s.time != time {
			panic("should not happen")
		}
		oldProduct, ok := robotToProduct[s.robot]
		if !ok || smaller(oldProduct, s.product) {
			robotToProduct[s.robot] = s.product
		}
	}
	//fmt.Println("robotToProduct", robotToProduct)

	//if len(bag) != len(states) {
	//	fmt.Println(len(states))
	//	fmt.Println(len(bag))
	//	//panic("should not happen")
	//}

	res := []State{}
	//for _, s := range states {
	for s := range bag {
		maxProduct := robotToProduct[s.robot]
		if smaller(s.product, maxProduct) && s.product != maxProduct {
			//fmt.Println("duplicate found", s.product, "smaller than", maxProduct)
		} else {
			res = append(res, s)
		}
	}

	fmt.Println("removeDuplicates time", time, len(states), "->", len(res))
	return res

}

func solve(s State, maxTime int) int {
	todo := []State{s}
	seen := map[State]Production{}

	max := 0
	for len(todo) > 0 {
		s = todo[0]
		todo = todo[1:]
		next := neighbors(s)

		for len(todo) > 0 && todo[0].time == s.time {
			s = todo[0]
			todo = todo[1:]
			next = append(next, neighbors(s)...)
		}

		//fmt.Println("todo", len(todo), "next", len(next))
		if len(todo) > 0 {
			panic("should not happen")
		}

		next = removeDuplicates(next)

		for _, n := range next {
			//fmt.Println(n)
			if n.time < maxTime {
				//if p, ok := seen[n]; ok {
				//	if p.time <= n.time {
				//		if compare(p.product, n.product) >= 0 {
				//			//fmt.Println("seen", p, "skip", n)
				//			continue
				//		} else if compare(p.product, n.product) < 0 {
				//			seen[n] = Production{n.time, n.product}
				//			//fmt.Println("seen", p, "replace by", n)
				//		}
				//	}
				//}
				seen[n] = Production{n.time, n.product}
				todo = append(todo, n)
			} else if n.product[3] > max {
				//fmt.Println("new max", n)
				max = n.product[3]

			}
		}
	}
	return max
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	res := 0
	for i, line := range lines {
		var condition [4][3]int
		values := strings.Split(line, " ")
		condition[0][0], _ = strconv.Atoi(values[6])
		condition[1][0], _ = strconv.Atoi(values[12])
		condition[2][0], _ = strconv.Atoi(values[18])
		condition[2][1], _ = strconv.Atoi(values[21])
		condition[3][0], _ = strconv.Atoi(values[27])
		condition[3][2], _ = strconv.Atoi(values[30])
		s := State{
			time:    0,
			product: [4]int{0, 0, 0, 0},
			robot:   [4]int{1, 0, 0, 0},
			//condition: [4][3]int{{4, 0, 0}, {2, 0, 0}, {3, 14, 0}, {2, 0, 7}},
			condition: condition,
		}
		fmt.Println("part", i+1, s.condition)

		//max := solve(s, 24)
		max := 0
		fmt.Println("max", max, " --> ", max*(i+1))
		res = res + (max * (i + 1))
	}
	return res
}

func Part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	res := 1
	for i, line := range lines {
		if i < 3 {
			var condition [4][3]int
			values := strings.Split(line, " ")
			condition[0][0], _ = strconv.Atoi(values[6])
			condition[1][0], _ = strconv.Atoi(values[12])
			condition[2][0], _ = strconv.Atoi(values[18])
			condition[2][1], _ = strconv.Atoi(values[21])
			condition[3][0], _ = strconv.Atoi(values[27])
			condition[3][2], _ = strconv.Atoi(values[30])
			s := State{
				time:    0,
				product: [4]int{0, 0, 0, 0},
				robot:   [4]int{1, 0, 0, 0},
				//condition: [4][3]int{{4, 0, 0}, {2, 0, 0}, {3, 14, 0}, {2, 0, 7}},
				condition: condition,
			}
			fmt.Println("part", i+1, s.condition)

			max := solve(s, 32)
			fmt.Println("max", max, " --> ", max*(i+1))
			res = res * max
		}
	}
	return res

}

func main() {
	fmt.Println("--2022 day 19 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(input_day))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(input_day))
	fmt.Println(time.Since(start))
}
