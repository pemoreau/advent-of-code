# Advent Of Code 2022

# Comments

## [Day 01: Calorie Counting](https://adventofcode.com/2022/day/1)

Example of input:

```
1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
```

### [Go](./go/2022/01/day01.go)
Nothing special.

## [Day 02: Rock Paper Scissors](https://adventofcode.com/2022/day/2)

Example of input:

```
A Y
B X
C Z
```

### [Go](./go/2022/02/day02.go)
Convert letters into numbers 0, 1, and 2. Then use the modulo operator.


## [Day 03: Rucksack Reorganization](https://adventofcode.com/2022/day/3)

Example of input:

```
vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
```

### [Go](./go/2022/03/day03.go)
Use sets and intersections

## [Day 04: Camp Cleanup](https://adventofcode.com/2022/day/4)

Example of input:

```
2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
```

### [Go](./go/2022/04/day04.go)
Nothing special

## [Day 05: Supply Stacks](https://adventofcode.com/2022/day/5)

Example of input:

```
    [D]
[N] [C]
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
```

### [Go](./go/2022/05/day05.go)
Use +4 offset to collect columns and build stacks

Use a hand made parser for instructions (much more efficient than `sscanf`)

Use `pushN`, `popN` and `reverse` to move stacks

## [Day 06: Tuning Trouble](https://adventofcode.com/2022/day/6)

Example of input:

```
bvwbjplbgvbhsrlpgdmjqwftvncz
```

### [Go](./go/2022/06/day06.go)
Use an array to check the all-diff constraint

## [Day 07: No Space Left On Device](https://adventofcode.com/2022/day/7)

Example of input:

```
$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
```

### [Go](./go/2022/07/day07.go)
Not very proud of my solution.

Got a bit confused when combining pointers and interface

## [Day 08: Treetop Tree House](https://adventofcode.com/2022/day/8)

Example of input:

```
30373
25512
65332
33549
35390
```

### [Go](./go/2022/08/day08.go)
Not fan of my solution

## [Day 09: Rope Bridge](https://adventofcode.com/2022/day/9)

Example of input:

```
R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2
```

### [Go](./go/2022/09/day09.go)

The rope is represented by an array of positions `[]Pos`, while the path is represented by a `Set[Pos]`

## [Day 10: Cathode-Ray Tube](https://adventofcode.com/2022/day/10)

Example of input:

```
noop
addx 3
addx -5
```

### [Go](./go/2022/10/day10.go)
Use an interface to write the simulator once and use it for both parts

## [Day 11: Monkey in the Middle](https://adventofcode.com/2022/day/11)

Example of input:

```
Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

```

### [Go](./go/2022/11/day11.go)

Used function combinators to implement monkeys

Skip parsing

## [Day 12: Hill Climbing Algorithm](https://adventofcode.com/2022/day/12)

Example of input:

```
Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi
```

### [Go](./go/2022/12/day12.go)
Used (overkill) A*, because it was already implemented

For part 2 I considered that all 'a' are neighbors of 'a' (with cost 0).
The elevation's difference can be used as heuristic



## [Day 13: ](https://adventofcode.com/2022/day/13)

Example of input:

```
[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]
```

### [Go](./go/2022/13/day13.go)


## [Day 14: ](https://adventofcode.com/2022/day/14)

Example of input:

```
498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9
```

### [Go](./go/2022/14/day14.go)


## [Day 15: ](https://adventofcode.com/2022/day/15)

Example of input:

```
Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3
```

### [Go](./go/2022/15/day15.go)

Nice problem!

I have use a list of disjoint intervals to represent empty space for a given horizontal line `y`

Given `y`, each interval corresponds to `[s.x - r, s.x + r]` where `r` is the distance from the sensor to the line `y`

To solve part 2, the intervals are merge into a unique one for each `y`. When this results in more than a unique interval, a solution is found

Use go routines to speed up the computation :-)

## [Day 16: ](https://adventofcode.com/2022/day/16)

Example of input:

```

```

### [Go](./go/2022/16/day16.go)


## [Day 17: ](https://adventofcode.com/2022/day/17)

Example of input:

```

```

### [Go](./go/2022/17/day17.go)


## [Day 18: ](https://adventofcode.com/2022/day/18)

Example of input:

```

```

### [Go](./go/2022/18/day18.go)


## [Day 19: ](https://adventofcode.com/2022/day/19)

Example of input:

```

```

### [Go](./go/2022/19/day19.go)


## [Day 20: ](https://adventofcode.com/2022/day/20)

Example of input:

```

```

### [Go](./go/2022/20/day20.go)


## [Day 21: ](https://adventofcode.com/2022/day/21)

Example of input:

```

```

### [Go](./go/2022/21/day21.go)


## [Day 22: ](https://adventofcode.com/2022/day/22)

Example of input:

```

```

### [Go](./go/2022/22/day22.go)


## [Day 23: ](https://adventofcode.com/2022/day/23)

Example of input:

```

```

### [Go](./go/2022/23/day23.go)


## [Day 24: ](https://adventofcode.com/2022/day/24)

Example of input:

```

```

### [Go](./go/2022/24/day24.go)


## [Day 25: ](https://adventofcode.com/2022/day/25)

Example of input:

```

```

### [Go](./go/2022/25/day25.go)

