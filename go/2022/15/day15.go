package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"math"
	"strings"
	"time"
)

//go:embed input.txt
var input_day string

type sensor struct {
	x, y   int
	beacon *beacon
	dist   int
}

type beacon struct {
	x, y int
}

func manhattanDistance(xa, ya, xb, yb int) int {
	return utils.Abs(xa-xb) + utils.Abs(ya-yb)
}

func free(ty int, sensors []sensor, beacons utils.Set[beacon], min, max int) int {
	candidates := make([]sensor, 0)

	xmin := math.MaxInt
	xmax := math.MinInt
	for _, s := range sensors {
		if utils.Abs(s.y-ty) <= s.dist {
			candidates = append(candidates, s)
			xmin = utils.Min(xmin, s.x-s.dist)
			xmax = utils.Max(xmax, s.x+s.dist)
		}
	}
	//fmt.Println("sensors: ", len(sensors))
	//fmt.Println("candidates: ", len(candidates))
	//fmt.Println("xmin: ", xmin)
	//fmt.Println("xmax: ", xmax)

	res := 0
	for b := range beacons {
		if b.y == ty {
			res--
			//fmt.Printf("beacon at y=%d: %v\n", ty, b)
		}
	}
	for x := utils.Max(xmin, min); x <= utils.Min(xmax, max); x++ {
		for _, s := range candidates {
			d := manhattanDistance(x, ty, s.x, s.y)
			if d <= s.dist {
				res++
				break
			}
		}
	}
	return res
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	sensors := make([]sensor, 0)
	beacons := utils.BuildSet[beacon]()
	for _, line := range lines {
		var xs, ys, xb, yb int
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &xs, &ys, &xb, &yb)
		b := beacon{xb, yb}
		s := sensor{xs, ys, &b, manhattanDistance(xs, ys, xb, yb)}
		sensors = append(sensors, s)
		beacons.Add(b)
	}
	xmin := math.MinInt
	xmax := math.MaxInt
	return free(10, sensors, beacons, xmin, xmax)
	//return free(2000000)

}

type FreeSpace struct {
	intervals []utils.Interval
}

func (fs *FreeSpace) Add(i utils.Interval) {
	if len(fs.intervals) == 0 {
		fs.intervals = append(fs.intervals, i)
	}
	for j, iv := range fs.intervals {

	}
}

func occupied(ty int, sensors []sensor, beacons utils.Set[beacon], min, max int) int {
	line := &FreeSpace{}
	line.Add(utils.Interval{min, max})

	return 0
}

func Part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	sensors := make([]sensor, 0)
	beacons := utils.BuildSet[beacon]()
	for _, line := range lines {
		var xs, ys, xb, yb int
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &xs, &ys, &xb, &yb)
		b := beacon{xb, yb}
		s := sensor{xs, ys, &b, manhattanDistance(xs, ys, xb, yb)}
		sensors = append(sensors, s)
		beacons.Add(b)
	}

	MIN := 0
	//MAX := 20
	MAX := 4000000
	for ty := MIN; ty < MAX+1; ty++ {
		if ty%10000 == 0 {
			fmt.Println(ty)
		}
		nb := 0
		for b := range beacons {
			if b.y == ty {
				nb++
				//fmt.Printf("beacon at y=%d: %v\n", ty, b)
			}
		}

		f := free(ty, sensors, beacons, MIN, MAX)
		if nb+f < (MAX - MIN + 1) {
			fmt.Printf("ty=%d: %d\n", ty, f)
			for tx := MIN; tx <= MAX; tx++ {
				marked := false
				for _, s := range sensors {
					d := manhattanDistance(tx, ty, s.x, s.y)
					if d <= s.dist {
						marked = true
						break
					}
				}
				if !marked {
					return 4000000*tx + ty
				}
			}
		}
	}

	//	fmt.Println("ty: ", ty)
	//	candidates := make([]sensor, 0)
	//	xmin := math.MaxInt
	//	xmax := math.MinInt
	//	for _, s := range sensors {
	//		if utils.Abs(s.y-ty) <= s.dist {
	//			candidates = append(candidates, s)
	//			xmin = utils.Min(xmin, s.x-s.dist)
	//			xmax = utils.Max(xmax, s.x+s.dist)
	//		}
	//	}
	//	fmt.Println("sensors: ", len(sensors))
	//	fmt.Println("candidates: ", len(candidates))
	//	fmt.Println("xmin: ", xmin)
	//	fmt.Println("xmax: ", xmax)
	//
	//	res := 0
	//	for b := range beacons {
	//		if b.y == ty {
	//			res--
	//			fmt.Println("beacon at y=20: ", b)
	//		}
	//	}
	//
	//	for tx := utils.Max(0, xmin); tx < utils.Min(21, xmax); tx++ {
	//		fmt.Println("tx: ", tx)
	//		for _, s := range candidates {
	//			d := manhattanDistance(tx, ty, s.x, s.y)
	//			if d <= s.dist {
	//				res := 4000000*tx + ty
	//				return res
	//			}
	//		}
	//	}
	//}
	return 0

}

func main() {
	fmt.Println("--2022 day 15 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(input_day))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(input_day))
	fmt.Println(time.Since(start))
}
