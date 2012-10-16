package main

import (
	"github.com/aeons/euler/util"
	"math"
)

func init() {
	solvers[3] = Solve3
}

func Solve3() (solution int) {
	n := 600851475143
	max := int(math.Sqrt(float64(n)))
	primes := util.Sieve(max)
	for p := range primes {
		if n%p == 0 {
			n /= p
		}
		if n == 1 {
			return p
		}
	}
	return
}
