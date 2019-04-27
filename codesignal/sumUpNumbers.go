package main

import (
	"fmt"
	"strconv"
)

func main() {

	inputString := "there are some (12) digits 5566 in this 770 string 239"
	var numStrArray []string
	total := 0
	strNum := ""

	for _, x := range inputString {
		_, err := strconv.Atoi(string(x))

		if err != nil && strNum != "" {
			numStrArray = append(numStrArray, strNum)
			num, _ := strconv.Atoi(strNum)
			total += num
			strNum = ""
			continue
		} else if err != nil {
			continue
		} else {
			strNum += string(x)
		}

	}

	if strNum != "" {
		num, _ := strconv.Atoi(strNum)
		total += num
	}

	fmt.Println(strNum)

	fmt.Println(numStrArray)
	fmt.Println(total)

}
