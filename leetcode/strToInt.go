package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	//https://leetcode.com/problems/string-to-integer-atoi/description/
	str := "+91283472443342"

	str = strings.Trim(str, " ")

	test, _ := strconv.Atoi(str)

	fmt.Println(test)

}
