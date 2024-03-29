package main

import (
	"strconv"
	"strings"
)

func init() {
	solvers[18] = Solve18
}

func Solve18() (solution int) {
	pyrs := `75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23`

	pyr := make(map[int][]int)
	lines := strings.Split(pyrs, "\n")
	height := len(lines)
	for y, l := range lines {
		row := make([]int, height-(height-y)+1)
		for x, tok := range strings.Split(l, " ") {
			row[x], _ = strconv.Atoi(tok)
		}
		pyr[y] = row
	}

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	for y := height - 2; y >= 0; y-- {
		for x := 0; x <= y; x++ {
			val := pyr[y][x]
			pyr[y][x] = val + max(pyr[y+1][x], pyr[y+1][x+1])
		}
	}

	return pyr[0][0]
}
