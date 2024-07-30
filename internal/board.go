package internal

import (
	"strconv"
	"strings"
)

func CreateBoard(gridSize int) [][]string {
	board := make([][]string, gridSize)
	for i := range board {
		board[i] = make([]string, gridSize)
	}

	return board
}

func InitializeShips(board [][]string, positions string) [][]string {
	positionsList := strings.Split(positions, ":")

	for i := range positionsList {
		grids := strings.Split(positionsList[i], ",")
		row, _ := strconv.ParseInt(grids[0], 10, 64)
		col, _ := strconv.ParseInt(grids[1], 10, 64)
		board[row][col] = "B"
	}

	return board
}

func PlaceMissiles(board [][]string, positions string) [][]string {
	positionsList := strings.Split(positions, ":")

	for i := range positionsList {
		grids := strings.Split(positionsList[i], ",")
		row, _ := strconv.ParseInt(grids[0], 10, 64)
		col, _ := strconv.ParseInt(grids[1], 10, 64)

		if board[row][col] == "B" {
			board[row][col] = "X"
		} else {
			board[row][col] = "O"
		}
	}

	return board
}
