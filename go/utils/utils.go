package utils

import (
	"bufio"
	_ "embed"
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
	var res []string
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

func StringsToNumbers(inputs []string) []int {
	numbers := make([]int, 0, len(inputs))
	for _, input := range inputs {
		numbers = append(numbers, ToInt(input))
	}
	return numbers
}

func LinesToNumbers(input string) []int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	return StringsToNumbers(lines)
}
func CommaSeparatedToNumbers(input string) []int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), ",")
	return StringsToNumbers(lines)
}

// dedent supprime l'indentation commune des blocs multi-lignes pour plus de lisibilité.
func Dedent(s string) string {
	lines := strings.Split(strings.Trim(s, "\n"), "\n")
	// calcule l'indent min (ignore les lignes vides)
	min := -1
	for _, ln := range lines {
		if strings.TrimSpace(ln) == "" {
			continue
		}
		i := 0
		for i < len(ln) && ln[i] == ' ' {
			i++
		}
		if min == -1 || i < min {
			min = i
		}
	}
	if min > 0 {
		for i, ln := range lines {
			if len(ln) >= min {
				lines[i] = ln[min:]
			}
		}
	}
	return strings.Join(lines, "\n")
}
