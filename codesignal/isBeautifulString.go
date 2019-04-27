package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	inputString := "zyy"
	inputStringMap := make(map[string]int)
	orderedInputMap := make(map[string]int)
	var orderedInputArray []string

	alphabet := "abcdefghijklmnopqrstuvwxyz"

	for _, x := range alphabet {
		inputStringMap[string(x)] = 0
	}

	for _, x := range inputString {

		//fmt.Println(key)

		inputStringMap[string(x)]++
	}

	for _, x := range inputString {

		//fmt.Println(key)

		orderedInputMap[string(x)] = 0
	}

	for key, _ := range orderedInputMap {

		orderedInputArray = append(orderedInputArray, key)
	}

	sort.Strings(orderedInputArray)

	//	sort.Strings()

	fmt.Println(inputStringMap)
	fmt.Println(orderedInputArray)

	positionLast := strings.Index(alphabet, orderedInputArray[len(orderedInputArray)-1])

	fmt.Println(positionLast)
	orderedInputJoin := strings.Join(orderedInputArray, "")

	for i := positionLast; i >= 0; i-- {
		if strings.Contains(orderedInputJoin, string(alphabet[i])) {
			continue
		} else {
			orderedInputArray = append(orderedInputArray, string(alphabet[i]))
		}
	}

	sort.Strings(orderedInputArray)

	fmt.Println(orderedInputArray)

	for i := len(orderedInputArray) - 1; i > 0; i-- {
		//fmt.Println(orderedInputMap["a"], " a ")
		//fmt.Println(orderedInputMap["b"], " b")
		if inputStringMap[orderedInputArray[i]] > inputStringMap[orderedInputArray[i-1]] {
			fmt.Println("false")
			return
		}

	}

	fmt.Println("true")

}
