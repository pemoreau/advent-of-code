package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
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

func (b *box) step(from string, signal bool, table map[string]*box, w *wires) bool {
	switch b.model {
	case FLIP:
		//fmt.Printf("%s -%v-> %s (%v)\n", from, signal, b.name, b.state)
		if signal {
			// ignore
			return false
		}
		b.state = !b.state
		for _, name := range b.outputs {
			w.push(pulse{b.name, b.state, name})
		}

	case CONJ:
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
		if b.name == "jq" && signal {
			// gt, vr, nl, lr -> &jq
			// one entry is high
			return true
		}

	case WIRE:
		//fmt.Printf("%s -%v-> %s(%d)\n", from, signal, b.name, b.model)
		for _, name := range b.outputs {
			w.push(pulse{b.name, signal, name})
		}
	}
	return false
}

func parse(input string) (table map[string]*box, w *wires) {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")

	table = make(map[string]*box)
	q := make([]pulse, 0)
	w = &wires{queue: &q}

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
			//fmt.Println("not found", name)
			table[name] = NewBox(name, WIRE, []string{})
		}
	}

	for _, box := range table {
		box.init(table)
	}
	return table, w
}

func Part1(input string) int {
	table, w := parse(input)

	for i := 0; i < 1000; i++ {
		w.push(pulse{"button", false, "broadcaster"})
		for !w.isEmpty() {
			p := w.pop()
			table[p.to].step(p.from, p.value, table, w)
		}
	}
	return w.nbLow * w.nbHigh
}

func Part2(input string) int {
	table, w := parse(input)
	var gt, vr, nl, lr int
	for i := 1; i < 10000000; i++ {
		w.push(pulse{"button", false, "broadcaster"})
		for !w.isEmpty() {
			p := w.pop()
			found := table[p.to].step(p.from, p.value, table, w)
			if found {
				if p.from == "gt" && gt == 0 {
					gt = i
				}
				if p.from == "vr" && vr == 0 {
					vr = i
				}
				if p.from == "nl" && nl == 0 {
					nl = i
				}
				if p.from == "lr" && lr == 0 {
					lr = i
				}
				if gt != 0 && vr != 0 && nl != 0 && lr != 0 {
					return utils.LCM(gt, vr, nl, lr)
				}
			}
		}
	}
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
