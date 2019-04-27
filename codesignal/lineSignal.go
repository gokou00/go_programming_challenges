package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "fifteenletterrs"

	newStr := ""
	count := 0
	char := ""

	for i, x := range s {
		if i == 0 {
			count++
			char = string(x)
			continue
		}

		if char == string(x) {
			count++
			//	fmt.Println(char, string(x))
			//	fmt.Println(count)

		} else {
			//count++
			countStr := strconv.Itoa(count)
			newStr += countStr
			newStr += strings.Repeat(char, count)

			char = string(x)
			count = 1

			//	fmt.Println(char)
		}

	}

	countStr := strconv.Itoa(count)
	newStr += countStr
	newStr += strings.Repeat(char, count)

	//	fmt.Println(char)
	//	fmt.Println(count)

	fmt.Println(newStr)

	finalStr := ""
	//	temp := 0

	buildNum := ""

	for i := 0; i < len(newStr); i++ {

		_, err := strconv.Atoi(string(newStr[i]))

		if err == nil {
			buildNum += string(newStr[i])
		}

		if err != nil {
			if buildNum == "" {
				continue
			}

			tempNum, _ := strconv.Atoi(buildNum)

			if tempNum == 1 {
				finalStr += string(newStr[i])
				buildNum = ""
			} else if tempNum > 1 {
				strNum := strconv.Itoa(tempNum)
				finalStr += strNum
				finalStr += string(newStr[i])
				buildNum = ""
			}

		}

	}

	fmt.Println(finalStr)

}
