package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils/game2d"
	"github.com/pemoreau/advent-of-code/go/utils/set"
	"math"
	"time"
)

//go:embed sample.txt
var inputTest string

type Piece = set.Set[game2d.Pos]

func collectPiece(grid *game2d.GridChar, pos game2d.Pos) Piece {
	fmt.Println("collectPiece", pos)
	var char, _ = grid.GetPos(pos)
	var res = set.NewSet[game2d.Pos]()
	var visited = set.NewSet[game2d.Pos]()
	var todo = []game2d.Pos{pos}
	for len(todo) > 0 {
		var p = todo[0]
		todo = todo[1:]
		if visited.Contains(p) {
			continue
		}
		visited.Add(p)
		if c, ok := grid.GetPos(p); ok && c == char {
			res.Add(p)
			for n := range p.Neighbors4() {
				if !res.Contains(n) {
					todo = append(todo, n)
				}
			}
		}
	}
	return res
}

func extractPieces(grid *game2d.GridChar) map[uint8][]Piece {
	var res = make(map[uint8][]Piece)
	var visited = set.NewSet[game2d.Pos]()
	for pos, c := range grid.All() {
		if visited.Contains(pos) {
			continue
		}
		var piece = collectPiece(grid, pos)
		visited = visited.Union(piece)
		fmt.Println("add piece:", c)
		res[c] = append(res[c], piece)
	}
	return res
}

func areaPerimeter(piece Piece) (int, int) {
	var area = len(piece)
	var perimeter int
	for p := range piece {
		for n := range p.Neighbors4() {
			if !piece.Contains(n) {
				perimeter++
			}
		}
	}
	return area, perimeter
}

func nextPos(p game2d.Pos, dir int) game2d.Pos {
	switch dir {
	case 0:
		return p.N()
	case 1:
		return p.E()
	case 2:
		return p.S()
	case 3:
		return p.W()
	}
	return p
}

//func freemanCode(piece Piece, start game2d.Pos) []int {
//	var res []int
//	var visited = set.NewSet[game2d.Pos]()
//	var p = start
//	var dir = 0
//	for {
//		if visited.Contains(p) {
//			break
//		}
//		visited.Add(p)
//		for d := 0; d < 4; d++ {
//			var np = nextPos(p, (dir+d)%4)
//			if piece.Contains(np) {
//				dir = (dir + d) % 4
//				res = append(res, dir)
//				p = np
//				break
//			}
//		}
//	}
//	return res
//}

//O --> x
//|                                   3
//v                                 2 P 0
//y                                   1
//O --> x
//|                                   0
//v                                 3 P 1
//y                                   2

const FORME = 255

func effectuer_suivi_contours(piece Piece) {
	var image = map[game2d.Pos]int{}
	var minx, maxx, miny, maxy = math.MaxInt, math.MinInt, math.MaxInt, math.MinInt
	for p := range piece {
		image[p] = FORME
		minx = min(minx, p.X)
		maxx = max(maxx, p.X)
		miny = min(miny, p.Y)
		maxy = max(maxy, p.Y)
	}

	var num_contours = 1
	for y := miny; y <= maxy; y++ {
		for x := minx; x <= maxx; x++ {
			p := game2d.Pos{x, y}
			if v, ok := image[p]; ok && v == FORME {
				dir := -1
				if !piece.Contains(p.N()) {
					dir = 0
				} else if !piece.Contains(p.E()) {
					dir = 1
				} else if !piece.Contains(p.S()) {
					dir = 2
				} else if !piece.Contains(p.W()) {
					dir = 3
				}
				if dir != -1 {
					suivre_un_contour(image, p, dir, num_contours)
					num_contours++
				}
			}
		}
	}
}

func suivre_un_contour(image map[game2d.Pos]int, A game2d.Pos, dirA int, num_contour int) {
	fmt.Printf("suivre_un_contour %v dirA: %d num_contour: %d\n", A, dirA, num_contour)
	// 1) recherche de la direction d'arrivée sur xA,yA :
	//    on tourne autour de xA,yA dans le sens croissant à partir de dirA
	var dir_finale int
	for i := range 4 {
		var d = (dirA + i) % 4
		if v, ok := image[nextPos(A, d)]; !ok || v > 0 {
			dir_finale = (d + 2) % 4
			break
		}
	}

	// 2) suivi et marquage du contour
	fmt.Printf("dir_finale: %d\n", dir_finale)
	var p = A
	var dir = dir_finale
	var continuer = true
	for continuer {
		fmt.Println(DisplayImage(image, true, p, dir))
		image[p] = num_contour
		//dir = (dir + 4 - 1) % 4
		found := false
		for i := 0; i < 4 && !found; i++ {
			d := (dir + 4 - i) % 4
			Q := nextPos(p, d)
			if v, ok := image[Q]; ok && v > 0 {
				p = Q
				dir = d
				fmt.Println("freeman: ", d)
				found = true
			}
		}
		if !found {
			return
		}
		continuer = p != A || dir != dir_finale
	}

}

func Part1(input string) int {
	var grid = game2d.BuildGridCharFromString(input)
	var allpieces = extractPieces(grid)

	var res int
	for c, pieces := range allpieces {
		fmt.Println("pieces pour", c, ":", len(pieces))
		for _, p := range pieces {
			var area, perimeter = areaPerimeter(p)
			fmt.Println("area:", area, "perimeter:", perimeter)
			res += area * perimeter
		}

	}

	return res
}

const (
	N, E, S, W = 0, 1, 2, 3
)

// returns false if is marked or not a new face
func markedFaceLeft(piece Piece, p game2d.Pos, dir int, visited []set.Set[game2d.Pos]) int {
	fmt.Printf("markedFaceLeft %v dir: %d\n", p, dir)
	var dirp = []game2d.Pos{p.N(), p.E(), p.S(), p.W()}
	if visited[dir].Contains(p) {
		//fmt.Printf("visited p %v\n", p)
		fmt.Println("return 0")
		return 0
	}
	visited[dir].Add(p)
	if !piece.Contains(p) {
		//fmt.Printf("not in piece\n")
		fmt.Println("return 0")
		return 0
	}
	if visited[dir].Contains(dirp[dir]) {
		//fmt.Printf("visited dirp %v\n", dirp[dir])
		fmt.Println("return 1")
		return 1
	}
	visited[dir].Add(dirp[dir])
	if piece.Contains(dirp[dir]) {
		//fmt.Printf("dirp[dir] in piece\n")
		fmt.Println("return 1")
		return 1
	}
	vleft := markedFaceLeft(piece, dirp[(dir+3)%4], dir, visited)
	//vright := borderLength(piece, dirp[(dir+1)%4], dir, visited)
	fmt.Printf("return %d\n", vleft)
	return vleft
}

func borderLength(piece Piece, p game2d.Pos, dir int, visited []set.Set[game2d.Pos]) int {
	fmt.Printf("borderLength %v dir: %d\n", p, dir)
	var dirp = []game2d.Pos{p.N(), p.E(), p.S(), p.W()}
	if visited[dir].Contains(p) {
		//fmt.Printf("visited p %v\n", p)
		fmt.Println("return 0")
		return 0
	}
	visited[dir].Add(p)

	if !piece.Contains(p) {
		//fmt.Printf("not in piece\n")
		fmt.Println("return 0")
		return 0
	}
	if piece.Contains(dirp[dir]) {
		//fmt.Printf("dirp[dir] in piece\n")
		fmt.Println("return 1")
		return 0
	}

	if visited[dir].Contains(dirp[dir]) {
		//fmt.Printf("visited dirp %v\n", dirp[dir])
		fmt.Println("return 1")
		return 0
	}
	visited[dir].Add(dirp[dir])

	vleft := borderLength(piece, dirp[(dir+3)%4], dir, visited)
	vright := borderLength(piece, dirp[(dir+1)%4], dir, visited)
	//fmt.Printf("vleft:%d vright:%d\n", vleft, vright)
	fmt.Printf("return %d\n", vright+vleft)
	return 1 + vright + vleft
}

func nbFaces(piece Piece) int {

	var res int
	var visited = make([]set.Set[game2d.Pos], 4)
	for i := 0; i < 4; i++ {
		visited[i] = set.NewSet[game2d.Pos]()
	}
	for d := 0; d < 4; d++ {
		for p := range piece {
			bl := borderLength(piece, p, d, visited)
			res += bl - 1
		}
	}

	return res
}

func DisplayPiece(piece Piece, c uint8, position bool, pos game2d.Pos, dir int) string {
	var res string
	var minx = math.MaxInt
	var maxx = math.MinInt
	var miny = math.MaxInt
	var maxy = math.MinInt
	for p := range piece {
		minx = min(minx, p.X)
		maxx = max(maxx, p.X)
		miny = min(miny, p.Y)
		maxy = max(maxy, p.Y)
	}
	for y := miny; y <= maxy; y++ {
		for x := minx; x <= maxx; x++ {
			if position && pos.X == x && pos.Y == y {
				res += string("^>v<"[dir])
			} else if piece.Contains(game2d.Pos{x, y}) {
				res += string(c)
			} else {
				res += "."
			}
		}
		res += "\n"
	}
	return res
}

func DisplayImage(image map[game2d.Pos]int, position bool, pos game2d.Pos, dir int) string {
	var res string
	var minx = math.MaxInt
	var maxx = math.MinInt
	var miny = math.MaxInt
	var maxy = math.MinInt
	for p := range image {
		minx = min(minx, p.X)
		maxx = max(maxx, p.X)
		miny = min(miny, p.Y)
		maxy = max(maxy, p.Y)
	}
	for y := miny; y <= maxy; y++ {
		for x := minx; x <= maxx; x++ {
			if position && pos.X == x && pos.Y == y {
				res += string("^>v<"[dir])

			} else if v, ok := image[game2d.Pos{x, y}]; ok {
				if v == FORME {
					res += "X"
				} else {
					res += "_"
				}
			} else {
				res += "."
			}
		}
		res += "\n"
	}
	return res
}

func Part2(input string) int {
	var grid = game2d.BuildGridCharFromString(input)
	var allpieces = extractPieces(grid)

	var res int
	for c, pieces := range allpieces {
		for _, p := range pieces {
			//var area, perimeter = areaPerimeter(p)
			//var nbFaces = perimeter - nbFaces(p)
			//fmt.Printf("nb faces pour %c: %d\n", c, nbFaces)
			//res += area * nbFaces
			fmt.Println(DisplayPiece(p, c, false, game2d.Pos{0, 0}, 0))
			fmt.Printf("effectuer_suivi_contours pour %c\n", c)
			effectuer_suivi_contours(p)
		}

	}

	return res
}

func main() {
	fmt.Println("--2024 day 12 solution--")
	//var inputDay = utils.Input()
	var inputDay = inputTest
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
