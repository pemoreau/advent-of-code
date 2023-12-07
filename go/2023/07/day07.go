package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

type hand string

const (
	HIGHCARD = iota
	PAIR
	TWOPAIRS
	THREEOFAKIND
	FULLHOUSE
	FOUROFAKIND
	STRAIGHT
)

func (h hand) kind(withJoker bool) int {
	var mult = make([]int, 15)
	for i := 0; i < len(h); i++ {
		mult[index(h[i])]++
	}

	var indexJ = index('J')
	var nbJoker = mult[indexJ]
	if withJoker {
		mult[indexJ] = 0
	}

	slices.SortFunc(mult, func(a, b int) int { return b - a })

	if withJoker {
		mult[0] += nbJoker
	}

	if mult[0] == 5 {
		return STRAIGHT
	}
	if mult[0] == 4 {
		return FOUROFAKIND
	}
	if mult[0] == 3 && mult[1] == 2 {
		return FULLHOUSE
	}
	if mult[0] == 3 {
		return THREEOFAKIND
	}
	if mult[0] == 2 && mult[1] == 2 {
		return TWOPAIRS
	}
	if mult[0] == 2 {
		return PAIR
	}
	return HIGHCARD
}

func index(c byte) int {
	if c >= '2' && c <= '9' {
		return int(c - '0')
	}
	switch c {
	case 'T':
		return 10
	case 'J':
		return 11
	case 'Q':
		return 12
	case 'K':
		return 13
	case 'A':
		return 14
	}
	panic("invalid card")
}

func compareString(s1, s2 hand, withJoker bool) int {
	if s1 == s2 {
		return 0
	}
	var e1, e2 string
	for i := 0; i < len(s1); i++ {
		if s1[i] == 'J' && withJoker {
			e1 = e1 + "A"
		} else {
			e1 = e1 + string('A'+uint8(index(s1[i])))
		}
		if s2[i] == 'J' && withJoker {
			e2 = e2 + "A"
		} else {
			e2 = e2 + string('A'+uint8(index(s2[i])))
		}
	}
	if e1 > e2 {
		return 1
	}
	return -1
}

func compare(h1, h2 hand, withJoker bool) int {
	if h1 == h2 {
		return 0
	}

	var k1 = h1.kind(withJoker)
	var k2 = h2.kind(withJoker)
	if k1 > k2 {
		return 1
	}
	if k1 < k2 {
		return -1
	}

	return compareString(h1, h2, withJoker)
}

func solve(input string, withJoker bool) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	var hands []hand
	var bid = make(map[hand]int)
	for _, line := range lines {
		s := strings.Split(line, " ")
		hand := hand(s[0])
		hands = append(hands, hand)
		v, _ := strconv.Atoi(s[1])
		bid[hand] = v
	}

	slices.SortFunc(hands, func(a, b hand) int { return compare(a, b, withJoker) })
	//for _, h := range hands {
	//	fmt.Println(h)
	//}
	var res int
	for i, h := range hands {
		res += bid[h] * (i + 1)
	}
	return res
}

func Part1(input string) int {
	return solve(input, false)
}

func Part2(input string) int {
	return solve(input, true)
}

func main() {
	fmt.Println("--2023 day 07 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
