package main

import (
	"fmt"
)

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

	for _, v := range X {
		for key, vv := range Y {
			for _, vvv := range vv {
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
	isNullYStruct := YStruct{0, 0, 0}
	for _, v := range Y[r] {
		for _, vv := range Z[v] {
			if vv == isNullYStruct {
				continue
			}
			for _, vvv := range Y[vv] {
				if vvv != v {
					ZV := Z[vvv]
					for j := 0; j < len(ZV); j++ {
						if ZV[j] == vv {
							ZV[j] = YStruct{0, 0, 0}
							break
						}
					}
					Z[vvv] = ZV
				}
			}
		}
		delete(Z, v)
	}
}

func solve(Z map[XStruct][9]YStruct, Y map[YStruct][4]XStruct, solution []YStruct) {
	i := 0
	for i < len(solution) {
		min := 10
		c := XStruct{}
		counter := 0
		isNullYStruct := YStruct{0, 0, 0}
		for key, v := range Z {
			for _, vv := range v {
				if vv != isNullYStruct {
					counter++
				}
			}
			if counter < min {
				min = counter
				c = key
				if min == 1 {
					break
				}
			}
			counter = 0
		}
		r := YStruct{}
		for _, v := range Z[c] {
			if v != isNullYStruct {
				r = v
				break
			}
		}
		solution[i] = r
		choice(Z, Y, r)
		i++
	}
}

func sudokuSolve(grid [9][9]int) [9][9]int {
	R, C := 3, 3
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
			}
		}
	}
	Z := exactCover(X, Y)
	zeroCount := 0
	for i, row := range grid {
		for j, n := range row {
			if n == 0 {
				zeroCount++
			} else {
				yStruct := YStruct{i, j, n}
				choice(Z, Y, yStruct)
			}
		}
	}
	solution := make([]YStruct, zeroCount)
	solve(Z, Y, solution)
	for _, v := range solution {
		grid[v.R][v.C] = v.N
	}
	return grid
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
	results := sudokuSolve(grid)
	fmt.Println(results)
}
