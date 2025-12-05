package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"

	"github.com/pemoreau/advent-of-code/go/utils"
	"github.com/pemoreau/advent-of-code/go/utils/interval"
	"github.com/pemoreau/advent-of-code/go/utils/set"
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

func getLine(sensors []sensor, ty int) interval.FreeSpace {
	line := interval.FreeSpace{}
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
				//if len(line.intervals) > 1 {
				if line.Len() > 1 {
					tx := line.Get(0).Max + 1
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
