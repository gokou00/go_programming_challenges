package main

import "fmt"

func main() {
	// divide the string into equal parts.

	inputString := "truetruetrue"

	if len(inputString)%2 != 0 {
		fmt.Println("false")
		return
	}

	split := len(inputString) / 2

	splitText := inputString[:split]

	splitText += splitText

	if splitText == inputString {
		fmt.Println("true")
		return
	} else {
		fmt.Println("false")
		return
	}

}
