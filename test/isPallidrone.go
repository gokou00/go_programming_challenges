package main

import "fmt"

func main() {
	test := []int{1, 2, 2, 3}

	count := 0

	for i := len(test) - 1; i >= 0; i-- {

		if test[i] != test[count] {
			fmt.Println("false")
			return
		}
		count++
	}

	fmt.Println("true")
}
