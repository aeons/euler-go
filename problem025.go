package main

import (
	"math/big"
)

func init() {
	solvers[25] = Solve25
}

func Solve25() (solution int) {
	x := big.NewInt(1)
	y := big.NewInt(1)
	z := big.NewInt(0)

	for solution = 3; len(x.String()) < 1000; solution++ {
		z.Set(x)
		x.Add(x, y)
		y.Set(z)
	}
	solution--

	return
}
