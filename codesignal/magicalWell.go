package main

import "fmt"

func main() {
	a := 1275
	b := 362
	n := 2
	total := 0

	for i := 0; i < n; i++ {
		total += a * b
		a++
		b++

	}

	fmt.Println(total)

}
