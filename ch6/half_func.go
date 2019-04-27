package main

import "fmt"

func half(x int) (int, bool) {
	numAns := x / 2

	if x%2 == 0 {
		return numAns, true
	} else {
		return numAns, false
	}
}

func main() {
	x, ok := half(1)
	fmt.Println(x, ok)
}
