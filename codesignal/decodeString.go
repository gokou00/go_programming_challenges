package main

//create a function that takes in the string and return the repeated char minus the count and brackets

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	test := "100[codefight]"

	for strings.Contains(test, "[") {

		test = decodeString2(test)
		fmt.Println(test)

	}

	fmt.Println(test)
}

func decodeString2(s string) string {

	lhs := ""
	rhs := ""
	indexLeft := 0
	indexNext := 0
	repeatStr := ""
	//	repeatNum := 0
	finalStr := ""
	num := 0
	numStr := ""
	finalNum := 0

	indexLeft = strings.LastIndex(s, "[")
	indexNum := indexLeft - 1

	//	repeatNum, _ = strconv.Atoi(string(s[indexLeft-1]))

	for num = indexNum; num > 0; num-- {
		_, err := strconv.Atoi(string(s[num]))
		fmt.Println("Starting loop", num)

		if err != nil {
			fmt.Println("print num")
			num = num + 1
			break
		} else if num == 1 {
			_, errTest := strconv.Atoi(string(s[num-1]))
			if errTest != nil {
				break
			}

		} else {
			continue
		}
	}

	fmt.Println(num)

	for i := num; i <= indexNum; i++ {
		numStr += string(s[i])
	}

	fmt.Println("num String", numStr)

	finalNum, _ = strconv.Atoi(numStr)

	fmt.Println(finalNum)

	//newStr := ""

	if num <= 0 {
		fmt.Println("no LHS")
	} else {
		for i := 0; i < num; i++ {
			lhs += string(s[i])
		}
	}

	for i := indexLeft; i < len(s); i++ {
		rhs += string(s[i])
	}

	fmt.Println(rhs)

	indexNext = strings.Index(rhs, "]")

	//	fmt.Println(indexLeft)
	//	fmt.Println(indexNext)

	for i := 1; i < indexNext; i++ {
		repeatStr += string(rhs[i])
		//fmt.Println(repeatStr)
	}

	fmt.Println(repeatStr)

	repeatStr = strings.Repeat(repeatStr, finalNum)

	finalStr = lhs + repeatStr

	if indexNext < len(rhs) {

		for i := indexNext + 1; i < len(rhs); i++ {
			finalStr += string(rhs[i])
		}

	}

	return finalStr

}
