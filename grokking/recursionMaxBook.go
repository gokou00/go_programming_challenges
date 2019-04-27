package main

import "fmt"

func max(a []int) int {
	if len(a) == 2 {
		if a[0] > a[1] {
			return a[0]
		} else {
			return a[1]
		}
	} else {
		subMax := max(a[1:])
		fmt.Println(subMax)
		fmt.Println(a)
		if a[0] > subMax {
			return a[0]
		} else {
			return subMax
		}
	}
}

func main() {
	fmt.Println(max([]int{21, 6, 5, 10, 20}))
}
