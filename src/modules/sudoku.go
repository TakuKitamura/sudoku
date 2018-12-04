package modules

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

func exactCover(X [324]XStruct, Y map[YStruct][4]XStruct, Z map[XStruct][9]YStruct) {
	i := 0
	yStruct := [9]YStruct{}
	for j := 0; j < len(X); j++ {
		for key, vv := range Y {
			for k := 0; k < len(vv); k++ {
				if X[j] == vv[k] {
					yStruct[i] = key
					if i == 8 {
						Z[X[j]] = yStruct
						i = 0
					} else {
						i++
					}
				}
			}
		}
	}
}

func choice(Z map[XStruct][9]YStruct, Y map[YStruct][4]XStruct, r YStruct) {
	isNullYStruct := YStruct{0, 0, 0}
	for i := 0; i < len(Y[r]); i++ {
		for _, vv := range Z[Y[r][i]] {
			if vv == isNullYStruct {
				continue
			}
			for _, vvv := range Y[vv] {
				if vvv != Y[r][i] {
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
		delete(Z, Y[r][i])
	}
}

func solve(Z map[XStruct][9]YStruct, Y map[YStruct][4]XStruct, solution []YStruct) {
	i := 0
	for i < len(solution) {
		min := 10
		// c := XStruct{}
		counter := 0
		isNullYStruct := YStruct{0, 0, 0}
		for key, v := range Z {
			for i := 0; i < len(v); i++ {
				if v[i] != isNullYStruct {
					counter++
				}
			}
			if counter < min {
				min = counter
				if min == 1 {
					r := YStruct{}
					for j := 0; j < len(Z[key]); j++ {
						if Z[key][j] != isNullYStruct {
							r = Z[key][j]
							break
						}
					}
					solution[i] = r
					choice(Z, Y, r)
					i++
					break
				}
			}
			counter = 0
		}

	}
}

func SudokuSolve(grid [9][9]int) [9][9]int {
	R, C := 3, 3
	N := 9
	X := [324]XStruct{} // [N * N * 4]
	keyNames := [4]string{"rc", "rn", "cn", "bn"}
	// nameIndex := 0
	k := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			X[k].Values = [2]int{i, j}
			X[k].KeyName = "rc"
			k++
		}
	}
	for h := 0; h < 3; h++ {
		for i := 0; i < N; i++ {
			for j := 1; j < N+1; j++ {
				X[k] = XStruct{KeyName: keyNames[h+1], Values: [2]int{i, j}}
				k++
			}
		}
	}
	Y := map[YStruct][4]XStruct{}
	yStructArray := [729]YStruct{}
	k = 0
	for h := 0; h < N; h++ {
		for i := 0; i < N; i++ {
			for j := 1; j < N+1; j++ {
				b := (h/R)*R + (i / C)
				// yStruct := YStruct{h, i, j}
				yStructArray[k] = YStruct{h, i, j}

				v := [2]int{h, i}

				rc := XStruct{KeyName: keyNames[0], Values: v}

				v = [2]int{h, j}

				rn := XStruct{KeyName: keyNames[1], Values: v}

				v = [2]int{i, j}

				cn := XStruct{KeyName: keyNames[2], Values: v}

				v = [2]int{b, j}

				bn := XStruct{KeyName: keyNames[3], Values: v}

				Y[yStructArray[k]] = [4]XStruct{rc, rn, cn, bn}

				k++
			}
		}
	}

	Z := map[XStruct][9]YStruct{}
	exactCover(X, Y, Z)
	zeroCount := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				zeroCount++
			} else {
				choice(Z, Y, YStruct{i, j, grid[i][j]})
			}
		}
	}
	solution := make([]YStruct, zeroCount)
	solve(Z, Y, solution)
	for i := 0; i < len(solution); i++ {
		grid[solution[i].R][solution[i].C] = solution[i].N
	}
	return grid
}
