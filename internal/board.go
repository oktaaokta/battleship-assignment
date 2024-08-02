package internal

import (
	"strconv"
	"strings"
)

func NewBoard() *Board {
	return &Board{
		GameBoard: [][]byte{},
	}
}

func (board *Board) CreateBoard(gridSize int) {
	gameBoard := make([][]byte, gridSize)
	for i := range gameBoard {
		gameBoard[i] = make([]byte, gridSize)
		for j := range gameBoard[i] {
			gameBoard[i][j] = '_'
		}
	}
	board.GameBoard = gameBoard
}

func (board *Board) InitShips(ships int) {
	board.AvailableShips = ships
}

func (board *Board) PlaceShips(positions string) {
	gameBoard := board.GameBoard
	positionsList := strings.Split(positions, ":")

	for i := range positionsList {
		grids := strings.Split(positionsList[i], ",")
		row, _ := strconv.ParseInt(grids[0], 10, 64)
		col, _ := strconv.ParseInt(grids[1], 10, 64)
		gameBoard[row][col] = '_'
	}

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
