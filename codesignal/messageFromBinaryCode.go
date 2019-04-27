package main

import (
	"fmt"
	"strconv"
)

func main() {
	code := "010010000110010101101100011011000110111100100001"

	strbuild := ""

	var binarray []string
	var ascii []int64

	for i := 0; i < len(code); i++ {
		if i%8 == 0 && i != 0 {
			binarray = append(binarray, strbuild)
			strbuild = ""

		}

		strbuild += string(code[i])

	}

	fmt.Println(strbuild)

	binarray = append(binarray, strbuild)

	fmt.Println(binarray)

	for _, x := range binarray {
		i, _ := strconv.ParseInt(x, 2, 64)
		ascii = append(ascii, i)
	}

	fmt.Println(ascii)
	finalStr := ""

	for _, x := range ascii {
		finalStr += string(x)
	}

	fmt.Println(finalStr)

}
