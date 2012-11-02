package main

func init() {
	solvers[31] = Solve31
}

func Solve31() (solution int) {
	coins := []int{1, 2, 5, 10, 20, 50, 100, 200}
	type pair struct{ r, m int }

	cache := make(map[pair]int)
	for _, i := range coins {
		cache[pair{1, i}] = 1
		cache[pair{0, i}] = 1
	}

	var ways func(remaining, max int) int
	ways = func(remaining, max int) (w int) {
		if max == 1 {
			return 1
		}
		hit, ok := cache[pair{remaining, max}]
		if ok {
			return hit
		}

		for _, c := range coins {
			if c <= remaining && c <= max {
				w += ways(remaining-c, c)
			}
		}
		cache[pair{remaining, max}] = w
		return
	}

	for r := 2; r <= 200; r++ {
		for _, c := range coins {
			if c < r {
				ways(r, c)
			}
		}
	}

	return ways(200, 200)
}
