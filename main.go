package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Problem struct{}

type Solver interface {
	Solve() int
}

var (
	solvers map[int]Solver = make(map[int]Solver)
)

const (
	usage string = "usage: euler.exe problem-number"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println(usage)
		os.Exit(0)
	}

	var problemNumber int
	if p, err := strconv.ParseInt(os.Args[1], 10, 32); err != nil {
		fmt.Println(usage)
		fmt.Println("problem-number must be an integer")
		os.Exit(0)
	} else {
		problemNumber = int(p)
	}

	if problemNumber < 1 {
		fmt.Println(usage)
		fmt.Println("problem-number must be positive")
		os.Exit(0)
	}

	solver, ok := solvers[problemNumber]
	if !ok {
		fmt.Printf("solution for problem %d not found\n", problemNumber)
		os.Exit(0)
	}

	t0 := time.Now()
	solution := solver.Solve()
	ttotal := time.Since(t0)

	fmt.Printf("Solution for problem %d: %d found in %v\n", problemNumber, solution, ttotal)
}
