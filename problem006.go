package main

func init() {
	solvers[6] = Solve6
}

func Solve6() (solution int) {
	n := 100
	sqsu := (n + 1) * (n / 2)
	sqsu *= sqsu
	susq := 0
	for i := 1; i <= n; i++ {
		susq += i * i
	}
	return sqsu - susq
}
