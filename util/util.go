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
				c <- prime
			}
		}
		close(c)
	}()
	return c
}

func Triangles() chan int {
	c := make(chan int)
	go func() {
		t, i := 0, 1
		for {
			t += i
			i++
			c <- t
		}
	}()
	return c
}

func LexicographicsPermutations(init []int) chan []int {
	c := make(chan []int)
	go func(p []int) {
		length := len(p)
		{
			p1 := make([]int, length)
			copy(p1, p)
			c <- p1
		}
		for {
			// Find the largest index k such that a[k] < a[k + 1]. If no such index exists, the permutation is the last permutation.
			k := -1
			for i := length - 2; i >= 0; i-- {
				if p[i] < p[i+1] {
					k = i
					break
				}
			}
			if k == -1 {
				close(c)
				break
			}

			// Find the largest index l such that a[k] < a[l]. Since k + 1 is such an index, l is well defined and satisfies k < l.
			l := -1
			for i := length - 1; i >= 0; i-- {
				if p[k] < p[i] {
					l = i
					break
				}
			}

			// Swap a[k] with a[l].
			p[k], p[l] = p[l], p[k]

			// Reverse the sequence from a[k + 1] up to and including the final element a[n].
			for i, j := k+1, length-1; i < j; i, j = i+1, j-1 {
				p[i], p[j] = p[j], p[i]
			}

			{
				p1 := make([]int, length)
				copy(p1, p)
				c <- p1
			}
		}
	}(init)
	return c
}
