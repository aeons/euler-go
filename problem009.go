package main

func init() {
	solvers[9] = Solve9
}

func Solve9() (solution int) {
	for c := 0; c < 1000; c++ {
		for b := 0; b < c; b++ {
			for a := 0; a < b; a++ {
				if a+b+c == 1000 && a*a+b*b == c*c {
					return a * b * c
				}
			}
		}
	}
	return
}
