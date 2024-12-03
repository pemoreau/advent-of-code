package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"github.com/pemoreau/advent-of-code/go/utils/interval"
	"github.com/pemoreau/advent-of-code/go/utils/set"
	"strings"
	"time"
)

//go:embed sample.txt
var inputTest string

type sensor struct {
	x, y   int
	beacon *beacon
	dist   int
}

type beacon struct {
	x, y int
}

// ordered list of disjoint intervals
type FreeSpace struct {
	intervals []interval.Interval
}

func (fs *FreeSpace) String() string {
	res := ""
	//res += fmt.Sprintf("[%d] ", len(fs.intervals))
	for _, i := range fs.intervals {
		res += fmt.Sprintf("%v ", i)
	}
	return res
}

func (fs *FreeSpace) Add(ab interval.Interval) {
	// https://coderbyte.com/algorithm/insert-interval-into-list-of-sorted-disjoint-intervals
	var newSet []interval.Interval
	var endSet []interval.Interval
	i := 0
	// add intervals that come before the new interval
	for i < len(fs.intervals) && fs.intervals[i].Max < ab.Min {
		newSet = append(newSet, fs.intervals[i])
		i++
	}

	// add our new interval to this final list
	newSet = append(newSet, ab)

	// check each interval that comes after the new interval to determine if we can merge
	// if no merges are required then populate a list of the remaining intervals
	for i < len(fs.intervals) {
		var last = newSet[len(newSet)-1]
		if fs.intervals[i].Min < last.Max {
			newInterval := interval.Interval{min(last.Min, fs.intervals[i].Min), max(last.Max, fs.intervals[i].Max)}
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
	var newSet []interval.Interval
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

func (fs *FreeSpace) Intersect(ab interval.Interval) {
	var newSet []interval.Interval
	for _, i := range fs.intervals {
		if i.Max < ab.Min {
			continue
		}
		if i.Min > ab.Max {
			break
		}
		newSet = append(newSet, interval.Interval{max(i.Min, ab.Min), min(i.Max, ab.Max)})
	}
	fs.intervals = newSet
}

func (fs *FreeSpace) Cardinality() int {
	res := 0
	for _, i := range fs.intervals {
		res += i.Max - i.Min + 1
	}
	return res
}

func getLine(sensors []sensor, ty int) FreeSpace {
	line := FreeSpace{}
	for _, s := range sensors {
		r := s.dist - utils.Abs(s.y-ty)
		if r > 0 {
			line.Add(interval.Interval{s.x - r, s.x + r})
		}
	}
	line.Merge()
	return line
}

func manhattanDistance(xa, ya, xb, yb int) int {
	return utils.Abs(xa-xb) + utils.Abs(ya-yb)
}

func parse(input string) []sensor {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	var sensors []sensor
	for _, line := range lines {
		var xs, ys, xb, yb int
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &xs, &ys, &xb, &yb)
		b := beacon{xb, yb}
		s := sensor{xs, ys, &b, manhattanDistance(xs, ys, xb, yb)}
		sensors = append(sensors, s)
	}
	return sensors
}

func Part1(input string) int {
	sensors := parse(input)
	ty := 2000000
	beacons := set.NewSet[beacon]()
	for _, s := range sensors {
		if s.beacon.y == ty {
			beacons.Add(*s.beacon)
		}
	}
	line := getLine(sensors, ty)
	return line.Cardinality() - len(beacons)
}

func Part2(input string) int {
	sensors := parse(input)
	MAX := 4000000

	//for ty := 0; ty < MAX+1; ty++ {
	//	line := getLine(sensors, ty)
	//	line.Intersect(utils.Interval{0, MAX})
	//	if len(line.intervals) > 1 {
	//		tx := line.intervals[0].Max + 1
	//		return 4000000*tx + ty
	//	}
	//}
	N := 100
	STEP := MAX / N
	messages := make(chan int)

	for i := 0; i < MAX; i = i + STEP {
		go func(messages chan int, min, max int) {
			//fmt.Println("Starting", min, max)
			for ty := min; ty < max+1; ty++ {
				line := getLine(sensors, ty)
				line.Intersect(interval.Interval{0, MAX})
				if len(line.intervals) > 1 {
					tx := line.intervals[0].Max + 1
					//fmt.Println("Found", tx, ty)
					messages <- 4000000*tx + ty
				}
			}
		}(messages, i, (i+STEP)-1)
	}

	res := <-messages
	return res
}

func main() {
	fmt.Println("--2022 day 15 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
