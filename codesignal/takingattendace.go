package main

import (
	"fmt"
	"strings"
)

func main() {
	//	const min = 5
	classList := []string{"Chruschtschov",
		"Hristo",
		"Nguyen",
		"Dmitry",
		"Madchen",
		"Fujiyama",
		"Connor"}

	total := 5
	count := 0
	final := 0

	isConstant := false

	for _, x := range classList {
		x = strings.ToLower(x)
		//	fmt.Println(isConstant, "what is the con")
		for _, y := range x {
			//	fmt.Println(string(y))
			if string(y) != "a" && string(y) != "e" && string(y) != "i" && string(y) != "o" && string(y) != "u" && string(y) != "y" && isConstant == true {
				count++
				fmt.Println(string(y), "isConstant is already set to true")
				//	fmt.Println(count, "isConstant")
			} else if string(y) != "a" && string(y) != "e" && string(y) != "i" && string(y) != "o" && string(y) != "u" && string(y) != "y" {
				count++
				isConstant = true
				fmt.Println(string(y), "isConstant is set to true")
				//	fmt.Println(count, "isConstant")
			} else {
				fmt.Println(string(y), "is  not a constant", "this is the count", count)
				isConstant = false
				switch count {
				case 1:
					//total += min
					total += 1
				case 2:
					//	total += min
					total += 2
				case 3:
					//	total += min
					total += 4
				case 4:
					//	total += min
					total += 8
				case 7:
					total += 64
				}
				count = 0
			}

		}

		if count > 0 {
			switch count {
			case 1:
				//total += min
				total += 1
			case 2:
				//	total += min
				total += 2
			case 3:
				//	total += min
				total += 4
			case 4:
				//	total += min
				total += 8
			default:
				total += 8
			}

		}

		fmt.Println(total)
		isConstant = false
		final += total
		total = 5
		count = 0
	}

	fmt.Println(final)
}
