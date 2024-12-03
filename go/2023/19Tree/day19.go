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

type Object map[uint8]int

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

func (o Object) String() string {
	return fmt.Sprintf("x=%d, m=%d, a=%d, s=%d", o['x'], o['m'], o['a'], o['s'])
}

func parseInstructions(rules map[string][]string, instr []string) Expr {
	if len(instr) == 0 {
		panic("empty list")
	}
	head := instr[0]
	if len(instr) == 1 {
		if head == "A" {
			return Accepted{}
		}
		if head == "R" {
			return Rejected{}
		}

		return parseInstructions(rules, rules[head])
	}
	subject := head[0]
	cond := head[1]
	value, _ := strconv.Atoi(head[2:])
	then := instr[1]
	tail := instr[2:]

	if cond == '<' {
		return IfThenElse{
			cond:  Lt{subject, value},
			then:  parseInstructions(rules, rules[then]),
			else_: parseInstructions(rules, tail),
		}
	}
	if cond == '>' {
		return IfThenElse{
			cond:  Gt{subject, value},
			then:  parseInstructions(rules, rules[then]),
			else_: parseInstructions(rules, tail),
		}
	}
	panic("not implemented")
}

func parseRules(lines []string, start string) Expr {
	var rules = make(map[string][]string)
	for _, line := range lines {
		name, after, _ := strings.Cut(line, "{")
		fields := strings.FieldsFunc(after, func(c rune) bool { return c == ',' || c == ':' || c == '}' })
		rules[name] = fields
	}
	rules["A"] = []string{"A"}
	rules["R"] = []string{"R"}
	return parseInstructions(rules, rules[start])
}

// AST

type Expr interface {
	isExpr()
	apply(obj Object) int
	propagate(constraint Constraint) int
}
type Cond interface {
	isCond()
	apply(obj Object) bool
	propagate(constraint Constraint) (Constraint, Constraint)
}

type IfThenElse struct {
	cond  Cond
	then  Expr
	else_ Expr
}

type Accepted struct{}
type Rejected struct{}

type Lt struct {
	xmas  uint8
	value int
}
type Gt struct {
	xmas  uint8
	value int
}

// func (_ Var) isExpr()        {}
func (_ IfThenElse) isExpr() {}
func (_ Accepted) isExpr()   {}
func (_ Rejected) isExpr()   {}
func (_ Lt) isCond()         {}
func (_ Gt) isCond()         {}

func (e Accepted) String() string {
	return "A"
}
func (e Rejected) String() string {
	return "R"
}
func (e IfThenElse) String() string {
	return fmt.Sprintf("if %v then %v else %v", e.cond, e.then, e.else_)
}
func (c Lt) String() string {
	return fmt.Sprintf("%c < %d", c.xmas, c.value)
}
func (c Gt) String() string {
	return fmt.Sprintf("%c > %d", c.xmas, c.value)
}

func (c Lt) apply(obj Object) bool {
	return obj[c.xmas] < c.value
}
func (c Gt) apply(obj Object) bool {
	return obj[c.xmas] > c.value
}
func (e Accepted) apply(obj Object) int {
	return obj['x'] + obj['m'] + obj['a'] + obj['s']
}
func (e Rejected) apply(obj Object) int {
	return 0
}

func (e IfThenElse) apply(obj Object) int {
	if e.cond.apply(obj) {
		return e.then.apply(obj)
	} else {
		return e.else_.apply(obj)
	}
}

func (c Lt) propagate(constraint Constraint) (Constraint, Constraint) {
	pos := constraint.copy()
	neg := constraint.copy()
	name := c.xmas
	pos[name] = constraint[name].Inter(interval.Interval{1, c.value - 1})
	neg[name] = constraint[name].Inter(interval.Interval{c.value, 4000})
	return pos, neg
}

func (c Gt) propagate(constraint Constraint) (Constraint, Constraint) {
	pos := constraint.copy()
	neg := constraint.copy()
	name := c.xmas
	pos[name] = constraint[name].Inter(interval.Interval{c.value + 1, 4000})
	neg[name] = constraint[name].Inter(interval.Interval{1, c.value})
	return pos, neg
}

func (e Accepted) propagate(constraint Constraint) int {
	res := 1
	for _, v := range constraint {
		res *= v.Len()
	}
	return res
}

func (e Rejected) propagate(constraint Constraint) int {
	return 0
}

func (e IfThenElse) propagate(constraint Constraint) int {
	pos, neg := e.cond.propagate(constraint)
	return e.then.propagate(pos) + e.else_.propagate(neg)
}

type Constraint map[uint8]interval.Interval

func (c Constraint) String() string {
	return fmt.Sprintf("x=%v, m=%v, a=%v, s=%v", c['x'], c['m'], c['a'], c['s'])
}

func (c Constraint) copy() Constraint {
	return Constraint{'x': c['x'], 'm': c['m'], 'a': c['a'], 's': c['s']}
}

func Part1(input string) int {
	input = strings.TrimSpace(input)
	parts := strings.Split(input, "\n\n")
	lines := strings.Split(parts[0], "\n")

	var expr = parseRules(lines, "in")

	var res int
	lines = strings.Split(parts[1], "\n")
	for _, line := range lines {
		object := parseObject(line)
		res += expr.apply(object)
	}

	return res
}

func Part2(input string) int {
	input = strings.TrimSpace(input)
	parts := strings.Split(input, "\n\n")
	lines := strings.Split(parts[0], "\n")

	var expr = parseRules(lines, "in")

	start := interval.Interval{1, 4000}
	var c = Constraint{'x': start, 'm': start, 'a': start, 's': start}
	return expr.propagate(c)
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
