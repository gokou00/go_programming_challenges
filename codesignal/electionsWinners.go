package main

import (
	"fmt"
	//"sort"
)

func main() {
	votes := []int{3, 1, 1, 3, 1}
	k := 2
	isTrue := true
	count := 0

	for i, val := range votes {
		temp := val + k
		for j, val2 := range votes {
			if i == j {
				continue
			}
			if temp <= val2 {
				isTrue = false
				break
			}

		}

		if isTrue {
			count++
		}
		isTrue = true

	}

	fmt.Println(count)

}
