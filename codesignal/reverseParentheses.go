package main

import (
	"fmt"
	"strings"
)

func main() {

	// check for ()

	driver := "a(bcdefghijkl(mno)p)q"
	str := driver

	if strings.Contains(str, "(") == false || strings.Contains(str, ")") == false {
		fmt.Println(str)
	}

	for i := 0; i < len(str); i++ {
		if strings.Contains(str, "(") || strings.Contains(str, ")") {
			str = rev(str)
		}
	}

	fmt.Println(str)

}

func rev(s string) string {
	retStr := ""

	firstStr := ""

	restStr := ""

	revStr := ""

	indexFirst := strings.LastIndex(s, "(")

	for i := 0; i < indexFirst; i++ {
		firstStr += string(s[i])
	}

	fmt.Println(firstStr)

	for i := indexFirst; i < len(s); i++ {
		restStr += string(s[i])
	}

	fmt.Println(restStr)

	indexLast := strings.Index(restStr, ")")

	//	fmt.Println(indexLast)

	for i := indexLast - 1; i > 0; i-- {

		revStr += string(restStr[i])
	}

	fmt.Println(revStr)

	retStr = firstStr + revStr

	//fmt.Println(retStr)

	for i := indexLast + 1; i < len(restStr); i++ {
		retStr += string(restStr[i])
	}

	return retStr

}
