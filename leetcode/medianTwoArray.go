package main

import (
	"fmt"
	"sort"
)

func main() {
	test1 := []int{1, 2}
	test2 := []int{3, 4, 5}

	//var test3 []int

	for _, x := range test2 {

		test1 = append(test1, x)
	}

	sort.Ints(test1)

	if len(test1)%2 == 1 {
		med := len(test1) / 2
		fmt.Println(test1[med])
		fmt.Println("test1")
		return
	}
	index := len(test1) / 2
	fmt.Println(index)

	var ans float64

	ans = float64((test1[index] + test1[index-1])) / 2.0

	fmt.Println(ans)
	fmt.Println(5 / 2.0)

	//fmt.Println(ans)

	//fmt.Println(med)
	//fmt.Println(test1[med])

}
