package main

import (
	_ "embed"
	"fmt"
	"sort"
	"time"
)

//go:embed input.txt
var input_day string

type Monkey struct {
	name  int
	data  []int
	op    byte
	arg   int
	div   int
	dest1 int
	dest2 int
	done  int
}

func (m *Monkey) String() string {
	return fmt.Sprintf("Monkey %d: %v inspected %d", m.name, m.data, m.done)
}

func (m *Monkey) run1(monkey []*Monkey) {
	for i := 0; i < len(m.data); i++ {
		var new int
		if m.op == '*' {
			new = int(float64(m.arg*m.data[i]) / float64(3))
		} else if m.op == '+' {
			new = int((float64(m.arg+m.data[i]) / float64(3)))
		} else {
			new = int((float64(m.data[i]*m.data[i]) / float64(3)))
		}
		if new%m.div == 0 {
			monkey[m.dest1].data = append(monkey[m.dest1].data, new)
		} else {
			monkey[m.dest2].data = append(monkey[m.dest2].data, new)
		}
		m.done++
	}
	m.data = nil
}

//func (m *Monkey) run0(monkey []Monkey) {
//	const mult = 19
//	const div = 23
//	const dest1 = 2
//	const dest2 = 3
//	for i := 0; i < len(m.data); i++ {
//		new := int(math.Round(float64(mult*m.data[i]) / float64(3)))
//		if new%div == 0 {
//			monkey[dest1].data = append(monkey[dest1].data, new)
//		} else {
//			monkey[dest2].data = append(monkey[dest2].data, new)
//		}
//	}
//}

func Part1(input string) int {
	// input = strings.TrimSuffix(input, "\n")
	// lines := strings.Split(input, "\n")
	//monkeys := []*Monkey{
	//	&Monkey{name: 0, data: []int{79, 98}, op: '*', arg: 19, div: 23, dest1: 2, dest2: 3},
	//	&Monkey{name: 1, data: []int{54, 65, 75, 74}, op: '+', arg: 6, div: 19, dest1: 2, dest2: 0},
	//	&Monkey{name: 2, data: []int{79, 60, 97}, op: '2', div: 13, dest1: 1, dest2: 3},
	//	&Monkey{name: 3, data: []int{74}, op: '+', arg: 3, div: 17, dest1: 0, dest2: 1},
	//}
	monkeys := []*Monkey{
		&Monkey{name: 0, data: []int{83, 62, 93}, op: '*', arg: 17, div: 2, dest1: 1, dest2: 6},
		&Monkey{name: 1, data: []int{90, 55}, op: '+', arg: 1, div: 17, dest1: 6, dest2: 3},
		&Monkey{name: 1, data: []int{91, 78, 80, 97, 79, 88}, op: '+', arg: 3, div: 19, dest1: 7, dest2: 5},
		&Monkey{name: 1, data: []int{64, 80, 83, 89, 59}, op: '+', arg: 5, div: 3, dest1: 7, dest2: 2},
		&Monkey{name: 1, data: []int{98, 92, 99, 51}, op: '2', div: 5, dest1: 0, dest2: 1},
		&Monkey{name: 1, data: []int{68, 57, 95, 85, 98, 75, 98, 75}, op: '+', arg: 2, div: 13, dest1: 4, dest2: 0},
		&Monkey{name: 1, data: []int{74}, op: '+', arg: 4, div: 7, dest1: 3, dest2: 2},
		&Monkey{name: 1, data: []int{68, 64, 60, 68, 87, 80, 82}, op: '*', arg: 19, div: 11, dest1: 4, dest2: 5},
	}

	for i := 0; i < 20; i++ {
		for _, m := range monkeys {
			m.run1(monkeys)
		}
	}
	work := make([]int, len(monkeys))
	for i, m := range monkeys {
		work[i] = m.done
	}
	sort.Ints(work)
	return work[len(work)-1] * work[len(work)-2]
}

func (m *Monkey) run2(monkey []*Monkey) {
	for i := 0; i < len(m.data); i++ {
		var new int
		if m.op == '*' {
			new = m.arg * m.data[i]
		} else if m.op == '+' {
			new = m.arg + m.data[i]
		} else {
			new = m.data[i] * m.data[i]
		}
		if new%m.div == 0 {
			monkey[m.dest1].data = append(monkey[m.dest1].data, new)
		} else {
			monkey[m.dest2].data = append(monkey[m.dest2].data, new)
		}
		m.done++
	}
	m.data = nil
}
func Part2(input string) int {
	// input = strings.TrimSuffix(input, "\n")
	// lines := strings.Split(input, "\n")
	monkeys := []*Monkey{
		&Monkey{name: 0, data: []int{79, 98}, op: '*', arg: 19, div: 23, dest1: 2, dest2: 3},
		&Monkey{name: 1, data: []int{54, 65, 75, 74}, op: '+', arg: 6, div: 19, dest1: 2, dest2: 0},
		&Monkey{name: 2, data: []int{79, 60, 97}, op: '2', div: 13, dest1: 1, dest2: 3},
		&Monkey{name: 3, data: []int{74}, op: '+', arg: 3, div: 17, dest1: 0, dest2: 1},
	}
	//monkeys := []*Monkey{
	//	&Monkey{name: 0, data: []int{83, 62, 93}, op: '*', arg: 17, div: 2, dest1: 1, dest2: 6},
	//	&Monkey{name: 1, data: []int{90, 55}, op: '+', arg: 1, div: 17, dest1: 6, dest2: 3},
	//	&Monkey{name: 1, data: []int{91, 78, 80, 97, 79, 88}, op: '+', arg: 3, div: 19, dest1: 7, dest2: 5},
	//	&Monkey{name: 1, data: []int{64, 80, 83, 89, 59}, op: '+', arg: 5, div: 3, dest1: 7, dest2: 2},
	//	&Monkey{name: 1, data: []int{98, 92, 99, 51}, op: '2', div: 5, dest1: 0, dest2: 1},
	//	&Monkey{name: 1, data: []int{68, 57, 95, 85, 98, 75, 98, 75}, op: '+', arg: 2, div: 13, dest1: 4, dest2: 0},
	//	&Monkey{name: 1, data: []int{74}, op: '+', arg: 4, div: 7, dest1: 3, dest2: 2},
	//	&Monkey{name: 1, data: []int{68, 64, 60, 68, 87, 80, 82}, op: '*', arg: 19, div: 11, dest1: 4, dest2: 5},
	//}

	for i := 0; i < 10000; i++ {
		for _, m := range monkeys {
			m.run2(monkeys)
		}
		fmt.Println("after round", i+1)
		for _, m := range monkeys {
			fmt.Println(m.String())
		}
	}
	work := make([]int, len(monkeys))
	for i, m := range monkeys {
		work[i] = m.done
	}
	sort.Ints(work)
	return work[len(work)-1] * work[len(work)-2]
}

func main() {
	fmt.Println("--2022 day 11 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(input_day))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(input_day))
	fmt.Println(time.Since(start))
}
