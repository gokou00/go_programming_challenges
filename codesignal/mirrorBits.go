package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 5
	temp := strconv.FormatInt(int64(a), 2)

	revStr := ""

	for i := len(temp) - 1; i >= 0; i-- {
		revStr += string(temp[i])
	}

	fmt.Println(revStr)

	finalAns, _ := strconv.ParseInt(revStr, 2, 64)

	fmt.Println(int(finalAns))

}
