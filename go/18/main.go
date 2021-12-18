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

func (v value) String() string {
	if v.depth >= 0 {
		return fmt.Sprintf("%d(%d)", v.v, v.depth)
	}
	return fmt.Sprintf("%d", v.v)
}

func add(l1 []value, l2 []value) []value {
	if len(l1) == 0 {
		return l2
	}
	if len(l2) == 0 {
		return l1
	}
	res := append(l1, l2...)
	for i, _ := range res {
		res[i].depth += 1
	}
	return res
}

func buildPair(a, b int) []value {
	return []value{
		{a, 1},
		{b, 1},
	}
}

func Parse(s string) []value {
	res := []value{}
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

func removeIndex(s []value, index int) []value {
	ret := make([]value, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
func replaceIndex(s []value, index int, a value, b value) []value {
	ret := []value{}
	ret = append(ret, s[:index]...)
	ret = append(ret, a)
	ret = append(ret, b)
	return append(ret, s[index+1:]...)
}

func explode(l []value) ([]value, bool) {
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

func split(l []value) ([]value, bool) {
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

// func normalize(l []value) []value {
// 	for {
// 		var re, rs bool
// 		l, re = explode(l)
// 		l, rs = split(l)
// 		if !re && !rs {
// 			return l
// 		}
// 	}
// }

func normalize(l []value) []value {
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

func Magnitude(l []value) int {
	stack := BuildStack()
	for _, v := range l {
		stack.PushMagnitude(v)

	}
	top, _ := stack.Pop()
	return top.v
}

func Part1(input string) int {
	//s := strings.TrimSuffix(input, "\n")

	// 	input = `[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]
	// [7,[[[3,7],[4,3]],[[6,3],[8,8]]]]
	// [[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]
	// [[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]
	// [7,[5,[[3,8],[1,4]]]]
	// [[2,[2,2]],[8,[8,1]]]
	// [2,9]
	// [1,[[[9,3],9],[[9,0],[0,7]]]]
	// [[[5,[7,4]],7],1]
	// [[[[4,2],2],6],[8,7]]`

	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")

	exp := []value{}
	for _, l := range lines {
		// fmt.Println("add:", l)

		exp = add(exp, Parse(l))
		exp = normalize(exp)
		// fmt.Println("res add:", exp)
	}
	// fmt.Println("result:", exp)

	// fmt.Println("magnitude:", Magnitude(Parse("[9,1]")))         // 29
	// fmt.Println("magnitude:", Magnitude(Parse("[1,9]")))         // 21
	// fmt.Println("magnitude:", Magnitude(Parse("[[9,1],[1,9]]"))) // 129
	// fmt.Println("magnitude:", Magnitude(Parse("[[1,2],[[3,4],5]]"))) // 143
	// fmt.Println("magnitude:", Magnitude(Parse("[[[0,7],4],[[7,8],[6,0]]],[8,1]]")))
	// l := buildPair(1, 1)
	// l = add(l, buildPair(2, 2))
	// l = normalize(l)
	// l = add(l, buildPair(3, 3))
	// l = normalize(l)
	// l = add(l, buildPair(4, 4))
	// l = normalize(l)
	// l = add(l, buildPair(5, 5))
	// l = normalize(l)
	// l = add(l, buildPair(6, 6))
	// l = normalize(l)

	// fmt.Println("l", l)
	// fmt.Println("parse", parse("[[1,1],2]"))

	return Magnitude(exp)
}

func Part2(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	values := [][]value{}

	for _, l := range lines {
		values = append(values, normalize(Parse(l)))
	}

	max := 0
	for i, la := range lines {
		for j, lb := range lines {
			a := Parse(la)
			b := Parse(lb)
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

type stack []value

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

func (s *stack) Size() int {
	return len(*s)
}

func (s *stack) String() string {
	return fmt.Sprintf("%v", *s)
}
