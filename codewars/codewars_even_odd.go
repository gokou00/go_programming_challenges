package main

import "fmt"

func main() {
	input := 0
	fmt.Println("Please enter a number")
	fmt.Scanf("%d", &input)
	fmt.Println(input)

	if input%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

}
