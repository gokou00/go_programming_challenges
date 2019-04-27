package main

import (
	"fmt"
	"strings"
)

func main() {
	inputString := "12-34-56-78-9A-BG"

	inputArray := strings.Split(inputString, "-")

	fmt.Println(inputArray)

	if len(inputArray) != 6 {
		fmt.Println("false")
		return
	}

	isMac := true
	for _, x := range inputArray {
		if len(x) != 2 {
			fmt.Println("false")
			return
		}

		for _, y := range x {
			if string(y) >= "0" && string(y) <= "9" || string(y) >= "A" && string(y) <= "F" {
				continue
			} else {
				fmt.Println("false")
				return
			}

		}
	}

	fmt.Println(isMac)

}
