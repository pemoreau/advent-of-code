# Advent Of Code 2020

# Comments

## [Day 01: Report Repair](https://adventofcode.com/2020/day/1)

Example of input:

```
1721
979
366
299
675
1456
```

### [Rust](./rust/2020/day01)

### [Go](./go/2020/01/day01.go)

## [Day 02: Password Philosophy](https://adventofcode.com/2020/day/2)

Example of input:

```
1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc
```

### [Rust](./rust/2020/day02)

Used `peg` for parsing the input.

## [Day 03: Toboggan Trajectory](https://adventofcode.com/2020/day/3)

Example of input:

```
..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#
```

### [Rust](./rust/2020/day03)

## [Day 04: Passport Processing](https://adventofcode.com/2020/day/4)

Example of input:

```
ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in
```

### [Rust](./rust/2020/day04)

Used `nom` for parsing the input.

## [Day 05: Binary Boarding](https://adventofcode.com/2020/day/5)

Example of input:

```
BFBBFFFLRR
FFBFBBBLLL
FBFBFBFLLL
BBFFFBFLLR
FBFFBBFLRR
```

### [Rust](./rust/2020/day05)

## [Day 06: Custom Customs](https://adventofcode.com/2020/day/6)

Example of input:

```
abc

a
b
c

ab
ac

a
a
a
a

b
```

### [Rust](./rust/2020/day06)

## [Day 07: Handy Haversacks](https://adventofcode.com/2020/day/7)

Example of input:

```
light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.
```

### [Go](./go/2020/07/day07.go)

## [Day 08: Handheld Halting](https://adventofcode.com/2020/day/8)

Example of input:

```
nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6
```

### [Go](./go/2020/08/day08.go)

## [Day 09: Encoding Error](https://adventofcode.com/2020/day/9)

Example of input:

```
35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576
```

### [Go](./go/2020/09/day09.go)

## [Day 10: Adapter Array](https://adventofcode.com/2020/day/10)

Example of input:

```
16
10
15
5
1
11
7
19
6
12
4
```

### [Go](./go/2020/10/day10.go)

## [Day 11: Seating System](https://adventofcode.com/2020/day/11)

Example of input:

```
L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL
```

### [Rust](./rust/2020/day11)

## [Day 12: Rain Risk](https://adventofcode.com/2020/day/12)

Example of input:

```
F10
N3
F7
R90
F11
```

### [Rust](./rust/2020/day12)

## [Day 13: Shuttle Search](https://adventofcode.com/2020/day/13)

Example of input:

```
time   bus 7   bus 13  bus 59  bus 31  bus 19
929      .       .       .       .       .
930      .       .       .       D       .
931      D       .       .       .       D
932      .       .       .       .       .
933      .       .       .       .       .
934      .       .       .       .       .
935      .       .       .       .       .
936      .       D       .       .       .
937      .       .       .       .       .
938      D       .       .       .       .
939      .       .       .       .       .
940      .       .       .       .       .
941      .       .       .       .       .
942      .       .       .       .       .
943      .       .       .       .       .
944      .       .       D       .       .
945      D       .       .       .       .
946      .       .       .       .       .
947      .       .       .       .       .
948      .       .       .       .       .
949      .       D       .       .       .
```

### [Rust](./rust/2020/day13)

## [Day 14: Docking Data](https://adventofcode.com/2020/day/14)

Example of input:

```
mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0
```

### [Rust](./rust/2020/day14)

## [Day 15: Rambunctious Recitation](https://adventofcode.com/2020/day/15)

Example of input:

```
5,1,9,18,13,8,0
```

### [Rust](./rust/2020/day15)

## [Day 16: Ticket Translation](https://adventofcode.com/2020/day/16)

Example of input:

```
class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12
```

### [Rust](./rust/2020/day16)

Used `peg` for parsing the input.

## [Day 17: Conway Cubes](https://adventofcode.com/2020/day/17)

Example of input:

```
.#.
..#
###
```

### [Rust](./rust/2020/day17)


## [Day 18: Operation Order](https://adventofcode.com/2020/day/18)

Example of input:

```
7 + (9 * 8 + 5 + 5 * (3 * 4 * 7 + 6 * 4)) * ((3 * 6 + 3 * 4 * 7 * 4) + 4 * 3 * 5 + 5 * (5 * 6 + 7)) * 2 + 6 * 4
9 * 4 * ((9 * 8 + 9 + 2 + 9) + 2 * 9 + 2 + 2) * 5 * 6
3 * ((9 * 3 * 8 * 6 * 6 * 7) + 8) * 2 * 9 + 4 * 8
(3 + 4 + 4 * 4 + 9) + (7 + 6 + 2 * 8) * 9 + 7 * 8
```

### [Rust](./rust/2020/day18)

Used `peg` for parsing the input.

## [Day 19: ](https://adventofcode.com/2020/day/19)

Example of input:

```
0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"

ababbb
bababa
abbbab
aaabbb
aaaabbb
```

### [Rust](./rust/2020/day19)

### [Go](./go/2020/19/day19.go)

## [Day 20: Jurassic Jigsaw](https://adventofcode.com/2020/day/20)

Example of input:

```
Tile 2311:
..##.#..#.
##..#.....
#...##..#.
####.#...#
##.##.###.
##...#.###
.#.#.#..##
..#....#..
###...#.#.
..###..###

Tile 1951:
#.##...##.
#.####...#
.....#..##
#...######
.##.#....#
.###.#####
###.##.##.
.###....#.
..#.#..#.#
#...##.#..

```

### [Go](./go/2020/20/day20.go)

Classic backtracking. Tee border of each tile has been encoded by an unsigned integer.
I don't know if it is a good idea but that simplified comparisons

## [Day 21: ](https://adventofcode.com/2020/day/21)

Example of input:

```

```

### [Rust](./rust/2020/day21)

### [Go](./go/2020/21/day21.go)

## [Day 22: ](https://adventofcode.com/2020/day/22)

Example of input:

```

```

### [Rust](./rust/2020/day22)

### [Go](./go/2020/22/day22.go)

## [Day 23: ](https://adventofcode.com/2020/day/23)

Example of input:

```

```

### [Rust](./rust/2020/day23)

### [Go](./go/2020/23/day23.go)

## [Day 24: ](https://adventofcode.com/2020/day/24)

Example of input:

```

```

### [Rust](./rust/2020/day24)

### [Go](./go/2020/24/day24.go)

## [Day 25: ](https://adventofcode.com/2020/day/25)

Example of input:

```

```

### [Rust](./rust/2020/day25)

### [Go](./go/2020/25/day25.go)

