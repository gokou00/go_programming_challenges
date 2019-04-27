package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 99
	numStr := strconv.Itoa(n)
	buildStr := ""
	increase := false
	finalStr := ""

	for i := len(numStr) - 1; i > 0; i-- {
		temp, _ := strconv.Atoi(string(numStr[i]))

		if increase {
			temp++
			increase = false
		}

		if temp >= 5 {
			increase = true
		}

		//tempStr := strconv.Itoa(temp)

		buildStr += "0"

	}

	if increase {
		temp, _ := strconv.Atoi(string(numStr[0]))
		temp++
		tempStr := strconv.Itoa(temp)
		finalStr = tempStr + buildStr

	} else {
		temp, _ := strconv.Atoi(string(numStr[0]))
		tempStr := strconv.Itoa(temp)
		finalStr = tempStr + buildStr

	}

	//fmt.Println(finalStr)
	finalAns, _ := strconv.Atoi(finalStr)

	fmt.Println(finalAns)

}
