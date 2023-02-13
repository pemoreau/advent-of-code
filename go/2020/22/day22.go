package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"strconv"
	"strings"
	"time"
)

//go:embed input_test.txt
var input_day string

type Game struct {
	deck1  []int
	deck2  []int
	past1  utils.Set[int]
	past2  utils.Set[int]
	winner int
}

func endOfGame(g Game) bool {
	return len(g.deck1) == 0 || len(g.deck2) == 0 || g.winner != 0
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
	if g.winner == 1 {
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

	game := Game{deck1, deck2, utils.NewSet[int](), utils.NewSet[int](), 0}
	for !endOfGame(game) {
		game = playRound(game)
	}
	if len(game.deck1) > 0 {
		game.winner = 1
	} else {
		game.winner = 2
	}

	return score(game)
}

func hash(deck []int) int {
	var h int
	for _, card := range deck {
		h = h*31 + card
	}
	return h
}

func playRound2(g Game) Game {
	h1 := hash(g.deck1)
	h2 := hash(g.deck2)
	// rule 1
	if g.past1.Contains(h1) && g.past2.Contains(h2) {
		g.winner = 1
		return g
	}
	g.past1.Add(h1)
	g.past2.Add(h2)

	card1 := g.deck1[0]
	card2 := g.deck2[0]
	g.deck1 = g.deck1[1:]
	g.deck2 = g.deck2[1:]

	// rule 2
	if len(g.deck1) >= card1 && len(g.deck2) >= card2 {
		newDeck1 := make([]int, card1)
		newDeck2 := make([]int, card2)
		copy(newDeck1, g.deck1[:card1])
		copy(newDeck2, g.deck2[:card2])
		newGame := Game{newDeck1, newDeck2, utils.NewSet[int](), utils.NewSet[int](), 0}
		for !endOfGame(newGame) {
			newGame = playRound2(newGame)
		}
		if newGame.winner == 1 {
			//g.winner = 1
			g.deck1 = append(g.deck1, card1, card2)
		} else if newGame.winner == 2 {
			//g.winner = 2
			g.deck2 = append(g.deck2, card2, card1)
		}
		return g
	}

	if card1 > card2 {
		g.deck1 = append(g.deck1, card1, card2)
	} else {
		g.deck2 = append(g.deck2, card2, card1)
	}
	return g
}

func Part2(input string) int {
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

	game := Game{deck1, deck2, utils.NewSet[int](), utils.NewSet[int](), 0}
	for !endOfGame(game) {
		game = playRound2(game)
	}
	return score(game)
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
