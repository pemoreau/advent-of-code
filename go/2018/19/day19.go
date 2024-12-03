package main

import (
	_ "embed"
	"fmt"
	. "github.com/pemoreau/advent-of-code/go/2018/device"
	"time"
)

//go:embed input.txt
var inputDay string

func Part1(input string) int {
	var m = CreateMachine(input)

	var s = map[int]int{}
	for m.Run(false) {
		//fmt.Println(m.registers[0])
		s[m.Register(0)]++
	}

	//for k, v := range s {
	//	fmt.Printf("%d: %d\n", k, v)
	//}

	return m.Register(0)
}

func Part2(input string) int {
	var m = CreateMachine(input)
	m.SetRegister(0, 1)
	var s = map[int]int{}
	for m.Run(false) && m.Register(0) < 7 {
		s[m.Register(0)]++
	}

	//for k, v := range s {
	//	fmt.Printf("%d: %d\n", k, v)
	//}

	return m.Register(0)
}

func main() {
	fmt.Println("--2018 day 19 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", Part2(inputDay))
	fmt.Println(time.Since(start))
}

/*
--2018 day 19 solution--
0: 7720
1: 3860
3: 13504
7: 1826796
248: 1859540
730: 3719104
1694: 7712
part1:  1694

964 = 2x2x241

diviseurs = 1, 2, 4, 241, 482, 964
somme =

1+(1x2) = 3
3+(2x2) = 7
7+(1x241) = 248
248 + (2x241) = 730
730 + (2x2x241) = 1694

1+(1x2)+(2x2)+(1x241)+(2x241)+(2x2x241)
1 + (1+2)x2 + (1+2+4)x241

10551364 = 2x2x37x71293

diviseurs = 1,2,4,37, 74, 148, 71293,
            142586,
1
2
4
37
74
148
71293
142586
285172
2637841
5275682
10551364

somme = 18964204
1+(1x2)+(2x2)+(1x37)+(2x37)+(2x2x37)
 +(1x71293)+(2x71293)+(2x2x71293)+(2x2x37x71293)

1+(1x2) = 3
3+(2x2) = 7
7+(1x37) = 44
44 + (2x37) = 118
118 + (2x2x37) = 266
266 + (1x71293)





...
0 964 . 1 . 964
1   1 . 2 . 964
...
1 482 . 2 . 964
3 482 . 2 . 964
...
3 965 . 2 . 964
3 965 . 3 . 964
...
3   1 . 3 . 964
...
3 241 . 4 . 964
7 241 . 4 . 964
...
7 965 . 4 . 964
7 965 . 5 . 964
...
7   1 . 5 . 964
...
7   4 . 241 . 964
248 4 . 241 . 964
...
248 2 . 482 . 964
730 2 . 482 . 964







0: 84410911
1: 42205477
3: 147719104
part2:  7

248 = 2x2x2x31
730 = 2x5x73
1694 = 2x7x11x11

*/
