package main

import "fmt"

func main() {
	inputArray := []int{5, 3, 2, 3, 4}
	l := 1
	r := 1
	var finalArr []int

	for i := 0; i < len(inputArray); i++ {
		if i >= l && i <= r {
			continue
		}

		finalArr = append(finalArr, inputArray[i])

	}

	fmt.Println(finalArr)

}
