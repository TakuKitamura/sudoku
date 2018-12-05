package api

import (
	"net/http"
	"sudoku/src/util"
	"time"

	"github.com/gin-gonic/gin"
)

type SudokuSolveRequest struct {
	Problem [9][9]uint8 `json:"problem"`
}

type XStruct struct {
	KeyName string
	Values  [2]uint8
}

type XArray []XStruct

type YStruct struct {
	R uint8
	C uint8
	N uint8
}

func exactCover(X [324]XStruct, Y map[YStruct][4]XStruct, Z map[XStruct][9]YStruct) {
	for j := uint16(0); j < uint16(len(X)); j++ {
		XJ := &X[j]
		yStruct := [9]YStruct{}
		i := uint8(0)
		for key, vv := range Y {
			for k := uint8(0); k < uint8(len(vv)); k++ {
				if *XJ == vv[k] {
					yStruct[i] = key
					if i == 8 {
						Z[*XJ] = yStruct
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
	for i := uint8(0); i < uint8(len(Y[r])); i++ {
		for _, vv := range Z[Y[r][i]] {
			if vv == isNullYStruct {
				continue
			}
			for _, vvv := range Y[vv] {
				if vvv != Y[r][i] {
					ZV := Z[vvv]
					for j := uint8(0); j < uint8(len(ZV)); j++ {
						if ZV[j] == vv {
							ZV[j] = YStruct{0, 0, 0}
							Z[vvv] = ZV
							break
						}
					}
				}
			}
		}
		delete(Z, Y[r][i])
	}
}

func solve(Z map[XStruct][9]YStruct, Y map[YStruct][4]XStruct, solution []YStruct) {
	i := uint8(0)
	for i < uint8(len(solution)) {
		isNullYStruct := YStruct{0, 0, 0}
		for key, v := range Z {
			counter := uint8(0)
			for i := uint8(0); i < uint8(len(v)); i++ {
				if v[i] != isNullYStruct {
					counter++
				}
			}
			min := uint8(10)
			if counter < min {
				min = counter
				if min == 1 {
					for j := uint8(0); j < uint8(len(Z[key])); j++ {
						if Z[key][j] != isNullYStruct {
							solution[i] = Z[key][j]
							choice(Z, Y, solution[i])
							i++
							break
						}
					}
					break
				}
			}
		}
	}
}

func sudokuSolve(problem [9][9]uint8) [9][9]uint8 {
	N := uint8(9)
	X := [324]XStruct{} // [N * N * 4]
	keyNames := [4]string{"rc", "rn", "cn", "bn"}
	k := uint16(0)
	for i := uint8(0); i < uint8(N); i++ {
		for j := uint8(0); j < uint8(N); j++ {
			X[k] = XStruct{KeyName: keyNames[0], Values: [2]uint8{i, j}}
			k++
		}
	}
	for h := uint8(0); h < 3; h++ {
		for i := uint8(0); i < uint8(N); i++ {
			for j := uint8(1); j < uint8(N+1); j++ {
				X[k] = XStruct{KeyName: keyNames[h+1], Values: [2]uint8{i, j}}
				k++
			}
		}
	}
	Y := map[YStruct][4]XStruct{}
	for h := uint8(0); h < uint8(N); h++ {
		for i := uint8(0); i < uint8(N); i++ {
			for j := uint8(1); j < uint8(N+1); j++ {
				R, C := uint8(3), uint8(3)
				b := (h/R)*R + (i / C)
				yStruct := YStruct{h, i, j}
				v := [2]uint8{h, i}
				rc := XStruct{KeyName: keyNames[0], Values: v}
				v = [2]uint8{h, j}
				rn := XStruct{KeyName: keyNames[1], Values: v}
				v = [2]uint8{i, j}
				cn := XStruct{KeyName: keyNames[2], Values: v}
				v = [2]uint8{b, j}
				bn := XStruct{KeyName: keyNames[3], Values: v}
				Y[yStruct] = [4]XStruct{rc, rn, cn, bn}
			}
		}
	}

	Z := map[XStruct][9]YStruct{}
	exactCover(X, Y, Z)
	zeroCount := uint8(0)
	for i := uint8(0); i < uint8(len(problem)); i++ {
		for j := uint8(0); j < uint8(len(problem[i])); j++ {
			if problem[i][j] == 0 {
				zeroCount++
			} else {
				choice(Z, Y, YStruct{i, j, problem[i][j]})
			}
		}
	}
	solution := make([]YStruct, zeroCount)
	solve(Z, Y, solution)
	for i := uint8(0); i < uint8(len(solution)); i++ {
		problem[solution[i].R][solution[i].C] = solution[i].N
	}
	return problem
}

func SudokuSolveAPI(c *gin.Context) {
	sudokuSolveRequest := SudokuSolveRequest{}
	err := c.BindJSON(&sudokuSolveRequest)
	if err != nil {
		util.LogUnexpectedErr(err)
		return
	}

	startTime := time.Now()
	answer := sudokuSolve(sudokuSolveRequest.Problem)
	time := time.Now().Sub(startTime).Seconds()
	c.JSON(http.StatusOK, gin.H{
		"answer": answer,
		"time":   time,
	})
	// errStruct := util.ErrStruct{}
	// util.RequestInfo(c, util.ErrStruct{}, sudokuSolveRequest)
	// util.APIErr(c, errStruct, sudokuSolveRequest)
	return
}
