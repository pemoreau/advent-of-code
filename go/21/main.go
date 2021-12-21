package main

import (
	_ "embed"
	"fmt"
	"time"
)

//go:embed input.txt
var input_day string

type Player struct {
	space uint16
	score uint16
}

type Dice struct {
	value  uint8
	nbRoll uint16
}

func (p Player) String() string {
	return fmt.Sprintf("space=%d score=%d", p.space, p.score)
}

func (p *Player) move(value uint16) {
	p.space = 1 + ((p.space + value - 1) % 10)
	p.score += p.space
}

func (d *Dice) deterministicRoll(n uint8) (sum uint16) {
	var i uint8
	for i = 0; i < n; i++ {
		sum += uint16(d.value)
		d.nbRoll += 1
		d.value += 1
		if d.value > 100 {
			d.value = 1
		}
	}
	return
}

func Part1(input string) int {
	// input := strings.TrimSuffix(input, "\n")
	// lines := strings.Split(input, "\n")
	p1 := Player{space: 10}
	p2 := Player{space: 3}
	d := Dice{value: 1}

	for {
		m := d.deterministicRoll(3)
		p1.move(m)
		// fmt.Printf("p1 dice #%d=%d %s\n", d.nbRoll, m, p1)
		if p1.score >= 1000 {
			return int(d.nbRoll) * int(p2.score)
		}
		m = d.deterministicRoll(3)
		p2.move(m)
		// fmt.Printf("p2 dice #%d=%d %s\n", d.nbRoll, m, p2)
		if p2.score >= 1000 {
			return int(d.nbRoll) * int(p1.score)
		}
	}
}

func explore(roll, absoluteRoll, dice uint, score0, score1 uint, space0, space1 uint, win0, win1 uint, player uint8) (uint, uint) {
	fmt.Printf("player=%d roll=%d[%d] dice=%d score0=%d score1=%d space0=%d space1=%d\n", player, roll, absoluteRoll, dice, score0, score1, space0, space1)
	if player == 0 {
		space0 = 1 + ((space0 + dice - 1) % 10)
		score0 += space0
	} else if player == 1 {
		space1 = 1 + ((space1 + dice - 1) % 10)
		score1 += space1
	}
	fmt.Printf("***                       score0=%d score1=%d space0=%d space1=%d\n", score0, score1, space0, space1)

	if score0 >= 21 {
		fmt.Printf("player0 win: %d, %d\n", win0+1, win1)
		return win0 + 1, win1
	}
	if score1 >= 21 {
		fmt.Printf("player1 win: %d, %d\n", win0, win1+1)
		return win0, win1 + 1
	}

	roll = (roll + 1) % 3
	if roll == 0 {
		player = (player + 1) % 2
	}

	resw0, resw1 := win0, win1
	w0, w1 := explore(roll, absoluteRoll+1, 1, score0, score1, space0, space1, win0, win1, player)
	resw0 += w0
	resw1 += w1
	w0, w1 = explore(roll, absoluteRoll+1, 2, score0, score1, space0, space1, win0, win1, player)
	resw0 += w0
	resw1 += w1
	w0, w1 = explore(roll, absoluteRoll+1, 3, score0, score1, space0, space1, win0, win1, player)
	resw0 += w0
	resw1 += w1
	return resw0, resw1
}

func Part2(input string) int {
	// input := strings.TrimSuffix(input, "\n")
	// lines := strings.Split(input, "\n")
	var win0, win1 uint
	w0, w1 := explore(0, 0, 1, 0, 0, 4, 8, 0, 0, 0)
	win0 += w0
	win1 += w1
	w0, w1 = explore(0, 0, 2, 0, 0, 4, 8, 0, 0, 0)
	win0 += w0
	win1 += w1
	w0, w1 = explore(0, 0, 3, 0, 0, 4, 8, 0, 0, 0)
	win0 += w0
	win1 += w1

	if win0 > win1 {
		return int(win0)
	} else {
		return int(win1)
	}
}

func main() {
	fmt.Println("--2021 day 21 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(string(input_day)))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(string(input_day)))
	fmt.Println(time.Since(start))
}
