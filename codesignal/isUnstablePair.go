package main

import (
	"fmt"
	//"sort"
	"strings"
)

func main() {
	filename1 := "aaZ"
	filename2 := "AAzz"

	//filename1 = strings.ToUpper(filename1)
	//fmt.Println(strings.EqualFold(filename1, filename2))

	if strings.Compare(filename1, filename2) > 0 {
		fmt.Println(strings.Compare(strings.ToUpper(filename1), strings.ToUpper(filename2)) < 0)
		return
	} else if strings.Compare(filename1, filename2) < 0 {
		fmt.Println(strings.Compare(strings.ToLower(filename1), strings.ToLower(filename2)) > 0)
		return
	} else {
		fmt.Println("false")
		return
	}

}
