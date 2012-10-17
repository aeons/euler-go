package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func init() {
	solvers[16] = Solve16
}

func Solve16() (solution int) {
	two := big.NewInt(2)
	mille := big.NewInt(1000)
	pow := big.NewInt(0)
	pow.Exp(two, mille, nil)
	powstr := fmt.Sprint(pow)
	toks := strings.Split(powstr, "")
	for _, t := range toks {
		n, _ := strconv.Atoi(t)
		solution += n
	}
	return
}
