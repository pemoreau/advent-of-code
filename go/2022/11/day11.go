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
			new = int(float64(m.arg+m.data[i]) / float64(3))
		} else {
			new = int(float64(m.data[i]*m.data[i]) / float64(3))
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

func Part1(input string) int {
	// input = strings.TrimSuffix(input, "\n")
	// lines := strings.Split(input, "\n")
	//monkeys := []*Monkey{
	//	{name: 0, data: []int{79, 98}, op: '*', arg: 19, div: 23, dest1: 2, dest2: 3},
	//	{name: 1, data: []int{54, 65, 75, 74}, op: '+', arg: 6, div: 19, dest1: 2, dest2: 0},
	//	{name: 2, data: []int{79, 60, 97}, op: '2', div: 13, dest1: 1, dest2: 3},
	//	{name: 3, data: []int{74}, op: '+', arg: 3, div: 17, dest1: 0, dest2: 1},
	//}
	monkeys := []*Monkey{
		{name: 0, data: []int{83, 62, 93}, op: '*', arg: 17, div: 2, dest1: 1, dest2: 6},
		{name: 1, data: []int{90, 55}, op: '+', arg: 1, div: 17, dest1: 6, dest2: 3},
		{name: 1, data: []int{91, 78, 80, 97, 79, 88}, op: '+', arg: 3, div: 19, dest1: 7, dest2: 5},
		{name: 1, data: []int{64, 80, 83, 89, 59}, op: '+', arg: 5, div: 3, dest1: 7, dest2: 2},
		{name: 1, data: []int{98, 92, 99, 51}, op: '2', div: 5, dest1: 0, dest2: 1},
		{name: 1, data: []int{68, 57, 95, 85, 98, 75, 98, 75}, op: '+', arg: 2, div: 13, dest1: 4, dest2: 0},
		{name: 1, data: []int{74}, op: '+', arg: 4, div: 7, dest1: 3, dest2: 2},
		{name: 1, data: []int{68, 64, 60, 68, 87, 80, 82}, op: '*', arg: 19, div: 11, dest1: 4, dest2: 5},
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

//type BigMonkey struct {
//	name  int
//	data  []big.Int
//	op    byte
//	arg   *big.Int
//	div   *big.Int
//	dest1 int
//	dest2 int
//	done  int
//}
//
//func (m *BigMonkey) String() string {
//	return fmt.Sprintf("Monkey %d: inspected %d", m.name, m.done)
//}
//func (m *BigMonkey) run2(monkey []*BigMonkey) {
//	for i := 0; i < len(m.data); i++ {
//		var new big.Int
//		if m.op == '*' {
//			new.Mul(m.arg, &m.data[i])
//		} else if m.op == '+' {
//			new.Add(m.arg, &m.data[i])
//		} else {
//			new.Mul(&m.data[i], &m.data[i])
//		}
//		var mod big.Int
//		mod.Mod(&new, m.div)
//		if mod.Cmp(big.NewInt(int64(0))) == 0 {
//			monkey[m.dest1].data = append(monkey[m.dest1].data, new)
//		} else {
//			monkey[m.dest2].data = append(monkey[m.dest2].data, new)
//		}
//		m.done++
//	}
//	m.data = nil
//}
//
//func (m *Monkey) toBigMonkey() *BigMonkey {
//	data := make([]big.Int, len(m.data))
//	for i, d := range m.data {
//		data[i].SetInt64(int64(d))
//	}
//	return &BigMonkey{
//		name:  m.name,
//		data:  data,
//		op:    m.op,
//		arg:   big.NewInt(int64(m.arg)),
//		div:   big.NewInt(int64(m.div)),
//		dest1: m.dest1,
//		dest2: m.dest2,
//	}
//}

type SmartMonkey struct {
	name  int
	data  []*Value
	op    byte
	arg   int
	div   int
	dest1 int
	dest2 int
	done  int
}

func (m *SmartMonkey) String() string {
	//return fmt.Sprintf("SmartMonkey %d: %v inspected %d", m.name, m.data, m.done)
	return fmt.Sprintf("SmartMonkey %d: inspected %d", m.name, m.done)
}

func (v *Value) String() string {
	return fmt.Sprintf("Value %d: %v", v.init, v.history)
}

func (m *SmartMonkey) evalmod(old int, div int) int {
	var new int
	old = old % div
	if m.op == '*' {
		new = int(float64(m.arg * old))
	} else if m.op == '+' {
		new = int(float64(m.arg + old))
	} else {
		new = int(float64(old * old))
	}
	return new % div
}

func (m *SmartMonkey) run2(monkeys []*SmartMonkey) {
	for i := 0; i < len(m.data); i++ {
		value := m.data[i]
		new := value.init
		for _, h := range value.history {
			new = monkeys[h].evalmod(new, m.div)
		}
		if new%m.div == 0 {
			value.history = append(value.history, m.dest1)
			monkeys[m.dest1].data = append(monkeys[m.dest1].data, value)
		} else {
			value.history = append(value.history, m.dest2)
			monkeys[m.dest2].data = append(monkeys[m.dest2].data, value)
		}
		m.done++
	}
	m.data = nil
}

type Value struct {
	init    int
	history []int
}

func Part2(input string) int {
	// input = strings.TrimSuffix(input, "\n")
	// lines := strings.Split(input, "\n")

	//monkeys := []*SmartMonkey{
	//	{name: 0, data: []*Value{{init: 79}, {init: 98}}, op: '*', arg: 19, div: 23, dest1: 2, dest2: 3},
	//	{name: 1, data: []*Value{{init: 54}, {init: 65}, {init: 75}, {init: 74}}, op: '+', arg: 6, div: 19, dest1: 2, dest2: 0},
	//	{name: 2, data: []*Value{{init: 79}, {init: 60}, {init: 97}}, op: '2', div: 13, dest1: 1, dest2: 3},
	//	{name: 3, data: []*Value{{init: 74}}, op: '+', arg: 3, div: 17, dest1: 0, dest2: 1},
	//}
	monkeys := []*SmartMonkey{
		{name: 0, data: []*Value{{init: 83}, {init: 62}, {init: 93}}, op: '*', arg: 17, div: 2, dest1: 1, dest2: 6},
		{name: 1, data: []*Value{{init: 90}, {init: 55}}, op: '+', arg: 1, div: 17, dest1: 6, dest2: 3},
		{name: 2, data: []*Value{{init: 91}, {init: 78}, {init: 80}, {init: 97}, {init: 79}, {init: 88}}, op: '+', arg: 3, div: 19, dest1: 7, dest2: 5},
		{name: 3, data: []*Value{{init: 64}, {init: 80}, {init: 83}, {init: 89}, {init: 59}}, op: '+', arg: 5, div: 3, dest1: 7, dest2: 2},
		{name: 4, data: []*Value{{init: 98}, {init: 92}, {init: 99}, {init: 51}}, op: '2', div: 5, dest1: 0, dest2: 1},
		{name: 5, data: []*Value{{init: 68}, {init: 57}, {init: 95}, {init: 85}, {init: 98}, {init: 75}, {init: 98}, {init: 75}}, op: '+', arg: 2, div: 13, dest1: 4, dest2: 0},
		{name: 6, data: []*Value{{init: 74}}, op: '+', arg: 4, div: 7, dest1: 3, dest2: 2},
		{name: 7, data: []*Value{{init: 68}, {init: 64}, {init: 60}, {init: 68}, {init: 87}, {init: 80}, {init: 82}}, op: '*', arg: 19, div: 11, dest1: 4, dest2: 5},
	}

	//data := map[int][]*Monkey{79: {monkeys[0]}, 98: {monkeys[0]}, 54: {monkeys[1]}, 65: {monkeys[1]}, 75: {monkeys[1]}, 74: {monkeys[1]}, 79: {monkeys[2]}, 60: {monkeys[2]}, 97: {monkeys[2]}, 74: {monkeys[3]}}
	//data := []Value{{79, []*Monkey{monkeys[0]}}, {98, []*Monkey{monkeys[0]}}, {54, []*Monkey{monkeys[1]}}, {65, []*Monkey{monkeys[1]}}, {75, []*Monkey{monkeys[1]}}, {74, []*Monkey{monkeys[1]}}, {79, []*Monkey{monkeys[2]}}, {60, []*Monkey{monkeys[2]}}, {97, []*Monkey{monkeys[2]}}, {74, []*Monkey{monkeys[3]}}}
	//data := []Value{{79, []int{0}}, {98, []int{0}}, {54, []int{1}}, {65, []int{1}}, {75, []int{1}}, {74, []int{1}}, {79, []int{2}}, {60, []int{2}}, {97, []int{2}}, {74, []int{3}}}
	//data := []*Value{}
	for i, m := range monkeys {
		for _, v := range m.data {
			v.history = append(v.history, i)
			//data = append(data, v)
		}
	}

	//monkeys := []*Monkey{
	//	{name: 0, data: []int{83, 62, 93}, op: '*', arg: 17, div: 2, dest1: 1, dest2: 6},
	//	{name: 1, data: []int{90, 55}, op: '+', arg: 1, div: 17, dest1: 6, dest2: 3},
	//	{name: 1, data: []int{91, 78, 80, 97, 79, 88}, op: '+', arg: 3, div: 19, dest1: 7, dest2: 5},
	//	{name: 1, data: []int{64, 80, 83, 89, 59}, op: '+', arg: 5, div: 3, dest1: 7, dest2: 2},
	//	{name: 1, data: []int{98, 92, 99, 51}, op: '2', div: 5, dest1: 0, dest2: 1},
	//	{name: 1, data: []int{68, 57, 95, 85, 98, 75, 98, 75}, op: '+', arg: 2, div: 13, dest1: 4, dest2: 0},
	//	{name: 1, data: []int{74}, op: '+', arg: 4, div: 7, dest1: 3, dest2: 2},
	//	{name: 1, data: []int{68, 64, 60, 68, 87, 80, 82}, op: '*', arg: 19, div: 11, dest1: 4, dest2: 5},
	//}

	for i := 0; i < 10000; i++ {
		for _, m := range monkeys {
			m.run2(monkeys)
		}
		if (i+1)%1000 == 0 {
			fmt.Println("after round", i+1)
			for _, m := range monkeys {
				fmt.Println(m.String())
			}
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
