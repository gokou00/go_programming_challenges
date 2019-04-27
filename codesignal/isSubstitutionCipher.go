package main

import (
	"fmt"
	"strconv"
)

func main() {

	string1 := "aaabbbccc"
	string2 := "aaabbbccc"

	string1Map := make(map[string]int)
	string2Map := make(map[string]int)

	string1Build := ""
	string2Build := ""

	count := 1

	for _, x := range string1 {
		_, ok := string1Map[string(x)]
		if ok {
			continue
		} else {
			string1Map[string(x)] = count
			count++
		}

	}

	for _, x := range string1 {

		temp := strconv.Itoa(string1Map[string(x)])

		string1Build += temp

	}

	//fmt.Println(string1Map)
	//fmt.Println(string1Build)

	count = 1

	for _, x := range string2 {
		_, ok := string2Map[string(x)]
		if ok {
			continue
		} else {
			string2Map[string(x)] = count
			count++
		}

	}

	for _, x := range string2 {

		temp := strconv.Itoa(string2Map[string(x)])

		string2Build += temp

	}

	//fmt.Println(string2Map)
	//fmt.Println(string2Build)

	if string1Build == string2Build {
		fmt.Println("true")
		return
	} else {
		fmt.Println("false")
		return
	}

}
