package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pemoreau/advent-of-code/go/utils"
	"github.com/pemoreau/advent-of-code/go/utils/set"
)

type machine struct {
	goal    int
	buttons []int
	counter []int
}

func buildGoal(s string) int {
	var res int
	for i := 0; i < len(s); i++ {
		if s[i] == '#' {
			res = 2*res + 1
		} else {
			res = 2 * res
		}
	}

	return res
}

func buildButton(size int, s string) int {
	var res int
	//fmt.Printf("s=%s size=%d\n", s, size)
	var l = strings.Split(s, ",")
	//fmt.Printf("s=%s l=%v\n", s, l)
	for _, v := range l {
		var n, _ = strconv.Atoi(v)
		//fmt.Printf("v=%s n=%d len(l)=%d\n", v, n, len(l))
		res += (1 << (size - 1 - n))
	}
	return res
}
func buildCounter(s string) []int {
	var res []int
	var l = strings.Split(s, ",")
	for _, v := range l {
		var n, _ = strconv.Atoi(v)
		res = append(res, n)
	}
	return res
}

type state struct {
	config  int
	buttons []int
	step    int
}

func remove(buttons []int, index int) []int {
	var res []int
	for i, b := range buttons {
		if i != index {
			res = append(res, b)
		}
	}
	return res
}

func bfs(m machine) (int, bool) {
	var todo []state
	var config = 0
	todo = append(todo, state{config, m.buttons, 0})

	for len(todo) > 0 {
		var s = todo[0]
		todo = todo[1:]

		//fmt.Printf("state = %v\n", s)

		for i, b := range s.buttons {
			var nextConfig = s.config ^ b
			//var nextButtons = append(s.buttons[:i], s.buttons[i+1:]...)
			var nextButtons = remove(s.buttons, i)
			if nextConfig == m.goal {
				return s.step + 1, true
			}
			todo = append(todo, state{nextConfig, nextButtons, s.step + 1})
		}
	}
	return 0, false
}

type stateCounter struct {
	counter []int
	buttons []int
	step    int
}

func bfsCounter(m machine) (int, bool) {
	var visited = set.NewSet[string]()
	var todo []stateCounter
	todo = append(todo, stateCounter{m.counter, m.buttons, 0})

	for len(todo) > 0 {
		var s = todo[0]
		todo = todo[1:]

		var str = fmt.Sprintf("%v", s)
		if visited.Contains(str) {
			continue
		} else {
			visited.Add(str)
		}

		fmt.Printf("state = %v\n", s)

		for _, b := range s.buttons {
			var discard = false
			var nextCounter []int
			for i, c := range s.counter {
				var toAdd = c
				//fmt.Printf("check %d\n", 1<<(len(s.counter)-i-1))
				if b&(1<<(len(s.counter)-i-1)) != 0 {
					toAdd = toAdd - 1
				}
				nextCounter = append(nextCounter, toAdd)
				//fmt.Printf("i=%d b=%d c=%d nextCounter=%v\n", i, b, c, nextCounter)
				if toAdd < 0 {
					discard = true
				}
			}
			//fmt.Printf("button %v, next counter %v\n", b, nextCounter)

			var sum int
			for _, n := range nextCounter {
				sum += n
			}
			if sum == 0 {
				return s.step + 1, true
			}
			if !discard {
				todo = append(todo, stateCounter{nextCounter, s.buttons, s.step + 1})
			}
		}
	}
	return 0, false
}

func parse(input string) []machine {
	input = strings.TrimSuffix(input, "\n")
	var lines = strings.Split(input, "\n")
	var machines []machine
	for _, line := range lines {
		var fields = strings.Fields(line)
		var size = len(fields[0]) - 2
		var goal = buildGoal(fields[0][1 : len(fields[0])-1])
		//fmt.Printf("goal: %v -- %d\n", fields[0], goal)
		var buttons []int
		for j := 1; j < len(fields)-1; j++ {
			var button = buildButton(size, fields[j][1:len(fields[j])-1])
			//fmt.Printf("button: %v -- %d\n", fields[j], button)
			buttons = append(buttons, button)
		}
		var lastField = fields[len(fields)-1]
		var counter = buildCounter(lastField[1 : len(lastField)-1])
		machines = append(machines, machine{goal, buttons, counter})
	}
	return machines
}

func Part1(input string) int {
	var machines = parse(input)
	var res int
	for _, m := range machines {
		var n, _ = bfs(m)
		res += n
	}

	return res
}

func Part2(input string) int {
	var machines = parse(input)
	var res int
	for _, m := range machines {
		fmt.Printf("machines = %v\n", m)
		var n, ok = bfsCounter(m)
		fmt.Printf("n = %v ok = %v\n", n, ok)
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
