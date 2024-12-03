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

type hand struct {
	initial string
	kind    int
	encoded string
}

const (
	HIGHCARD = iota
	PAIR
	TWOPAIRS
	THREEOFAKIND
	FULLHOUSE
	FOUROFAKIND
	STRAIGHT
)

func newHand(s string, withJoker bool) hand {
	var res = hand{
		initial: s,
	}
	res.kind = res.computeKind(withJoker)
	res.encoded = encode(s, withJoker)
	return res
}

func (h hand) computeKind(withJoker bool) int {
	var mult = make([]int, 15)
	for i := 0; i < len(h.initial); i++ {
		mult[index(h.initial[i])]++
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

	var res int
	if mult[0] == 5 {
		res = STRAIGHT
	} else if mult[0] == 4 {
		res = FOUROFAKIND
	} else if mult[0] == 3 && mult[1] == 2 {
		res = FULLHOUSE
	} else if mult[0] == 3 {
		res = THREEOFAKIND
	} else if mult[0] == 2 && mult[1] == 2 {
		res = TWOPAIRS
	} else if mult[0] == 2 {
		res = PAIR
	} else {
		res = HIGHCARD
	}

	return res
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

func encode(s string, withJoker bool) string {
	var res string
	for i := 0; i < len(s); i++ {
		if s[i] == 'J' && withJoker {
			res = res + "A"
		} else {
			res = res + string('A'+uint8(index(s[i])))
		}
	}
	return res
}

func compare(h1, h2 hand) int {
	var k1 = h1.kind
	var k2 = h2.kind

	if k1 > k2 {
		return 1
	}
	if k1 < k2 {
		return -1
	}

	if h1.encoded == h2.encoded {
		return 0
	}
	if h1.encoded > h2.encoded {
		return 1
	}
	return -1
}

func solve(input string, withJoker bool) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	var hands = make([]hand, 0, len(lines))
	var bid = make(map[hand]int)
	for _, line := range lines {
		s := strings.Split(line, " ")
		hand := newHand(s[0], withJoker)
		hands = append(hands, hand)
		v, _ := strconv.Atoi(s[1])
		bid[hand] = v
	}

	slices.SortFunc(hands, compare)

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
