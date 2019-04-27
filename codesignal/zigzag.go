package main

import "fmt"

func main() {
	//test := []int{9, 8, 8, 5, 3, 5, 3, 2, 8, 6}
	//	test := []int{1, 7, 4, 9, 2, 5}

	test := []int{2, 1, 4, 4, 1, 4, 4, 1, 2, 0, 1, 0, 0, 3, 1, 3, 4, 1, 3, 4}
	var zigArry []int

	temp := 0

	for i := 1; i < len(test); i++ {
		temp = test[i] - test[i-1]

		zigArry = append(zigArry, temp)
	}

	fmt.Println(zigArry)

}
