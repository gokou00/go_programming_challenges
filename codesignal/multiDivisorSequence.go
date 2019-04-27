package main

import "fmt"

func main() {

	isDivior := false
	isMultiple := false
	failed := 0

	nums := []int{1, 1, 3}

	for i := len(nums) - 1; i >= 1; i-- {
		//divisor test

		if nums[i-1]%nums[i] == 0 {
			isDivior = true

		}

		if i-2 >= 0 {
			if nums[i]%nums[i-2] == 0 || nums[i]%nums[i-2] == nums[i-2] {
				isMultiple = true
			}

		}

		if isDivior || isMultiple {
			isDivior = false
			isMultiple = false
		} else {
			failed = nums[i]
		}

	}

	if failed != 0 {
		fmt.Println(failed)
		return
	} else {
		fmt.Println(-1)
	}

}
