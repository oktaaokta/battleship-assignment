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

	scanner.Scan()
	line = scanner.Text()
	totalMissiles, _ := strconv.Atoi(line)
	boardPlayer1.InitMissiles(totalMissiles)
	boardPlayer2.InitMissiles(totalMissiles)

	for i := 0; i < 2; i++ {
		scanner.Scan()
		line := scanner.Text()
		if i == 0 {
			boardPlayer1.PlaceMissiles(line, boardPlayer2)
		} else {
			boardPlayer2.PlaceMissiles(line, boardPlayer1)
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
			_, err = writer.WriteString(fmt.Sprintf("%v ", string(boardPlayer1.GameBoard[i][j])))
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}
		}
		_, err = writer.WriteString("\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	_, err = writer.WriteString("Player2" + "\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	for i := 0; i < len(boardPlayer2.GameBoard); i++ {
		for j := 0; j < len(boardPlayer2.GameBoard[0]); j++ {
			_, err = writer.WriteString(fmt.Sprintf("%v ", string(boardPlayer2.GameBoard[i][j])))
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}
		}
		_, err = writer.WriteString("\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	_, err = writer.WriteString(fmt.Sprintf("\nP1:%v \n", boardPlayer2.TotalMissiles-boardPlayer2.AvailableShips))
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	_, err = writer.WriteString(fmt.Sprintf("P2:%v \n", boardPlayer1.TotalMissiles-boardPlayer1.AvailableShips))
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	var text string
	if boardPlayer1.AvailableShips > boardPlayer2.AvailableShips {
		text = "Player 1 Wins"
	} else if boardPlayer1.AvailableShips < boardPlayer2.AvailableShips {
		text = "Player 2 Wins"
	} else {
		text = "It is a draw"
	}
	_, err = writer.WriteString(text)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing buffer:", err)
		return
	}
}
