package main

import (
	"fmt"
	//"strconv"
)

func main() {
	l := 10
	r := 12
	addedI := 0
	addedJ := 0
	count := 0
	//temp := ""

	//testMap := make(map[string]int)
	//var test []int

	for i := l; i <= r; i++ {
		addedI = adddigits(i)
		//fmt.Println(added)
		for j := i + 1; j <= r; j++ {
			addedJ = adddigits(j)

			if (i <= j+addedJ) && (i >= j-addedJ) && (j <= i+addedI) && (j >= i-addedI) {
				count++
			}

			//fmt.Println(j)

			//test = append(test, j)
		}
		//fmt.Println(test)

		//test = []int{}
	}
	fmt.Println(count)
	//fmt.Println(len(testMap))

}

func adddigits(num int) int {
	total := 0
	pop := 0
	for num != 0 {
		pop = num % 10
		total += pop
		num /= 10
	}
	return total
}
