package main

import (
	"fmt"
	"strings"
	"time"
	"unicode"
)

type Graph struct {
	nodes map[string]*Node
}

type Node struct {
	label   string
	visited int
	next    []string
}

func (g *Graph) getNode(label string) *Node {
	return g.nodes[label]
}

func (g *Graph) addNode(label string) *Node {
	if _, ok := g.nodes[label]; !ok {
		g.nodes[label] = &Node{label: label}
	}
	return g.nodes[label]
}

func (g *Graph) addEdge(src, dest string) {
	n := g.addNode(src)
	n.next = append(n.next, dest)
	n = g.addNode(dest)
	n.next = append(n.next, src)
}

func (g *Graph) String() string {
	var sb strings.Builder
	for _, n := range g.nodes {
		sb.WriteString(fmt.Sprintf("%s: %s\n", n.label, n.next))
	}
	return sb.String()
}

func (n *Node) isVisited() bool {
	if unicode.IsUpper(rune(n.label[0])) {
		return false
	}
	return n.visited
}

func (n *Node) setVisited(value bool) {
	if !unicode.IsUpper(rune(n.label[0])) {
		n.visited = value
	}
}

func BuildGraph(lines []string) Graph {
	g := Graph{nodes: map[string]*Node{}}
	for _, l := range lines {
		relation := strings.SplitN(l, "-", 2)
		left := strings.TrimSpace(relation[0])
		right := strings.TrimSpace(relation[1])
		g.addEdge(left, right)
	}
	return g
}

func (g *Graph) explore(src, dest string, count int) int {
	if src == dest {
		println("found", src, dest)
		return count + 1
	}
	n := g.nodes[src]
	n.setVisited(true)
	for _, label := range n.next {
		if !g.nodes[label].isVisited() {
			count = g.explore(label, dest, count)
		}
	}
	n.setVisited(false)

	return count
}

func Part1(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	g := BuildGraph(lines)
	println(g.String())
	res := g.explore("start", "end", 0)

	return res
}

func Part2(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	g := BuildGraph(lines)
	println(g.String())
	return 0
}

const input1 = `start-A
start-b
A-c
A-b
b-d
A-end
b-end`

//
//const input2 = `dc-end
//HN-start
//start-kj
//dc-start
//dc-HN
//LN-dc
//HN-end
//kj-sa
//kj-HN
//kj-dc`
//
//const input3 = `fs-end
//he-DX
//fs-he
//start-DX
//pj-DX
//end-zg
//zg-sl
//zg-pj
//pj-he
//RW-he
//fs-DX
//pj-RW
//zg-RW
//start-pj
//he-WI
//zg-he
//pj-fs
//start-RW`

func main() {
	//content, _ := ioutil.ReadFile("../../inputs/day12.txt")

	start := time.Now()
	//fmt.Println("part1: ", Part1(string(content)))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(string(input1)))
	fmt.Println(time.Since(start))
}
