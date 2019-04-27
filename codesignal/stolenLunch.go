package main

import "fmt"

func main() {
	note := "just 63jd73 some random note jkhdf83 ds823 that you, dfj238 dsf38 little one?, will abjk38 s83    skdu3 29never get!"
	alpha := "abcdefghij"
	num := "0123456789"
	finalStr := ""

	alphaMap := make(map[string]string)
	numMap := make(map[string]string)

	for i := 0; i < len(alpha); i++ {
		alphaMap[string(alpha[i])] = string(num[i])
		numMap[string(num[i])] = string(alpha[i])

	}

	for _, x := range note {
		if string(x) >= "0" && string(x) <= "9" {
			finalStr += numMap[string(x)]
		} else if string(x) >= "a" && string(x) <= "j" {
			finalStr += alphaMap[string(x)]
		} else {
			finalStr += string(x)
		}

	}

	fmt.Println(finalStr)

}
