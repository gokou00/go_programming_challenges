package main

import (
	"fmt"
	"strconv"
	"strings"
)

//X'x go first 0

func main() {
	game := []string{"102",
		"457",
		"386"}

	//buildStr := ""

	gameCombine := ""

	for _, x := range game {
		gameCombine += string(x)
	}

	gameBoard := "........."

	for strings.Contains(gameBoard, ".") {
		for i := 0; i < len(gameCombine); i++ {
			tempStr := strconv.Itoa(i)
			idx := strings.Index(gameCombine, tempStr)
			gameBoardSplit := strings.Split(gameBoard, "")

			if i%2 == 0 {
				gameBoardSplit[idx] = "X"

			} else {
				gameBoardSplit[idx] = "O"
			}

			gameBoard = strings.Join(gameBoardSplit, "")

			if string(gameBoard[0]) == "X" && string(gameBoard[1]) == "X" && string(gameBoard[2]) == "X" {
				fmt.Println("X")
				return
			}

			if string(gameBoard[0]) == "O" && string(gameBoard[1]) == "O" && string(gameBoard[2]) == "O" {
				fmt.Println("O")
				return
			}

			if string(gameBoard[3]) == "X" && string(gameBoard[4]) == "X" && string(gameBoard[5]) == "X" {
				fmt.Println("X")
				return
			}

			if string(gameBoard[3]) == "O" && string(gameBoard[4]) == "O" && string(gameBoard[5]) == "O" {
				fmt.Println("O")
				return
			}

			if string(gameBoard[6]) == "X" && string(gameBoard[7]) == "X" && string(gameBoard[8]) == "X" {
				fmt.Println("X")
				return
			}

			if string(gameBoard[6]) == "O" && string(gameBoard[7]) == "O" && string(gameBoard[8]) == "O" {
				fmt.Println("O")
				return
			}

			if string(gameBoard[0]) == "X" && string(gameBoard[3]) == "X" && string(gameBoard[6]) == "X" {
				fmt.Println("X")
				return
			}

			if string(gameBoard[0]) == "O" && string(gameBoard[3]) == "O" && string(gameBoard[6]) == "O" {
				fmt.Println("O")
				return
			}

			if string(gameBoard[1]) == "X" && string(gameBoard[4]) == "X" && string(gameBoard[7]) == "X" {
				fmt.Println("X")
				return
			}

			if string(gameBoard[1]) == "O" && string(gameBoard[4]) == "O" && string(gameBoard[7]) == "O" {
				fmt.Println("O")
				return
			}

			if string(gameBoard[2]) == "X" && string(gameBoard[5]) == "X" && string(gameBoard[8]) == "X" {
				fmt.Println("X")
				return
			}

			if string(gameBoard[2]) == "O" && string(gameBoard[5]) == "O" && string(gameBoard[8]) == "O" {
				fmt.Println("O")
				return
			}

			if string(gameBoard[0]) == "X" && string(gameBoard[4]) == "X" && string(gameBoard[8]) == "X" {
				fmt.Println("X")
				return
			}

			if string(gameBoard[0]) == "O" && string(gameBoard[4]) == "O" && string(gameBoard[8]) == "O" {
				fmt.Println("O")
				return
			}

			if string(gameBoard[2]) == "X" && string(gameBoard[4]) == "X" && string(gameBoard[6]) == "X" {
				fmt.Println("X")
				return
			}

			if string(gameBoard[2]) == "O" && string(gameBoard[4]) == "O" && string(gameBoard[6]) == "O" {
				fmt.Println("O")
				return
			}

		}

	} // end of loop

	fmt.Println("draw")

}
