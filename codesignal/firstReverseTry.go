package main

import (
	"fmt"
	//"sort"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	//var revArr []int

	/*

		for i := len(arr) - 1; i >= 0; i-- {
			revArr = append(revArr, arr[i])
		}

		fmt.Println(revArr)

	*/

	temp := arr[0]

	arr[0] = arr[len(arr)-1]

	arr[len(arr)-1] = temp

	fmt.Println(arr)

}
