package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils/set"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

const (
	ORE      = 0
	CLAY     = 1
	OBSIDIAN = 2
	GEODE    = 3
	NBROBOT  = 3
)

type State struct {
	time    int8
	product Product
	robot   Robot
}

type Product [4]int16
type Robot [NBROBOT]int8
type Condition [4][3]int16 // needed ORE, CLAY, OBSIDIAN

func (s State) String() string {
	return fmt.Sprintf("time: %d, product: %v, robot: %v\n", s.time, s.product, s.robot)
}

func neighbors(s State, condition Condition, maxTime int8) []State {
	var res []State
	remainingTime := int16(maxTime - s.time)

	if s.product[ORE] >= condition[GEODE][ORE] && s.product[CLAY] >= condition[GEODE][CLAY] && s.product[OBSIDIAN] >= condition[GEODE][OBSIDIAN] {
		// if we can build a GEODE we do it
		newState := s
		newState.time += 1
		for j := 0; j < NBROBOT; j++ {
			newState.product[j] += int16(s.robot[j])
		}
		newState.product[0] -= condition[GEODE][0]
		newState.product[1] -= condition[GEODE][1]
		newState.product[2] -= condition[GEODE][2]
		//newState.robot[GEODE] += 1
		newState.product[GEODE] += remainingTime - 1 // we consider production until the end
		res = append(res, newState)
		return res
	}

	// generate products
	newState := s
	newState.time += 1
	for i := 0; i < NBROBOT; i++ {
		newState.product[i] += int16(s.robot[i])
	}
	res = append(res, newState)

	if remainingTime == 2 {
		// remaining time is not enough to build anything,
		// except a geode robot
		return res
	}
	if remainingTime == 1 {
		// remaining time is not enough to build anything
		return res
	}

	for i := 0; i < NBROBOT; i++ {
		// build new robots
		prod := int16(s.robot[i])
		if prod > condition[0][i] && prod > condition[1][i] && prod > condition[2][i] && prod > condition[3][i] {
			// do not product a new robot if the production is higher than needed
			continue
		}

		//For any resource R that's not geode:
		//  if you already have X robots creating resource R, a current stock of Y for that resource, remainingTime minutes left,
		//   and no robot requires more than Z of resource R to build,
		//   and X * remainingTime+Y >= remainingTime * Z, then you never need to build another robot mining R anymore.
		Z := max(condition[0][i], condition[1][i], condition[2][i], condition[3][i])
		if int16(s.robot[i])*remainingTime+s.product[i] >= remainingTime*Z {
			continue
		}

		if s.product[ORE] >= condition[i][ORE] && s.product[CLAY] >= condition[i][CLAY] && s.product[OBSIDIAN] >= condition[i][OBSIDIAN] {
			newState := s
			newState.time += 1
			for j := 0; j < NBROBOT; j++ {
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
	if a[GEODE] < b[GEODE] {
		return true
	}
	// not sure this is correct for any input, but it works for my input
	//if a[3] == b[3] && a[2] < b[2] {
	//	return true
	//}
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
	t := states[0].time
	robotToProduct := map[Robot]Product{}
	bag := set.Set[State]{}
	// compute max product for a given robot configuration
	for _, s := range states {
		bag.Add(s)
		if s.time != t {
			panic("should not happen")
		}
		oldProduct, ok := robotToProduct[s.robot]
		if !ok || smaller(oldProduct, s.product) {
			robotToProduct[s.robot] = s.product
		}
	}

	var res []State
	//for _, s := range states {
	for s := range bag {
		maxProduct := robotToProduct[s.robot]
		if smaller(s.product, maxProduct) && s.product != maxProduct {
			//fmt.Println("duplicate found", s.product, "smaller than", maxProduct)
		} else {
			res = append(res, s)
		}
	}

	//fmt.Println("removeDuplicates t", t, len(states), "->", len(res))
	return res

}

func removeDuplicates2(states []State) []State {
	if len(states) == 0 {
		return states
	}

	//bag := utils.Set[State]{}
	byRobot := map[Robot]set.Set[State]{}
	for _, s := range states {
		//bag.Add(s)
		_, ok := byRobot[s.robot]
		if !ok {
			byRobot[s.robot] = set.Set[State]{}
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

	var res []State
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
	var res int16 = 0
	for len(todo) > 0 {
		s = todo[0]
		todo = todo[1:]
		next := neighbors(s, condition, maxTime)

		for len(todo) > 0 && todo[0].time == s.time {
			s = todo[0]
			todo = todo[1:]
			next = append(next, neighbors(s, condition, maxTime)...)
		}

		if len(todo) > 0 {
			panic("should not happen")
		}

		next = removeDuplicates2(next)

		for _, n := range next {
			if n.time < maxTime {
				todo = append(todo, n)
			} else if n.product[3] > res {
				res = n.product[3]

			}
		}
	}
	return int(res)
}

func parse(input string) ([]State, []Condition) {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	var states []State
	var conditions []Condition
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
			//robot:   Robot{1, 0, 0, 0},
			robot: Robot{1, 0, 0},
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
		m := solve(states[i], conditions[i], 24)
		res = res + (m * (i + 1))
	}
	return res
}

func Part2(input string) int {
	states, conditions := parse(input)
	res := 1
	for i := 0; i < min(3, len(states)); i++ {
		m := solve(states[i], conditions[i], 32)
		res = res * m
	}
	return res
}

func main() {
	fmt.Println("--2022 day 19 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
