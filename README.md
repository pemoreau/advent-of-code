# Advent Of Code 2021

Learning Rust and Go

# Commands

## Rust

- `cargo run <n>` to run the n-th puzzle (example: `cargo run 1`)
- `cargo run --release <n>` to run the n-th puzzle in release mode
- `cargo test` to run unit tests

## Go

- `cd src/go07` and `go run day07.go `
- `go test -v` to run unit tests

# Timings

Execution time on an old Mac Pro (Late 2013), 3,7 GHz Quad-Core Intel Xeon E5

| Rust                          | part A     | part B     | Go                                      | part A     | part B     |
| :---------------------------- | :--------- | :--------- | --------------------------------------- | ---------- | ---------- |
| [day 01](./rust/src/day01.rs) | ` 0.089ms` | ` 0.067ms` |                                         |            |            |
| [day 02](./rust/src/day02.rs) | ` 0.092ms` | ` 0.063ms` |                                         |            |            |
| [day 03](./rust/src/day03.rs) | ` 0.157ms` | ` 0.084ms` |                                         |            |            |
| [day 04](./rust/src/day04.rs) | ` 1.048ms` | ` 0.841ms` |                                         |            |            |
| [day 05](./rust/src/day05.rs) | ` 45.94ms` | ` 46.03ms` |                                         |            |            |
| [day 06](./rust/src/day06.rs) | ` 0.010ms` | ` 0.008ms` | [day 06](./go/go06/day06.go)            | ` 0.061ms` | ` 0.016ms` |
| [day 07](./rust/src/day07.rs) | ` 0.274ms` | ` 0.795ms` | [day 07](./go/go07/day07.go)            | ` 4.713ms` | ` 5.214ms` |
| [day 08](./rust/src/day08.rs) | ` 0.198ms` | ` 1.786ms` |                                         |            |            |
|                               |            |            | [day 09](./go/go09_simplified/day09.go) | ` 0.189ms` | ` 1.183ms` |
| [day 10](./rust/src/day10.rs) | ` 0.137ms` | ` 0.134ms` | [day 10](./go/go10/day10.go)            | ` 0.152ms` | ` 0.151ms` |
| [day 11](./rust/src/day11.rs) | ` 0.186ms` | ` 0.420ms` | [day 11](./go/go11/day11.go)            | ` 0.211`   | ` 0.422ms` |
|                               |            |            | [day 12](./go/go12/day12.go)            | ` 0.161ms` | ` 3.944ms` |
| [day 13](./rust/src/day13.rs) | ` 0.156ms` | ` 0.118ms` | [day 13](./go/go13/day13.go)            | ` 0.441ms` | ` 0.706ms` |
|                               |            |            | [day 14](./go/go14/day14.go)            | ` 0.075ms` | ` 0.056ms` |
|                               |            |            | [day 15](./go/go15/day15.go)            | ` 11.64ms` | ` 344.0ms` |
|                               |            |            | [day 16](./go/go16/day16.go)            | ` 0.121ms` | ` 0.071ms` |
|                               |            |            | [day 17](./go/17/day17.go)              | ` 0.133ms` | ` 0.424ms` |
|                               |            |            | [day 18](./go/18/day18.go)              | ` 4.071ms` | ` 24.39ms` |
|                               |            |            | [day 19](./go/19/day19.go)              |            |            |
|                               |            |            | [day 20](./go/20/day20.go)              | ` 13.37ms` | ` 683.0ms` |

# Comments

## Day 01

### Rust

nothing special except the use of `windows` function

## Day 02

### Rust

use `split_one` instead of regex to speed-up parsing

use `for-loop` style and then `fold` for the second part

## Day 03

### Rust

nothing special

## Day 04

### Rust

use `split("\n\n")` to separate parts (instead of counting the number of entries)

use `array2d::Array2D` to represent the board but this may be not the best choice

the program contains too many `for-loop` in my opinion

## Day 05

### Rust

a bit slow due the the use of regex for parsing

discovered the `signum` function

use references to mutable structures

## Day 06

### Rust

very simple solution thanks to `rotate_left` function

## Day 07

### Rust

use a cost function as parameter

## Day 08

### Rust

use a HashMap with sorted letters as keys

## Day 09

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

`[]int{}` and `make([]int, 0)` are equivalent but the later can be used to defined the capacity of the underlaying array. For instance: `make([]int, 0, len(collectedBassin))`.

## Day 10

### Go

Use a stack for parsing the input

## Day 11

### Go

Use 2d array to represent the board. The code involves many nested `for-loop` but it is quite natural in Golang

### Rust

Also used 2d array in Rust. The code is similar to the one written in Go, making the result a bit awkward in Rust.

Tried to represent the board using a 1d-array, but this did not improved the code so much.

I am not very satisfied by this solution.

## Day 12

### Go

I have lost a lot of time because I did not immediately understood the second part.

I first came with a solution that builds the list of paths, but it was a bit slow.

I then simplified the code to get a more efficient solution (~3.8ms for part 2 on my 2013 mac).

## Day 13

### Rust

Used a list of tuples to avoid creating a 2d-array

But the second part needed the construction of such a 2d-array

### Go

Done in Go after the first Rust implementation.

Use a set (`map[Pos]struct{}`) to store the positions

And then a 2d-array to display the result

Once again I found the Go version easier to write

I have written a code of the form:

```
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

### Go

I have used a 2d-array for the rules (`[26][26]byte`) and a 2d-array for counting the pairs (`[26][26]int{}`). Maybe a `map` would have been more efficient. I did not have time to compare the two solutions.

Also used the following construct to statically include the input file :

```
//go:embed input.txt
var input string
```

## Day 15

### Go

Lost a lot of time because I made a mistake when building the mega-matrix but I have discovered a very nice website: https://www.redblobgames.com/

And in partiular articles from Amit (https://theory.stanford.edu/~amitp/GameProgramming/)

## Day 16

### Go

Not too difficult today but I had several problems with Golang.
My function `extract` takes a position (`index`) as an argument and returns the new value for this index.

I would have liked to write my code as follows:

```
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

```
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

```
    for index < end {
		var res Packet
		res, index = decode(bytes, index)
		packets = append(packets, res)
	}
```

where I cannot use the short declaration syntax for `res, index :=` becasue the `index` variable should not be shallowed.

The second part of the problem is very interesting because we have to find a way to construct an expression-tree and to define evaluation functions.

In general we can use inheritance. In Golang I have defined structs and an interface (`Packet`).

As explained here: https://eli.thegreenplace.net/2018/go-and-algebraic-data-types/, a key point is to define a function (`isTree` for example) which is common to all struct that belong to the interface.

In a second step, the type discrimination (instead of dynamic dispathing) is done by the `switch` statement and the use of `t.(type)` construct

After finishing the puzzle I discovered the [bitio](https://github.com/icza/bitio) library. I will try it when I have time.

## Day 17

### Go

Brute force approach. Not very proud of it.

## Day 18

### Go

A good day for me. I have used a list of (value, depth) tuples to represent the tree. With this representation, the normalization wrt. explode can be done in one pass. `split` rule is also efficient.

One difficulty is to compute the magnitude. Fopr that I use a stack of `(value,depth)`. when the top of the stack contains two values with the same depth, theyn can be replaced by a new tuple `(3*left+2*right, depth-1)`. This is quite efficient.

The current implementation doesn ot perform side effect. This should be possible to improve the efficiency by doing transformations in place.

## Day 19

### Go

not a good day for me

## Day 20

### Go

I have use a hashset to store the positions. I had to store the `.` positions so this is not very optimized.
The interest is that I can extend the border of the image without moviing everything, but it is possible that an array of pixels would be more efficient
