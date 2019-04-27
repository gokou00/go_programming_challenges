package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "abdc"
	greatest := ""

	pairs := [][]int{{1, 4}, {3, 4}}

	for _, x := range pairs {
		s1 := x[0] - 1
		s2 := x[1] - 1

		tempArr := strings.Split(str, "")

		tempStr := tempArr[s1]

		tempArr[s1] = tempArr[s2]
		tempArr[s2] = tempStr

		combinStr := strings.Join(tempArr, "")

		if combinStr > greatest {
			greatest = combinStr
		}

		//fmt.Println(x)
	}

	fmt.Println(greatest)

}
