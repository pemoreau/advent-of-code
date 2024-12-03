package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"time"
)

//go:embed input.txt
var inputDay string

func transformSubjectNumber(subjectNumber, loopSize int) int {
	res := 1
	for i := 0; i < loopSize; i++ {
		res *= subjectNumber
		res %= 20201227
	}
	return res
}

func findLoopSize(subjectNumber, publicKey int) int {
	value := 1
	for i := 1; ; i++ {
		value *= subjectNumber
		value %= 20201227
		if value == publicKey {
			return i
		}
	}
}

func Part1(input string) int {
	keys := utils.LinesToNumbers(input)
	cardPublicKey, doorPublicKey := keys[0], keys[1]
	doorLoopSize := findLoopSize(7, doorPublicKey)
	cardLoopSize := findLoopSize(7, cardPublicKey)
	encryptionKey1 := transformSubjectNumber(doorPublicKey, cardLoopSize)
	encryptionKey2 := transformSubjectNumber(cardPublicKey, doorLoopSize)
	if encryptionKey1 != encryptionKey2 {
		panic("encryption keys don't match")
	}
	return encryptionKey1
}

func main() {
	fmt.Println("--2020 day 25 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))
}
