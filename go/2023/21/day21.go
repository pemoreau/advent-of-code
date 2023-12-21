package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils/game2d"
	"github.com/pemoreau/advent-of-code/go/utils/set"
	"time"
)

//go:embed input_test.txt
var inputDay string

func findStart(m game2d.MatrixChar) game2d.Pos {
	for y, line := range m {
		for x, c := range line {
			if c == 'S' {
				return game2d.Pos{x, y}
			}
		}
	}
	return game2d.Pos{-1, -1}
}

func displayGrid(m game2d.MatrixChar, elves set.Set[game2d.Pos]) {
	for y, line := range m {
		for x, c := range line {
			if elves.Contains(game2d.Pos{x, y}) {
				fmt.Print("O")
			} else {
				fmt.Print(string(c))
			}
		}
		fmt.Println()
	}
}

func Part1(input string) int {
	grid := game2d.BuildMatrixCharFromString(input)
	start := findStart(grid)

	elves := set.NewSet[game2d.Pos]()
	elves.Add(start)
	var reached = set.NewSet[game2d.Pos]()

	for i := 0; i < 64; i++ {
		for e := range elves {
			neighbours := e.Neighbors4()
			for _, n := range neighbours {
				if !grid.IsValidPos(n) {
					continue
				}
				if grid[n.Y][n.X] == '#' {
					continue
				}
				if reached.Contains(n) {
					continue
				}
				reached.Add(n)
			}
		}
		//displayGrid(grid, elves)
		//fmt.Println()
		tmp := elves
		elves = reached
		reached = tmp
		clear(reached)
	}
	return len(elves)
}

func mod(a, b int) int {
	m := a % b
	if a < 0 && b < 0 {
		m -= b
	}
	if a < 0 && b > 0 {
		m += b
	}
	return m % b
}

func posModulo(p game2d.Pos, m game2d.MatrixChar) game2d.Pos {
	lx := m.LenX()
	ly := m.LenY()
	return game2d.Pos{X: mod(p.X, lx), Y: mod(p.Y, ly)}
}

func Part2(input string) int {
	grid := game2d.BuildMatrixCharFromString(input)
	start := findStart(grid)

	elves := set.NewSet[game2d.Pos]()
	elves.Add(start)
	var reached = set.NewSet[game2d.Pos]()

	var values []int
	var delta1 []int
	var delta2 []int
	n := grid.LenX()
	for i := 0; i < 300; i++ {
		for e := range elves {
			for _, n := range e.Neighbors4() {
				m := posModulo(n, grid)
				if !grid.IsValidPos(m) {
					continue
				}
				if grid[m.Y][m.X] == '#' {
					continue
				}
				if reached.Contains(n) {
					continue
				}
				reached.Add(n)
			}
		}
		values = append(values, len(reached))
		fmt.Printf("values_%d=%d\t", len(values)-1, values[len(values)-1])

		delta1 = append(delta1, len(reached)-len(elves))
		fmt.Printf("delta1_%d=%d\t", len(delta1)-1, delta1[len(delta1)-1])

		if len(delta1) > n {
			last := delta1[len(delta1)-1]
			previous := delta1[len(delta1)-1-n]
			diff := last - previous
			//fmt.Printf("i=%d last=%d previous=%d diff=%d\n", i, last, previous, diff)
			delta2 = append(delta2, diff)
			//delta3 = append(delta3, delta[len(delta2)-1]-delta[len(delta2)-grid.LenX()-1])
		} else {
			delta2 = append(delta2, 0)
		}
		fmt.Printf("delta2_%d=%d\n", len(delta2)-1, delta2[len(delta2)-1])

		// cycle delta2 : 16 16 14 18 16 9 11 14 16 14 18
		// [10 13 8 13 10 13 14 19 13 9 14 22 17 17 26 18 17 11 19 13 13 12 12 19 17 20 16 17 14 18 17 8 11 15 17 15 19 16 16 14 18 16 9 11 14 16 14 18 16 16 14 18 16 9 11 14 16 14 18 16 16 14 18 16 9 11 14 16 14 18 16 16 14 18 16 9 11 14 16 14 18 16 16 14 18 16 9 11 14 16 14 18 16 16 14 18 16 9 11 14 16 14 18 16 16 14 18 16 9 11 14 16 14 18 16 16 14 18 16 9 11 14 16 14 18 16 16 14 18 16 9 11 14 16 14 18 16 16 14 18 16 9 11 14 16 14 18 16 16 14 18 16 9 11 14 16 14 18 16 16 14 18 16 9 11 14 16 14 18 16 16 14 18 16 9 11 14 16 14 18 16 16 14 18 16 9 11 14 16 14 18 16 16 14 18 16 9 11 14 16 14 18 16 16 14 18 16 9 11 14 16 14 18 16 16 14 18 16 9 11 14 16 14 18 16 16 14 18 16 9 11 14 16 14 18 16 16 14 18 16 9 11 14 16 14 18 16 16 14 18 16 9 11 14 16 14 18 16 16 14 18 16 9 11 14 16 14 18 16 16 14 18 16 9 11 14 16 14 18 16 16 14 18 16 9 11 14 16]

		//219 221 205 226 204 238 177 265 187 291 176 258 164 255 174 256 193 258 193 261 171 284 161 281 160 288 171 273 167 253 180 257 188 249 185 260 189 262 189 261 187 243 191 253 189 256 195 247 196 256 180 245 164 273 155 285 158 277 173 274 170 321 215 330 209 316 208 316 169 277 164 272 149 280 152 282 158 271 167 248 188 262 193 259 194 244 192 247 187 259 185 262 185 264 195 250 180 252 183 260 165 260 184 293 160 281 145 282 174 267 188 259 186 260 183 254 168 264 162 269 181 271 186 251 189 227 207 236 210 222 216

		//fmt.Println(i, len(reached))
		tmp := elves
		elves = reached
		reached = tmp
		clear(reached)
	}
	fmt.Println("values", values)
	fmt.Println("delta1", delta1)
	fmt.Println("delta2", delta2)
	N := 10 * n
	values = values[:N]
	delta1 = delta1[:N]
	delta2 = delta2[:N]
	//fmt.Println("delta1", delta1)
	//fmt.Println("delta2", delta2)

	value_i := values[len(values)-1]
	fmt.Printf("last values_i(%d)=%d\n", len(values)-1, value_i)

	LAST := 5000
	//for i := N + 1; i < N+10*n; i++ {
	for i := N + 1; i < LAST+1; i++ {
		delta2_i := delta2[i-n-1]
		delta1_i := delta1[i-n-1] + delta2_i
		value_i += delta1_i

		delta1 = append(delta1, delta1_i)
		delta2 = append(delta2, delta2_i)

		fmt.Printf("values_%d=%d\t", i-1, value_i)
		fmt.Printf("delta1_%d=%d\t", len(delta1)-1, delta1[len(delta1)-1])
		fmt.Printf("delta2_%d=%d\n", len(delta2)-1, delta2[len(delta2)-1])

		//values = append(values, values[len(values)-1]+delta_i)

	}

	return len(elves)
}

func main() {
	fmt.Println("--2023 day 21 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
