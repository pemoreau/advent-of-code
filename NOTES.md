Some interesting repositories

### Go

https://github.com/bozdoz/advent-of-code-2021
https://github.com/Skarlso/aoc2021
https://github.com/danvk/aoc2021
https://github.com/Maitgon/Advent2021/tree/master/Go
https://github.com/alexchao26/advent-of-code-go/tree/main/2021
https://github.com/sebnyberg/aoc/tree/main/aoc2021
https://github.com/jdrst/adventofgo/tree/main/2021
https://github.com/sekullbe/advent-of-code
https://github.com/derat/advent-of-code/tree/main/2021

https://github.com/janetschel/advent-of-go-2020

### Rust

https://github.com/jontmy/aoc-rust/tree/master/src
https://github.com/AxlLind/adventofcode2021

### Python (smart)

https://github.com/r0f1/adventofcode2021

pour info pour profiler j'ai fait ça :

- j'ai ajouté un test qui appelle Part1 au lieu du main
- j'ai fait go test -cpuprofile cpu.prof -bench .
- puis go tool pprof -http=:8800 cpu.prof
