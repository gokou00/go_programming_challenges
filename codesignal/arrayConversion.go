package main

import "fmt"

func main() {
	inputArray := []int{1, 2, 3, 4, 5, 6, 7, 8}
	count := 1

	for len(inputArray) != 1 {
		if count%2 == 0 {
			var temp []int
			for i := 0; i < len(inputArray)-1; i += 2 {
				temp = append(temp, inputArray[i]*inputArray[i+1])
			}
			inputArray = make([]int, len(temp))
			copy(inputArray, temp)
		} else {
			var temp []int
			for i := 0; i < len(inputArray)-1; i += 2 {
				temp = append(temp, inputArray[i]+inputArray[i+1])
			}
			inputArray = make([]int, len(temp))
			copy(inputArray, temp)

		}

		//fmt.Println(inputArray)
		count++

	}

	fmt.Println(inputArray)

}
