package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	cipher := "11211111911310110810910097107108115111112119113101106107971101021101061021041149710511411497"
	finalStr := ""

	inc := 0
	num := 0

	if len(cipher) >= 3 {
		inc = 3

		temp := cipher[0:3]
		num1, _ := strconv.Atoi(temp)
		num = num1
		//finalStr += string(num1)

	} else {
		//inc = 2
		temp := cipher[0:2]
		num1, _ := strconv.Atoi(temp)

		finalStr += string(num1)
		fmt.Println(finalStr)
		return

	}

	for cipher != "" {
		if (num >= 65 && num <= 90) || (num >= 97 && num <= 122) {
			inc = 3
			finalStr += string(num)
		} else {
			test := cipher[0:2]
			num1, _ := strconv.Atoi(test)
			finalStr += string(num1)
			inc = 2

		}

		arr := strings.Split(cipher, "")
		arr = arr[inc:len(arr)]
		cipher = strings.Join(arr, "")
		tempStr := ""
		if len(cipher) >= 3 {
			tempStr = cipher[0:3]
			num1, _ := strconv.Atoi(tempStr)
			num = num1
		} else if len(cipher) == 2 {
			tempStr = cipher[0:2]
			num1, _ := strconv.Atoi(tempStr)
			//num = num1
			finalStr += string(num1)
			fmt.Println(finalStr)
			return
		}

	}

	fmt.Println(finalStr)

}
