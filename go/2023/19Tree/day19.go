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

type Rule []Instruction

type Object map[uint8]int

func (i Instruction) String() string {
	if i.cond == 'T' {
		return fmt.Sprintf("else: %s", i.then)
	}
	return fmt.Sprintf("%c%c%d => %s", i.subject, i.cond, i.value, i.then)
}

func (o Object) String() string {
	return fmt.Sprintf("x=%d, m=%d, a=%d, s=%d", o['x'], o['m'], o['a'], o['s'])
}

func parseRule(line string) (string, Rule) {
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
	return name, instructions
}

func parseRules(lines []string) map[string]Rule {
	var rules = make(map[string]Rule)
	for _, line := range lines {
		name, rule := parseRule(line)
		rules[name] = rule
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

type Var uint8

type Lt struct {
	v     Var
	value int
}
type Gt struct {
	v     Var
	value int
}

func (_ Var) isExpr()        {}
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
	return fmt.Sprintf("%c < %d", c.v, c.value)
}
func (c Gt) String() string {
	return fmt.Sprintf("%c > %d", c.v, c.value)
}

func buildExpr(rules map[string]Rule, pc string) Expr {
	if pc == "A" {
		return Accepted{}
	}
	if pc == "R" {
		return Rejected{}
	}
	return buildRule(rules[pc], rules)
}

func buildRule(instr []Instruction, rules map[string]Rule) Expr {
	if len(instr) == 0 {
		panic("empty list")
	}
	if len(instr) == 1 {
		name := instr[0].then
		return buildExpr(rules, name)
	}
	head := instr[0]
	tail := instr[1:]
	if head.cond == '<' {
		return IfThenElse{
			cond:  Lt{Var(head.subject), head.value},
			then:  buildExpr(rules, head.then),
			else_: buildRule(tail, rules),
		}
	}
	if head.cond == '>' {
		return IfThenElse{
			cond:  Gt{Var(head.subject), head.value},
			then:  buildExpr(rules, head.then),
			else_: buildRule(tail, rules),
		}
	}
	panic("not implemented")
}

func (c Lt) apply(obj Object) bool {
	return obj[uint8(c.v)] < c.value
}
func (c Gt) apply(obj Object) bool {
	return obj[uint8(c.v)] > c.value
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
	name := uint8(c.v)
	pos[name] = constraint[name].Inter(interval.Interval{1, c.value - 1})
	neg[name] = constraint[name].Inter(interval.Interval{c.value, 4000})
	return pos, neg
}

func (c Gt) propagate(constraint Constraint) (Constraint, Constraint) {
	pos := constraint.copy()
	neg := constraint.copy()
	name := uint8(c.v)
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

	var rules = parseRules(lines)
	var expr = buildExpr(rules, "in")

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

	var rules = parseRules(lines)
	var expr = buildExpr(rules, "in")

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
