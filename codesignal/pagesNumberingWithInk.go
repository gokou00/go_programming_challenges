package main

import (
	"fmt"
	"strconv"
)

func main() {

	current := 80
	numberOfDigits := 1000
	strBuild := ""
	i := 0

	for i = current; len(strBuild) < numberOfDigits; i++ {
		temp := strconv.Itoa(i)

		if len(strBuild+temp) > numberOfDigits {
			break
		}
		strBuild += temp
		//fmt.Println(strBuild)
	}

	fmt.Println(i - 1)

}
