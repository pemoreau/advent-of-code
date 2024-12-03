package main

import (
	_ "embed"
	"fmt"
	"sort"
	"time"
)

//go:embed input.txt
var inputDay string

type Monkey struct {
	name  int8
	data  []int
	op    transformer
	div   int
	dest1 int8
	dest2 int8
	done  int
}

func (m *Monkey) String() string {
	return fmt.Sprintf("Monkey %d: %v inspected %d", m.name, m.data, m.done)
}

func (m *Monkey) step(monkey []*Monkey) {
	for i := 0; i < len(m.data); i++ {
		n := m.op(m.data[i])
		if n%m.div == 0 {
			monkey[m.dest1].data = append(monkey[m.dest1].data, n)
		} else {
			monkey[m.dest2].data = append(monkey[m.dest2].data, n)
		}
		m.done++
	}
	m.data = m.data[:0] // faster than nil since underlying array is not deallocated
}

type transformer func(int) int

func plus(a int) transformer {
	return func(b int) int {
		return a + b
	}
}

func mult(a int) transformer {
	return func(b int) int {
		return a * b
	}
}

func square() transformer {
	return func(b int) int {
		return b * b
	}
}

func div(a int, t transformer) transformer {
	return func(b int) int {
		return int(float64(t(b)) / float64(a))
	}
}

func mod(a int, t transformer) transformer {
	return func(b int) int {
		return t(b) % a
	}
}

func GenMonkeys(part int) []*Monkey {
	//monkeys := []*Monkey{
	//	{0, []int{79, 98}, div(3, mult(19)), 23, 2, 3, 0},
	//	{1, []int{54, 65, 75, 74}, div(3, plus(6)), 19, 2, 0, 0},
	//	{2, []int{79, 60, 97}, div(3, square()), 13, 1, 3, 0},
	//	{3, []int{74}, div(3, plus(3)), 17, 0, 1, 0},
	//}
	monkeys := []*Monkey{
		{0, []int{83, 62, 93}, mult(17), 2, 1, 6, 0},
		{1, []int{90, 55}, plus(1), 17, 6, 3, 0},
		{2, []int{91, 78, 80, 97, 79, 88}, plus(3), 19, 7, 5, 0},
		{3, []int{64, 80, 83, 89, 59}, plus(5), 3, 7, 2, 0},
		{4, []int{98, 92, 99, 51}, square(), 5, 0, 1, 0},
		{5, []int{68, 57, 95, 85, 98, 75, 98, 75}, plus(2), 13, 4, 0, 0},
		{6, []int{74}, plus(4), 7, 3, 2, 0},
		{7, []int{68, 64, 60, 68, 87, 80, 82}, mult(19), 11, 4, 5, 0},
	}
	for _, m := range monkeys {
		if part == 1 {
			m.op = div(3, m.op)
		} else {
			var p int = 2 * 17 * 19 * 3 * 5 * 13 * 7 * 11
			m.op = mod(p, m.op)
		}
	}
	return monkeys
}
func run(part int, n int) int {
	monkeys := GenMonkeys(part)
	for i := 0; i < n; i++ {
		for _, m := range monkeys {
			m.step(monkeys)
		}
	}
	work := make([]int, len(monkeys))
	for i, m := range monkeys {
		work[i] = m.done
	}
	sort.Ints(work)
	return work[len(work)-1] * work[len(work)-2]
}

func Part1(input string) int {
	return run(1, 20)
}

func Part2(input string) int {
	return run(2, 10000)
}

func main() {
	fmt.Println("--2022 day 11 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
