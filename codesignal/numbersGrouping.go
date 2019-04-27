package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {

	a := []int{102382103, 21039898, 39823, 433, 30928398, 40283209, 23234, 342534, 98473483, 498398424, 9384984, 9839239}

	var total float64

	sort.Ints(a)

	var left float64
	var right float64

	var leftArr []float64
	var rightArr []float64

	for _, x := range a {

		temp := (x - 1) / 10000
		temp1 := float64(temp)

		left = math.Floor(temp1) * 10000
		right = left + 10000

		leftArr = append(leftArr, left)
		rightArr = append(rightArr, right)

	}

	testMap := make(map[float64]int)

	for i := 0; i < len(leftArr); i++ {
		total = leftArr[i] + rightArr[i]

		testMap[total]++

	}

	//fmt.Println(testMap)
	//fmt.Println(total)

	//fmt.Println(leftArr)
	//fmt.Println(rightArr)

	fmt.Println(len(testMap) + len(a))

}
