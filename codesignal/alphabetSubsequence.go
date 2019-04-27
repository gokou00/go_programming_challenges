package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s := "fz"

	letterMap := make(map[string]int)

	for _, x := range s {
		letterMap[string(x)]++

		if letterMap[string(x)] > 1 {
			fmt.Println("false")
			return
		}

	}

	strArr := strings.Split(s, "")
	//strArrCopy := strArr[:]

	strArrCopy := make([]string, len(strArr))

	copy(strArrCopy, strArr)

	sort.Strings(strArrCopy)

	//fmt.Println(strArr)
	//fmt.Println(strArrCopy)

	for i := 0; i < len(strArr); i++ {
		if strArr[i] != strArrCopy[i] {
			fmt.Println("false")
			return
		}

	}

	fmt.Println("true")
	return

}
