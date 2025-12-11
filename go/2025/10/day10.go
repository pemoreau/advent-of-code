package main

import (
	"cmp"
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/pemoreau/advent-of-code/go/utils"
	"github.com/pemoreau/advent-of-code/go/utils/priorityqueue"
	"github.com/pemoreau/advent-of-code/go/utils/set"
)

type machine struct {
	goal     int
	buttons  []int
	buttons2 [][]int
	counter  []int
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
		var size = len(fields[0]) - 2
		var goal = buildGoal(fields[0][1 : len(fields[0])-1])
		//fmt.Printf("goal: %v -- %d\n", fields[0], goal)
		var buttons []int
		var buttons2 [][]int
		for j := 1; j < len(fields)-1; j++ {
			var button = buildButton(size, fields[j][1:len(fields[j])-1])
			//fmt.Printf("button: %v -- %d\n", fields[j], button)
			buttons = append(buttons, button)
			buttons2 = append(buttons2, buildIntList(fields[j][1:len(fields[j])-1]))
		}
		var lastField = fields[len(fields)-1]
		var counter = buildIntList(lastField[1 : len(lastField)-1])
		machines = append(machines, machine{goal, buttons, buttons2, counter})
	}
	return machines
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

func buildCounter(s string) []int {
	var res []int
	var l = strings.Split(s, ",")
	for i := len(l) - 1; i >= 0; i-- {
		var n, _ = strconv.Atoi(l[i])
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

func pressButton(counter []int, button int) ([]int, bool) {
	var res []int
	for _, c := range counter {
		res = append(res, c)
	}
	var index int
	for button > 0 {
		if button&1 == 1 {
			res[index]--
			if res[index] < 0 {
				return res, false
			}
		}
		button = button >> 1
		index++
	}
	return res, true
}
func pressButton2(counter []int, button []int) ([]int, bool) {
	var res []int
	for _, c := range counter {
		res = append(res, c)
	}
	for _, b := range button {
		res[b]--
		if res[b] < 0 {
			return res, false
		}
	}
	return res, true
}

func goalFunction(s string) bool {
	var sum int
	for _, n := range stringToCounter(s) {
		sum += n
	}
	if sum == 0 {
		return true
	}
	return false
}

func counterToString(counter []int) string {
	var l []string
	for _, n := range counter {
		l = append(l, strconv.Itoa(n))
	}
	//fmt.Printf("counterToString: %v -- %s\n", counter, strings.Join(l, "-"))
	return strings.Join(l, "-")
}

func stringToCounter(str string) []int {
	var l = strings.Split(str, "-")
	var res []int
	for _, n := range l {
		v, _ := strconv.Atoi(n)
		res = append(res, v)
	}
	//fmt.Printf("stringToCounter: %s -- %v\n", str, res)
	return res
}

type stateCounter struct {
	counter  []int
	buttons2 [][]int
	step     int
}

func nextStates(s stateCounter) []stateCounter {
	var occurences = make(map[int]int)
	var smallest = math.MaxInt32
	var maxi int
	var index int
	for _, b := range s.buttons2 {
		for _, n := range b {
			occurences[n] = occurences[n] + 1
		}
	}
	for k, v := range occurences {
		if v > 0 && v < smallest {
			smallest = v
			index = k
		}
		if v > maxi {
			maxi = v
		}
	}
	if smallest > 2 {
		fmt.Println("smallest is too large", smallest, maxi)
	}

	fmt.Printf("smallest occurence: index=%d #%d\n", index, smallest)
	var res []stateCounter
	return res
}

func cost(counter []int) int {
	var res int
	for _, v := range counter {
		res += v * v
	}
	return res
}

func bfsCounter(m machine) (int, bool) {
	var visited = set.NewSet[string]()
	itemCostCmp := func(a, b stateCounter) int {
		return -cmp.Compare(cost(a.counter), cost(b.counter))
	}
	var todo = priorityqueue.New(itemCostCmp)
	todo.Insert(stateCounter{m.counter, m.buttons2, 0})

	for todo.Len() > 0 {
		var s = todo.PopMax()

		var str = fmt.Sprintf("%v", s.counter)
		if visited.Contains(str) {
			continue
		} else {
			visited.Add(str)
		}

		//fmt.Printf("counter = %v cost = %d\n", s.counter, cost(s.counter))

		var nextCounters [][]int
		for _, b := range s.buttons2 {
			var nextCounter, ok = pressButton2(s.counter, b)
			if ok {
				nextCounters = append(nextCounters, nextCounter)
			}
			if cost(nextCounter) == 0 {
				return s.step + 1, true
			}
		}
		//fmt.Printf("nextCounters = %v\n", nextCounters)
		for _, nextCounter := range nextCounters {
			todo.Insert(stateCounter{nextCounter, s.buttons2, s.step + 1})
		}
	}
	return 0, false
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

	//for i := 0; i < len(machines); i++ {
	//	var m = machines[i]
	//	var s = stateCounter{m.counter, m.buttons2, 0}
	//	fmt.Printf("state = %v\n", s)
	//	fmt.Printf("next = %v\n", nextStates(s))
	//}

	return res
}

func Part2Dijkstra(input string) int {
	var machines = parse(input)
	var res int
	for _, m := range machines {
		var costFunction = func(a, b string) int {
			var res int
			var c1 = stringToCounter(a)
			var c2 = stringToCounter(b)
			for i := 0; i < len(c1); i++ {
				//var v = m.counter[i] - utils.Abs(c1[i]-c2[i])
				var v = m.counter[i] - c2[i]
				res += v * v
			}
			return res
		}

		var neighborsFunction = func(s string) []string {
			var res []string
			for _, b := range m.buttons {
				var nextCounter, ok = pressButton(stringToCounter(s), b)
				if ok {
					res = append(res, counterToString(nextCounter))
				}
			}
			return res
		}

		fmt.Printf("machines = %v\n", m)
		var start = counterToString(m.counter)
		var path, cost = utils.Dijkstra[string](start, goalFunction, neighborsFunction, costFunction)
		fmt.Printf("path = %v\n", path)
		fmt.Printf("cost = %v -- %d\n", len(path)-1, cost)
		res += len(path) - 1
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
