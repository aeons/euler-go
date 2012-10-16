package main

import (
	"github.com/aeons/euler/util"
)

func init() {
	solvers[2] = Solve2
}

func Solve2() (solution int) {
	fib := util.Fibonacci(4 * 1000 * 1000)
	for f := range fib {
		if f%2 == 0 {
			solution += f
		}
	}
	return
}
