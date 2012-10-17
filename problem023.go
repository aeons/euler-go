package main

import (
	"math"
)

func init() {
	solvers[23] = Solve23
}

func Solve23() (solution int) {
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

	max := 28124

	// Find all abundant integers below 28123
	abmap := make(map[int]struct{})
	for i := 1; i < max; i++ {
		if sd := d(i); sd > i {
			abmap[i] = struct{}{}
		}
	}
	// Put them into a slice and sort it
	abundant := make([]int, len(abmap))
	{
		i := 0
		for k, _ := range abmap {
			abundant[i] = k
			i++
		}
	}

	// All integers below max are possibly the sum of two abundant integers
	possibles := make(map[int]struct{})
	for i := 1; i < max; i++ {
		possibles[i] = struct{}{}
	}

	// For each possible sum of two abundant integers, 
	// that sum is removed from the map of possible sums
	for i := 0; i < len(abundant); i++ {
		for j := 0; j <= i; j++ {
			delete(possibles, abundant[i]+abundant[j])
		}
	}

	// Sum the relevant numbers
	for k, _ := range possibles {
		solution += k
	}

	return
}
