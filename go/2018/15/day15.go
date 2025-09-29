package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/pemoreau/advent-of-code/go/utils/set"

	"github.com/pemoreau/advent-of-code/go/utils"
	"github.com/pemoreau/advent-of-code/go/utils/game2d"
)

var dirs = []game2d.Pos{
	{X: 0, Y: -1}, // Up
	{X: -1, Y: 0}, // Left
	{X: 1, Y: 0},  // Right
	{X: 0, Y: 1},  // Down
} // Up, Left, Right, Down (ordre de lecture)

var invalidPos = game2d.Pos{X: -1, Y: -1}

func lessReadingOrder(a, b game2d.Pos) bool {
	if a.Y != b.Y {
		return a.Y < b.Y
	}
	return a.X < b.X
}

func moveUnit(u *Unit, step game2d.Pos, occ map[game2d.Pos]*Unit) {
	delete(occ, u.Pos)
	u.Pos = step
	occ[u.Pos] = u
}

func turnOrder(units []Unit) []*Unit {
	order := make([]*Unit, 0, len(units))
	for i := range units {
		if units[i].Alive {
			order = append(order, &units[i])
		}
	}
	sort.Slice(order, func(i, j int) bool {
		return lessReadingOrder(order[i].Pos, order[j].Pos)
	})
	return order
}

type Unit struct {
	Pos   game2d.Pos
	HP    int
	AP    int
	Race  byte // 'E' ou 'G'
	Alive bool
}

func Part1(input string) int {
	_, outcome, _ := simulate(input, 3, false)
	return outcome
}

func Part2(input string) int {
	// Cherche le plus petit AP des Elfes tel qu'aucun Elfe ne meure
	elfAP := 4
	for {
		_, outcome, elvesDied := simulate(input, elfAP, true)
		if !elvesDied {
			return outcome
		}
		elfAP++
	}
}

func main() {
	fmt.Println("--2018 day 15 solution--")
	var inputDay = utils.Input()

	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}

// ─── Simulation ─────────────────────────────────────────────────────────────────

// Teste si la case (x, y) contient un ennemi vivant de u
func isEnemy(occ map[game2d.Pos]*Unit, pos game2d.Pos, u *Unit) bool {
	v := occ[pos]
	return v != nil && v.Alive && v.Race != u.Race
}

func simulate(input string, elfAP int, stopIfElfDies bool) (rounds int, outcome int, elvesDied bool) {
	grid, units := parse(input)
	// set AP
	for i := range units {
		units[i].HP = 200
		units[i].Alive = true
		if units[i].Race == 'E' {
			units[i].AP = elfAP
		} else {
			units[i].AP = 3
		}
	}
	occ := make(map[game2d.Pos]*Unit, len(units))
	for i := range units {
		u := &units[i]
		occ[u.Pos] = u
	}

	for {
		for _, u := range turnOrder(units) {
			if !u.Alive {
				continue
			}
			// Combat s'arrête si aucune cible pour cette unité au début de son tour
			if !hasEnemyAlive(units, u.Race) {
				sum := sumHP(units)
				return rounds, rounds * sum, elvesDied
			}
			// Si pas d'ennemi adjacent, tenter de se déplacer
			if !hasAdjacentEnemy(occ, u) {
				dest, step := chooseDestination(grid, occ, units, u)
				if dest != invalidPos && step != invalidPos {
					moveUnit(u, step, occ)
				}
			}
			// Attaquer si possible
			target := chooseAttackTarget(occ, u)
			if target != nil {
				target.HP -= u.AP
				if target.HP <= 0 {
					target.Alive = false
					delete(occ, target.Pos)
					if target.Race == 'E' && stopIfElfDies {
						return rounds, 0, true
					}
				}
			}
		}
		rounds++
	}
}

// ─── Parsing ────────────────────────────────────────────────────────────────────

func parse(input string) (grid *game2d.GridChar, units []Unit) {
	grid = game2d.BuildGridCharFromString(input)
	for p := range grid.AllPos() {
		v, ok := grid.Get(p.X, p.Y)
		if !ok {
			continue
		}
		switch v {
		case 'E', 'G':
			units = append(units, Unit{Pos: p, Race: v, HP: 200, AP: 3, Alive: true})
			grid.SetPos(p, '.')
		}
	}
	return
}

// ─── Helpers combat ────────────────────────────────────────────────────────────

func hasEnemyAlive(units []Unit, my byte) bool {
	opp := byte('E')
	if my == 'E' {
		opp = 'G'
	}
	for i := range units {
		if units[i].Alive && units[i].Race == opp {
			return true
		}
	}
	return false
}

func hasAdjacentEnemy(occ map[game2d.Pos]*Unit, u *Unit) bool {
	for next := range u.Pos.Neighbors4() {
		if isEnemy(occ, next, u) {
			return true
		}
	}
	return false
}

func chooseAttackTarget(occ map[game2d.Pos]*Unit, u *Unit) *Unit {
	var best *Unit
	for _, d := range dirs { // ordre lecture: Up, Left, Right, Down
		pos := game2d.Pos{X: u.Pos.X + d.X, Y: u.Pos.Y + d.Y}
		if !isEnemy(occ, pos, u) {
			continue
		}
		v := occ[pos]
		if best == nil || v.HP < best.HP || (v.HP == best.HP && lessReadingOrder(v.Pos, best.Pos)) {
			best = v
		}
	}
	return best
}

func collectTargets(grid *game2d.GridChar, occ map[game2d.Pos]*Unit, units []Unit, u *Unit) set.Set[game2d.Pos] {
	targets := set.NewSet[game2d.Pos]()
	for i := range units {
		v := &units[i]
		if !v.Alive || v.Race == u.Race {
			continue
		}
		for n := range v.Pos.Neighbors4() {
			vGrid, ok := grid.Get(n.X, n.Y)
			if !ok {
				continue
			}
			if vGrid == '.' && occ[n] == nil {
				targets.Add(n)
			}
		}
	}
	return targets
}

// Renvoie la destination et le premier pas optimal
func chooseDestination(grid *game2d.GridChar, occ map[game2d.Pos]*Unit, units []Unit, u *Unit) (game2d.Pos, game2d.Pos) {
	// 1. Collecter toutes les cases "in range" (adjacentes aux ennemis) qui sont libres '.'
	targets := collectTargets(grid, occ, units, u)
	if targets.Len() == 0 {
		return invalidPos, invalidPos
	}

	// 2. BFS unique depuis la position de l'unité
	type bfsItem struct {
		Pos       game2d.Pos
		Dist      int
		FirstStep game2d.Pos // Premier pas depuis la position de départ
	}
	queue := []bfsItem{{Pos: u.Pos, Dist: 0, FirstStep: u.Pos}}
	visited := map[game2d.Pos]bool{u.Pos: true}
	var bestTarget bfsItem
	minDist := -1
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		if minDist != -1 && item.Dist > minDist {
			break
		}
		if targets.Contains(item.Pos) {
			if item.Dist == 0 {
				item.FirstStep = invalidPos
			}
			if minDist == -1 {
				minDist = item.Dist
				bestTarget = item
			} else if lessReadingOrder(item.Pos, bestTarget.Pos) {
				bestTarget = item
			}
			continue
		}
		for _, d := range dirs {
			nx, ny := item.Pos.X+d.X, item.Pos.Y+d.Y
			p := game2d.Pos{X: nx, Y: ny}
			if visited[p] {
				continue
			}
			vGrid, ok := grid.Get(nx, ny)
			if !ok {
				continue
			}
			if vGrid == '.' && (occ[p] == nil || p == u.Pos) {
				visited[p] = true
				first := item.FirstStep
				if item.Dist == 0 {
					first = p // Premier pas
				}
				queue = append(queue, bfsItem{Pos: p, Dist: item.Dist + 1, FirstStep: first})
			}
		}
	}
	if minDist == -1 {
		return invalidPos, invalidPos
	}
	return bestTarget.Pos, bestTarget.FirstStep
}

func sumHP(units []Unit) int {
	total := 0
	for i := range units {
		if units[i].Alive {
			total += units[i].HP
		}
	}
	return total
}
