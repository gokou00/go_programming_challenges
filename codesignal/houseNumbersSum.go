package main

import "fmt"

func main() {
	inputArray := []int{10, 10, 10, 10, 10, 10, 10, 10, 0, 10}

	total := 0

	for _, x := range inputArray {
		if x == 0 {
			break
		}
		total += x

	}

	fmt.Println(total)

}
