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

I have used a list of disjoint intervals to represent empty space for a given horizontal line `y`

Given `y`, each interval corresponds to `[s.x - r, s.x + r]` where `r` is the distance from the sensor to the line `y`

To solve part 2, the intervals are merged into a unique one for each `y`. When this results in more than a unique interval, a solution is found

Used go routines to speed up the computation :-)

## [Day 16: Proboscidea Volcanium](https://adventofcode.com/2022/day/16)

Example of input:

```
Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
Valve BB has flow rate=13; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=22; tunnel leads to valve GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=21; tunnel leads to valve II
```

### [Go](./go/2022/16/day16.go)

The problem was quite difficult until finding a nice way to model it

Used Floyd–Warshall algorithm to compute the shortest path between 
each pair of active valves

Then used a simple BFS to compute all possible paths

For part 2, for each path found by part 1, I search another path which
does not intersect with the first one. 
The max is the sum of the two max

## [Day 17: Pyroclastic Flow](https://adventofcode.com/2022/day/17)

Example of input:

```
####

.#.
###
.#.

..#
..#
###

#
#
#
#

##
```

### [Go](./go/2022/17/day17.go)

Fun Tetris problem!

Part 1 is not too difficult, but part 2 is a bit tricky

For part 2, we keep in a list, the difference of maximal height each time a rock reach the ground.

Then, in this list of integers we look for a cycle, using a naive algorithm (but Floyd Tortoise algorithm should have
been used instead)

From that list, the result can be easily found

## [Day 18: Boiling Boulders](https://adventofcode.com/2022/day/18)

Example of input:

```
2,2,2
1,2,2
3,2,2
2,1,2
2,3,2
2,2,1
2,2,3
2,2,4
2,2,6
1,2,5
3,2,5
2,1,5
2,3,5
```

### [Go](./go/2022/18/day18.go)

Minecraft problem

Not too difficult: collect free space of the bounding box
then substract the number of rocks

## [Day 19: Not Enough Minerals](https://adventofcode.com/2022/day/19)

Example of input:

```
Blueprint 1:
  Each ore robot costs 4 ore.
  Each clay robot costs 2 ore.
  Each obsidian robot costs 3 ore and 14 clay.
  Each geode robot costs 2 ore and 7 obsidian.

Blueprint 2:
  Each ore robot costs 2 ore.
  Each clay robot costs 3 ore.
  Each obsidian robot costs 3 ore and 8 clay.
  Each geode robot costs 3 ore and 12 obsidian.
```

### [Go](./go/2022/19/day19.go)

Another difficult problem.

I have used an order to compare collected resources and removed redundant states.
I.e. those who have the same amount of robots but fewer collected resources

To be efficient, some smart tests have to be implemented to cut down the search space

## [Day 20: Grove Positioning System](https://adventofcode.com/2022/day/20)

Example of input:

```
1
2
-3
3
-2
0
4
```

### [Go](./go/2022/20/day20.go)

Used `container/ring`

## [Day 21: Monkey Math](https://adventofcode.com/2022/day/21)

Example of input:

```
root: pppw + sjmn
dbpl: 5
cczh: sllz + lgvd
zczc: 2
ptdq: humn - dvpt
dvpt: 3
lfqf: 4
humn: 5
ljgn: 2
sjmn: drzm * dbpl
sllz: 4
pppw: cczh / lfqf
lgvd: ljgn * ptdq
drzm: hmdt - zczc
hmdt: 32
```

### [Go](./go/2022/21/day21.go)

I have used a full symbolic approach to compute the value of `root` 
and to isolate the variable `humn`

## [Day 22: Monkey Map](https://adventofcode.com/2022/day/22)

Example of input:

```
        ...#
        .#..
        #...
        ....
...#.......#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.

10R5L5R10L4R5L5
```

### [Go](./go/2022/22/day22.go)

Not too difficult but a bit boring

## [Day 23: Unstable Diffusion](https://adventofcode.com/2022/day/23)

Example of input:

```
....#..
..###.#
#...#.#
.#...##
#.###..
##.#.##
.#..#..
```

### [Go](./go/2022/23/day23.go)

Nothing special

## [Day 24: Blizzard Basin](https://adventofcode.com/2022/day/24)

Example of input:

```
#.######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#
```

### [Go](./go/2022/24/day24.go)
Quite interesting.

Used a generic A* (using Go generics)

Blizzards are precalculated since they repeat every LCM(height,width) steps

## [Day 25: Full of Hot Air](https://adventofcode.com/2022/day/25)

Example of input:

```
1=-0-2
12111
2=0=
21
2=01
111
20012
112
1=-1=
1-12
12
1=
122
```

### [Go](./go/2022/25/day25.go)

Funny
