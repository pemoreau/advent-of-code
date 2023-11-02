package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
	"time"
)

//go:embed input.txt
var inputDay string

type Pos struct{ x, y int16 }
type Image map[Pos]bool

func step(img Image, convo string, defaultValue bool) Image {

	new_img := make(Image, len(img))
	for key := range img {
		neighbors := []Pos{{key.x - 1, key.y - 1},
			{key.x, key.y - 1},
			{key.x + 1, key.y - 1},
			{key.x - 1, key.y},
			key,
			{key.x + 1, key.y},
			{key.x - 1, key.y + 1},
			{key.x, key.y + 1},
			{key.x + 1, key.y + 1}}

		index := 0
		for _, p := range neighbors {
			v, found := img[p]
			index = index << 1
			if found && v {
				index += 1
			} else if !found && defaultValue {
				index += 1
			}
		}
		new_img[key] = (convo[index] == '#')
	}
	return new_img
}

func countPixel(img Image) int {
	count := 0
	for _, v := range img {
		if v {
			count++
		}
	}
	return count
}

func findMinMax(img Image) (minX, maxX, minY, maxY int16) {
	minX = math.MaxInt16
	maxX = math.MinInt16
	minY = math.MaxInt16
	maxY = math.MinInt16
	for key := range img {
		if key.x < minX {
			minX = key.x
		}
		if key.x > maxX {
			maxX = key.x
		}
		if key.y < minY {
			minY = key.y
		}
		if key.y > maxY {
			maxY = key.y
		}
	}
	return
}

func addExtraLayer(img Image, minX, maxX, minY, maxY int16, b bool) (newMinX, newMaxX, newMinY, newMaxY int16) {
	newMinX = minX - 1
	newMaxX = maxX + 1
	for x := newMinX; x <= newMaxX; x++ {
		img[Pos{x, minY - 1}] = b
		img[Pos{x, maxY + 1}] = b
	}
	newMinY = minY - 1
	newMaxY = maxY + 1
	for y := newMinY; y <= newMaxY; y++ {
		img[Pos{minX - 1, y}] = b
		img[Pos{maxX + 1, y}] = b
	}
	return
}

func solve(input string, n uint16) int {
	input = strings.TrimSuffix(input, "\n")
	parts := strings.Split(input, "\n\n")
	convo := parts[0]

	img := make(Image)
	for j, line := range strings.Split(parts[1], "\n") {
		for i, c := range strings.TrimSpace(line) {
			img[Pos{int16(i), int16(j)}] = (c == '#')
		}
	}

	minX, maxX, minY, maxY := findMinMax(img)
	minX, maxX, minY, maxY = addExtraLayer(img, minX, maxX, minY, maxY, false)

	var i uint16
	for i = 0; i < n; i++ {
		defaultValue := img[Pos{minX, minY}]
		minX, maxX, minY, maxY = addExtraLayer(img, minX, maxX, minY, maxY, defaultValue)
		// addExtraLayer(img, minX-1, maxX+1, minY-1, maxY+1, defaultValue)
		// addExtraLayer(img, minX-2, maxX+2, minY-2, maxY+2, defaultValue)
		img = step(img, convo, defaultValue)
	}

	return countPixel(img)
}

func Part1(input string) int {
	return solve(input, 2)
}

func Part2(input string) int {
	return solve(input, 50)
}

func main() {
	fmt.Println("--2021 day 20 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
