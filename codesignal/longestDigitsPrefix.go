package main

import (
	"fmt"
	"strconv"
)

func main() {
	inputString := "the output is 42"
	final := ""

	for _, x := range inputString {
		_, err := strconv.Atoi(string(x))
		if err != nil {
			break
		}

		final += string(x)

	}

	fmt.Println(final)

}
