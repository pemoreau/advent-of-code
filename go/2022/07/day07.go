package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"strings"
	"time"
)

//go:embed input_test.txt
var input_day string

type node interface {
	name() string
	size() int
	sumSize(max int) int
}

type file struct {
	fname string
	fsize int
}

type directory struct {
	dname    string
	children []node
}

func (f file) size() int {
	return f.fsize
}

func (d directory) name() string {
	return d.dname
}

func (f file) name() string {
	return f.fname
}

type fs struct {
	root *directory
	path utils.Stack[*directory]
}

func NewFS() fs {
	root := directory{dname: "/", children: []node{}}
	res := fs{root: &root, path: utils.Stack[*directory]{}}
	res.path.Push(&root)
	return res
}

func (fs fs) addDir(name string) {
	fmt.Println("add dir", name)
	dir := directory{dname: name, children: []node{}}
	top, _ := fs.path.Peek()
	exist := false
	for _, child := range top.children {
		if child.name() == name {
			exist = true
			break
		}
	}
	if !exist {
		top.children = append(top.children, dir)
	}
}

func (fs fs) addFile(name string, size int) {
	fmt.Println("add file", name, size, fs)
	file := file{fname: name, fsize: size}
	top, _ := fs.path.Peek()
	exist := false
	for _, child := range top.children {
		if child.name() == name {
			exist = true
			break
		}
	}
	if !exist {
		top.children = append(top.children, file)
	} else {
		panic("file already exist")
	}
}

func (fs fs) String() string {
	top, _ := fs.path.Peek()
	return fmt.Sprintf("root: %v\n   -> %v", fs.root, top)
}

func (d directory) String() string {
	return fmt.Sprintf("dir %s %v", d.dname, d.children)
}

func (f file) String() string {
	return fmt.Sprintf("file %d %v", f.fsize, f.fname)
}

func cd(fs fs, dir string) {
	fmt.Println(fs)
	fmt.Println("cd", dir)
	if dir == "/" {
		fmt.Println(fs)
		for !fs.path.IsEmpty() {
			fs.path.Pop()
		}
		return
	}
	if dir == "." {
		return
	}
	if dir == ".." {
		fs.path.Pop()
		return
	}
	// cd name
	fmt.Println("cd into: ", dir)
	current, _ := fs.path.Peek()
	for _, child := range current.children {
		if child.name() == dir {
			switch child.(type) {
			case directory:
				d := child.(directory)
				fs.path.Push(&d)
				fmt.Println("Pushed", d)
				fmt.Println(fs)
			default:
				panic("not a directory")
			}
			return
		}
	}
}

func (f file) sumSize(max int) int {
	//if f.fsize < max {
	//	return f.fsize
	//}
	return 0
}

func (d directory) size() int {
	size := 0
	for _, child := range d.children {
		size += child.size()
	}
	return size
}

func (d directory) sumSize(max int) int {
	sum := 0
	mysize := d.size()
	if mysize < max {
		sum += mysize
		fmt.Println("count1", mysize)
	}
	for _, child := range d.children {
		switch child.(type) {
		case file:
			break
		case directory:
			s := child.size()
			if s < max {
				sum += s
				fmt.Println("count2", s, child)
			}
			break
		default:
			panic("unknown type")
		}
	}

	return sum
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")
	fs := NewFS()
	for _, line := range lines {
		if strings.HasPrefix(line, "$ ") {
			command := strings.TrimPrefix(line, "$ ")
			if strings.HasPrefix(command, "cd ") {
				dir := strings.TrimPrefix(command, "cd ")
				cd(fs, dir)
			} else {
				fmt.Println("ls")
			}
		} else {
			if strings.HasPrefix(line, "dir ") {
				dir := strings.TrimPrefix(line, "dir ")
				fs.addDir(dir)
			} else {
				var size int
				var name string
				fmt.Sscanf(line, "%d %s", &size, &name)
				fs.addFile(name, size)
			}
		}
	}

	fmt.Println("total size", fs.root.size())
	fmt.Println("max 100000", fs.root.sumSize(100000))
	return 0
}

func Part2(input string) int {
	// input = strings.TrimSuffix(input, "\n")
	// lines := strings.Split(input, "\n")
	return 0

}

func main() {
	fmt.Println("--2022 day 07 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(input_day))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(input_day))
	fmt.Println(time.Since(start))
}
