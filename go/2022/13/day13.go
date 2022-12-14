package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input_day string

const (
	VALUE = 0
	LIST  = 1
)

type node struct {
	id       int
	value    int
	children []node
}

func (n node) String() string {
	if n.id == VALUE {
		return strconv.Itoa(n.value)
	} else {
		return fmt.Sprintf("%v", n.children)
	}
}

func parseInt(s string, i int) (int, int) {
	res := 0
	for i < len(s) {
		if s[i] == ',' || s[i] == ']' || s[i] == '[' {
			return res, i
		}
		res = res*10 + int(s[i]-'0')
		i++
	}
	return res, i
}

func parseList(s string, i int) (node, int) {
	var current []node
	i++ // skip [
	for s[i] != ']' {
		if s[i] == '[' {
			n, j := parseList(s, i)
			current = append(current, n)
			i = j
		} else if s[i] == ',' {
			i++
		} else {
			v, j := parseInt(s, i)
			current = append(current, node{id: VALUE, value: v})
			i = j
		}
	}
	i++ // skip ]
	return node{id: LIST, children: current}, i
}

func equals(a, b node) bool {
	if a.id == VALUE && b.id == VALUE {
		return a.value == b.value
	}
	if a.id == LIST && b.id == LIST {
		if len(a.children) != len(b.children) {
			return false
		}
		for i := 0; i < len(a.children); i++ {
			if !equals(a.children[i], b.children[i]) {
				return false
			}
		}
		return true
	}
	return false
}

func compare(a, b node) int {
	if a.id == VALUE && b.id == VALUE {
		return b.value - a.value
	}
	if a.id == LIST && b.id == LIST {
		for i := 0; i < len(a.children) && i < len(b.children); i++ {
			res := compare(a.children[i], b.children[i])
			if res != 0 {
				return res
			}
		}
		return len(b.children) - len(a.children)
	}
	if a.id == VALUE && b.id == LIST {
		return compare(node{id: LIST, children: []node{a}}, b)
	}
	if a.id == LIST && b.id == VALUE {
		return compare(a, node{id: LIST, children: []node{b}})
	}
	return 0
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	parts := strings.Split(input, "\n\n")
	index := 1
	res := 0
	for _, part := range parts {
		lines := strings.Split(part, "\n")
		t1, _ := parseList(lines[0], 0)
		t2, _ := parseList(lines[1], 0)
		b := compare(t1, t2) > 0
		if b {
			res += index
		}
		index++
	}
	//lines := strings.Split(input, "\n")
	return res
}

func Part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	lines = append(lines, "[[2]]")
	lines = append(lines, "[[6]]")
	l := []node{}
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		t, _ := parseList(line, 0)
		l = append(l, t)
	}

	sort.Slice(l, func(i, j int) bool { return compare(l[i], l[j]) > 0 })
	//fmt.Println(l)
	i1, i2 := 0, 0
	//t1 := node{id: LIST, children: []node{node{id: LIST, children: []node{node{id: VALUE, value: 2}}}}}
	//t2 := node{id: LIST, children: []node{node{id: LIST, children: []node{node{id: VALUE, value: 6}}}}}
	t1, _ := parseList("[[2]]", 0)
	t2, _ := parseList("[[6]]", 0)
	for i := 0; i < len(l); i++ {
		if equals(l[i], t1) {
			i1 = i
		}
		if equals(l[i], t2) {
			i2 = i
		}
	}
	//t1, _ := parseList("[[2]]", 0)
	//t2, _ := parseList("[[6]]", 0)
	//i1 := sort.Search(len(l), func(i int) bool { return equals(l[i], t1) })
	//i2 := sort.Search(len(l), func(i int) bool { return equals(l[i], t2) })

	//lines := strings.Split(input, "\n")
	return (i1 + 1) * (i2 + 1)

}

func main() {
	fmt.Println("--2022 day 13 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(input_day))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(input_day))
	fmt.Println(time.Since(start))
}
