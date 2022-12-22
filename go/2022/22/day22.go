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

//go:embed input_test.txt
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

func (s State3D) String() string {
	dir := []string{">", "v", "<", "^"}
	return fmt.Sprintf("face %d -- %d,%d  %s", s.face+1, s.pos.Y+1, s.pos.X+1, dir[s.dir])
}

func (s *State3D) rotateRight() {
	s.dir = (s.dir + 1) % 4
	s.pos.X = s.N - s.pos.Y - 1
	s.pos.Y = s.pos.X
}

func (s *State3D) rotateLeft() {
	s.dir = (s.dir + 3) % 4
	s.pos.X = s.pos.Y
	s.pos.Y = s.N - s.pos.X - 1
}

func (s *State3D) rotate180() {
	s.dir = (s.dir + 2) % 4
	s.pos.X = s.N - s.pos.X - 1
	s.pos.Y = s.N - s.pos.Y - 1
}

func (s *State3D) move(cube *Cube, transitionTable [6][4]struct{ face, rot int }) bool {
	start := *s
	switch s.dir {
	case 0:
		s.pos.X++
	case 1:
		s.pos.Y++
	case 2:
		s.pos.X--
	case 3:
		s.pos.Y--
	}
	if s.pos.X >= 0 && s.pos.X < s.N && s.pos.Y >= 0 && s.pos.Y < s.N {
		if cube.grid[s.face][s.pos] == '#' {
			*s = start
			return false
		}
		return true
	}

	if s.pos.X < 0 {
		s.pos.X = s.N - 1
	}
	if s.pos.X >= s.N {
		s.pos.X = 0
	}
	if s.pos.Y < 0 {
		s.pos.Y = s.N - 1
	}
	if s.pos.Y >= s.N {
		s.pos.Y = 0
	}
	//fmt.Println("transition", s.face, s.dir, transitionTable[s.face][s.dir])
	s.face = transitionTable[start.face][start.dir].face
	switch transitionTable[start.face][start.dir].rot {
	case 90:
		fmt.Println("rotate 90")
		s.rotateLeft()
	case 180:
		fmt.Println("rotate 180")
		s.rotate180()
	case 270:
		fmt.Println("rotate 270")
		s.rotateRight()
	}
	if cube.grid[s.face][s.pos] == '#' {
		*s = start
		return false
	}
	return true
}

// 0 for right (>), 1 for down (v), 2 for left (<), and 3 for up (^)

type TransitionTable [6][4]struct{ face, rot int }

func step3D(s *State3D, cube *Cube, order string, transitionTable TransitionTable) {
	fmt.Println("order", order)
	switch order {
	case "L":
		s.dir = (s.dir + 3) % 4
	case "R":
		s.dir = (s.dir + 1) % 4
	default:
		nbSteps, _ := strconv.Atoi(order)
		for i := 0; i < nbSteps; i++ {
			before := *s
			ok := s.move(cube, transitionTable)
			if !ok {
				break
			}
			fmt.Println("step", i+1, "move", before, "to", *s)
		}

	}
}

func Part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	parts := strings.Split(input, "\n\n")
	lines := strings.Split(parts[0], "\n")
	path := parts[1]

	N := 4
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

	transition := TransitionTable{
		{{face: 6 - 1, rot: 180}, {face: 4 - 1, rot: 0}, {face: 3 - 1, rot: 90}, {face: 2 - 1, rot: 180}},
		{{face: 3 - 1, rot: 0}, {face: 5 - 1, rot: 180}, {face: 6 - 1, rot: 270}, {face: 1 - 1, rot: 180}},
		{{face: 4 - 1, rot: 0}, {face: 5 - 1, rot: 90}, {face: 2 - 1, rot: 0}, {face: 1 - 1, rot: 270}},
		{{face: 6 - 1, rot: 270}, {face: 5 - 1, rot: 0}, {face: 3 - 1, rot: 0}, {face: 1 - 1, rot: 0}},
		{{face: 6 - 1, rot: 0}, {face: 2 - 1, rot: 180}, {face: 3 - 1, rot: 270}, {face: 4 - 1, rot: 0}},
		{{face: 1 - 1, rot: 180}, {face: 2 - 1, rot: 90}, {face: 5 - 1, rot: 0}, {face: 4 - 1, rot: 90}},
	}

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
	path = strings.ReplaceAll(path, "L", " L ")
	path = strings.ReplaceAll(path, "R", " R ")
	orders := strings.Split(path, " ")

	state := State3D{N: N, face: 0, pos: Pos{0, 0}, dir: 0}
	fmt.Printf("start from state: %v\n", state)
	for _, order := range orders {
		step3D(&state, &cube, order, transition)
		fmt.Printf("state: %v\n", state)
	}
	res := 1000*(state.pos.Y+1) + 4*(state.pos.X+1) + state.dir

	return res

}

func main() {
	fmt.Println("--2022 day 22 solution--")
	start := time.Now()
	//fmt.Println("part1: ", Part1(input_day))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(input_day))
	fmt.Println(time.Since(start))
}
