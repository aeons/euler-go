package main

func init() {
	solvers[1] = Solve1
}

func Solve1() (solution int) {
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			solution += i
		}
	}
	return
}
