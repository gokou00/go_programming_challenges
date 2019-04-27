package main

import (
	"fmt"
	"strconv"
)

func main() {

	/*
		matrix := [][]int{{9, 9, 9, 9, 9},
			{9, 9, 9, 9, 9},
			{9, 9, 9, 9, 9},
			{9, 9, 9, 9, 9},
			{9, 9, 9, 9, 9},
			{9, 9, 9, 9, 9},
		}

	*/

	matrix := [][]int{{1, 1, 1, 1, 1, 1, 2, 2, 2, 3, 3, 3, 9, 9, 9, 2, 3, 9}}

	total := 0
	testMap := make(map[string]int)
	buildStr := ""

	for i := 0; i < len(matrix)-2+1; i++ {
		for j := 0; j < len(matrix[0])-2+1; j++ {

			for p := i; p < 2+i; p++ {
				for m := j; m < 2+j; m++ {
					//fmt.Print(matrix[p][m])
					temp := strconv.Itoa(matrix[p][m])
					buildStr += temp

				}
				//fmt.Println()
			}

			testMap[buildStr]++
			buildStr = ""
			total++

		}

	}

	//fmt.Println(total)
	//fmt.Println(testMap)
	fmt.Println(len(testMap))

}
