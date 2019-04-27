package main

import "fmt"

func main() {
	maxLength := 5
	text := "Now and then, however, he is horribly thoughtless, and seems to take a real delight in giving me pain."
	strBuild := ""

	var strArr []string
	var strArr2 []string

	for _, x := range text {
		if string(x) == " " {
			strArr = append(strArr, strBuild)
			strBuild = ""
		} else if (string(x) >= "a" && string(x) <= "z") || (string(x) >= "A" && string(x) <= "Z") {
			strBuild += string(x)
		}

	}

	if len(strBuild) != 0 {
		strArr = append(strArr, strBuild)
	}

	fmt.Println(strArr)

	for _, x := range strArr {
		if len(x) <= maxLength {
			strArr2 = append(strArr2, x)
		}
	}

	fmt.Println(len(strArr2))

}
