package main

import (
	"fmt"
	"strings"
)

func main() {

	vowels := "aeiou"

	str := "ALL cow eat grass and moo"

	newStr := strings.ToLower(str)
	count := 0

	for _, x := range newStr {
		if strings.Contains(vowels, string(x)) {
			count++
		}
	}

	fmt.Println(count)

}
