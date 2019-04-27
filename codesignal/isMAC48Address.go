package main

import (
	"fmt"
	"strings"
)

func main() {

	inputString := "12-34-56-78-9A-BG"
	inputStringArr := strings.Split(inputString, "-")

	if len(inputStringArr) != 6 {
		fmt.Println("False")
		return
	}

	for _, x := range inputStringArr {
		if len(string(x)) != 2 {
			fmt.Println("False")
			return
		}
		for _, j := range string(x) {

			if string(j) >= "0" && string(j) <= "9" || string(j) >= "A" && string(j) <= "F" {
				continue
			} else {
				fmt.Println("False")
				return
			}
		}
	}

	fmt.Println("True")

}
