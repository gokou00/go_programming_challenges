package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "88Hello 3World!"
	strBuild := ""
	total := 0

	for _, x := range str {
		_, err := strconv.Atoi(string(x))
		if err != nil {
			num, _ := strconv.Atoi(strBuild)
			total += num
			strBuild = ""
			continue
		} else {
			strBuild += string(x)
		}

	}

	if strBuild != "" {
		num, _ := strconv.Atoi(strBuild)
		total += num

	}

	fmt.Println(total)

}
