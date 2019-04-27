package main

import "fmt"

//https://stackoverflow.com/questions/31891022/maximum-sum-of-n-consecutive-elements-of-array
//https://www.geeksforgeeks.org/largest-sum-contiguous-subarray/
//Kadane’s Algorithm
func main() {

	test := []int{2, 3, 5, 1, 6}

	largestSum := 0
	previousSum := 0

	k := 2

	for i := 0; i <= len(test)-k; i++ {
		if i == 0 {
			for j := 0; j < k; j++ {
				largestSum += test[j]
			}
			previousSum = largestSum
		} else {
			currentSum := previousSum - test[i-1] + test[i+k-1]
			if currentSum > largestSum {
				largestSum = currentSum
			}
			previousSum = currentSum
		}
	}

	fmt.Println(largestSum)

}
