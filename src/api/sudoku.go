package api

import (
	"net/http"
	"sudoku/src/util"
	"time"

	"github.com/gin-gonic/gin"
)

type SudokuSolveRequest struct {
	Problem [9][9]int `json:"problem"`
}

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
	for j := 0; j < len(X); j++ {
		XJ := &X[j]
		yStruct := [9]YStruct{}
		i := 0
		for key, vv := range Y {
			for k := 0; k < len(vv); k++ {
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
	i := 0
	for i < len(solution) {
		isNullYStruct := YStruct{0, 0, 0}
		for key, v := range Z {
			counter := 0
			for i := 0; i < len(v); i++ {
				if v[i] != isNullYStruct {
					counter++
				}
			}
			min := 10
			if counter < min {
				min = counter
				if min == 1 {
					for j := 0; j < len(Z[key]); j++ {
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

func sudokuSolve(problem [9][9]int) [9][9]int {
	N := 9
	X := [324]XStruct{} // [N * N * 4]
	keyNames := [4]string{"rc", "rn", "cn", "bn"}
	k := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			X[k] = XStruct{KeyName: keyNames[0], Values: [2]int{i, j}}
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
				R, C := 3, 3
				b := (h/R)*R + (i / C)
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
	for i := 0; i < len(problem); i++ {
		for j := 0; j < len(problem[i]); j++ {
			if problem[i][j] == 0 {
				zeroCount++
			} else {
				choice(Z, Y, YStruct{i, j, problem[i][j]})
			}
		}
	}
	solution := make([]YStruct, zeroCount)
	solve(Z, Y, solution)
	for i := 0; i < len(solution); i++ {
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
	util.RequestInfo(c, util.ErrStruct{}, sudokuSolveRequest)

	startTime := time.Now()
	answer := sudokuSolve(sudokuSolveRequest.Problem)
	time := time.Now().Sub(startTime).Seconds()
	c.JSON(http.StatusOK, gin.H{
		"answer": answer,
		"time":   time,
	})
	errStruct := util.ErrStruct{}

	util.APIErr(c, errStruct, sudokuSolveRequest)
	return
}
