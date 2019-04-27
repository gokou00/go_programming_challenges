package main

import "fmt"

func main() {

	a := []int{6, 2, 7, 3, 1, 0, 4, 5}

	if len(a) == 1 {
		a[0] = -1
		fmt.Println(a)
		return
	}

	for i := 0; i < len(a); i++ {
		a[i] = findLargest(a[i], a[i+1:])
	}

	fmt.Println(a)

}

func findLargest(num int, numArray []int) int {
	for _, x := range numArray {
		if x > num {
			return x
		}
	}

	return -1

}
