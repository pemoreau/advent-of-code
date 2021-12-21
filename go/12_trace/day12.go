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
	visit   []string
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

func isVisited1(n *Node, g *Graph, path []string) bool {
	if unicode.IsUpper(rune(n.label[0])) {
		return false
	}
	return n.visited >= 1
}

func isVisited2(n *Node, g *Graph, path []string) bool {
	if unicode.IsUpper(rune(n.label[0])) {
		return false
	}
	if n.label == "start" {
		return n.visited >= 1
	}
	nbLower := 0
	sum := 0
	for _, l := range path {
		if unicode.IsLower(rune(l[0])) {
			nbLower++
			sum += g.nodes[l].visited
		}
	}
	//fmt.Println(path, nbLower, sum)

	if nbLower == sum {
		return n.visited >= 2
	}

	return n.visited >= 1
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

func BuildSet() set {
	return make(map[string]struct{})
}

func (s set) Add(value string) {
	s[value] = struct{}{}
}

func (s set) Contains(value string) bool {
	_, ok := s[value]
	return ok
}
func (s set) Len() int {
	return len(s)
}

type fn func(n *Node, g *Graph, path []string) bool
type set map[string]struct{}

func (g *Graph) explore(src, dest string, count int, isVisited fn, path []string, allPaths *set) int {
	if src == dest {
		allPaths.Add(strings.Join(path, " "))

		//fmt.Println("path", path)
		//fmt.Println("found", dest, count+1, allPaths)
		return count + 1
	}
	srcNode := g.nodes[src]
	//srcNode.visited += 1
	//fmt.Printf("%s(%d,%s) --> \n", src, srcNode.visited, srcNode.visit)

	for _, nextLabel := range srcNode.next {
		nextNode := g.nodes[nextLabel]
		if !isVisited(nextNode, g, path) {
			//println("visiting", src, nextLabel, g.nodes[nextLabel].visited)
			//fmt.Printf("%s(%d) --> %s(%d)\n", src, srcNode.visited, nextLabel, nextNode.visited)
			//fmt.Println("will visit", nextLabel)

			nextNode.visited += 1
			path = append(path, nextLabel)
			count = g.explore(nextLabel, dest, count, isVisited, path, allPaths)
			nextNode.visited -= 1
			path = path[:len(path)-1]
		}
	}
	//fmt.Printf("leave %s(%d)\n", src, srcNode.visited)
	//srcNode.visited -= 1

	return count
}

func Part1(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	g := BuildGraph(lines)
	g.nodes["start"].visited = 1
	path := []string{"start"}
	allPaths := BuildSet()
	return g.explore("start", "end", 0, isVisited1, path, &allPaths)
}

func Part2(input string) int {
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")
	g := BuildGraph(lines)
	g.nodes["start"].visited = 1
	path := []string{"start"}
	allPaths := BuildSet()
	return g.explore("start", "end", 0, isVisited2, path, &allPaths)
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
