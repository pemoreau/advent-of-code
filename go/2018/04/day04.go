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

func computeAsleep(input string) map[int][]interval.Interval {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	slices.SortFunc(lines, func(a, b string) int {
		partsA := strings.Split(a, "]")
		partsB := strings.Split(b, "]")
		return strings.Compare(partsA[0], partsB[0])
	})

	var asleep = make(map[int][]interval.Interval)
	var guard = 0
	var fallAsleep = 0
	var wakeup = 0
	for _, line := range lines {
		var parts = strings.Split(line, " ")
		if strings.Contains(parts[3], "#") {
			fmt.Sscanf(parts[3], "#%d", &guard)
		} else if strings.Contains(parts[2], "falls") {
			// Handle falls asleep
			fmt.Sscanf(parts[1], "00:%d", &fallAsleep)
		} else if strings.Contains(parts[2], "wakes") {
			// Handle wakes up
			fmt.Sscanf(parts[1], "00:%d", &wakeup)
			asleep[guard] = append(asleep[guard], interval.Interval{Min: fallAsleep, Max: wakeup - 1})
		}
	}
	return asleep
}

func Part1(input string) int {
	var asleep = computeAsleep(input)
	var mostAsleepGuard = 0
	var mostAsleepTime = 0
	for guard, intervals := range asleep {
		var total = 0
		for _, interval := range intervals {
			total += interval.Len()
		}
		if total > mostAsleepTime {
			mostAsleepTime = total
			mostAsleepGuard = guard
		}
	}

	// most asleep minute
	var minutes = make([]int, 60)
	var mostAsleepMinute = 0
	for _, interval := range asleep[mostAsleepGuard] {
		for i := interval.Min; i <= interval.Max; i++ {
			minutes[i]++
			if minutes[i] > minutes[mostAsleepMinute] {
				mostAsleepMinute = i
			}
		}
	}

	return mostAsleepGuard * mostAsleepMinute
}

func Part2(input string) int {
	var asleep = computeAsleep(input)
	var mostAsleepGuard = 0
	var mostAsleepMinute = 0
	var mostAsleepTime = 0
	for guard, intervals := range asleep {
		var minutes = make([]int, 60)
		for _, interval := range intervals {
			for i := interval.Min; i <= interval.Max; i++ {
				minutes[i]++
				if minutes[i] > mostAsleepTime {
					mostAsleepTime = minutes[i]
					mostAsleepGuard = guard
					mostAsleepMinute = i
				}
			}
		}
	}
	return mostAsleepGuard * mostAsleepMinute
}

func main() {
	fmt.Println("--2018 day 04 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
