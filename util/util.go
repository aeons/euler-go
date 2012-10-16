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

	p := 2
	for p < max {
		// Find starting point
		for sieve[p] == 0 {
			p++
		}

		// Mark multiples
		for i := 2 * p; i < max; i = i + p {
			sieve[i] = 0
		}

		// Ready to find next starting point
		p *= p
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
