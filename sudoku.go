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

func exactCover(X [324]XStruct, Y map[YStruct][4]XStruct) map[XStruct][9]YStruct {

	Z := map[XStruct][9]YStruct{}
	i := 0
	yStruct := [9]YStruct{}

	for _, v := range X { // ok
		for key, vv := range Y { // ok
			for _, vvv := range vv { //ok
				if v == vvv {
					yStruct[i] = key
					if i == 8 {
						Z[v] = yStruct
						i = 0
					} else {
						i++
					}
				}
			}
		}
	}

	return Z
}

func choice(Z map[XStruct][9]YStruct, Y map[YStruct][4]XStruct, r YStruct) {
	for _, v := range Y[r] {
		for _, vv := range Z[v] {
			for _, vvv := range Y[vv] {
				if vvv != v {
					ZV := Z[vvv]
					for i := 0; i < len(ZV); i++ {
						if ZV[i] == vv {
							ZV[i] = YStruct{0, 0, 0}
							break
						}
					}
					Z[vvv] = ZV
				}
			}

		}
	}
}

// def select(X, Y, r):
//     for j in Y[r]:
//         for i in X[j]:
//             for k in Y[i]:
//                 if k != j:
//                     X[k].remove(i)
//         cols.append(X.pop(j))

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

	// for i, row in enumerate(grid):
	// for j, n in enumerate(row):

	Z := exactCover(X, Y)

	for i, row := range grid {
		fmt.Println(i, row)
		for j, n := range row {
			if n != 0 {
				yStruct := YStruct{i, j, n}
				// fmt.Println(yStruct)
				choice(Z, Y, yStruct)
				// fmt.Println(Z, yStruct)
			}
		}
	}
	fmt.Println(Z)

}

func main() {
	grid := [9][9]int{
		{0, 6, 1, 0, 0, 7, 0, 0, 3},
		{0, 9, 2, 0, 0, 3, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 8, 5, 3, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 5, 0, 4},
		{5, 0, 0, 0, 0, 8, 0, 0, 0},
		{0, 4, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 1, 6, 0, 8, 0, 0},
		{6, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	sudokuSolve(grid)
}
