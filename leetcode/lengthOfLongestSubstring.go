package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	//s := "pwwkew"
	//s := "abcabcbb"
	//s := "pwwkew"
	s := ""

	strBuild := ""

	var arrSize []int
	count := 0

	for i := 0; i < len(s); i++ {

		if strings.Contains(strBuild, string(s[i])) {
			if i == len(s)-1 {
				break
			}

			if string(s[i-1]) == string(s[i]) {
				arrSize = append(arrSize, len(strBuild))
				strBuild = ""
			} else {
				i = count
				count++
				arrSize = append(arrSize, len(strBuild))
				strBuild = ""
				continue

			}

			//strBuild += s[i]

		}

		strBuild += string(s[i])

		fmt.Println(strBuild)

	}

	//fmt.Println(len(strBuild))

	if len(strBuild) != 0 {
		arrSize = append(arrSize, len(strBuild))
	}

	sort.Ints(arrSize)

	fmt.Println(arrSize)

	if len(arrSize) == 0 {
		fmt.Println(0)
	} else {
		fmt.Println(arrSize[len(arrSize)-1])
	}

	//fmt.Println(arrSize[len(arrSize)-1])

}
