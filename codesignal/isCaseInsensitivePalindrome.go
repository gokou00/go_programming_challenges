package main

import (
	"fmt"
	"strings"
)

func main() {
	inputString := "AaBaa"

	inputString = strings.ToLower(inputString)
	strBuild := ""

	for i := len(inputString) - 1; i >= 0; i-- {
		strBuild += string(inputString[i])
	}

	fmt.Println(strBuild)

}
