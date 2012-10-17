package main

func init() {
	solvers[15] = Solve15
}

func Solve15() (solution int) {
	const edge = 20

	type pair struct{ x, y int }

	cache := make(map[pair]int)
	cache[pair{edge, edge}] = 1

	var r func(x, y int) int
	r = func(x, y int) int {
		l, ok := cache[pair{x, y}]
		if !ok {
			if x == edge || y == edge {
				l = 1
			} else {
				l = r(x+1, y) + r(x, y+1)
			}
			cache[pair{x, y}] = l
		}
		return l
	}
	return r(0, 0)
}
