package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "abcabcabc"
	b := "aaaaaaaaaaabbbbbbbbbbcccccccccc"

	str1Map := make(map[string]int)

	for _, x := range a {

		count := strings.Count(b, string(x))
		str1Map[string(x)] = count
	}

	if len(str1Map) == 1 {
		fmt.Println(str1Map[string(a[0])] / len(a))
		return

	}

	fmt.Println(str1Map)

	least := 999999999993

	for _, x := range str1Map {

		if x < least {
			least = x
		}

	}

	if least < len(a) {
		fmt.Println(least)
		return
	} else {
		fmt.Println(least / len(str1Map))
	}

	//fmt.Println(least)

}
