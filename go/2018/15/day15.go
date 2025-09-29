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
func isEnemy(occ *game2d.Matrix[*Unit], x, y int, u *Unit) bool {
	if y >= 0 && y < occ.LenY() && x >= 0 && x < occ.LenX() {
		v := occ.Get(x, y)
		return v != nil && v.Alive && v.Race != u.Race
	}
	return false
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
	minY, maxY := grid.MinY(), grid.MaxY()
	minX, maxX := grid.MinX(), grid.MaxX()
	h, w := maxY-minY+1, maxX-minX+1
	occ := game2d.NewMatrix[*Unit](w, h, func(_ *Unit) string { return "." })
	for i := range units {
		u := &units[i]
		occ.Set(u.Pos.X-minX, u.Pos.Y-minY, u)
	}

	for {
		// Ordre de jeu : positions en ordre de lecture, figées au début du round
		order := make([]*Unit, 0, len(units))
		for i := range units {
			if units[i].Alive {
				order = append(order, &units[i])
			}
		}
		sort.Slice(order, func(i, j int) bool {
			if order[i].Pos.Y != order[j].Pos.Y {
				return order[i].Pos.Y < order[j].Pos.Y
			}
			return order[i].Pos.X < order[j].Pos.X
		})
		for _, u := range order {
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
				dest, _, firstStepMap := chooseDestination(grid, occ, units, u)
				if dest.X != -1 {
					step := firstStepMap[dest]
					if step.X != -1 {
						occ.Set(u.Pos.X, u.Pos.Y, nil)
						u.Pos = step
						occ.Set(u.Pos.X, u.Pos.Y, u)
					}
				}
			}
			// Attaquer si possible
			target := chooseAttackTarget(occ, u)
			if target != nil {
				target.HP -= u.AP
				if target.HP <= 0 {
					target.Alive = false
					occ.Set(target.Pos.X, target.Pos.Y, nil)
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

func hasAdjacentEnemy(occ *game2d.Matrix[*Unit], u *Unit) bool {
	for _, d := range dirs {
		x, y := u.Pos.X+d.X, u.Pos.Y+d.Y
		if isEnemy(occ, x, y, u) {
			return true
		}
	}
	return false
}

func chooseAttackTarget(occ *game2d.Matrix[*Unit], u *Unit) *Unit {
	var best *Unit
	for _, d := range dirs { // ordre lecture: Up, Left, Right, Down
		x, y := u.Pos.X+d.X, u.Pos.Y+d.Y
		if !isEnemy(occ, x, y, u) {
			continue
		}
		v := occ.Get(x, y)
		if best == nil || v.HP < best.HP ||
			(v.HP == best.HP && (v.Pos.Y < best.Pos.Y || (v.Pos.Y == best.Pos.Y && v.Pos.X < best.Pos.X))) {
			best = v
		}
	}
	return best
}

// Renvoie la destination, la distance, et une map pour retrouver le premier pas optimal
func chooseDestination(grid *game2d.GridChar, occ *game2d.Matrix[*Unit], units []Unit, u *Unit) (game2d.Pos, int, map[game2d.Pos]game2d.Pos) {
	// 1. Collecter toutes les cases "in range" (adjacentes aux ennemis) qui sont libres '.'
	targets := set.NewSet[game2d.Pos]()
	for i := range units {
		v := &units[i]
		if !v.Alive || v.Race == u.Race {
			continue
		}
		for n := range v.Pos.Neighbors4() {
			vGrid, ok := grid.GetPos(n)
			if !ok {
				continue
			}
			if vGrid == '.' && occ.Get(n.X-grid.MinX(), n.Y-grid.MinY()) == nil {
				targets.Add(n)
			}
		}
	}
	if targets.Len() == 0 {
		return game2d.Pos{X: -1, Y: -1}, -1, nil
	}

	// 2. BFS unique depuis la position de l'unité
	type bfsItem struct {
		Pos       game2d.Pos
		Dist      int
		FirstStep game2d.Pos // Premier pas depuis la position de départ
	}
	queue := []bfsItem{{Pos: u.Pos, Dist: 0, FirstStep: u.Pos}}
	visited := map[game2d.Pos]bool{u.Pos: true}
	firstStepMap := make(map[game2d.Pos]game2d.Pos) // Pour chaque case atteinte, le premier pas depuis la position de départ
	var foundTargets []bfsItem
	minDist := -1
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		if item.Dist > 0 {
			firstStepMap[item.Pos] = item.FirstStep
		}
		if targets.Contains(item.Pos) {
			if minDist == -1 || item.Dist == minDist {
				foundTargets = append(foundTargets, item)
				minDist = item.Dist
			} else if item.Dist > minDist {
				break // On a déjà trouvé toutes les cibles à distance minimale
			}
		}
		if minDist != -1 && item.Dist > minDist {
			break
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
			if vGrid == '.' && (occ.Get(nx-grid.MinX(), ny-grid.MinY()) == nil || (nx == u.Pos.X && ny == u.Pos.Y)) {
				visited[p] = true
				first := item.FirstStep
				if item.Dist == 0 {
					first = p // Premier pas
				}
				queue = append(queue, bfsItem{Pos: p, Dist: item.Dist + 1, FirstStep: first})
			}
		}
	}
	if len(foundTargets) == 0 {
		return game2d.Pos{X: -1, Y: -1}, -1, nil
	}
	// 3. Choisir la cible la plus haute, puis la plus à gauche
	best := foundTargets[0].Pos
	for _, t := range foundTargets[1:] {
		if t.Pos.Y < best.Y || (t.Pos.Y == best.Y && t.Pos.X < best.X) {
			best = t.Pos
		}
	}
	return best, minDist, firstStepMap
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
