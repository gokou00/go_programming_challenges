package main

import "fmt"

func main() {
	inputString := "codesignal"

	stringMap := make(map[string]string)

	str := "abcdefghijklmnopqrstuvwxyz"
	revstr := ""

	for i := len(str) - 1; i >= 0; i-- {
		revstr += string(str[i])
	}

	for i := 0; i < len(str); i++ {
		stringMap[string(str[i])] = string(revstr[i])

	}

	finalStr := ""

	for _, x := range inputString {
		finalStr += stringMap[string(x)]

	}

	fmt.Println(finalStr)

}
