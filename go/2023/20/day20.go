package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

const (
	FLIP = iota
	CONJ
	WIRE
)

type box struct {
	name     string
	model    int
	state    bool
	remember map[string]bool
	outputs  []string
	//nbLow   int
	//nbHigh  int
	//queue   *[]pulse
}

type pulse struct {
	from  string
	value bool
	to    string
}

type wires struct {
	queue  *[]pulse
	nbLow  int
	nbHigh int
}

func (w *wires) push(p pulse) {
	*w.queue = append(*w.queue, p)
	//fmt.Println("push", *w.queue)
	if p.value {
		w.nbHigh++
	} else {
		w.nbLow++
	}
}

func (w *wires) isEmpty() bool {
	return len(*w.queue) == 0
}

func (w *wires) pop() pulse {
	p := (*w.queue)[0]
	*w.queue = (*w.queue)[1:]
	return p
}

func NewBox(name string, model int, outputs []string) *box {
	return &box{name: name, model: model, remember: make(map[string]bool), outputs: outputs}
}

func (b *box) init(table map[string]*box) {
	for _, name := range b.outputs {
		outputBox := table[name]
		outputBox.remember[b.name] = false
	}
}

//var nbLow int
//var nbHigh int

//func (b *box) push(from string, signal bool, table map[string]*box) {
//	b.inputs[from] = append(b.inputs[from], signal)
//	if signal {
//		nbHigh++
//	} else {
//		nbLow++
//	}
//	fmt.Printf("%s -%v-> %s(%d)\n", from, signal, b.name, b.model)
//	b.step(from, table)
//}

func (b *box) step(from string, signal bool, table map[string]*box, w *wires) {
	switch b.model {
	case FLIP:
		fmt.Printf("%s -%v-> %s (%v)\n", from, signal, b.name, b.state)
		if signal {
			// ignore
			return
		}
		b.state = !b.state
		for _, name := range b.outputs {
			w.push(pulse{b.name, b.state, name})
		}
		//for _, name := range b.outputs {
		//	table[name].step(b.name, table)
		//}

	case CONJ:
		fmt.Printf("%s -%v-> %s(%d)\n", from, signal, b.name, b.model)
		b.remember[from] = signal
		// check all state
		var res bool
		for _, state := range b.remember {
			if !state {
				res = true
				break
			}
		}
		for _, name := range b.outputs {
			w.push(pulse{b.name, res, name})
		}
		//for _, name := range b.outputs {
		//	table[name].step(b.name, table)
		//}
	case WIRE:
		fmt.Printf("%s -%v-> %s(%d)\n", from, signal, b.name, b.model)
		for _, name := range b.outputs {
			w.push(pulse{b.name, signal, name})
		}
		//for _, name := range b.outputs {
		//	table[name].step(b.name, table)
		//}
	}

}

func Part1(input string) int {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")

	table := make(map[string]*box)
	q := make([]pulse, 0)
	w := &wires{queue: &q}

	var names []string
	for _, line := range lines {
		name, after, _ := strings.Cut(line, " -> ")
		outputs := strings.Split(after, ", ")
		names = append(names, outputs...)
		if name[0] == '%' {
			name = name[1:]
			//fmt.Printf("flip %s -> %v\n", name, outputs)
			table[name] = NewBox(name, FLIP, outputs)
		} else if name[0] == '&' {
			name = name[1:]
			//fmt.Printf("conj %s -> %v\n", name, outputs)
			table[name] = NewBox(name, CONJ, outputs)
		} else {
			//fmt.Printf("wire %s -> %v\n", name, outputs)
			table[name] = NewBox(name, WIRE, outputs)
		}
	}

	for _, name := range names {
		if _, ok := table[name]; !ok {
			fmt.Println("not found", name)
			table[name] = NewBox(name, WIRE, []string{})
		}
	}

	for _, box := range table {
		box.init(table)
	}

	for i := 0; i < 1000; i++ {
		w.push(pulse{"button", false, "broadcaster"})
		for !w.isEmpty() {
			p := w.pop()
			table[p.to].step(p.from, p.value, table, w)
		}
	}
	fmt.Printf("nbLow: %d, nbHigh: %d\n", w.nbLow, w.nbHigh)
	//table["broadcaster"].push("button", false, table)
	//table["broadcaster"].step("button", table)

	return w.nbLow * w.nbHigh
}

func Part2(input string) int {
	return 0
}

func main() {
	fmt.Println("--2023 day 20 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
