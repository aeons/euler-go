package main

import (
	"strings"
)

func init() {
	solvers[17] = Solve17
}

func Solve17() (solution int) {
	letters := make(map[int]string)
	letters[1] = "one"
	letters[2] = "two"
	letters[3] = "three"
	letters[4] = "four"
	letters[5] = "five"
	letters[6] = "six"
	letters[7] = "seven"
	letters[8] = "eight"
	letters[9] = "nine"
	letters[10] = "ten"
	letters[11] = "eleven"
	letters[12] = "twelve"
	letters[13] = "thirteen"
	letters[14] = "fourteen"
	letters[15] = "fifteen"
	letters[16] = "sixteen"
	letters[17] = "seventeen"
	letters[18] = "eighteen"
	letters[19] = "nineteen"
	letters[20] = "twenty"
	letters[30] = "thirty"
	letters[40] = "forty"
	letters[50] = "fifty"
	letters[60] = "sixty"
	letters[70] = "seventy"
	letters[80] = "eighty"
	letters[90] = "ninety"
	letters[100] = "hundred"
	letters[1000] = "thousand"

	var asWords func(n int) string
	asWords = func(n int) string {
		if n < 21 {
			return letters[n]
		}
		if n < 100 {
			w := letters[n/10*10]
			if n%10 != 0 {
				w += " " + asWords(n%10)
			}
			return w
		}
		if n < 1000 {
			w := letters[n/100] + " " + letters[100]
			if n%100 != 0 {
				w += " and " + asWords(n%100)
			}
			return w
		}
		return letters[n/1000] + " " + letters[1000]
	}

	length := 0
	for i := 0; i <= 1000; i++ {
		length += len(strings.Replace(asWords(i), " ", "", -1))
	}

	return length
}
