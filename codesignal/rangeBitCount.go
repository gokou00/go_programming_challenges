package main

import (
	"fmt"
	"math/bits"
)

func main() {
	a := 2
	b := 7
	total := 0

	for i := a; i <= b; i++ {
		total += bits.OnesCount16(i)
	}

	fmt.Println(total)

}
