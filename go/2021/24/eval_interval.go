package main

import (
	"github.com/pemoreau/advent-of-code/go/utils/interval"
)

type EnvInterval [4]interval.Interval

func createEnvInterval(env Env) EnvInterval {
	var res EnvInterval
	for i, v := range env {
		res[i] = interval.Interval{v, v}
	}
	return res
}

func eqlInterval(a, b interval.Interval) interval.Interval {
	if a.Min == a.Max && a.Min == b.Min && a.Max == b.Max {
		return interval.Interval{1, 1}
	}
	if a.Max < b.Min || b.Max < a.Min {
		return interval.Interval{0, 0}
	}
	return interval.Interval{0, 1}
}

func abstractInterpretation(e Expr, env EnvInterval) interval.Interval {
	switch exp := e.(type) {
	case Value:
		return interval.Interval{int(exp), int(exp)}
	case Reg:
		return env[exp]
	case Add:
		return env[exp.reg].Add(abstractInterpretation(exp.arg, env))
	case Mul:
		return env[exp.reg].Mul(abstractInterpretation(exp.arg, env))
	case Div:
		return env[exp.reg].Div(abstractInterpretation(exp.arg, env))
	case Mod:
		return env[exp.reg].Mod2(abstractInterpretation(exp.arg, env))
	case Eql:
		return eqlInterval(env[exp.reg], abstractInterpretation(exp.arg, env))
	default:
		panic("unknown exp")
	}
}

func abstractInterpretationInstr(i Instr, env EnvInterval) EnvInterval {
	newEnv := env
	switch ins := i.(type) {
	case Assign:
		newEnv[ins.reg] = abstractInterpretation(ins.rhs, env)
	case Input:
		wIndex := 0
		newEnv[wIndex] = interval.Interval{1, 9}
	default:
		panic("unknown instr")
	}
	return newEnv
}

type key struct {
	len int
	env EnvInterval
}

func reachable(program []Instr, env EnvInterval) bool {
	z := regIndex('z')

	// fmt.Println(key{len(program), env})
	for _, i := range program {
		env = abstractInterpretationInstr(i, env)
	}
	// fmt.Println(key{len(program), env})
	return env[z].Min <= 0 && 0 <= env[z].Max
}
