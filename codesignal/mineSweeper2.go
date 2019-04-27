package main

import "fmt"

func main() {
	matrix := [][]bool{
		{true, false, false, true},
		{false, false, true, false},
		{true, true, false, true},
	}

	//creating the 2d int array
	result := make([][]int, len(matrix)+2)

	for i := range result {
		result[i] = make([]int, len(matrix[0])+2)
	}

	//creating the 2d int array
	result2 := make([][]int, len(matrix))

	for i := range result2 {
		result2[i] = make([]int, len(matrix[0]))
	}

	fmt.Println(matrix)
	fmt.Println(result)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == true {
				resRow := i + 1
				resCol := j + 1
				//	result[resRow][resCol] = 1
				result[resRow][resCol+1] += 1
				result[resRow-1][resCol+1] += 1
				result[resRow-1][resCol] += 1
				result[resRow-1][resCol-1] += 1
				result[resRow][resCol-1] += 1
				result[resRow+1][resCol-1] += 1
				result[resRow+1][resCol] += 1
				result[resRow+1][resCol+1] += 1
			}
		}
	}

	fmt.Println(result)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			result2[i][j] = result[i+1][j+1]
		}
	}

	/*
		for i := 0; i < len(matrix); i++ {
			for j := 0; j < len(matrix[0]); j++ {
				if matrix[i][j] == true {
					result2[i][j] = 1
				}
			}
		}
	*/
	fmt.Println(result2)

}
