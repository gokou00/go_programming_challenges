package main

import "fmt"

func main() {
	a := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{4, 7, 8},
	}

	fmt.Println(len(a))
	fmt.Println(len(a[0]))

	finalArry := make([][]int, len(a[0]))

	for i := range finalArry {
		finalArry[i] = make([]int, len(a))
	}

	var numArry []int
	count := 0

	//fmt.Println(len(test))
	//fmt.Println(len(test[0]))

	for i := 0; i < len(a[0]); i++ {
		for j := len(a) - 1; j >= 0; j-- {
			//fmt.Println(test[j][i])

			numArry = append(numArry, a[j][i])
		}
	}

	for i := 0; i < len(finalArry); i++ {
		for j := 0; j < len(finalArry[0]); j++ {
			finalArry[i][j] = numArry[count]
			count++
		}
	}

	fmt.Println(finalArry)

	// create another for loop to put all the values in final array

}
