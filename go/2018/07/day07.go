package main

import (
	"cmp"
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"math"
	"slices"
	"strings"
	"time"
)

//go:embed sample.txt
var inputTest string

type Graph map[uint8][]uint8

func (g Graph) String() string {
	var res = ""
	for k, v := range g {
		res += fmt.Sprintf("%c: %c\n", k, v)
	}
	return res
}

func parseInput(input string) Graph {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	var graph = make(Graph)
	for _, line := range lines {
		// Step <previous> must be finished before step <next> can begin.
		var previous, next uint8
		fmt.Sscanf(line, "Step %c must be finished before step %c can begin.", &previous, &next)
		if _, ok := graph[previous]; !ok {
			graph[previous] = make([]uint8, 0)
		}
		graph[next] = append(graph[next], previous)
	}
	return graph
}

func selectNextTask(graph Graph) uint8 {
	var smallest uint8 = math.MaxUint8
	// select the smallest activable
	for k, previous := range graph {
		if len(previous) == 0 && k < smallest {
			smallest = k
		}
	}
	return smallest
}

func removeTaskFromPrevious(graph Graph, task uint8) {
	for k, previous := range graph {
		if slices.Contains(previous, task) {
			index := slices.Index(previous, task)
			graph[k] = append(previous[:index], previous[index+1:]...)
		}
	}
}

func Part1(input string) string {
	var graph = parseInput(input)
	var res = ""
	for len(graph) > 0 {
		var task = selectNextTask(graph)
		res += string(task)
		delete(graph, task)
		removeTaskFromPrevious(graph, task)
	}
	return res
}

func duration(c uint8, offset int) int {
	return offset + int(c) - 'A' + 1
}

type Worker struct {
	number       int
	taskId       uint8
	durationLeft int
}

func (w Worker) String() string {
	return fmt.Sprintf("worker #%d: task: %c duration left: %d", w.number, w.taskId, w.durationLeft)
}

func Part2(input string) int {
	var graph = parseInput(input)

	var offset, nbWorkers = 60, 5
	var clock = 0
	var available, active []*Worker
	for i := range nbWorkers {
		available = append(available, &Worker{number: i})
	}

	for len(graph) > 0 {
		// assign tasks to available workers
		for task := selectNextTask(graph); len(available) > 0 && task < math.MaxUint8; task = selectNextTask(graph) {
			var worker = available[0]
			worker.taskId = task
			worker.durationLeft = duration(worker.taskId, offset)
			available = available[1:]
			active = append(active, worker)
			delete(graph, worker.taskId)
		}

		clock++
		var newActive []*Worker
		for _, worker := range active {
			worker.durationLeft--
			if worker.durationLeft == 0 { // worker finished
				removeTaskFromPrevious(graph, worker.taskId)
				available = append(available, worker)
			} else {
				newActive = append(newActive, worker)
			}
		}
		active = newActive
	}

	clock += slices.MaxFunc(active, func(a, b *Worker) int { return cmp.Compare(a.durationLeft, b.durationLeft) }).durationLeft
	return clock
}

func main() {
	fmt.Println("--2018 day 07 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
