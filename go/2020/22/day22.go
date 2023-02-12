package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input_day string

type Game struct {
	deck1 []int
	deck2 []int
}

func endOfGame(g Game) bool {
	return len(g.deck1) == 0 || len(g.deck2) == 0
}

func playRound(g Game) Game {
	card1 := g.deck1[0]
	card2 := g.deck2[0]
	g.deck1 = g.deck1[1:]
	g.deck2 = g.deck2[1:]
	if card1 > card2 {
		g.deck1 = append(g.deck1, card1, card2)
	} else {
		g.deck2 = append(g.deck2, card2, card1)
	}
	return g
}

func score(g Game) int {
	var score int
	if len(g.deck1) > 0 {
		for i, card := range g.deck1 {
			score += card * (len(g.deck1) - i)
		}
	} else {
		for i, card := range g.deck2 {
			score += card * (len(g.deck2) - i)
		}
	}
	return score
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	parts := strings.Split(input, "\n\n")
	lines1 := strings.Split(parts[0], "\n")[1:]
	lines2 := strings.Split(parts[1], "\n")[1:]

	var deck1 []int
	var deck2 []int
	for _, line := range lines1 {
		num, _ := strconv.Atoi(line)
		deck1 = append(deck1, num)
	}
	for _, line := range lines2 {
		num, _ := strconv.Atoi(line)
		deck2 = append(deck2, num)
	}

	game := Game{deck1, deck2}
	for !endOfGame(game) {
		game = playRound(game)
	}
	return score(game)
}

func Part2(input string) int {
	// input = strings.TrimSuffix(input, "\n")
	// lines := strings.Split(input, "\n")
	return 0

}

func main() {
	fmt.Println("--2020 day 21 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(input_day))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(input_day))
	fmt.Println(time.Since(start))
}
