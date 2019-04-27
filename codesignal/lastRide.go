package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 8
	//div by 60 to get the hours
	//use mod to get the mins
	numStr := ""
	hours := 0

	hours = n / 60

	mins := 0

	mins = n % 60

	temp := strconv.Itoa(hours)
	temp2 := strconv.Itoa(mins)

	numStr = temp + temp2

	num, _ := strconv.Atoi(numStr)

	fmt.Println(addnum(num))

}

func addnum(num int) int {
	pop := 0
	total := 0

	for num != 0 {
		pop = num % 10
		total += pop
		num /= 10
	}

	return total

}
