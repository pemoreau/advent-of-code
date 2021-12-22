package main

import "fmt"

type World map[Cuboid]int

func (w World) String() string {
	var s string
	for c, v := range w {
		// if v != 0 {
		s += fmt.Sprintf("%v: %v\n", c, v)
		// }
	}
	return s
}

func (w World) Add(c Cuboid, info int) {
	for key, ki := range w {
		if ki > 0 {
			if info == ki && Include(c, key) {
				// do nothing
				return
			}
			if info == ki && Include(key, c) {
				// key subsumed
				w[key] = 0 // remove key
			} else if _, ok := Intersection(key, c); ok {
				w[key] = 0 // remove key
				list := Overlap(key, c)
				for _, e := range list {
					if !Include(e, c) {
						w[e] = ki
					}
				}
			}
		}
	}
	w[c] = info
}

func (w World) Count(info int) int {
	count := 0
	for c, v := range w {
		if v == info {
			count += Size(c)
		}
	}
	return count
}
