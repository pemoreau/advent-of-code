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
	var res []Pos
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

type Scanner struct {
	name     string
	view     []Pos
	aligned  bool
	rotation uint8
	position Pos
	distance []int
}

// each position on a new line
func (s Scanner) String() string {
	b := strings.Builder{}
	b.WriteString(s.name)
	b.WriteString(fmt.Sprintf(" aligned %t", s.aligned))
	b.WriteString(fmt.Sprintf(" rotation %v", s.rotation))
	b.WriteString(fmt.Sprintf(" translation %v", s.position))
	b.WriteString("\n")
	for _, p := range s.view {
		b.WriteString(p.String())
		b.WriteString("\n")
	}
	return b.String()
}

func ParseScanner(lines []string) Scanner {
	var view []Pos
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
	var res []int
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

	var res []int
	for p, c := range count {
		if c > 1 {
			res = append(res, p)
		}
	}
	return res
}

func arrangementView(s []Pos) [][]Pos {
	points := [][]Pos{}
	for _, p := range s {
		points = append(points, allRotations(p))
	}

	var res [][]Pos
	for i := 0; i < len(points[0]); i++ { // 24 rotations
		var view []Pos
		for j := range points {
			view = append(view, points[j][i])
		}
		res = append(res, view)
	}
	return res
}

func sortView(view []Pos) {
	sort.Slice(view, func(i, j int) bool {
		p, q := view[i], view[j]
		switch {
		case p.x < q.x:
			return true
		case p.x > q.x:
			return false
		case p.y < q.y:
			return true
		case p.y > q.y:
			return false
		case p.z < q.z:
			return true
		case p.z > q.z:
			return false
		default:
			return false
		}
	})
}

func commonBeamer(s1, s2 Scanner) int {
	commonDistance := intersectionIntList(s1.distance, s2.distance)
	return len(commonDistance)
}

func extractCommonBeamer(s1, s2 Scanner) (b1 []Pos, b2 []Pos) {
	m1 := map[Pos]struct{}{}
	m2 := map[Pos]struct{}{}
	commonDistance := intersectionIntList(s1.distance, s2.distance)
	for _, d := range commonDistance {
		for i, p := range s1.view {
			for j := i + 1; j < len(s1.view); j++ {
				q := s1.view[j]
				n := norme(p.x-q.x, p.y-q.y, p.z-q.z)
				if d == n {
					m1[p] = struct{}{}
					m1[q] = struct{}{}
				}
			}
		}
		for i, p := range s2.view {
			for j := i + 1; j < len(s2.view); j++ {
				q := s2.view[j]
				n := norme(p.x-q.x, p.y-q.y, p.z-q.z)
				if d == n {
					m2[p] = struct{}{}
					m2[q] = struct{}{}
				}
			}
		}
	}
	for k := range m1 {
		b1 = append(b1, k)
	}
	for k := range m2 {
		b2 = append(b2, k)
	}
	return
}

func checkSame(v1, v2 []Pos) bool {
	for i := 1; i < len(v1)-1; i++ {
		d1 := Pos{
			v1[i].x - v1[i-1].x,
			v1[i].y - v1[i-1].y,
			v1[i].z - v1[i-1].z,
		}
		d2 := Pos{
			v2[i].x - v2[i-1].x,
			v2[i].y - v2[i-1].y,
			v2[i].z - v2[i-1].z,
		}
		if d1 != d2 {
			return false
		}
	}
	return true
}

func searchTranslation(s1, s2 Scanner) (Pos, uint8, bool) {
	v1, v2 := extractCommonBeamer(s1, s2)
	sortView(v1)
	allv2 := arrangementView(v2)
	for r, v2 := range allv2 {
		sortView(v2)
		if checkSame(v1, v2) {
			// we have found rotation r and translation t from s1 to s2
			t := Pos{
				v1[0].x - v2[0].x,
				v1[0].y - v2[0].y,
				v1[0].z - v2[0].z,
			}
			return t, uint8(r), true
		}
	}
	return Pos{}, 0, false
}

func solve(input string) (int, []Pos) {
	input = strings.TrimSuffix(input, "\n")
	parts := strings.Split(input, "\n\n")
	var scanners []Scanner
	for _, part := range parts {
		lines := strings.Split(part, "\n")
		scanners = append(scanners, ParseScanner(lines))
	}

	scanners[0].aligned = true
	beacons := map[Pos]bool{}
	for _, p := range scanners[0].view {
		beacons[p] = true
	}
	aligned := []Scanner{scanners[0]}
	for len(aligned) > 0 {
		s1 := &aligned[0]
		aligned = aligned[1:]
		for j := 1; j < len(scanners); j++ {
			s2 := &scanners[j]

			if !s2.aligned && commonBeamer(*s1, *s2) >= (12*11)/2 {
				// fmt.Printf("common %d --> %d\n", j, commonBeamer(*s1, *s2))
				if t, r, ok := searchTranslation(*s1, *s2); ok {
					s2.aligned = true
					s2.position = t
					s2.rotation = r
					s2.view = arrangementView(s2.view)[r]
					var alignedv2 []Pos
					for _, p2 := range s2.view {
						alignedv2 = append(alignedv2, Pos{p2.x + t.x, p2.y + t.y, p2.z + t.z})
					}
					s2.view = alignedv2
					aligned = append(aligned, *s2)
					// fmt.Println(s2)
					for _, p := range s2.view {
						beacons[p] = true
					}
				}
			}
		}
	}
	var res = make([]Pos, 0, len(scanners))
	for _, s := range scanners {
		res = append(res, s.position)
	}
	return len(beacons), res
}

func manhattanDistance(from, to Pos) int {
	absX := from.x - to.x
	if absX < 0 {
		absX = -absX
	}
	absY := from.y - to.y
	if absY < 0 {
		absY = -absY
	}
	absZ := from.z - to.z
	if absZ < 0 {
		absZ = -absZ
	}
	return absX + absY + absZ
}

func Part1(input string) int {
	nb, _ := solve(input)
	return nb
}

func Part2(input string) int {
	_, pos := solve(input)
	max := 0
	for i := 0; i < len(pos); i++ {
		for j := i + 1; j < len(pos); j++ {
			d := manhattanDistance(pos[i], pos[j])
			if d > max {
				max = d
			}
		}
	}
	return max
}

func main() {
	fmt.Println("--2021 day 19 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
