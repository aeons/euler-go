package main

import (
	"bufio"
	"os"
	"sort"
	"strings"
)

func init() {
	solvers[22] = Solve22
}

func Solve22() (solution int) {
	score := func(s string) (sc int) {
		for _, c := range s {
			sc += int(c) - 64
		}
		return
	}

	file, _ := os.Open("data/22.txt")
	defer file.Close()

	reader := bufio.NewReader(file)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\"", "", -1)
	names := strings.Split(text, ",")
	sort.Strings(names)

	for rank, name := range names {
		solution += (rank + 1) * score(name)
	}
	return
}
