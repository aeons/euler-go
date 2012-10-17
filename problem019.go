package main

import (
	"time"
)

func init() {
	solvers[19] = Solve19
}

func Solve19() (solution int) {
	d := time.Date(1901, 1, 6, 0, 0, 0, 0, time.UTC)
	end := time.Date(2000, 12, 31, 0, 0, 0, 0, time.UTC)
	for d.Before(end) {
		if d.Day() == 1 {
			solution++
		}
		d = d.AddDate(0, 0, 7)
	}
	return
}
