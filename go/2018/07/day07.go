package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

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
			graph[previous] = []uint8{}
		}
		graph[next] = append(graph[next], previous)
	}
	return graph
}

func Part1(input string) string {
	var graph = parseInput(input)
	var res = ""
	for len(graph) > 0 {
		var smallest uint8 = math.MaxUint8
		// select the smallest activable
		for k, previous := range graph {
			if len(previous) == 0 && k < smallest {
				smallest = k
			}
		}
		res += string(smallest)
		delete(graph, smallest)
		// remove the smallest from all previous
		for k, previous := range graph {
			if slices.Contains(previous, smallest) {
				index := slices.Index(previous, smallest)
				graph[k] = append(previous[:index], previous[index+1:]...)
			}
		}
	}

	return res
}

func duration(c uint8, offset int) int {
	return offset + int(c) - 'A' + 1
}

type Worker struct {
	number       int
	taskId       uint8
	start        int
	durationLeft int
}

func (w Worker) String() string {
	return fmt.Sprintf("worker #%d: task: %c start at: %d duration left: %d", w.number, w.taskId, w.start, w.durationLeft)
}

func Part2(input string) int {
	var graph = parseInput(input)

	//var offset, nbWorkers = 0, 2
	var offset, nbWorkers = 60, 5

	var time = 0
	var workers = make([]*Worker, nbWorkers)
	for i := range workers {
		workers[i] = &Worker{number: i}
	}

	for len(graph) > 0 {
		// while there are workers available
		for _, worker := range workers {
			if worker.durationLeft == 0 {
				//fmt.Printf("worker %d is available\n", worker.number)
				// select the taskId activable
				var taskId uint8 = math.MaxUint8
				for id, previous := range graph {
					if len(previous) == 0 && id < taskId {
						taskId = id
					}
				}
				if taskId == math.MaxUint8 {
					break
				}
				//fmt.Printf("worker %d is working on %c\n", worker.number, taskId)
				worker.taskId = taskId
				worker.start = time
				worker.durationLeft = duration(taskId, offset)
				delete(graph, taskId)
			}
		}

		// advance time
		time++
		//fmt.Printf("time: %d\n", time)
		for _, worker := range workers {
			if worker.durationLeft > 0 {
				worker.durationLeft--
				//fmt.Println(worker)
				if worker.durationLeft == 0 {
					//fmt.Printf("worker %d finished %c\n", worker.number, worker.taskId)
					// remove the task from all previous
					for k, previous := range graph {
						if slices.Contains(previous, worker.taskId) {
							index := slices.Index(previous, worker.taskId)
							graph[k] = append(previous[:index], previous[index+1:]...)
						}
					}
				}
			}
		}
	}
	var maxDuration = 0
	var workerWithMaxDurationLeft Worker
	for _, worker := range workers {
		if worker.durationLeft > maxDuration {
			maxDuration = worker.durationLeft
			workerWithMaxDurationLeft = *worker
		}
	}
	time += workerWithMaxDurationLeft.durationLeft
	return time
}

// too high 1126
func main() {
	fmt.Println("--2018 day 07 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
