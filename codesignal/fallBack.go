package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	time := "12:23PM"

	end := strings.Index(time, ":")
	typ := time[end+3 : len(time)]
	fmt.Println(typ)

	change := string(time[0:end])
	fmt.Println(change)

	num, _ := strconv.Atoi(change)
	changeType := false

	if num == 1 {
		num = 12
	} else if num == 12 {
		changeType = true
		num--
	} else {
		num--
	}

	finalStr := ""
	numStr := strconv.Itoa(num)

	finalStr += numStr

	if changeType == false {

		for i := end; i < len(time); i++ {
			finalStr += string(time[i])
		}

	} else {
		for i := end; i < len(time)-2; i++ {
			finalStr += string(time[i])
		}
		if typ == "PM" {
			finalStr += "AM"
		} else {
			finalStr += "PM"
		}

	}

	fmt.Println(finalStr)
}
