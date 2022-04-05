package main

import (
	"github.com/pemoreau/advent-of-code/go/utils"
)

type EnvInterval [4]utils.Interval

func createEnvInterval(env Env) EnvInterval {
	var res EnvInterval
	for i, v := range env {
		res[i] = utils.Interval{v, v}
	}
	return res
}

func eqlInterval(a, b utils.Interval) utils.Interval {
	if a.Min == a.Max && a.Min == b.Min && a.Max == b.Max {
		return utils.Interval{1, 1}
	}
	if a.Max < b.Min || b.Max < a.Min {
		return utils.Interval{0, 0}
	}
	return utils.Interval{0, 1}
}

func abstractInterpretation(e Expr, env EnvInterval) utils.Interval {
	switch exp := e.(type) {
	case Value:
		return utils.Interval{int(exp), int(exp)}
	case Reg:
		return env[regIndex(exp)]
	case Add:
		return env[regIndex(exp.reg)].Add(abstractInterpretation(exp.arg, env))
	case Mul:
		return env[regIndex(exp.reg)].Mul(abstractInterpretation(exp.arg, env))
	case Div:
		return env[regIndex(exp.reg)].Div(abstractInterpretation(exp.arg, env))
	case Mod:
		return env[regIndex(exp.reg)].Mod2(abstractInterpretation(exp.arg, env))
	case Eql:
		return eqlInterval(env[regIndex(exp.reg)], abstractInterpretation(exp.arg, env))
	default:
		panic("unknown exp")
	}
}

func abstractInterpretationInstr(i Instr, env EnvInterval) EnvInterval {
	newEnv := env
	switch ins := i.(type) {
	case Assign:
		newEnv[regIndex(ins.reg)] = abstractInterpretation(ins.rhs, env)
	case Input:
		newEnv[regIndex(ins.reg)] = utils.Interval{1, 9}
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
	z := regIndex("z")

	// fmt.Println(key{len(program), env})
	for _, i := range program {
		env = abstractInterpretationInstr(i, env)
	}
	// fmt.Println(key{len(program), env})
	return env[z].Min <= 0 && 0 <= env[z].Max
}
