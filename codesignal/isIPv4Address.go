package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	test := ".16.254.1"

	ipArry := strings.Split(test, ".")

	fmt.Println(ipArry[0])
	fmt.Println(len(ipArry))

	num, err := strconv.Atoi(ipArry[0])

	if err != nil {
		fmt.Println("error")
	}

	//fmt.Println(num)

}
