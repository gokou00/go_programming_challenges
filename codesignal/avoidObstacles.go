package main

//Needs to sort array first
//create a array with the gaps if any from the input array
//first element in array is the first distance from 0.
//check the different from i+1 to i.
//if the distance is greater than jump add that to count.

import (
	"fmt"
	"sort"
)

func main() {
	test := []int{5, 3, 6, 7, 9}

	//	gap := 0

	count := 0
	final := 0

	//	var gapsArry []int

	sort.Ints(test)

	max := test[len(test)-1]

	for i := 1; i <= max; i++ {
		for _, x := range test {
			if x%i == 0 {
				break
			} else {
				count++
			}
		}

		fmt.Println("count", count, "i", i)
		if count == len(test) {
			final = i
			break
		}

		count = 0

	}

	if final == 0 {
		final = test[len(test)-1] + 1
	}

	fmt.Println(final)

}
