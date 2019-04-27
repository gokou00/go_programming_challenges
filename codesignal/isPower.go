package main

import (
	"fmt"
	"math"
)

func main() {
	n := 361

	for i := 1; i <= 100; i++ {
		for j := 2; j <= 100; j++ {
			temp := math.Pow(float64(i), float64(j))
			if n == int(temp) {
				fmt.Println("true")
				return
			}
		}
	}

	fmt.Println("false")

}
