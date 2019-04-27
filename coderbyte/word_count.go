package main

import (
	"fmt"
)

func main() {
	str := "all mmm sds sd\nsdsd"

	var strArray []string

	strBuild := ""

	for _, x := range str {
		if string(x) >= "a" && string(x) <= "z" || string(x) >= "A" && string(x) <= "Z" {
			strBuild += string(x)
		} else {
			strArray = append(strArray, strBuild)
			strBuild = ""
		}
	}

	if strBuild != "" {
		strArray = append(strArray, strBuild)
	}

	fmt.Println(strBuild)
	fmt.Println(strArray)
	fmt.Println(len(strArray))

}
