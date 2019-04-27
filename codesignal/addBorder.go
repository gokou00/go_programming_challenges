package main

import (
	"fmt"
	"strings"
)

func main() {
	picture := []string{"abc", "ded"}

	borderLength := len(picture[0])
	borderLength += 2

	border := strings.Repeat("*", borderLength)

	var finalArr []string

	finalArr = append(finalArr, border)

	for _, x := range picture {
		temp := "*"
		temp += x
		temp += "*"
		finalArr = append(finalArr, temp)
		temp = ""

	}
	finalArr = append(finalArr, border)

	fmt.Println(finalArr)

}
