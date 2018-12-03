package main

import (
	"fmt"
)

// Y[(r, c, n)] = [
// 	("rc", (r, c)),
// 	("rn", (r, n)),
// 	("cn", (c, n)),
// 	("bn", (b, n))]

type XStruct struct {
	KeyName string
	Values  [2]int
}

type XArray []XStruct

type YStruct struct {
	R int
	C int
	N int
}

func exactCover(X [324]XStruct, Y map[YStruct][4]XStruct) {

	Z := map[XStruct][9]YStruct{}
	i := 0
	ystruct := [9]YStruct{}

	for _, v := range X { // ok
		for key, vv := range Y { // ok
			for _, vvv := range vv { //ok
				if v == vvv {
					ystruct[i] = key
					if i == 8 {
						Z[v] = ystruct
						i = 0
					} else {
						i++
					}
				}
			}
		}
	}

	fmt.Println(Z)
}

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

	Y := map[YStruct][4]XStruct{}

	for h := 0; h < N; h++ {
		for i := 0; i < N; i++ {
			for j := 1; j < N+1; j++ {
				b := (h/R)*R + (i / C)
				yStruct := YStruct{h, i, j}
				rc := XStruct{}
				rc.KeyName = "rc"
				rc.Values = [2]int{h, i}

				rn := XStruct{}
				rn.KeyName = "rn"
				rn.Values = [2]int{h, j}

				cn := XStruct{}
				cn.KeyName = "cn"
				cn.Values = [2]int{i, j}

				bn := XStruct{}
				bn.KeyName = "bn"
				bn.Values = [2]int{b, j}

				Y[yStruct] = [4]XStruct{rc, rn, cn, bn}
				// yStruct.R = h
				// yStruct.C = i
				// yStruct.N = j
			}
		}
	}

}

func main() {

}
