package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/pemoreau/advent-of-code/go/utils"
)

func maxBank(bank string) int {
	var index int
	var value1 uint8
	for i := 0; i < len(bank)-1; i++ {
		var c = bank[i]
		if c > value1 {
			value1 = c
			index = i
		}
	}

	var value2 uint8
	for i := index + 1; i < len(bank); i++ {
		var c = bank[i]
		if c > value2 {
			value2 = c
			index = i
		}
	}

	return 10*int(value1-'0') + int(value2-'0')
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	var lines = strings.Split(input, "\n")

	var res int
	for _, l := range lines {
		res += maxBank(l)
	}
	return res
}

//func maxBank12(bank string) int {
//	var index1, index2, index3, index4, index5, index6, index7, index8, index9, index10, index11 int
//	var value1, value2, value3, value4, value5, value6, value7, value8, value9, value10, value11, value12 uint8
//	for i := 0; i < len(bank)-11; i++ {
//		var c = bank[i]
//		if c > value1 {
//			value1 = c
//			index1 = i
//		}
//	}
//	for i := index1 + 1; i < len(bank)-10; i++ {
//		var c = bank[i]
//		if c > value2 {
//			value2 = c
//			index2 = i
//		}
//	}
//	for i := index2 + 1; i < len(bank)-9; i++ {
//		var c = bank[i]
//		if c > value3 {
//			value3 = c
//			index3 = i
//		}
//	}
//	for i := index3 + 1; i < len(bank)-8; i++ {
//		var c = bank[i]
//		if c > value4 {
//			value4 = c
//			index4 = i
//		}
//	}
//	for i := index4 + 1; i < len(bank)-7; i++ {
//		var c = bank[i]
//		if c > value5 {
//			value5 = c
//			index5 = i
//		}
//	}
//	for i := index5 + 1; i < len(bank)-6; i++ {
//		var c = bank[i]
//		if c > value6 {
//			value6 = c
//			index6 = i
//		}
//	}
//	for i := index6 + 1; i < len(bank)-5; i++ {
//		var c = bank[i]
//		if c > value7 {
//			value7 = c
//			index7 = i
//		}
//	}
//	for i := index7 + 1; i < len(bank)-4; i++ {
//		var c = bank[i]
//		if c > value8 {
//			value8 = c
//			index8 = i
//		}
//	}
//	for i := index8 + 1; i < len(bank)-3; i++ {
//		var c = bank[i]
//		if c > value9 {
//			value9 = c
//			index9 = i
//		}
//	}
//	for i := index9 + 1; i < len(bank)-2; i++ {
//		var c = bank[i]
//		if c > value10 {
//			value10 = c
//			index10 = i
//		}
//	}
//	for i := index10 + 1; i < len(bank)-1; i++ {
//		var c = bank[i]
//		if c > value11 {
//			value11 = c
//			index11 = i
//		}
//	}
//	for i := index11 + 1; i < len(bank); i++ {
//		var c = bank[i]
//		if c > value12 {
//			value12 = c
//			//index12 = i
//		}
//	}
//
//	var res = 100000000000*int(value1-'0') + 10000000000*int(value2-'0') + 1000000000*int(value3-'0') + 100000000*int(value4-'0') + 10000000*int(value5-'0') + 1000000*int(value6-'0') + 100000*int(value7-'0') + 10000*int(value8-'0') + 1000*int(value9-'0') + 100*int(value10-'0') + 10*int(value11-'0') + int(value12-'0')
//	//fmt.Printf("max=%d\n", res)
//	return res
//}

func rmax(bank []uint8, digits int) int {
	if digits <= 0 {
		return 0
	}
	var maxValue uint8
	var index int
	for i := 0; i < len(bank); i++ {
		if bank[i] > maxValue {
			maxValue = bank[i]
			index = i
		}
	}
	maxValue -= '0' // from '0' to 0
	if digits == 1 {
		return int(maxValue)
	}
	var v = rmax(bank[index+1:len(bank)+1], digits-1)
	return int(maxValue)*int(math.Pow10(digits-1)) + v
}

func Part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	var lines = strings.Split(input, "\n")

	var res int
	for _, l := range lines {
		var bank = []uint8(l)
		res += rmax(bank[:len(bank)-11], 12)
	}
	return res
}

func main() {
	fmt.Println("--2025 day 03 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
