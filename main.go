package main

import (
	"log"
	"os"
	"sudoku/src/api"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.DisableConsoleColor()
	os.Setenv("DEBUG", "true")
}

func main() {

	r := gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	v00 := r.Group("/v0.0")
	{
		v00.POST("/sudoku/solve", api.SudokuSolveAPI)
		v00.POST("/sudoku/check", api.SudokuValidCheckAPI)
	}

	r.Run(":8080")
}
