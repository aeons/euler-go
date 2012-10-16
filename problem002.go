package main

func init() {
	solvers[2] = Solve2
}

func Solve2() (solution int) {
	fib := fibGenerator(4 * 1000 * 1000)
	for f := range fib {
		if f%2 == 0 {
			solution += f
		}
	}
	return
}

func fibGenerator(max int) chan int {
	x, y := 1, 1
	c := make(chan int)
	go func() {
		for x < max {
			c <- x
			x, y = y, x+y
		}
		close(c)
	}()
	return c
}
