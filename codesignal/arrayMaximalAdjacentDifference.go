package main

import (
	"fmt"
	"sort"
)

func main() {
	test := []int{2, 4, 1, 0}

	var diffArry []int

	diff := 0

	if len(test) <= 1 {
		return
	}

	diff = test[0] - test[1]

	diffArry = append(diffArry, diff)

	for i := 1; i < len(test)-1; i++ {
		diff = test[i] - test[i-1]

		diffArry = append(diffArry, diff)

		diff = test[i] - test[i+1]

		diffArry = append(diffArry, diff)
	}

	diff = test[len(test)-1] - test[len(test)-2]

	diffArry = append(diffArry, diff)

	sort.Ints(diffArry)

	fmt.Println(diffArry)

}
