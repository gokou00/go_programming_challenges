package main

import (
	"fmt"
	"sort"
)

func main() {
	test := []int{3, 2, 1, 1}

	k := 1

	count := 0
	var ansArray []int
	total := 0

	for i := 0; i < len(test); i++ {
		for j := i; j < k+i && j < len(test); j++ {

			//fmt.Println("test")
			total += test[j]
			count++
		}

		if count == k {
			ansArray = append(ansArray, total)
		}
		total = 0
		count = 0

	}

	fmt.Println(ansArray)

	sort.Ints(ansArray)
	fmt.Println(ansArray)

	fmt.Println(ansArray[len(ansArray)-1])

}
