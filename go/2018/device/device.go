package device

import (
	"fmt"
	"strings"
)

type instruction struct {
	opcode string
	args   [3]int
}

type Machine struct {
	binding   int
	ip        int
	registers [6]int
	program   []instruction
}

func (m *Machine) Register(n int) int {
	if n < 0 || n >= len(m.registers) {
		return 0
	}
	return m.registers[n]
}

func (m *Machine) SetRegister(n int, v int) {
	if n < 0 || n >= len(m.registers) {
		return
	}
	m.registers[n] = v
}

func (m *Machine) Ip() int {
	return m.ip
}

func bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}

func (m *Machine) Run(debug bool) bool {
	if m.ip >= len(m.program) || m.ip < 0 {
		return false
	}

	m.registers[m.binding] = m.ip
	var inst = m.program[m.ip]
	var v0 = inst.args[0]
	var v1 = inst.args[1]
	var v2 = inst.args[2]

	if debug {
		//fmt.Printf("ip=%d %v %s %d %d %d ", m.ip, m.registers, inst.opcode, inst.args[0], inst.args[1], inst.args[2])
	}

	switch inst.opcode {
	case "addr":
		m.registers[v2] = m.registers[v0] + m.registers[v1]
	case "addi":
		m.registers[v2] = m.registers[v0] + v1
	case "mulr":
		m.registers[v2] = m.registers[v0] * m.registers[v1]
	case "muli":
		m.registers[v2] = m.registers[v0] * v1
	case "banr":
		m.registers[v2] = m.registers[v0] & m.registers[v1]
	case "bani":
		m.registers[v2] = m.registers[v0] & v1
	case "borr":
		m.registers[v2] = m.registers[v0] | m.registers[v1]
	case "bori":
		m.registers[v2] = m.registers[v0] | v1
	case "setr":
		m.registers[v2] = m.registers[v0]
	case "seti":
		m.registers[v2] = v0
	case "gtir":
		m.registers[v2] = bool2int(v0 > m.registers[v1])
	case "gtri":
		m.registers[v2] = bool2int(m.registers[v0] > v1)
	case "gtrr":
		m.registers[v2] = bool2int(m.registers[v0] > m.registers[v1])
	case "eqir":
		m.registers[v2] = bool2int(v0 == m.registers[v1])
	case "eqri":
		m.registers[v2] = bool2int(m.registers[v0] == v1)
	case "eqrr":
		m.registers[v2] = bool2int(m.registers[v0] == m.registers[v1])
	}

	if debug {
		//fmt.Printf("%v\n", m.registers)
	}

	m.ip = m.registers[m.binding]
	m.ip++
	return true
}

func (m *Machine) init(binding int) {
	m.binding = binding
	m.ip = 0
}

func (m *Machine) loadProgram(lines []string) {
	for _, line := range lines {
		var i instruction
		fmt.Sscanf(line, "%s %d %d %d", &i.opcode, &i.args[0], &i.args[1], &i.args[2])
		m.program = append(m.program, i)
	}
}

func CreateMachine(input string) *Machine {
	input = strings.Trim(input, "\n")
	var lines = strings.Split(input, "\n")

	var binding int
	fmt.Sscanf(lines[0], "#ip %d", &binding)
	lines = lines[1:]

	var m = &Machine{}
	m.init(binding)
	m.loadProgram(lines)
	return m
}
