package main

import "fmt"

func main() {
	inputArray := []string{"aba", "aa", "ad", "vcd", "aba"}

	var finalArr []string

	largest := 0

	for _, x := range inputArray {

		if len(x) > largest {
			largest = len(x)
		}

	}

	for _, x := range inputArray {
		if len(x) == largest {
			finalArr = append(finalArr, x)

		}
	}

	fmt.Println(finalArr)

}
