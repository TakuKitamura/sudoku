package api

import (
	"net/http"
	"sudoku/src/util"

	"github.com/gin-gonic/gin"
)

type CanSolveSudokuResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func SudokuValidCheckAPI(c *gin.Context) {
	sudokuSolveRequest := SudokuSolveRequest{}
	err := c.BindJSON(&sudokuSolveRequest)
	if err != nil {
		util.LogUnexpectedErr(err)
		return
	}

	_, _, cannotSolveSudokuResponse, err := sudokuValidCheck(sudokuSolveRequest.Problem, true)

	if err != nil {
		c.JSON(http.StatusOK, cannotSolveSudokuResponse)
		return
	}

	sudokuSolveOKResponse := CanSolveSudokuResponse{
		Status:  "ok",
		Message: "can solve sudoku!",
	}

	util.APIStatusOK(c, sudokuSolveOKResponse)
	return
}
