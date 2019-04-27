package main

import (
	"fmt"
	"sort"
)

func main() {
	test := []int{2, 3, 5, 1, 6}

	k := 2

	count := 0
	count2 := 0
	//var ansArray []int
	total := 0
	maxvar := 0

	if k == 1 {
		sort.Ints(test)
		fmt.Println(test[len(test)-1])
		return
	}

	for i := 0; i < len(test); i++ {
		count++
		if count == k {
			count2++
			total += test[i]
			if maxvar == 0 {
				maxvar = total
			}

			if total > maxvar {
				maxvar = total
			}

			//fmt.Println(total)
			//ansArray = append(ansArray, total)
			count = 0
			total = 0
			count++
			i = count2
			//	continue
		}

		total += test[i]

		//fmt.Println(i)

	}

	//fmt.Println(ansArray)

	//sort.Ints(ansArray)
	//fmt.Println(ansArray)

	fmt.Println(maxvar)

}
