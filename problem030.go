package main

import (
	"math"
)

func init() {
	solvers[30] = Solve30
}

func Solve30() (solution int) {
	powers := [10]int{
		0,
		1,
		int(math.Pow(2, 5)),
		int(math.Pow(3, 5)),
		int(math.Pow(4, 5)),
		int(math.Pow(5, 5)),
		int(math.Pow(6, 5)),
		int(math.Pow(7, 5)),
		int(math.Pow(8, 5)),
		int(math.Pow(9, 5)),
	}

	fifthSum := func(n int) (s int) {
		for ; n > 0; n /= 10 {
			s += powers[n%10]
		}
		return
	}

	for i := 2; i < 200000; i++ {
		if fifthSum(i) == i {
			solution += i
		}
	}
	return
}
