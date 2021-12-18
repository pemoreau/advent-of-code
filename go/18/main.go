package main

import (
	_ "embed"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
	"unicode"
)

//go:embed input.txt
var input_day string

type value struct {
	v     int
	depth int
}

type flattree []value

func (v value) String() string {
	if v.depth >= 0 {
		return fmt.Sprintf("%d(%d)", v.v, v.depth)
	}
	return fmt.Sprintf("%d", v.v)
}

func Parse(s string) flattree {
	res := flattree{}
	depth := 0
	i := 0
	for i < len(s) {
		switch {
		case s[i] == '[':
			depth += 1
			i += 1
		case s[i] == ']':
			depth -= 1
			i += 1
		case s[i] == ',':
			i += 1
		case unicode.IsDigit(rune(s[i])):
			v := int(s[i] - '0')
			j := i + 1
			for j < len(s) && unicode.IsDigit(rune(s[j])) {
				v = v*10 + int(s[j]-'0')
				j += 1
			}
			i = j
			res = append(res, value{v, depth})
		default:
			log.Fatalf("unexpected char %c", s[i])
		}
	}
	return res
}

func add(l1 flattree, l2 flattree) flattree {
	if len(l1) == 0 {
		return l2
	}
	if len(l2) == 0 {
		return l1
	}

	res := flattree{}
	res = append(res, l1...)
	res = append(res, l2...)
	for i, _ := range res {
		res[i].depth += 1
	}
	return res
}

func removeIndex(s flattree, i int) flattree {
	res := make(flattree, i, len(s)-1)
	copy(res, s[:i])
	return append(res, s[i+1:]...)
}

func replaceIndex(s flattree, i int, a value, b value) flattree {
	res := make(flattree, i, len(s)+1)
	copy(res, s[:i])
	res = append(res, a)
	res = append(res, b)
	return append(res, s[i+1:]...)
}

func explode(l flattree) (flattree, bool) {
	for i := 0; i < len(l)-1; i++ {
		if l[i].depth >= 5 && l[i].depth == l[i+1].depth {
			if i > 0 {
				l[i-1].v += l[i].v
			}
			if i < len(l)-2 {
				l[i+2].v += l[i+1].v
			}
			l[i].v = 0
			l[i].depth -= 1
			res := removeIndex(l, i+1)
			// fmt.Println("explode", res)
			return res, true
		}
	}
	return l, false
}

func split(l flattree) (flattree, bool) {
	for i := 0; i < len(l); i++ {
		if l[i].v >= 10 {
			a := l[i].v / 2
			b := l[i].v - a
			newDepth := l[i].depth + 1
			res := replaceIndex(l, i, value{a, newDepth}, value{b, newDepth})
			// fmt.Printf("split %d %v\n", l[i].v, res)
			return res, true
		}
	}
	return l, false
}

func normalize(l flattree) flattree {
	res := l
	reduced := true
	for reduced {
		res, reduced = explode(res)
		if !reduced {
			res, reduced = split(res)
		}
	}
	return res
}

func (s *stack) PushMagnitude(v value) {
	if s.IsEmpty() {
		s.Push(v)
		return
	}

	top, _ := s.Peek()
	if v.depth == top.depth {
		s.Pop()
		s.PushMagnitude(value{3*top.v + 2*v.v, v.depth - 1})
	} else {
		s.Push(v)
	}
}

func Magnitude(l flattree) int {
	stack := BuildStack()
	for _, v := range l {
		stack.PushMagnitude(v)

	}
	top, _ := stack.Pop()
	return top.v
}

func Part1(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")

	exp := flattree{}
	for _, l := range lines {
		exp = add(exp, Parse(l))
		exp = normalize(exp)
	}
	return Magnitude(exp)
}

func Part2(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	values := []flattree{}

	for _, l := range lines {
		values = append(values, normalize(Parse(l)))
	}

	max := 0
	for i, a := range values {
		for j, b := range values {
			if i != j {
				m := Magnitude(normalize(add(a, b)))
				if m > max {
					max = m
				}
				m = Magnitude(normalize(add(b, a)))
				if m > max {
					max = m
				}
			}
		}
	}
	return max

}

func main() {
	fmt.Println("--2021 day 18 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(string(input_day)))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(string(input_day)))
	fmt.Println(time.Since(start))
}

type stack flattree

func BuildStack() stack {
	return make([]value, 0)
}

func (s *stack) Push(c value) {
	*s = append(*s, c)
}

func (s *stack) Pop() (value, error) {
	l := len(*s)
	if l == 0 {
		return value{}, errors.New("stack is empty")
	}
	top := (*s)[l-1]
	*s = (*s)[:l-1]
	return top, nil
}

func (s *stack) Peek() (value, error) {
	if s.IsEmpty() {
		return value{}, errors.New("stack is empty")
	}
	return (*s)[len(*s)-1], nil
}

func (s *stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *stack) String() string {
	return fmt.Sprintf("%v", *s)
}
