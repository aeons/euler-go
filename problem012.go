package main

import (
	"github.com/aeons/euler/util"
	"math"
)

func init() {
	solvers[12] = Solve12
}

func Solve12() (solution int) {
	ndiv := func(n int) int {
		// 1 and n
		d := 2
		for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
			if n%i == 0 {
				d += 2
			}
		}
		return d
	}

	triangles := util.Triangles()
	for t := range triangles {
		if d := ndiv(t); d > 500 {
			return t
		}
	}
	return
}
