package main

import "fmt"

func main() {
	str := "abban"

	revStr := ""

	for i := len(str) - 1; i >= 0; i-- {
		revStr += string(str[i])
	}

	fmt.Println(str == revStr)

}
