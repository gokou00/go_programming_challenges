package main

import (
	"fmt"
	"sort"
)

func main() {
	statues := []int{1}
	count := 0
	isFound := false

	sort.Ints(statues)

	start := statues[0]
	end := statues[len(statues)-1]

	for i := start; i <= end; i++ {
		for j := 0; j < len(statues); j++ {
			if i == statues[j] {
				isFound = true
			}

		}

		if isFound {
			isFound = false
		} else {
			count++
		}

	}

	fmt.Println(count)

}
