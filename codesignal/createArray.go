package main

import "fmt"

func main() {

	size := 49
	var numArray []int

	for i := 0; i < size; i++ {
		numArray = append(numArray, 1)
	}

	fmt.Println(numArray)

}
