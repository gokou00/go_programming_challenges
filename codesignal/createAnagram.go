package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s := "OVGHK"
	t := "RPGUC"

	sMap := make(map[string]int)
	total := 0
	strS := ""
	strT := ""

	for _, x := range s {
		sMap[string(x)] = 0

	}

	for key, _ := range sMap {

		if !(strings.Contains(t, key)) {
			total += strings.Count(s, key)
		} else {

			numS := strings.Count(s, key)
			numT := strings.Count(t, key)

			strS += strings.Repeat(key, numS)
			strT += strings.Repeat(key, numT)
		}

	}

	fmt.Println(total)
	//fmt.Println(strS)
	//fmt.Println(strT)

	sArr := strings.Split(strS, "")
	tArr := strings.Split(strT, "")

	sort.Strings(sArr)
	sort.Strings(tArr)

	fmt.Println(sArr)
	fmt.Println(tArr)

	for i := 0; i < len(sArr); i++ {
		if string(sArr[i]) != string(tArr[i]) {
			total++
		}

	}

	fmt.Println(total)

}
