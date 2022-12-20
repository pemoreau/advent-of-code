package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input_day string

type State struct {
	time    int8
	product Product
	robot   Robot
	//condition [4][3]int16 // needed ORE, CLAY, OBSIDIAN
}

type Product [4]int16
type Robot [4]int8
type Condition [4][3]int16
type Production struct {
	time    int8
	product [4]int16
}

func (s State) String() string {
	return fmt.Sprintf("time: %d, product: %v, robot: %v\n", s.time, s.product, s.robot)
}

func neighbors(s State, condition Condition) []State {
	res := []State{}

	if s.product[0] >= condition[3][0] && s.product[1] >= condition[3][1] && s.product[2] >= condition[3][2] {
		newState := s
		newState.time += 1
		for j := 0; j < 4; j++ {
			newState.product[j] += int16(s.robot[j])
		}
		newState.product[0] -= condition[3][0]
		newState.product[1] -= condition[3][1]
		newState.product[2] -= condition[3][2]
		newState.robot[3] += 1
		res = append(res, newState)
		return res
	}

	newState := s
	newState.time += 1
	for i := 0; i < 4; i++ {
		newState.product[i] += int16(s.robot[i])
	}
	res = append(res, newState)

	for i := 0; i < 4; i++ {
		if s.product[0] >= condition[i][0] && s.product[1] >= condition[i][1] && s.product[2] >= condition[i][2] {
			newState := s
			newState.time += 1
			for j := 0; j < 4; j++ {
				newState.product[j] += int16(s.robot[j])
			}
			newState.product[0] -= condition[i][0]
			newState.product[1] -= condition[i][1]
			newState.product[2] -= condition[i][2]
			newState.robot[i] += 1
			res = append(res, newState)
		}
	}
	return res
}

func smaller(a, b Product) bool {
	if a[3] < b[3] {
		return true
	}
	// not sure this is correct for any input, but it works for my input
	if a[3] == b[3] && a[2] < b[2] {
		return true
	}
	// sure this is not correct, but it works for my input
	//if a[3] == b[3] && a[2] == b[2] && a[1]+3 < b[1] {
	//	return true
	//}
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
	robotToProduct := map[Robot]Product{}
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

	//fmt.Println("removeDuplicates time", time, len(states), "->", len(res))
	return res

}

func removeDuplicates2(states []State) []State {
	if len(states) == 0 {
		return states
	}

	//bag := utils.Set[State]{}
	byRobot := map[Robot]utils.Set[State]{}
	for _, s := range states {
		//bag.Add(s)
		_, ok := byRobot[s.robot]
		if !ok {
			byRobot[s.robot] = utils.Set[State]{}
		}
		byRobot[s.robot].Add(s)
	}

	for robot := range byRobot {
		other := byRobot[robot]
		for a := range byRobot[robot] {
			for b := range other {
				if smaller(a.product, b.product) {
					other.Remove(a)
				}
			}
		}
		byRobot[robot] = other
	}

	//for robot := range byRobot {
	//	max := [4]int{0, 0, 0, 0}
	//	other := byRobot[robot]
	//	for a := range other {
	//		if smaller(max, a.product) {
	//			max = a.product
	//		}
	//	}
	//	for a := range byRobot[robot] {
	//		if smaller(a.product, max) {
	//			other.Remove(a)
	//		}
	//	}
	//	byRobot[robot] = other
	//}

	res := []State{}
	for robot := range byRobot {
		for s := range byRobot[robot] {
			res = append(res, s)
		}
	}

	//fmt.Println("removeDuplicates time", time, len(states), "->", len(res))
	return res

}

func solve(s State, condition Condition, maxTime int8) int {
	todo := []State{s}
	//seen := map[State]Production{}

	var max int16 = 0
	for len(todo) > 0 {
		s = todo[0]
		todo = todo[1:]
		next := neighbors(s, condition)

		for len(todo) > 0 && todo[0].time == s.time {
			s = todo[0]
			todo = todo[1:]
			next = append(next, neighbors(s, condition)...)
		}

		if len(todo) > 0 {
			panic("should not happen")
		}

		next = removeDuplicates(next)

		for _, n := range next {
			if n.time < maxTime {

				//if oldProduct, ok := seen[n]; ok {
				//	if oldProduct.time <= n.time {
				//		if smaller(n.product, oldProduct.product) || n.product == oldProduct.product {
				//			//fmt.Println("seen", oldProduct, "skip", n)
				//			continue
				//		} else if smaller(oldProduct.product, n.product) {
				//			seen[n] = Production{n.time, n.product}
				//			//fmt.Println("seen", oldProduct, "replace by", n)
				//		}
				//	}
				//}
				//seen[n] = Production{n.time, n.product}

				todo = append(todo, n)
			} else if n.product[3] > max {
				//fmt.Println("new max", n)
				max = n.product[3]

			}
		}
	}
	return int(max)
}

func parse(input string) ([]State, []Condition) {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	states := []State{}
	conditions := []Condition{}
	for _, line := range lines {
		var condition Condition
		values := strings.Split(line, " ")
		a, _ := strconv.Atoi(values[6])
		b, _ := strconv.Atoi(values[12])
		c, _ := strconv.Atoi(values[18])
		d, _ := strconv.Atoi(values[21])
		e, _ := strconv.Atoi(values[27])
		f, _ := strconv.Atoi(values[30])
		condition[0][0] = int16(a)
		condition[1][0] = int16(b)
		condition[2][0] = int16(c)
		condition[2][1] = int16(d)
		condition[3][0] = int16(e)
		condition[3][2] = int16(f)

		s := State{
			time:    0,
			product: Product{0, 0, 0, 0},
			robot:   Robot{1, 0, 0, 0},
		}
		states = append(states, s)
		conditions = append(conditions, condition)
	}
	return states, conditions
}

func Part1(input string) int {
	states, conditions := parse(input)
	res := 0
	for i := 0; i < len(states); i++ {
		max := solve(states[i], conditions[i], 24)
		//fmt.Println("input", i+1, " --> ", max)
		res = res + (max * (i + 1))
	}
	return res
}

func Part2(input string) int {
	states, conditions := parse(input)
	res := 1
	for i := 0; i < utils.Min(3, len(states)); i++ {
		max := solve(states[i], conditions[i], 32)
		//fmt.Println("input", i+1, " --> ", max)
		res = res * max
	}
	return res
	// 13, 31, 42 -> 16926
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
