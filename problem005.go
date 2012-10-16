package main

var (
	divisors = []int{7, 8, 9, 11, 12, 13, 14, 15, 16, 17, 18, 19}
)

func init() {
	solvers[5] = Solve5
}

func Solve5() (solution int) {
	for i := 2520; ; i = i + 3*7*20 {
		if isValid(i) {
			return i
		}
	}
	return
}

func isValid(n int) bool {
	for _, d := range divisors {
		if n%d != 0 {
			return false
		}
	}
	return true
}
