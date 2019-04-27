package main

import (
	"fmt"
	"strconv"
)

func main() {
	num1 := 123456789
	num2 := 100

	ans := num1 / num2

	ansStr := strconv.Itoa(ans)

	if len(ansStr) <= 3 {
		fmt.Println(ansStr)
		return
	}

	revStr := ""

	for i := len(ansStr) - 1; i >= 0; i-- {
		revStr += string(ansStr[i])

	}

	buildStr := ""

	for i := 0; i < len(revStr); i++ {
		if i%3 == 0 && i != 0 {
			buildStr += ","
		}

		buildStr += string(revStr[i])
	}

	rightStr := ""

	for i := len(buildStr) - 1; i >= 0; i-- {
		rightStr += string(buildStr[i])
	}

	fmt.Println(rightStr)

}
