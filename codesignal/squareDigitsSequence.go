package main

import (
	"fmt"
	"math"
)

func main() {

	a0 := 1
	total := 0
	var arr []int
	count := 0

	for contains(a0, arr) == false {
		arr = append(arr, a0)
		count++

		total = 0
		for a0 != 0 {
			pop := a0 % 10
			total += int(math.Pow(float64(pop), 2))
			a0 /= 10

		}

		a0 = total
		//fmt.Println(a0)
		//arr = append(arr, total)

	}

	fmt.Println(arr)
	fmt.Println(count + 1)

}

func contains(num int, arr []int) bool {
	for _, x := range arr {
		if x == num {
			//fmt.Println("test")
			return true
		}
	}

	return false

}
