package main

import "fmt"

func Solution(str string) []string {
	var strSlice []string
	letters := ""
	for i, x := range str {
		if len(letters) <= 2 {
			letters += string(x)
			fmt.Println(letters)
		}

		if len(letters) == 2 {
			strSlice = append(strSlice, letters)
			//fmt.Println(letters)
			letters = ""
		}

		if len(letters)%2 != 0 && i == len(str)-1 {
			letters += "_"
			strSlice = append(strSlice, letters)
		}
	}

	return strSlice

}

func main() {

	var finalAns []string
	finalAns = Solution("h")
	fmt.Println(finalAns)

}
