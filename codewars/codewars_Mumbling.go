package main

import (
	"fmt"
	"strings"
)

func main() {
	input := strings.ToLower("RqaEzty")
	final := ""
	letter := ""

	for i := 0; i < len(input); i++ {
		letter = string([]rune(input)[i])
		//fmt.Println(final)
		letter = strings.ToUpper(letter)
		final += letter

		if i > 0 {
			final += strings.Repeat(string([]rune(input)[i]), i)
			final += "-"
		} else if i == 0 {
			final += "-"
		}
	}

	//fmt.Println(len(final))
	final = final[0 : len(final)-1]

	fmt.Println(final)
	fmt.Println(len(final))
}
