package main

import "fmt"

func main() {
	n := 8

	var sumArry []int

	total := 0

	for i := 1; i < n; i++ {
		for j := i; j < n; j++ {
			if total == n {
				sumArry = append(sumArry, total)
				total = 0
				break
			}

			total += j
			//fmt.Println(total)

		}

		total = 0
	}

	fmt.Println(len(sumArry))

}
