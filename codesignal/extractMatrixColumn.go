package main

import "fmt"

func main() {

	matrix := [][]int{{1, 1, 1, 2}, {0, 5, 0, 4}, {2, 1, 3, 6}}
	column := 2

	var arr []int

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if j == column {
				arr = append(arr, matrix[i][j])
			}
		}
	}

	fmt.Println(arr)

}
