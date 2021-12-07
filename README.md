# Advent Of Code 2021

Learning Rust

# Command

`cargo run 1` to run the first puzzle

`cargo test` to run unit tests

# Timings

Execution time on an old Mac Pro (Late 2013), 3,7 GHz Quad-Core Intel Xeon E5

|                         | part A     | part B     |
| :---------------------- | :--------- | :--------- |
| [day 1](./src/day01.rs) | ` 0.089ms` | ` 0.067ms` |
| [day 2](./src/day02.rs) | ` 0.092ms` | ` 0.063ms` |
| [day 3](./src/day03.rs) | ` 0.157ms` | ` 0.084ms` |
| [day 4](./src/day04.rs) | ` 1.048ms` | ` 0.841ms` |
| [day 5](./src/day05.rs) | ` 45.94ms` | ` 46.03ms` |
| [day 6](./src/day06.rs) | ` 0.010ms` | ` 0.008ms` |
| [day 7](./src/day07.rs) | ` 0.274ms` | ` 0.795ms` |

# Comments

## Day 01

nothing special except the use of `windows` function

## Day 02

use `split_one` instead of regex to speed-up parsing

use `for-loop` style and then `fold` for the second part

## Day 03

nothing special

## Day 04

use `split("\n\n")` to separate parts (instead of counting the number of entries)

use `array2d::Array2D` to represent the board but this may be not the best choice

the program contains too many `for-loop` in my opinion

## Day 05

a bit slow due the the use of regex for parsing

discovered the `signum` function

use references to mutable structures

## Day 06

very simple solution thanks to `rotate_left` function

## Day 07

use a cost function as parameter
