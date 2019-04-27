package main

import "fmt"
import "strconv"

func main() {
	colMap := make(map[string]int)
	colMap["a"] = 1
	colMap["b"] = 2
	colMap["c"] = 3
	colMap["d"] = 4
	colMap["e"] = 5
	colMap["f"] = 6
	colMap["g"] = 7
	colMap["h"] = 8

	cell := "a1"
	col := 0
	row := 0
	count := 0
	tempCol := 0
	tempRow := 0

	col = colMap[string(cell[0])]

	num, _ := strconv.Atoi(string(cell[1]))

	//fmt.Println(row, col)

	row = num

	//	fmt.Println(row, col)

	//row = row - 1
	//	col = col - 1

	// hori

	tempCol = col + 2
	tempRow = row + 1

	if tempRow > 0 && tempRow <= 8 && tempCol > 0 && tempCol <= 8 {
		count++
	}

	tempCol = col + 2
	tempRow = row - 1

	if tempRow > 0 && tempRow <= 8 && tempCol > 0 && tempCol <= 8 {
		count++
	}

	tempCol = col - 2
	tempRow = row + 1

	if tempRow > 0 && tempRow <= 8 && tempCol > 0 && tempCol <= 8 {
		count++
	}

	tempCol = col - 2
	tempRow = row - 1

	if tempRow > 0 && tempRow <= 8 && tempCol > 0 && tempCol <= 8 {
		count++
	}

	//vert
	tempRow = row + 2
	tempCol = col + 1

	if tempRow > 0 && tempRow <= 8 && tempCol > 0 && tempCol <= 8 {
		count++
	}

	tempRow = row + 2
	tempCol = col - 1

	if tempRow > 0 && tempRow <= 8 && tempCol > 0 && tempCol <= 8 {
		count++
	}

	tempRow = row - 2
	tempCol = col + 1

	if tempRow > 0 && tempRow <= 8 && tempCol > 0 && tempCol <= 8 {
		count++
	}

	tempRow = row - 2
	tempCol = col - 1

	if tempRow > 0 && tempRow <= 8 && tempCol > 0 && tempCol <= 8 {
		count++
	}

	fmt.Println(count)

}
