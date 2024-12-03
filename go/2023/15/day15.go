package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"time"
)

//go:embed sample.txt
var inputTest string

//func hash(input string) int {
//	var res int
//	for _, c := range input {
//		res += int(c)
//		res *= 17
//		res &= 0xff // % 256
//	}
//	return res
//}

type Box []Lens

type Lens struct {
	name  string
	digit uint8
}

func index(b Box, name string) int {
	for i, l := range b {
		if l.name == name {
			return i
		}
	}
	return -1
}

func Part1(input string) int {
	// Slower version
	//input = strings.TrimSpace(input)
	//var parts = strings.Split(input, ",")
	//var res int
	//for _, part := range parts {
	//	res += hash(part)
	//}

	var res int
	var h int
	for _, c := range input {
		if c == ',' {
			res += h
			h = 0
			continue
		}
		h += int(c)
		h *= 17
		h &= 0xff // % 256
	}
	res += h
	return res
}

func Part2(input string) int {
	var boxes = make([]Box, 256)
	var h uint
	var start int
	var i int
	for i < len(input) {
		c := input[i]
		if c == ',' {
			h = 0
			i++
			start = i
		} else if c == '-' {
			box := boxes[h]
			name := input[start:i]
			if i := index(box, name); i != -1 {
				boxes[h] = append(box[:i], box[i+1:]...)
			}
			i++
		} else if c == '=' {
			box := boxes[h]
			name := input[start:i]
			i++
			digit := input[i]
			i++
			if i := index(box, name); i != -1 {
				boxes[h][i].digit = digit
			} else {
				boxes[h] = append(box, Lens{name, digit})
			}
		} else if c >= 'a' && c <= 'z' {
			h += uint(c)
			h *= 17
			h &= 0xff // % 256
			i++
		} else {
			panic("unexpected char")
		}
	}

	var res int
	for i, box := range boxes {
		for j, lens := range box {
			value := lens.digit - '0'
			res += (1 + i) * (j + 1) * int(value)
		}
	}
	return res
}

// Slower version
//func Part2(input string) int {
//	input = strings.TrimSpace(input)
//	var parts = strings.Split(input, ",")
//	var boxes = make([]Box, 256)
//	for _, part := range parts {
//		if part[len(part)-1] == '-' {
//			name := part[:len(part)-1]
//			h := hash(name)
//			box := boxes[h]
//			if i := index(box, name); i != -1 {
//				boxes[h] = append(box[:i], box[i+1:]...)
//			}
//		} else if name, number, found := strings.Cut(part, "="); found {
//			h := hash(name)
//			box := boxes[h]
//			if i := index(box, name); i != -1 {
//				boxes[h] = append(box[:i], Lens{name, number})
//				boxes[h] = append(boxes[h], box[i+1:]...)
//			} else {
//				boxes[h] = append(box, Lens{name, number})
//			}
//		}
//	}
//
//	var res int
//	for i, box := range boxes {
//		for j, lens := range box {
//			value, _ := strconv.Atoi(lens.number)
//			res += (1 + i) * (j + 1) * value
//		}
//	}
//
//	return res
//}

func main() {
	fmt.Println("--2023 day 15 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
