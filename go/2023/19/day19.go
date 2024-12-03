package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils/interval"
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

func (i Instruction) String() string {
	if i.cond == 'T' {
		return fmt.Sprintf("else: %s", i.then)
	}
	return fmt.Sprintf("%c%c%d => %s", i.subject, i.cond, i.value, i.then)
}

func (r Rule) String() string {
	return fmt.Sprintf("%s: %v", r.name, r.instr)
}

func (o Object) String() string {
	return fmt.Sprintf("x=%d, m=%d, a=%d, s=%d", o['x'], o['m'], o['a'], o['s'])
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
	return Rule{name: name, instr: instructions}
}

func parseRules(lines []string) map[string]Rule {
	var rules = make(map[string]Rule)
	for _, line := range lines {
		rule := parseRule(line)
		rules[rule.name] = rule
	}
	return rules
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

func (r Rule) apply(obj Object) string {
	for _, instr := range r.instr {
		switch instr.cond {
		case '<':
			if obj[instr.subject] < instr.value {
				return instr.then
			}
		case '>':
			if obj[instr.subject] > instr.value {
				return instr.then
			}
		case 'T':
			return instr.then
		}
	}
	panic("no instruction found")
}

func run(rules map[string]Rule, object Object) int {
	var pc = "in"
	for {
		rule := rules[pc]
		label := rule.apply(object)
		if label == "R" {
			return 0
		}
		if label == "A" {
			return object['x'] + object['m'] + object['a'] + object['s']
		}
		pc = label
	}
}

func Part1(input string) int {
	input = strings.TrimSpace(input)
	parts := strings.Split(input, "\n\n")
	lines := strings.Split(parts[0], "\n")

	var rules = parseRules(lines)

	var objects []Object
	lines = strings.Split(parts[1], "\n")
	for _, line := range lines {
		object := parseObject(line)
		objects = append(objects, object)
	}

	var res int
	for _, object := range objects {
		v := run(rules, object)
		res += v
	}

	return res
}

type Constraint map[uint8]interval.Interval

func (c Constraint) String() string {
	return fmt.Sprintf("x=%v, m=%v, a=%v, s=%v", c['x'], c['m'], c['a'], c['s'])
}

func (c Constraint) copy() Constraint {
	return Constraint{'x': c['x'], 'm': c['m'], 'a': c['a'], 's': c['s']}
}

func (i Instruction) apply(c Constraint) (pos Constraint, neg Constraint) {
	if i.cond == 'T' {
		return c, Constraint{'x': interval.Empty(), 'm': interval.Empty(), 'a': interval.Empty(), 's': interval.Empty()}
	}

	var i1, i2 interval.Interval
	switch i.cond {
	case '<':
		i1 = interval.Interval{1, i.value - 1}
		i2 = interval.Interval{i.value, 4000}
	case '>':
		i1 = interval.Interval{i.value + 1, 4000}
		i2 = interval.Interval{1, i.value}
	}
	pos = c.copy()
	neg = c.copy()
	pos[i.subject] = c[i.subject].Inter(i1)
	neg[i.subject] = c[i.subject].Inter(i2)
	return pos, neg
}

func (c Constraint) combinaison() int {
	res := 1
	for _, v := range c {
		res *= v.Len()
	}
	return res
}

func propagate(rules map[string]Rule, c Constraint, pc string) int {
	if pc == "A" {
		return c.combinaison()
	}
	if pc == "R" {
		return 0
	}
	var res int
	rule := rules[pc]
	for _, instr := range rule.instr {
		pos, neg := instr.apply(c)
		res += propagate(rules, pos, instr.then)
		c = neg
	}
	return res
}

func Part2(input string) int {
	input = strings.TrimSpace(input)
	parts := strings.Split(input, "\n\n")
	lines := strings.Split(parts[0], "\n")

	var rules = make(map[string]Rule)

	for _, line := range lines {
		rule := parseRule(line)
		rules[rule.name] = rule
	}
	start := interval.Interval{1, 4000}
	var c = Constraint{'x': start, 'm': start, 'a': start, 's': start}
	res := propagate(rules, c, "in")
	return res
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
