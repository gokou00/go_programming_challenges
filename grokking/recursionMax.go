package main

import (
	"fmt"
	"sort"
)

func MaxNum(a []int) int {
	//count := -1
	if len(a) == 1 {
		//temp := 0
		//temp = a[0]
		return a[0]
	} else {
		//sort.Ints(a)
		sort.Sort(sort.Reverse(sort.IntSlice(a)))
		a = a[0 : len(a)-1]

		return MaxNum(a)

	}
}

func main() {
	fmt.Println(MaxNum([]int{2}))
}
