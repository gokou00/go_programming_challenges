package main

import "fmt"

func main() {
	matrix := [][]bool{
		{true, false, false},
		{false, true, false},
		{false, false, false},
	}

	//creating the 2d int array
	result := make([][]int, len(matrix))

	for i := range result {
		result[i] = make([]int, len(matrix[0]))
	}

	fmt.Println(matrix)
	fmt.Println(result)

	fmt.Println(result)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == true {
				if i == 0 && j == 0 {
					result[i][j+1] += 1
					result[i+1][j] += 1
					result[i+1][j+1] += 1
					result[i][j] += 1
				} else if i == 0 && j > 0 && j < len(matrix[0]) {
					result[i][j] += 1
					result[i][j-1] += 1
					result[i][j+1] += 1
					result[i+1][j-1] += 1
					result[i+1][j+1] += 1
				} else if i == 0 && j == len(matrix[0])-1 {
					result[i][j] += 1
					result[i][j-1] += 1
					result[i+1][j-1] += 1
					result[i+1][j] += 1
				}

			}
		}
	}

}
