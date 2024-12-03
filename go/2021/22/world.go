package main

import (
	"fmt"
	"github.com/pemoreau/advent-of-code/go/utils/interval"
)

type World map[interval.Cuboid]uint8

func (w World) String() string {
	var s string
	for c, v := range w {
		s += fmt.Sprintf("%v: %v\n", c, v)
	}
	return s
}

func (w World) Add(c interval.Cuboid, info uint8) {
	for key, ki := range w {
		if !key.Disjoint(c) {
			if info == ki && key.Contains(c) {
				return // do nothing
			}
			delete(w, key) // remove key
			for _, e := range c.Overlap(key) {
				if !c.Contains(e) {
					w[e] = ki
				}
			}
		}
	}
	w[c] = info
}

func (w World) Count(info uint8) int {
	count := 0
	for c, v := range w {
		if v == info {
			count += c.Size()
		}
	}
	return count
}
