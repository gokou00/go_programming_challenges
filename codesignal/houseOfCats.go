package main

import (
	"fmt"
	//"sort"
)

func main() {

	legs := 44
	people := 0

	var finalArr []int

	for legs >= 0 {
		cats := 0
		if legs%4 == 2 || legs%4 == 0 {
			if legs%4 == 2 {
				cats = 1
			} else {
				cats = 0
			}
		}
		finalArr = append(finalArr, cats+people)
		people += 2
		legs -= 4

	}

	fmt.Println(finalArr)

}
