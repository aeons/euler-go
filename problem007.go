package main

import (
	"github.com/aeons/euler/util"
)

func init() {
	solvers[7] = Solve7
}

func Solve7() (solution int) {
	primes := util.Sieve(105000)
	i := 1
	for p := range primes {
		if i == 10001 {
			return p
		}
		i++
	}
	return
}
