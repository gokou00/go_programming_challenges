package main

import (
	"fmt"
	"strings"
)

func main() {
	startTag := "<li class='test'>"

	//endTagMap := make(map[string]string)

	idx := 0

	endingTag := "</"

	if !strings.Contains(startTag, " ") {

		for i := 1; i < len(startTag); i++ {
			endingTag += string(startTag[i])
		}
		fmt.Println(endingTag)
		return
	}

	idx = strings.Index(startTag, " ")

	for i := 1; i < idx; i++ {
		endingTag += string(startTag[i])
	}

	endingTag += ">"
	fmt.Println(endingTag)
	return

}
