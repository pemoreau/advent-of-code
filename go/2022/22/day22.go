package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"math"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input_day string

type Pos struct {
	X, Y int
}

type Topology struct {
	rockByY map[int]utils.Set[Pos]
	rockByX map[int]utils.Set[Pos]
	row     map[int]utils.Interval
	column  map[int]utils.Interval
	grid    map[Pos]uint8
}

type State struct {
	Pos
	Dir  int //0 for right (>), 1 for down (v), 2 for left (<), and 3 for up (^)
	Grid *Topology
}

func (s State) String() string {
	dir := []string{">", "v", "<", "^"}
	return fmt.Sprintf("%d,%d  %s", s.Pos.Y+1, s.Pos.X+1, dir[s.Dir])
}

func parse(input string) (topology *Topology, path string) {
	input = strings.TrimSuffix(input, "\n")
	parts := strings.Split(input, "\n\n")
	lines := strings.Split(parts[0], "\n")

	grid := make(map[Pos]uint8)
	rockByY := make(map[int]utils.Set[Pos])
	rockByX := make(map[int]utils.Set[Pos])
	row := make(map[int]utils.Interval)
	column := make(map[int]utils.Interval)

	for j, line := range lines {
		if _, ok := rockByX[j]; !ok {
			rockByX[j] = utils.BuildSet[Pos]()
		}
		ymin := math.MaxInt
		ymax := 0
		xmin := math.MaxInt
		xmax := 0
		for i, c := range line {
			if _, ok := rockByY[i]; !ok {
				rockByY[i] = utils.BuildSet[Pos]()
			}
			if _, ok := column[i]; !ok {
				column[i] = utils.Interval{math.MaxInt, 0}
			}
			if c == '#' || c == '.' {
				p := Pos{i, j}
				grid[p] = uint8(c)
				if c == '#' {
					rockByY[i].Add(p)
					rockByX[j].Add(p)
				}
				xmin = utils.Min(xmin, i)
				xmax = utils.Max(xmax, i)
				ymin = utils.Min(column[i].Min, j)
				ymax = utils.Max(column[i].Max, j)
				column[i] = utils.Interval{ymin, ymax}
			}
		}
		row[j] = utils.Interval{xmin, xmax}
	}

	return &Topology{rockByY, rockByX, row, column, grid}, parts[1]
}

func step(state *State, order string) {
	fmt.Println("order", order)
	switch order {
	case "L":
		state.Dir = (state.Dir + 3) % 4
	case "R":
		state.Dir = (state.Dir + 1) % 4
	default:
		nbSteps, _ := strconv.Atoi(order)
		for i := 0; i < nbSteps; i++ {
			X, Y := state.Pos.X, state.Pos.Y
			var nextX, nextY int
			switch state.Dir {
			case 0:
				nextX = X + 1
				nextY = Y
				if nextX > state.Grid.row[Y].Max {
					nextX = state.Grid.row[Y].Min
				}
			case 1:
				nextX = X
				nextY = Y + 1
				if nextY > state.Grid.column[X].Max {
					nextY = state.Grid.column[X].Min
				}
			case 2:
				nextX = X - 1
				nextY = Y
				if nextX < state.Grid.row[Y].Min {
					nextX = state.Grid.row[Y].Max
				}
			case 3:
				nextX = X
				nextY = Y - 1
				if nextY < state.Grid.column[X].Min {
					nextY = state.Grid.column[X].Max
				}
			}
			if c, ok := state.Grid.grid[Pos{nextX, nextY}]; ok && c == '.' {
				state.Pos.X = nextX
				state.Pos.Y = nextY
			} else {
				break
			}
		}

	}
}

func Part1(input string) int {
	topology, path := parse(input)
	state := State{Pos{topology.row[0].Min, 0}, 0, topology}
	fmt.Println(path)
	fmt.Println("row", topology.row)
	fmt.Println("column", topology.column)
	fmt.Println("state", state)

	path = strings.ReplaceAll(path, "L", " L ")
	path = strings.ReplaceAll(path, "R", " R ")
	orders := strings.Split(path, " ")
	for _, order := range orders {
		step(&state, order)
		fmt.Printf("state: %v\n", state)
	}

	res := 1000*(state.Pos.Y+1) + 4*(state.Pos.X+1) + state.Dir
	return res
}

type Cube struct {
	N    int
	grid [6]map[Pos]uint8
}

type State3D struct {
	N    int
	face int
	pos  Pos
	dir  int //0 for right (>), 1 for down (v), 2 for left (<), and 3 for up (^)
}

func step3D(s *State3D, cube *Cube, order string) {
	switch order {
	case "L":
		s.dir = (s.dir + 3) % 4
	case "R":
		s.dir = (s.dir + 1) % 4
	default:
		nbSteps, _ := strconv.Atoi(order)
		for i := 0; i < nbSteps; i++ {
			X, Y := s.pos.X, s.pos.Y
			var nextX, nextY int
			switch s.dir {
			case 0:
				nextX = X + 1
				nextY = Y
				if nextX >= cube.N {

				}
			}
			if c, ok := cube.grid[s.face][Pos{nextX, nextY}]; ok && c == '.' {
				s.pos.X = nextX
				s.pos.Y = nextY
			} else {
				break
			}
		}

	}
}

// 0 for right (>), 1 for down (v), 2 for left (<), and 3 for up (^)
func transition(s *State3D) {
	X, Y, N := s.pos.X, s.pos.Y, s.N

	switch s.face {
	case 0:
		switch s.dir {
		case 0:
			s.face = 5
			s.dir = 1
			s.pos = Pos{N - 1, N - 1 - Y}
		case 1:
			s.face = 3
			s.pos = Pos{X, 0}
		case 2:
			s.face = 2
			s.dir = 2
			s.pos = Pos{Y, 0}
		case 3:
			s.face = 1
			s.dir = 1
			s.pos = Pos{N - 1 - X, 0}

		}
	}

}

func Part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	parts := strings.Split(input, "\n\n")
	lines := strings.Split(parts[0], "\n")

	faces := [6]struct {
		X utils.Interval
		Y utils.Interval
	}{
		{X: utils.Interval{8, 11}, Y: utils.Interval{0, 3}},
		{X: utils.Interval{0, 3}, Y: utils.Interval{4, 7}},
		{X: utils.Interval{4, 7}, Y: utils.Interval{4, 7}},
		{X: utils.Interval{8, 11}, Y: utils.Interval{4, 7}},
		{X: utils.Interval{8, 11}, Y: utils.Interval{8, 11}},
		{X: utils.Interval{12, 15}, Y: utils.Interval{8, 11}},
	}

	N := 4
	cube := Cube{N: N}
	for i := 0; i < 6; i++ {
		cube.grid[i] = make(map[Pos]uint8)
	}
	for j, line := range lines {
		for i, c := range line {
			for k := 0; k < 6; k++ {
				if faces[k].X.Contains(i) && faces[k].Y.Contains(j) {
					cube.grid[k][Pos{i % N, j % N}] = uint8(c)
				}
			}
		}
	}

	return 0

}

func main() {
	fmt.Println("--2022 day 22 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(input_day))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(input_day))
	fmt.Println(time.Since(start))
}
