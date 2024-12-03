package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"github.com/pemoreau/advent-of-code/go/utils/interval"
	"slices"
	"strings"
	"time"
)

//go:embed sample.txt
var inputTest string

type brick struct {
	x, y, z    int
	nx, ny, nz int
}

type pos struct {
	x, y, z int
}

//func (b brick) cubes() set.Set[pos] {
//	var cubes = set.NewSet[pos]()
//	for x := b.x; x < b.x+b.nx; x++ {
//		for y := b.y; y < b.y+b.ny; y++ {
//			for z := b.z; z < b.z+b.nz; z++ {
//				cubes.Add(pos{x, y, z})
//			}
//		}
//	}
//	return cubes
//}

//func (b brick) bottom() set.Set[pos] {
//	var cubes = set.NewSet[pos]()
//	for x := b.x; x < b.x+b.nx; x++ {
//		for y := b.y; y < b.y+b.ny; y++ {
//			cubes.Add(pos{x, y, b.z})
//		}
//	}
//	return cubes
//}

//func (b brick) underBottom() set.Set[pos] {
//	var cubes = set.NewSet[pos]()
//	for x := b.x; x < b.x+b.nx; x++ {
//		for y := b.y; y < b.y+b.ny; y++ {
//			cubes.Add(pos{x, y, b.z - 1})
//		}
//	}
//	return cubes
//}
//
//func (b brick) top() set.Set[pos] {
//	var cubes = set.NewSet[pos]()
//	for x := b.x; x < b.x+b.nx; x++ {
//		for y := b.y; y < b.y+b.ny; y++ {
//			cubes.Add(pos{x, y, b.z + b.nz - 1})
//		}
//	}
//	return cubes
//}

func (b brick) sustainedBy(b2 brick) bool {
	if b.z-(b2.z+b2.nz) != 0 {
		return false
	}

	i1 := interval.Interval{b.x, b.x + b.nx - 1}
	i2 := interval.Interval{b2.x, b2.x + b2.nx - 1}
	i3 := interval.Interval{b.y, b.y + b.ny - 1}
	i4 := interval.Interval{b2.y, b2.y + b2.ny - 1}
	if i1.Inter(i2) != interval.Empty() && i3.Inter(i4) != interval.Empty() {
		return true
	}

	//under := b.underBottom()
	//top := b2.top()
	//for p := range top {
	//	if under.Contains(p) {
	//		return true
	//	}
	//}
	return false
}

func (b brick) drop(bricks []brick) brick {
	if b.z == 1 {
		return b
	}
	for _, b2 := range bricks {
		if b.sustainedBy(b2) {
			return b
		}
	}
	b.z -= 1
	return b.drop(bricks)
}

func dropBricks(bricks []brick) ([]brick, int) {
	var dropped []brick
	var n int
	for _, b := range bricks {
		newBrick := b.drop(dropped)
		if newBrick != b {
			n++
		}
		dropped = append(dropped, newBrick)
	}
	return dropped, n
}

func nbFall(bricks []brick, i int, b brick) int {
	var bricks2 []brick
	for _, b2 := range bricks {
		if b2 != b {
			bricks2 = append(bricks2, b2)
		}
	}

	//var bricks2 = make([]brick, i, len(bricks)-1)
	//copy(bricks2, bricks[:i])
	//bricks2 = append(bricks2, bricks[i+1:]...)

	_, n := dropBricks(bricks2)
	return n
}

func solve(input string, part2 bool) int {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")
	var bricks []brick
	for _, line := range lines {
		var x, y, z, a, b, c int
		fmt.Sscanf(line, "%d,%d,%d~%d,%d,%d", &x, &y, &z, &a, &b, &c)
		bb := brick{x, y, z, a - x + 1, b - y + 1, c - z + 1}
		bricks = append(bricks, bb)
	}

	slices.SortFunc(bricks, func(a brick, b brick) int { return a.z - b.z })
	bricks, _ = dropBricks(bricks)
	slices.SortFunc(bricks, func(a brick, b brick) int { return a.z - b.z })

	var canDesintegrate int
	var cpt int
	for i, b := range bricks {
		nb := nbFall(bricks, i, b)
		//fmt.Println(b, can)
		if nb == 0 {
			canDesintegrate++
		}
		cpt += nb
	}
	if part2 {
		return cpt
	}

	return canDesintegrate
}

func Part1(input string) int {
	return solve(input, false)
}

func Part2(input string) int {
	return solve(input, true)
}

func main() {
	fmt.Println("--2023 day 22 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
