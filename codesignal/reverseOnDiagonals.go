package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	i := 0
	j := len(matrix)

	for i < j {
		temp1 := matrix[i][i]
		matrix[i][i] = matrix[j-1][j-1]
		matrix[j-1][j-1] = temp1

		temp2 := matrix[i][j-1]
		matrix[i][j-1] = matrix[j-1][i]
		matrix[j-1][i] = temp2

		i++
		j--
	}

	fmt.Println(matrix)

}
