package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 103456789

	numStr := strconv.Itoa(n)

	isNonZero := false

	for i := len(numStr) - 1; i >= 0; i-- {
		if isNonZero && string(numStr[i]) == "0" {
			fmt.Println("true")
			return
		}

		if string(numStr[i]) != "0" {
			isNonZero = true
		}

	}

	fmt.Println("false")

}
