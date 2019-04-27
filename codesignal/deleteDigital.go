package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	n := 45553333
	tempStr := ""
	//tempNum := 0
	numStr := strconv.Itoa(n)

	var numArray []int

	//try using a nested loop

	for i := 0; i < len(numStr); i++ {
		for j := 0; j < len(numStr); j++ {
			if j == i {
				continue
			}

			tempStr += string(numStr[j])

		}

		num, _ := strconv.Atoi(tempStr)

		numArray = append(numArray, num)
		tempStr = ""

	}

	sort.Ints(numArray)

	fmt.Println(numArray)
	fmt.Println(numArray[len(numArray)-1])

}
