package main

import "fmt"

func main() {
	mask := []string{"     ",
		" @ @ ",
		"  <  ",
		" O   ",
		"     "}
	strRev := ""
	count := 0
	finalcount := 0

	for _, x := range mask {
		strRev = revStr(x)
		for _, j := range x {
			if string(j) != string(strRev[count]) {
				finalcount++
			}

			count++
		}
		count = 0

	}

	fmt.Println(finalcount / 2)

}

func revStr(str string) string {

	temp := ""

	for i := len(str) - 1; i >= 0; i-- {
		temp += string(str[i])
	}

	return temp

}
