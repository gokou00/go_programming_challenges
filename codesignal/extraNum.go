package main

import "fmt"

func main() {
	a := 500000000
	b := 3
	c := 500000000

	if a == b {
		fmt.Println(c)
		return
	}

	if a == c {
		fmt.Println(b)
		return
	}

	if b == c {
		fmt.Println(a)
		return
	}

	fmt.Println(-1)

}
