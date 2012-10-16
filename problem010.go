package main

import (
	"github.com/aeons/euler/util"
)

func init() {
	solvers[10] = Solve10
}

func Solve10() (solution int) {
	for prime := range util.Sieve(2 * 1000 * 1000) {
		solution += prime
	}
	return
}
