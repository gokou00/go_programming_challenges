package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	a := []int{187, 99, 42, 43}
	var revArry []string

	for i := len(a) - 1; i >= 0; i-- {
		testTemp := strconv.FormatInt(int64(a[i]), 2)
		StrTemp := ""

		if len(testTemp) < 8 {
			StrTemp = AddZeros(testTemp)
		} else {
			StrTemp = testTemp
		}

		revArry = append(revArry, StrTemp)

	}

	temp := strings.Join(revArry, "")
	fmt.Println(revArry)
	//fmt.Println(temp)

	num, _ := strconv.ParseInt(temp, 2, 64) //("1001", 2)

	fmt.Println(int(num))

}

func AddZeros(str string) string {

	tempStr := ""

	length := 8 - len(str)

	for i := 0; i < length; i++ {
		tempStr += "0"
	}

	return tempStr + str

}
