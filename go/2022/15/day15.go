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

// ordered list of disjoint intervals
type FreeSpace struct {
	intervals []utils.Interval
}

func (fs FreeSpace) String() string {
	res := ""
	//res += fmt.Sprintf("[%d] ", len(fs.intervals))
	for _, i := range fs.intervals {
		res += fmt.Sprintf("%v ", i)
	}
	return res
}

func (fs *FreeSpace) Add(interval utils.Interval) {
	// https://coderbyte.com/algorithm/insert-interval-into-list-of-sorted-disjoint-intervals
	newSet := make([]utils.Interval, 0)
	endSet := make([]utils.Interval, 0)
	i := 0
	// add intervals that come before the new interval
	for i < len(fs.intervals) && fs.intervals[i].Max < interval.Min {
		newSet = append(newSet, fs.intervals[i])
		i++
	}

	// add our new interval to this final list
	newSet = append(newSet, interval)

	// check each interval that comes after the new interval to determine if we can merge
	// if no merges are required then populate a list of the remaining intervals
	for i < len(fs.intervals) {
		var last = newSet[len(newSet)-1]
		if fs.intervals[i].Min < last.Max {
			newInterval := utils.Interval{utils.Min(last.Min, fs.intervals[i].Min), utils.Max(last.Max, fs.intervals[i].Max)}
			newSet[len(newSet)-1] = newInterval
		} else {
			endSet = append(endSet, fs.intervals[i])
		}
		i++
	}
	fs.intervals = append(newSet, endSet...)
}

func (fs *FreeSpace) Merge() {
	if len(fs.intervals) == 0 {
		return
	}
	newSet := make([]utils.Interval, 0)
	var last = fs.intervals[0]
	i := 1

	for i < len(fs.intervals) {
		if last.Max+1 >= fs.intervals[i].Min {
			last.Max = fs.intervals[i].Max
		} else {
			newSet = append(newSet, last)
			last = fs.intervals[i]
		}
		i++
	}
	newSet = append(newSet, last)
	fs.intervals = newSet
}

func (fs *FreeSpace) Intersect(interval utils.Interval) {
	newSet := make([]utils.Interval, 0)
	for _, i := range fs.intervals {
		if i.Max < interval.Min {
			continue
		}
		if i.Min > interval.Max {
			break
		}
		newSet = append(newSet, utils.Interval{utils.Max(i.Min, interval.Min), utils.Min(i.Max, interval.Max)})
	}
	fs.intervals = newSet
}

//func (fs *FreeSpace) Add(interval utils.Interval) {
//	// https://coderbyte.com/algorithm/insert-interval-into-list-of-sorted-disjoint-intervals
//	newSet := make([]utils.Interval, 0)
//	run := true
//	for _, v := range fs.intervals {
//		if run {
//			if (v.Min <= interval.Min && interval.Min <= v.Max) || interval.Min <= v.Min {
//				newSet = append(newSet, utils.Interval{utils.Min(v.Min, interval.Min), utils.Max(v.Max, interval.Max)})
//				run = false
//			} else {
//				newSet = append(newSet, v)
//			}
//		} else {
//			newSet = append(newSet, v)
//		}
//	}
//	if run {
//		newSet = append(newSet, interval)
//	}
//	fs.intervals = newSet
//}

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

	MAX := 4000000

	for ty := 0; ty < MAX+1; ty++ {
		line := FreeSpace{}
		for _, s := range sensors {
			r := s.dist - utils.Abs(s.y-ty)
			if r > 0 {
				line.Add(utils.Interval{s.x - r, s.x + r})
			}
		}
		line.Merge()
		line.Intersect(utils.Interval{0, MAX})
		if len(line.intervals) > 1 {
			tx := line.intervals[0].Max + 1
			fmt.Println(tx, ty, line)
			return 4000000*tx + ty
		}

	}
	//x = 2721114
	//y = 3367718
	//10884459367718

	//MIN := 0
	////MAX := 20
	//MAX := 4000000
	//for ty := MIN; ty < MAX+1; ty++ {
	//	if ty%10000 == 0 {
	//		fmt.Println(ty)
	//	}
	//	nb := 0
	//	for b := range beacons {
	//		if b.y == ty {
	//			nb++
	//			//fmt.Printf("beacon at y=%d: %v\n", ty, b)
	//		}
	//	}
	//
	//	f := free(ty, sensors, beacons, MIN, MAX)
	//	if nb+f < (MAX - MIN + 1) {
	//		fmt.Printf("ty=%d: %d\n", ty, f)
	//		for tx := MIN; tx <= MAX; tx++ {
	//			marked := false
	//			for _, s := range sensors {
	//				d := manhattanDistance(tx, ty, s.x, s.y)
	//				if d <= s.dist {
	//					marked = true
	//					break
	//				}
	//			}
	//			if !marked {
	//				return 4000000*tx + ty
	//			}
	//		}
	//	}
	//}

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
