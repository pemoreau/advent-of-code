package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/pemoreau/advent-of-code/go/utils"
)

type Counter []int
type machine struct {
	goal    Counter
	buttons [][]int
	counter Counter
}

type ButtonCombinaison struct {
	counter          Counter
	nbPressedButtons int
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

func (a Counter) smallerOrEqual(b Counter) bool {
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			return false
		}
	}
	return true
}

func (a Counter) equalsModulo2(b Counter) bool {
	for i := 0; i < len(a); i++ {
		if a[i]%2 != b[i]%2 {
			return false
		}
	}
	return true
}
func (a Counter) isZero() bool {
	for i := 0; i < len(a); i++ {
		if a[i] != 0 {
			return false
		}
	}
	return true
}

func allCombinaisons(buttons [][]int, m int) []ButtonCombinaison {
	var nbButtons = len(buttons)
	if nbButtons == 0 {
		return []ButtonCombinaison{{counter: make([]int, m), nbPressedButtons: 0}}
	}

	var res = make([]ButtonCombinaison, 0, 1<<nbButtons)
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
func (goal Counter) solve1(combinaisons []ButtonCombinaison) int {
	var res = math.MaxInt
	for _, comb := range combinaisons {
		if comb.counter.equalsModulo2(goal) {
			if comb.nbPressedButtons < res {
				res = comb.nbPressedButtons
			}
		}
	}
	return res
}

func newCounter(size int) Counter {
	return make([]int, size)
}

func (counter Counter) solve2(combinaisons []ButtonCombinaison) (int, bool) {
	if counter.isZero() {
		return 0, true
	}

	var res = math.MaxInt
	for _, comb := range combinaisons {
		if !comb.counter.smallerOrEqual(counter) {
			continue
		}
		if !comb.counter.equalsModulo2(counter) {
			continue
		}

		var nextCounter = newCounter(len(counter))
		for i := 0; i < len(counter); i++ {
			nextCounter[i] = (counter[i] - comb.counter[i]) / 2
		}
		rec, ok := nextCounter.solve2(combinaisons)
		if !ok {
			continue
		}

		if n := 2*rec + comb.nbPressedButtons; n < res {
			res = n
		}
	}
	if res < math.MaxInt {
		return res, true
	}
	return 0, false
}

func Part1(input string) int {
	var machines = parse(input)
	var res int
	for _, m := range machines {
		var combinaisons = allCombinaisons(m.buttons, len(m.counter))
		var n = m.goal.solve1(combinaisons)
		res += n
	}

	return res
}

func Part2(input string) int {
	var machines = parse(input)
	var res int
	for _, m := range machines {
		var combinaisons = allCombinaisons(m.buttons, len(m.counter))
		var n, _ = m.counter.solve2(combinaisons)
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
