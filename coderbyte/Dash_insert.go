package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "56730"
	strbuild := ""

	for i := 0; i < len(str)-1; i++ {
		temp1, _ := strconv.Atoi(string(str[i]))
		temp2, _ := strconv.Atoi(string((str[i+1])))
		if temp1%2 == 1 && temp2%2 == 1 {
			strbuild += string(str[i]) + "-"
		} else {
			strbuild += string(str[i])
		}
	}

	strbuild += string(str[len(str)-1])
	fmt.Println(strbuild)

}
