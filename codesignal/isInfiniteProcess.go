package main

import "fmt"

func main() {
	a := 5
	b := 5
	count := 0

	for a != b {
		count++
		if count > b+a {
			fmt.Println("true")
			return
		}

		a++
		b--

	}

	fmt.Println(count)
	fmt.Println("false")

}
