package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	internal "github.com/machine-coding/internal"
)

func main() {
	file, err := os.Open("../input/input.txt")
	if err != nil {
		log.Printf("Error when opening file: %v", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	line := scanner.Text()
	size, _ := strconv.Atoi(line)

	boardPlayer1 := internal.NewBoard()
	boardPlayer1.CreateBoard(int(size))
	boardPlayer2 := internal.NewBoard()
	boardPlayer2.CreateBoard(int(size))

	//move to next line
	scanner.Scan()
	line = scanner.Text()
	availableShips, _ := strconv.Atoi(line)
	boardPlayer1.InitShips(availableShips)
	boardPlayer2.InitShips(availableShips)

	fmt.Println("Player1")
	for i := 0; i < len(boardPlayer1.GameBoard); i++ {
		for j := 0; j < len(boardPlayer1.GameBoard[0]); j++ {
			fmt.Printf("%v ", string(boardPlayer1.GameBoard[i][j]))
		}
		fmt.Print("\n")
	}

	fmt.Println("Player2")
	for i := 0; i < len(boardPlayer2.GameBoard); i++ {
		for j := 0; j < len(boardPlayer2.GameBoard[0]); j++ {
			fmt.Printf("%v ", string(boardPlayer2.GameBoard[i][j]))
		}
		fmt.Print("\n")
	}
}
