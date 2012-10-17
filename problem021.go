package main

import (
	"math"
)

func init() {
	solvers[21] = Solve21
}

func Solve21() (solution int) {
	divisors := func(n int) []int {
		// 1 and n
		d := []int{1}
		for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
			if n%i == 0 {
				div := n / i
				if i == div {
					d = append(d, i)
				} else {
					d = append(d, i, n/i)
				}

			}
		}
		return d
	}

	d := func(n int) (sum int) {
		for _, div := range divisors(n) {
			sum += div
		}
		return
	}

	ds := make([]int, 10000)
	for i := 1; i < 10000; i++ {
		di := d(i)
		ds[i] = di
		if di < i && ds[di] == i {
			solution += i + di
		}

	}

	return
}
