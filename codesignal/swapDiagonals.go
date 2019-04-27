package main

import "fmt"

func main() {
	matrix := [][]int{{3, 2, 1}, {4, 5, 6}, {9, 8, 7}}
	j := len(matrix)

	for i := 0; i < j; i++ {
		matrix[i][i], matrix[i][j-i-1] = matrix[i][j-i-1], matrix[i][i]
	}

	return matrix

}
