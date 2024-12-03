package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"strings"
	"time"
	"unicode"
)

type Graph struct {
	nodes map[string]*Node
}

type Node struct {
	label   string
	small   bool
	visited int
	next    []*Node
}

func (g *Graph) addNode(label string) *Node {
	if _, ok := g.nodes[label]; !ok {
		g.nodes[label] = &Node{label: label, small: unicode.IsLower(rune(label[0]))}
	}
	return g.nodes[label]
}

func (g *Graph) addEdge(src, dest string) {
	srcNode := g.addNode(src)
	destNode := g.addNode(dest)
	srcNode.next = append(srcNode.next, destNode)
	destNode.next = append(destNode.next, srcNode)
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

// twice is used for 2nd part only: should be true for 1st part
func (g *Graph) explore(src *Node, dest *Node, count int, twice bool) int {
	if src == dest {
		return count + 1
	}

	for _, nextNode := range src.next {
		visitable := !nextNode.small || (nextNode.visited == 0) || (nextNode.label != "start" && !twice)
		if visitable {
			twiceUsed := nextNode.small && nextNode.visited == 1
			nextNode.visited += 1
			count = g.explore(nextNode, dest, count, twice || twiceUsed)
			nextNode.visited -= 1
		}
	}
	return count
}

func Part1(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	g := BuildGraph(lines)
	start := g.nodes["start"]
	end := g.nodes["end"]
	start.visited = 1
	return g.explore(start, end, 0, true)
}

func Part2(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	g := BuildGraph(lines)
	start := g.nodes["start"]
	end := g.nodes["end"]
	start.visited = 1
	return g.explore(start, end, 0, false)
}

func main() {

	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
