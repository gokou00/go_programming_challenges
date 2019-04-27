package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	ver1 := "5"
	ver2 := "1"

	if !strings.Contains(ver1, ".") && !strings.Contains(ver2, ".") {
		num1, _ := strconv.Atoi(ver1)
		num2, _ := strconv.Atoi(ver2)

		if num1 > num2 {
			fmt.Println("true")
			return
		}
	}

	ver1Arr := strings.Split(ver1, ".")
	ver2Arr := strings.Split(ver2, ".")

	for i := 0; i < len(ver1Arr); i++ {
		num1, _ := strconv.Atoi(ver1Arr[i])
		num2, _ := strconv.Atoi(ver2Arr[i])
		if num1 == num2 {
			continue
		} else if num1 > num2 {
			fmt.Println("true")
			return
		} else if num1 < num2 {
			fmt.Println("false")
			return
		}

	}

	ver1Str := ""
	ver2Str := ""

	for _, x := range ver1 {
		if string(x) == "." {
			continue
		}
		ver1Str += string(x)
	}

	for _, x := range ver2 {
		if string(x) == "." {
			continue
		}
		ver2Str += string(x)
	}

	num1, _ := strconv.Atoi(ver1Str)
	num2, _ := strconv.Atoi(ver2Str)

	if num1 > num2 {
		fmt.Println("true")
		return
	}

	fmt.Println("false")
	return
}
