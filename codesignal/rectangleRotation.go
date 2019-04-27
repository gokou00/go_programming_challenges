package main

import (
	"fmt"
	"math"
)

func main() {
	a := 16
	b := 20

	result := 0

	for x := -(a + 100); x < a+111; x++ {
		for y := -(b + 100); y < b+111; y++ {
			x1 := (float64(x) + float64(y)) / math.Sqrt(2)
			y1 := (float64(y) - float64(x)) / math.Sqrt(2)

			if x1 < float64(a)/2.0 && x1 > float64(-a)/2.0 && y1 < float64(b)/2.0 && y1 > float64(-b)/2.0 {
				result++
			}

		}

	}
	fmt.Println(result)

}
