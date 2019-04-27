package main

import "fmt"

func main() {
	left := 1
	right := 22

	var arr []int
	isOk := true

	for i := left; i <= right; i++ {
		num := i
		for num != 0 {
			pop := num % 10
			if pop == 0 {
				isOk = false
				break
			}
			if i%pop != 0 {
				isOk = false
				break
			}
			num /= 10
		}

		if isOk {
			arr = append(arr, i)

		} else {
			isOk = true
		}

	}

	fmt.Println(arr)

}
