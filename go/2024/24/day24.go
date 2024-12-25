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

//go:embed sample2.txt
var inputTest2 string

type Node struct {
	op  string
	lhs string
	rhs string
	val int
}

func eval(variable string, values map[string]Node, visited map[string]int) (int, error) {
	if visited[variable] > 100 {
		return 0, fmt.Errorf("circular reference %s", variable)
	}
	//fmt.Printf("eval %s\n", variable)
	if v, ok := values[variable]; ok {
		visited[variable]++
		if v.op == "" && (v.val == 1 || v.val == 0) {
			//fmt.Printf("return %d\n", v.val)
			return v.val, nil
		}
		var lhs, rhs int
		var err error

		if lhs, err = eval(v.lhs, values, visited); err != nil {
			return 0, err
		}
		if rhs, err = eval(v.rhs, values, visited); err != nil {
			return 0, err
		}
		switch v.op {
		case "AND":
			//fmt.Printf("return %d & %d = %d\n", lhs, rhs, lhs&rhs)
			return lhs & rhs, nil
		case "OR":
			//fmt.Printf("return %d | %d = %d\n", lhs, rhs, lhs|rhs)
			return lhs | rhs, nil
		case "XOR":
			//fmt.Printf("return %d ^ %d = %d\n", lhs, rhs, lhs^rhs)
			return lhs ^ rhs, nil
		}
	}
	return 0, fmt.Errorf("variable %s not found", variable)
}

func run(wires map[string]Node) (int, error) {
	var z []string
	for k, _ := range wires {
		if strings.HasPrefix(k, "z") {
			z = append(z, k)
		}
	}
	slices.Sort(z)

	//fmt.Println(z)

	var res int
	for _, k := range slices.Backward(z) {
		var b, err = eval(k, wires, map[string]int{})
		if err != nil {
			return 0, err
		}
		//fmt.Println(k, b)
		res = res<<1 + b
	}
	return res, nil
}

func setValue(name string, value int, wires map[string]Node) {
	var index = 0
	for value != 0 {
		wires[fmt.Sprintf("%s%02d", name, index)] = Node{val: value % 2}
		//fmt.Printf("%s%02d = %d\n", name, index, value%2)
		value = value >> 1
		index++
	}
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
		//fmt.Println(line)
		fmt.Sscanf(line, "%s %d", &variable, &value)
		wires[variable[:len(variable)-1]] = Node{val: value}
		//fmt.Printf("add %s %d: %v\n", variable, value, wires[variable])
	}

	for _, line := range strings.Split(parts[1], "\n") {
		var op, lhs, rhs, variable string
		fmt.Sscanf(line, "%s %s %s -> %s", &lhs, &op, &rhs, &variable)
		wires[variable] = Node{op: op, lhs: lhs, rhs: rhs}
	}

	return wires
}

func checkZ(z string, wires map[string]Node) bool {
	xor_output, ok := wires[z]
	// z output
	if !ok || xor_output.op != "XOR" {
		fmt.Printf("%s: bad output expected XOR got: %v\n", z, xor_output)
		return false
	}
	var lhs = wires[xor_output.lhs]
	var rhs = wires[xor_output.rhs]
	if !((lhs.op == "XOR" && rhs.op == "OR") || (lhs.op == "OR" && rhs.op == "XOR")) {
		fmt.Printf("%s: bad input for xor_output expected: XOR and OR got: %v %v\n", z, lhs, rhs)
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
	xname := "x" + z[1:]
	yname := "y" + z[1:]
	if !((and_input.lhs == xname && and_input.rhs == yname) || (and_input.lhs == yname && and_input.rhs == xname)) {
		fmt.Printf("%s: bad input for and_input expected:%s %s got: %v\n", z, xname, yname, and_input)
		return false
	}
	if !((xor_input.lhs == xname && xor_input.rhs == yname) || (xor_input.lhs == yname && xor_input.rhs == xname)) {
		fmt.Printf("%s: bad input for xor_input expected:%s %s got: %v\n", z, xname, yname, xor_input)
		return false
	}
	return true
}

func checkSomme(zname string, wires map[string]Node) bool {
	xname := "x" + zname[1:]
	yname := "y" + zname[1:]
	xvalue := wires[xname]
	yvalue := wires[yname]
	inputsX := []int{0, 0, 1, 1}
	inputsY := []int{0, 1, 0, 1}
	expected := []int{0, 1, 1, 0}
	var res = true
	for i, inputX := range inputsX {
		inputY := inputsY[i]
		wires[xname] = Node{val: inputX}
		wires[yname] = Node{val: inputY}
		var b, err = eval(zname, wires, map[string]int{})
		if err != nil {
			fmt.Printf("checkSomme %s: %s=%d %s=%d error: %v\n", zname, xname, inputX, yname, inputY, err)
			res = false
			break
		}
		if b != expected[i] {
			fmt.Printf("checkSomme %s: %s=%d %s=%d expected %d got %d\n", zname, xname, inputX, yname, inputY, expected[i], b)
			res = false
			break
		}
	}
	wires[xname] = xvalue
	wires[yname] = yvalue
	return res
}

func Part1(input string) int {
	var wires = parse(input)
	v, err := run(wires)
	if err != nil {
		fmt.Println(err)
	}
	return v
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
	for _, zname := range slices.Backward(znodes) {
		// save x and y values
		for i, zname := range znodes {
			xname := "x" + zname[1:]
			yname := "y" + zname[1:]
			xvalues[i] = wires[xname]
			yvalues[i] = wires[yname]
			// set to 0
			wires[xname] = Node{val: 0}
			wires[yname] = Node{val: 0}
		}

		ok := checkSomme(zname, wires)
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
