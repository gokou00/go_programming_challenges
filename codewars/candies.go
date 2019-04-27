package main

import "fmt"

func main() {
	m := 5  // candies
	n := 10 //kids

	finalArray := make([]int, n)
	//	fmt.Println(finalArray)

	for m > 0 {
		for i := 0; i < len(finalArray) && m > 0; i++ {
			finalArray[i]++
			m--
		}
	}
	fmt.Println(finalArray)

}
