package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"math"
	"strconv"
	"strings"
	"time"
)

//go:embed sample.txt
var inputTest string

func hash(secret int) int {
	n := secret * 64
	n = n ^ secret
	n = n % 16777216
	secret = n
	n = n / 32
	n = n ^ secret
	n = n % 16777216
	secret = n
	n = n * 2048
	n = n ^ secret
	n = n % 16777216
	return n
}

func extractSequences(secret int) map[[4]int]int {
	var res = make(map[[4]int]int)
	var diff []int

	N := 2000
	for i := 0; i < N; i++ {
		var newSecret = hash(secret)
		diff = append(diff, (newSecret%10)-(secret%10))
		if len(diff) > 4 {
			diff = diff[1:]
		}
		if len(diff) == 4 {
			if _, ok := res[[4]int{diff[0], diff[1], diff[2], diff[3]}]; !ok {
				res[[4]int{diff[0], diff[1], diff[2], diff[3]}] = newSecret % 10
			}
		}
		secret = newSecret
	}
	return res
}

func Part1(input string) int {
	var lines = strings.Split(input, "\n")
	var res int
	var N = 2000
	for _, line := range lines {
		var secret, _ = strconv.Atoi(line)
		for range N {
			secret = hash(secret)
		}
		res += secret
	}
	return res
}

func Part2(input string) int {

	var secrets []int
	var sequences = make(map[[4]int]int)

	for _, line := range strings.Split(input, "\n") {
		var secret, _ = strconv.Atoi(line)
		secrets = append(secrets, secret)

		for seq, price := range extractSequences(secret) {
			sequences[seq] += price
		}
	}

	var res = math.MinInt
	for _, v := range sequences {
		if v > res {
			res = v
		}
	}

	return res
}

func main() {
	fmt.Println("--2024 day 22 solution--")
	var inputDay = utils.Input()
	//var inputDay = inputTest
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
