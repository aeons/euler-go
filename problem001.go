package main

type Problem001 struct{}

func init() {
	solvers[1] = new(Problem001)
}

func (p *Problem001) Solve() (solution int) {
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			solution += i
		}
	}
	return
}
