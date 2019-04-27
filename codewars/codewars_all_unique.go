package main

import (
	"fmt"
	"strings"
)

func main() {
	input := " nAa"
	input = strings.ToLower(input)

	var strMap map[string]int
	strMap = make(map[string]int)

	for _, letter := range input {
		fmt.Println(string(letter))
		strMap[string(letter)] = 0

	}

	for _, letter := range input {
		fmt.Println(string(letter))
		strMap[string(letter)]++

	}

	// go through map and find anything greater than 1 if so not unique

	for _, value := range strMap {

		if value > 1 {
			fmt.Println("false")
		}

	}

	fmt.Println(strMap)
}
