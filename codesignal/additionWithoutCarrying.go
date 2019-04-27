package main

import (
	"fmt"
	"strconv"
)

func main() {
	param1 := 54321
	param2 := 54321

	if param1 == 0 && param2 == 0 {
		fmt.Println(0)
		return
	}

	var finalans []int

	for param1 != 0 || param2 != 0 {
		pop1 := param1 % 10
		pop2 := param2 % 10
		param1 /= 10
		param2 /= 10

		ans := (pop1 + pop2) % 10

		finalans = append(finalans, ans)

	}

	fmt.Println(finalans)

	strbuilder := ""

	for i := len(finalans) - 1; i >= 0; i-- {
		temp := strconv.Itoa(finalans[i])
		strbuilder += temp
	}

	ans2, _ := strconv.Atoi(strbuilder)

	fmt.Println(ans2)

}
