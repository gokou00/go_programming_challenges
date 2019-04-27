package main

import (
	"fmt"
	//"sort"
	"strings"
)

func main() {

	text := "You are the best!!!!!!!!!!!! CodeFighter ever!"
	count := 0
	temp := ""
	largestWord := ""

	textArray := strings.Split(text, " ")

	//var wordLength []int

	//fmt.Println(textArray)

	for _, x := range textArray {
		for _, y := range x {
			if string(y) >= "a" && string(y) <= "z" || string(y) >= "A" && string(y) <= "Z" {
				temp += string(y)

			}
		}

		if count == 0 {
			largestWord = temp
			temp = ""
			count++
			continue

		}

		if len(temp) > len(largestWord) {
			largestWord = temp
		}

		//wordLength = append(wordLength, len(temp))
		temp = ""

	}

	fmt.Println(largestWord)

}
