package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func init() {
	solvers[20] = Solve20
}

func Solve20() (solution int) {
	fact := big.NewInt(0)
	fact.MulRange(1, 100)
	for _, tok := range strings.Split(fmt.Sprint(fact), "") {
		n, _ := strconv.Atoi(tok)
		solution += n
	}
	return
}
