package main

import (
	"fmt"
	"sort"
)

func main() {
	n := 99999

	var testArr []int

	freqMap := make(map[int]int)

	for n > 0 {

		temp1, temp2 := testFunc(n)

		n = n - temp1

		//fmt.Println(n)

		testArr = append(testArr, temp2)

	}

	//fmt.Println(testArr)

	for _, x := range testArr {
		freqMap[x]++

	}

	//fmt.Println(freqMap)

	largest := 0
	mostFreq := 0

	sort.Ints(testArr)

	for _, x := range testArr {

		if freqMap[x] >= largest {
			largest = freqMap[x]
			mostFreq = x

		}

	}

	fmt.Println(mostFreq)
}

func testFunc(num int) (int, int) {

	total := 0
	pop := 0

	for num > 0 {
		pop = num % 10
		total += pop
		num /= 10

	}

	//fmt.Println(num)

	return total, total

}
