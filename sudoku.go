package main

import (
	"fmt"
)

type XStruct struct {
	KeyName string
	Values  [2]int
}

type XArray []XStruct

func sudokuSolve(grid [9][9]int) {
	R, C := 3, 3
	fmt.Println(R, C)
	N := 9

	X := [324]XStruct{} // [N * N * 4]

	keyNames := [4]string{"rc", "rn", "cn", "bn"}

	nameIndex := 0
	k := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			X[k].Values = [2]int{i, j}
			X[k].KeyName = "rc"
			k++
		}
	}

	for h := 0; h < 3; h++ {
		nameIndex++
		for i := 0; i < N; i++ {
			for j := 1; j < N+1; j++ {
				X[k].Values = [2]int{i, j}
				X[k].KeyName = keyNames[nameIndex]
				k++
			}
		}
	}

	for h := 0; h < N; h++ {
		for i := 0; i < N; i++ {
			for j := 1; j < N+1; j++ {
				b := (h/R)*R + (i / C)
			}
		}
	}

}

func main() {

}
