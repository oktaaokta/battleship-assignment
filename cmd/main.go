package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

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

	boardPlayer1 := make([][]string, 0)
	boardPlayer2 := make([][]string, 0)
	count := 1

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		switch parts[0] {
		case "create_board":
			gridSize, _ := strconv.ParseInt(parts[1], 10, 64)
			boardPlayer1 = internal.CreateBoard(int(gridSize))
			boardPlayer2 = internal.CreateBoard(int(gridSize))
		default:
			if count == 3 || count == 4 {
				if count%2 == 0 {
					boardPlayer2 = internal.InitializeShips(boardPlayer2, parts[0])
				} else {
					boardPlayer1 = internal.InitializeShips(boardPlayer1, parts[0])
				}
			} else if count == 6 || count == 7 {
				if count%2 == 1 {
					boardPlayer1 = internal.PlaceMissiles(boardPlayer1, parts[0])
				} else {
					boardPlayer2 = internal.PlaceMissiles(boardPlayer2, parts[0])
				}
			}
		}
		count++
	}

	var hit1, hit2 int

	fmt.Println("Player1")
	for i := 0; i < len(boardPlayer1); i++ {
		for j := 0; j < len(boardPlayer1[0]); j++ {
			if boardPlayer1[i][j] == "" {
				fmt.Print("_ ")
				continue
			} else if boardPlayer1[i][j] == "X" {
				hit2++
			}
			fmt.Printf("%v ", boardPlayer1[i][j])
		}
		fmt.Print("\n")
	}

	fmt.Println("Player2")
	for i := 0; i < len(boardPlayer2); i++ {
		for j := 0; j < len(boardPlayer2[0]); j++ {
			if boardPlayer2[i][j] == "" {
				fmt.Print("_ ")
				continue
			} else if boardPlayer2[i][j] == "X" {
				hit1++
			}
			fmt.Printf("%v ", boardPlayer2[i][j])
		}
		fmt.Print("\n")
	}
	fmt.Printf("P1: %v \n", hit1)
	fmt.Printf("P2: %v \n", hit2)
	if hit1 > hit2 {
		fmt.Print("Player 1 Wins")
	} else if hit1 < hit2 {
		fmt.Print("Player 2 Wins")
	} else {
		fmt.Print("It is a draw")
	}
}
