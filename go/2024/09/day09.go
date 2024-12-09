package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"strings"
	"time"
)

//go:embed sample.txt
var inputTest string

type file struct {
	id  int
	len int
	adr int
}

func (f file) String() string {
	if f.id == -1 {
		return fmt.Sprintf("{FREE, len: %d, adr: %d}", f.len, f.adr)
	}
	return fmt.Sprintf("{#%d, len: %d, adr: %d}", f.id, f.len, f.adr)
}

type block = []*file

//type block struct {
//	files []file
//}
//
//func (b block) String() string {
//	var res strings.Builder
//	res.WriteString("[")
//	for _, f := range b.files {
//		res.WriteString(f.String())
//	}
//	res.WriteString("]")
//	return res.String()
//}

func parse(input string) []block {
	var blocks = make([]block, 0, len(input))
	var adr, id int
	for i, c := range input {
		var d = int(c - '0')
		var f file
		if i%2 == 0 {
			f = file{id, d, adr}
			id++
		} else {
			f = file{-1, d, adr}
		}
		var b = block{&f}
		blocks = append(blocks, b)
		adr += d
	}
	return blocks
}

func findNextFree(blocks []block, indexFree int) int {
	if indexFree < 0 || indexFree >= len(blocks) {
		panic("indexFree out of range")
	}
	for i := indexFree; i < len(blocks); i++ {
		var b = blocks[i]
		if len(b) > 0 && b[len(b)-1].id == -1 {
			return i
		}
	}
	panic("no free block")
}

//func compact(blocks []block) bool {
//	for lastIndex := len(blocks) - 1; lastIndex >= 1; lastIndex-- {
//		if len(blocks[lastIndex]) == 0 || blocks[lastIndex][0].id == -1 {
//			continue
//		}
//		var f = blocks[lastIndex][0]
//		var ff = blocks[lastIndex-1][len(blocks[lastIndex-1])-1]
//		if ff.id == -1 || ff.id != f.id {
//			return false
//		}
//		ff.len += f.len
//		f.id = -1
//
//	}
//
//}

func defragFile(blocks []block, indexFree int, f *file) bool {
	if indexFree < 0 || indexFree >= len(blocks) {
		return false
	}
	var freeBlock = blocks[indexFree]
	var freeFile = freeBlock[len(freeBlock)-1]
	if freeFile.id != -1 {
		fmt.Println("freeFile.id != -1")
		return false
	}

	if f.len > freeFile.len {
		// fill freeFile and split f
		freeFile.id = f.id
		f.len -= freeFile.len
		indexFree = findNextFree(blocks, indexFree)
		return true
	} else if f.len == freeFile.len {
		// fill freeFile and make f free
		freeFile.id = f.id
		f.id = -1
		indexFree = findNextFree(blocks, indexFree)
		return true
	} else if f.len < freeFile.len {
		// fill freeFile, split freeFile and make f free
		var newFile = &file{f.id, f.len, freeFile.adr}
		freeFile.len -= f.len
		freeFile.adr += f.len
		f.id = -1
		// insert f in freeBlock before last position
		freeBlock = append(freeBlock[:len(freeBlock)-1], newFile, freeBlock[len(freeBlock)-1])
		blocks[indexFree] = freeBlock
		return true
	}

	return false

}

func checksum(blocks []block) int {
	var res int
	var adr int
	for _, b := range blocks {
		for _, f := range b {
			if f.id != -1 {
				for i := 0; i < f.len; i++ {
					res += (adr * f.id)
					adr++
				}
			}
		}
	}
	return res
}

func Part1(input string) int {
	input = strings.TrimSpace(input)
	var blocks = parse(input)

	fmt.Println(blocks)

	var lastIndex = len(blocks) - 1
	if lastIndex%2 != 0 {
		lastIndex--
	}

	var freeIndex = 1
	var ok = true
	for ok {
		fmt.Printf("freeIndex: %d lastIndex: %d\n", freeIndex, lastIndex)
		var f = blocks[lastIndex][0]
		fmt.Printf("file to defrag: %v\n", f)
		ok = defragFile(blocks, freeIndex, f)
		freeIndex = findNextFree(blocks, freeIndex)
		if freeIndex >= lastIndex {
			break
		}
		if !ok {
			break
		}
		if f.id == -1 {
			lastIndex -= 2
		}
	}

	fmt.Println(blocks)

	return checksum(blocks)
}

func findFreeBlock(blocks []block, maxIndex int, size int) int {
	for i := 0; i < maxIndex; i++ {
		var b = blocks[i]
		if len(b) > 0 {
			var last = b[len(b)-1]
			if last.id == -1 && last.len >= size {
				return i
			}
		}
	}
	return -1
}

func defragFile2(blocks []block, maxIndex int, f *file) bool {
	var indexFree = findFreeBlock(blocks, maxIndex, f.len)
	if indexFree == -1 {
		return false
	}

	var freeBlock = blocks[indexFree]
	fmt.Printf("found free block: %v\n", freeBlock)
	var freeFile = freeBlock[len(freeBlock)-1]
	if freeFile.id != -1 {
		fmt.Println("freeFile.id != -1")
		return false
	}

	if f.len > freeFile.len {
		// fill freeFile and split f
		freeFile.id = f.id
		f.len -= freeFile.len
		return true
	} else if f.len == freeFile.len {
		// fill freeFile and make f free
		freeFile.id = f.id
		f.id = -1
		return true
	} else if f.len < freeFile.len {
		// fill freeFile, split freeFile and make f free
		var newFile = &file{f.id, f.len, freeFile.adr}
		freeFile.len -= f.len
		freeFile.adr += f.len
		f.id = -1
		// insert f in freeBlock before last position
		freeBlock = append(freeBlock[:len(freeBlock)-1], newFile, freeBlock[len(freeBlock)-1])
		blocks[indexFree] = freeBlock
		return true
	}

	return false

}

func checksum2(blocks []block) int {
	var res int
	var adr int
	for _, b := range blocks {
		for _, f := range b {
			if f.id != -1 {
				for i := 0; i < f.len; i++ {
					res += (adr * f.id)
					adr++
				}
			}
			if f.id == -1 {
				adr += f.len
			}
		}
	}
	return res
}

func Part2(input string) int {
	input = strings.TrimSpace(input)
	var blocks = parse(input)

	fmt.Println(blocks)

	var lastIndex = len(blocks) - 1
	if lastIndex%2 != 0 {
		lastIndex--
	}

	for lastIndex >= 1 {
		fmt.Printf("lastIndex: %d\n", lastIndex)
		var f = blocks[lastIndex][0]
		fmt.Printf("file to move: %v\n", f)
		defragFile2(blocks, lastIndex, f)
		lastIndex -= 2
	}

	fmt.Println(blocks)

	return checksum2(blocks)
}

func main() {
	fmt.Println("--2024 day 09 solution--")
	var inputDay = utils.Input()
	//var inputDay = inputTest
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
