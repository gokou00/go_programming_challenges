package main

import "fmt"

func main() {

	k := 3
	nums := []int{0, 1, 2, 3, 5, 2}
	/*
		var nums []int

		for i := -24500; i <= 29999; i++ {
			nums = append(nums, i)
		}
	*/

	hasTwo := false
	dup := 0

	var indexOfDup []int

	numOfNums := make(map[int]int)
	posOfNums := make(map[int]int)

	for _, x := range nums {
		numOfNums[x]++
	}

	if len(nums) == len(numOfNums) {
		fmt.Println("no dups")
		return
	}

	for key, x := range numOfNums {

		if x >= 2 {
			hasTwo = true
			dup = key
		}
	}

	if !hasTwo {
		fmt.Println("array does not have 2 of the same")
		return
	}

	for i, x := range nums {
		posOfNums[x] = i
	}

	lastPos := posOfNums[dup]

	fmt.Println(lastPos, "lastPos")

	//	firstPos := 0

	for i := 0; i < lastPos; i++ {
		if dup == nums[i] {
			indexOfDup = append(indexOfDup, i)
			//break
		}

	}

	indexOfDup = append(indexOfDup, lastPos)

	//fmt.Println(indexOfDup)

	for i := len(indexOfDup) - 1; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			sub := indexOfDup[i] - indexOfDup[j]

			if sub <= k {
				fmt.Println("true")
				return
			}
		}
	}

	/*
		sub := lastPos - firstPos

		if sub <= k {
			fmt.Println("true")
			return
		}
	*/

	fmt.Println("false")
	return

}
