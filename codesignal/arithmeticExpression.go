package main

import "fmt"

func main() {
	a := 8
	b := 2
	c := 4

	if a+b == c {
		fmt.Println("true")
		return
	} else if a-b == c {
		fmt.Println("true")
		return
	} else if a*b == c {
		fmt.Println("true")
		return
	}

	if b != 0 {
		if float32(a)/float32(b) == float32(c) {
			fmt.Println("true: div")
			fmt.Println(float32(a) / float32(b))
			return
		}
	}

	fmt.Println("false")

}
