package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	test := []string{"push 10", "min", "push 5", "min", "push 8", "min", "pop", "min", "pop", "min"}

	//var numStrArray []string
	instruction := ""
	numStr := ""
	var stack []int
	var final []int

	for _, x := range test {
		for _, y := range x {
			_, err := strconv.Atoi(string(y))

			if err == nil {
				numStr += string(y)
			} else if err != nil {
				instruction += string(y)
			}

		}

		//	fmt.Println(instruction)
		//	fmt.Println(numStr)

		switch instruction {

		case "push ":
			fmt.Println("test")
			temp, _ := strconv.Atoi(numStr)

			stack = append(stack, temp)

		case "min":
			tempArray := make([]int, len(stack))
			copy(tempArray, stack)
			sort.Ints(tempArray)
			final = append(final, tempArray[0])

		case "pop":
			stack = stack[0 : len(stack)-1]
			fmt.Println(stack)

		}

		numStr = ""
		instruction = ""

	}

	fmt.Println(final)
	//fmt.Println(stack)

}
