package main

import "fmt"

func main() {
	inputArray := []int{1, 2, 1}
	elemToReplace := 1
	substitutionElem := 3

	var finalArry []int

	for _, x := range inputArray {
		if x == elemToReplace {
			finalArry = append(finalArry, substitutionElem)
			continue
		}

		finalArry = append(finalArry, x)
	}

	fmt.Println(finalArry)

}
