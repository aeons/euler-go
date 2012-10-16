package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Solver func() int

var (
	solvers  = make(map[int]Solver)
	solveAll bool
)

func main() {
	flag.BoolVar(&solveAll, "all", false, "solve all problems with solutions")
	flag.Usage = usage
	flag.Parse()

	if solveAll {
		fmt.Printf("Solving %d problems that currently have solutions.\n", len(solvers))
		for number, solver := range solvers {
			solve(number, solver)
		}
		os.Exit(0)
	}

	if flag.NArg() != 1 {
		usage()
		os.Exit(0)
	}

	var problemNumber int
	if p, err := strconv.ParseInt(flag.Arg(0), 10, 32); err != nil {
		fmt.Println("problem-number must be an integer")
		usage()
		os.Exit(0)
	} else {
		problemNumber = int(p)
	}

	if problemNumber < 1 {
		fmt.Println("problem-number must be positive")
		usage()
		os.Exit(0)
	}

	solver, ok := solvers[problemNumber]
	if !ok {
		fmt.Printf("solution for problem %d not found\n", problemNumber)
		os.Exit(0)
	}

	solve(problemNumber, solver)

}

func usage() {
	fmt.Fprint(os.Stderr, "usage: euler [-all] problem-number\n")
	flag.PrintDefaults()
	os.Exit(1)
}

func solve(n int, s Solver) {
	t0 := time.Now()
	solution := s()
	ttotal := time.Since(t0)

	fmt.Printf("Solution for problem %d: %d found in %v\n", n, solution, ttotal)
}
