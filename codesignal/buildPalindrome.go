package main

import (
	"fmt"
	"strings"
)

func main() {
	//reverse string then build a str to check to see if it contains with the orig str when it does not not the pos add to another str that will be the palindrome

	st := "abaa"
	rev := ""
	buildStr := ""

	drome := ""

	for i := len(st) - 1; i >= 0; i-- {
		rev += string(st[i])
	}

	fmt.Println(rev)

	for _, x := range rev {
		buildStr += string(x)
		if strings.Contains(st, buildStr) {
			continue
		} else {
			drome += string(x)
		}

	}

	fmt.Println(st + drome)

}
