package main

import "fmt"

func main() {
	value1 := 3
	weight1 := 5
	value2 := 3
	weight2 := 8
	maxW := 10

	totalWeight := weight1 + weight2

	if totalWeight <= maxW {
		fmt.Println(value1 + value2)
		return
	}

	if value1 >= value2 && weight1 <= maxW {
		fmt.Println(value1)
		return
	}

	if value2 >= value1 && weight2 <= maxW {
		fmt.Println(value2)
		return
	}

	if weight1 <= maxW {
		fmt.Println(value1)
		return
	}

	if weight2 <= maxW {
		fmt.Println(value2)
		return
	}

	fmt.Println("0")

}
