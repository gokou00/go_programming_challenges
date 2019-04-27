package main

import "fmt"

func main() {
	n := 91
	count := 0

	total := addFunction(n)
	//	fmt.Println(total)

	if total == n {
		fmt.Println(count)
		return
	} else {
		count++
	}

	for total > 9 {
		total = addFunction(total)
		count++
	}

	fmt.Println(count)
	return

}

func addFunction(number int) int {
	pop := 0
	total := 0
	num := number

	for num > 0 {
		pop = num % 10
		total += pop
		num /= 10
	}

	return total

}
