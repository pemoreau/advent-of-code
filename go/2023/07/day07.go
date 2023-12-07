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
		mult[index(h[i])]++
	}

	var orderedOccurrence []int
	for _, e := range mult {
		if e > 0 {
			orderedOccurrence = append(orderedOccurrence, e)
		}
	}
	slices.SortFunc(orderedOccurrence, func(a, b int) int { return b - a })
	//fmt.Println("orderedOccurrence", h, orderedOccurrence)

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

func (h hand) kind2() int {
	var mult [15]int
	for i := 0; i < len(h); i++ {
		mult[index(h[i])]++
	}
	var nbJoker = mult[index('J')]
	mult[index('J')] = 0

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
	fmt.Println("orderedOccurrence2", h, orderedOccurrence)

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

func compareString(s1, s2 hand) int {
	if s1 == s2 {
		return 0
	}
	var e1, e2 string
	for i := 0; i < len(s1); i++ {
		e1 = e1 + string('A'+uint8(index(s1[i])))
		e2 = e2 + string('A'+uint8(index(s2[i])))
	}
	if e1 > e2 {
		return 1
	}
	return -1
}

func compareString2(s1, s2 hand) int {
	if s1 == s2 {
		return 0
	}
	var e1, e2 string
	for i := 0; i < len(s1); i++ {
		if s1[i] == 'J' {
			e1 = e1 + string('A')

		} else {
			e1 = e1 + string('A'+uint8(index(s1[i])))
		}
		if s2[i] == 'J' {
			e2 = e2 + string('A')

		} else {
			e2 = e2 + string('A'+uint8(index(s2[i])))
		}
	}
	if e1 > e2 {
		return 1
	}
	return -1
}

func compare(h1, h2 hand) int {
	if h1 == h2 {
		return 0
	}

	k1 := h1.kind()
	k2 := h2.kind()

	if k1 > k2 {
		return 1
	}
	if k1 < k2 {
		return -1
	}

	return compareString(h1, h2)
}

func compare2(h1, h2 hand) int {
	if h1 == h2 {
		return 0
	}

	k1 := h1.kind2()
	k2 := h2.kind2()

	if k1 > k2 {
		return 1
	}
	if k1 < k2 {
		return -1
	}

	return compareString2(h1, h2)
}
func Part1(input string) int {
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

	//h1 := newHand("99999")
	//fmt.Println(h1, h1.toInteger())
	//h2 := newHand("KKKKK")
	//fmt.Println(h2, h2.toInteger())
	//fmt.Println(compare(h1, h2))

	slices.SortFunc(hands, compare)
	for _, h := range hands {
		fmt.Println(h)
	}
	var res int
	for i, h := range hands {
		res += bid[h] * (i + 1)
	}
	return res
}

func Part2(input string) int {
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

	slices.SortFunc(hands, compare2)
	for _, h := range hands {
		fmt.Println(h)
	}
	var res int
	for i, h := range hands {
		res += bid[h] * (i + 1)
	}
	return res
}

// 253718982 too low
// 254083736
func main() {
	fmt.Println("--2023 day 06 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
