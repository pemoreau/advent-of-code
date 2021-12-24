package main

import (
	"fmt"
	"time"
)

func increment(inp []int, index int) []int {
	if index < 0 {
		return []int{}
	}
	if inp[index] == 9 {
		inp[index] = 1
		increment(inp, index-1)
	} else {
		inp[index]++
	}
	return inp
}
func decrement(inp []int, index int) []int {
	if index < 0 {
		return []int{}
	}
	if inp[index] == 1 {
		inp[index] = 9
		decrement(inp, index-1)
	} else {
		inp[index]--
	}
	if index < 6 {
		fmt.Println(inp)
	}
	return inp
}

func mod(a, b int) int {
	return a % b
}

func eql(a, b int) int {
	if a == b {
		return 1
	}
	return 0
}

func run(ww, xx, yy, zz, index int, inp []int) (w, x, y, z, i int) {
	w = ww
	x = ww
	y = yy
	z = zz
	i = index
	w = inp[i]
	i++
	x *= 0
	x += z
	x = mod(x, 26)
	z /= 1
	x += 14
	x = eql(x, w)
	x = eql(x, 0)
	y *= 0
	y += 25
	y *= x
	y += 1
	z *= y
	y *= 0
	y += w
	y += 12
	y *= x
	z += y
	w = inp[i]
	i++
	x *= 0
	x += z
	x = mod(x, 26)
	z /= 1
	x += 15
	x = eql(x, w)
	x = eql(x, 0)
	y *= 0
	y += 25
	y *= x
	y += 1
	z *= y
	y *= 0
	y += w
	y += 7
	y *= x
	z += y
	w = inp[i]
	i++
	x *= 0
	x += z
	x = mod(x, 26)
	z /= 1
	x += 12
	x = eql(x, w)
	x = eql(x, 0)
	y *= 0
	y += 25
	y *= x
	y += 1
	z *= y
	y *= 0
	y += w
	y += 1
	y *= x
	z += y
	w = inp[i]
	i++
	x *= 0
	x += z
	x = mod(x, 26)
	z /= 1
	x += 11
	x = eql(x, w)
	x = eql(x, 0)
	y *= 0
	y += 25
	y *= x
	y += 1
	z *= y
	y *= 0
	y += w
	y += 2
	y *= x
	z += y
	w = inp[i]
	i++
	x *= 0
	x += z
	x = mod(x, 26)
	z /= 26
	x += -5
	x = eql(x, w)
	x = eql(x, 0)
	y *= 0
	y += 25
	y *= x
	y += 1
	z *= y
	y *= 0
	y += w
	y += 4
	y *= x
	z += y
	w = inp[i]
	i++
	x *= 0
	x += z
	x = mod(x, 26)
	z /= 1
	x += 14
	x = eql(x, w)
	x = eql(x, 0)
	y *= 0
	y += 25
	y *= x
	y += 1
	z *= y
	y *= 0
	y += w
	y += 15
	y *= x
	z += y
	w = inp[i]
	i++
	x *= 0
	x += z
	x = mod(x, 26)
	z /= 1
	x += 15
	x = eql(x, w)
	x = eql(x, 0)
	y *= 0
	y += 25
	y *= x
	y += 1
	z *= y
	y *= 0
	y += w
	y += 11
	y *= x
	z += y
	w = inp[i]
	i++
	x *= 0
	x += z
	x = mod(x, 26)
	z /= 26
	x += -13
	x = eql(x, w)
	x = eql(x, 0)
	y *= 0
	y += 25
	y *= x
	y += 1
	z *= y
	y *= 0
	y += w
	y += 5
	y *= x
	z += y
	w = inp[i]
	i++
	x *= 0
	x += z
	x = mod(x, 26)
	z /= 26
	x += -16
	x = eql(x, w)
	x = eql(x, 0)
	y *= 0
	y += 25
	y *= x
	y += 1
	z *= y
	y *= 0
	y += w
	y += 3
	y *= x
	z += y
	w = inp[i]
	i++
	x *= 0
	x += z
	x = mod(x, 26)
	z /= 26
	x += -8
	x = eql(x, w)
	x = eql(x, 0)
	y *= 0
	y += 25
	y *= x
	y += 1
	z *= y
	y *= 0
	y += w
	y += 9
	y *= x
	z += y
	w = inp[i]
	i++
	x *= 0
	x += z
	x = mod(x, 26)
	z /= 1
	x += 15
	x = eql(x, w)
	x = eql(x, 0)
	y *= 0
	y += 25
	y *= x
	y += 1
	z *= y
	y *= 0
	y += w
	y += 2
	y *= x
	z += y
	w = inp[i]
	i++
	x *= 0
	x += z
	x = mod(x, 26)
	z /= 26
	x += -8
	x = eql(x, w)
	x = eql(x, 0)
	y *= 0
	y += 25
	y *= x
	y += 1
	z *= y
	y *= 0
	y += w
	y += 3
	y *= x
	z += y
	w = inp[i]
	i++
	x *= 0
	x += z
	x = mod(x, 26)
	z /= 26
	x += 0
	x = eql(x, w)
	x = eql(x, 0)
	y *= 0
	y += 25
	y *= x
	y += 1
	z *= y
	y *= 0
	y += w
	y += 3
	y *= x
	z += y
	w = inp[i]
	i++
	x *= 0
	x += z
	x = mod(x, 26)
	z /= 26
	x += -4
	x = eql(x, w)
	x = eql(x, 0)
	y *= 0
	y += 25
	y *= x
	y += 1
	z *= y
	y *= 0
	y += w
	y += 11
	y *= x
	z += y

	return
}

func Part1() int {

	inp := []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}
	// inp := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 9, 9, 9, 9, 9}
	for {
		w, x, y, z, i := 0, 0, 0, 0, 0

		w, x, y, z, i = run(w, x, y, z, 0, inp)

		if i != len(inp) {
			fmt.Printf("i=%d len=%d\n", i, len(inp))
			panic("i != len(inp)")
		}
		// fmt.Printf("inp=%v\n", inp)
		// fmt.Printf("w: %d, x: %d, y: %d, z: %d, i: %d\n", w, x, y, z, i)

		if z == 0 {
			fmt.Printf("inp=%v\n", inp)
			break
		}
		inp = decrement(inp, len(inp)-1)
	}

	return 0
}

func Part2() int {
	return 0

}

func main() {
	fmt.Println("--2021 day 24 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1())
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2())
	fmt.Println(time.Since(start))
}
