package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils/game2d"
	"slices"
	"time"
)

//go:embed input.txt
var inputDay string

type Cart struct {
	pos          game2d.Pos
	dir          uint8
	intersection uint8
	alive        bool
}

const (
	UP    = 0
	RIGHT = 1
	DOWN  = 2
	LEFT  = 3
)

var directions = []uint8{'^', '>', 'v', '<'}

func parseInput(input string) (*game2d.GridChar, []Cart) {
	var grid = game2d.BuildGridCharFromString(input)
	var carts []Cart
	for pos, v := range grid.All() {
		if v == '^' {
			carts = append(carts, Cart{pos, UP, 0, true})
			grid.SetPos(pos, '|')
		} else if v == 'v' {
			carts = append(carts, Cart{pos, DOWN, 0, true})
			grid.SetPos(pos, '|')
		} else if v == '<' {
			carts = append(carts, Cart{pos, LEFT, 0, true})
			grid.SetPos(pos, '-')
		} else if v == '>' {
			carts = append(carts, Cart{pos, RIGHT, 0, true})
			grid.SetPos(pos, '-')
		}
	}

	return grid, carts
}

func tick(grid *game2d.GridChar, cart Cart) Cart {
	switch cart.dir {
	case UP:
		cart.pos.Y--
	case RIGHT:
		cart.pos.X++
	case DOWN:
		cart.pos.Y++
	case LEFT:
		cart.pos.X--
	}
	tile, _ := grid.GetPos(cart.pos)
	if tile == '/' && cart.dir == UP {
		cart.dir = RIGHT
	} else if tile == '/' && cart.dir == LEFT {
		cart.dir = DOWN
	} else if tile == '/' && cart.dir == DOWN {
		cart.dir = LEFT
	} else if tile == '/' && cart.dir == RIGHT {
		cart.dir = UP
	} else if tile == '\\' && cart.dir == UP {
		cart.dir = LEFT
	} else if tile == '\\' && cart.dir == LEFT {
		cart.dir = UP
	} else if tile == '\\' && cart.dir == DOWN {
		cart.dir = RIGHT
	} else if tile == '\\' && cart.dir == RIGHT {
		cart.dir = DOWN
	} else if tile == '+' {
		if cart.intersection == 0 {
			cart.dir = (cart.dir + 3) % 4
		} else if cart.intersection == 2 {
			cart.dir = (cart.dir + 1) % 4
		}
		cart.intersection = (cart.intersection + 1) % 3
	}

	return cart
}

func collision(carts []Cart, i int) (bool, int) {
	for j := range carts {
		if i != j && carts[i].alive && carts[j].alive && carts[i].pos == carts[j].pos {
			return true, j
		}
	}
	return false, -1
}

func step(grid *game2d.GridChar, carts []Cart, part1 bool) (game2d.Pos, []Cart, bool) {
	slices.SortFunc(carts, func(i, j Cart) int {
		if i.pos.Y < j.pos.Y {
			return -1
		} else if i.pos.Y > j.pos.Y {
			return 1
		}
		return i.pos.X - j.pos.X
	})
	for i := range carts {
		if carts[i].alive {
			carts[i] = tick(grid, carts[i])
			if c, j := collision(carts, i); c {
				if part1 {
					return carts[i].pos, carts, false
				} else {
					carts[i].alive = false
					carts[j].alive = false
				}
			}
		}
	}
	// remove carts
	if !part1 {
		carts = slices.DeleteFunc(carts, func(c Cart) bool { return !c.alive })
	}
	return game2d.Pos{}, carts, true
}

func Part1(input string) string {
	var grid, carts = parseInput(input)
	p, carts, ok := step(grid, carts, true)
	for ok {
		p, carts, ok = step(grid, carts, true)
	}

	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

func Part2(input string) string {
	var grid, carts = parseInput(input)
	p, carts, ok := step(grid, carts, false)
	for ok && len(carts) > 1 {
		p, carts, ok = step(grid, carts, false)
	}
	p = carts[0].pos

	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

func main() {
	fmt.Println("--2018 day 13 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
