package main

func init() {
	solvers[28] = Solve28
}

func Solve28() (solution int) {
	makeSteps := func() chan int {
		c := make(chan int)
		go func() {
			i, j := 1, 1
			for {
				c <- i
				if i == j {
					j++
				} else {
					i++
				}
			}
		}()
		return c
	}

	makeSpiral := func(n int) []int {
		if n%2 == 0 {
			panic("can only make odd-sized spiral")
		}
		spiral := make([]int, n*n)
		x, y, val := n/2, n/2, 1
		spiral[y*n+x] = val
		stepsGen := makeSteps()
		for {
			// Go right
			steps := <-stepsGen
			for i := 0; i < steps; i++ {
				x++
				val++
				spiral[y*n+x] = val
				if val == n*n {
					return spiral
				}
			}

			// Go down
			steps = <-stepsGen
			for i := 0; i < steps; i++ {
				y++
				val++
				spiral[y*n+x] = val
				if val == n*n {
					return spiral
				}
			}

			// Go left
			steps = <-stepsGen
			for i := 0; i < steps; i++ {
				x--
				val++
				spiral[y*n+x] = val
				if val == n*n {
					return spiral
				}
			}

			// Go up
			steps = <-stepsGen
			for i := 0; i < steps; i++ {
				y--
				val++
				spiral[y*n+x] = val
				if val == n*n {
					return spiral
				}
			}
		}
		return spiral
	}

	size := 1001
	spiral := makeSpiral(size)
	for i := 0; i < size; i++ {
		solution += spiral[i*size+i]
		solution += spiral[(size-i-1)*size+i]
	}
	// Subtract the double added center cell
	return solution - 1
}
