package main

import (
	"math/big"
)

func init() {
	solvers[33] = Solve33
}

func Solve33() (solution int) {
	prod := big.NewRat(1, 1)
	k := big.NewRat(1, 1)
	l := big.NewRat(1, 1)
	for i := int64(10); i < 100; i++ {
		for j := i + 1; j < 100; j++ {
			k.SetFrac64(i, j)
			fi := i / 10
			li := i % 10
			fj := j / 10
			lj := j % 10

			if fi == lj && fj != 0 {
				l.SetFrac64(li, fj)
				if k.Cmp(l) == 0 {
					prod.Mul(prod, k)
				}
			}

			if li == fj && lj != 0 {
				l.SetFrac64(fi, lj)
				if k.Cmp(l) == 0 {
					prod.Mul(prod, k)
				}
			}

			if li == lj && li != 0 && fj != 0 {
				l.SetFrac64(fi, fj)
				if k.Cmp(l) == 0 {
					prod.Mul(prod, k)
				}
			}

			if fi == fj && lj != 0 {
				l.SetFrac64(li, lj)
				if k.Cmp(l) == 0 {
					prod.Mul(prod, k)
				}
			}
		}
	}
	return int(prod.Denom().Int64())
}
