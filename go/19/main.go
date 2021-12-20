package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed input.txt
var input_day string

//go:embed input_test.txt
var input_test string

type Pos struct{ x, y, z int }

func (p Pos) String() string {
	return fmt.Sprintf("%d,%d,%d", p.x, p.y, p.z)
}

func ParsePos(s string) Pos {
	var x, y, z int
	fmt.Sscanf(s, "%d,%d,%d", &x, &y, &z)
	return Pos{x, y, z}
}

func roll(p Pos) Pos {
	return Pos{p.x, p.z, -p.y}
}
func turn(p Pos) Pos {
	return Pos{-p.y, p.x, p.z}
}

func allRotations(p Pos) []Pos {
	res := make([]Pos, 0)
	for cycle := 0; cycle < 2; cycle++ {
		for step := 0; step < 3; step++ {
			p = roll(p)
			res = append(res, p)
			for i := 0; i < 3; i++ {
				p = turn(p)
				res = append(res, p)
			}
		}
		p = roll(turn(roll(p)))
	}
	return res
}

func translation(p Pos, x, y, z int) Pos {
	return Pos{p.x + x, p.y + y, p.z + z}
}

func intersection(a, b []Pos) []Pos {
	count := make(map[Pos]uint8)
	for _, p := range a {
		count[p]++
	}
	for _, q := range b {
		count[q]++
	}

	res := make([]Pos, 0)
	for p, c := range count {
		if c > 1 {
			res = append(res, p)
		}
	}
	return res
}

type Scanner struct {
	name     string
	view     []Pos
	rotation uint8
	distance []int
}

// each position on a new line
func (s Scanner) String() string {
	b := strings.Builder{}
	b.WriteString(s.name)
	b.WriteString("\n")
	for _, p := range s.view {
		b.WriteString(p.String())
		b.WriteString("\n")
	}
	b.WriteString("distances: ")
	b.WriteString(fmt.Sprintf("%v", s.distance))
	b.WriteString("\n")

	return b.String()
}

func ParseScanner(lines []string) Scanner {
	view := make([]Pos, 0)
	for i, line := range lines {
		if i > 0 {
			view = append(view, ParsePos(line))
		}
	}
	return Scanner{name: lines[0], view: view, rotation: 0, distance: computeDistance(view)}
}

func norme(x, y, z int) int {
	return x*x + y*y + z*z
}

func computeDistance(view []Pos) []int {
	res := make([]int, 0)
	for i, p := range view {
		for j := i + 1; j < len(view); j++ {
			q := view[j]
			res = append(res, norme(p.x-q.x, p.y-q.y, p.z-q.z))
		}
	}
	return res
}

func intersectionIntList(a, b []int) []int {
	count := make(map[int]uint8)
	for _, p := range a {
		count[p]++
	}
	for _, q := range b {
		count[q]++
	}

	res := make([]int, 0)
	for p, c := range count {
		if c > 1 {
			res = append(res, p)
		}
	}
	return res
}

func TranslationScanner(s Scanner, x, y, z int) Scanner {
	view := make([]Pos, 0)
	for _, p := range s.view {
		view = append(view, translation(p, x, y, z))
	}
	return Scanner{name: s.name, view: view, rotation: s.rotation, distance: s.distance}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func maxAbs(s Scanner) int {
	max := 0
	for _, p := range s.view {
		if Abs(p.x) > max {
			max = Abs(p.x)
		}
		if Abs(p.y) > max {
			max = Abs(p.y)
		}
		if Abs(p.z) > max {
			max = Abs(p.z)
		}
	}
	return max
}

func allX(s Scanner) []int {
	res := make([]int, 0)
	for _, p := range s.view {
		res = append(res, p.x)
	}
	return res
}
func allY(s Scanner) []int {
	res := make([]int, 0)
	for _, p := range s.view {
		res = append(res, p.y)
	}
	return res
}
func allZ(s Scanner) []int {
	res := make([]int, 0)
	for _, p := range s.view {
		res = append(res, p.z)
	}
	return res
}

func arrangementScanner(s Scanner) []Scanner {
	points := [][]Pos{}
	for _, p := range s.view {
		points = append(points, allRotations(p))
	}

	res := make([]Scanner, 0)
	for i := 0; i < len(points[0]); i++ { // 24 rotations
		view := make([]Pos, 0)
		for j := range points {
			view = append(view, points[j][i])
		}
		res = append(res, Scanner{name: s.name, view: view, rotation: uint8(i), distance: s.distance})
	}
	return res
}

func searchTranslation(s1, s2 Scanner) []Pos {
	for _, x := range allX(s2) {
		for _, y := range allY(s2) {
			for _, z := range allZ(s2) {
				i := searchCommon(s1, s2, x, y, z)
				if len(i) > 0 {
					return i
				}
			}
		}
	}
	return make([]Pos, 0)
}

func commonBeamer(s1, s2 Scanner) int {
	commonDistance := intersectionIntList(s1.distance, s2.distance)
	return len(commonDistance)
}

func searchCommon(s1, s2 Scanner, x, y, z int) []Pos {
	s1x := s1.view[0].x
	s1y := s1.view[0].y
	s1z := s1.view[0].z
	sign := []int{1, -1}
	// fmt.Println("try", x, y, z)
	as2 := arrangementScanner(s2)
	// fmt.Println(s1)
	// fmt.Println(as2)
	for _, os2 := range as2 {
		for _, sp := range sign {
			for _, sx := range sign {
				for _, sy := range sign {
					for _, sz := range sign {
						tos2 := TranslationScanner(os2, sp*(s1x+sx*x), sp*(s1y+sy*y), sp*(s1z+sz*z))
						i := intersection(s1.view, tos2.view)
						// fmt.Printf("len=%d %v\n", len(i), i)
						if len(i) >= 12 {
							fmt.Println("translation", sp*(s1x+sx*x), sp*(s1y+sy*y), sp*(s1z+sz*z))
							return i
						}
					}

				}
			}
		}
	}
	return make([]Pos, 0)
}

func Part1(input string) int {

	input = `--- scanner 0a ---
-1,-1,1
-2,-2,2
-3,-3,3
-2,-3,1
5,6,-4
8,0,7

--- scanner 0b ---
1,-1,1
2,-2,2
3,-3,3
2,-1,3
-5,4,-6
-8,-7,0

--- scanner 0c ---
-1,-1,-1
-2,-2,-2
-3,-3,-3
-1,-3,-2
4,6,5
-7,0,8

--- scanner 0d ---
1,1,-1
2,2,-2
3,3,-3
1,3,-2
-4,-6,5
7,0,8

--- scanner 0e ---
1,1,1
2,2,2
3,3,3
3,1,2
-6,-4,-5
0,7,-8`

	input = input_test
	input = strings.TrimSuffix(input, "\n")
	parts := strings.Split(input, "\n\n")
	scanners := make([]Scanner, 0)
	for _, part := range parts {
		lines := strings.Split(part, "\n")
		scanners = append(scanners, ParseScanner(lines))
		// fmt.Println(scanners[len(scanners)-1])
	}

	// fmt.Println(allRotations(Pos{3, 2, 1}))
	// fmt.Println(allRotations(Pos{-1, -1, 1}))
	// for _, p := range allRotations(Pos{8, 0, 7}) {
	// 	fmt.Println(p)
	// }
	// scanner := ParseScanner(lines)
	// fmt.Println(Scanner{view: arrangement(ParsePos("3,1,2"))})
	// fmt.Println(arrangementScanner(scanners[0]))

	// fmt.Printf("common %d,%d --> %t\n", 0, 1, commonBeamer(scanners[0], scanners[1]))
	// fmt.Printf("common %d,%d --> %v\n", 0, 1, searchCommon(scanners[0], scanners[1], 68, -1246, -43))
	// fmt.Printf("common %d,%d --> %v\n", 0, 1, searchTranslation(scanners[0], scanners[1]))
	for i, _ := range scanners {
		for j := i + 1; j < len(scanners); j++ {
			fmt.Printf("common %d,%d --> %d\n", i, j, commonBeamer(scanners[i], scanners[j]))
			// fmt.Printf("common %d,%d --> %v\n", i, j, searchTranslation(s, scanners[j]))
		}
	}

	return 0
}

func Part2(input string) int {
	//s := strings.TrimSuffix(input, "\n")
	return 0

}

func main() {
	fmt.Println("--2021 day 19 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(string(input_day)))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(string(input_day)))
	fmt.Println(time.Since(start))
}
