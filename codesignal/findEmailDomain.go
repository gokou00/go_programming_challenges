package main

import (
	"fmt"
	"strings"
)

func main() {
	//use the strings method to find the last occ of @
	address := "\" \"@xample.org"

	finalStr := ""

	index := strings.LastIndex(address, "@")

	index++

	for i := index; i < len(address); i++ {
		finalStr += string(address[i])

	}

	fmt.Println(finalStr)

}
