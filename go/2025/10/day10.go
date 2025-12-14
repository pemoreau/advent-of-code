package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/pemoreau/advent-of-code/go/utils"
)

type machine struct {
	goal    []int
	buttons [][]int
	counter []int
}

func createMachine(goal []int, buttons [][]int, counter []int) machine {
	copyGoal := make([]int, len(goal))
	copy(copyGoal, goal)
	var copyButtons = make([][]int, len(buttons))
	copy(copyButtons, buttons)
	var copyCounter = make([]int, len(counter))
	copy(copyCounter, counter)
	return machine{copyGoal, copyButtons, copyCounter}
}

func buildIntList(s string) []int {
	var res []int
	var l = strings.Split(s, ",")
	for _, v := range l {
		var n, _ = strconv.Atoi(v)
		res = append(res, n)
	}
	return res
}

func parse(input string) []machine {
	input = strings.TrimSuffix(input, "\n")
	var lines = strings.Split(input, "\n")
	var machines []machine
	for _, line := range lines {
		var fields = strings.Fields(line)
		var goal []int
		for i := 1; i < len(fields[0])-1; i++ {
			if fields[0][i] == '#' {
				goal = append(goal, 1)
			} else {
				goal = append(goal, 0)
			}
		}
		//fmt.Printf("goal: %v -- %d\n", fields[0], goal)
		var buttons [][]int
		for j := 1; j < len(fields)-1; j++ {
			buttons = append(buttons, buildIntList(fields[j][1:len(fields[j])-1]))
		}
		var lastField = fields[len(fields)-1]
		var counter = buildIntList(lastField[1 : len(lastField)-1])
		machines = append(machines, machine{goal, buttons, counter})
	}
	return machines
}

func remove(buttons [][]int, index int) [][]int {
	var res [][]int
	for i, b := range buttons {
		if i != index {
			res = append(res, b)
		}
	}
	return res
}

type state struct {
	buttons [][]int
	counter []int
	step    int
}

func bfs(m machine) (int, bool) {
	var todo []state
	todo = append(todo, state{m.buttons, make([]int, len(m.goal)), 0})

	//fmt.Printf("machine: %v\n", m)
	for len(todo) > 0 {
		var s = todo[0]
		todo = todo[1:]

		for i, button := range s.buttons {
			var nextCounter = make([]int, len(s.counter))
			copy(nextCounter, s.counter)
			//fmt.Printf("state = %v\n", s)
			//fmt.Printf("apply button[%d] = %v\n", i, button)
			for _, b := range button {
				nextCounter[b] = (nextCounter[b] + 1) % 2
			}
			var nextButtons = remove(s.buttons, i)
			//fmt.Printf("goal = %v nextConfig = %v\n", m.goal, nextConfig)
			var nextState = state{nextButtons, nextCounter, s.step + 1}
			//fmt.Printf("nextState = %v\n", nextState)
			//if slices.Equal(nextConfig, m.goal) {
			if slices.Equal(nextCounter, m.goal) {
				//fmt.Printf("GOAL = %d\n", s.step+1)
				//fmt.Printf("nextState = %v\n", nextState)
				return s.step + 1, true
			}
			todo = append(todo, nextState)
		}
	}
	return 0, false
}

func Part1(input string) int {
	var machines = parse(input)
	var res int
	for _, m := range machines {
		m.counter = m.goal
		var n, _ = bfs(m)
		res += n
	}

	return res
}

func smallerOrEqual(a, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			return false
		}
	}
	return true
}

func equalsModulo2(a, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i]%2 != b[i]%2 {
			return false
		}
	}
	return true
}

type ButtonCombinaison struct {
	counter          []int
	nbPressedButtons int
}

func allCombinaisons(buttons [][]int, m int) []ButtonCombinaison {
	var nbButtons = len(buttons)
	if nbButtons == 0 {
		return []ButtonCombinaison{{counter: make([]int, m), nbPressedButtons: 0}}
	}

	var res []ButtonCombinaison
	for n := 0; n < (1 << nbButtons); n++ {
		var counter = make([]int, m)
		var nbPressedButtons = 0
		for j := 0; j < nbButtons; j++ {
			if (n & (1 << j)) != 0 {
				nbPressedButtons++
				for _, idx := range buttons[j] {
					counter[idx]++
				}
			}
		}
		res = append(res, ButtonCombinaison{counter, nbPressedButtons})
	}
	return res
}

func solve2(counter []int, combinaisons []ButtonCombinaison) (int, bool) {
	var stop = true
	for _, x := range counter {
		stop = stop && x == 0
	}
	if stop {
		return 0, true
	}

	var mini = math.MaxInt
	var found = false

	for _, comb := range combinaisons {
		if !smallerOrEqual(comb.counter, counter) {
			continue
		}
		if !equalsModulo2(comb.counter, counter) {
			continue
		}

		var nextCounter = make([]int, len(counter))
		for i := 0; i < len(counter); i++ {
			nextCounter[i] = (counter[i] - comb.counter[i]) / 2
		}
		rec, ok := solve2(nextCounter, combinaisons)
		if !ok {
			continue
		}

		if n := 2*rec + comb.nbPressedButtons; n < mini {
			mini = n
		}
		found = true
	}

	if !found {
		return 0, false
	}
	return mini, true
}

func Part2(input string) int {
	var machines = parse(input)
	var res int

	for _, m := range machines {
		//fmt.Printf("machines = %v\n", m)
		var combinaisons = allCombinaisons(m.buttons, len(m.counter))
		var n, _ = solve2(m.counter, combinaisons)
		//fmt.Printf("n = %v\n", n)
		res += n
	}
	return res
}

func main() {
	fmt.Println("--2025 day 10 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
