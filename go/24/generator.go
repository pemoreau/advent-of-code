package main

import (
	"fmt"
	"strings"
)

func compile(input string, index int) (r string, i int) {
	cmd := strings.Split(input, " ")
	i = index
	switch cmd[0] {
	case "inp":
		r = fmt.Sprintf("%v = inp[%d]", cmd[1], index)
		i = index + 1
	case "add":
		r = fmt.Sprintf("%v = %v + %v", cmd[1], cmd[1], cmd[2])
	case "mul":
		if cmd[2] == "0" {
			r = fmt.Sprintf("%v = %v", cmd[1], cmd[2])
		} else {
			r = fmt.Sprintf("%v = %v * %v", cmd[1], cmd[1], cmd[2])
		}
	case "div":
		if cmd[2] != "1" {
			r = fmt.Sprintf("%v = %v / %v", cmd[1], cmd[1], cmd[2])
		}
	case "mod":
		r = fmt.Sprintf("%v = %v %% %v", cmd[1], cmd[1], cmd[2])
	case "eql":
		r = fmt.Sprintf("if %v == %v {\n %v=1\n} else {\n %v=0\n}", cmd[1], cmd[2], cmd[1], cmd[1])
	}
	return
}

func genCode(lines []string) {
	fmt.Println(`package main

	import (
		"fmt"
	)`)
	fmt.Println(`func Run(inp []int) (w, x, y, z int) {
		w = 0
		x = 0
		y = 0
		z = 0
		`)
	i := 0
	for _, line := range lines {
		var s string
		s, i = compile(line, i)
		fmt.Println(s)
	}
	fmt.Println("return \n}")

}
