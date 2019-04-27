package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{-1000000, -10000, -10000, -1000, -100, -10, -1, 0, 1, 10, 100, 1000, 10000, 100000, 1000000}
	var test []int
	//	var sortedTest []int

	// test := make(map[int]int)

	ans := 0

	if len(a) == 1 {
		fmt.Println(a[0])
		return
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			ans += abs(a[j] - a[i])
		}

		test = append(test, ans)
		ans = 0
	}

	//fmt.Println(test)

	sortedTest := make([]int, len(test))

	copy(sortedTest, test)

	sort.Ints(sortedTest)

	fmt.Println("Test", test)
	fmt.Println("sortedTest", sortedTest)

	final := position(test, sortedTest[0])

	fmt.Println(a[final])

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func position(a []int, b int) int {

	for i, x := range a {
		if x == b {
			return i
		}

	}

	return -1

}
