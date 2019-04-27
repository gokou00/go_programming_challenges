package main

import "fmt"

func main() {

	arr := []int{-5, 3, -1, 9}

	mid := 0

	idx := 0

	idx = len(arr) / 2

	//fmt.Println(mid)
	if len(arr)%2 == 0 {
		mid = arr[idx-1] + arr[idx]
	} else {
		mid = arr[idx]
	}

	if arr[0] == mid && mid == arr[len(arr)-1] {
		fmt.Println("true")
		return
	} else {
		fmt.Println("false")
		return
	}

}
