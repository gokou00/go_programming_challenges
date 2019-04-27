package main

import "fmt"

func main() {
	a, b, c := 5, 1, 2
	isTriangle := true

	if a+b < c {
		isTriangle = false
	}

	if !isTriangle {
		fmt.Println("Not a triangle")
	}

	if a+c < b {
		isTriangle = false
	}

	if !isTriangle {
		fmt.Println("Not a triangle")
	}

	if b+c < a {
		isTriangle = false
	}

	if !isTriangle {
		fmt.Println("Not a triangle")
	} else {
		fmt.Println("is a triangle")
	}

}
