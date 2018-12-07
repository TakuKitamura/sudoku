package api

import (
	"math/rand"
)

func shuffle(oneRow [9]uint8) {
	for i := 8; i >= 0; i-- {
		j := rand.Intn(i + 1)
		oneRow[i], oneRow[j] = oneRow[j], oneRow[i]
	}
}

func sudokuGenerate() {
	oneRow := [9]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9}
	shuffle(oneRow)
}
