package main

import (
	"github.com/aeons/euler/util"
)

func init() {
	solvers[27] = Solve27
}

func Solve27() (solution int) {
	primes := make(map[int]struct{})
	for p := range util.Sieve(10709) {
		primes[p] = struct{}{}
	}

	max := 0
	for a := -999; a < 1000; a++ {
		for b := -999; b < 1000; b++ {
			for n := 0; ; n++ {
				form := n*n + a*n + b
				if form < 0 || form%2 == 0 {
					break
				}
				_, isPrime := primes[form]
				if !isPrime {
					if n > max {
						max = n
						solution = a * b
					}
					break
				}
			}
		}
	}
	return
}
