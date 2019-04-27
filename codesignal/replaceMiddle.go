package main

import "fmt"

func main() {
	arr := []int{-5, -5, 10}
	idx := 0
	mid := 0
	var finalArr []int

	arrLen := len(arr)
	isEven := false

	if arrLen%2 == 0 {
		idx = arrLen / 2
		mid = arr[idx] + arr[idx-1]
		isEven = true

	} else {
		fmt.Println(arr)
		return
	}

	if isEven {
		for i := 0; i < len(arr); i++ {
			if i == idx-1 {
				finalArr = append(finalArr, mid)
				continue
			}

			if i == idx {
				continue
			}

			finalArr = append(finalArr, arr[i])

		}
	}

	fmt.Println(finalArr)

}
