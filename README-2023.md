# Advent Of Code 2023

# Comments

## [Day 01: Trebuchet](https://adventofcode.com/2023/day/1)

Example of input:

```
1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
```

### [Go](./go/2023/01/day01.go)
Used a discrimination tree for efficiency

## [Day 02: Cube Conundrum](https://adventofcode.com/2023/day/2)

Example of input:

```
Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
```

### [Go](./go/2023/02/day02.go)
Nothing special

## [Day 03: Gear Ratios](https://adventofcode.com/2023/day/3)

Example of input:

```
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
```

### [Go](./go/2023/03/day03.go)

## [Day 04: Scratchcards](https://adventofcode.com/2023/day/4)

Example of input:

```
Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
```

### [Go](./go/2023/04/day04.go)
Not too difficult using set, and in particular [bitset](github.com/bits-and-blooms/bitset)

## [Day 05: If You Give A Seed A Fertilizer](https://adventofcode.com/2023/day/5)

Example of input:

```
seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
```

### [Go](./go/2023/05/day05.go)
Used a list of intervals for both parts

## [Day 06: Wait For It](https://adventofcode.com/2023/day/6)

Example of input:

```
Time:      7  15   30
Distance:  9  40  200
```

### [Go](./go/2023/06/day06.go)

## [Day 07: Camel Cards](https://adventofcode.com/2023/day/7)

Example of input:

```
32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
```

### [Go](./go/2023/07/day07.go)
Misread the problem and got a bit confused by the order when comparing two hands of the same kind


## [Day 08: Haunted Wasteland](https://adventofcode.com/2023/day/8)

Example of input:

```
RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)
```

### [Go](./go/2023/08/day08.go)
An AOC classic where LCM is needed

## [Day 09: Mirage Maintenance](https://adventofcode.com/2023/day/9)

Example of input:

```
0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
```

### [Go](./go/2023/09/day09.go)


## [Day 10: Pipe Maze](https://adventofcode.com/2023/day/10)

Example of input:

```
.....
.F-7.
.|.|.
.L-J.
.....
```

### [Go](./go/2023/10/day10.go)
Interresting. Used a naive ray-tracing technique.
Efficiency could be improved using Shoeslace formula and Pick's theorem (see day 18 below)

## [Day 11: Cosmic Expansion ](https://adventofcode.com/2023/day/11)

Example of input:

```
...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....
```

### [Go](./go/2023/11/day11.go)


## [Day 12: Hot Springs](https://adventofcode.com/2023/day/12)

Example of input:

```
#.#.### 1,1,3
.#...#....###. 1,1,3
.#.###.#.###### 1,3,1,6
####.#...#... 4,1,1
#....######..#####. 1,6,5
.###.##....# 3,2,1
```

### [Go](./go/2023/12/day12.go)
Used recursion and a cache.
A good point si to used immutable string for patterns (instead of lists)


## [Day 13: Point of Incidence](https://adventofcode.com/2023/day/13)

Example of input:

```
#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#
```

### [Go](./go/2023/13/day13.go)

Brute force to find smudges

## [Day 14: Parabolic Reflector Dish](https://adventofcode.com/2023/day/14)

Example of input:

```
O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....
```

### [Go](./go/2023/14/day14.go)
Nice. Have to find a cycle for second part

## [Day 15: Lens Library](https://adventofcode.com/2023/day/15)

Example of input:

```
rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7
```

### [Go](./go/2023/15/day15.go)

Implementation of an hashtable

## [Day 16: The Floor Will Be Lava](https://adventofcode.com/2023/day/16)

Example of input:

```
.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....
```

### [Go](./go/2023/16/day16.go)

Nothing special. Quite easy for a day 16

## [Day 17: Clumsy Crucible](https://adventofcode.com/2023/day/17)

Example of input:

```
2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533
```

### [Go](./go/2023/17/day17.go)

First use of Dijkstra/A* algorithm.
No too difficult when using a neighboor function which computes all reachable states

## [Day 18: Lavaduct Lagoon](https://adventofcode.com/2023/day/18)

Example of input:

```
R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)
```

### [Go](./go/2023/18/day18.go)

Used Shoeslace formula and Pick's theorem
Not fan of this problem

## [Day 19: Aplenty](https://adventofcode.com/2023/day/19)

Example of input:

```
px{a<2006:qkq,m>2090:A,rfg}
pv{a>1716:R,A}
lnx{m>1548:A,A}
rfg{s<537:gd,x>2440:R,A}
qs{s>3448:A,lnx}
qkq{x<1416:A,crn}
crn{x>2662:A,R}
in{s<1351:px,qqz}
qqz{s>2770:qs,m<1801:hdj,R}
gd{a>3333:R,R}
hdj{m>838:A,pv}

{x=787,m=2655,a=1222,s=2876}
{x=1679,m=44,a=2067,s=496}
{x=2036,m=264,a=79,s=2244}
{x=2461,m=1339,a=466,s=291}
{x=2127,m=1623,a=2188,s=1013}
```

### [Go](./go/2023/19/day19.go)

Used a nice tree data-structure and a propagate function

## [Day 20: Pulse Propagation](https://adventofcode.com/2023/day/20)

Example of input:

```
button -low-> broadcaster
broadcaster -low-> a
broadcaster -low-> b
broadcaster -low-> c
a -high-> b
b -high-> c
c -high-> inv
inv -low-> a
a -low-> b
b -low-> c
c -low-> inv
inv -high-> a
```

### [Go](./go/2023/20/day20.go)

Interresting problem. Part 2 was a bit difficult

## [Day 21: Step Counter](https://adventofcode.com/2023/day/21)

Example of input:

```
...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........
```

### [Go](./go/2023/21/day21.go)

Another interresting but difficult problem where a cycle has to be found to solve part 2
Quite happy with my solution found without any help

## [Day 22: Sand Slabs](https://adventofcode.com/2023/day/22)

Example of input:

```
1,0,1~1,2,1
0,0,2~2,0,2
0,2,3~2,2,3
0,0,4~0,2,4
2,0,5~2,2,5
0,1,6~2,1,6
1,1,8~1,1,9
```

### [Go](./go/2023/22/day22.go)

A nice 3D problem. Not too difficult but my approach is not very efficient.
I have to use a Z-buffer data struture to improve the implementation

## [Day 23: A Long Walk](https://adventofcode.com/2023/day/23)

Example of input:

```
#.#####################
#.......#########...###
#######.#########.#.###
###.....#.>.>.###.#.###
###v#####.#v#.###.#.###
###.>...#.#.#.....#...#
###v###.#.#.#########.#
###...#.#.#.......#...#
#####.#.#.#######.#.###
#.....#.#.#.......#...#
#.#####.#.#.#########v#
#.#...#...#...###...>.#
#.#.#v#######v###.###v#
#...#.>.#...>.>.#.###.#
#####v#.#.###v#.#.###.#
#.....#...#...#.#.#...#
#.#########.###.#.#.###
#...###...#...#...#.###
###.###.#.###v#####v###
#...#...#.#.>.>.#.>.###
#.###.###.#.###.#.#v###
#.....###...###...#...#
#####################.#
```

### [Go](./go/2023/23/day23.go)

Another interresting problem where a graph abstraction has to be used to find the longest path in a reasonable time

## [Day 24: Never Tell Me The Odds](https://adventofcode.com/2023/day/24)

Example of input:

```
19, 13, 30 @ -2,  1, -2
18, 19, 22 @ -1, -1, -2
20, 25, 34 @ -2, -2, -4
12, 31, 28 @ -1, -2, -1
20, 19, 15 @  1, -5, -3
```

### [Go](./go/2023/24/day24.go)

A very difficult problem (for me). 

## [Day 25: Snowverload](https://adventofcode.com/2023/day/25)

Example of input:

```
jqt: rhn xhk nvd
rsh: frs pzl lsr
xhk: hfx
cmg: qnr nvd lhk bvb
rhn: xhk bvb hfx
bvb: xhk hfx
pzl: lsr hfx nvd
qnr: nvd
ntq: jqt hfx bvb xhk
nvd: lhk
lsr: lhk
rzs: qnr cmg lsr rsh
frs: qnr lhk lsr
```

### [Go](./go/2023/25/day25.go)

Funny (but not so easy)
