package main

import (
	"math/big"
)

func init() {
	solvers[29] = Solve29
}

func Solve29() (solution int) {
	terms := make(map[string]struct{})
	lim := 100
	abig := big.NewInt(0)
	bbig := big.NewInt(0)
	pow := big.NewInt(0)
	for a := 2; a <= lim; a++ {
		abig.SetInt64(int64(a))
		for b := 2; b <= lim; b++ {
			bbig.SetInt64(int64(b))
			terms[pow.Exp(abig, bbig, nil).String()] = struct{}{}
		}
	}
	return len(terms)
}
