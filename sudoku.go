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

	j := 0

	for i := 0; i < 81; i++ {
		k := i % N
		if i != 0 && k == 0 {
			j++
		}
		X[i].KeyName = keyNames[nameIndex]
		X[i].Values = [2]int{j, k}
	}

	nameIndex++

	j = 0
	for i := 81; i < len(X); i++ {
		k := (i % N) + 1

		if i != 0 && k == 1 {
			j++
		}

		X[i].KeyName = keyNames[nameIndex]
		X[i].Values = [2]int{j - 1, k}

		if i == 160 {
			nameIndex++
			j = 0
		} else if i == 240 {
			nameIndex++
			j = 0
		}
	}

}

func main() {

}
