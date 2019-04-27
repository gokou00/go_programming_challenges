package main

import "fmt"
import "strconv"

func main() {
	//create a 2d array to mirror the chess board the values are strings

	chessBoard := make([][]string, 8)

	for i := range chessBoard {
		chessBoard[i] = make([]string, 8)
	}

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if i%2 == 0 {
				if j%2 == 0 {
					chessBoard[i][j] = "B"
				} else {
					chessBoard[i][j] = "W"
				}
			} else {
				if j%2 == 1 {
					chessBoard[i][j] = "B"
				} else {
					chessBoard[i][j] = "W"
				}
			}

		}
	}

	fmt.Println(chessBoard)

	//Parse the different strings

	cell1 := "D2"
	cell2 := "D2"

	row1 := string(cell1[0])

	var cell1Row int
	var cell2Row int

	switch row1 {
	case "A":
		cell1Row = 0
	case "B":
		cell1Row = 1
	case "C":
		cell1Row = 2
	case "D":
		cell1Row = 3
	case "E":
		cell1Row = 4
	case "F":
		cell1Row = 5
	case "G":
		cell1Row = 6
	case "H":
		cell1Row = 7
	}

	row2 := string(cell2[0])

	switch row2 {
	case "A":
		cell2Row = 0
	case "B":
		cell2Row = 1
	case "C":
		cell2Row = 2
	case "D":
		cell2Row = 3
	case "E":
		cell2Row = 4
	case "F":
		cell2Row = 5
	case "G":
		cell2Row = 6
	case "H":
		cell2Row = 7
	}

	col1 := string(cell1[1])
	col2 := string(cell2[1])

	col1Num, _ := strconv.Atoi(col1)
	col2Num, _ := strconv.Atoi(col2)

	ans1 := chessBoard[cell1Row][col1Num-1]
	ans2 := chessBoard[cell2Row][col2Num-1]

	fmt.Println(ans1, ans2)

}
