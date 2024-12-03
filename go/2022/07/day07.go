package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"sort"
	"strings"
	"time"
)

//go:embed sample.txt
var inputTest string

func createFS(input string) fs {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	fs := NewFS()
	for _, line := range lines {
		if strings.HasPrefix(line, "$ ") {
			command := strings.TrimPrefix(line, "$ ")
			if strings.HasPrefix(command, "cd ") {
				dir := strings.TrimPrefix(command, "cd ")
				fs.cd(dir)
			} else {
				// ls
			}
		} else {
			if strings.HasPrefix(line, "dir ") {
				dir := strings.TrimPrefix(line, "dir ")
				fs.mkdir(dir)
			} else {
				var size int
				var name string
				fmt.Sscanf(line, "%d %s", &size, &name)
				fs.mkfile(name, size)
			}
		}
	}
	return fs
}

func Part1(input string) int {
	fs := createFS(input)
	res := 0
	for _, n := range fs.root.search(100000) {
		res += n.size()
	}
	return res
}

func Part2(input string) int {
	fs := createFS(input)

	totalSize := fs.root.size()
	maxSize := 70000000
	needed := 30000000
	unused := maxSize - totalSize
	toDelete := needed - unused

	sizes := fs.root.dirsizes()
	sort.Ints(sizes)
	i := sort.Search(len(sizes), func(i int) bool { return sizes[i] >= toDelete })
	return sizes[i]
}

func main() {
	fmt.Println("--2022 day 07 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
