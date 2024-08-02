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

	for i := 0; i < 2; i++ {
		scanner.Scan()
		line := scanner.Text()
		if i == 0 {
			boardPlayer1.PlaceShips(line)
		} else {
			boardPlayer2.PlaceShips(line)
		}
	}

	outputFile, err := os.Create("../output/output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)

	_, err = writer.WriteString("Player1" + "\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	for i := 0; i < len(boardPlayer1.GameBoard); i++ {
		for j := 0; j < len(boardPlayer1.GameBoard[0]); j++ {
			fmt.Printf("%v ", string(boardPlayer1.GameBoard[i][j]))
			_, err = writer.WriteString(fmt.Sprintf("%v ", string(boardPlayer1.GameBoard[i][j])))
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}
		}
		fmt.Print("\n")
	}

	fmt.Println("Player2")
	for i := 0; i < len(boardPlayer2.GameBoard); i++ {
		for j := 0; j < len(boardPlayer2.GameBoard[0]); j++ {
			fmt.Printf("%v ", string(boardPlayer2.GameBoard[i][j]))
			_, err = writer.WriteString(fmt.Sprintf("%v ", string(boardPlayer2.GameBoard[i][j])))
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}
		}
		fmt.Print("\n")
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing buffer:", err)
		return
	}
}
