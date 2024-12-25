package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"slices"
	"strings"
	"time"
)

//go:embed sample.txt
var inputTest string

type Node struct {
	op  string
	lhs string
	rhs string
	val int
}

func eval(variable string, values map[string]Node) int {
	if v, ok := values[variable]; ok {
		switch v.op {
		case "":
			return v.val
		case "AND":
			return eval(v.lhs, values) & eval(v.rhs, values)
		case "OR":
			return eval(v.lhs, values) | eval(v.rhs, values)
		case "XOR":
			return eval(v.lhs, values) ^ eval(v.rhs, values)
		}
	}
	return 0
}

func run(wires map[string]Node) int {
	var z []string
	for k, _ := range wires {
		if strings.HasPrefix(k, "z") {
			z = append(z, k)
		}
	}
	slices.Sort(z)

	var res int
	for _, k := range slices.Backward(z) {
		var b = eval(k, wires)
		res = res<<1 + b
	}
	return res
}

func swapWires(w1, w2 string, wires map[string]Node) {
	var tmp = wires[w1]
	wires[w1] = wires[w2]
	wires[w2] = tmp
}

func parse(input string) map[string]Node {
	var parts = strings.Split(input, "\n\n")
	var wires = make(map[string]Node)

	for _, line := range strings.Split(parts[0], "\n") {
		var value int
		var variable string
		fmt.Sscanf(line, "%s %d", &variable, &value)
		wires[variable[:len(variable)-1]] = Node{val: value}
	}

	for _, line := range strings.Split(parts[1], "\n") {
		var op, lhs, rhs, variable string
		fmt.Sscanf(line, "%s %s %s -> %s", &lhs, &op, &rhs, &variable)
		wires[variable] = Node{op: op, lhs: lhs, rhs: rhs}
	}

	return wires
}

// x01 XOR y01 -> A
// y01 AND x01 -> B
// A AND C -> D
// A XOR C -> z01
// B OR D -> E

func checkZ(zname string, wires map[string]Node) bool {
	xname := "x" + zname[1:]
	yname := "y" + zname[1:]

	xor_output, ok := wires[zname]
	// zname output
	if !ok || xor_output.op != "XOR" {
		fmt.Printf("%s: bad output expected XOR got: %v\n", zname, xor_output)
		return false
	}
	var lhs = wires[xor_output.lhs]
	var rhs = wires[xor_output.rhs]
	if !((lhs.op == "XOR" && rhs.op == "OR") || (lhs.op == "OR" && rhs.op == "XOR")) {
		fmt.Printf("%s: bad input for xor_output expected: XOR and OR got: %v %v\n", zname, lhs, rhs)
		return false
	}
	var xor_input, and_input Node
	if lhs.op == "XOR" {
		xor_input = wires[lhs.lhs]
		and_input = wires[rhs.lhs]
	} else {
		xor_input = wires[lhs.rhs]
		and_input = wires[rhs.rhs]
	}

	// check xor inputs

	if !((and_input.lhs == xname && and_input.rhs == yname) || (and_input.lhs == yname && and_input.rhs == xname)) {
		fmt.Printf("%s: bad input for and_input expected:%s %s got: %v\n", zname, xname, yname, and_input)
		return false
	}
	if !((xor_input.lhs == xname && xor_input.rhs == yname) || (xor_input.lhs == yname && xor_input.rhs == xname)) {
		fmt.Printf("%s: bad input for xor_input expected:%s %s got: %v\n", zname, xname, yname, xor_input)
		return false
	}
	return true
}

func checkSomme(zname string, wires map[string]Node) bool {
	xname := "x" + zname[1:]
	yname := "y" + zname[1:]
	inputsX := []int{0, 0, 1, 1}
	inputsY := []int{0, 1, 0, 1}
	expected := []int{0, 1, 1, 0}
	for i, inputX := range inputsX {
		inputY := inputsY[i]
		wires[xname] = Node{val: inputX}
		wires[yname] = Node{val: inputY}
		if b := eval(zname, wires); b != expected[i] {
			fmt.Printf("checkSomme %s: %s=%d %s=%d expected %d got %d\n", zname, xname, inputX, yname, inputY, expected[i], b)
			return false
		}
	}
	return true
}

func Part1(input string) int {
	var wires = parse(input)
	return run(wires)
}

func Part2(input string) string {
	var wires = parse(input)

	var znodes []string
	for k, _ := range wires {
		if strings.HasPrefix(k, "z") {
			znodes = append(znodes, k)
		}
	}
	slices.Sort(znodes)

	var xvalues = make([]Node, len(znodes))
	var yvalues = make([]Node, len(znodes))

	swapWires("z11", "rpv", wires) // z11
	swapWires("rpb", "ctg", wires) // z15
	swapWires("z31", "dmh", wires) // z31
	swapWires("z38", "dvq", wires) // z38

	//ctg,dmh,dvq,rpb,rpv,z11,z31,z38
	for i := len(znodes) - 2; i >= 0; i-- {
		// save x and y values
		for i, zname := range znodes {
			xname := "x" + zname[1:]
			yname := "y" + zname[1:]
			xvalues[i] = wires[xname]
			yvalues[i] = wires[yname]
			// set x and y to 0
			wires[xname] = Node{val: 0}
			wires[yname] = Node{val: 0}
		}

		ok := checkSomme(znodes[i], wires)
		if ok {
			//fmt.Printf("check %s: %v\n", zname, ok)
		}

		// restore x and y values
		for i, zname := range znodes {
			xname := "x" + zname[1:]
			yname := "y" + zname[1:]
			wires[xname] = xvalues[i]
			wires[yname] = yvalues[i]
		}

	}

	return "ctg,dmh,dvq,rpb,rpv,z11,z31,z38"
}

func main() {
	fmt.Println("--2024 day 24 solution--")
	var inputDay = utils.Input()
	//var inputDay = inputTest
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
