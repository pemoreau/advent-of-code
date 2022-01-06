# Advent Of Code 2021

Learning Rust and Go

# Commands

## Rust

- `cargo run <n>` to run the n-th puzzle (example: `cargo run 1`)
- `cargo run --release <n>` to run the n-th puzzle in release mode
- `cargo test` to run unit tests

## Go

- `cd go/<n>` and `go run .`
- `go test` to run unit tests
- `go test -bench .` to run unit tests

# Timings

Execution time on an old Mac Pro (Late 2013), 3,7 GHz Quad-Core Intel Xeon E5

| Rust                          | part A     | part B     | Go                                    | part A      | part B        |
| :---------------------------- | :--------- | :--------- | ------------------------------------- | ----------- | ------------- |
| [day 01](./rust/src/day01.rs) | ` 0.089ms` | ` 0.067ms` | [day 01](./go/01/day01.go)            | ` 0.047 ms` | ` 0.048 ms`   |
| [day 02](./rust/src/day02.rs) | ` 0.092ms` | ` 0.063ms` | [day 02](./go/02/day02.go)            | ` 0.102 ms` | ` 0.103 ms`   |
| [day 03](./rust/src/day03.rs) | ` 0.157ms` | ` 0.084ms` |                                       |             |               |
| [day 04](./rust/src/day04.rs) | ` 1.048ms` | ` 0.841ms` |                                       |             |               |
| [day 05](./rust/src/day05.rs) | ` 45.94ms` | ` 46.03ms` |                                       |             |               |
| [day 06](./rust/src/day06.rs) | ` 0.010ms` | ` 0.008ms` | [day 06](./go/06/day06.go)            | ` 0.007 ms` | ` 0.008 ms`   |
| [day 07](./rust/src/day07.rs) | ` 0.274ms` | ` 0.795ms` | [day 07](./go/07/day07.go)            | ` 1.711 ms` | ` 2.841 ms`   |
| [day 08](./rust/src/day08.rs) | ` 0.198ms` | ` 1.786ms` |                                       |             |               |
|                               |            |            | [day 09](./go/09_simplified/day09.go) | ` 0.146 ms` | ` 0.670 ms`   |
| [day 10](./rust/src/day10.rs) | ` 0.137ms` | ` 0.134ms` | [day 10](./go/10/day10.go)            | ` 0.158 ms` | ` 0.160 ms`   |
| [day 11](./rust/src/day11.rs) | ` 0.186ms` | ` 0.420ms` | [day 11](./go/11/day11.go)            | ` 0.152 ms` | ` 0.432 ms`   |
|                               |            |            | [day 12](./go/12/day12.go)            | ` 0.161 ms` | ` 3.944 ms`   |
| [day 13](./rust/src/day13.rs) | ` 0.156ms` | ` 0.118ms` | [day 13](./go/13/day13.go)            | ` 0.441 ms` | ` 0.706 ms`   |
|                               |            |            | [day 14](./go/14/day14.go)            | ` 0.075 ms` | ` 0.056 ms`   |
|                               |            |            | [day 15](./go/15/day15.go)            | ` 11.64 ms` | ` 344.0 ms`   |
|                               |            |            | [day 16](./go/16/day16.go)            | ` 0.121 ms` | ` 0.071 ms`   |
|                               |            |            | [day 17](./go/17/day17.go)            | ` 0.133 ms` | ` 0.424 ms`   |
|                               |            |            | [day 18](./go/18/day18.go)            | ` 4.071 ms` | ` 24.39 ms`   |
|                               |            |            | [day 19](./go/19/day19.go)            | ` 20.20 ms` | ` 20.88 ms`   |
|                               |            |            | [day 20](./go/20/day20.go)            | ` 9.035 ms` | ` 491.576 ms` |
|                               |            |            | [day 21](./go/21/day21.go)            | ` 2.342 µs` | ` 137.152 ms` |
|                               |            |            | [day 22](./go/22/day22.go)            | ` 2.237 ms` | ` 56.162 ms`  |
|                               |            |            | day 23                                |             |               |
|                               |            |            | [day 24](./go/24/day24.go)            | ` 101.5 s`  | ` 0.003 ms`   |
|                               |            |            | [day 25](./go/25/day25.go)            | ` 98.0 ms`  | ` 0.003 ms`   |

# Comments

## Day 01: Sonar Sweep

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

### Rust

Nothing special except the use of `windows` function

### Go

Used the following construct to embed the input file as a string in the source code:

```go
//go:embed input.txt
var input_day string
```

## Day 02: Dive!

Example of input:

```
forward 5
down 5
forward 8
up 3
down 8
forward 2
```

### Rust

Used `split_one` instead of regex to speed-up parsing

Used `for-loop` style and then `fold` for the second part

### Go

Nothing special

## Day 03: Binary Diagnostic

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

### Rust

Nothing special

## Day 04: Giant Squid

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

### Rust

Used `split("\n\n")` to separate parts (instead of counting the number of entries)

Used `array2d::Array2D` to represent the board but this may be not the best choice

The program contains too many `for-loop` in my opinion

## Day 05: Hydrothermal Venture

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

### Rust

A bit slow due the the use of regex for parsing

Discovered the `signum` function

Used references to mutable structures

## Day 06: Lanternfish

Example of input:

```
3,4,3,1,2
```

### Rust

Very simple solution thanks to `rotate_left` function

### Go

Used `mult = append(mult[1:], mult[0])` to rotate left the slice

## Day 07: The Treachery of Whales

Example of input:

```
16,1,2,0,4,2,7,1,2,14
```

### Rust

Used a cost function as parameter

### Go

Using a cost function as parameter is 2 to 3 times slower than coding the cost expression directly in the for-loop

Tried to use `int32` instead of `int` but it was slower

## Day 08: Seven Segment Search

Example of input:

```
acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab |
cdfeb fcadb cdfeb cdbaf
```

### Rust

use a HashMap with sorted letters as keys

## Day 09: Smoke Basin

Example of input:

```
2199943210
3987894921
9856789892
8767896789
9899965678
```

### Go

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

## Day 10: Syntax Scoring

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

### Go

Use a stack for parsing the input

## Day 11

Example of input:

```

```

### Go

Use 2d array to represent the board. The code involves many nested `for-loop` but it is quite natural in Golang

### Rust

Also used 2d array in Rust. The code is similar to the one written in Go, making the result a bit awkward in Rust.

Tried to represent the board using a 1d-array, but this did not improved the code so much.

I am not very satisfied by this solution.

## Day 12

Example of input:

```

```

### Go

I have lost a lot of time because I did not immediately understood the second part.

I first came with a solution that builds the list of paths, but it was a bit slow.

I then simplified the code to get a more efficient solution (~3.8ms for part 2 on my 2013 mac).

## Day 13

Example of input:

```

```

### Rust

Used a list of tuples to avoid creating a 2d-array

But the second part needed the construction of such a 2d-array

### Go

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

## Day 14

Example of input:

```

```

### Go

I have used a 2d-array for the rules (`[26][26]byte`) and a 2d-array for counting the pairs (`[26][26]int{}`). Maybe a `map` would have been more efficient. I did not have time to compare the two solutions.

Also used the following construct to statically include the input file :

```go
//go:embed input.txt
var input string
```

## Day 15

Example of input:

```

```

### Go

Lost a lot of time because I made a mistake when building the mega-matrix but I have discovered a very nice website: https://www.redblobgames.com/

And in particular articles from Amit (https://theory.stanford.edu/~amitp/GameProgramming/)

## Day 16

Example of input:

```

```

### Go

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

As explained here: https://eli.thegreenplace.net/2018/go-and-algebraic-data-types/, a key point is to define a function (`isTree` for example) which is common to all struct that belong to the interface.

In a second step, the type discrimination (instead of dynamic dispatching) is done by the `switch` statement and the use of `t.(type)` construct

After finishing the puzzle I discovered the [bitio](https://github.com/icza/bitio) library. I will try it when I have time.

## Day 17

Example of input:

```

```

### Go

Brute force approach. Not very proud of it.

## Day 18

Example of input:

```

```

### Go

A good day for me. I have used a list of (value, depth) tuples to represent the tree. With this representation, the normalization wrt. explode can be done in one pass. `split` rule is also efficient.

One difficulty is to compute the magnitude. For that I use a stack of `(value,depth)`. when the top of the stack contains two values with the same depth, they can be replaced by a new tuple `(3*left+2*right, depth-1)`. This is quite efficient.

The current implementation does not perform side effect. This should be possible to improve the efficiency by doing transformations in place.

## Day 19

Example of input:

```

```

### Go

not a good day for me

## Day 20:: Trench Map

Example of input:

```

```

### Go

I have use a hashset to store the positions. I had to store the `.` positions so this is not very optimized.
The interest is that I can extend the border of the image without moving everything, but it is possible that an array of pixels would be more efficient

## Day 21: Dirac Dice

Example of input:

```

```

### Go

Using recursion and a cache.

Not so easy to get it correct. I made a mistake: starting with `uint8` for space. I took me some time to the overflow problem.

This is a good lesson: use `int` instead of `int8`, `int16`, `uint8`, ...,and do not do premature optimization. Then implement unit tests, and only after that, narrow integer types to speedup.

## Day 22: Reactor Reboot

Example of input:

```

```

### Go

Implemented a Cuboid data-structure with an `overlap(c1,c2 Cuboid)` function (which splits `c1` into smaller disjoint ones when `c2` overlaps)

## Day 23: Amphipod

Example of input:

```
#############
#...........#
###B#C#B#D###
  #A#D#C#A#
  #########
```

Not yet implemented. Solved it using a spreadsheet.

## Day 24: Arithmetic Logic Unit

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

### Go

First tried brute force (using a compiled approach) but it is way too slow.

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

## Day 25: Sea Cucumber

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

### Go

Nothing special: just uses a 2d-array of bytes and a step function to simulate the movement.
