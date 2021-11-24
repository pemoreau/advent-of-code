#!/bin/bash
# Advent of Code Rust Template

# Instructions
# For each day:
# 1. Create src/dayxx.rs. Pad 0 for 1-9 so that files sort properly.
#  Write a function like `pub fn part1(input: String) {}`
# 2. In src/lib.rs, add a line with `pub mod dayxx;` as shown
# 3. In src/lib.rs, add a case to the match, for example:
#   ```
#   1 => (day01::part1, day01::part2),
#   2 => (day02::part1, noop),
#   ```
#   Use noop whenever a part is not yet implemented.
# 4. Create inputs/dayxx.txt and add your puzzle input
# 5. Update the argument below to the current day.
#  If no arguments are given, user will be prompted for the day on stdin.
# Your code will be passed the input and timed automatically.

cargo run 1
