package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	pattern := "abba"
	str := "dog dog dog dog"
	patternStr := ""
	strBuild := ""
	count := 1

	strArr := strings.Split(str, " ")

	patternMap := make(map[string]int)
	strMap := make(map[string]int)

	for _, x := range pattern {
		_, ok := patternMap[string(x)]

		if !ok {
			patternMap[string(x)] = count
			count++
		}

	}

	for _, x := range pattern {
		temp := strconv.Itoa(patternMap[string(x)])

		patternStr += temp

	}

	count = 1

	for _, x := range strArr {
		_, ok := strMap[string(x)]

		if !ok {
			strMap[string(x)] = count
			count++
		}

	}

	for _, x := range strArr {
		//_, ok := patternMap[string(x)]
		temp := strconv.Itoa(strMap[string(x)])
		strBuild += temp

	}

	fmt.Println(patternStr)
	fmt.Println(strBuild)

}
