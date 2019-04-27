package main

import "fmt"

func main() {
	divisors := []int{6, 2, 2, 8, 9, 2, 2, 2, 2}
	k := 10
	//var numArr []int

	numMap := make(map[string]int)
	numMap2 := make(map[int]string)

	for i := 1; i <= k; i++ {
		//numArr = append(numArr,i)
		numMap2[i] = ""
	}

	for key, _ := range numMap2 {
		for _, m := range divisors {

			if key%m == 0 {
				numMap2[key] += fmt.Sprintf("%d is a div", m)
			} else {
				numMap2[key] += fmt.Sprintf("%d is not a div", m)
			}
		}

	}

	for _, val := range numMap2 {
		numMap[val]++

	}

	fmt.Println(len(numMap))

}
