// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 198.

// Package eval provides an expression evaluator.
package eval

import (
	"fmt"
	"math"
)

//!+env

type Env map[Var]float64

//!-env

//!+Eval1

// Eval returns the value of this Expr in the environment env.
func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

//!-Eval1

//!+Eval2

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

func (t ternary) Eval(env Env) float64 {
	switch t.op1 {
	case '?':
		switch t.op2 {
		case ':':
			x := t.x.Eval(env)
			if math.IsInf(x, 0) || math.IsNaN(x) || x == 0 {
				return t.z.Eval(env)
			}
			return t.y.Eval(env)
		}
		panic(fmt.Sprintf("unsupported ternary operator: %q", t.op2))
	}
	panic(fmt.Sprintf("unsupported ternary operator: %q", t.op1))
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}

//!-Eval2
