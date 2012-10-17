package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func init() {
	solvers[63] = Solve63
}

func Solve63() (solution int) {
	// Load data
	file, _ := os.Open("data/63.txt")
	defer file.Close()

	reader := bufio.NewReader(file)
	lines := make([]string, 100)
	for i := 0; i < 100; i++ {
		l, _ := reader.ReadString('\n')
		// remove \r\n
		lines[i] = l[:len(l)-2]
	}

	pyr := make(map[int][]int)
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
