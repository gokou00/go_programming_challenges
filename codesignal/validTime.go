package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	time := "09:56"

	timeArray := strings.Split(time, ":")

	fmt.Println(timeArray)

	isValid := true

	temp, _ := strconv.Atoi(timeArray[0])

	if temp < 0 || temp > 23 {
		fmt.Println("false")
		return
	}

	temp2, _ := strconv.Atoi(timeArray[1])

	if temp2 < 0 || temp2 > 59 {
		fmt.Println("false")
		return
	}

	fmt.Println(isValid)

}
