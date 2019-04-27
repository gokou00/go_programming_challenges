package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello World"
	strBuild := ""

	for _, x := range str {
		if string(x) >= "a" && string(x) <= "z" {
			strBuild += strings.ToUpper(string(x))
		} else if string(x) >= "A" && string(x) <= "Z" {
			strBuild += strings.ToLower(string(x))
		} else {
			strBuild += string(x)
		}
	}

	fmt.Println(strBuild)

}
