package main

import (
	"fmt"
	"strconv"
)

func main() {
	red := uint64(0)
	red = uint64(red)

	green := uint64(0)
	green = uint64(green)

	blue := uint64(0)
	blue = uint64(blue)

	var finalArr []int

	//colors := []string{"A0FC77", "90CACA"}

	for i, x := range colors {

		red, _ := strconv.ParseUint(x[0:2], 16, 32)

		green, _ := strconv.ParseUint(x[2:4], 16, 32)

		blue, _ := strconv.ParseUint(x[4:], 16, 32)

		if green > red && green > blue {
			finalArr = append(finalArr, i)

		}

		//fmt.Println(x[4:7])

	}

	fmt.Println(finalArr)

}
