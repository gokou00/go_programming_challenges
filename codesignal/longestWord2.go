package main

import (
	"fmt"
	//"sort"
	//	"strings"
)

func main() {

	text := "ABCd"
	//count := 0
	temp := ""
	largestWord := ""
	var wordArray []string

	//textArray := strings.Split(text, " ")

	for _, x := range text {
		if string(x) >= "a" && string(x) <= "z" || string(x) >= "A" && string(x) <= "Z" {
			temp += string(x)

		} else {
			wordArray = append(wordArray, temp)
			temp = ""
		}
	}

	if temp != "" {
		wordArray = append(wordArray, temp)
	}

	//fmt.Println(len(wordArray))

	if len(wordArray) == 1 {
		fmt.Println(wordArray[0])
		return
	}

	//var wordLength []int

	//fmt.Println(textArray)

	largestWord = wordArray[0]

	for i := 1; i < len(wordArray); i++ {

		if len(wordArray[i]) > len(largestWord) {
			largestWord = wordArray[i]
		}

	}

	fmt.Println(largestWord)

	/*

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
	*/
}
