package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"strconv"
	"strings"
	"time"
)

//go:embed sample.txt
var inputTest string

type Rule interface {
	isRule()
}

type TerminalRule struct {
	value string
}

type SequenceRule struct {
	rules []int
}

type AlternativeRule struct {
	rules []Rule
}

func (TerminalRule) isRule()    {}
func (SequenceRule) isRule()    {}
func (AlternativeRule) isRule() {}

func parseLine(line string) (int, Rule) {
	parts := strings.Split(line, ":")
	ruleNumber, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}
	rule := parseRule(parts[1])
	return ruleNumber, rule
}

func parseRule(rule string) Rule {
	rule = strings.TrimSpace(rule)
	if strings.Contains(rule, "\"") {
		return TerminalRule{value: strings.Trim(rule, "\"")}
	}
	if strings.Contains(rule, "|") {
		parts := strings.Split(rule, "|")
		var rules []Rule
		for _, part := range parts {
			rules = append(rules, parseRule(part))
		}
		return AlternativeRule{rules: rules}
	}
	parts := strings.Split(rule, " ")
	var rules []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		rules = append(rules, n)
	}
	return SequenceRule{rules: rules}
}

type Grammar = map[int]Rule

func buildGrammar(lines []string) Grammar {
	grammar := make(Grammar)
	for _, line := range lines {
		ruleNumber, rule := parseLine(line)
		grammar[ruleNumber] = rule
	}
	return grammar
}

func recognize(s string, index int, grammar Grammar, rule Rule) (newIndex int, err error) {
	if index >= len(s) {
		return index, fmt.Errorf("out of bounds")
	}
	switch rule.(type) {
	case TerminalRule:
		if s[index:index+1] != rule.(TerminalRule).value {
			return -1, fmt.Errorf("not a match: %s expected", rule.(TerminalRule).value)
		}
		return index + 1, nil
	case SequenceRule:
		newIndex = index
		for _, ruleNumber := range rule.(SequenceRule).rules {
			newIndex, err = recognize(s, newIndex, grammar, grammar[ruleNumber])
			if err != nil {
				return -1, err
			}
		}
		return newIndex, nil
	case AlternativeRule:
		for _, rule := range rule.(AlternativeRule).rules {
			newIndex, err = recognize(s, index, grammar, rule)
			if err == nil {
				return newIndex, nil
			}
		}
		return -1, fmt.Errorf("no match")
	}
	return -1, fmt.Errorf("unknown rule type")
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	parts := strings.Split(input, "\n\n")
	grammar := buildGrammar(strings.Split(parts[0], "\n"))

	messages := strings.Split(parts[1], "\n")
	count := 0
	for _, message := range messages {
		index, err := recognize(message, 0, grammar, grammar[0])
		if err == nil && index == len(message) {
			//fmt.Println("match: ", message)
			count++
		} else {
			//fmt.Println("no match: ", message)
		}
	}
	return count
}

func Part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	parts := strings.Split(input, "\n\n")
	grammar := buildGrammar(strings.Split(parts[0], "\n"))

	messages := strings.Split(parts[1], "\n")
	count := 0
	for _, message := range messages {
		var index int
		var m, n int
		for {
			newIndex, err := recognize(message, index, grammar, grammar[42])
			if err != nil {
				break
			}
			m++
			index = newIndex
		}
		for {
			newIndex, err := recognize(message, index, grammar, grammar[31])
			if err != nil {
				break
			}
			n++
			index = newIndex
		}

		if m > n && n >= 1 && index == len(message) {
			//fmt.Println("match: ", message)
			count++
		} else {
			//fmt.Println("no match: ", message)
		}
	}
	return count

}

func main() {
	fmt.Println("--2020 day 19 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
