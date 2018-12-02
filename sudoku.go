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

	j := -1

	for i := 0; i < len(X); i++ {
		var k int
		if nameIndex == 0 {
			k = i % N
			if k == 0 {
				j++
			}
		} else {
			k = (i % N) + 1
			if k == 1 {
				j++
			}
		}
		X[i].Values = [2]int{j, k}
		X[i].KeyName = keyNames[nameIndex]
		if i == 80 || i == 160 || i == 240 {
			nameIndex++
			j = -1
		}
	}
}

func main() {

}
