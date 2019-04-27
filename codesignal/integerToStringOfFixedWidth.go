package main

import (
	"fmt"
	"strconv"
)

func main() {
	number := 0
	width := 1
	numStr := ""
	revStr := ""

	for (number != 0) && (len(numStr) != width) {
		pop := number % 10
		numStr += strconv.Itoa(pop)
		number /= 10

	}

	if len(numStr) < width {
		for len(numStr) < width {
			numStr += "0"
		}
	}

	for i := len(numStr) - 1; i >= 0; i-- {
		revStr += string(numStr[i])
	}

	fmt.Println(numStr)
	fmt.Println(revStr)

}
