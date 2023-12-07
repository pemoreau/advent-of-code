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

func (h hand) kind() int {
	var mult [15]int
	for i := 0; i < len(h); i++ {
		mult[index(h[i], true)]++
	}

	var orderedOccurrence []int
	for _, e := range mult {
		if e > 0 {
			orderedOccurrence = append(orderedOccurrence, e)
		}
	}
	slices.SortFunc(orderedOccurrence, func(a, b int) int { return b - a })

	if orderedOccurrence[0] == 5 {
		return STRAIGHT
	}
	if orderedOccurrence[0] == 4 {
		return FOUROFAKIND
	}
	if orderedOccurrence[0] == 3 && orderedOccurrence[1] == 2 {
		return FULLHOUSE
	}
	if orderedOccurrence[0] == 3 {
		return THREEOFAKIND
	}
	if orderedOccurrence[0] == 2 && orderedOccurrence[1] == 2 {
		return TWOPAIRS
	}
	if orderedOccurrence[0] == 2 {
		return PAIR
	}
	return HIGHCARD
}

// with joker
func (h hand) kind2() int {
	var mult [15]int
	var indexJ = index('J', false)
	for i := 0; i < len(h); i++ {
		mult[index(h[i], false)]++
	}
	var nbJoker = mult[indexJ]
	mult[indexJ] = 0

	var orderedOccurrence []int
	for _, e := range mult {
		if e > 0 {
			orderedOccurrence = append(orderedOccurrence, e)
		}
	}
	slices.SortFunc(orderedOccurrence, func(a, b int) int { return b - a })
	if len(orderedOccurrence) == 0 {
		orderedOccurrence = append(orderedOccurrence, nbJoker)
	} else {
		orderedOccurrence[0] += nbJoker
	}

	if orderedOccurrence[0] == 5 {
		return STRAIGHT
	}
	if orderedOccurrence[0] == 4 {
		return FOUROFAKIND
	}
	if orderedOccurrence[0] == 3 && orderedOccurrence[1] == 2 {
		return FULLHOUSE
	}
	if orderedOccurrence[0] == 3 {
		return THREEOFAKIND
	}
	if orderedOccurrence[0] == 2 && orderedOccurrence[1] == 2 {
		return TWOPAIRS
	}
	if orderedOccurrence[0] == 2 {
		return PAIR
	}
	return HIGHCARD
}

func index(c byte, withJoker bool) int {
	if c >= '2' && c <= '9' {
		return int(c - '0')
	}
	switch c {
	case 'T':
		return 10
	case 'J':
		if withJoker {
			return 1
		} else {
			return 11
		}
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
		e1 = e1 + string('A'+uint8(index(s1[i], withJoker)))
		e2 = e2 + string('A'+uint8(index(s2[i], withJoker)))
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
	var k1, k2 int
	if withJoker {
		k1 = h1.kind2()
		k2 = h2.kind2()
	} else {
		k1 = h1.kind()
		k2 = h2.kind()
	}

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
