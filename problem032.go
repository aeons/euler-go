package main

import (
	"github.com/aeons/euler/util"
)

func init() {
	solvers[32] = Solve32
}

func Solve32() (solution int) {
	perms := util.LexicographicsPermutations([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	prods := make(map[int]struct{})

	for p := range perms {
		oneFourFour := p[0] * (p[1]*1000 + p[2]*100 + p[3]*10 + p[4])
		if oneFourFour < 10000 {
			prod := p[5]*1000 + p[6]*100 + p[7]*10 + p[8]
			if oneFourFour == prod {
				prods[prod] = struct{}{}
			}
		}

		twoThreeFour := (p[0]*10 + p[1]) * (p[2]*100 + p[3]*10 + p[4])
		if twoThreeFour > 1000 && twoThreeFour < 10000 {
			prod := p[5]*1000 + p[6]*100 + p[7]*10 + p[8]
			if twoThreeFour == prod {
				prods[prod] = struct{}{}
			}
		}
	}

	for p, _ := range prods {
		solution += p
	}

	return
}
