package main

import (
	"github.com/aeons/euler/util"
)

func init() {
	solvers[24] = Solve24
}

func Solve24() (solution int) {
	p := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// This could be done faster doing in-place permutations on a single slice
	i := 0
	for p := range util.LexicographicsPermutations(p) {
		if i == 999999 {
			for _, n := range p {
				solution *= 10
				solution += n
			}
			return solution
		}
		i++
	}
	return
}
