package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "JhBkPBaozMnBqEWiIaOEje"
	strBuild := string(s[0])

	for _, x := range s[1:] {

		if x >= 65 && x <= 90 {
			strBuild += " "
			//strBuild += string(x)

		}

		strBuild += string(x)

	}

	finalStr := strings.ToLower(strBuild)

	fmt.Println(strBuild)
	fmt.Println(finalStr)

}
