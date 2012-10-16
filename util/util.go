package util

func Fibonacci(max int) chan int {
	c := make(chan int)
	x, y := 1, 1

	go func() {
		for x < max {
			c <- x
			x, y = y, x+y
		}
		close(c)
	}()
	return c
}

func Sieve(max int) chan int {
	sieve := make([]int, max, max)

	for i := range sieve {
		sieve[i] = i
	}
	sieve[1] = 0

	for p := 2; p < max; p++ {
		// Find starting point
		for p < max && sieve[p] == 0 {
			p++
		}
		if p == max {
			break
		}

		// Mark multiples
		for i := p * p; i < max; i = i + p {
			sieve[i] = 0
		}
	}

	c := make(chan int)
	go func() {
		for i := 2; i < max; i++ {
			if prime := sieve[i]; prime != 0 {
				c <- sieve[i]
			}
		}
		close(c)
	}()
	return c
}
