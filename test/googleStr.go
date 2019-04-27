package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s1 := "abcdedadf"
	s2 := "cae"

	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)

	firstPart := ""
	lastPart := ""

	for _, x := range s1 {
		if strings.Contains(s2, string(x)) {
			firstPart += string(x)
		} else {
			lastPart += string(x)
		}
	}

	testMap1 := make(map[string]int)
	testMap2 := make(map[int]string)

	for i, x := range s2 {
		testMap1[string(x)] = i
		testMap2[i] = string(x)
	}

	var placementArr []int

	for _, x := range firstPart {
		placementArr = append(placementArr, testMap1[string(x)])
	}

	sort.Ints(placementArr)
	sortedFirst := ""

	for _, x := range placementArr {
		sortedFirst += testMap2[x]
	}

	fmt.Println(sortedFirst + lastPart)

}
