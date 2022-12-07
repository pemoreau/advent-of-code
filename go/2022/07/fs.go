package main

import (
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
)

type fs struct {
	root *node
	path *utils.Stack[*node]
}

type node struct {
	isdir    bool
	name     string
	_size    int
	children []*node
}

func NewFS() fs {
	root := NewDir("/")
	res := fs{root: root, path: &utils.Stack[*node]{}}
	res.path.Push(root)
	return res
}

func NewDir(name string) *node {
	return &node{isdir: true, name: name, children: []*node{}}
}

func NewFile(name string, size int) *node {
	return &node{name: name, _size: size}
}

func (n *node) size() int {
	size := n._size
	if n.isdir {
		for _, child := range n.children {
			size += child.size()
		}
	}
	return size
}

func (fs fs) addNode(node *node) {
	current, _ := fs.path.Peek()
	exist := false
	for _, child := range current.children {
		if child.name == node.name {
			exist = true
			break
		}
	}
	if !exist {
		current.children = append(current.children, node)
	}
}

func (fs fs) mkdir(name string) {
	fs.addNode(NewDir(name))
}

func (fs fs) mkfile(name string, size int) {
	fs.addNode(NewFile(name, size))
}

func (fs fs) cd(dir string) {
	if dir == "/" {
		for len(*fs.path) > 1 {
			fs.path.Pop()
		}
		return
	}
	if dir == "." {
		return
	}
	if dir == ".." {
		if len(*fs.path) > 1 {
			fs.path.Pop()
		}
		return
	}
	// cd name
	current, _ := fs.path.Peek()
	for _, child := range current.children {
		if child.name == dir {
			fs.path.Push(child)
			return
		}
	}
}

func (fs fs) String() string {
	current, _ := fs.path.Peek()
	return fmt.Sprintf("root: %v\n   -> %v", fs.root, current)
}

func (n *node) String() string {
	if n.isdir {
		return fmt.Sprintf("dir %s %v", n.name, n.children)
	}
	return fmt.Sprintf("(%d %v)", n._size, n.name)
}

func (n *node) search(max int) []*node {
	res := []*node{}
	if n.isdir {
		if n.size() < max {
			res = append(res, n)
		}
		for _, child := range n.children {
			res = append(res, child.search(max)...)
		}
	}
	return res
}

func (n *node) dirsizes() []int {
	res := []int{}
	if n.isdir {
		res = append(res, n.size())
		for _, child := range n.children {
			res = append(res, child.dirsizes()...)
		}
	}
	return res
}
