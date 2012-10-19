package main

import (
	"math/big"
)

func init() {
	solvers[26] = Solve26
}

func Solve26() (solution int) {
	cycleLen := func(s string) int {
		length := len(s)
		// We will try finding a pattern starting at i
		for i := 0; i < length; i++ {
			// Start point for the moving target
			for j := i + 1; j < length; j++ {
				// If we have a repeat of the character at the start
				// we investigate further
				if s[i] == s[j] {
					// k is the start of the second repetition of the pattern
					k := j + (j - i)
					// Walk up both i and j through two complete repetitions
					cl := 0
					for ; s[i] == s[j] && i < k; i, j = i+1, j+1 {
						cl++
					}
					// If i == k then we have the pattern two times in a row
					// and we can return half the length as the answer
					if i == k {
						return cl / 2
					}
				}
			}
		}
		return -1
	}

	for i, max := 0, 0; i < 1000; i++ {
		rs := big.NewRat(1, int64(i+1)).FloatString(10000)[2:10001]
		cl := cycleLen(rs)
		if cl > max {
			max = cl
			solution = i + 1
		}
	}
	return
}
