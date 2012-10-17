package main

func init() {
	solvers[14] = Solve14
}

func Solve14() (solution int) {
	cache := make(map[int]int)

	var seq func(n int) int
	seq = func(n int) int {
		if n == 1 {
			return 1
		}
		l, ok := cache[n]
		if !ok {
			if n%2 == 0 {
				l = 1 + seq(n/2)
			} else {
				l = 1 + seq(3*n+1)
			}
			cache[n] = l
		}
		return l
	}

	max := 0
	for i := 1; i < 1000000; i++ {
		if s := seq(i); s > max {
			max = s
			solution = i
		}
	}
	return
}
