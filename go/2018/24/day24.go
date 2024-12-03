package main

import (
	_ "embed"
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils"
	"slices"
	"strings"
	"time"
)

//go:embed sample.txt
var inputTest string

type Unit struct {
	hitPoints    int
	attackDamage int
	attackType   string
	initiative   int
	weaknesses   string
	immunities   string
}

type Group struct {
	number int
	team   string
	units  int
	unit   Unit
}

func parseWeaknessesAndImmunities(wi string) (weaknesses string, immunities string) {
	if !strings.Contains(wi, ";") {
		wi = wi + ";"
	}

	var parts = strings.Split(wi, ";")

	for _, part := range parts {
		part = strings.Trim(part, " ")
		if strings.Contains(part, "weak to") {
			weaknesses = part
		} else if strings.Contains(part, "immune to") {
			immunities = part
		}
	}
	return weaknesses, immunities
}

func parseGroup(team string, number int, line string) *Group {

	var unit Unit
	var group Group
	var wi string
	var before, after, ok = strings.Cut(line, "(")
	if ok {
		wi, after, _ = strings.Cut(after, ") with an")
	} else {
		before, after, _ = strings.Cut(line, "with an")
	}
	before = strings.Trim(before, " ")
	after = strings.Trim(after, " ")

	fmt.Sscanf(before, "%d units each with %d hit points",
		&group.units, &unit.hitPoints)
	fmt.Sscanf(after, "attack that does %d %s damage at initiative %d",
		&unit.attackDamage, &unit.attackType, &unit.initiative)
	unit.weaknesses, unit.immunities = parseWeaknessesAndImmunities(wi)
	group.unit = unit
	group.team = team
	group.number = number

	return &group
}

func parseInput(input string) ([]*Group, []*Group) {
	var parts = strings.Split(input, "\n\n")
	var linesImmuneSystem = strings.Split(parts[0], "\n")[1:]
	var linesInfection = strings.Split(parts[1], "\n")[1:]

	var immuneSystem []*Group
	var infection []*Group
	for i, line := range linesImmuneSystem {
		immuneSystem = append(immuneSystem, parseGroup("Immune System", i+1, line))
	}
	for i, line := range linesInfection {
		infection = append(infection, parseGroup("Infection", i+1, line))
	}

	return immuneSystem, infection
}

func (g Group) effectivePower() int {
	return g.units * g.unit.attackDamage
}

func (g Group) damageTo(target Group) int {
	if strings.Contains(target.unit.immunities, g.unit.attackType) {
		return 0
	}
	var damage = g.effectivePower()
	if strings.Contains(target.unit.weaknesses, g.unit.attackType) {
		damage *= 2
	}
	return damage
}

func targetSelection(immuneSystem, infection []*Group) map[*Group]*Group {
	immuneSystem = slices.Clone(immuneSystem)
	infection = slices.Clone(infection)
	var assignment = make(map[*Group]*Group)

	var all = slices.Concat(immuneSystem, infection)
	// sort all by decreasing effective power and initiative
	slices.SortFunc(all, func(a, b *Group) int {
		if a.effectivePower() == b.effectivePower() {
			return b.unit.initiative - a.unit.initiative
		}
		return b.effectivePower() - a.effectivePower()
	})

	for _, attacker := range all {
		var targets []*Group
		if attacker.team == "Immune System" {
			targets = infection
		} else {
			targets = immuneSystem
		}

		// sort targets by damageTo
		slices.SortFunc(targets, func(a, b *Group) int {
			if attacker.damageTo(*a) == attacker.damageTo(*b) {
				if a.effectivePower() == b.effectivePower() {
					return b.unit.initiative - a.unit.initiative
				}
				return b.effectivePower() - a.effectivePower()
			}
			return attacker.damageTo(*b) - attacker.damageTo(*a)
		})

		// fmt.Printf("attacker: %v\n", attacker)

		// select first target
		for len(targets) > 0 {
			var target = targets[0]
			targets = targets[1:]
			if target.units == 0 {
				continue
			}
			if attacker.damageTo(*target) > 0 {
				if attacker.team == "Immune System" {
					assignment[attacker] = target
					index := slices.Index(infection, target)
					infection = slices.Delete(infection, index, index+1)
				} else {
					assignment[attacker] = target
					index := slices.Index(immuneSystem, target)
					immuneSystem = slices.Delete(immuneSystem, index, index+1)
				}
				// fmt.Printf("target: %v\n", target)
				break
			}
		}
	}

	return assignment
}

func attacking(immuneSystem, infection []*Group, assignement map[*Group]*Group) int {
	var arena = slices.Concat(immuneSystem, infection)

	// sort all by decreasing initiative
	slices.SortFunc(arena, func(a, b *Group) int {
		return b.unit.initiative - a.unit.initiative
	})

	var killedUnits = 0
	for _, attacker := range arena {
		if attacker.units == 0 {
			continue
		}
		//fmt.Printf("attacker: %v\n", attacker)
		var target = assignement[attacker]
		if target == nil || target.units == 0 {
			continue
		}
		//fmt.Printf("target: %v\n", target)
		var damage = attacker.damageTo(*target)
		var killed = damage / target.unit.hitPoints
		if killed > target.units {
			killed = target.units
		}
		target.units -= killed
		killedUnits += killed
		//fmt.Printf("%s group %d attacks defending group %d, killing %d units\n",
		//	attacker.team, attacker.number, target.number, killed)
	}
	return killedUnits
}

func display(groups []*Group) {
	if len(groups) == 0 {
		fmt.Println("no groups")
		return
	}
	fmt.Println(groups[0].team)
	for _, group := range groups {
		fmt.Printf("Group %d contains %d units\n", group.number, group.units)
	}
}

func alive(groups []*Group) int {
	var count = 0
	for _, group := range groups {
		count += group.units
	}
	return count
}

func Part1(input string) int {
	immuneSystem, infection := parseInput(input)

	for alive(immuneSystem) > 0 && alive(infection) > 0 {
		//display(immuneSystem)
		//display(infection)
		var assignement = targetSelection(immuneSystem, infection)
		attacking(immuneSystem, infection, assignement)
		//fmt.Println("----------------")

	}

	return alive(immuneSystem) + alive(infection)
}

func deepCopy(groups []*Group) []*Group {
	var copy []*Group
	for _, group := range groups {
		var unit Unit
		unit.attackDamage = group.unit.attackDamage
		unit.attackType = group.unit.attackType
		unit.hitPoints = group.unit.hitPoints
		unit.initiative = group.unit.initiative
		unit.immunities = group.unit.immunities
		unit.weaknesses = group.unit.weaknesses
		var newGroup = Group{
			number: group.number,
			team:   group.team,
			units:  group.units,
			unit:   unit,
		}
		copy = append(copy, &newGroup)
	}
	return copy
}

func Part2(input string) int {
	originImmuneSystem, originInfection := parseInput(input)

	var boost int

	for {
		boost++
		var immuneSystem = deepCopy(originImmuneSystem)
		var infection = deepCopy(originInfection)

		// boost immune system
		for _, group := range immuneSystem {
			group.unit.attackDamage += boost
		}

		for alive(immuneSystem) > 0 && alive(infection) > 0 {
			//display(immuneSystem)
			//display(infection)
			var assignement = targetSelection(immuneSystem, infection)
			killedUnits := attacking(immuneSystem, infection, assignement)
			if killedUnits == 0 {
				break
			}
			//fmt.Println("----------------")
		}
		if alive(immuneSystem) > 0 && alive(infection) == 0 {
			return alive(immuneSystem)
		}
	}
	return 0
}

func main() {
	fmt.Println("--2018 day 24 solution--")
	var inputDay = utils.Input()
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}
