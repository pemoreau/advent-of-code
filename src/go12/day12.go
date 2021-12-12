package main

import (
	"fmt"
	"io/ioutil"
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

func (g *Graph) explore1(src, dest string, count int) int {
	if src == dest {
		return count + 1
	}
	srcNode := g.nodes[src]
	for _, nextLabel := range srcNode.next {
		nextNode := g.nodes[nextLabel]
		visitable := unicode.IsUpper(rune(nextNode.label[0])) || nextNode.visited == 0
		if visitable {
			nextNode.visited += 1
			count = g.explore1(nextLabel, dest, count)
			nextNode.visited -= 1
		}
	}

	return count
}

func (g *Graph) explore2(src, dest string, count int, twice bool) int {
	if src == dest {
		return count + 1
	}
	srcNode := g.nodes[src]
	for _, nextLabel := range srcNode.next {
		nextNode := g.nodes[nextLabel]
		lower := unicode.IsLower(rune(nextNode.label[0]))
		if !lower || (lower && nextNode.visited == 0) || (lower && nextLabel != "start" && !twice) {
			if lower && nextNode.visited == 1 {
				twice = true
			}
			nextNode.visited += 1
			count = g.explore2(nextLabel, dest, count, twice)
			nextNode.visited -= 1
			if lower && nextNode.visited == 1 {
				twice = false
			}
		}
	}

	return count
}

func Part1(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	g := BuildGraph(lines)
	g.nodes["start"].visited = 1
	return g.explore1("start", "end", 0)
}

func Part2(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	g := BuildGraph(lines)
	g.nodes["start"].visited = 1
	return g.explore2("start", "end", 0, false)
}

func main() {
	content, _ := ioutil.ReadFile("../../inputs/day12.txt")

	start := time.Now()
	fmt.Println("part1: ", Part1(string(content)))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(string(content)))
	fmt.Println(time.Since(start))
}
