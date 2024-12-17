package main

import (
	"fmt"
)

const (
	A = 0
	B = 1
	C = 2
)

type instruction = int

type Machine struct {
	ip        int
	registers [3]int
	program   []instruction
}

type Output struct {
	out []int
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

func (m *Machine) Run(output *Output, debug bool) bool {
	if m.ip >= len(m.program) || m.ip < 0 {
		return false
	}

	var inst = m.program[m.ip]
	var operand = m.program[m.ip+1]

	var combo = func(v int) int {
		if v == 7 {
			panic("not implemented")
		}
		if v >= 0 && v <= 3 {
			return v
		}
		return m.Register(v - 4)
	}

	var dv = func(reg int) {
		var num = m.Register(A)
		var den = 1 << combo(operand)
		var res = num / den
		//fmt.Printf("%d-dv operand: %d num: %d deb: %d res: %d\n", reg, operand, num, den, res)
		m.SetRegister(reg, res)
	}

	if debug {
		fmt.Printf("ip=%d A=%d B=%d C=%d inst: %d operand: %d combo: %d ", m.ip, m.Register(A), m.Register(B), m.Register(C), inst, operand, combo(operand))
	}
	switch inst {
	case 0: // adv
		dv(A)
	case 1: // bxl
		m.SetRegister(B, m.Register(B)^operand)
	case 2: // bst
		var res = combo(operand) % 8
		m.SetRegister(B, res)
	case 3: // jnz
		if m.Register(A) != 0 {
			m.ip = operand
			return true
		}
	case 4: // bxc
		m.SetRegister(B, m.Register(B)^m.Register(C))
	case 5: // out
		output.out = append(output.out, combo(operand)%8)
		if debug {
			fmt.Println()
			fmt.Printf("out: %d\n", combo(operand)%8)
		}
	case 6: //bdv
		dv(B)
	case 7: // cdv
		dv(C)
	}

	if debug {
		fmt.Printf("A=%d B=%d C=%d\n", m.Register(A), m.Register(B), m.Register(C))
	}

	m.ip += 2
	return true
}

func (m *Machine) init() {
	m.ip = 0
}

func CreateMachine(inst []instruction, a, b, c int) *Machine {
	var m = &Machine{}
	m.registers[A] = a
	m.registers[B] = b
	m.registers[C] = c
	m.program = inst
	return m
}
