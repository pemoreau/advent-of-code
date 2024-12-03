# Advent Of Code 2021

Learning Rust and Go

# Comments

## [Day 01: Sonar Sweep](https://adventofcode.com/2021/day/1)

Example of input:

```
199
200
208
210
200
207
240
269
260
263
```

### [Rust](./rust/2021/day01)

Nothing special except the use of `windows` function

### [Go](./go/2021/01/day01.go)

Used the following construct to embed the input file as a string in the source code:

```go
//go:embed sample.txt
var inputTest string
```

## [Day 02: Dive!](https://adventofcode.com/2021/day/2)

Example of input:

```
forward 5
down 5
forward 8
up 3
down 8
forward 2
```

### [Rust](./rust/2021/day02)

Used `split_one` instead of regex to speed-up parsing

Used `for-loop` style and then `fold` for the second part

### [Go](./go/2021/02/day02.go)

Nothing special

## [Day 03: Binary Diagnostic](https://adventofcode.com/2021/day/3)

Example of input:

```
00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010
```

### [Rust](./rust/2021/day03)

Nothing special

## [Day 04: Giant Squid](https://adventofcode.com/2021/day/4)

Example of input:

```
7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7
```

### [Rust](./rust/2021/day04)

Used `split("\n\n")` to separate parts (instead of counting the number of entries)

Used `array2d::Array2D` to represent the board but this may be not the best choice

The program contains too many `for-loop` in my opinion

## [Day 05: Hydrothermal Venture](https://adventofcode.com/2021/day/5)

Example of input:

```
0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2
```

### [Rust](./rust/2021/day05)

A bit slow due to the use of regex for parsing

Discovered the `signum` function

Used references to mutable structures

## [Day 06: Lanternfish](https://adventofcode.com/2021/day/6)

Example of input:

```
3,4,3,1,2
```

### [Rust](./rust/2021/day06)

Very simple solution thanks to `rotate_left` function

### [Go](./go/2021/06/day06.go)

Used `mult = append(mult[1:], mult[0])` to rotate left the slice

## [Day 07: The Treachery of Whales](https://adventofcode.com/2021/day/7)

Example of input:

```
16,1,2,0,4,2,7,1,2,14
```

### [Rust](./rust/2021/day07)

Used a cost function as parameter

### [Go](./go/2021/07/day07.go)

Using a cost function as parameter is 2 to 3 times slower than coding the cost expression directly in the for-loop

Tried to use `int32` instead of `int` but it was slower

## [Day 08: Seven Segment Search](https://adventofcode.com/2021/day/8)

Example of input:

```
acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab |
cdfeb fcadb cdfeb cdbaf
```

### [Rust](./rust/2021/day08)

use a HashMap with sorted letters as keys

## [Day 09: Smoke Basin](https://adventofcode.com/2021/day/9)

Example of input:

```
2199943210
3987894921
9856789892
8767896789
9899965678
```

### [Go](./go/2021/09_simplified/day09.go)

explore search space using a set of visited positions (`type set map[Pos]struct{}`), with positions defined as

```go
type Pos struct {
	i, j int
}
```

Today I learned that `struct` are always passed by value (i.e. function parameters are mutable copies), but reference types (which correspond to slice, map, channel, interface, and function types) are passed by reference.

To be more precise, slice and map are passed by value (like any other value in Go), but these values are references to underlying data-structures.

This is why `matrix` and `set` can be passed without `*` as function arguments, and still be mutable.

See [Dave Cheney's article](https://dave.cheney.net/2017/04/30/if-a-map-isnt-a-reference-variable-what-is-it) for more explanations

`[]int{}` and `make([]int, 0)` are equivalent but the later can be used to defined the capacity of the underlying array. For instance: `make([]Pos, 0, 4)`.

## [Day 10: Syntax Scoring](https://adventofcode.com/2021/day/10)

Example of input:

```
[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]
```

### [Go](./go/2021/10/day10.go)

Use a stack for parsing the input

### [Rust](./rust/2021/day10)

Nothing special

## [Day 11: Dumbo Octopus](https://adventofcode.com/2021/day/11)

Example of input:

```
5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526
```

### [Go](./go/2021/11/day11.go)

Used 2d array to represent the board. The code involves many nested `for-loop` but it is quite natural in Golang

### [Rust](./rust/2021/day11)

Also used 2d array in Rust. The code is similar to the one written in Go, making the result a bit awkward in Rust.

Tried to represent the board using a 1d-array, but this did not improved the code so much.

I am not very satisfied by this solution.

## [Day 12: Passage Pathing](https://adventofcode.com/2021/day/12)

Example of input:

```
start-A
start-b
A-c
A-b
b-d
A-end
b-end
```

### [Go](./go/2021/12/day12.go)

I have lost a lot of time because I did not immediately understood the second part.

I first came with a solution that builds the list of paths, but it was a bit slow.

I then simplified the code to get a more efficient solution (~3.8ms for part 2 on my 2013 mac).

## [Day 13: Transparent Origami](https://adventofcode.com/2021/day/13)

Example of input:

```
6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5
```

### [Rust](./rust/2021/day13)

Used a list of tuples to avoid creating a 2d-array

But the second part needed the construction of such a 2d-array

### [Go](./go/2021/13/day13.go)

Done in Go after the first Rust implementation.

Use a set (`map[Pos]struct{}`) to store the positions

And then a 2d-array to display the result

Once again I found the Go version easier to write

I have written a code of the form:

```go
func step(screen map[Pos]struct{}, inst Instr) {
	for p := range screen {
        ...
		screen.Add(Pos{2*d - p.x, p.y})
		screen.Remove(p)
    }
}
```

This shows that you can mutate a map while iterating over it.

## [Day 14: Extended Polymerization](https://adventofcode.com/2021/day/14)

Example of input:

```
NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C
```

### [Go](./go/2021/14/day14.go)

I have used a 2d-array for the rules (`[26][26]byte`) and a 2d-array for counting the pairs (`[26][26]int{}`). Maybe a `map` would have been more efficient. I did not have time to compare the two solutions.

## [Day 15: Chiton](https://adventofcode.com/2021/day/15)

Example of input:

```
1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581
```

### [Go](./go/2021/15/day15.go)

Lost a lot of time because I made a mistake when building the mega-matrix but I have discovered a very nice website: https://www.redblobgames.com/

And in particular articles from Amit (https://theory.stanford.edu/~amitp/GameProgramming/)

## [Day 16: Packet Decoder](https://adventofcode.com/2021/day/16)

Example of input:

```
9C0141080250320F1802104A08
```

### [Go](./go/2021/16/day16.go)

Not too difficult today but I had several problems with Golang.
My function `extract` takes a position (`index`) as an argument and returns the new value for this index.

I would have liked to write my code as follows:

```go
    next, index := bit(bytes, index)
	value, index = extract(bytes, index, 4)
	for next {
		next, index = bit(bytes, index)
		v, index := extract(bytes, index, 4)
		value = value<<4 + v
	}
```

but this is not possible because in the `for-loop`, using `v, index :=` declares a new `index` variable which shallows the previous one.

The solution is to declare `v` before the assignment:

```go
    next, index = bit(bytes, index)
	value, index = extract(bytes, index, 4)
	for next {
		next, index = bit(bytes, index)
		var v uint64
		v, index = extract(bytes, index, 4)
		value = value<<4 + v
	}
```

Another example is:

```go
    for index < end {
		var res Packet
		res, index = decode(bytes, index)
		packets = append(packets, res)
	}
```

where I cannot use the short declaration syntax for `res, index :=` because the `index` variable should not be shallowed.

The second part of the problem is very interesting because we have to find a way to construct an expression-tree and to define evaluation functions.

In general we can use inheritance. In Golang I have defined structs and an interface (`Packet`).

As explained [here](https://eli.thegreenplace.net/2018/go-and-algebraic-data-types/), a key point is to define a function (`isTree` for example) which is common to all struct that belong to the interface.

In a second step, the type discrimination (instead of dynamic dispatching) is done by the `switch` statement and the use of `t.(type)` construct

After finishing the puzzle I discovered the [bitio](https://github.com/icza/bitio) library. I will try it when I have time.

## [Day 17: Trick Shot](https://adventofcode.com/2021/day/17)

Example of input:

```
target area: x=20..30, y=-10..-5
```

### [Go](./go/2021/17/day17.go)

Brute force approach. Not very proud of it.

Used `regexp` to parse the input.

## [Day 18: Snailfish](https://adventofcode.com/2021/day/18)

Example of input:

```
[1,2]
[[1,2],3]
[9,[8,7]]
[[1,9],[8,5]]
[[[[1,2],[3,4]],[[5,6],[7,8]]],9]
[[[9,[3,8]],[[0,9],6]],[[[3,7],[4,9]],3]]
[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]
```

### [Go](./go/2021/18/day18.go)

A good day for me. I have used a list of `(value, depth)` tuples to represent the tree. With this representation, the normalization wrt. explode can be done in one pass. `split` rule is also efficient.

One difficulty is to compute the magnitude. For that I use a stack of `(value,depth)`. when the top of the stack contains two values with the same depth, they can be replaced by a new tuple `(3*left+2*right, depth-1)`. This is quite efficient.

The current implementation does not perform side effect. This should be possible to improve the efficiency by doing transformations in place.

## [Day 19: Beacon Scanner](https://adventofcode.com/2021/day/19)

Example of input:

```
--- scanner 0 ---
0,2
4,1
3,3

--- scanner 1 ---
-1,-1
-5,0
-2,1
```

### [Go](./go/2021/19/day19.go)

Not a good day for me

## [Day 20: Trench Map](https://adventofcode.com/2021/day/20)

Example of input:

```
..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..##
#..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###
.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#.
.#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#.....
.#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#..
...####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.....
..##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#

#..#.
#....
##..#
..#..
..###
```

### [Go](./go/2021/20/day20.go)

I have use a hashset to store the positions. I had to store the `.` positions so this is not very optimized.
The interest is that I can extend the border of the image without moving everything, but it is possible that an array of pixels would be more efficient

## [Day 21: Dirac Dice](https://adventofcode.com/2021/day/21)

Example of input:

```
Player 1 starting position: 4
Player 2 starting position: 8
```

### [Go](./go/2021/21/day21.go)

Using recursion and a cache.

Not so easy to get it correct. I made a mistake: starting with `uint8` for space. It took me some time to see the overflow problem.

This is a good lesson: use `int` instead of `int8`, `int16`, `uint8`, ...,and do not do premature optimization. Then implement unit tests, and only after that, narrow integer types to speedup.

## [Day 22: Reactor Reboot](https://adventofcode.com/2021/day/22)

Example of input:

```
on x=10..12,y=10..12,z=10..12
on x=11..13,y=11..13,z=11..13
off x=9..11,y=9..11,z=9..11
on x=10..10,y=10..10,z=10..1
```

### [Go](./go/2021/22/day22.go)

Implemented a Cuboid data-structure with an `overlap(c1,c2 Cuboid)` function (which splits `c1` into smaller disjoint ones when `c2` overlaps)

Used `fmt.Sscanf` to parse the input.

## [Day 23: Amphipod](https://adventofcode.com/2021/day/23)

Example of input:

```
#############
#...........#
###B#C#B#D###
  #A#D#C#A#
  #########
```

### [Go](./go/2021/23/day23.go)

I spent a lot of time on this solution.

First I tried to use a `map[Position]byte` to represent the game but it was a bit too slow, in particular when compared to [1e9y](https://github.com/1e9y/adventofcode/blob/main/2021/day23/day23.go) approach.

1e9y has a very simple representation (a string) with a function to convert a position `{x,y}` into an `int` index. This representation is very efficient for this problem. His code is very smart.

In my approach I use an A\* algorithm with a heuristic based on manhattan distance.

I have also noted that storing (instead of computing each time) the fact that an occupant is "at home" can help. I use a lowercase in my string based implementation.

With these optimisations, part 1 can be solved in 7ms, and part 2 in 70ms on a 2013 Mac.

## [Day 24: Arithmetic Logic Unit](https://adventofcode.com/2021/day/24)

Example of input:

```
inp w
add z w
mod z 2
div w 2
add y w
mod y 2
div w 2
add x w
mod x 2
div w 2
mod w 2
```

The goal is to find inputs (`w0`...`w13`) such that `z=0`.

### [Go](./go/2021/24/day24.go)

First tried brute force (using a compiled approach) but it was way too slow.

In a second attempt, I tried to use abstract interpretation to associate an
interval (min,max values) to each variable.
Unfortunately, this does not restrict the search space enough.
Then I wanted to use back-propagation (i.e. assign the interval `[0,0]` to `z`) and interpret the program starting from the last instruction to compute conditions that the input variable should satisfy.

Since I was not convince to get something at the end, I put some energy in the brute force approach.
First I used channels and go routines (using a producer-consumer pattern) to explore all the search space.
In a couple of minutes I got the lower bound but it took much more time (several days) to get the upper one.

So, I tried another approach: I used a hashmap to store all explored states (i.e. all `w, x, y, z` values). And to each state we associate the `min` and the `max` input values which lead to this state.
After each `inp` instruction, the search space is increase by a factor 9. Fortunately we can merge all the states whose values `w, x, y, z` are the same.

This approach solved the problems in ~100s, using 70.10^6 states.

Later, I reused the code using abstract interpretation to cut search branches when z cannot reach the value 0. This reduced the computation to less than 8 seconds.

I noted that the exploration can be done in parallel. I just spawn some goroutines and the time decreased to 2.5 seconds.

## [Day 25: Sea Cucumber](https://adventofcode.com/2021/day/25)

Example of input:

```
v...>>.vv>
.vv>>.vv..
>>.>v>...v
>>v>>.>.v.
v>v.vv.v..
>.>>..v...
.vv..>.>v.
v.v..>>v.v
....v..v.>
```

### [Go](./go/2021/25/day25.go)

Nothing special: just uses a 2d-array of bytes and a step function to simulate the movement.
