package main

import (
	"fmt"
	"sudoku/src/modules"
	"time"
)

func main() {
	startTime := time.Now()
	grid := [9][9]int{
		{4, 6, 1, 0, 0, 7, 0, 0, 3},
		{0, 9, 2, 0, 0, 3, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 8, 5, 3, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 5, 0, 4},
		{5, 0, 0, 0, 0, 8, 0, 0, 0},
		{0, 4, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 1, 6, 0, 8, 0, 0},
		{6, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	answer := modules.SudokuSolve(grid)

	fmt.Println(answer)
	finishTime := time.Now()
	fmt.Printf("time: %f[Sec]", (finishTime.Sub(startTime)).Seconds())
}
