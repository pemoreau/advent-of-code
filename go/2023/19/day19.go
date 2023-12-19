package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

type Instruction struct {
	cond    uint8 // <, > or T
	subject uint8 // x, m, a or s
	value   int
	then    string
}

type Rule struct {
	name  string
	instr []Instruction
}

type Object map[uint8]int

func (o Object) String() string {
	return fmt.Sprintf("x=%d, m=%d, a=%d, s=%d", o['x'], o['m'], o['a'], o['s'])
}

func (r Rule) String() string {
	return fmt.Sprintf("%s: %v", r.name, r.instr)
}

func (i Instruction) String() string {
	if i.cond == 'T' {
		return fmt.Sprintf("else: %s", i.then)
	}
	return fmt.Sprintf("%c%c%d => %s", i.subject, i.cond, i.value, i.then)
}

func (r Rule) apply(obj Object) string {
	fmt.Printf("apply %s to %v\n", r, obj)
	for _, instr := range r.instr {
		switch instr.cond {
		case '<':
			fmt.Printf("compare %c:%d < %d\n", instr.subject, obj[instr.subject], instr.value)
			if obj[instr.subject] < instr.value {
				return instr.then
			}
		case '>':
			fmt.Printf("compare %c:%d > %d\n", instr.subject, obj[instr.subject], instr.value)
			if obj[instr.subject] > instr.value {
				return instr.then
			}
		case 'T':
			return instr.then
		}
	}
	panic("no instruction found")
}

func parseRule(line string) Rule {
	splitFunc := func(c rune) bool {
		return c == '{' || c == '}' || c == ',' || c == ':'
	}
	fields := strings.FieldsFunc(line, splitFunc)
	name := fields[0]
	other := fields[len(fields)-1]
	var instructions []Instruction
	for i := 1; i < len(fields)-1; i += 2 {
		//fmt.Printf("i=%d, fields[i]=%s\n", i, fields[i])
		subject := fields[i][0]
		cond := fields[i][1]
		value, _ := strconv.Atoi(fields[i][2:])
		then := fields[i+1]
		instructions = append(instructions, Instruction{
			cond:    cond,
			subject: subject,
			value:   value,
			then:    then,
		})
	}
	instructions = append(instructions, Instruction{
		cond: 'T',
		then: other,
	})

	//fmt.Println(fields)
	return Rule{name: name, instr: instructions}
}

func parseObject(line string) Object {
	splitFunc := func(c rune) bool {
		return c == '{' || c == '}' || c == ',' || c == '='
	}
	fields := strings.FieldsFunc(line, splitFunc)
	var res = make(Object)
	for i := 0; i < len(fields); i += 2 {
		value, _ := strconv.Atoi(fields[i+1])
		res[fields[i][0]] = value
	}
	return res
}

func run(rules map[string]Rule, object Object) int {
	var pc = "in"
	for {
		rule := rules[pc]
		label := rule.apply(object)
		if label == "R" {
			fmt.Println("Rejected")
			return 0
		}
		if label == "A" {
			fmt.Println("Accepted")
			return object['x'] + object['m'] + object['a'] + object['s']
		}
		fmt.Println("pc=", label)
		pc = label
	}
}

func Part1(input string) int {
	input = strings.TrimSpace(input)
	parts := strings.Split(input, "\n\n")
	lines := strings.Split(parts[0], "\n")

	var rules = make(map[string]Rule)
	var objects []Object

	for _, line := range lines {
		rule := parseRule(line)
		rules[rule.name] = rule
	}

	lines = strings.Split(parts[1], "\n")
	for _, line := range lines {
		object := parseObject(line)
		objects = append(objects, object)
	}

	var res int
	for _, object := range objects {
		v := run(rules, object)
		fmt.Println(object, v)
		res += v
	}

	return res
}

func Part2(input string) int {
	return 0
}

func main() {
	fmt.Println("--2023 day 19 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
