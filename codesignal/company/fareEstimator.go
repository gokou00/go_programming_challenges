package main

import (
	"fmt"
	"math"
)

func main() {

	ride_time := 30
	ride_distance := 7

	cost_per_minute := []float64{.2, .35, .4, .45}
	cost_per_mile := []float64{1.1, 1.8, 2.3, 3.5}

	var finalarr []float64

	for i := 0; i < len(cost_per_mile); i++ {
		temp := cost_per_minute[i]*float64(ride_time) + cost_per_mile[i]*float64(ride_distance)
		temp = math.Ceil(temp) // should be temp = math.Round(temp)

		finalarr = append(finalarr, temp)

	}

	fmt.Println(finalarr)

}
