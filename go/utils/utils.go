package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	res := []string{}
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	return res
}

func ToInt(s string) int {
	result, err := strconv.Atoi(s)
	Check(err)
	return result
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadNumbers(filename string) []int {
	file, err := os.Open(filename)
	Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var numbers []int
	for scanner.Scan() {
		numbers = append(numbers, ToInt(scanner.Text()))
	}
	return numbers
}

func stringsToNumbers(inputs []string) []int {
	numbers := make([]int, 0, len(inputs))
	for _, input := range inputs {
		numbers = append(numbers, ToInt(input))
	}
	return numbers
}

func LinesToNumbers(input string) []int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	return stringsToNumbers(lines)
}
func CommaSeparatedToNumbers(input string) []int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), ",")
	return stringsToNumbers(lines)
}
