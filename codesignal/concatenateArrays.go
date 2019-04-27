package main

import "fmt"

func main() {

	a := []int{2, 2, 1}
	b := []int{10, 11}

	var finalArr []int

	for _, x := range a {
		finalArr = append(finalArr, x)
	}

	for _, x := range b {
		finalArr = append(finalArr, x)
	}

	fmt.Println(finalArr)

}
