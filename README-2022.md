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

## [Day 09: ](https://adventofcode.com/2022/day/9)

Example of input:

```

```

### [Go](./go/2022/09/day09.go)

## [Day 10: ](https://adventofcode.com/2022/day/10)

Example of input:

```

```

### [Go](./go/2022/10/day10.go)


## [Day 11: ](https://adventofcode.com/2022/day/11)

Example of input:

```

```

### [Go](./go/2022/11/day11.go)


## [Day 12: ](https://adventofcode.com/2022/day/12)

Example of input:

```

```

### [Go](./go/2022/12/day12.go)

## [Day 13: ](https://adventofcode.com/2022/day/13)

Example of input:

```

```

### [Go](./go/2022/13/day13.go)


## [Day 14: ](https://adventofcode.com/2022/day/14)

Example of input:

```

```

### [Go](./go/2022/14/day14.go)


## [Day 15: ](https://adventofcode.com/2022/day/15)

Example of input:

```

```

### [Go](./go/2022/15/day15.go)


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

