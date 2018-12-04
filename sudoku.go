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

// def select(X, Y, r):
//     cols = []
//     for j in Y[r]:
//         for i in X[j]:
//             for k in Y[i]:
//                 if k != j:
//                     X[k].remove(i)
//         cols.append(X.pop(j))
//     print(cols)
//     return cols

func choice(Z map[XStruct][9]YStruct, Y map[YStruct][4]XStruct, r YStruct) {
	// cols := make([][9]YStruct, len(Y[r]))
	for _, v := range Y[r] {
		for _, vv := range Z[v] {
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
		// if isReturnedPop == false {
		// 	delete(Z, v)
		// } else {
		// 	cols[i] = Z[v]
		// }
	}

	// if isReturnedPop == false {
	// 	return nil
	// } else {
	// 	return cols
	// }
}

// def solve(X, Y, solution):
//     if not X:
//         yield list(solution)
//     else:
//         c = min(X, key=lambda c: len(X[c]))
//         for r in list(X[c]):
//             solution.append(r)
//             cols = select(X, Y, r)
//             for s in solve(X, Y, solution):
//                 yield s
//             deselect(X, Y, r, cols)
//             solution.pop()

func solve(Z map[XStruct][9]YStruct, Y map[YStruct][4]XStruct, solution []YStruct) {
	i := 0
	// fmt.Println(Z)
	for i < len(solution) {
		min := 10
		c := XStruct{}
		counter := 0
		isNullYStruct := YStruct{0, 0, 0}
		// fmt.Println(Z)
		for key, v := range Z {
			// fmt.Println(key, v)
			for _, vv := range v {
				if vv != isNullYStruct {
					counter++
				}
			}
			if counter == 0 {
				continue
			}
			fmt.Println(counter, min, key, v)
			if counter < min {
				min = counter
				c = key

				if min == 1 {
					break
				}
			} else {
				fmt.Println(counter, min)
			}
			counter = 0
		}
		// fmt.Println()

		// if min == 0 {
		// 	break
		// }
		// fmt.Println(min, c, Z[c])
		// fmt.Println()
		r := YStruct{}
		// fmt.Println(c, Z[c])
		for _, v := range Z[c] {
			if v != isNullYStruct {
				r = v
				break
			}
		}
		// r := Z[c][0]
		solution[i] = r
		// fmt.Println(solution)
		choice(Z, Y, r)
		i++

	}
}

func sudokuSolve(grid [9][9]int) [9][9]int {
	R, C := 3, 3
	// fmt.Println(R, C)
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
	zeroCount := 0
	for i, row := range grid {
		// fmt.Println(i, row)
		for j, n := range row {
			if n == 0 {
				zeroCount++
			} else {
				yStruct := YStruct{i, j, n}
				choice(Z, Y, yStruct)
			}
		}
	}
	// fmt.Println(Z)
	solution := make([]YStruct, zeroCount)
	// fmt.Println(solution)
	solve(Z, Y, solution)
	fmt.Println(solution)
	for _, v := range solution {
		grid[v.R][v.C] = v.N
	}
	return grid

	// 	for (r, c, n) in a:
	// 	grid[r][c] = n
	// return grid

	// solution := make([]YStruct, zeroCount)
	// fmt.Println(solution)

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
	fmt.Println(sudokuSolve(grid))
}
