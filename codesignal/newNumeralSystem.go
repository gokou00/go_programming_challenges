package main

import (
	"fmt"
	"sort"
)

func main() {
	number := "G"
	alphaStr := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphaMap := make(map[string]int)
	numberMap := make(map[int]string)
	testMap := make(map[string]int)
	var finalAns []string
	var testArr []string

	for i, x := range alphaStr {

		alphaMap[string(x)] = i
	}

	for key, val := range alphaMap {
		numberMap[val] = key
	}

	target := alphaMap[number]

	//orderMap := make(map[string]int)

	for _, x := range alphaStr {
		key1 := string(x)
		val1 := alphaMap[string(x)]

		find := target - val1

		_, ok := numberMap[find]

		if ok {
			testArr = append(testArr, key1)
			testArr = append(testArr, numberMap[find])

			sort.Strings(testArr)

			strBuild := testArr[0]
			strBuild += " + "
			strBuild += testArr[1]
			finalAns = append(finalAns, strBuild)
			testMap[strBuild] = 0
			strBuild = ""
			testArr = nil
		}

	}

	fmt.Println(finalAns)
	fmt.Println(testMap)

}
